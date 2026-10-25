package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 bool }
type GenS3 struct { F0 string }
func genFn0(p0 uint) uint32 {
    p0 -= p0
    p0 = (^(p0))
    v6 := []int{127, 48, 799, -291}
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 int16) uint16 {
    fmt.Println(p0)
    if true {
        v7 := GenS1{F0: 40}
        if (761 <= 319) {
            switch "hello" {
            case "a":
                var v8 int16 = p0
                _ = v8
            case "":
                var v9 float32 = 1e10
                _ = v9
            }
        }
        switch ("hello" + "a") {
        case "hello":
            p0 -= 808
        case "\n\t":
            v10 := GenS2{F0: int(p0), F1: 1, F2: (65535 >= 4294967295)}
            _ = v10
        default:
            fmt.Println(p0)
        }
        _ = v7
    } else {
        p0 += p0
    }
    v11 := []int{32768, 0, 42}
    if (!(true || true)) {
        v12 := GenS2{F0: 128, F1: 255, F2: true}
        var v13 float64 = 3.14159
        _ = v12
        _ = v13
    } else {
        var v14 uint32 = 1
        p0 = -777
        _ = v14
    }
    _ = p0
    _ = v11
    return 1
}

func main() {
    var v1 int8 = 27
    var v2 int32 = -2147483648
    var v3 int16 = -846
    v4 := make(map[uint32]string)
    var v5 float32 = -2.25
    v15 := []int{}
    fmt.Println(v2, v1)
    switch 2147483647 {
    case 186:
        v16 := GenS3{F0: "foo bar"}
        _ = v16
    case 1:
        for fi17 := 0; fi17 < 13; fi17++ {
            switch (-(fi17)) {
            case -172:
                fmt.Println(v2)
            case 1:
                v3 += v3
            default:
                _ = len(v15)
            }
            _ = fi17
        }
    default:
        var v18 int8 = v1
        _ = v18
    }
    v19 := GenS3{F0: (("a" + "a") + "a")}
    var v20 float64 = -2.25
    v3 += v3
    v2 += v2
    var v21 uint64 = 0
    _ = len(v15)
    var v22 int64 = 1
    fmt.Println(v1, v3)
    fmt.Println(v2, v20)
    v20 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
