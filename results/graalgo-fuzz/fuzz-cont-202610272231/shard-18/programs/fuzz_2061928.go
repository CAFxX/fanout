package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint16 }

func main() {
    v1 := []uint32{65535, 1}
    v2 := []uint32{2147483648, 2, 4294967295, 42}
    var v3 int32 = 2147483647
    v4 := []int{-1, 266}
    v5 := GenS1{F0: (!((false && false) && true)), F1: 2147483648}
    var v6 int64 = 1
    switch "a" {
    case "a":
        var v7 uint16 = 255
        _ = v7
    }
    _ = len(v4)
    v8 := GenS2{F0: 255, F1: 42}
    if ((!(0 >= -9223372036854775808)) || ((!(false || true)) || true)) {
        if true {
            v6 = (v6 << ((42) & 63))
        }
        if (0.0 <= -2.25) {
            switch -749 {
            case 127:
                var v9 uint8 = 1
                _ = v9
            case -2147483648:
                fmt.Println(v6, v6)
            }
        }
    }
    if false {
        v10 := GenS3{F0: 13, F1: 18446744073709551615, F2: 32767}
        _ = v10
    } else {
        v11 := GenS1{F0: (!true), F1: 1}
        _ = v11
    }
    if (!false) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v12 := GenS3{F0: 1, F1: 0, F2: 255}
        _ = v12
    } else {
        if (!((false && false) || false)) {
            v4 = append(v4, 984)
            var v13 uint8 = 88
            if false {
                var v14 uint = 1
                _ = v14
            } else {
                var v15 int64 = -9223372036854775808
                _ = v15
            }
            _ = v13
        }
    }
    v3 = (v3 % (((v3 | v3)) | 1))
    v6 -= v6
    var v16 uint16 = 42
    v16 = v16
    var v17 int64 = (v6 / (((v6 / ((v6) | 1))) | 1))
    switch "foo bar" {
    case "foo bar":
        switch ("\n\t" + "foo bar") {
        case "a":
            if len(v2) > 0 {
                v2[0] = 691
            }
        case "":
            v18 := GenS2{F0: 18446744073709551615, F1: 0}
            _ = v18
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
