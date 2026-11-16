package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
func genFn0(p0 int32) uint16 {
    switch 325 {
    case -1:
        v6 := []int{42, 32768}
        _ = v6
    default:
        p0 = p0
    }
    switch "hello" {
    case "a":
        var v7 uint8 = 42
        _ = v7
    }
    p0 = (p0 + p0)
    var v8 int32 = (p0 >> ((1) & 63))
    _ = p0
    _ = v8
    return 65535
}
func genFn1(p0 uint) uint32 {
    switch -2147483648 {
    case 23:
        v9 := GenS1{F0: "\n\t", F1: 1000, F2: -9223372036854775808}
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!true) {
            v11 := GenS1{F0: "foo bar", F1: 401, F2: 0}
            _ = v11
        }
        if (1e10 > 0.1) {
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        if ((!(!((("\n\t" + ("foo bar" + ("" + "foo bar"))) + "") > "\n\t"))) || (fi10 == -660)) {
            var v13 float32 = (1e10 - (0.0 * -2.25))
            _ = v13
        }
        _ = fi10
    }
    if (-11 > 0) {
        var v14 string = ("" + "\n\t")
        v14 = v14
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 uint8 = 127
    var v2 int16 = 1000
    var v3 int8 = -1
    var v4 int8 = v3
    v5 := make(map[int]string)
    v15 := GenS2{F0: v3}
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 float32 = float32(0.0)
        _ = fi16
        _ = v17
    }
    v3 += (v4 / ((51) | 1))
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 string = ""
        var v20 uint16 = 297
        _ = fi18
        _ = v19
        _ = v20
    }
    v4 = (v4 >> ((255) & 63))
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint8 = v1
        _ = fi21
        _ = v22
    }
    v23 := v5[84]
    v24 := GenS3{F0: v4, F1: 1, F2: 904}
    fmt.Println(v3, v1, v3)
    v25 := GenS4{F0: (v2 / ((v2) | 1)), F1: 42, F2: 680}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
