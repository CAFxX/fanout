package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0() uint64 {
    var v4 int = 175
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float32 = 0.0
        _ = fi5
        _ = v6
    }
    _ = v4
    return 18446744073709551615
}

func main() {
    v1 := []int{-9223372036854775808, 32768}
    var v2 float64 = 1.5
    var v3 uint64 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 int = 1
    if len(v1) > 0 {
        v1[0] = (v7 % (((v7 | v7)) | 1))
    }
    if (128 == 127) {
        if false {
            if false {
                fmt.Println(v2, v2)
                fmt.Println(v2, v2)
                v8 := GenS1{F0: v7, F1: 42}
                _ = v8
            }
            v3 += (v3 - 1)
        } else {
            switch 2147483647 {
            case 9223372036854775807:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 32767:
                switch ("hello" + "hello") {
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "foo bar":
                    v7 = (^(v7))
                }
            }
        }
        var v9 uint16 = 255
        if (v3 < v3) {
            if true {
                v3 += v3
            }
            v9 = uint16(v3)
        }
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v3, fi10, fi10)
        fmt.Println(v7, fi10)
        _ = fi10
    }
    if len(v1) > 0 {
        v1[0] = (-(v7))
    }
    var v11 float32 = 0.1
    if false {
        switch v7 {
        case 42:
            if false {
                v12 := GenS1{F0: v7, F1: -1}
                _ = v12
            }
        case 128:
            var v13 float32 = 1.5
            _ = v13
        default:
            var v14 int8 = 100
            _ = v14
        }
    }
    switch "" {
    case "a":
        switch ("" + "\n\t") {
        case "\n\t":
            var v15 uint16 = 1
            _ = v15
        case "a":
            _ = len(v1)
        default:
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int8 = 100
                fmt.Println(v2)
                _ = fi16
                _ = v17
            }
        }
    case "foo bar":
        var v18 float32 = (float32(0.1) + 0.1)
        _ = v18
    case "":
        if (true || true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v19 int8 = 42
            v1 = append(v1, (v7 / ((-1) | 1)))
            _ = v19
        }
    }
    if true {
        var v20 bool = false
        if (v2 > 3.14159) {
            v21 := GenS2{F0: 2}
            switch (v7 & (v7 ^ 32767)) {
            case 127:
                v20 = (65535 == 826)
            case 32767:
                _ = len(v1)
            default:
                v20 = (v11 <= v11)
            }
            v3 -= v3
            _ = v21
        }
        fmt.Println(v7, v20)
        _ = v20
    } else {
        var v22 int8 = -128
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint = 975
            if false {
                v2 = (v2 + (float64(0.1) - float64((v11 / v11))))
                v25 := GenS1{F0: (fi23 % ((32767) | 1)), F1: 1}
                _ = v25
            } else {
                v26 := GenS1{F0: (-(fi23)), F1: 42}
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v27 := GenS3{F0: 0, F1: -61}
    var v28 uint64 = (v3 - 42)
    for fi29 := 0; fi29 < 6; fi29++ {
        if len(v1) > 0 {
            v1[0] = (v7 ^ (fi29 >> ((1000) & 63)))
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v7
}
