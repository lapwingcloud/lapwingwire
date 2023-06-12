export default function Login() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="xs:p-0 mx-auto p-10 md:w-full md:max-w-md">
        <h1 className="mb-5 text-center text-2xl font-bold">Lawpwingwire</h1>
        {/* <div className="w-full divide-y divide-gray-200 rounded-lg bg-white shadow"> */}
        <div className="w-full rounded-lg">
          <div className="flex flex-col space-y-6 px-6 py-8">
            <button className="flex w-full bg-white rounded-lg border border-slate-300 px-4 py-2 text-slate-700 transition duration-150 hover:border-slate-400 hover:text-slate-900 hover:shadow">
              <img
                className="h-6 w-6"
                src="https://www.svgrepo.com/show/475656/google-color.svg"
                loading="lazy"
                alt="google logo"
              />
              <span className="flex-1">Login with Google</span>
              <span className="h-6 w-6"></span>
            </button>
          </div>
        </div>
      </div>
    </main>
  );
}
