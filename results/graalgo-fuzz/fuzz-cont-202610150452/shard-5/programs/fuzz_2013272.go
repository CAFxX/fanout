package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    var v3 bool = true
    v4 := []string{"a", "\n\t", (("" + "\n\t") + "\n\t"), "hello"}
    v5 := make(map[uint32]bool)
    if v3 {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 65535, F1: 39, F2: fi6}
            v2 = false
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi6
            _ = v7
        }
        switch 32767 {
        case 32767:
            var v8 uint32 = 42
            _ = v8
        case 0:
            fmt.Println(v2, v3)
        case 143:
            switch ("a" + "\n\t") {
            case "a":
                if (!false) {
                    v3 = v2
                }
            case "\n\t":
                v3 = v2
            }
        default:
            switch -855 {
            case 413:
                if v3 {
                    if (!(v3 || v3)) {
                        var v9 int32 = 1
                        _ = v9
                    } else {
                        if (!v3) {
                            var v10 int16 = 1
                            v3 = (-1 < 127)
                            _ = v10
                        }
                    }
                } else {
                    v2 = true
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 float64 = -2.25
                        v13 := GenS1{F0: 855, F1: 42, F2: (fi11 - -953)}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            default:
                var v14 float32 = (3.14159 + 0.0)
                _ = v14
            }
        }
    } else {
        var v15 int64 = -1
        _ = len(v4)
        v3 = false
        _ = v15
    }
    fmt.Println(v2, v3)
    switch 9223372036854775807 {
    case -246:
        v16 := GenS1{F0: 42, F1: 42, F2: 32767}
        _ = v16
    case 128:
        _ = len(v5)
    default:
        fmt.Println(v3)
    }
    v17 := GenS2{F0: ""}
    var v18 uint16 = 65535
    v19 := GenS3{F0: 2, F1: 228}
    if (127 != -89) {
        for fi20 := 0; fi20 < 8; fi20++ {
            switch "a" {
            case "\n\t":
                v21 := GenS1{F0: (v18 ^ (v18 | (^(v18)))), F1: 61, F2: -135}
                _ = v21
            case "foo bar":
                v2 = true
            }
            if false {
                v3 = (!v2)
                v2 = false
            } else {
                v22 := GenS3{F0: 2, F1: 788}
                fmt.Println(fi20)
                var v23 float64 = (0.1 - 0.0)
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        fmt.Println(v18)
    }
    var v24 int32 = 0
    var v25 int64 = 42
    var v26 float64 = 1.5
    for fi27 := 0; fi27 < 10; fi27++ {
        if (!false) {
            var v28 uint16 = 65535
            _ = v28
        } else {
            if true {
                fmt.Println(v18, fi27, v3)
            } else {
                v29 := GenS4{F0: v24}
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 int64 = 42
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            v26 = (((0.1 + v26) / (0.1 / 1e10)) * (0.1 / v26))
        }
        var v32 int16 = 0
        _ = fi27
        _ = v32
    }
    if len(v4) > 0 {
        v4[0] = ("" + "hello")
    }
    switch -352 {
    case 2147483647:
        var v33 uint32 = 0
        _ = v33
    case 127:
        v4 = append(v4, ("a" + ("" + ("\n\t" + "a"))))
    }
    v2 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
