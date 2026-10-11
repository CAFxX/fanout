package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint16, p1 int16, p2 uint64) int8 {
    var v5 uint32 = 651
    switch 128 {
    case -254:
        var v6 int = 454
        _ = v6
    case 930:
        p2 -= 4294967295
    case 0:
        v5 = 986
    default:
        switch ("a" + "") {
        case "\n\t":
            switch "\n\t" {
            case "a":
                v5 = 0
            case "foo bar":
                fmt.Println(p1)
            case "\n\t":
                p1 -= -245
            default:
                if false {
                    var v7 int64 = 9223372036854775807
                    for fi8 := 0; fi8 < 15; fi8++ {
                        v9 := []int{fi8, (fi8 % ((42) | 1)), fi8, (fi8 / ((128) | 1))}
                        var v10 uint8 = 1
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = v7
                }
            }
        case "":
            if (true && ((p1 >> ((1) & 63)) != (^(p1)))) {
                p0 = (p0 ^ (p0 - (^(p0))))
                v11 := []int{4294967295, -2147483648, 122}
                fmt.Println(p0, p0, p2)
                _ = v11
            }
        case "a":
            var v12 int = -2147483648
            _ = v12
        }
    }
    var v13 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    return 50
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 65535
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v2 = 255
    switch -418 {
    case 32767:
        var v14 uint = 0
        _ = v14
    }
    var v15 int8 = 0
    v2 = 42
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(v2, v2)
        _ = fi16
    }
    v17 := GenS1{F0: -123, F1: 39}
    v18 := []int{}
    v19 := GenS1{F0: (v15 ^ -8), F1: 2}
    v15 = (v15 << ((4294967295) & 63))
    delete(v3, 128)
    switch ("hello" + "a") {
    case "":
        v2 -= (v2 / ((42) | 1))
    }
    fmt.Println(v15, v2)
    fmt.Println(v15)
    v15 = 27
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int64 = 0
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
