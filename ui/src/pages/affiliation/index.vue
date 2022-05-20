<template>
  <div class="aff">
    <div class="aff-header"></div>
    <div class="aff-body">
      <el-tree
        :data="affiliationInfo.affiliations"
        :props="props"
        @node-click="handleNodeClick"
        default-expand-all
      ></el-tree>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      props: {
        label: "name",
        children: "affiliations"
      },
      affiliationInfo: {}
    };
  },
  created() {
    this.getAllAffiliations();
  },
  methods: {
    getAllAffiliations() {
      this.$request.get("/affiliations").then(res => {
        this.affiliationInfo = res.data;
      });
    },
    handleNodeClick() {}
  }
};
</script>

<style lang="postcss" scoped>
.aff {
  text-align: left;
}

.aff-header {
  height: 100px;
}

.aff-body {
  padding: 0 20px;
  text-align: left;
}
</style>
