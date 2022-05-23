<template>
  <div class="cert">
    <div class="cert-header">
      <div class="cert-search">
        <el-input placeholder="Please input id" v-model="search.id">
          <template slot="prepend">Id</template>
        </el-input>
        <el-input placeholder="Please input serial" v-model="search.serial">
          <template slot="prepend">Serial</template>
        </el-input>
        <el-button @click="reset">Reset</el-button>
        <el-button type="primary" @click="fetchCertificates">Search</el-button>
      </div>
      <el-button type="primary" @click="showEnroll">Enroll</el-button>
    </div>
    <div class="cert-body">
      <el-table :data="certificates" border height="80vh" class="cert-table">
        <el-table-column prop="id" label="id" width="100px" />
        <el-table-column prop="serialNumber" label="serial number" />
        <el-table-column prop="aki" label="Authority Key Id" />
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
        <el-table-column prop="notBefore" label="not before" />
        <el-table-column prop="notAfter" label="not after" />
        <el-table-column label="operation">
          <template slot-scope="{ row }">
            <el-button type="text" @click="requestReenroll(row.id)"
              >Reenroll</el-button
            >
            <el-button
              type="text"
              class="cert-btn-del"
              @click="deleteCertificate(row)"
              >Delete</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- enroll dialog -->
    <Enroll ref="enroll" @enroll="requestEnroll" />
  </div>
</template>

<script>
import Enroll from "./__Enroll_.vue";
export default {
  components: { Enroll },
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
    this.fetchCertificates();
  },
  methods: {
    fetchCertificates() {
      this.$request
        .post("/cert/list", this.search)
        .then(res => {
          this.certificates = res.data;
        })
        .catch(() => {});
    },
    requestEnroll(postData) {
      this.$request
        .post("/cert/enroll", postData)
        .then(() => {
          this.hideEnroll();
          return this.fetchCertificates();
        })
        .catch(() => {});
    },
    requestReenroll(id) {
      this.$confirm("Certificate will be reenroll. Continuer?", "Warning")
        .then(() => {
          return this.$request.post("/cert/reenroll", { user: id });
        })
        .then(() => {
          this.$notify({
            title: "Success",
            message: "Reenroll success",
            type: "success"
          });
        })
        .catch(() => {});
    },
    showEnroll() {
      this.$refs["enroll"].show();
    },
    hideEnroll() {
      this.$refs["enroll"].hide();
    },
    reset() {
      this.search = {};
      return this.fetchCertificates();
    },
    onDateChange() {
      console.log(this.dateRange);
    },
    deleteCertificate(row) {
      this.$request
        .post("/cert/del", row)
        .then(() => {
          this.fetchCertificates();
        })
        .catch(() => {});
    }
  }
};
</script>

<style lang="postcss" scoped>
.cert-header {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
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

.cert-table {
  width: 99%;
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
  justify-content: flex-start;
}

.cert-search .el-input {
  margin-right: 10px;
}

.cert-btn-del {
  color: red;
}
</style>
