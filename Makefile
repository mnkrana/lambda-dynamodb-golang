commit:
	git add .
	git commit -m "$m"
	git push origin main
	git tag "$n"
	git push --tag