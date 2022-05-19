<template>
  <div class="cert">
    <div class="cert-header">
      <div class="cert-search">
        <el-input placeholder="Please input id" v-model="search.id">
          <template slot="prepend">Id</template>
        </el-input>
        <el-input
          placeholder="Please input serialNumber"
          v-model="search.serial"
        >
          <template slot="prepend">SerialNumber</template>
        </el-input>
        <el-button @click="reset">Reset</el-button>
        <el-button type="primary" @click="fetchIdentities">Search</el-button>
      </div>
      <el-button type="primary">Create Certificate</el-button>
    </div>
    <div class="cert-body">
      <h3 class="cert-title">Certificates</h3>
      <el-table :data="certificates" border>
        <el-table-column prop="id" label="id" />
        <el-table-column prop="serialNumber" label="serialNumber" />
        <el-table-column prop="pem" label="pem" width="230px">
          <div slot-scope="{ row }" class="cert-column-attr">
            <span class="cert-attr">{{ row.pem }}</span>
            <el-popover trigger="hover" placement="top-start" width="600">
              <div class="cert-pem">{{ row.pem }}</div>
              <i
                slot="reference"
                class="iden-attr-icon el-icon-warning-outline"
              ></i>
            </el-popover>
          </div>
        </el-table-column>
        <el-table-column prop="notBefore" label="notBefore" />
        <el-table-column prop="notAfter" label="notAfter" />
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      certificates: [],
      search: {
        id: "",
        serial: ""
      }
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
        .post("/certificates", this.search)
        .then(res => {
          this.certificates = res.data;
        })
        .catch(() => {});
    },
    reset() {
      this.search = {};
      return this.fetchIdentities();
    }
  }
};
</script>

<style lang="postcss" scoped>
.cert-header {
  width: 100%;
  height: 10vh;
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  padding: 0 20px;
  box-sizing: border-box;
}

.cert-body {
  padding: 0 20px;
  box-sizing: border-box;
}

.cert-title {
  text-align: left;
}

.cert-attr {
  width: 200px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.cert-attr-icon {
  font-size: 18px;
  cursor: pointer;
  color: #409eff;
}

.cert-column-attr {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.cert-pem {
  font-weight: 100;
  color: #909399;
  font-size: 14px;
  white-space: pre;
  font-family: "Avenir", Helvetica, Arial, sans-serif;
}

.cert-search {
  display: flex;
}
</style>
