package constants

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var frameworkMap map[string]string = map[string]string{
	"TRITON":                 "triton",
	"PYTORCH":                "pytorch",
	"LLAMA":                  "llma",
	"CODELAMA":               "codellama",
	"STABLE_DIFFUSION":       "stable_diffusion",
	"STABLE_DIFFUSION_XL":    "stable_diffusion_xl",
	"MPT":                    "mpt",
	"CUSTOM":                 "custom",
	"MIXTRAL8X7B":            "mixtral-8x7b-instruct",
	"MIXTRAL7B":              "mistral-7b-instruct",
	"TENSOR_RT":              "tensorrt",
	"GEMMA_2B":               "gemma-2b",
	"GEMMA_2B_IT":            "gemma-2b-it",
	"GEMMA_7B":               "gemma-7b",
	"GEMMA_7B_IT":            "gemma-7b-it",
	"LLAMA_3":                "llama-3-8b-instruct",
	"LLAMA_3_1":              "llama-3_1-8b-instruct",
	"LLAMA_3_2":              "llama-3_2-3b-instruct",
	"LLAMA_3_2_VISION":       "llama-3_2-11b-vision-instruct",
	"VLLM":                   "vllm",
	"STARCODER":              "starcoder2-7b",
	"PHI_3_MINI":             "Phi-3-mini-128k-instruct",
	"NEMO":                   "nemo-rag",
	"STABLE_VIDEO_DIFFUSION": "stable-video-diffusion-img2vid-xt",
	"YOLO_V8":                "yolov8",
	"NEMOTRON":               "nemotron-3-8b-chat-4k-rlhf",
	"NV_EMBED":               "nvidia-nv-embed-v1",
	"BAAI_LARGE":             "bge-large-en-v1_5",
	"BAAI_RERANKER":          "bge-reranker-large",
	"PIXTRAL":                "pixtral-12b-2409",
	"SGLANG" : 				  "sglang",
	"DYNAMO" : 				  "dynamo",	
}

var FrameworkContainerNames = map[string]map[string]string{
	"TRITON": {
		"v24.02": "aimle2e/tritonserver:24.02-py3-01",
		"v24.01": "aimle2e/tritonserver:24.01-py3-01",
		"v23.12": "aimle2e/tritonserver:23.12-py3-01",
		"v23.11": "aimle2e/tritonserver:23.01-py3-01",
		"v23.10": "aimle2e/tritonserver:23.10-py3-01",
	},
	"PYTORCH": {
		"v0.9.0": "pytorch/torchserve:0.9.0",
		"v0.8.2": "pytorch/torchserve:0.8.2",
		"v0.8.1": "pytorch/torchserve:0.8.1",
	},
	"LLAMA": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"LLAMA_3": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"SGLANG": {
		"MODEL_SELECTED":     "lmsysorg/sglang:latest",
		"MODEL_NOT_SELECTED": "lmsysorg/sglang:latest",
	},
	"LLAMA_3_1": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"LLAMA_3_2": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"LLAMA_3_2_VISION": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"CODELAMA": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"STABLE_DIFFUSION": {
		"MODEL_SELECTED":     "registry.e2enetworks.net/aimle2e/stable-diffusion-2-1:eos-v1",
		"MODEL_NOT_SELECTED": "registry.e2enetworks.net/aimle2e/stable-diffusion-2-1:hf-v1",
	},
	"STABLE_DIFFUSION_XL": {
		"MODEL_SELECTED":     "registry.e2enetworks.net/aimle2e/stable-diffusion-xl-base-1.0:eos",
		"MODEL_NOT_SELECTED": "registry.e2enetworks.net/aimle2e/stable-diffusion-xl-base-1.0:hf",
	},
	"MPT": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"MIXTRAL8X7B": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"MIXTRAL7B": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"TENSOR_RT": {
		"v24.02":  "aimle2e/tritonserver:24.02-trtllm-python-py3-01",
		"v24.01":  "aimle2e/tritonserver:24.01-trtllm-python-py3-01",
		"v23.12":  "aimle2e/tritonserver:23.12-trtllm-python-py3-01",
		"v23.11":  "aimle2e/tritonserver:23.11-trtllm-python-py3-01",
		"v23.10":  "aimle2e/tritonserver:23.10-trtllm-python-py3-01",
		"v0.10.0": "aimle2e/triton_trt_llm:v0.10.0",
		"v0.9.0":  "aimle2e/triton_trt_llm:0.9.0",
		"v0.7.2":  "aimle2e/triton_trt_llm:0.7.2",
	},
	"GEMMA_2B": {
		"MODEL_NOT_SELECTED": "registry.e2enetworks.net/aimle2e/triton_trt_llm:gemma-v1",
	},
	"GEMMA_2B_IT": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"GEMMA_7B": {
		"MODEL_NOT_SELECTED": "registry.e2enetworks.net/aimle2e/triton_trt_llm:gemma-v1",
	},
	"GEMMA_7B_IT": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"VLLM": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"STARCODER": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"PHI_3_MINI": {
		"MODEL_SELECTED":     "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED": "vllm/vllm-openai:latest",
	},
	"NEMO": {
		"v0.9.0": "registry.e2enetworks.net/aimle2e/nemo-rag:0.9.0",
	},
	"STABLE_VIDEO_DIFFUSION": {
		"MODEL_SELECTED":     "aimle2e/stable-video-diffusion:v1_eos",
		"MODEL_NOT_SELECTED": "aimle2e/stable-video-diffusion:v1",
	},
	"YOLO_V8": {
		"MODEL_SELECTED":     "registry.e2enetworks.net/aimle2e/yolov8:v1",
		"MODEL_NOT_SELECTED": "registry.e2enetworks.net/aimle2e/yolov8:v1",
	},
	"NEMOTRON": {
		"MODEL_SELECTED":     "aimle2e/nemotron:3-8b-chat-4k-rlhf",
		"MODEL_NOT_SELECTED": "aimle2e/nemotron:3-8b-chat-4k-rlhf",
	},
	"NV_EMBED": {
		"MODEL_SELECTED":     "aimle2e/nv_embed_v1:v1_eos",
		"MODEL_NOT_SELECTED": "aimle2e/nv_embed_v1:v1",
	},
	"BAAI_LARGE": {
		"MODEL_SELECTED":     "ghcr.io/huggingface/text-embeddings-inference:1.5",
		"MODEL_NOT_SELECTED": "ghcr.io/huggingface/text-embeddings-inference:1.5",
	},
	"BAAI_RERANKER": {
		"MODEL_SELECTED":     "ghcr.io/huggingface/text-embeddings-inference:1.5",
		"MODEL_NOT_SELECTED": "ghcr.io/huggingface/text-embeddings-inference:1.5",
	},
	"DYNAMO": {
		"MODEL_SELECTED":     "aimle2e/dynamo:latest-vllm",
		"MODEL_NOT_SELECTED": "aimle2e/dynamo:latest-vllm",
	},
	"PIXTRAL" : {
		"MODEL_SELECTED" : "vllm/vllm-openai:latest",
		"MODEL_NOT_SELECTED" : "vllm/vllm-openai:latest",
	},
}


