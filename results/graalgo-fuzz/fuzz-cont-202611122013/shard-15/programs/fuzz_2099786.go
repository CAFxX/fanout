package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    var v3 int16 = 1
    for fi4 := 0; fi4 < 13; fi4++ {
        v1 += v1
        _ = fi4
    }
    var v5 int32 = 2147483647
    fmt.Println(v5, v1, v1)
    var v6 int64 = v1
    var v7 uint16 = 42
    v8 := GenS1{F0: int16(v7), F1: v7, F2: (-(v3))}
    var v9 float32 = 1e10
    var v10 int = -9223372036854775808
    if true {
        fmt.Println(v7)
    }
    if (v7 == v7) {
        switch "" {
        case "foo bar":
            v11 := GenS1{F0: v3, F1: 65535, F2: v3}
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            v3 += v3
            var v13 int8 = 0
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: v3, F1: v7, F2: -32768}
        _ = v14
    } else {
        v15 := GenS1{F0: v3, F1: v7, F2: (-(v3))}
        for fi16 := 0; fi16 < 16; fi16++ {
            switch v10 {
            case 4294967295:
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS2{F0: v10, F1: 924, F2: false}
                    v1 -= (v6 | (^(v1)))
                    var v19 uint8 = 42
                    if false {
                        v20 := []int{fi17, 9223372036854775807}
                        var v21 uint = 42
                        if len(v20) > 0 {
                            v20[0] = 1
                        }
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int32 = (v5 & v5)
                        _ = v22
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 32767:
                v10 = -2147483648
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    fmt.Println(v10, fi23, fi23)
                    if false {
                        if false {
                            v24 := GenS1{F0: 1, F1: 2, F2: (-(v3))}
                            v1 = int64(v5)
                            _ = v24
                        }
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                v6 -= (v1 % ((1) | 1))
                            }
                        case "hello":
                            var v25 string = "a"
                            _ = v25
                        }
                        var v26 uint32 = 439
                        _ = v26
                    }
                    _ = fi23
                }
            }
            v7 = v7
            if true {
                v27 := GenS1{F0: (v3 - (v3 | v3)), F1: (v7 * v7), F2: v3}
                var v28 bool = (false || ((true && (280 == 744)) || (v3 != (v3 >> ((0) & 63)))))
                _ = v27
                _ = v28
            }
            _ = fi16
        }
        _ = v15
    }
    if (!true) {
        delete(v2, (-(v10)))
    }
    fmt.Println(v9, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
