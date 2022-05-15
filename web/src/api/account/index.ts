import { request, BaseResponse } from '@/utils/request';

export function updateAccountInfo(data: any) {
  return request<BaseResponse<any>>({
    url: 'users/update',
    method: 'post',
    data,
  });
}

export function updatePassword(data: any) {
  return request({
    url: 'users/password',
    method: 'post',
    data,
  });
}

export function getInfo() {
  return request<API.AdminUserInfo>({
    url: 'users/info',
    method: 'get',
  });
}

export function logout() {
  return request({
    url: 'users/logout',
    method: 'post',
  });
}
