package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
func genFn0(p0 bool) uint64 {
    v5 := GenS1{F0: "a", F1: 244, F2: 1}
    switch 127 {
    case 42:
        for fi6 := 0; fi6 < 10; fi6++ {
            switch "" {
            case "foo bar":
                v7 := []int{(fi6 - fi6)}
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: (("foo bar" + "hello") + (("\n\t" + "\n\t") + "\n\t")), F1: 2, F2: 42}
                _ = v8
            }
            _ = fi6
        }
    }
    p0 = p0
    _ = p0
    _ = v5
    return 4294967295
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 0
    var v3 bool = (780 <= 18446744073709551615)
    var v4 uint64 = 2
    v9 := GenS1{F0: "a", F1: v2, F2: 715}
    v10 := GenS1{F0: v1, F1: v2, F2: -2147483648}
    v11 := []int{-226, 1}
    v12 := GenS1{F0: "", F1: v2, F2: 1}
    switch -9223372036854775808 {
    case 121:
        var v13 uint64 = 688
        _ = v13
    case -9223372036854775808:
        for fi14 := 0; fi14 < 19; fi14++ {
            if ((v3 || true) && v3) {
                for fi15 := 0; fi15 < 17; fi15++ {
                    v4 = v4
                    var v16 string = v1
                    var v17 uint32 = 983
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                var v18 int64 = 90
                v18 = (v18 - (v18 / (((-(v18))) | 1)))
                _ = v18
            }
            _ = fi14
        }
    case 0:
        v4 = genFn0((2147483647 >= 698))
    default:
        var v19 int16 = 32767
        _ = v19
    }
    var v20 string = ("" + ("" + v1))
    var v21 int16 = -1
    v20 = v1
    fmt.Println(v3, v21, v1)
    v4 = (v4 << ((55) & 63))
    var v22 uint64 = 2
    var v23 float64 = 0.0
    v24 := GenS1{F0: ("hello" + v20), F1: (v2 + v2), F2: 2147483647}
    v1 = ""
    if len(v11) > 0 {
        v11[0] = -9223372036854775808
    }
    fmt.Println(v22, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
