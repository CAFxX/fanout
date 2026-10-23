package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0() uint64 {
    v6 := GenS1{F0: -1}
    v7 := GenS1{F0: 0}
    var v8 string = ("" + "\n\t")
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int) int {
    var v9 uint = 4294967295
    switch "a" {
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 bool = (uint32(v9) < 946)
            v12 := []int{(p0 | (p0 >> ((v9) & 63))), fi10, fi10, -2147483648}
            p0 = 0
            _ = fi10
            _ = v11
            _ = v12
        }
    case "foo bar":
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := []int{4294967295, p0, fi13, fi13}
            p0 = (p0 << (((v9 & 65535)) & 63))
            v9 = v9
            _ = fi13
            _ = v14
        }
    default:
        p0 = p0
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v9 = (v9 | 162)
        v16 := []int{(p0 % ((fi15) | 1)), (p0 >> ((v9) & 63)), fi15, (fi15 & (fi15 & p0))}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 100, F1: false, F2: "foo bar"}
    if false {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS3{F0: "hello", F1: 42}
            var v20 bool = (((0.1 + 0.0) != -2.25) || (v9 < (v9 << ((v9) & 63))))
            _ = fi18
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: "foo bar", F1: 0}
        _ = v21
    } else {
        v22 := []int{}
        var v23 float32 = (0.0 - -2.25)
        v9 = v9
        _ = v22
        _ = v23
    }
    _ = p0
    _ = v17
    _ = v9
    return 128
}

func main() {
    var v1 float32 = -2.25
    v2 := []int{}
    var v3 bool = (false || true)
    v4 := []string{}
    var v5 uint32 = 2
    v3 = false
    var v24 uint = 115
    if ((v3 && (668 >= (v5 | 0))) || (809 != 0)) {
        for fi25 := 0; fi25 < 17; fi25++ {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            v26 := GenS2{F0: 0, F1: (v3 || ("foo bar" < "hello")), F2: "a"}
            _ = fi25
            _ = v26
        }
        v3 = (!v3)
    } else {
        _ = len(v4)
    }
    v2 = append(v2, 42)
    v5 = (^(v5))
    var v27 uint64 = 4294967295
    v1 += ((-2.25 + -2.25) + (v1 / (v1 - ((v1 - float32(((float64(((v1 + 3.14159) + v1)) / 1.0) * float64(3.14159)))) / (v1 / v1)))))
    switch -128 {
    case 128:
        v28 := GenS4{F0: (^(v24))}
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v24 = (v24 * (v24 & (v24 & (v24 / ((v24) | 1)))))
        _ = fi29
    }
    v5 -= (^(v5))
    var v30 uint8 = uint8(v24)
    v1 = (((1e10 + v1) - float32(((0.1 / 3.14159) / 3.14159))) * v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v24 -= (v24 + 0)
    v30 += v30
    switch "hello" {
    case "":
        fmt.Println(v24, v27, v27)
    case "foo bar":
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 float32 = 1.5
            fmt.Println(v3)
            _ = fi31
            _ = v32
        }
    default:
        if v3 {
            var v33 int8 = 100
            _ = v33
        }
    }
    fmt.Println(v3, v27)
    var v34 uint32 = 968
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
