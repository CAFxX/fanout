package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
func genFn0() bool {
    var v5 string = ""
    v6 := []int{}
    var v7 int = 4294967295
    if false {
        v7 += v7
    }
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 bool = false
    var v2 string = "hello"
    v3 := []string{v2, "a", ((v2 + (v2 + (((v2 + ((v2 + (v2 + ("foo bar" + v2))) + ("a" + v2))) + v2) + v2))) + "a")}
    v4 := []uint32{2147483648, 2, 4294967295}
    if v1 {
        var v8 uint = 2
        if (v1 || v1) {
            v4 = append(v4, 42)
            v8 = v8
        } else {
            v9 := GenS1{F0: -1, F1: 42, F2: v8}
            if (!false) {
                v8 += v8
                v4 = append(v4, 4294967295)
            }
            if (1 < 32767) {
                var v10 int64 = -9223372036854775808
                var v11 int = -2147483648
                _ = v10
                _ = v11
            }
            _ = v9
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        if (!(!(-2.25 < float32((-2.25 / float64((0.1 / 1.5))))))) {
            switch v2 {
            case "":
                v1 = v1
            case "foo bar":
                fmt.Println(fi12)
            default:
                var v13 uint = 42
                _ = v13
            }
            var v14 float64 = 1e10
            var v15 int = fi12
            _ = v14
            _ = v15
        }
        var v16 int8 = 100
        _ = fi12
        _ = v16
    }
    switch "hello" {
    case "a":
        switch "foo bar" {
        case "a":
            var v17 int = 128
            _ = v17
        case "":
            switch v2 {
            case "hello":
                var v18 uint8 = 0
                _ = v18
            default:
                for fi19 := 0; fi19 < 9; fi19++ {
                    v2 = (v2 + ("foo bar" + ("\n\t" + v2)))
                    var v20 bool = v1
                    for fi21 := 0; fi21 < 6; fi21++ {
                        fmt.Println(fi19, v2)
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
            }
        default:
            v3 = append(v3, "foo bar")
        }
    case "":
        v2 = v2
    case "hello":
        var v22 int16 = -1
        _ = v22
    }
    var v23 int32 = 42
    fmt.Println(v2)
    _ = len(v4)
    v1 = (422 >= 4294967295)
    v2 = "\n\t"
    var v24 uint32 = 42
    fmt.Println(v23, v24, v1)
    var v25 int8 = 0
    var v26 string = v2
    for fi27 := 0; fi27 < 2; fi27++ {
        v3 = append(v3, v2)
        v24 = 372
        var v28 int8 = v25
        _ = fi27
        _ = v28
    }
    if v1 {
        _ = len(v4)
    }
    if ((v25 & 1) <= (v25 / ((-128) | 1))) {
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 = v2
            _ = fi29
        }
    } else {
        for fi30 := 0; fi30 < 4; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                fmt.Println(v23, fi31, v1)
                _ = fi31
            }
            v1 = v1
            v4 = append(v4, v24)
            _ = fi30
        }
    }
    var v32 string = v2
    v33 := GenS2{F0: (^(v24)), F1: 4294967295, F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
