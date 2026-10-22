package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint; F1 string; F2 int8 }

func main() {
    var v1 uint8 = 1
    v2 := []string{("hello" + "")}
    var v3 float64 = (1.5 + 1.5)
    for fi4 := 0; fi4 < 20; fi4++ {
        switch fi4 {
        case 313:
            var v5 float32 = -2.25
            _ = v5
        case 2147483647:
            v3 = 0.1
        case -92:
            v1 -= (v1 + 40)
        default:
            for fi6 := 0; fi6 < 3; fi6++ {
                v7 := GenS1{F0: v1, F1: 2147483647, F2: 100}
                _ = fi6
                _ = v7
            }
        }
        v2 = append(v2, ("" + ((("a" + ("a" + "a")) + "hello") + ("a" + ""))))
        if (-32768 >= int16(v1)) {
            v8 := GenS2{F0: 42, F1: 4294967295}
            _ = v8
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
        }
        _ = fi4
    }
    v9 := GenS3{F0: 443, F1: 0}
    var v10 uint = uint(v1)
    fmt.Println(v1, v3)
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 float64 = 3.14159
        if (!(false || true)) {
            if (false || (0 < 128)) {
                fmt.Println(v3, v1, v10)
            } else {
                v13 := GenS1{F0: (v1 - v1), F1: 0, F2: -1}
                if false {
                    v10 = v10
                } else {
                    v12 += -2.25
                    var v14 uint16 = 2
                    v1 = (v1 + v1)
                    _ = v14
                }
                _ = v13
            }
            fmt.Println(v3, v10, v1)
            v2 = append(v2, "\n\t")
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3, v1, v1)
    v3 -= v3
    var v15 uint16 = 1
    fmt.Println(v3, v10, v10)
    var v16 int8 = -46
    v1 += (v1 | v1)
    for fi17 := 0; fi17 < 6; fi17++ {
        v3 += (((v3 / (v3 * 0.1)) + -2.25) + 0.1)
        _ = fi17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 -= (v1 - v1)
        _ = fi18
    }
    switch "foo bar" {
    case "\n\t":
        v19 := GenS4{F0: (v10 / (((v10 * v10)) | 1)), F1: "a", F2: v16}
        _ = v19
    case "a":
        var v20 uint8 = (v1 - (v1 + (v1 - (v1 >> ((543) & 63)))))
        _ = v20
    case "":
        var v21 uint64 = 1
        _ = v21
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            v10 = (v10 / (((v10 / ((18446744073709551615) | 1))) | 1))
            v23 := GenS1{F0: (v1 - 128), F1: 977, F2: (v16 >> ((v10) & 63))}
            switch 127 {
            case 32768:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            case 32767:
                v3 = (v3 / 3.14159)
            case 4294967295:
                _ = len(v2)
            default:
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 uint64 = 42
                    v16 = v16
                    v26 := GenS4{F0: 4294967295, F1: "\n\t", F2: 127}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi22
            _ = v23
        }
    }
    var v27 float32 = (0.1 + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v9
}
