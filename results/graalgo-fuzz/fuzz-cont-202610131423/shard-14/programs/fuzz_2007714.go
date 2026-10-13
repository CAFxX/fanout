package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int; F2 bool }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int16; F1 int16; F2 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int; F1 uint8 }
func genFn0() int64 {
    v5 := GenS1{F0: -97}
    if (42 != 566) {
        v6 := []int{903, -9223372036854775808, 336, -9223372036854775808}
        var v7 uint8 = 1
        if true {
            v6 = append(v6, 42)
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(fi8, fi8)
                _ = fi8
            }
            fmt.Println(v7)
        }
        _ = v6
        _ = v7
    }
    var v9 float64 = 3.14159
    var v10 int = -2147483648
    _ = v10
    _ = v5
    _ = v9
    return 402
}

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS2{F0: 18446744073709551615, F1: -564}
        if true {
            if (true || false) {
                if false {
                    if false {
                        if (255 > v2) {
                            fmt.Println(fi11)
                            var v13 int32 = -499
                            v2 = v2
                            _ = v13
                        }
                        var v14 uint32 = 42
                        _ = v14
                    }
                    var v15 float64 = 1e10
                    v16 := GenS3{F0: 65535}
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint8 = (v2 ^ v2)
                    v18 := []int{128, fi11}
                    _ = v17
                    _ = v18
                }
                fmt.Println(v2)
            }
            v19 := []int{fi11, (^(fi11))}
            _ = v19
        } else {
            var v20 string = "a"
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            fmt.Println(fi21)
            _ = fi21
        }
        _ = fi11
        _ = v12
    }
    if (!(!(127 <= 127))) {
        var v22 uint64 = 4294967295
        switch "" {
        case "":
            v23 := GenS4{F0: v1, F1: 127, F2: true}
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 5; fi24++ {
                var v25 uint8 = 0
                v2 = (^(v25))
                if (true && false) {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        fmt.Println(v25)
                        var v27 int32 = -2147483648
                        v2 = v25
                        _ = fi26
                        _ = v27
                    }
                    v1 = (^(v1))
                    fmt.Println(v25, v22, v25)
                }
                _ = fi24
                _ = v25
            }
        case "hello":
            if false {
                for fi28 := 0; fi28 < 16; fi28++ {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        var v30 int64 = 1
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
                fmt.Println(v2, v22)
                var v31 uint = 18446744073709551615
                _ = v31
            }
        default:
            var v32 float32 = (-2.25 * 0.0)
            _ = v32
        }
        var v33 uint16 = v1
        _ = v22
        _ = v33
    }
    v2 -= 2
    var v34 int32 = 53
    v34 = v34
    v35 := GenS5{F0: 127, F1: -2147483648}
    _ = len(v3)
    v36 := GenS6{F0: 1, F1: 1000, F2: 32767}
    for fi37 := 0; fi37 < 15; fi37++ {
        v38 := GenS7{F0: -128}
        v39 := []int{fi37}
        v40 := GenS8{F0: fi37, F1: v2}
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
