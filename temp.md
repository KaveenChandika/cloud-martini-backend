```yaml
  # build-martini-agent:
  #   runs-on: ubuntu-latest
  #   name: Build Martini Agent
  #   needs: [build, deploy]
    
  #   env:
  #     REGISTRY_OWNER: chamodshehanka
  #     IMAGE_TAG: ${{ needs.build.outputs.image_tag }}

  #   steps:
  #     - name: Checkout code
  #       uses: actions/checkout@v4

  #     - name: Set up QEMU
  #       uses: docker/setup-qemu-action@v2

  #     - name: Set up Docker Buildx
  #       uses: docker/setup-buildx-action@v2

  #     - name: Log in to GitHub Container Registry
  #       run: echo ${{ secrets.GHCR_TOKEN }} | docker login ghcr.io -u ${{ github.actor }} --password-stdin

  #     - name: Cache Docker layers for martini-agent
  #       uses: actions/cache@v3
  #       with:
  #         path: /tmp/.buildx-cache-martini-agent
  #         key: ${{ runner.os }}-buildx-martini-agent-${{ github.sha }}
  #         restore-keys: |
  #           ${{ runner.os }}-buildx-martini-agent-

  #     - name: Build and push martini-agent Docker image
  #       uses: docker/build-push-action@v2
  #       with:
  #         context: ./martini-agent
  #         push: true
  #         tags: ghcr.io/${{ env.REGISTRY_OWNER }}/martini-agent:${{ env.IMAGE_TAG }}
  #         cache-from: type=local,src=/tmp/.buildx-cache-martini-agent
  #         cache-to: type=local,dest=/tmp/.buildx-cache-martini-agent
  # deploy-martini-agent:
  #   name: Deploy Martini Agent
  #   runs-on: ubuntu-latest
  #   needs: [build, deploy, build-martini-agent]

  #   env:
  #     GKE_CLUSTER: circles-cluster
  #     GKE_ZONE: us-central1
  #     IMAGE_TAG: ${{ needs.build.outputs.image_tag }}
  #     REGISTRY_OWNER: chamodshehanka

  #   steps:
  #     - name: Checkout code
  #       uses: actions/checkout@v4

  #     - name: Set up Kubernetes
  #       uses: azure/setup-kubectl@v1
  #       with:
  #         version: v1.21.0

  #     - name: Set up Helm
  #       uses: azure/setup-helm@v1
  #       with:
  #         version: v3.5.4

  #     # Setup gcloud CLI
  #     - id: 'auth'
  #       uses: 'google-github-actions/auth@v2'
  #       with:
  #         credentials_json: '${{ secrets.GKE_SA_KEY }}'

  #     # Get the GKE credentials so we can deploy to the cluster
  #     - uses: google-github-actions/get-gke-credentials@v2
  #       with:
  #         cluster_name: ${{ env.GKE_CLUSTER }}
  #         location: ${{ env.GKE_ZONE }}
  #         project_id: ${{ secrets.GOOGLE_PROJECT_ID }}
      
  #     - name: Test Cluster Connectivity
  #       run: kubectl cluster-info

  #     - name: Deploy to Kubernetes
  #       run: |
  #         helm upgrade --install martini-agent-job ./helm/martini-agent \
  #           --set martiniAgent.image.repository=ghcr.io/${{ env.REGISTRY_OWNER }}/martini-agent \
  #           --set martiniAgent.image.tag=${{ env.IMAGE_TAG }}
```