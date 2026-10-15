package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
type GenS3 struct { F0 uint16; F1 int64; F2 int32 }
func genFn0(p0 uint64) int64 {
    if false {
        v4 := []int{0}
        if ((-954 == 9223372036854775807) || false) {
            var v5 int16 = 1
            var v6 uint16 = 32767
            _ = v5
            _ = v6
        } else {
            fmt.Println(p0)
        }
        _ = v4
    } else {
        p0 -= p0
        p0 += (p0 ^ (p0 * (p0 / (((p0 ^ p0)) | 1))))
    }
    var v7 uint32 = 42
    v8 := GenS1{F0: -273}
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{(fi9 / ((32767) | 1)), (^(fi9))}
        var v11 uint = 42
        if false {
            switch "foo bar" {
            case "foo bar":
                switch "hello" {
                case "":
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := GenS2{F0: 672, F1: (!(!false)), F2: -307}
                        _ = fi12
                        _ = v13
                    }
                case "foo bar":
                    v10 = append(v10, (-(fi9)))
                case "a":
                    p0 += (p0 ^ p0)
                }
            default:
                fmt.Println(v11, v11, v7)
            }
            var v14 int64 = -478
            _ = v14
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v7)
    _ = p0
    _ = v7
    _ = v8
    return 820
}

func main() {
    var v1 float64 = (0.1 * 0.0)
    v2 := []int{4294967295}
    var v3 int = 32768
    var v15 string = ("" + "")
    v16 := GenS2{F0: 1000, F1: true, F2: 1}
    v3 = 1
    var v17 uint = 65535
    fmt.Println(v1, v17, v17)
    for fi18 := 0; fi18 < 12; fi18++ {
        switch ("hello" + ("hello" + ("" + v15))) {
        case "foo bar":
            v1 = v1
        case "":
            v15 = ""
        default:
            if len(v2) > 0 {
                v2[0] = 32768
            }
        }
        switch ((v15 + ((v15 + ((v15 + (v15 + "hello")) + ("hello" + v15))) + v15)) + v15) {
        case "hello":
            v15 = "\n\t"
        case "":
            if (!(42 == 240)) {
                if (false && false) {
                    v3 += (v3 >> (((v17 * (v17 >> (((v17 >> (((v17 % (((v17 << ((v17) & 63))) | 1))) & 63))) & 63)))) & 63))
                    var v19 string = "\n\t"
                    _ = v19
                } else {
                    var v20 uint16 = 2
                    _ = v20
                }
                switch "foo bar" {
                case "":
                    v21 := GenS3{F0: 394, F1: -9223372036854775808, F2: 0}
                    _ = v21
                }
                var v22 int32 = -1
                _ = v22
            }
        case "a":
            v23 := GenS2{F0: 32767, F1: (v1 > float64(1.5)), F2: 9223372036854775807}
            _ = v23
        }
        _ = fi18
    }
    v1 = v1
    v17 -= v17
    var v24 uint32 = 65535
    var v25 uint64 = 0
    for fi26 := 0; fi26 < 15; fi26++ {
        fmt.Println(fi26, v1, v3)
        for fi27 := 0; fi27 < 19; fi27++ {
            fmt.Println(fi27, v3, v3)
            v3 -= (v3 % ((v3) | 1))
            var v28 uint = 4294967295
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
