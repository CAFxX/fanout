package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 uint8, p1 uint64) uint64 {
    switch 76 {
    case -673:
        v6 := GenS1{F0: true, F1: -527, F2: (("\n\t" + "hello") + "hello")}
        _ = v6
    case 4294967295:
        v7 := []int{-267}
        _ = v7
    default:
        p0 -= p0
    }
    if true {
        p1 -= 141
        var v8 uint = 65535
        if (!true) {
            fmt.Println(p1)
            v9 := GenS1{F0: false, F1: 1, F2: "hello"}
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 float64 = -2.25
    var v4 float64 = v3
    var v5 uint16 = 682
    v10 := GenS1{F0: (genFn0(0, 2) > 752), F1: -212, F2: ("a" + "hello")}
    switch "hello" {
    case "a":
        v11 := GenS2{F0: 0, F1: 1000}
        _ = v11
    case "hello":
        if false {
            var v12 uint8 = 128
            _ = v12
        }
    case "foo bar":
        switch "a" {
        case "foo bar":
            v3 -= (((0.0 / v4) + (((v3 * float64(1e10)) - v4) * 3.14159)) / float64(float32((v4 + 0.0))))
        default:
            for fi13 := 0; fi13 < 7; fi13++ {
                v5 = (^(v5))
                v1[492] = 0
                _ = fi13
            }
        }
    }
    var v14 int16 = 0
    if (!("\n\t" <= "foo bar")) {
        v15 := GenS3{F0: 4294967295, F1: genFn0(128, 42), F2: 150}
        v16 := GenS4{F0: v5}
        v14 = v14
        _ = v15
        _ = v16
    }
    if (true && ((-2.25 < float32((v3 / 1e10))) || true)) {
        var v17 uint64 = 2
        var v18 int = -9223372036854775808
        for fi19 := 0; fi19 < 3; fi19++ {
            _ = len(v2)
            v20 := GenS4{F0: v5}
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    }
    if (((32767 != (v5 / ((970) | 1))) || (!("" >= (("hello" + "") + "\n\t")))) || (0 <= v14)) {
        var v21 uint16 = (v5 ^ 2)
        if (0 != 2) {
            var v22 int8 = 127
            _ = v22
        }
        _ = v21
    }
    v14 = v14
    v3 = v3
    switch -9223372036854775808 {
    case -900:
        var v23 int32 = -1
        _ = v23
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := v2["hello"]
            v26 := GenS1{F0: (2147483647 >= 971), F1: 693, F2: "hello"}
            var v27 int8 = 42
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case "a":
        v3 -= v4
    case "hello":
        v28 := GenS5{F0: 1, F1: ((1e10 / (v3 * v3)) < ((3.14159 * -2.25) + (float64((1.5 / 1.5)) * 0.1)))}
        _ = v28
    default:
        for fi29 := 0; fi29 < 17; fi29++ {
            v4 = v3
            fmt.Println(v3)
            fmt.Println(v14, v4)
            _ = fi29
        }
    }
    v30 := GenS1{F0: false, F1: 0, F2: "a"}
    var v31 uint16 = 866
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
