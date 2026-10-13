package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 uint32 = 0
    if true {
        if false {
            var v6 uint8 = 206
            var v7 string = ("foo bar" + "a")
            _ = v6
            _ = v7
        }
    } else {
        var v8 uint32 = v5
        _ = v8
    }
    var v9 float64 = 0.0
    var v10 uint = 4294967295
    switch "" {
    case "a":
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: v1, F1: "hello", F2: "foo bar"}
            v5 = 1
            var v13 int8 = 42
            _ = fi11
            _ = v12
            _ = v13
        }
    case "\n\t":
        var v14 float32 = 1e10
        _ = v14
    case "hello":
        var v15 uint16 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS1{F0: v1, F1: "foo bar", F2: (("hello" + "foo bar") + "")}
        if false {
            switch ("a" + ("\n\t" + "\n\t")) {
            case "hello":
                for fi18 := 0; fi18 < 2; fi18++ {
                    v19 := []int{(fi16 + fi16), (^(fi16))}
                    _ = fi18
                    _ = v19
                }
            default:
                v5 -= v5
            }
        }
        _ = fi16
        _ = v17
    }
    if false {
        if (42 > -822) {
            switch 0 {
            case 0:
                var v20 int8 = -47
                _ = v20
            default:
                for fi21 := 0; fi21 < 7; fi21++ {
                    v1 = v1
                    _ = fi21
                }
            }
            var v22 int8 = -128
            _ = v22
        }
    } else {
        v23 := []int{326}
        if (true || false) {
            switch -339 {
            case -262:
                v24 := GenS2{F0: true, F1: -62, F2: 1}
                _ = v24
            }
            fmt.Println(v1, v5, v5)
            v25 := v4[""]
            _ = v25
        } else {
            var v26 int32 = -2147483648
            v27 := GenS2{F0: false, F1: 100, F2: -1}
            _ = v26
            _ = v27
        }
        v9 += -2.25
        _ = v23
    }
    fmt.Println(v9)
    delete(v2, "\n\t")
    v28 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
