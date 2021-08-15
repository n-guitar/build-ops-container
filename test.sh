echo "# top"
curl http://127.0.0.1:3000/;echo

echo "# get all"
curl http://127.0.0.1:3000/api/v1/build;echo

echo "# form multi"
curl -X POST 'http://127.0.0.1:3000/api/v1/build' \
--form 'buildname="テスト form build"' \
--form 'gitrepo="https://github.com/n-guitar/build-ops-container.git"' \
--form 'imgtag="build-ops-container:form-test"';echo

echo "# get all"
curl http://127.0.0.1:3000/api/v1/build;echo
