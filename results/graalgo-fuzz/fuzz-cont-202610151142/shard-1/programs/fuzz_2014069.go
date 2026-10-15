package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0() uint64 {
    switch "hello" {
    case "":
        switch 340 {
        case 42:
            for fi7 := 0; fi7 < 2; fi7++ {
                if (990 <= 820) {
                    v8 := []int{(fi7 % (((fi7 << ((171) & 63))) | 1)), fi7, (fi7 >> ((1) & 63))}
                    _ = v8
                } else {
                    var v9 int64 = 1
                    _ = v9
                }
                var v10 uint64 = 838
                _ = fi7
                _ = v10
            }
        case 4294967295:
            switch "a" {
            case "a":
                var v11 uint16 = 2
                _ = v11
            case "":
                for fi12 := 0; fi12 < 12; fi12++ {
                    if true {
                        var v13 uint = 763
                        if true {
                            fmt.Println(v13, v13)
                        }
                        _ = v13
                    }
                    var v14 float64 = 1.5
                    v14 -= (float64(float32(v14)) / float64(1.5))
                    _ = fi12
                    _ = v14
                }
            }
        case 1:
            var v15 uint32 = 4294967295
            _ = v15
        default:
            var v16 int8 = 4
            _ = v16
        }
    case "hello":
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
    }
    var v19 uint = 378
    var v20 int = -1
    if true {
        switch ((("a" + ("hello" + "\n\t")) + ("" + "a")) + "foo bar") {
        case "":
            switch 817 {
            case 218:
                v21 := []int{}
                _ = v21
            default:
                fmt.Println(v19, v19)
            }
        case "a":
            v22 := []int{v20, -176, (v20 >> ((v19) & 63))}
            _ = v22
        }
        var v23 uint = v19
        v20 = v20
        _ = v23
    }
    if (!(0.1 < 0.1)) {
        fmt.Println(v19, v19)
    }
    _ = v19
    _ = v20
    return 206
}

func main() {
    v1 := []int{2147483647}
    var v2 int64 = 0
    v3 := make(map[string]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 uint32 = 2147483648
    v6 := make(map[int]uint32)
    v5 = (v5 / ((v5) | 1))
    v2 = int64(v4)
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int16 = -1
            var v27 uint16 = 2
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 uint64 = (^(v4))
        _ = fi24
        _ = v28
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v4 = v4
    case 2147483647:
        var v29 uint = 18446744073709551615
        _ = v29
    }
    if true {
        var v30 float32 = 1.5
        _ = v30
    }
    var v31 uint64 = genFn0()
    v2 = (v2 - v2)
    for fi32 := 0; fi32 < 9; fi32++ {
        v3["foo bar"] = v5
        v5 -= (^(v5))
        _ = fi32
    }
    v33 := GenS1{F0: "hello", F1: 42, F2: (v31 * v4)}
    var v34 int = 2147483647
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 bool = true
        v37 := GenS1{F0: "hello", F1: 255, F2: (v31 / ((v31) | 1))}
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println(v34, v31, v5)
    v2 -= -1
    switch ("a" + "\n\t") {
    case "hello":
        v4 = v31
    }
    _ = len(v3)
    v31 = (v4 >> ((586) & 63))
    var v38 uint32 = v5
    delete(v6, (v34 >> ((18446744073709551615) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
