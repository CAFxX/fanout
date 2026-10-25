package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []uint32{327, 2147483648, 699, 590}
    var v2 int32 = 451
    v3 := []int{1, 315, -830, 42}
    v4 := []int{733, 855, 9223372036854775807, 1}
    var v5 int16 = 0
    v6 := GenS1{F0: 4294967295, F1: -222, F2: ("" + ("hello" + ("a" + ((("a" + "hello") + "foo bar") + ((("foo bar" + ("a" + "")) + "foo bar") + "a")))))}
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
    }
    v5 -= 32767
    fmt.Println(v2)
    if true {
        v2 = v2
    }
    var v7 float64 = 3.14159
    fmt.Println(v2, v5, v5)
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 2, F1: 346, F2: "\n\t"}
        v7 = float64(0.1)
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 312, F1: 2147483648}
    v11 := GenS3{F0: 746, F1: v5, F2: ((v7 != 0.0) || false)}
    var v12 int8 = 69
    if true {
        v13 := GenS4{F0: (^(v12))}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            v12 += (v12 ^ v12)
            var v15 float64 = v7
            for fi16 := 0; fi16 < 17; fi16++ {
                if len(v4) > 0 {
                    v4[0] = 21
                }
                if (4294967295 > 1) {
                    v17 := GenS1{F0: 0, F1: -156, F2: "a"}
                    _ = v17
                } else {
                    fmt.Println(v7)
                    v7 = v15
                    if true {
                        if false {
                            var v18 uint8 = 0
                            var v19 bool = ((fi14 ^ 936) > (^(fi14)))
                            _ = v18
                            _ = v19
                        }
                        var v20 uint8 = 1
                        _ = v20
                    }
                }
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        fmt.Println(v12)
        var v21 float32 = float32(v7)
        _ = v21
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
