package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 int32; F1 int; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0(p0 uint8, p1 bool, p2 int8) uint32 {
    v5 := GenS1{F0: 183, F1: (p0 / (((p0 + 80)) | 1)), F2: 106}
    var v6 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 int8 = -1
    var v2 uint = 0
    v3 := []string{"hello", "hello", ("foo bar" + "\n\t"), ("foo bar" + "hello")}
    var v4 int8 = v1
    v7 := GenS1{F0: 593, F1: 127, F2: 0}
    if true {
        v1 = v1
    }
    if (true || true) {
        v8 := GenS2{F0: 400, F1: 824, F2: 0}
        if (32767 == 780) {
            v9 := GenS2{F0: 1, F1: 2147483647, F2: 1}
            fmt.Println(v2, v1)
            var v10 int16 = 32767
            _ = v10
            _ = v9
        }
        fmt.Println(v1)
        _ = v8
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v3)
    if (!false) {
        _ = len(v3)
    } else {
        if len(v3) > 0 {
            v3[0] = ("" + ("foo bar" + "foo bar"))
        }
        _ = len(v3)
        switch (("hello" + "\n\t") + "a") {
        case "foo bar":
            v1 = v1
        case "hello":
            if len(v3) > 0 {
                v3[0] = "a"
            }
        case "a":
            if false {
                if true {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        var v12 uint32 = 116
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v14 := GenS1{F0: 757, F1: 34, F2: -128}
                        _ = fi13
                        _ = v14
                    }
                } else {
                    var v15 uint64 = 42
                    for fi16 := 0; fi16 < 14; fi16++ {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            var v18 int64 = 0
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                    switch "hello" {
                    case "":
                        v19 := GenS1{F0: 413, F1: 109, F2: v4}
                        _ = v19
                    }
                    _ = v15
                }
                for fi20 := 0; fi20 < 15; fi20++ {
                    switch "" {
                    case "":
                        v4 = v1
                    case "hello":
                        var v21 uint32 = 42
                        _ = v21
                    default:
                        var v22 string = "\n\t"
                        _ = v22
                    }
                    _ = fi20
                }
                var v23 int16 = -32768
                _ = v23
            }
        default:
            if true {
                var v24 bool = false
                v25 := GenS2{F0: 2147483647, F1: 0, F2: 616}
                v26 := GenS3{F0: 9223372036854775807, F1: 2}
                _ = v24
                _ = v25
                _ = v26
            }
        }
    }
    var v27 float64 = 0.1
    v1 = v1
    if (!false) {
        v28 := GenS4{F0: 431, F1: 65535}
        _ = v28
    } else {
        var v29 int16 = -12
        _ = v29
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
