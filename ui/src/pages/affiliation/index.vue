<template>
  <div class="aff">
    <div class="aff-header">
      <el-button type="primary">Add</el-button>
    </div>
    <div class="aff-body">
      <el-table :data="affilliationList">
        <el-table-column prop="name" label="name"></el-table-column>
      </el-table>
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
      affiliationInfo: {},
      affilliationList: []
    };
  },
  created() {
    this.getAllAffiliations();
  },
  methods: {
    getAllAffiliations() {
      this.$request.get("/affi/all").then(res => {
        this.affiliationInfo = res.data;
        const affilliationList = [];
        this.parseAffiliation(this.affiliationInfo, affilliationList);
        this.affilliationList = affilliationList;
      });
    },
    parseAffiliation(affInfo, affList) {
      if (!affInfo) return;
      if (affInfo.name) {
        affList.push({ name: affInfo.name });
      }
      if (!affInfo.affiliations) return;
      affInfo.affiliations.forEach(element => {
        this.parseAffiliation(element, affList);
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
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 0 20px;
  box-sizing: border-box;
}

.aff-body {
  padding: 0 20px;
  text-align: left;
}
</style>
