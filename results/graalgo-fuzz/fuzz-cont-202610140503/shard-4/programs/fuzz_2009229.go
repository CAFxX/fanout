package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint16 }
func genFn0(p0 uint64, p1 int8, p2 uint8) string {
    for fi4 := 0; fi4 < 11; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            if false {
                fmt.Println(fi5, p0)
            }
            _ = fi5
        }
        p0 += p0
        p0 = p0
        _ = fi4
    }
    if ((!(false && (4294967295 > p0))) && true) {
        v6 := GenS1{F0: 641, F1: -9223372036854775808, F2: 1}
        p0 -= 2
        v7 := []int{-9223372036854775808, 1, -9223372036854775808}
        _ = v6
        _ = v7
    } else {
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 uint = 2
    v2 := []string{}
    v3 := []uint32{}
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = 42
        switch (("hello" + "\n\t") + "hello") {
        case "hello":
            v9 := GenS2{F0: 18446744073709551615, F1: 472, F2: 18446744073709551615}
            _ = v9
        case "":
            _ = len(v2)
        default:
            switch (fi8 << ((v1) & 63)) {
            case 857:
                v10 := GenS3{F0: -9223372036854775808, F1: false}
                _ = v10
            default:
                fmt.Println(fi8, v1)
            }
        }
        v11 := GenS2{F0: 42, F1: 512, F2: 2}
        _ = fi8
        _ = v11
    }
    if (!false) {
        v12 := GenS1{F0: 2, F1: -303, F2: 0}
        if (1e10 < 1e10) {
            v1 = v1
            v13 := GenS4{F0: -1, F1: 848, F2: 0}
            fmt.Println(v1)
            _ = v13
        }
        if false {
            var v14 int32 = -2147483648
            _ = v14
        }
        _ = v12
    } else {
        v1 -= v1
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int32 = 42
        _ = fi15
        _ = v16
    }
    v1 += v1
    v1 = (^(v1))
    for fi17 := 0; fi17 < 13; fi17++ {
        v3 = append(v3, 0)
        _ = fi17
    }
    if ((0.1 > (3.14159 / 1.0)) && false) {
        if true {
            v18 := GenS5{F0: (((false && false) && (65535 >= 370)) && (("a" + "foo bar") < ("foo bar" + ("a" + "\n\t")))), F1: -2147483648, F2: 65535}
            _ = v18
        }
    }
    v19 := GenS3{F0: 42, F1: ((!true) && (32768 != 2147483647))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
