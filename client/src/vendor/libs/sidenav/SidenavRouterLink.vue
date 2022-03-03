<template>
  <router-link custom v-slot="{ href, navigate, isActive, isExactActive }"
    :to="to"
    :replace="replace"
    :append="append"
    :exact="exact">
    <div class="sidenav-item" :class="{active: active || isActive || isExactActive, disabled: disabled}">
      <a class="sidenav-link" :class="linkClass" :href="href" @click="navigate" @keypress.enter="navigate">
        <i class="sidenav-icon" v-if="icon" :class="icon"></i>
        <div><slot></slot></div>
        <div v-if="badge" class="pl-1 ml-auto"><div class="badge" :class="badgeClass">{{badge}}</div></div>
      </a>
    </div>
  </router-link>
</template>

<script>
export default {
  name: 'sidenav-router-link',
  props: {
    to: null,
    replace: {
      type: Boolean,
      default: false
    },
    append: {
      type: Boolean,
      default: false
    },
    exact: {
      type: Boolean,
      default: false
    },
    icon: String,
    linkClass: {
      type: String,
      default: ''
    },
    badge: {
      default: null
    },
    badgeClass: {
      type: String,
      default: ''
    },
    disabled: {
      type: Boolean,
      default: false
    },
    active: {
      type: Boolean,
      default: false
    }
  }
}
</script>
