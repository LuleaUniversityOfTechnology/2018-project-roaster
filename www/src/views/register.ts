import m, { ClassComponent, CVnode} from "mithril";
import base from "./base";
import user from "../models/user";

function onsubmit() {
    console.log("On submit function called, but not implemented");
    user.create();
}

function validate() {
    console.log(this.value);
}

function input(fieldName: string, type: string) {
    return m(".field",
        m("label", fieldName + ":"),
        m("input", {
            type: type,
            placeholder: fieldName,
            oninput: m.withAttr("value", validate),
            // TODO: inplement on input
        })
    );
}
                        
export default class Register implements ClassComponent {
    view(vnode: CVnode) {
        return m(base, 
            m(".ui.grid",
                m(".ui.container.six.wide.column.centered",
                    m(".ui.segments",
                        m(".ui.segment", m("h2", "REGISTER")),
                        m(".ui.segment", m("form.ui.form",  { onsubmit }, [
                            input("Full name", "text"),
                            input("Username", "text"),
                            input("E-mail", "text"),
                            input("Password", "password"),
                            m("button.ui.teal.basic.button", "GET ROASTED!")
                        ]))
                    )
                )
            )
        );
    }
};
