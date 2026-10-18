package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 bool) int64 {
    if ((("" + ("foo bar" + "\n\t")) > ("foo bar" + ("hello" + ("" + "a")))) && (0 != 870)) {
        if p0 {
            if p0 {
                var v4 uint64 = 2
                var v5 bool = false
                _ = v4
                _ = v5
            }
        }
        var v6 int = -985
        _ = v6
    }
    if (128 < 1) {
        if (0 != 128) {
            var v7 int = -1
            var v8 int8 = 1
            v8 = v8
            _ = v7
            _ = v8
        }
    }
    switch "" {
    case "\n\t":
        p0 = true
    default:
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: "a", F1: 1, F2: 375}
            v11 := GenS2{F0: (fi9 >> ((524) & 63))}
            var v12 int64 = 619
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    return 42
}

func main() {
    var v1 string = ("\n\t" + "hello")
    v2 := make(map[string]string)
    v3 := []string{}
    for fi13 := 0; fi13 < 16; fi13++ {
        if false {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS1{F0: v1, F1: 0, F2: 65535}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch (fi14 / (((fi14 + fi13)) | 1)) {
                case 32768:
                    var v16 string = v1
                    _ = v16
                case -66:
                    v2[((v1 + v1) + (v1 + v1))] = v1
                case -438:
                    fmt.Println(fi13, fi14, v1)
                default:
                    v1 = (("a" + ("foo bar" + v1)) + (((v1 + ("\n\t" + v1)) + v1) + v1))
                }
                if len(v3) > 0 {
                    v3[0] = v1
                }
                var v17 uint64 = 82
                _ = fi14
                _ = v15
                _ = v17
            }
        }
        v1 = v1
        var v18 uint = 65535
        _ = fi13
        _ = v18
    }
    v1 = (((v1 + ("hello" + (("\n\t" + v1) + ""))) + v1) + v1)
    var v19 int32 = 0
    v1 = ""
    var v20 uint16 = 65535
    fmt.Println(v1, v19)
    v19 += v19
    v1 = v1
    for fi21 := 0; fi21 < 13; fi21++ {
        v1 = "hello"
        _ = fi21
    }
    if false {
        v22 := GenS2{F0: 4294967295}
        if true {
            switch v1 {
            case "\n\t":
                v19 += v19
            default:
                fmt.Println(v20)
            }
            v2[v1] = v1
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS3{F0: (!(-949 < 9223372036854775807)), F1: ((fi23 | (-(fi23))) > -1)}
                v19 -= 42
                v25 := GenS2{F0: (fi23 % (((fi23 ^ fi23)) | 1))}
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            var v26 uint8 = 0
            _ = v26
        }
        v19 -= v19
        _ = v22
    } else {
        v1 = "hello"
        v27 := GenS4{F0: 128, F1: 203}
        var v28 uint = 65535
        _ = v27
        _ = v28
    }
    v3 = append(v3, v1)
    for fi29 := 0; fi29 < 16; fi29++ {
        if (253 == 42) {
            if (9223372036854775807 >= -1) {
                var v30 uint8 = 2
                _ = v30
            } else {
                switch "" {
                case "hello":
                    var v31 bool = false
                    _ = v31
                default:
                    delete(v2, ("" + v1))
                }
                switch (fi29 - fi29) {
                case 32768:
                    v32 := GenS5{F0: 1, F1: 2, F2: ("foo bar" + (v1 + v1))}
                    _ = v32
                }
                if false {
                    var v33 int32 = (v19 & (^(v19)))
                    var v34 int64 = genFn0((308 == 956))
                    var v35 uint64 = 896
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    var v36 string = v1
                    _ = v36
                }
            }
            fmt.Println(v1, v1, fi29)
        }
        _ = fi29
    }
    var v37 float32 = ((0.1 / 3.14159) * 0.0)
    delete(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
}
