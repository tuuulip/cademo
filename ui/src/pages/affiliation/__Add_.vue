<template>
  <el-dialog
    class="add"
    title="Add affiliation"
    :visible.sync="dialogVisible"
    width="30%"
    :before-close="hide"
  >
    <el-form class="add-form" ref="form" :model="form" label-width="80px">
      <el-form-item prop="organization" label="Name">
        <el-input
          v-model="form.name"
          placeholder="Please input affiliation name"
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
  props: ["add"],
  data() {
    return {
      dialogVisible: false,
      form: {
        name: ""
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
      if (this.form.name === "") {
        this.$message.warning("Please input affiliation.");
        return;
      }
      this.$emit("add", this.form);
    }
  }
};
</script>

<style lang="postcss" scoped>
.add-form .el-select {
  width: 220px;
}
.add-form .el-input {
  width: 220px;
}
</style>
