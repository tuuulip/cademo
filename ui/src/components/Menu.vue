<template>
  <div class="menu-container">
    <el-menu class="menu" :default-active="defaultActive" @select="onSelect">
      <el-menu-item
        v-for="(item, i) in menuItems"
        :key="i"
        :index="item.index"
        class="menu-item"
        :class="active(item.index) ? 'menu-item-active' : ''"
      >
        <template slot="title">
          <i :class="item.icon"></i>
          <span>{{ item.title }}</span>
        </template>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  data() {
    return {
      defaultActive: "1",
      menuItems: [
        { index: "identity", icon: "el-icon-user", title: "Identity" },
        {
          index: "certificate",
          icon: "el-icon-bank-card",
          title: "Certificate"
        },
        {
          index: "affiliation",
          icon: "el-icon-office-building",
          title: "Affiliation"
        },
        {
          index: "credential",
          icon: "el-icon-money",
          title: "Credential"
        }
      ]
    };
  },
  methods: {
    active(index) {
      return this.$route.path.includes(index);
    },
    onSelect(index) {
      switch (index) {
        case "identity":
          this.$router.push("/identity").catch(() => {});
          break;
        case "certificate":
          this.$router.push("/certificate").catch(() => {});
          break;
        case "affiliation":
          this.$router.push("/affiliation").catch(() => {});
          break;
        case "credential":
          this.$router.push("/credential").catch(() => {});
          break;
      }
    }
  }
};
</script>

<style lang="postcss" scoped>
.menu {
  background-color: #f3f6fb;
  border: none;
}

.menu-item {
  display: flex;
  align-items: center;
}

.menu-item-active {
  background-color: #fff;
  font-weight: bold;
}
</style>
