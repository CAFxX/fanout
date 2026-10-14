package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 65535
    var v3 int8 = 42
    var v4 int32 = 802
    v5 := make(map[int]bool)
    switch 241 {
    case 1:
        v6 := []int{(v1 * v1)}
        _ = v6
    case 4294967295:
        v3 += (v3 + v3)
    case 32767:
        fmt.Println(v2, v4, v1)
    }
    var v7 uint32 = 153
    if (32 >= 770) {
        v8 := []int{}
        _ = v8
    }
    v4 = 42
    v2 += v2
    v7 = 4294967295
    switch "" {
    case "":
        fmt.Println(v2)
    default:
        v1 = v1
    }
    v9 := GenS1{F0: false, F1: v4}
    v7 = 2147483648
    v10 := GenS2{F0: 1, F1: (!(18446744073709551615 > 1))}
    if false {
        if (!false) {
            var v11 bool = (1 < 255)
            v12 := []int{}
            var v13 int32 = (v4 << ((1) & 63))
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := []int{v1, (v1 & (v1 + v1)), (v1 + 562)}
            switch "a" {
            case "foo bar":
                var v15 float64 = 3.14159
                _ = v15
            case "hello":
                fmt.Println(v1, v2, v2)
            case "a":
                if true {
                    v16 := GenS3{F0: 71}
                    _ = v16
                } else {
                    v7 = v7
                }
            }
            if false {
                v1 = (v1 << ((4294967295) & 63))
            } else {
                v5[v1] = (-1 != (^(v3)))
                v3 = v3
            }
            _ = v14
        }
        v17 := GenS1{F0: true, F1: v4}
        fmt.Println(v1)
        _ = v17
    }
    var v18 string = ("a" + "")
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS3{F0: 195}
        _ = fi19
        _ = v20
    }
    if false {
        v5[v1] = ((v4 - (v4 % ((v4) | 1))) == (v4 / ((2147483647) | 1)))
    } else {
        v21 := []int{}
        if true {
            for fi22 := 0; fi22 < 11; fi22++ {
                v21 = append(v21, 0)
                var v23 float64 = float64(-2.25)
                var v24 uint32 = v7
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            var v25 uint = 652
            _ = v25
        }
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
