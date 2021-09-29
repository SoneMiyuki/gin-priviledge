const telephoneValidator = (value) => /^1[3|4|5|7|9]\d{9}$/.test(value);

export default {
  telephoneValidator,
};
