package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32; F2 int8 }
type GenS6 struct { F0 uint8; F1 uint; F2 string }
type GenS7 struct { F0 int32; F1 string; F2 string }

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    var v3 int8 = -128
    var v4 int = -1
    v3 += (^(v3))
    var v5 uint16 = 42
    switch int(v5) {
    case 2147483647:
        v6 := GenS1{F0: (v4 ^ v4), F1: 42}
        _ = v6
    case -2147483648:
        v7 := GenS2{F0: 255, F1: (v1 - -1)}
        _ = v7
    }
    var v8 bool = false
    v8 = v8
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = (1e10 - 0.1)
        var v11 string = (("hello" + "\n\t") + "")
        fmt.Println(fi9, v1)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint32 = 2147483648
    switch "" {
    case "foo bar":
        v13 := []int{(v4 >> ((0) & 63)), (v4 << ((426) & 63)), v4}
        _ = v13
    case "hello":
        v14 := GenS2{F0: 0, F1: v1}
        _ = v14
    default:
        v15 := GenS1{F0: v4, F1: 1000}
        _ = v15
    }
    v12 -= (v12 + v12)
    var v16 int64 = 0
    fmt.Println(v1, v3, v1)
    var v17 int8 = -44
    switch v4 {
    case 42:
        v8 = false
    case -2147483648:
        v18 := GenS3{F0: 109}
        _ = v18
    case -1:
        if v8 {
            if true {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v12 = 871
                    v16 += 0
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 int16 = -322
                        fmt.Println(v21, v5, fi19)
                        v22 := GenS4{F0: v5, F1: -2}
                        if v8 {
                            v23 := []int{}
                            var v24 int16 = (v21 & (^(v21)))
                            _ = v23
                            _ = v24
                        } else {
                            if (145 <= 255) {
                                v25 := GenS5{F0: 0, F1: -881, F2: (v3 ^ 100)}
                                if (!v8) {
                                    if (true && true) {
                                        v26 := []int{}
                                        v27 := GenS1{F0: 4294967295, F1: v21}
                                        _ = v26
                                        _ = v27
                                    }
                                } else {
                                    fmt.Println(v5, v5)
                                }
                                v8 = v8
                                _ = v25
                            }
                            v28 := GenS6{F0: 2, F1: 56, F2: "\n\t"}
                            if v8 {
                                v8 = v8
                                var v29 float32 = 0.0
                                _ = v29
                            } else {
                                var v30 float32 = ((1e10 + 1e10) * ((0.0 / -2.25) * 0.0))
                                var v31 string = ("a" + "hello")
                                _ = v30
                                _ = v31
                            }
                            _ = v28
                        }
                        fmt.Println(fi19, v16)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                }
                fmt.Println(v2)
            }
            var v32 float32 = (0.0 - 3.14159)
            v33 := GenS7{F0: v2, F1: "a", F2: "foo bar"}
            _ = v32
            _ = v33
        } else {
            v34 := []int{32767}
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
