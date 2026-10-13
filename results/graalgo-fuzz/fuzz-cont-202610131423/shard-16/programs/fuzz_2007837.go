package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8) int32 {
    v4 := GenS1{F0: p0, F1: 4294967295, F2: 255}
    switch -2147483648 {
    case 2147483647:
        for fi5 := 0; fi5 < 18; fi5++ {
            if (-150 != 9223372036854775807) {
                var v6 uint = 640
                if true {
                    var v7 int8 = (p0 * p0)
                    v8 := GenS1{F0: p0, F1: 0, F2: (^(v6))}
                    _ = v7
                    _ = v8
                }
                _ = v6
            } else {
                fmt.Println(fi5, p0)
            }
            v9 := GenS2{F0: fi5, F1: 2}
            if (1 != 127) {
                p0 = (p0 >> ((78) & 63))
                if false {
                    fmt.Println(p0, fi5)
                } else {
                    fmt.Println(fi5, fi5)
                }
            } else {
                for fi10 := 0; fi10 < 2; fi10++ {
                    fmt.Println(p0, fi5, fi5)
                    p0 = -49
                    _ = fi10
                }
                v11 := GenS3{F0: -32768}
                _ = v11
            }
            _ = fi5
            _ = v9
        }
    case 657:
        if false {
            var v12 string = "a"
            v12 = v12
            _ = v12
        } else {
            p0 = p0
            var v13 float32 = ((0.1 - 0.1) / 0.1)
            _ = v13
        }
    default:
        var v14 uint8 = 255
        _ = v14
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 uint16 = 32767
    var v2 int = 668
    v3 := make(map[int]bool)
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS2{F0: (-(v2)), F1: 18446744073709551615}
        _ = fi15
        _ = v16
    }
    v1 = 255
    v17 := GenS1{F0: -128, F1: 4294967295, F2: 2}
    if false {
        v18 := GenS1{F0: -69, F1: uint64(v1), F2: 681}
        _ = v18
    } else {
        v19 := GenS2{F0: (-(v2)), F1: 4294967295}
        v20 := GenS2{F0: (v2 - 32768), F1: 351}
        _ = v19
        _ = v20
    }
    switch "a" {
    case "\n\t":
        var v21 int16 = -756
        _ = v21
    case "a":
        var v22 float64 = -2.25
        _ = v22
    }
    fmt.Println(v1)
    v23 := []int{}
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS2{F0: (v2 * (fi25 * 0)), F1: 21}
            _ = fi25
            _ = v26
        }
        v23 = append(v23, v2)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
