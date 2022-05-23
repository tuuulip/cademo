<template>
  <div class="cred">
    <div class="cred-header"></div>
    <div class="cred-body">
      <el-table :data="credentials" border>
        <el-table-column prop="id" label="id"></el-table-column>
        <el-table-column prop="revocationHandle" label="revocationHandle">
          <template slot-scope="{ row }">
            <span class="cred-revo">{{ row.revocationHandle }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="cred" label="cred">
          <div class="cred-cred-item" slot-scope="{ row }">
            <span class="cred-cred">{{ row.cred }}</span>
            <el-popover trigger="hover" placement="top-start" width="500">
              <div class="cred-cred__all">{{ row.cred }}</div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
        <el-table-column prop="status" label="status"></el-table-column>
        <el-table-column prop="expiry" label="expiry"></el-table-column>
        <el-table-column label="operation">
          <template slot-scope="{ row }">
            <el-button
              type="text"
              class="cred-del"
              @click="delCredential(row.revocationHandle)"
              >Delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      credentials: []
    };
  },
  created() {
    this.fetchCredentials();
  },
  methods: {
    fetchCredentials() {
      this.$request
        .get("/credential/list")
        .then(res => {
          this.credentials = res.data;
        })
        .catch(() => {});
    },
    delCredential(rh) {
      this.$request
        .post("/credential/del", { revocationHandle: rh })
        .then(() => {
          return this.fetchCredentials();
        })
        .catch(() => {});
    }
  }
};
</script>

<style lang="postcss" scoped>
.cred-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 0 20px;
  box-sizing: border-box;
}

.cred-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.cred-revo {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.cred-cred {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.cred-cred-item {
  display: flex;
}

.cred-cred__all {
  font-weight: 100;
  color: #909399;
  font-size: 14px;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.cred-del {
  color: red;
}
</style>
