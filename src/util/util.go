package util

import (
	utilFile "agent_office/src/util/file"
	utilString "agent_office/src/util/string"
)

type UtilContainers struct {
	UString utilString.UtilsString
	UHash   utilString.UtilsHash
	UFile   utilFile.UtilsFile
}

func RequireUtil() UtilContainers {
	return UtilContainers{
		UString: utilString.UseUtilsString(),
		UHash:   utilString.UseHashedString(),
		UFile:   utilFile.UseFileHelper(),
	}
}
