package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint8) int32 {
    p0 -= (p0 % ((p0) | 1))
    var v4 bool = false
    var v5 uint16 = 2
    var v6 float64 = float64(1.5)
    var v7 int32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 0
    var v3 int16 = 42
    v1 = 0
    fmt.Println(v2)
    if (genFn0(128) >= (v2 << ((0) & 63))) {
        var v8 uint16 = 27
        var v9 int64 = (v1 >> ((4294967295) & 63))
        var v10 uint64 = 64
        _ = v10
        _ = v8
        _ = v9
    } else {
        switch -386 {
        case 777:
            fmt.Println(v2, v3)
        case -91:
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint = 1
                _ = fi11
                _ = v12
            }
        case 822:
            v13 := GenS1{F0: v1, F1: (true || true)}
            _ = v13
        }
    }
    var v14 float32 = float32(1.5)
    v14 = ((v14 + v14) * v14)
    switch "foo bar" {
    case "":
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := []int{-2147483648}
            v1 -= v1
            v14 = v14
            _ = fi15
            _ = v16
        }
    case "hello":
        var v17 uint16 = 42
        _ = v17
    }
    if true {
        for fi18 := 0; fi18 < 17; fi18++ {
            v3 -= v3
            _ = fi18
        }
        v3 = v3
    } else {
        fmt.Println(v1, v14)
        for fi19 := 0; fi19 < 7; fi19++ {
            v14 = (v14 * (3.14159 * ((v14 - v14) * (float32(1e10) / 1.5))))
            v20 := GenS2{F0: 42, F1: 1}
            _ = fi19
            _ = v20
        }
        switch ("\n\t" + ("a" + "foo bar")) {
        case "hello":
            v21 := GenS3{F0: 197, F1: 182}
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v3 += -1
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int16 = (^(v3))
            switch (fi22 / (((fi23 << ((0) & 63))) | 1)) {
            case 32768:
                v1 = (v1 % (((-(v1))) | 1))
            }
            v25 := []int{fi22, fi23, (^(fi23))}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    v3 = v3
    for fi26 := 0; fi26 < 13; fi26++ {
        v2 -= (v2 / (((^(v2))) | 1))
        v14 += v14
        _ = fi26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint8 = 128
        v2 = v2
        _ = fi27
        _ = v28
    }
    var v29 uint8 = 130
    v29 -= v29
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
}
