import {apiPost} from '@/utils/request'

export function login(data) {
  return apiPost('/auth', data)
}

export function getInfo() {
  return apiPost('/v1/user/info')
}

export function logout() {
  return apiPost('/v1/user/logout')
}
