package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int16, p1 uint64, p2 int16) int32 {
    v6 := []int{-29, -1}
    v7 := GenS1{F0: 2147483647, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -242
}

func main() {
    var v1 float32 = (1.5 * (-2.25 - (float32((float64(1e10) + -2.25)) * 3.14159)))
    var v2 int = 4294967295
    var v3 bool = (-946 == 735)
    v4 := []string{"\n\t", "a", ""}
    var v5 float64 = 1e10
    if true {
        var v8 bool = (true || true)
        v3 = (65535 <= 2147483648)
        _ = v8
    } else {
        var v9 string = ("a" + "")
        _ = v9
    }
    fmt.Println(v2, v1, v1)
    v4 = append(v4, "hello")
    if (42 == 2) {
        for fi10 := 0; fi10 < 2; fi10++ {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            var v11 float32 = 0.0
            _ = len(v4)
            _ = fi10
            _ = v11
        }
    } else {
        switch v2 {
        case 128:
            v12 := GenS1{F0: 2147483648, F1: 42}
            _ = v12
        default:
            v2 = (v2 | v2)
        }
        switch ("" + "") {
        case "a":
            fmt.Println(v5, v2)
        }
        var v13 string = "hello"
        _ = v13
    }
    v3 = ((-1 <= 861) || v3)
    fmt.Println(v1, v5, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "foo bar"
        v5 = v5
        v5 = v5
        _ = fi15
        _ = v16
    }
    if (((-2.25 * v5) > float64(1e10)) && (!(!true))) {
        switch 4294967295 {
        case -627:
            v2 = v2
        case 128:
            switch "foo bar" {
            case "\n\t":
                var v17 uint8 = 175
                _ = v17
            default:
                v4 = append(v4, ("hello" + "\n\t"))
            }
        case 1:
            v18 := GenS1{F0: 65535, F1: -1}
            _ = v18
        }
    }
    var v19 uint16 = 188
    switch v2 {
    case 267:
        var v20 int16 = 1
        _ = v20
    case 32767:
        for fi21 := 0; fi21 < 20; fi21++ {
            v19 += (v19 / (((v19 * (v19 | 0))) | 1))
            _ = fi21
        }
    case -1:
        v22 := GenS2{F0: (0 < -3)}
        _ = v22
    default:
        if (18446744073709551615 <= 65535) {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS3{F0: 1, F1: (v2 % ((v2) | 1)), F2: v3}
                if (42 > 933) {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v26 uint8 = 222
                        var v27 bool = v3
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    var v28 uint32 = 2147483647
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
            if true {
                fmt.Println(v14)
                if v3 {
                    fmt.Println(v14, v1)
                }
                var v29 float64 = ((0.0 + 1.5) + 1.5)
                _ = v29
            } else {
                v30 := GenS2{F0: false}
                _ = v30
            }
        }
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
