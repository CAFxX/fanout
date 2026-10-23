package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 int32) int64 {
    p0 += -650
    v5 := []int{42, -9223372036854775808, -1, 32767}
    var v6 int32 = (p1 - (-(p1)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 685
}
func genFn1(p0 uint) int {
    fmt.Println(p0)
    var v7 int = -9223372036854775808
    var v8 string = "hello"
    _ = p0
    _ = v7
    _ = v8
    return -692
}

func main() {
    var v1 uint8 = 1
    var v2 int = 1
    var v3 int16 = -468
    var v4 int64 = -827
    v2 = -890
    if (true || true) {
        var v9 bool = false
        fmt.Println(v1)
        if (!(((v4 + v4) < v4) && (!(!v9)))) {
            v10 := []int{}
            if v9 {
                v2 = v2
                v10 = append(v10, (v2 % (((^(v2))) | 1)))
            }
            _ = v10
        }
        _ = v9
    } else {
        var v11 float32 = 1e10
        var v12 int64 = v4
        v13 := GenS1{F0: 247}
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 string = ("a" + "hello")
    var v15 int64 = (v4 | (v4 - (v4 ^ 338)))
    v16 := []int{v2, genFn1(255)}
    switch genFn1(1) {
    case 357:
        v3 += 42
    case 4294967295:
        for fi17 := 0; fi17 < 18; fi17++ {
            v2 += 32767
            v3 = v3
            _ = fi17
        }
    case -1:
        var v18 uint8 = v1
        _ = v18
    }
    _ = len(v16)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint8 = 179
        _ = fi19
        _ = v20
    }
    if false {
        if true {
            if len(v16) > 0 {
                v16[0] = v2
            }
        }
        for fi21 := 0; fi21 < 13; fi21++ {
            v14 = ("" + (v14 + v14))
            switch v14 {
            case "":
                var v22 string = v14
                _ = v22
            }
            _ = fi21
        }
    }
    switch "hello" {
    case "hello":
        v23 := GenS2{F0: 100}
        _ = v23
    case "a":
        var v24 int32 = -777
        _ = v24
    case "foo bar":
        switch "" {
        case "\n\t":
            fmt.Println(v15, v14)
        case "foo bar":
            v2 = v2
        }
    default:
        for fi25 := 0; fi25 < 4; fi25++ {
            v4 = v4
            _ = fi25
        }
    }
    if (genFn1(42) > v2) {
        fmt.Println(v15)
        v1 -= v1
    } else {
        v16 = append(v16, v2)
    }
    var v26 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
