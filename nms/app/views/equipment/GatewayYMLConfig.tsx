/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import Button from '@material-ui/core/Button';
import FormLabel from '@material-ui/core/FormLabel';
import Grid from '@material-ui/core/Grid';
import ListItemText from '@material-ui/core/ListItemText';
import LoadingFiller from '../../components/LoadingFiller';
import MagmaAPI from '../../api/MagmaAPI';
import MenuItem from '@material-ui/core/MenuItem';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import React from 'react';
import Select from '@material-ui/core/Select';
import Text from '../../theme/design-system/Text';
import TextareaAutosize from 'react-textarea-autosize';
import nullthrows from '../../../shared/util/nullthrows';
import useMagmaAPI from '../../api/useMagmaAPI';
import withAlert from '../../components/Alert/withAlert';
import {AltFormField} from '../../components/FormField';
import {GenericCommandResponse} from '../../../generated-ts';
import {RUNNING_SERVICES} from '../../components/GatewayUtils';
import {Theme} from '@material-ui/core/styles';
import {getErrorMessage} from '../../util/ErrorUtils';
import {makeStyles} from '@material-ui/styles';
import {useCallback, useState} from 'react';
import {useEnqueueSnackbar} from '../../hooks/useSnackbar';
import {useParams} from 'react-router-dom';
import {useSnackbar} from '../../hooks';
import type {WithAlert} from '../../components/Alert/withAlert';

const useStyles = makeStyles<Theme>(theme => ({
  dashboardRoot: {
    margin: theme.spacing(3),
    flexGrow: 1,
  },
  ymlEditor: {
    minWidth: '80%',
    padding: '10px',
  },
  description: {
    marginRight: '15px',
  },
}));

type Props = WithAlert;
function GatewayConfigYml(props: Props) {
  const classes = useStyles();
  const params = useParams();
  const enqueueSnackbar = useEnqueueSnackbar();
  const networkId = nullthrows(params.networkId);
  const gatewayId: string = nullthrows(params.gatewayId);
  const [selectedService, setSelectedService] = useState(RUNNING_SERVICES[0]);
  const [serviceConfig, setServiceConfig] = useState<string>('');
  const SERVICE_CONTENT = `cat /etc/magma/${selectedService}.yml`;
  const serviceParams = {
    command: 'bash',
    params: {
      shell_params: [`-c '${SERVICE_CONTENT}'`],
    },
  };

  const onSave = () => {
    const config = serviceConfig.replace(/\\/g, '');
    void props
      .confirm('Are you sure you want to save this config?')
      .then(async confirm => {
        if (!confirm) return;
        try {
          await MagmaAPI.commands.networksNetworkIdGatewaysGatewayIdCommandGenericPost(
            {
              networkId,
              gatewayId,
              parameters: {
                command: 'bash',
                params: {
                  shell_params: [
                    `-c "cat >/etc/magma/${selectedService}.yml <<EOL \n${config} \nEOL"`,
                  ],
                },
              },
            },
          );
        } catch (e) {
          enqueueSnackbar(getErrorMessage(e), {
            variant: 'error',
          });
        }
      });
  };

  // fetch service if selectedService changes
  const {
    isLoading: serviceConfigLoading,
    error: serviceConfigError,
  } = useMagmaAPI(
    MagmaAPI.commands.networksNetworkIdGatewaysGatewayIdCommandGenericPost,
    {networkId, gatewayId, parameters: serviceParams},
    useCallback(
      (response: GenericCommandResponse) => {
        enqueueSnackbar(`${selectedService} config fetched successfully`, {
          variant: 'success',
        });
        // TODO[TS-migration] What is the real response here?
        // @ts-ignore
        setServiceConfig(response?.response?.['stdout'] ?? '');
      },
      [selectedService, enqueueSnackbar],
    ),
  );

  useSnackbar(
    `Error fetching ${selectedService} config`,
    {variant: 'error'},
    !!serviceConfigError,
  );

  if (serviceConfigLoading) {
    return <LoadingFiller />;
  }

  return (
    <div className={classes.dashboardRoot}>
      <Grid container direction="column" alignItems="center" spacing={4}>
        <Grid item xs={12}>
          <Text
            weight="medium"
            variant="subtitle2"
            className={classes.description}>
            {'Select a service:'}
          </Text>
          <Select
            variant={'outlined'}
            displayEmpty={true}
            value={selectedService}
            onChange={({target}) => setSelectedService(target.value as string)}
            input={<OutlinedInput />}>
            {RUNNING_SERVICES.map(service => (
              <MenuItem key={service} value={service}>
                <ListItemText primary={service} />
              </MenuItem>
            ))}
          </Select>
        </Grid>
        {serviceConfigError && (
          <Grid item xs={12}>
            <AltFormField label={''}>
              <FormLabel
                error>{`Error fetching ${selectedService} config`}</FormLabel>
            </AltFormField>
          </Grid>
        )}
        <TextareaAutosize
          disabled={!!serviceConfigError}
          value={serviceConfig}
          className={classes.ymlEditor}
          onChange={event => {
            setServiceConfig(event.target.value);
          }}
        />
        <Grid item xs={12}>
          <Button
            disabled={!!serviceConfigError}
            variant="contained"
            onClick={onSave}
            color="primary">
            {'Save'}
          </Button>
        </Grid>
      </Grid>
    </div>
  );
}
export default withAlert(GatewayConfigYml);
