package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int8; F1 bool; F2 int8 }
func genFn0() string {
    v4 := []int{0, 127, -2147483648}
    var v5 uint32 = 932
    v5 = 564
    v6 := GenS1{F0: 42, F1: 0}
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 int32 = -961
    fmt.Println(v3)
    _ = len(v1)
    if (29 != -9) {
        v3 += v3
        v7 := GenS2{F0: 2}
        _ = v7
    } else {
        _ = len(v1)
        if ((-968 <= -345) || ((-1 >= 9223372036854775807) || true)) {
            if true {
                if ((true || (-1 <= 695)) || false) {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        v1[42] = ((1e10 >= (float32((float64(1.5) - (1e10 * float64((-2.25 * (float32(3.14159) * 1.5)))))) - (float32((((3.14159 + 3.14159) - (float64(1.5) - -2.25)) - 0.0)) * 0.1))) && (1 >= 65535))
                        v3 = v3
                        var v9 uint64 = 0
                        _ = fi8
                        _ = v9
                    }
                }
                switch "foo bar" {
                case "":
                    var v10 uint = 65535
                    _ = v10
                }
            }
        } else {
            if true {
                v11 := GenS3{F0: 42}
                v12 := GenS3{F0: 9223372036854775807}
                _ = v11
                _ = v12
            } else {
                v13 := []int{-287, 928, 1}
                fmt.Println(v3)
                v3 -= v3
                _ = v13
            }
            if true {
                var v14 int8 = -19
                _ = v14
            }
        }
        if (432 <= 417) {
            v15 := GenS4{F0: -119, F1: v3}
            var v16 uint64 = 1
            _ = v15
            _ = v16
        } else {
            v17 := []int{-9223372036854775808}
            _ = v17
        }
    }
    var v18 uint = 66
    for fi19 := 0; fi19 < 20; fi19++ {
        fmt.Println(v18)
        _ = len(v2)
        _ = fi19
    }
    var v20 uint = (v18 ^ v18)
    v21 := GenS4{F0: -9223372036854775808, F1: v3}
    v22 := GenS5{F0: 100, F1: false, F2: 1}
    var v23 uint16 = 0
    var v24 int32 = 0
    var v25 int8 = 127
    if ((true && (225 < 0)) && (0 >= 128)) {
        if (0 != (v25 << ((v20) & 63))) {
            v20 = v18
        } else {
            for fi26 := 0; fi26 < 7; fi26++ {
                v24 = (v3 - int32(v25))
                var v27 int8 = v25
                v28 := GenS3{F0: -9223372036854775808}
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        v24 = v24
        var v29 bool = false
        _ = v29
    } else {
        v18 -= (v18 >> (((^(v18))) & 63))
    }
    var v30 int16 = -569
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS2{F0: 716}
        _ = fi31
        _ = v32
    }
    v23 = v23
    fmt.Println(v30)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
