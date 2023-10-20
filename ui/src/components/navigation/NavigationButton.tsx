import React from "react";

export interface NavigationButtonProps {
  variant?: "light" | "dark";
  icon: React.ReactNode;
  selected?: boolean;
  onClick?: () => void;
}

const NavigationButton: React.FC<NavigationButtonProps> = ({
  variant = "dark",
  icon,
  selected,
  onClick,
}) => {
  const styles = {
    base: {
      light: `bg-primary-400
        fill-primary-600
        hover:bg-primary-500
        focus:outline-primary-600`,

      dark: `bg-nuances-400
        fill-nuances-300
        hover:bg-nuances-black
        focus:outline-nuances-300`,
    },

    selected: {
      light: `bg-nuances-black
        fill-nuances-white
        focus:outline-nuances-black`,

      dark: `bg-nuances-50
        fill-nuances-black
        focus:outline-nuances-50`,
    },
  };

  return (
    <button
      className={`
        flex
        justify-center
        items-center
        w-8
        h-8
        flex-shrink
        rounded-lg
        focus:outline
        focus:outline-1
        focus:outline-offset-2
        ${selected ? styles.selected[variant] : styles.base[variant]}
      `}
      tabIndex={0}
      onClick={onClick}
    >
      {icon}
    </button>
  );
};

export default NavigationButton;
