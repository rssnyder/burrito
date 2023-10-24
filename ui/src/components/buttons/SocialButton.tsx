import React from "react";

import Button from "@/components/buttons/Button";
import GithubIcon from "@/assets/icons/GithubIcon";
import GitlabIcon from "@/assets/icons/GitlabIcon";

export interface SocialButtonProps {
  variant: "github" | "gitlab";
  isLoading?: boolean;
  onClick?: () => void;
}

const SocialButton: React.FC<SocialButtonProps> = ({
  variant,
  isLoading,
  onClick,
}) => {
  const getContent = () => {
    switch (variant) {
      case "github":
        return (
          <>
            <GithubIcon />
            <span>Login with GitHub</span>
          </>
        );
      case "gitlab":
        return (
          <>
            <GitlabIcon />
            <span>Login with Gitlab</span>
          </>
        );
    }
  };

  return (
    <Button
      className="w-[300px]"
      variant="secondary"
      isLoading={isLoading}
      onClick={onClick}
    >
      <div className="flex items-center gap-4 justify-center">
        {getContent()}
      </div>
    </Button>
  );
};

export default SocialButton;
