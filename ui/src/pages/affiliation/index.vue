<template>
  <div class="aff">
    <div class="aff-header">
      <el-button type="primary" @click="showDialog">Add Affiliation</el-button>
    </div>
    <div class="aff-body">
      <el-table :data="affilliationList" border>
        <el-table-column prop="name" label="name"></el-table-column>
        <el-table-column label="operation">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              class="aff-del"
              @click="delAffiliation(row.name)"
              >delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <Add ref="add" @add="addAffiliation" />
  </div>
</template>

<script>
import Add from "./__Add_.vue";
export default {
  components: { Add },
  data() {
    return {
      affiliationInfo: {},
      affilliationList: []
    };
  },
  created() {
    this.getAllAffiliations();
  },
  methods: {
    showDialog() {
      this.$refs["add"].show();
    },
    hideDialog() {
      this.$refs["add"].hide();
    },
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
    addAffiliation(postData) {
      this.$request
        .post("/affi/add", postData)
        .then(() => {
          this.hideDialog();
          this.$notify({
            title: "success",
            message: "Add affiliation success.",
            type: "success"
          });
          return this.getAllAffiliations();
        })
        .catch(() => {});
    },
    delAffiliation(name) {
      this.$request
        .post("/affi/del", { name })
        .then(() => {
          this.$notify({
            title: "success",
            message: "Delete affiliation success.",
            type: "success"
          });
        })
        .catch(() => {});
    }
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

.aff-del {
  color: red;
}
</style>
