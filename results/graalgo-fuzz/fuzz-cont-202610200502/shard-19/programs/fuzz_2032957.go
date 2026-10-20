package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 string) bool {
    p0 = "hello"
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 0, F1: 0, F2: 255}
        p0 = (p0 + "\n\t")
        _ = fi7
        _ = v8
    }
    if false {
        p0 = p0
        switch -412 {
        case 178:
            v9 := GenS1{F0: 934, F1: 0, F2: 503}
            _ = v9
        case 127:
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := []int{(fi10 >> ((18446744073709551615) & 63)), (fi10 & (^(fi10))), (^(fi10)), (-(fi10))}
                _ = fi10
                _ = v11
            }
        default:
            fmt.Println(p0)
        }
    }
    var v12 string = (p0 + (p0 + ""))
    _ = p0
    _ = v12
    return ((((1.5 / 1.0) == 0.0) && (172 < 223)) && true)
}

func main() {
    var v1 float64 = -2.25
    var v2 int16 = -32768
    var v3 int64 = 1
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    var v6 uint = 463
    v3 -= (v3 | v3)
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS2{F0: (("a" + "hello") + "hello")}
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: 1, F1: 42}
    var v16 int16 = (v2 - (v2 ^ (v2 + (v2 << (((v6 | 0)) & 63)))))
    v1 = (3.14159 - (-2.25 - v1))
    fmt.Println(v6, v6, v2)
    v1 = v1
    for fi17 := 0; fi17 < 8; fi17++ {
        switch ("a" + "a") {
        case "a":
            var v18 uint = 42
            _ = v18
        }
        _ = fi17
    }
    var v19 int8 = 42
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 int16 = (^(v16))
        v22 := GenS4{F0: (v19 - (v19 & (-(v19))))}
        _ = len(v5)
        _ = fi20
        _ = v21
        _ = v22
    }
    v19 = 127
    var v23 int = 128
    switch (("hello" + ("\n\t" + "\n\t")) + "a") {
    case "a":
        if ((v6 & (v6 / ((v6) | 1))) > (v6 | (v6 / ((0) | 1)))) {
            if false {
                fmt.Println(v16, v6)
                var v24 uint64 = 18446744073709551615
                switch "foo bar" {
                case "\n\t":
                    switch "hello" {
                    case "a":
                        if ((!((v3 / (((v3 / (((-(v3))) | 1))) | 1)) > (-(v3)))) || (65535 > 32767)) {
                            var v25 int = (v23 % (((v23 + v23)) | 1))
                            _ = v25
                        } else {
                            v26 := []int{int(v19), v23, -930}
                            v1 = 3.14159
                            v23 -= int(v16)
                            _ = v26
                        }
                    }
                case "foo bar":
                    v27 := GenS5{F0: v6, F1: (("a" + ("foo bar" + ("" + (("a" + "foo bar") + "")))) + ("a" + "a"))}
                    _ = v27
                }
                _ = v24
            }
            v28 := GenS4{F0: int8(v3)}
            for fi29 := 0; fi29 < 4; fi29++ {
                for fi30 := 0; fi30 < 12; fi30++ {
                    fmt.Println(v23, fi29, fi29)
                    _ = fi30
                }
                _ = fi29
            }
            _ = v28
        }
    case "hello":
        var v31 int8 = int8(v6)
        _ = v31
    default:
        fmt.Println(v2, v16)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
