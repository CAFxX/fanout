package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int16; F1 int; F2 string }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 bool; F2 int }

func main() {
    var v1 float32 = (0.0 - (float32(3.14159) - 1.5))
    var v2 int32 = -401
    var v3 float64 = 0.1
    var v4 int16 = -32768
    var v5 float64 = v3
    var v6 float64 = 3.14159
    v7 := GenS1{F0: ((("\n\t" + "foo bar") + "") + ""), F1: 1, F2: -672}
    v8 := []int{}
    var v9 uint = 2
    v6 -= (v5 + v6)
    if true {
        fmt.Println(v3, v9, v9)
        if len(v8) > 0 {
            v8[0] = 1
        }
    }
    v6 = (float64(3.14159) + ((((v5 * v5) + (v3 + (((float64(float32(0.0)) - v6) + ((v5 * v6) + v5)) + v5))) / float64(0.1)) - (v5 * float64(v1))))
    if true {
        v6 -= 0.1
    } else {
        if (127 == 114) {
            v3 -= (float64(float32((v5 - 1e10))) - v3)
        } else {
            var v10 int64 = 68
            v2 = v2
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            if (2 < 22) {
                if true {
                    v12 := GenS2{F0: 1}
                    var v13 int16 = (^(v4))
                    _ = v12
                    _ = v13
                }
                v2 += v2
                v9 -= (v9 ^ v9)
            }
            _ = fi11
        }
    }
    _ = len(v8)
    switch "" {
    case "foo bar":
        var v14 uint8 = 79
        _ = v14
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 float32 = 1.5
        if false {
            v17 := GenS2{F0: 1}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 bool = ((v2 << ((4294967295) & 63)) == 0)
    v19 := GenS1{F0: "a", F1: uint32(v9), F2: 1}
    v20 := GenS3{F0: 42, F1: 52}
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if true {
                var v23 int32 = (v2 + (^(v2)))
                v18 = v18
                if v18 {
                    v24 := GenS3{F0: 2, F1: 2}
                    v25 := GenS2{F0: -1}
                    v26 := GenS4{F0: (v4 << ((v9) & 63)), F1: fi22, F2: ""}
                    _ = v24
                    _ = v25
                    _ = v26
                }
                _ = v23
            }
            _ = fi22
        }
        v27 := GenS5{F0: fi21, F1: (-2.25 == -2.25)}
        v28 := GenS6{F0: 42}
        _ = fi21
        _ = v27
        _ = v28
    }
    v29 := GenS7{F0: "foo bar", F1: false, F2: -434}
    switch "" {
    case "foo bar":
        v1 = v1
    case "\n\t":
        _ = len(v8)
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v5 = 3.14159
        var v31 int8 = -1
        v9 -= (v9 ^ (^(v9)))
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
