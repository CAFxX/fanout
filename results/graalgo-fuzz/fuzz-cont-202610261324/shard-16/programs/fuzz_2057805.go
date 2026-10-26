package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
func genFn0() uint {
    switch "foo bar" {
    case "foo bar":
        if true {
            var v5 int = -1
            if false {
                v5 += v5
            } else {
                for fi6 := 0; fi6 < 3; fi6++ {
                    v5 -= (fi6 % ((v5) | 1))
                    _ = fi6
                }
                for fi7 := 0; fi7 < 6; fi7++ {
                    fmt.Println(fi7)
                    _ = fi7
                }
                var v8 int16 = 42
                _ = v8
            }
            _ = v5
        } else {
            var v9 int64 = -9223372036854775808
            _ = v9
        }
    case "a":
        var v10 uint32 = 0
        _ = v10
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int8 = 86
            v12 += v12
            switch "hello" {
            case "":
                var v13 string = "foo bar"
                _ = v13
            case "\n\t":
                v12 += (v12 / ((v12) | 1))
            }
            _ = fi11
            _ = v12
        }
    }
    var v14 uint = 42
    var v15 int16 = -1
    _ = v14
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{65535, 4294967295, 42, 4294967295}
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v3[683] = 628
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint64 = 42
        var v18 uint64 = v17
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 2
    if false {
        v19 -= v19
        if false {
            v20 := GenS1{F0: (0 != 0), F1: 0, F2: 0}
            _ = v20
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            v22 := GenS1{F0: false, F1: 32767, F2: -1}
            _ = fi21
            _ = v22
        }
    }
    switch "a" {
    case "foo bar":
        v23 := GenS1{F0: (((!false) && (42 <= 0)) && true), F1: 0, F2: -9223372036854775808}
        _ = v23
    case "a":
        v19 = (^(v19))
    case "\n\t":
        var v24 uint8 = 127
        _ = v24
    }
    if false {
        v3[65535] = 65535
    } else {
        var v25 uint32 = 1
        v19 = v19
        _ = v25
    }
    fmt.Println(v19)
    var v26 float32 = 1e10
    fmt.Println(v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
