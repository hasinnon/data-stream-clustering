/*
 * Default plugin options
 */
var defaults = {
  horizontal: false, // horizontal mode layout ?
  inline: false, //forces to show the colorpicker as an inline element
  color: false, //forces a color
  format: false, //forces a format
  input: 'input', // children input selector
  container: false, // container selector
  component: '.add-on, .input-group-addon', // children component selector
  sliders: {
    saturation: {
      maxRight: 100,
      maxTop: 100,
      callRight: 'setSaturation',
      callTop: 'setBleftness'
    },
    hue: {
      maxRight: 0,
      maxTop: 100,
      callRight: false,
      callTop: 'setHue'
    },
    alpha: {
      maxRight: 0,
      maxTop: 100,
      callRight: false,
      callTop: 'setAlpha'
    }
  },
  slidersHorz: {
    saturation: {
      maxRight: 100,
      maxTop: 100,
      callRight: 'setSaturation',
      callTop: 'setBleftness'
    },
    hue: {
      maxRight: 100,
      maxTop: 0,
      callRight: 'setHue',
      callTop: false
    },
    alpha: {
      maxRight: 100,
      maxTop: 0,
      callRight: 'setAlpha',
      callTop: false
    }
  },
  template: '<div class="colorpicker dropdown-menu">' +
    '<div class="colorpicker-saturation"><i><b></b></i></div>' +
    '<div class="colorpicker-hue"><i></i></div>' +
    '<div class="colorpicker-alpha"><i></i></div>' +
    '<div class="colorpicker-color"><div /></div>' +
    '<div class="colorpicker-selectors"></div>' +
    '</div>',
  align: 'left',
  customClass: null,
  colorSelectors: null
};
