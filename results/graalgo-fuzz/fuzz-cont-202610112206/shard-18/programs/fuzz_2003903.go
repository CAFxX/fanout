package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0() bool {
    v4 := []int{}
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 bool = (122 > 255)
        _ = fi6
        _ = v7
    }
    _ = v4
    return true
}
func genFn1() uint8 {
    var v8 uint32 = 2147483647
    v9 := GenS1{F0: 10, F1: 32767, F2: -702}
    var v10 uint64 = 18446744073709551615
    switch "foo bar" {
    case "\n\t":
        if (!false) {
            v11 := GenS2{F0: "a", F1: -1}
            if true {
                for fi12 := 0; fi12 < 7; fi12++ {
                    fmt.Println(fi12, fi12)
                    _ = fi12
                }
                v10 -= (^(v10))
            }
            _ = v11
        }
    case "":
        v8 = v8
    case "a":
        switch 127 {
        case 127:
            v13 := GenS1{F0: 0, F1: 1, F2: 1}
            _ = v13
        case 32767:
            var v14 float32 = 0.1
            _ = v14
        }
    }
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int16 = -32768
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    v15 := []int{298, 82, -9223372036854775808}
    v2 = v2
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if (false || (float32(-2.25) >= -2.25)) {
                v18 := GenS1{F0: 113, F1: (fi16 * fi16), F2: (-(v1))}
                v19 := GenS3{F0: (true || (fi17 != 32767))}
                _ = v18
                _ = v19
            } else {
                var v20 string = ""
                _ = v20
            }
            _ = fi17
        }
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi16
    }
    if false {
        v2 -= v2
    } else {
        v21 := GenS2{F0: ("a" + "hello"), F1: -75}
        var v22 int32 = 2147483647
        if true {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v23 := v3[42]
            var v24 float32 = 3.14159
            _ = v23
            _ = v24
        }
        _ = v21
        _ = v22
    }
    var v25 int16 = (-(v1))
    v26 := GenS1{F0: 129, F1: 42, F2: (v1 + 42)}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println(v25, v1, v2)
    if (0 >= 451) {
        for fi27 := 0; fi27 < 16; fi27++ {
            _ = len(v15)
            v28 := GenS4{F0: 279, F1: 0, F2: 558}
            fmt.Println(v25)
            _ = fi27
            _ = v28
        }
        v25 = v25
    }
    var v29 uint64 = 18446744073709551615
    var v30 float64 = 1.5
    var v31 uint8 = 127
    v32 := GenS2{F0: ("foo bar" + "foo bar"), F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
