package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int }

func main() {
    var v1 uint16 = 711
    var v2 uint16 = (v1 / ((493) | 1))
    var v3 uint16 = v1
    v4 := GenS1{F0: (v1 >> ((255) & 63)), F1: 255}
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(v2, v1, fi5)
        v3 += (v1 & (v2 ^ (v3 >> ((65535) & 63))))
        _ = fi5
    }
    var v6 uint64 = 2
    if (0 > 255) {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{(fi7 & -876), (-(fi7)), (fi7 & (-(fi7))), (-(fi7))}
            _ = fi7
            _ = v8
        }
        if ((42 == 2147483648) && true) {
            v2 = 42
            v1 = (^(v3))
        }
    } else {
        if ((!(int(v2) < 42)) && true) {
            v9 := []int{128, 42, 2147483647}
            var v10 int32 = 1
            _ = v10
            _ = v9
        }
        v11 := []int{385, 9223372036854775807}
        switch 175 {
        case 32767:
            for fi12 := 0; fi12 < 13; fi12++ {
                if len(v11) > 0 {
                    v11[0] = (-(fi12))
                }
                _ = fi12
            }
        }
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    case "a":
        v1 += 934
    case "\n\t":
        switch -9223372036854775808 {
        case 621:
            v13 := GenS1{F0: (v3 - uint16(v6)), F1: 127}
            _ = v13
        case 9223372036854775807:
            fmt.Println(v1)
        default:
            v14 := GenS2{F0: v1, F1: false}
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            v3 += v3
            v2 -= v1
            if true {
                v17 := GenS3{F0: 42, F1: (v1 | 42)}
                fmt.Println(v3, fi16, fi15)
                _ = v17
            }
            _ = fi16
        }
        if true {
            fmt.Println(fi15, v2, v1)
            v3 += v3
        }
        v1 += (v2 << ((255) & 63))
        _ = fi15
    }
    if false {
        var v18 uint = 4294967295
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint8 = 255
                v2 -= v3
                fmt.Println(v1, v6)
                _ = fi20
                _ = v21
            }
            v1 = v2
            if (-122 == -128) {
                if ((!(0 < 2)) && true) {
                    v1 -= v2
                    v1 = v3
                }
                var v22 int = 32768
                _ = v22
            } else {
                var v23 bool = (true || (0.1 <= 1e10))
                _ = v23
            }
            _ = fi19
        }
        switch "hello" {
        case "\n\t":
            var v24 bool = (!true)
            _ = v24
        }
        _ = v18
    } else {
        if (!true) {
            fmt.Println(v3, v3, v1)
            v25 := GenS2{F0: (^(v2)), F1: (false && false)}
            if (false && (1 > 4294967295)) {
                var v26 int32 = 42
                _ = v26
            }
            _ = v25
        } else {
            for fi27 := 0; fi27 < 15; fi27++ {
                switch "a" {
                case "foo bar":
                    v6 = (^(v6))
                }
                _ = fi27
            }
        }
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(v6)
        if (-2.25 <= (0.1 - 3.14159)) {
            v29 := []int{}
            var v30 string = "foo bar"
            _ = v29
            _ = v30
        } else {
            v31 := GenS4{F0: v1}
            _ = v31
        }
        _ = fi28
    }
    if (32767 > 0) {
        var v32 uint32 = 1
        var v33 uint64 = v6
        v3 += 343
        _ = v32
        _ = v33
    } else {
        switch "hello" {
        case "hello":
            v2 += 32767
        case "\n\t":
            v1 = v2
        case "a":
            v34 := []int{597, 128, 127, 127}
            _ = v34
        default:
            if (99 > 1) {
                if false {
                    if true {
                        v1 -= v1
                    } else {
                        var v35 uint64 = (v6 - v6)
                        _ = v35
                    }
                    var v36 uint64 = 269
                    var v37 uint64 = v6
                    _ = v36
                    _ = v37
                } else {
                    var v38 int = 505
                    v39 := GenS3{F0: 102, F1: v2}
                    _ = v38
                    _ = v39
                }
            }
        }
    }
    var v40 int16 = 42
    v41 := []int{32767, -847}
    var v42 int = 128
    switch "" {
    case "hello":
        v43 := GenS5{F0: v3, F1: (v42 << ((4294967295) & 63))}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v6
}
