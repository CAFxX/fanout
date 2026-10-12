package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16) uint8 {
    v4 := GenS1{F0: "foo bar"}
    p0 = 6
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 int8, p1 uint32, p2 int8) bool {
    v5 := []int{-1, 876}
    p0 -= (p2 ^ p0)
    _ = len(v5)
    var v6 float32 = (-2.25 - ((((-2.25 / float32(1.5)) * float32(-2.25)) - 0.1) - float32(1.5)))
    if len(v5) > 0 {
        v5[0] = 36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (!true)
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 29
    var v3 uint8 = 128
    if false {
        for fi7 := 0; fi7 < 2; fi7++ {
            v3 = (v3 ^ (v3 ^ v3))
            v3 -= v3
            _ = fi7
        }
        v8 := v1["foo bar"]
        var v9 float32 = 0.1
        _ = v8
        _ = v9
    } else {
        if ("foo bar" == "\n\t") {
            v10 := GenS1{F0: "a"}
            fmt.Println(v2, v3)
            _ = v10
        }
        fmt.Println(v2, v3)
    }
    if (789 != 18446744073709551615) {
        delete(v1, "")
        v2 -= (v2 ^ (^(v2)))
        v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
    }
    delete(v1, "")
    v2 = (v2 * v2)
    var v11 int = 47
    var v12 int = 0
    if (!true) {
        v13 := GenS2{F0: v11, F1: v2, F2: ("hello" + "hello")}
        _ = v13
    }
    var v14 float32 = (3.14159 * 1.5)
    var v15 uint = 742
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v2)
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 float32 = v14
            _ = fi17
            _ = v18
        }
        _ = len(v1)
        _ = fi16
    }
    if ("hello" == ((((("\n\t" + "foo bar") + "") + "foo bar") + "\n\t") + "a")) {
        v2 -= int8(v15)
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 uint8 = uint8(v15)
                v22 := GenS2{F0: (v11 - v11), F1: 65, F2: "hello"}
                switch ("" + "hello") {
                case "\n\t":
                    var v23 uint16 = 42
                    _ = v23
                case "hello":
                    switch "a" {
                    case "hello":
                        if (((1.5 - 0.1) == 0.1) || (!false)) {
                            if true {
                                var v24 int8 = v2
                                _ = v24
                            }
                        }
                    }
                case "":
                    var v25 uint = v15
                    _ = v25
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            if true {
                v1["a"] = 42
                v26 := GenS3{F0: 2}
                _ = v26
            }
            _ = fi19
        }
        var v27 uint = 73
        _ = v27
    } else {
        switch v12 {
        case 9223372036854775807:
            v2 = (-(v2))
        case 42:
            v2 += (v2 >> ((v15) & 63))
        default:
            fmt.Println(v14, v12, v15)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
