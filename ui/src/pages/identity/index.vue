<template>
  <div class="iden">
    <div class="iden-header">
      <el-button type="primary" @click="showDialog">Add Identity</el-button>
    </div>
    <div class="iden-body">
      <el-table :data="identities" border height="80vh">
        <el-table-column prop="id" label="id" />
        <el-table-column prop="type" label="type" />
        <el-table-column prop="affiliation" label="affiliation" />
        <el-table-column prop="max_enrollments" label="max_enrollments" />
        <el-table-column label="attrs" width="230px">
          <div slot-scope="{ row }" class="iden-column-attr">
            <span class="iden-attr">{{ row.attrs | attrJson }}</span>
            <el-popover trigger="hover" placement="top-start">
              <div>
                <div v-for="(item, i) in row.attrs" :key="i">{{ item }}</div>
              </div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
        <el-table-column label="operation">
          <template>
            <div>
              <el-button type="text">Edit</el-button>
              <el-button type="text" class="iden-btn-del">Delete</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- dialogs -->
    <Add ref="addDialog" @addIdentity="addIdentity" />
  </div>
</template>

<script>
import Add from "./__Add__.vue";
export default {
  components: { Add },
  data() {
    return {
      identities: []
    };
  },
  filters: {
    attrJson(obj) {
      return JSON.stringify(obj);
    }
  },
  created() {
    this.fetchIdentities();
  },
  methods: {
    fetchIdentities() {
      this.$request
        .get("/identities")
        .then(res => {
          this.identities = res.data;
        })
        .catch(() => {});
    },
    showDialog() {
      this.$refs["addDialog"].show();
    },
    hideDialog() {
      this.$refs["addDialog"].hide();
    },
    addIdentity(postData) {
      this.$request
        .post("/register", postData)
        .then(() => {
          this.hideDialog();
          return this.fetchIdentities();
        })
        .catch(() => {});
    }
  }
};
</script>

<style lang="postcss" scoped>
.iden-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 0 20px;
  box-sizing: border-box;
}

.iden-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.iden-title {
  text-align: left;
}

.iden-attr {
  width: 200px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.iden-column-attr {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.iden-btn-del {
  color: red;
}
</style>
