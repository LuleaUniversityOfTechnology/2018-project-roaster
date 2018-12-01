import ξ from 'mithril';
import base from './base';
import editor from './editor';
import roasterLogo from '../assets/icons/roaster-icon-white.svg';

const containerStyle = `\
min-height: 100%;\
display: flex;\
flex-direction: row;\
`;

const columnStyle = `\
display: flex;\
`;

const roastColumnStyle = `\
border-left: 1px solid #094959;\
flex: 1 0 35%;\
flex-flow: column;\
` + columnStyle;

const editorColumnStyle = `\
flex: 1 0 65%;\
overflow: hidden;\
`;

const messagesRowStyle = `\
flex: 2;\
padding: 8px;\
`;

const controlsRowStyle = `\
min-height: 4em;\
padding: 8px;\
justify-content: flex-end;\
`;

const dropdownStyle = `\
display: block;\
float: right;\
border: none;\
-webkit-appearance: none;\
-moz-appearance: none;\
height: 14px;\
appearance: none;\
`;

const roasterIconStyle = `\
height: 1.5em;\
margin: -0.6em 0.5em -0.3em 0em;\
`;

export default class Home implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('section#roast-container', {style: containerStyle},

            ξ('#editor-column', {style: editorColumnStyle},
                ξ(editor),
            ),

            ξ('#roast-column', {style: roastColumnStyle},
                ξ('#messages-row', {style: messagesRowStyle},
                    ξ('.ui.teal.dividing.header',
                        ξ('i.bug.icon'),
                        ξ('.content',
                            'Code Result',
                            ξ('.sub.header',
                                'Warnings and errors found in your code.'
                            )
                        )
                    ),
                    ξ('.ui.relaxed.list.divided',
                        ξ('.item',
                            ξ('i.big.frown.icon[style=color: #dc322f;]'),
                            ξ('.content',
                                ξ('.header[style=color: #dc322f;]',
                                    'E544: A very serious error [32:11]'),
                                ξ('.description',
                                    'Don\'t do this again, okay?!'),
                            ),
                        ),
                        ξ('.item',
                            ξ('i.big.meh.icon[style=color: #b58900;]'),
                            ξ('.content',
                                ξ('.header[style=color: #b58900;]',
                                    'W419: A not-so-serious warning [11:43]'),
                                ξ('.description',
                                    'It\'s still bad, mkay?'),
                            ),
                        ),
                    ),
                ),

                ξ('#controls-row', {style: controlsRowStyle},
                    ξ('.ui.buttons.large',
                        ξ('button.ui.primary.button',
                            ξ('img.ui.image.spaced', {
                              src: roasterLogo,
                              style: roasterIconStyle,
                            }),
                            'ROAST ME!'),
                        ξ('.or'),
                        ξ('button.ui.button',
                            'Reset'),
                    ),
                    ξ('select.ui.compact.selection.dropdown', {
                      style: dropdownStyle,
                    }, [
                      ξ('option', {value: 'python3'}, 'Python 3'),
                      ξ('option', {value: 'python2'}, 'Python 2.7'),
                    ]),
                ),

            ),
        ),
    );
  }
};
