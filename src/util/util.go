package util

import (
	utilFile "agent_office/src/util/file"
	utilJwt "agent_office/src/util/jwt"
	utilString "agent_office/src/util/string"
)

type UtilContainers struct {
	UString utilString.UtilsString
	UHash   utilString.UtilsHash
	UFile   utilFile.UtilsFile
	UJwt    utilJwt.UtilsJwt
}

func RequireUtil() UtilContainers {
	return UtilContainers{
		UString: utilString.UseUtilsString(),
		UHash:   utilString.UseHashedString(),
		UFile:   utilFile.UseFileHelper(),
		UJwt:    utilJwt.UseJwt(),
	}
}
