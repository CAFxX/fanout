package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
func genFn0() uint {
    switch -9223372036854775808 {
    case -341:
        var v4 uint = 962
        _ = v4
    case 127:
        switch "foo bar" {
        case "hello":
            var v5 float32 = -2.25
            _ = v5
        case "a":
            var v6 int16 = -32768
            _ = v6
        default:
            if true {
                var v7 int64 = -1
                var v8 uint16 = 65535
                v9 := []int{}
                _ = v7
                _ = v8
                _ = v9
            } else {
                v10 := GenS1{F0: 32767, F1: 0}
                v11 := []int{32768, 127, -230, -1}
                var v12 int32 = -2147483648
                _ = v10
                _ = v11
                _ = v12
            }
        }
    case 2147483647:
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 float32 = 3.14159
            var v15 float32 = ((((float32(-2.25) + 0.1) / 1e10) - (v14 / 1.5)) - v14)
            var v16 uint64 = 42
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    if (619 <= 817) {
        v17 := []int{32768}
        _ = v17
    } else {
        var v18 float32 = 0.0
        v19 := []int{388, 9223372036854775807}
        _ = v18
        _ = v19
    }
    var v20 int16 = -32768
    var v21 string = ("hello" + (("foo bar" + ("a" + ("foo bar" + "foo bar"))) + "a"))
    v20 = (-(v20))
    _ = v20
    _ = v21
    return 18446744073709551615
}
func genFn1(p0 uint32, p1 uint64) bool {
    var v22 float32 = 3.14159
    p1 = (p1 % (((p1 >> ((4294967295) & 63))) | 1))
    v23 := []int{0, -1, -1, -316}
    v24 := GenS1{F0: 0, F1: 1}
    switch 236 {
    case -1:
        fmt.Println(v22)
    case 32767:
        var v25 int32 = -2147483648
        _ = v25
    default:
        v26 := GenS2{F0: 226}
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    return true
}

func main() {
    var v1 uint16 = 0
    var v2 uint16 = 385
    var v3 bool = (!((2147483647 == 2147483648) || (42 <= 2)))
    v1 = (v1 % (((^(v2))) | 1))
    v3 = ((v1 + v1) > (v2 | (v2 ^ v2)))
    var v27 uint8 = 42
    v27 = 97
    for fi28 := 0; fi28 < 9; fi28++ {
        fmt.Println(v2, v1, v27)
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int = (^(fi29))
            v30 -= (v30 % (((v30 << ((433) & 63))) | 1))
            v31 := GenS1{F0: 979, F1: 32767}
            v32 := GenS3{F0: -986, F1: 557, F2: 1}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        for fi34 := 0; fi34 < 11; fi34++ {
            if v3 {
                v35 := GenS1{F0: 2, F1: 0}
                _ = v35
            }
            _ = fi34
        }
        var v36 uint = 368
        v27 = (^(v27))
        _ = fi33
        _ = v36
    }
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 float64 = (1.5 + -2.25)
        _ = fi37
        _ = v38
    }
    v39 := []int{9223372036854775807, -479, 582}
    fmt.Println(v27, v3)
    v40 := GenS3{F0: 1, F1: -2147483648, F2: 1}
    if (32767 > 0) {
        _ = len(v39)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v40
}
