import { type RouteRecordRaw } from 'vue-router';
import { t } from '@/hooks/useI18n';

const moduleName = 'build';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/build',
    name: moduleName,
    component: () => import('@/views/build/index.vue'),
    meta: {
      title: t('routes.build.build'),
      icon: 'icon-yaml',
    },
  },
];

export default routes;
