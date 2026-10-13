package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int; F1 uint8 }

func main() {
    var v1 bool = (!false)
    v2 := []int{-2147483648, 984}
    var v3 uint64 = 0
    var v4 int = 42
    if ((v4 * 675) != v4) {
        fmt.Println(v1, v4, v4)
        if ((100 >= -128) || (v3 <= v3)) {
            fmt.Println(v4, v3, v3)
            if (("foo bar" + "foo bar") < "foo bar") {
                if v1 {
                    v1 = v1
                    var v5 int8 = 42
                    v1 = (true && ((32767 <= 813) || v1))
                    _ = v5
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v6 := GenS1{F0: v4}
    v3 -= v3
    v4 -= v4
    v4 -= (v4 | (-(v4)))
    var v7 int64 = -1
    var v8 uint64 = (v3 % ((42) | 1))
    switch "hello" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        v9 := GenS2{F0: (v7 % ((v7) | 1))}
        _ = v9
    case "foo bar":
        var v10 string = "\n\t"
        _ = v10
    default:
        if true {
            v4 = v4
            var v11 uint16 = uint16(v3)
            var v12 uint32 = 2
            _ = v11
            _ = v12
        } else {
            if len(v2) > 0 {
                v2[0] = v4
            }
            if false {
                v8 = v8
            }
            v7 += -9223372036854775808
        }
    }
    v2 = append(v2, int(v7))
    fmt.Println(v8, v4, v1)
    var v13 int = (v4 << ((2) & 63))
    v4 = (-(v13))
    if false {
        var v14 int16 = -1
        _ = v14
    } else {
        v15 := GenS3{F0: -458, F1: -1}
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := GenS4{F0: 4294967295, F1: 2}
            var v18 uint64 = 1
            switch (fi16 % (((fi16 % ((917) | 1))) | 1)) {
            case 128:
                v2 = append(v2, -9223372036854775808)
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
