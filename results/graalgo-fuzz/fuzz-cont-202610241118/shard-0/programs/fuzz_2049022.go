package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 int64) int {
    var v5 bool = (!true)
    fmt.Println(p2)
    v6 := GenS1{F0: 325, F1: 100}
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]string)
    var v3 uint32 = (v1 & 2147483647)
    v4 := []string{}
    switch -168 {
    case -692:
        switch "a" {
        case "foo bar":
            if (42 != 127) {
                if false {
                    fmt.Println(v1, v3)
                    v1 -= v3
                }
                if false {
                    fmt.Println(v3, v3)
                } else {
                    var v8 string = (("foo bar" + "hello") + "\n\t")
                    for fi9 := 0; fi9 < 18; fi9++ {
                        var v10 int64 = -823
                        var v11 uint64 = 1
                        v2[(v3 - (v1 / (((v1 & v1)) | 1)))] = v8
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v8
                }
                var v12 uint = 18446744073709551615
                _ = v12
            } else {
                fmt.Println(v1, v1)
                _ = len(v4)
            }
        case "hello":
            var v13 string = (("hello" + ("foo bar" + "hello")) + ("" + "hello"))
            _ = v13
        }
    case 32768:
        fmt.Println(v3, v3)
    default:
        v1 += v1
    }
    switch 710 {
    case -2147483648:
        var v14 uint8 = 1
        _ = v14
    case 9223372036854775807:
        fmt.Println(v1, v3)
    case 3:
        var v15 uint8 = 255
        _ = v15
    }
    switch "foo bar" {
    case "hello":
        v16 := GenS1{F0: 2147483647, F1: -1}
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 7; fi17++ {
            switch "hello" {
            case "hello":
                if (false || false) {
                    v2[(v1 << ((1) & 63))] = "foo bar"
                } else {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        var v19 string = "a"
                        var v20 int = 166
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint32 = (v1 + 2)
                    _ = v21
                }
            }
            _ = len(v4)
            v3 += (v3 & v1)
            _ = fi17
        }
    case "":
        if (((0.0 + (((0.0 + 1e10) - ((((1.5 + 0.1) + (0.1 + 1.5)) + 0.1) * 3.14159)) + 3.14159)) == (0.0 + 3.14159)) && (42 < 42)) {
            v22 := GenS1{F0: 0, F1: -21}
            v23 := GenS2{F0: 127, F1: -9223372036854775808, F2: 42}
            _ = v22
            _ = v23
        } else {
            var v24 int8 = int8(v1)
            _ = len(v2)
            var v25 uint32 = 2147483648
            _ = v24
            _ = v25
        }
    default:
        var v26 uint16 = 255
        _ = v26
    }
    var v27 int8 = -1
    var v28 uint = 2
    v1 = (v3 / ((v1) | 1))
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS2{F0: (v27 ^ 0), F1: -868, F2: v3}
        fmt.Println(v3, v1)
        _ = fi29
        _ = v30
    }
    var v31 string = "hello"
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 bool = true
        for fi34 := 0; fi34 < 18; fi34++ {
            v35 := GenS3{F0: 1, F1: 0, F2: 0}
            _ = fi34
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
