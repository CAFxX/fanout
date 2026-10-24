package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 string }
type GenS3 struct { F0 uint16; F1 int16; F2 uint8 }
func genFn0() uint {
    var v5 float32 = float32(-2.25)
    v5 += float32(((0.0 + (float64(float32(3.14159)) / -2.25)) / 3.14159))
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 2
    var v3 float64 = 0.0
    v4 := make(map[uint32]uint32)
    var v6 int8 = 46
    if (2 == 186) {
        if (!false) {
            if (v3 >= (v3 - v3)) {
                v7 := v4[2147483648]
                _ = v7
            }
            v8 := GenS1{F0: 0, F1: 0}
            for fi9 := 0; fi9 < 20; fi9++ {
                v2 -= (v2 / ((genFn0()) | 1))
                var v10 int8 = (v6 % ((127) | 1))
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            var v11 float64 = v3
            _ = v11
        }
        var v12 uint64 = 42
        _ = v12
    } else {
        var v13 int = -665
        _ = v13
    }
    var v14 int = -1
    fmt.Println(v14, v2)
    v2 = 713
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float32 = ((-2.25 - -2.25) * 3.14159)
        v14 -= -1
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int16 = 32767
            if false {
                v19 := GenS2{F0: 3, F1: -2147483648, F2: "a"}
                _ = v19
            } else {
                var v20 int32 = -2147483648
                var v21 int = (fi17 / (((fi17 - v14)) | 1))
                delete(v1, 2)
                _ = v20
                _ = v21
            }
            v22 := []int{v14, fi15, (^(fi17))}
            _ = fi17
            _ = v18
            _ = v22
        }
        _ = fi15
        _ = v16
    }
    v23 := []int{128}
    if true {
        switch v14 {
        case 9223372036854775807:
            v24 := GenS3{F0: 65535, F1: 927, F2: 203}
            _ = v24
        case 4294967295:
            v25 := v1[465]
            _ = v25
        }
        v6 -= v6
    }
    if (true || (560 != 949)) {
        var v26 uint64 = 42
        _ = v26
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            v28 := GenS3{F0: 255, F1: 32767, F2: 100}
            _ = fi27
            _ = v28
        }
        var v29 int = (v14 / (((v14 << ((v2) & 63))) | 1))
        _ = v29
    }
    switch "foo bar" {
    case "foo bar":
        var v30 bool = false
        _ = v30
    default:
        var v31 uint32 = 2147483647
        _ = v31
    }
    v14 = v14
    if ((false && (true && (1e10 < (-2.25 / 1e10)))) || ((((0.1 + 0.0) / 3.14159) - 1e10) == (-2.25 + 1e10))) {
        v32 := GenS3{F0: 42, F1: -928, F2: 0}
        fmt.Println(v2, v3, v2)
        _ = v32
    }
    v33 := GenS2{F0: 1, F1: 945, F2: "foo bar"}
    if false {
        var v34 uint32 = 2
        _ = v34
    } else {
        v35 := v4[711]
        switch "" {
        case "":
            var v36 int = (v14 | v14)
            _ = v36
        case "foo bar":
            v1[v35] = v35
        case "hello":
            var v37 int32 = 1
            _ = v37
        }
        _ = v35
    }
    var v38 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v6
}
