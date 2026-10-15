package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 string, p2 int32) int32 {
    switch -1 {
    case -2147483648:
        switch -463 {
        case -541:
            for fi4 := 0; fi4 < 15; fi4++ {
                p0 = 0
                _ = fi4
            }
        default:
            var v5 int8 = -30
            _ = v5
        }
    case -361:
        var v6 uint16 = 285
        _ = v6
    }
    var v7 int8 = 1
    var v8 int8 = -44
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 % ((p2) | 1))
}

func main() {
    var v1 int64 = 945
    var v2 uint8 = 42
    var v3 int8 = -1
    v3 = v3
    v3 = v3
    fmt.Println(v2, v3, v3)
    v9 := GenS1{F0: 2147483647, F1: 663, F2: -963}
    if false {
        v10 := []int{-933, 9223372036854775807}
        fmt.Println(v3, v1, v3)
        switch "" {
        case "foo bar":
            v1 = (v1 ^ v1)
        case "a":
            for fi11 := 0; fi11 < 17; fi11++ {
                fmt.Println(v2, v1)
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch (fi11 * fi12) {
                    case -9223372036854775808:
                        fmt.Println(v3, v2)
                    case 127:
                        var v13 uint64 = 4294967295
                        _ = v13
                    }
                    switch fi11 {
                    case 1:
                        var v14 int = (fi11 - fi12)
                        _ = v14
                    case 0:
                        fmt.Println(v2, fi11, fi11)
                    case -1:
                        switch 128 {
                        case 917:
                            var v15 uint16 = 42
                            _ = v15
                        case 252:
                            var v16 uint8 = (v2 << ((2) & 63))
                            _ = v16
                        default:
                            if len(v10) > 0 {
                                v10[0] = 32767
                            }
                        }
                        v1 = (v1 | v1)
                        var v17 uint = 0
                        _ = v17
                    default:
                        fmt.Println(v3)
                    }
                    var v18 bool = true
                    _ = fi12
                    _ = v18
                }
                var v19 int16 = 0
                _ = fi11
                _ = v19
            }
        case "":
            var v20 uint32 = 4294967295
            _ = v20
        default:
            v21 := GenS2{F0: 65535}
            _ = v21
        }
        _ = v10
    } else {
        var v22 int32 = 42
        fmt.Println(v2, v2, v3)
        v23 := []int{-794}
        _ = v22
        _ = v23
    }
    var v24 int64 = v1
    var v25 uint8 = 127
    var v26 float32 = (float32(-2.25) + 1e10)
    for fi27 := 0; fi27 < 4; fi27++ {
        v24 = (v24 - int64(v25))
        v24 -= (v24 & v24)
        switch 1 {
        case 127:
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint32 = 1
                v1 -= -168
                _ = fi28
                _ = v29
            }
        case -1:
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := []int{}
                _ = fi30
                _ = v31
            }
        case 128:
            var v32 int8 = -33
            _ = v32
        }
        _ = fi27
    }
    v33 := []int{}
    v34 := GenS2{F0: 2}
    v1 -= v24
    v25 += (v2 - v2)
    if false {
        if false {
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            var v35 float32 = v26
            _ = v35
        }
        var v36 float32 = (1.5 - -2.25)
        _ = v36
    }
    fmt.Println(v2, v25)
    var v37 bool = (0 >= 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v9
}
