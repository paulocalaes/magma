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
import ActionTable from '../../components/ActionTable';
import Link from '@material-ui/core/Link';
import React from 'react';
import SubscriberContext from '../../components/context/SubscriberContext';
import nullthrows from '../../../shared/util/nullthrows';
import type {GatewayDetailType} from './GatewayDetailMain';

import {
  REFRESH_INTERVAL,
  useRefreshingContext,
} from '../../components/context/RefreshContext';
import {Subscriber} from '../../../generated-ts';
import {useContext} from 'react';
import {useNavigate, useParams, useResolvedPath} from 'react-router-dom';

type SubscriberRowType = {
  id: string;
  service: string;
};

export default function GatewayDetailSubscribers(props: GatewayDetailType) {
  const resolvedPath = useResolvedPath('');
  const navigate = useNavigate();
  const params = useParams();
  const networkId: string = nullthrows(params.networkId);
  // Auto refresh  every 30 seconds
  const subscriberState = useRefreshingContext({
    context: SubscriberContext,
    networkId: networkId,
    type: 'subscriber',
    interval: REFRESH_INTERVAL,
    refresh: props.refresh,
  });
  const subscriberCtx = useContext(SubscriberContext);
  const hardware_id = props.gwInfo?.device?.hardware_id;
  const gwSubscriberMap = hardware_id
    ? subscriberCtx.gwSubscriberMap[hardware_id] || []
    : [];

  const subscriberRows: Array<SubscriberRowType> = gwSubscriberMap.map(
    (serialNum: string) => {
      //TODO[TS-migration] Something is seriously wrong here?
      // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment,@typescript-eslint/no-unsafe-member-access
      const subscriberInfo = (subscriberState as any).state?.[
        serialNum
      ] as Subscriber;
      return {
        name: subscriberInfo?.name || serialNum,
        id: serialNum,
        service: subscriberInfo?.lte.state || '-',
      };
    },
  );

  return (
    <ActionTable
      title=""
      data={subscriberRows}
      columns={[
        {title: 'Name', field: 'name'},
        {
          title: 'Subscriber ID',
          field: 'id',
          render: currRow => (
            <Link
              variant="body2"
              component="button"
              onClick={() => {
                navigate(
                  resolvedPath.pathname.replace(
                    `equipment/overview/gateway/${props.gwInfo.id}`,
                    `subscribers/overview/${currRow.id}`,
                  ),
                );
              }}>
              {currRow.id}
            </Link>
          ),
        },
        {title: 'Service', field: 'service'},
      ]}
      options={{
        actionsColumnIndex: -1,
        pageSizeOptions: [10],
        toolbar: false,
      }}
    />
  );
}
