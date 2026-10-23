package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 string; F2 string }
type GenS4 struct { F0 int64; F1 bool; F2 int }
type GenS5 struct { F0 bool }
func genFn0() uint64 {
    switch "\n\t" {
    case "":
        v5 := GenS1{F0: ((("foo bar" + "a") + "a") + "a")}
        _ = v5
    case "a":
        var v6 int64 = 1
        _ = v6
    default:
        for fi7 := 0; fi7 < 18; fi7++ {
            switch fi7 {
            case 896:
                var v8 int16 = 42
                _ = v8
            case 0:
                var v9 int16 = 32767
                _ = v9
            }
            if false {
                v10 := GenS1{F0: "foo bar"}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS1{F0: (("" + "a") + ("\n\t" + "\n\t"))}
                    v13 := GenS2{F0: (-(fi7))}
                    switch fi7 {
                    case -2147483648:
                        var v14 string = ((("a" + "") + ("a" + "a")) + (("\n\t" + "\n\t") + ("" + "a")))
                        _ = v14
                    default:
                        var v15 float32 = 3.14159
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v16 := GenS1{F0: (((("hello" + (("" + "a") + "\n\t")) + ("\n\t" + ("foo bar" + "\n\t"))) + "") + ("a" + ("foo bar" + "hello")))}
                _ = v16
            }
            _ = fi7
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 int8 = -1
        if (false && ((fi17 - 645) < fi17)) {
            v18 += (v18 ^ (v18 % ((53) | 1)))
            v19 := GenS3{F0: 194, F1: "", F2: ("" + ("a" + "foo bar"))}
            fmt.Println(v18, fi17)
            _ = v19
        }
        v20 := []int{-813, fi17, fi17, fi17}
        _ = fi17
        _ = v18
        _ = v20
    }
    var v21 uint64 = 767
    var v22 uint64 = (v21 | (v21 * 383))
    _ = v21
    _ = v22
    return v22
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    v4 := make(map[string]int)
    if (1e10 == 1.5) {
        v3 = v3
    }
    v23 := GenS2{F0: 42}
    v24 := []int{42, 0, -9223372036854775808}
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS4{F0: 42, F1: ((32767 >= 299) || true), F2: -665}
        _ = fi25
        _ = v26
    }
    if true {
        v24 = append(v24, -108)
        v1 -= 70
        fmt.Println(v1, v1)
    }
    v27 := GenS4{F0: -1, F1: (83 != 1), F2: 1}
    v28 := GenS5{F0: (128 >= 2)}
    fmt.Println(v3)
    var v29 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
