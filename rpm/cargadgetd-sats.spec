%global _missing_build_ids_terminate_build 0

Name:           cargadgetd-plugin-sats
Version:        0.0.1
Release:        1
Summary:        Cargadgetd satellite toasts
Group:          Applications/System
License:        MIT
Source0:        %{name}-%{version}.tar.bz2
BuildRequires:  golang
BuildRequires:  gcc

%description
GPS lock and loss toasts.


%prep
%setup -q


%build
set -eu
export GOTOOLCHAIN=local
export GO111MODULE=on
export GOPROXY=off
export CGO_ENABLED=1
mkdir -p _out
go build -mod=vendor -buildvcs=false -buildmode=c-shared -ldflags "-s -w" \
    -o _out/libcargadgetd_app_sats.so .


%install
set -eu
rm -rf %{buildroot}
install -d %{buildroot}/usr/lib/cargadgetd/plugins
install -m 755 _out/libcargadgetd_app_sats.so %{buildroot}/usr/lib/cargadgetd/plugins/libcargadgetd_app_sats.so


%files
%defattr(-,root,root,-)
%dir /usr/lib/cargadgetd
%dir /usr/lib/cargadgetd/plugins
/usr/lib/cargadgetd/plugins/libcargadgetd_app_sats.so