func GetContainerName(server_option string, model_id string, framework string) (diag.Diagnostics, string) {
	_, ok := FrameworkContainerNames[framework]
	if !ok {
		log.Println("ok")
		return diag.Errorf("Error finding the framework, please enter the correct framework"), ""
	}

	if server_option != "" {
		log.Println("server", server_option)
		return nil, FrameworkContainerNames[framework][server_option]
	}

	if model_id != "" {
		log.Println("model_id",model_id)
		return nil, FrameworkContainerNames[framework]["MODEL_SELECTED"]
	} else {
		log.Println(FrameworkContainerNames[framework]["MODEL_NOT_SELECTED"])
		return nil, FrameworkContainerNames[framework]["MODEL_NOT_SELECTED"]
	}

}

func GetFrameworkName(framework string) (string, diag.Diagnostics) {
	frameName, ok := frameworkMap[framework]
	if !ok {
		return "", diag.Errorf("Please provide the framework name correctly")
	}
	return frameName, nil
}


func GetDefaultHuggingFaceID(framework string) string {
    switch framework {
    case "STARCODER":
        return "bigcode/starcoder2-7b"
    case "MIXTRAL7B":
        return "mistralai/Mistral-7B-Instruct-v0.1"
    case "MIXTRAL8X7B":
        return "mistralai/Mixtral-8x7B-Instruct-v0.1"
    case "PHI_3_MINI":
        return "microsoft/Phi-3-mini-128k-instruct"
    case "LLAMA":
        return "meta-llama/Llama-2-7b-chat-hf"
    case "LLAMA_3":
        return "meta-llama/Meta-Llama-3-8B-Instruct"
    case "LLAMA_3_1":
        return "meta-llama/Meta-Llama-3.1-8B-Instruct"
    case "LLAMA_3_2":
        return "meta-llama/Llama-3.2-3B-Instruct"
    case "LLAMA_3_2_VISION":
        return "meta-llama/Llama-3.2-11B-Vision-Instruct"
    case "GEMMA_2B_IT":
        return "google/gemma-2b-it"
    case "GEMMA_7B_IT":
        return "google/gemma-7b-it"
    case "CODELAMA":
        return "meta-llama/CodeLlama-7b-Instruct-hf"
    case "MPT":
        return "mosaicml/mpt-7b-instruct"
    case "NEMOTRON":
        return "nvidia/nemotron-3-8b-chat-4k-rlhf"
    case "NV_EMBED":
        return "nvidia/NV-Embed-v1"
    case "STABLE_DIFFUSION":
        return "stabilityai/stable-diffusion-2-1"
    case "STABLE_DIFFUSION_XL":
        return "stabilityai/stable-diffusion-xl-base-1.0"
    case "STABLE_VIDEO_DIFFUSION":
        return "stabilityai/stable-video-diffusion-img2vid-xt"
    case "PIXTRAL":
        return "mistralai/Pixtral-12B-2409"
    default:
        return ""
    }
}