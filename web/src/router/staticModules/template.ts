import { type RouteRecordRaw } from 'vue-router';
import { t } from '@/hooks/useI18n';

const moduleName = 'template';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/build',
    name: moduleName,
    component: () => import('@/views/template/index.vue'),
    meta: {
      title: t('routes.template.template'),
      icon: 'icon-template-l',
    },
  },
];

export default routes;
