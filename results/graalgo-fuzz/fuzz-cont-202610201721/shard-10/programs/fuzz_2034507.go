package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
func genFn0() int8 {
    switch -9223372036854775808 {
    case 127:
        v4 := GenS1{F0: (true || false), F1: false, F2: 42}
        _ = v4
    case 128:
        var v5 float64 = float64(((1.5 * 0.1) + 1e10))
        _ = v5
    case 2147483647:
        var v6 uint64 = 0
        _ = v6
    default:
        var v7 float64 = float64((-2.25 * 1.5))
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint = 78
        for fi10 := 0; fi10 < 14; fi10++ {
            if (-1 != -1) {
                var v11 int16 = -32768
                if (!(!false)) {
                    var v12 int = (fi8 + fi10)
                    v13 := []int{}
                    _ = v12
                    _ = v13
                }
                _ = v11
            }
            v14 := GenS1{F0: true, F1: false, F2: -2147483648}
            v15 := []int{}
            var v16 uint32 = 2147483648
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi10
            _ = v14
            _ = v15
            _ = v16
        }
        v9 = (^(v9))
        _ = fi8
        _ = v9
    }
    var v17 int64 = 46
    v18 := []int{1, -9223372036854775808}
    var v19 float32 = (1.5 / 1.5)
    _ = v17
    _ = v18
    _ = v19
    return 127
}
func genFn1() string {
    var v20 uint64 = 37
    var v21 uint64 = (^(v20))
    fmt.Println(v21)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int16 = -1
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v20, v20, fi22)
            fmt.Println(v23)
            fmt.Println(fi22, v23, fi22)
            _ = fi24
        }
        v25 := GenS1{F0: (("" + (("\n\t" + "hello") + "a")) < ""), F1: (true || true), F2: 42}
        _ = fi22
        _ = v23
        _ = v25
    }
    _ = v20
    _ = v21
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 int = 42
    v2 := []int{-9223372036854775808, (v1 % ((32768) | 1)), v1}
    v3 := make(map[uint32]bool)
    v1 = v1
    v26 := GenS1{F0: true, F1: false, F2: -1}
    for fi27 := 0; fi27 < 17; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            if true {
                var v29 int64 = 1
                var v30 float32 = 1.5
                switch "foo bar" {
                case "foo bar":
                    switch "a" {
                    case "":
                        if len(v2) > 0 {
                            v2[0] = (^(fi28))
                        }
                    }
                default:
                    var v31 string = "foo bar"
                    _ = v31
                }
                _ = v29
                _ = v30
            }
            _ = fi28
        }
        var v32 float64 = 0.0
        v3[2147483647] = false
        _ = fi27
        _ = v32
    }
    var v33 float64 = 0.1
    fmt.Println(v33)
    v34 := v3[4294967295]
    _ = len(v2)
    var v35 int64 = 9223372036854775807
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
