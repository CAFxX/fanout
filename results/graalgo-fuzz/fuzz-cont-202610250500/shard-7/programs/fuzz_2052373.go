package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int16; F2 string }
func genFn0() int64 {
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 1, F1: -855, F2: (("\n\t" + "foo bar") + ("a" + ("a" + "\n\t")))}
        _ = fi8
        _ = v9
    }
    _ = v7
    return -35
}

func main() {
    v1 := []string{("" + ("\n\t" + ("\n\t" + "")))}
    var v2 uint64 = 2
    v3 := make(map[uint32]bool)
    var v4 int8 = 42
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]string)
    v2 = v2
    v10 := GenS1{F0: 2147483648, F1: -1, F2: "foo bar"}
    var v11 int8 = v4
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := v3[2147483648]
        var v14 uint16 = 302
        var v15 int16 = 32767
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if true {
            switch fi16 {
            case -1:
                var v17 bool = false
                _ = v17
            case -9223372036854775808:
                v18 := GenS2{F0: v2, F1: uint8(v2), F2: 2147483648}
                _ = v18
            }
            var v19 float32 = 3.14159
            var v20 string = "hello"
            _ = v19
            _ = v20
        }
        _ = fi16
    }
    v4 -= -1
    var v21 bool = true
    v4 -= -128
    if true {
        var v22 uint32 = uint32(v2)
        for fi23 := 0; fi23 < 13; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                _ = len(v1)
                _ = fi24
            }
            var v25 float32 = (0.1 - -2.25)
            v26 := GenS1{F0: (^(v22)), F1: 0, F2: "a"}
            _ = fi23
            _ = v25
            _ = v26
        }
        _ = v22
    } else {
        if v21 {
            for fi27 := 0; fi27 < 4; fi27++ {
                _ = len(v1)
                _ = fi27
            }
        }
    }
    v28 := GenS3{F0: 1}
    switch -2147483648 {
    case -9223372036854775808:
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            var v29 int32 = -1
            _ = v29
        }
    case 32767:
        v30 := GenS4{F0: 1}
        _ = v30
    default:
        v5[("hello" + ("hello" + (("foo bar" + ("hello" + ("a" + "hello"))) + "hello")))] = 2147483647
    }
    v31 := GenS3{F0: 4294967295}
    v32 := v6[102]
    v33 := GenS5{F0: v21}
    v34 := GenS6{F0: 42, F1: -943, F2: v32}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
