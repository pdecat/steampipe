import ErrorMessage from "../ErrorMessage";
import Modal from "./index";
import { ErrorIcon } from "../../constants/icons";
import { useState } from "react";
import NeutralButton from "../forms/NeutralButton";

const ErrorModal = ({ error, title }) => {
  const [show, setShow] = useState(true);
  return show ? (
    <Modal
      actions={[
        <NeutralButton onClick={() => setShow(false)}>
          <>Close</>
        </NeutralButton>,
      ]}
      icon={<ErrorIcon className="h-6 w-6 text-red-600" aria-hidden="true" />}
      children={
        <p className="w-full sm:w-11/12 text-sm text-foreground-light break-words whitespace-pre-wrap">
          <div className="break-all">
            <ErrorMessage error={error} />
          </div>
        </p>
      }
      onClose={async () => {
        setShow(false);
      }}
      title={title}
    />
  ) : null;
};

export default ErrorModal;
