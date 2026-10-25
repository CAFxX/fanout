package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
func genFn0(p0 int32, p1 uint, p2 int32) int {
    v5 := []int{}
    if false {
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                var v7 int64 = -931
                fmt.Println(v7, v7, p0)
                for fi8 := 0; fi8 < 7; fi8++ {
                    var v9 int16 = 481
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
            switch "\n\t" {
            case "a":
                var v10 uint16 = 906
                _ = v10
            case "foo bar":
                p1 = p1
            default:
                p2 += p2
            }
        }
        if (127 == 255) {
            if false {
                v5 = append(v5, 4294967295)
            } else {
                p0 = p2
            }
            fmt.Println(p2, p2)
        }
        switch ("foo bar" + "a") {
        case "a":
            if false {
                p1 += (^(p1))
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -2147483648
}
func genFn1(p0 uint16) int64 {
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int16 = -1
            v12 -= v12
            switch -186 {
            case 1:
                v13 := []int{fi11, fi11, (fi11 >> ((255) & 63))}
                _ = v13
            case -1:
                var v14 int64 = -9223372036854775808
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    case "":
        v15 := GenS1{F0: -174, F1: -504}
        _ = v15
    default:
        v16 := GenS2{F0: 0, F1: 0}
        _ = v16
    }
    v17 := GenS1{F0: 0, F1: -256}
    v18 := GenS2{F0: 4294967295, F1: 76}
    v19 := GenS3{F0: 833, F1: -336, F2: 128}
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return -9223372036854775808
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 0
    var v3 bool = false
    var v4 int64 = -213
    fmt.Println(v2)
    v1 += v1
    v1 -= (v1 | 127)
    fmt.Println(v2)
    v20 := []int{4294967295, 128, 652}
    var v21 int32 = -176
    v2 += v2
    v21 = (^(v21))
    var v22 string = "foo bar"
    for fi23 := 0; fi23 < 9; fi23++ {
        _ = len(v20)
        var v24 float32 = float32(3.14159)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
