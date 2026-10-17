package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }

func main() {
    var v1 string = "a"
    v2 := []int{42, -2147483648}
    v3 := make(map[int]int)
    v4 := []string{(("foo bar" + ("a" + "\n\t")) + v1), v1, v1, (("a" + (v1 + "foo bar")) + v1)}
    var v5 int32 = 42
    var v6 uint32 = 2
    var v7 int64 = 391
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v7 -= -9223372036854775808
    switch "foo bar" {
    case "hello":
        v8 := GenS1{F0: 4294967295, F1: 42, F2: "a"}
        _ = v8
    case "a":
        v1 = (("hello" + ("hello" + v1)) + v1)
    default:
        var v9 uint8 = 77
        _ = v9
    }
    var v10 string = (("hello" + v1) + v1)
    if false {
        switch -2147483648 {
        case 32768:
            v11 := GenS1{F0: 18446744073709551615, F1: 449, F2: v1}
            _ = v11
        case -2147483648:
            switch ((("" + "") + ("\n\t" + "")) + "foo bar") {
            case "hello":
                v12 := GenS1{F0: 977, F1: 406, F2: v10}
                _ = v12
            case "a":
                for fi13 := 0; fi13 < 18; fi13++ {
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                    if true {
                        v10 = "foo bar"
                    } else {
                        var v14 float32 = -2.25
                        var v15 int = (^(fi13))
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
            case "\n\t":
                switch 2147483647 {
                case 656:
                    if (((float32(1.5) - (0.0 * (0.0 * -2.25))) < 1.5) && false) {
                        var v16 int8 = 42
                        _ = v16
                    }
                }
            default:
                var v17 int32 = v5
                _ = v17
            }
        default:
            fmt.Println(v10)
        }
        var v18 uint32 = 1
        v19 := v3[0]
        _ = v18
        _ = v19
    }
    v7 = (-(v7))
    var v20 int32 = (v5 & (v5 ^ 1))
    if true {
        v7 = -66
        var v21 int64 = (v7 & v7)
        for fi22 := 0; fi22 < 2; fi22++ {
            v21 = v21
            v23 := GenS1{F0: 18446744073709551615, F1: 2, F2: v1}
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                v5 = (v20 % ((v20) | 1))
                _ = fi25
            }
            _ = len(v3)
            _ = fi24
        }
        if (true || false) {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            switch 32767 {
            case 2147483647:
                v26 := GenS1{F0: 4294967295, F1: 2, F2: v1}
                _ = v26
            case -2147483648:
                if true {
                    fmt.Println(v7, v20)
                    var v27 bool = (255 >= 2)
                    _ = v27
                }
            case -379:
                v28 := GenS1{F0: 287, F1: uint64(v6), F2: ("" + ((v1 + (v10 + v1)) + "hello"))}
                _ = v28
            default:
                if (!false) {
                    v6 = v6
                }
            }
        }
        var v29 int32 = (^(v20))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
