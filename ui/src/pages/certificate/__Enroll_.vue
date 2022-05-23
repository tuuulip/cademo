<template>
  <el-dialog
    class="enroll"
    title="Enroll"
    :visible.sync="dialogVisible"
    width="30%"
    :before-close="hide"
  >
    <el-form class="enroll-form" ref="form" :model="form" label-width="80px">
      <el-form-item prop="user" label="ID">
        <el-input v-model="form.user" placeholder="Please input id"></el-input>
      </el-form-item>
      <el-form-item prop="password" label="Secret">
        <el-input
          show-password
          v-model="form.password"
          placeholder="Please input password"
        ></el-input>
      </el-form-item>
      <el-form-item prop="type" label="Type">
        <el-select v-model="form.type" placeholder="Please select type">
          <el-option
            v-for="item in typeOptions"
            :key="item"
            :label="item"
            :value="item"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item prop="organization" label="Organ">
        <el-input
          v-model="form.organization"
          placeholder="Please input organization"
        ></el-input>
      </el-form-item>
      <el-form-item prop="domain" label="Domain">
        <el-input
          v-model="form.domain"
          placeholder="Please input domain"
        ></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="hide">Cancel</el-button>
      <el-button type="primary" @click="onConfirm">Confirm</el-button>
    </span>
  </el-dialog>
</template>

<script>
export default {
  props: ["enroll"],
  data() {
    return {
      dialogVisible: false,
      typeOptions: ["client", "orderer", "peer", "user"],
      form: {
        user: "",
        password: "",
        organization: "",
        domain: "",
        type: ""
      }
    };
  },
  methods: {
    show() {
      this.dialogVisible = true;
    },
    hide() {
      this.beforeClose();
      this.dialogVisible = false;
    },
    beforeClose() {
      this.$refs["form"].resetFields();
    },
    onConfirm() {
      this.$emit("enroll", this.form);
    }
  }
};
</script>

<style lang="postcss" scoped>
.enroll-form .el-select {
  width: 220px;
}
.enroll-form .el-input {
  width: 220px;
}
</style>
