package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    v3 := []string{("hello" + "hello")}
    var v4 uint64 = 162
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 1}
        var v7 int = fi5
        var v8 int = 417
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1}
    if true {
        v4 -= v4
        for fi10 := 0; fi10 < 3; fi10++ {
            switch fi10 {
            case 42:
                v4 = 1
            case 326:
                var v11 int32 = 0
                _ = v11
            default:
                v12 := GenS2{F0: 500, F1: 284, F2: 396}
                _ = v12
            }
            fmt.Println(fi10, fi10)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi10
        }
    } else {
        v13 := GenS1{F0: 9223372036854775807}
        _ = v13
    }
    switch "" {
    case "":
        switch 42 {
        case 0:
            fmt.Println(v4)
        }
    case "hello":
        var v14 float64 = ((((0.1 * ((0.1 * (((1e10 + 3.14159) - (1e10 + 1e10)) - 0.0)) - float64(1e10))) + 3.14159) * 1.5) / 1.0)
        _ = v14
    case "\n\t":
        if ((true || true) && (2 != 65535)) {
            v15 := GenS2{F0: 1, F1: 908, F2: 2147483647}
            if true {
                for fi16 := 0; fi16 < 2; fi16++ {
                    for fi17 := 0; fi17 < 17; fi17++ {
                        _ = len(v1)
                        switch (fi16 & -141) {
                        case 32768:
                            var v18 uint64 = v4
                            _ = v18
                        }
                        _ = fi17
                    }
                    _ = fi16
                }
            }
            _ = v15
        }
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v4, v4)
        _ = len(v1)
        _ = fi19
    }
    var v20 string = ""
    v4 = 4294967295
    _ = len(v2)
    v21 := GenS1{F0: -875}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
