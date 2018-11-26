import m from 'mithril';

/**
 * Nav component provides a navigation bar for the top of the page.
 */
export default class Nav implements ClassComponent {
  /**
   * Creates a navigation bar.
   * @param {CVnode} vnode - Virtual node.
   * @return {CVnode}
   */
  view(vnode: CVnode) {
    return m('nav.ui.massive.borderless.stackable.menu[style=margin-top: 0;]', [
      m(
          '.ui.container',
          m(
              'a.header.item',
              {href: '/', oncreate: m.route.link},
              m('i.coffee.icon.logo'),
              'Roaster'
          ),
          m('a.item', {href: '/about', oncreate: m.route.link}, 'About'),
          m('a.item', {href: '/register', oncreate: m.route.link}, 'REGISTER'),
          m('a.item', {href: '/profile', oncreate: m.route.link}, 'Profile'),
          m(
              'a.item',
              {href: '/statistics', oncreate: m.route.link},
              'Statistics'
          )
      ),
    ]);
  }
}
