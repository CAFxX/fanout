package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 bool; F2 int64 }
type GenS4 struct { F0 uint16 }
func genFn0() int8 {
    v6 := []int{32768, 9223372036854775807, 0, 655}
    switch 128 {
    case 128:
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: "foo bar", F1: 127, F2: 204}
            if (-842 >= 1) {
                if false {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        switch 9223372036854775807 {
                        case 0:
                            var v10 int16 = 1
                            _ = v10
                        case 127:
                            var v11 int16 = 42
                            _ = v11
                        default:
                            v12 := GenS1{F0: "\n\t", F1: -66, F2: 128}
                            _ = v12
                        }
                        _ = fi9
                    }
                    var v13 int8 = -54
                    _ = v13
                } else {
                    switch "hello" {
                    case "a":
                        _ = len(v6)
                    }
                }
            }
            _ = fi7
            _ = v8
        }
    case -9223372036854775808:
        var v14 uint64 = 2
        _ = v14
    default:
        switch -2147483648 {
        case -9223372036854775808:
            if (2147483648 >= 0) {
                var v15 uint64 = 42
                var v16 float64 = (float64(1e10) * 0.0)
                v17 := GenS2{F0: 0, F1: 0, F2: 1}
                _ = v15
                _ = v16
                _ = v17
            }
        case 281:
            var v18 bool = (0 < 42)
            _ = v18
        case 4294967295:
            var v19 int64 = 1
            _ = v19
        default:
            var v20 int16 = 1
            _ = v20
        }
    }
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 32767
    var v3 uint64 = 2
    v4 := make(map[string]uint32)
    var v5 bool = false
    var v21 uint32 = 216
    switch ("a" + "") {
    case "a":
        v22 := GenS3{F0: 1, F1: v5, F2: 0}
        _ = v22
    }
    fmt.Println(v3, v3, v3)
    switch 32768 {
    case 2147483647:
        if true {
            v5 = (!(!v5))
            for fi23 := 0; fi23 < 3; fi23++ {
                fmt.Println(v21, fi23, v21)
                var v24 bool = v5
                var v25 uint16 = (v2 / ((65535) | 1))
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            fmt.Println(v2, v3)
        }
    default:
        var v26 uint64 = (v3 % ((v3) | 1))
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := v1[-654]
        _ = fi27
        _ = v28
    }
    fmt.Println(v2)
    v29 := GenS2{F0: 42, F1: -1, F2: 0}
    var v30 int32 = 42
    v3 = (v3 | v3)
    v31 := GenS4{F0: 1}
    var v32 float64 = (float64(0.1) + float64(3.14159))
    v2 = v2
    var v33 int = 4294967295
    var v34 string = (("foo bar" + "") + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
