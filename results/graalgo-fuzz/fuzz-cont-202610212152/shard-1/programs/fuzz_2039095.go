package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 bool }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 333
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{-9223372036854775808}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 0, F2: true}
    v7 := []int{2147483647, 524, 4294967295}
    for fi8 := 0; fi8 < 7; fi8++ {
        v3 = v3
        v3 = v3
        _ = fi8
    }
    var v9 float32 = float32(-2.25)
    v10 := GenS2{F0: 2, F1: true}
    var v11 uint64 = v3
    for fi12 := 0; fi12 < 7; fi12++ {
        v9 += v1
        var v13 int8 = 80
        v14 := v2[0]
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 bool = (-977 > -1)
    if len(v7) > 0 {
        v7[0] = -729
    }
    v15 = (1.5 == 1.5)
    v11 -= v11
    if (!(0.0 < (0.1 + (float64(v1) + (3.14159 + -2.25))))) {
        if (42 != 1) {
            fmt.Println(v11, v9)
        } else {
            var v16 uint32 = 65535
            v16 = v16
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            v11 += (v3 | v11)
            _ = fi17
        }
        fmt.Println(v9, v11)
    }
    var v18 int32 = 0
    delete(v2, 664)
    switch "hello" {
    case "":
        v1 -= float32(1e10)
    case "a":
        fmt.Println(v11)
    case "foo bar":
        if ((910 > 2147483647) || v15) {
            for fi19 := 0; fi19 < 19; fi19++ {
                v3 = (v11 + 4294967295)
                _ = fi19
            }
        } else {
            var v20 int = 63
            var v21 int = v20
            if ("hello" > "") {
                if false {
                    if v15 {
                        if true {
                            v20 -= (v20 << ((954) & 63))
                            v20 = (v20 ^ -390)
                            fmt.Println(v15)
                        }
                    } else {
                        var v22 int64 = 9223372036854775807
                        switch v20 {
                        case -1:
                            v23 := GenS1{F0: uint32(v3), F1: 42, F2: v15}
                            _ = v23
                        case 979:
                            for fi24 := 0; fi24 < 17; fi24++ {
                                var v25 int = v21
                                var v26 int16 = -32768
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        }
                        v20 += -295
                        _ = v22
                    }
                    fmt.Println(v1)
                }
                switch "hello" {
                case "hello":
                    var v27 uint64 = 4294967295
                    _ = v27
                case "foo bar":
                    _ = len(v2)
                default:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        fmt.Println(fi28, v3, v3)
                        var v29 int8 = 0
                        for fi30 := 0; fi30 < 8; fi30++ {
                            var v31 uint64 = v3
                            v11 += (v31 % ((1) | 1))
                            _ = fi30
                            _ = v31
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            }
            _ = v20
            _ = v21
        }
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        if ((2 > 1) || false) {
            v15 = ((fi32 << ((4294967295) & 63)) == int(v18))
        }
        _ = fi32
    }
    var v33 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v6
    _ = v7
    _ = v9
}
