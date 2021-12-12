import React, { Component } from 'react';

const Menu = () => {
  return (
    <nav
      id="sidebarMenu"
      className="col-md-3 col-lg-2 d-md-block bg-light sidebar cpllapse"
    >
      <div className="position-sticky pt-3">
        <ul className="nav flex-column">
          <li className="nav-item">
            <a className="nav-link active">Bashboard</a>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Menu;