package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]int)
    v3 := []string{}
    v4 := []int{9223372036854775807}
    var v5 uint = 65535
    var v6 int16 = 32767
    v6 = v6
    var v7 float64 = 1e10
    var v8 int8 = -50
    if (0.0 < (v7 * v7)) {
        switch 127 {
        case 182:
            var v9 int8 = (v8 % ((1) | 1))
            _ = v9
        case -454:
            var v10 uint = v5
            _ = v10
        case 42:
            _ = len(v3)
        }
        v1 = (v1 & (v1 ^ v1))
        switch "" {
        case "hello":
            v4 = append(v4, 128)
        case "":
            v8 -= 100
        case "foo bar":
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint32 = 1
                _ = fi11
                _ = v12
            }
        default:
            v2[128] = -1
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: (v5 << (((^(v5))) & 63)), F1: 4294967295}
        var v15 uint64 = 42
        v7 = float64(1e10)
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if ((v5 | (^(v5))) < (v5 << ((4294967295) & 63))) {
            v3 = append(v3, ("a" + "\n\t"))
            if false {
                if false {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        _ = len(v4)
                        v18 := GenS2{F0: 4294967295, F1: 367}
                        v5 -= uint(v1)
                        v19 := GenS2{F0: 1, F1: -2147483648}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            }
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint16 = v1
                fmt.Println(v5, v21)
                _ = fi20
                _ = v21
            }
        }
        var v22 uint32 = 4294967295
        if false {
            v23 := GenS3{F0: v22, F1: -690}
            fmt.Println(v1, v1)
            var v24 string = "a"
            _ = v23
            _ = v24
        }
        _ = fi16
        _ = v22
    }
    var v25 float64 = float64(1.5)
    v7 = v7
    if false {
        _ = len(v2)
        var v26 uint8 = 1
        v27 := v2[-122]
        _ = v26
        _ = v27
    } else {
        v1 = v1
        v28 := v2[127]
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
