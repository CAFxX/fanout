package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 string; F1 int32 }
func genFn0(p0 int32, p1 string, p2 int32) uint {
    var v4 int32 = p2
    v5 := GenS1{F0: true, F1: p1, F2: 42}
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            v7 := GenS1{F0: true, F1: p1, F2: 86}
            v8 := []int{fi6, fi6, 9223372036854775807, fi6}
            _ = v7
            _ = v8
        }
        p1 = ""
        v9 := []int{}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        switch ("hello" + "") {
        case "hello":
            v11 := GenS2{F0: true, F1: 471}
            _ = v11
        }
        var v12 bool = (!(1000 == 42))
        var v13 int8 = -128
        _ = fi10
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 uint8, p1 uint64) uint {
    v14 := []int{2147483647, 401, 9223372036854775807}
    for fi15 := 0; fi15 < 1; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float64 = ((-2.25 * (1.5 / 1e10)) + 0.0)
            switch 895 {
            case 32768:
                var v18 bool = true
                _ = v18
            }
            v17 += v17
            v19 := GenS3{F0: false, F1: (p0 * 187)}
            _ = fi16
            _ = v17
            _ = v19
        }
        v20 := GenS4{F0: 0, F1: (fi15 - (fi15 % ((-959) | 1)))}
        var v21 uint32 = 0
        _ = fi15
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    v1 := []int{32767, -485, -897}
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    var v22 float32 = (-2.25 - 0.1)
    v2 = append(v2, "")
    v22 = (((float32(float64(v22)) - (((0.1 * v22) - -2.25) * 1.5)) * 1e10) / (3.14159 - v22))
    v22 = v22
    if (480 == 42) {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        if ("a" <= "\n\t") {
            var v23 string = "\n\t"
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 string = v23
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            v22 = (float32(((-2.25 * float64(0.0)) / 1.0)) * v22)
        }
        if false {
            v26 := GenS5{F0: "a", F1: 979}
            _ = v26
        }
    } else {
        v22 = 0.1
        v22 = (3.14159 + 1.5)
    }
    switch "hello" {
    case "foo bar":
        v27 := v3[2]
        _ = v27
    case "hello":
        var v28 int = -124
        _ = v28
    }
    switch 186 {
    case 32767:
        v22 = 1.5
    case 42:
        if true {
            var v29 uint16 = 255
            _ = v29
        }
    default:
        fmt.Println(v22)
    }
    if len(v2) > 0 {
        v2[0] = ("a" + ("a" + "hello"))
    }
    if (false || ((!true) || false)) {
        fmt.Println(v22)
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        v31 := v3[42]
        var v32 int64 = -1
        v33 := v3[(v31 ^ (v31 - (v31 + (v31 >> ((genFn1(42, 918)) & 63)))))]
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    v22 = v22
    v22 = (v22 + v22)
    var v34 uint64 = 42
    var v35 bool = true
    v36 := GenS1{F0: (!(42 < 42)), F1: (("foo bar" + "hello") + "foo bar"), F2: 63}
    v1 = append(v1, -534)
    var v37 bool = (((("\n\t" + "\n\t") + ("foo bar" + ("" + ""))) + "a") != ("" + ("\n\t" + "\n\t")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
