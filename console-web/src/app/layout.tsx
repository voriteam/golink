import type { Metadata } from "next";
import ThemeRegistry from "./components/ThemeRegistry/ThemeRegistry";
import {
  Link as LinkIcon,
  List as ListIcon,
  Add as AddIcon,
} from "@mui/icons-material";
import {
  AppBar,
  Box,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Typography,
} from "@mui/material";
import zIndex from "@mui/material/styles/zIndex";

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

const drawerWidth = 240;

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <ThemeRegistry>
          <Box sx={{ display: "flex" }}>
            <AppBar position="fixed" sx={{ zIndex: zIndex.drawer + 1 }}>
              <Toolbar>
                <LinkIcon sx={{ mr: 2 }} />
                <Typography variant="h6" noWrap component="div">
                  Golink
                </Typography>
              </Toolbar>
            </AppBar>
            <Drawer
              variant="permanent"
              anchor="left"
              sx={{
                width: drawerWidth,
                flexShrink: 0,
                "& .MuiDrawer-paper": {
                  width: drawerWidth,
                  boxSizing: "border-box",
                },
              }}
            >
              <Toolbar />
              <Divider />
              <List>
                <ListItem disablePadding>
                  <ListItemButton>
                    <ListItemIcon>
                      <AddIcon />
                    </ListItemIcon>
                    <ListItemText primary="Create New Golink" />
                  </ListItemButton>
                </ListItem>
                <ListItem disablePadding>
                  <ListItemButton>
                    <ListItemIcon>
                      <ListIcon />
                    </ListItemIcon>
                    <ListItemText primary="My Golinks" />
                  </ListItemButton>
                </ListItem>
              </List>
            </Drawer>
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
              <Toolbar />
              {children}
            </Box>
          </Box>
        </ThemeRegistry>
      </body>
    </html>
  );
}
