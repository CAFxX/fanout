package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 uint64 }
func genFn0() int8 {
    var v5 uint64 = 0
    v5 -= (v5 - 22)
    v5 = (v5 / ((v5) | 1))
    v5 = v5
    v5 = v5
    _ = v5
    return 104
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = 42
    var v3 uint8 = 255
    var v4 uint32 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        v1 = v1
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{(^(fi7))}
            if ((!((true && (!(582 < -580))) || ((v3 & (v3 << ((4294967295) & 63))) >= v3))) && false) {
                v9 := GenS1{F0: -1, F1: 207}
                if len(v8) > 0 {
                    v8[0] = fi7
                }
                _ = v9
            }
            var v10 float32 = 0.0
            _ = fi7
            _ = v10
            _ = v8
        }
        _ = fi6
    }
    switch v1 {
    case "hello":
        var v11 int64 = -1
        _ = v11
    }
    if true {
        v1 = v1
        v2 += v2
    } else {
        switch "hello" {
        case "foo bar":
            if true {
                var v12 float32 = 1.5
                for fi13 := 0; fi13 < 5; fi13++ {
                    if ((v3 >> ((0) & 63)) != (v3 << ((390) & 63))) {
                        v14 := GenS2{F0: (-(v2)), F1: v1, F2: (false && true)}
                        var v15 string = v1
                        fmt.Println(v12, v2, v3)
                        _ = v14
                        _ = v15
                    } else {
                        v16 := []int{fi13, (fi13 + fi13)}
                        _ = v16
                    }
                    fmt.Println(v4, fi13, v1)
                    _ = fi13
                }
                v17 := GenS2{F0: (v2 + -189), F1: (("hello" + ("a" + ("foo bar" + ("a" + v1)))) + ("" + (((("hello" + v1) + "") + (v1 + ((v1 + "\n\t") + v1))) + ("" + "\n\t")))), F2: (false || (false && (v2 == v2)))}
                _ = v12
                _ = v17
            } else {
                var v18 float64 = 0.1
                v19 := []int{0, 4294967295, -1}
                _ = v18
                _ = v19
            }
        }
        v20 := GenS2{F0: 1, F1: (("" + "hello") + v1), F2: false}
        _ = v20
    }
    var v21 int32 = v2
    for fi22 := 0; fi22 < 11; fi22++ {
        v21 = v21
        v2 = (v21 | (v2 * 1))
        _ = fi22
    }
    var v23 bool = true
    fmt.Println(v1, v1)
    if (!(-1 != 598)) {
        var v24 int = 574
        _ = v24
    }
    v25 := GenS3{F0: v1, F1: 2}
    switch v1 {
    case "hello":
        var v26 int = 127
        _ = v26
    default:
        var v27 int = -165
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v4
}
