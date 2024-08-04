import { ComponentType } from "react";

export interface LocationStates {
  "/"?: object;
  "/register"?: object;
  "/login"?: object;
  "/home"?: object;
  "/about"?: object;
  "/help"?: object;
  "/donation-payment/:id"?: object;
}

export type PathName = keyof LocationStates;

export interface Page {
  path: PathName;
  exact?: boolean;
  component: ComponentType<object>;
}
