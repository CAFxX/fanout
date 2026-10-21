package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
func genFn0() uint32 {
    switch 4294967295 {
    case 944:
        if false {
            var v5 string = "\n\t"
            for fi6 := 0; fi6 < 10; fi6++ {
                var v7 uint32 = 2
                var v8 uint8 = 127
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        }
    case 32767:
        var v9 float32 = 1e10
        _ = v9
    case -705:
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int64 = 42
            _ = fi10
            _ = v11
        }
    default:
        var v12 float32 = 3.14159
        _ = v12
    }
    var v13 bool = false
    switch 9223372036854775807 {
    case 0:
        fmt.Println(v13)
    case 1:
        var v14 bool = true
        _ = v14
    case -296:
        fmt.Println(v13)
    default:
        switch "" {
        case "":
            var v15 int8 = 127
            _ = v15
        case "\n\t":
            if true {
                var v16 uint64 = 1
                _ = v16
            } else {
                if ("" < ("hello" + "\n\t")) {
                    var v17 int = 2147483647
                    _ = v17
                }
                v18 := GenS1{F0: (v13 && v13)}
                _ = v18
            }
        case "a":
            v13 = false
        }
    }
    _ = v13
    return 0
}
func genFn1() string {
    v19 := GenS1{F0: (2 != 127)}
    var v20 uint8 = 162
    _ = v19
    _ = v20
    return "foo bar"
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]uint32)
    var v3 int32 = 42
    v4 := []string{("hello" + ("foo bar" + "hello"))}
    v1 -= (v1 | v1)
    var v21 uint64 = 42
    var v22 int32 = -637
    v23 := GenS2{F0: v3, F1: 1, F2: uint32(v21)}
    var v24 uint16 = 0
    switch 32768 {
    case 128:
        var v25 uint64 = 455
        _ = v25
    case -1:
        for fi26 := 0; fi26 < 12; fi26++ {
            v4 = append(v4, "")
            if true {
                switch -1 {
                case 449:
                    v27 := GenS1{F0: false}
                    _ = v27
                case 32768:
                    v3 = v22
                default:
                    v22 -= v22
                }
            } else {
                fmt.Println(v24, v1, fi26)
            }
            _ = fi26
        }
    default:
        v1 = (v1 ^ v1)
    }
    switch -9223372036854775808 {
    case -1:
        if (!false) {
            v3 += v22
            v22 = v22
            if false {
                v3 = -845
                v28 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: "\n\t"}
                _ = v28
            }
        }
    case 4294967295:
        v1 += v1
    }
    var v29 int8 = -34
    v30 := GenS4{F0: 0, F1: v3, F2: v22}
    v1 = -63
    v31 := GenS3{F0: 2, F1: 1, F2: ("foo bar" + "a")}
    var v32 uint8 = uint8(v21)
    if (2 > (v24 | 255)) {
        v3 += -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
