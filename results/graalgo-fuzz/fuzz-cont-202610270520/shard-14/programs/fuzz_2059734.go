package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint64; F1 uint64 }
func genFn0() int16 {
    v6 := []int{127, -13, 4294967295, 208}
    var v7 uint = 125
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 1
    var v2 int32 = 820
    var v3 uint = v1
    var v4 bool = (2147483648 < 1)
    var v5 float32 = (0.0 * (1.5 - 3.14159))
    v2 += v2
    switch 32767 {
    case 4294967295:
        var v8 bool = v4
        _ = v8
    case 42:
        v2 -= v2
    default:
        v9 := GenS1{F0: 127, F1: 127, F2: v2}
        _ = v9
    }
    var v10 float32 = float32(0.1)
    switch ("foo bar" + ("\n\t" + "a")) {
    case "hello":
        var v11 uint32 = 2147483648
        _ = v11
    case "":
        var v12 int = 42
        _ = v12
    case "\n\t":
        v13 := []int{}
        _ = v13
    default:
        v5 = v5
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 bool = (false && ((fi14 + 2147483647) < 32768))
        _ = fi14
        _ = v15
    }
    var v16 int8 = -13
    v1 -= (v3 & (v1 + v3))
    v17 := GenS1{F0: 2147483647, F1: v16, F2: v2}
    if (!((v4 || (32767 >= -639)) || v4)) {
        v18 := GenS2{F0: -32768, F1: ""}
        for fi19 := 0; fi19 < 19; fi19++ {
            fmt.Println(v2, v3)
            _ = fi19
        }
        v20 := GenS2{F0: 1000, F1: ("" + "foo bar")}
        _ = v18
        _ = v20
    } else {
        v21 := GenS3{F0: v16, F1: v1}
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 uint16 = 0
            switch "" {
            case "\n\t":
                var v24 string = "hello"
                _ = v24
            case "a":
                v5 -= (v5 - -2.25)
            case "foo bar":
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := []int{(-(fi25)), (fi25 << ((v3) & 63)), (fi22 - (fi22 & (fi25 ^ fi22)))}
                    fmt.Println(v4, v5)
                    v5 = 3.14159
                    _ = fi25
                    _ = v26
                }
            }
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint32 = 686
                _ = fi27
                _ = v28
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    if (1 >= 32767) {
        var v29 int64 = -1
        fmt.Println(v5, v2, v10)
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 bool = (255 <= 0)
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    var v32 uint16 = 0
    var v33 int = -1
    switch (("hello" + "") + ("a" + "hello")) {
    case "\n\t":
        if v4 {
            v34 := []int{}
            v35 := GenS4{F0: 977, F1: 0}
            v10 = v10
            _ = v34
            _ = v35
        } else {
            v2 -= v2
            v36 := []int{v33, (v33 + v33), v33, (-(v33))}
            v3 = 4294967295
            _ = v36
        }
    case "a":
        v37 := []int{32768}
        _ = v37
    default:
        v38 := GenS1{F0: (v33 >> (((v1 % (((v1 * (v3 + (v1 & (v3 * 0))))) | 1))) & 63)), F1: (v16 ^ v16), F2: (^(v2))}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
