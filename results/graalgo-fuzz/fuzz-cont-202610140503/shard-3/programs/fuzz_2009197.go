package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0() uint {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint64 = 0
        _ = fi5
        _ = v6
    }
    var v7 int32 = 2147483647
    _ = v7
    return 4294967295
}
func genFn1() bool {
    switch 128 {
    case 2147483647:
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int64 = -9223372036854775808
                _ = fi9
                _ = v10
            }
            var v11 int8 = -128
            var v12 bool = true
            _ = fi8
            _ = v11
            _ = v12
        }
    case 32768:
        var v13 int8 = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(fi14)
        v15 := []int{fi14, 657}
        var v16 uint64 = 2
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 int32 = 0
    fmt.Println(v17)
    v17 -= v17
    _ = v17
    return ((!false) || (!("a" <= "\n\t")))
}

func main() {
    v1 := []int{-1, 343}
    var v2 uint64 = 1
    var v3 int64 = -33
    v4 := []int{42, 0}
    var v18 uint = 2
    for fi19 := 0; fi19 < 5; fi19++ {
        fmt.Println(fi19)
        v2 = v2
        fmt.Println(v2)
        _ = fi19
    }
    v3 = v3
    v3 = -270
    var v20 uint = v18
    v21 := GenS1{F0: v2, F1: 106, F2: 2}
    fmt.Println(v3, v18, v3)
    switch 4294967295 {
    case -9223372036854775808:
        v22 := GenS1{F0: 18446744073709551615, F1: 65535, F2: uint8(v2)}
        _ = v22
    case -610:
        v1 = append(v1, -2147483648)
    }
    switch "a" {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v18 -= v20
        for fi24 := 0; fi24 < 18; fi24++ {
            fmt.Println(v20)
            if len(v1) > 0 {
                v1[0] = fi23
            }
            v3 += v3
            v18 = 2
            _ = fi24
        }
        _ = fi23
    }
    var v25 float32 = 1e10
    if false {
        var v26 string = "a"
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int8 = -92
            v20 = (v20 >> (((v20 >> (((v20 | 65535)) & 63))) & 63))
            fmt.Println(v20)
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    if (437 > 65535) {
        fmt.Println(v25, v3, v18)
    }
    var v29 uint8 = 23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
