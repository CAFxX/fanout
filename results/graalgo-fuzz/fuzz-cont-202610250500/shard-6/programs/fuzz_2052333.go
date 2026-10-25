package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
func genFn0() int64 {
    var v6 uint = 4294967295
    fmt.Println(v6)
    v6 -= (^(v6))
    var v7 string = "foo bar"
    var v8 uint8 = 42
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() string {
    if true {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint64 = 2
            _ = fi9
            _ = v10
        }
        switch -2147483648 {
        case 32768:
            v11 := GenS1{F0: "hello", F1: -472, F2: ("\n\t" + "")}
            _ = v11
        case -672:
            var v12 int8 = 100
            _ = v12
        default:
            v13 := []int{-2147483648, 0}
            _ = v13
        }
        var v14 uint = 4294967295
        _ = v14
    } else {
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(fi15)
            switch fi15 {
            case -812:
                v16 := GenS1{F0: "hello", F1: -1, F2: ("foo bar" + "a")}
                _ = v16
            }
            _ = fi15
        }
        var v17 uint8 = 19
        switch 1 {
        case 9223372036854775807:
            if (!false) {
                var v18 string = "a"
                _ = v18
            }
        case -2147483648:
            v17 = (v17 % ((v17) | 1))
        default:
            v19 := GenS1{F0: "\n\t", F1: 1, F2: ""}
            _ = v19
        }
        _ = v17
    }
    var v20 int64 = genFn0()
    for fi21 := 0; fi21 < 9; fi21++ {
        v20 += v20
        if (971 != 2147483648) {
            var v22 float32 = -2.25
            v23 := []int{127, fi21, fi21, fi21}
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    _ = v20
    return "foo bar"
}

func main() {
    var v1 uint64 = 1
    var v2 bool = (false || ((!true) || (0.0 <= 1.5)))
    var v3 float32 = float32(1e10)
    v4 := make(map[string]int)
    v5 := []string{}
    _ = len(v4)
    var v24 uint32 = 42
    switch -130 {
    case 0:
        fmt.Println(v1, v1, v3)
    }
    v5 = append(v5, "foo bar")
    var v25 float64 = ((3.14159 * (3.14159 + -2.25)) / 3.14159)
    switch 32767 {
    case 42:
        v1 = 0
    case 661:
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := GenS1{F0: ("a" + ""), F1: -9223372036854775808, F2: ("\n\t" + "hello")}
            var v28 float64 = v25
            _ = fi26
            _ = v27
            _ = v28
        }
    case 128:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v5 = append(v5, "")
    if v2 {
        v3 += v3
        if (271 != v1) {
            var v29 uint64 = uint64(v24)
            if ((v24 << ((255) & 63)) != (v24 ^ v24)) {
                if (v2 || ((9223372036854775807 == -2147483648) && v2)) {
                    var v30 uint64 = (v29 << ((255) & 63))
                    v31 := GenS1{F0: "", F1: -261, F2: ""}
                    var v32 float32 = v3
                    _ = v30
                    _ = v31
                    _ = v32
                }
            }
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
