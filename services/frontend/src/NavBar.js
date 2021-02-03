import React from 'react'
import AppBar from '@material-ui/core/AppBar'
import Toolbar from '@material-ui/core/Toolbar'
import Typography from '@material-ui/core/Typography'

import { Link } from "react-router-dom";

const NavBar = () => {
    return(
        <div>
        <AppBar position="static">
            <Toolbar>
                <teste>
                    <Typography variant="title" color="inherit">
                    <h2><Link to="/" style={{fontfamily: "sans-serif"}}>EasyFood</Link></h2>
                    </Typography>
                </teste>
                <teste>
                    <a>
                        <Link to="/restaurant">Restaurantes</Link>
                        <Link to="/category">Categorias</Link>                
                        <Link to="/dish">Pratos</Link>
                    </a>
                </teste>
            </Toolbar>
        </AppBar>
        </div>
    )
}
export default NavBar;
