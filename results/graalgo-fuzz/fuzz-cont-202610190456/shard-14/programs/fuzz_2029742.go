package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }
func genFn0(p0 int16, p1 int, p2 uint64) uint32 {
    var v7 int = 244
    v8 := []int{}
    var v9 bool = false
    v9 = ((v7 - v7) > p1)
    for fi10 := 0; fi10 < 8; fi10++ {
        if (1e10 >= (1e10 * -2.25)) {
            var v11 uint8 = 127
            var v12 int8 = 100
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v9 = (1 > 1)
            _ = fi13
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 864
}
func genFn1() int64 {
    v14 := GenS1{F0: (2 <= 1), F1: 0, F2: 1}
    var v15 string = "foo bar"
    switch 128 {
    case 128:
        v16 := []int{94, 128, 42}
        _ = v16
    case 32767:
        v17 := GenS2{F0: -765, F1: (((v15 + (v15 + v15)) + "hello") + v15)}
        _ = v17
    case 42:
        var v18 uint = 833
        _ = v18
    default:
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                v15 = v15
                var v21 float64 = ((0.0 / 1e10) - (-2.25 + ((1e10 + ((1.5 - -2.25) - 1e10)) + 1e10)))
                _ = fi20
                _ = v21
            }
            v15 = ("\n\t" + "hello")
            v15 = ("\n\t" + (("foo bar" + v15) + v15))
            _ = fi19
        }
    }
    v15 = v15
    v22 := GenS3{F0: 244}
    _ = v14
    _ = v15
    _ = v22
    return 0
}

func main() {
    v1 := []int{4294967295, 32768}
    var v2 string = "foo bar"
    v3 := []int{-98, -2147483648}
    var v4 uint8 = 255
    var v5 int64 = 839
    var v6 uint64 = 848
    for fi23 := 0; fi23 < 6; fi23++ {
        v4 -= (v4 << ((1) & 63))
        _ = fi23
    }
    v4 = (v4 % ((v4) | 1))
    if true {
        var v24 float32 = ((-2.25 - 1e10) / -2.25)
        fmt.Println(v6)
        _ = v24
    } else {
        var v25 string = (v2 + ((v2 + ("a" + "hello")) + ("\n\t" + (v2 + ("\n\t" + v2)))))
        switch "foo bar" {
        case "foo bar":
            var v26 string = "foo bar"
            _ = v26
        case "\n\t":
            v6 = v6
        case "":
            v6 -= 1
        default:
            var v27 float64 = 1e10
            _ = v27
        }
        _ = v25
    }
    var v28 uint16 = 0
    v2 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = ((v2 + ((v2 + ((("foo bar" + "a") + v2) + (v2 + v2))) + ("" + ((v2 + "") + v2)))) + v2)
    v3 = append(v3, -1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS3{F0: 32767}
        var v31 uint64 = v6
        _ = fi29
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: "a"}
    v33 := GenS5{F0: true, F1: 444, F2: (v5 / (((v5 & (v5 % ((v5) | 1)))) | 1))}
    v6 -= v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
