import { ButtonHTMLAttributes, FC } from "react";

type ButtonProps = {
  onClick?: () => void;
  label: string;
} & ButtonHTMLAttributes<HTMLButtonElement>;

const Button: FC<ButtonProps> = ({ onClick, label, ...buttonProps }) => (
  <button
    {...buttonProps}
    className="rounded-md border border-solid bg-transparent px-3 py-2 transition-colors hover:border-emerald-800 hover:bg-emerald-800 hover:text-white"
    onClick={onClick}
  >
    {label}
  </button>
);

export default Button;
