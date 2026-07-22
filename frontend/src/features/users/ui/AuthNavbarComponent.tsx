
export function AuthNavbarComponent() {
  return (
    <div className="flex h-[65px] items-center justify-center px-5 pt-5 sm:justify-start">
      <div className="flex items-center gap-3">
        <img className="h-[45px] w-[45px] p-1" src="/logo.svg" />
        <div className="text-xl font-bold text-[#ea580c]">
          BackupForge
        </div>
      </div>
    </div>
  );
}
