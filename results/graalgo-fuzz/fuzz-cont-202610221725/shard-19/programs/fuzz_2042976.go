package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int32, p1 uint, p2 int) int8 {
    p1 += (p1 - (p1 % (((p1 << ((p1) & 63))) | 1)))
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4, p2)
        v5 := GenS1{F0: 2, F1: 744, F2: -9223372036854775808}
        _ = fi4
        _ = v5
    }
    var v6 int64 = -9223372036854775808
    v7 := GenS2{F0: p1, F1: 111}
    p1 += (p1 + (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 79
}
func genFn1(p0 uint16) uint16 {
    p0 += p0
    v8 := []int{127, 32768}
    var v9 bool = false
    var v10 int64 = 9223372036854775807
    v11 := GenS3{F0: 525}
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (p0 << ((18446744073709551615) & 63))
}

func main() {
    var v1 int8 = -72
    v2 := make(map[string]string)
    var v3 uint32 = 2
    v1 += (^(v1))
    var v12 uint8 = 119
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float64 = 0.0
        var v15 uint8 = (v12 * (^(v12)))
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 = v3
    v12 = v12
    if false {
        v1 += (v1 - v1)
        v1 += v1
        var v16 int = 32768
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint64 = 4294967295
        v2["hello"] = ("foo bar" + "foo bar")
        _ = fi17
        _ = v18
    }
    delete(v2, "a")
    var v19 int16 = 897
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 = (v3 / ((v3) | 1))
        switch (^(fi20)) {
        case 127:
            if true {
                var v21 uint8 = (v12 + 1)
                _ = v21
            } else {
                v22 := []int{}
                switch ((("a" + "a") + "a") + "") {
                case "\n\t":
                    switch ("hello" + "a") {
                    case "foo bar":
                        v3 += uint32(v12)
                    default:
                        v23 := v2[("foo bar" + "a")]
                        _ = v23
                    }
                }
                v24 := GenS4{F0: "hello", F1: false}
                _ = v22
                _ = v24
            }
        default:
            v25 := GenS2{F0: 0, F1: 250}
            _ = v25
        }
        _ = fi20
    }
    v1 = (v1 + (v1 >> ((1) & 63)))
    var v26 int = -884
    v26 += v26
    fmt.Println(v12, v3, v3)
    v27 := GenS1{F0: 18446744073709551615, F1: 1, F2: -9223372036854775808}
    var v28 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
