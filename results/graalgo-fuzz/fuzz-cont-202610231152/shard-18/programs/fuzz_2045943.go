package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS5 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 int8, p1 int16, p2 bool) uint64 {
    var v7 uint8 = 185
    if ((1e10 - 0.1) <= 1.5) {
        v8 := GenS1{F0: 0}
        p1 = (p1 << ((563) & 63))
        if false {
            v9 := []int{464, -102, 109}
            var v10 int32 = 0
            p2 = p2
            _ = v10
            _ = v9
        }
        _ = v8
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint64 = uint64(v7)
            _ = fi11
            _ = v12
        }
    }
    var v13 uint32 = 4294967295
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "" {
        case "a":
            switch "\n\t" {
            case "foo bar":
                var v15 int16 = p1
                _ = v15
            case "\n\t":
                v16 := GenS2{F0: ("\n\t" + "\n\t")}
                _ = v16
            case "a":
                v17 := GenS2{F0: ("a" + "hello")}
                _ = v17
            default:
                p2 = p2
            }
        default:
            p2 = true
        }
        if true {
            var v18 string = "hello"
            v18 = ""
            _ = v18
        } else {
            switch "a" {
            case "foo bar":
                var v19 uint32 = 4294967295
                _ = v19
            case "\n\t":
                fmt.Println(v13)
            default:
                var v20 bool = ((((0.0 / 1.5) + 1e10) + float64(1e10)) == ((float64((1e10 + 0.0)) - (0.1 - 0.0)) - -2.25))
                _ = v20
            }
        }
        v21 := []int{(^(fi14)), (^(fi14)), (fi14 ^ fi14)}
        _ = fi14
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 42
}
func genFn1(p0 uint16) int64 {
    v22 := []int{9223372036854775807, 32768}
    switch -335 {
    case 4294967295:
        p0 -= p0
    case 127:
        v23 := GenS1{F0: 1}
        _ = v23
    }
    _ = p0
    _ = v22
    return -1
}

func main() {
    v1 := []uint32{985, 98}
    var v2 string = "\n\t"
    var v3 uint16 = 0
    var v4 int8 = -128
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]bool)
    v2 = ("" + "a")
    var v24 uint8 = 113
    var v25 int64 = -871
    v4 = v4
    var v26 int64 = (v25 + (v25 * (v25 ^ v25)))
    switch (v2 + (v2 + v2)) {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        if (true || false) {
            _ = len(v1)
        }
    case "":
        v24 = v24
    default:
        fmt.Println(v25)
    }
    var v27 bool = (v4 != 0)
    for fi28 := 0; fi28 < 4; fi28++ {
        switch v2 {
        case "\n\t":
            switch (fi28 % ((849) | 1)) {
            case 1:
                var v29 int64 = v26
                _ = v29
            case 2147483647:
                if (476 >= -780) {
                    v30 := GenS3{F0: (-(fi28))}
                    _ = v30
                }
            }
        case "":
            v31 := GenS4{F0: 2, F1: 12, F2: 1}
            _ = v31
        case "hello":
            v25 -= v26
        default:
            switch v2 {
            case "a":
                v4 -= (v4 * v4)
            default:
                var v32 string = ((("foo bar" + v2) + (v2 + ("foo bar" + v2))) + "hello")
                _ = v32
            }
        }
        _ = fi28
    }
    var v33 int8 = 127
    var v34 bool = false
    v26 = 9223372036854775807
    var v35 string = "a"
    v26 += v25
    var v36 int16 = 1
    if (0 > genFn0(v33, v36, (!v34))) {
        v24 += (v24 >> ((157) & 63))
        for fi37 := 0; fi37 < 4; fi37++ {
            v5[227] = 65535
            v27 = true
            v38 := GenS3{F0: 32768}
            _ = fi37
            _ = v38
        }
    } else {
        v39 := GenS5{F0: "hello", F1: 46, F2: 65535}
        v3 -= (^(v3))
        _ = v39
    }
    v40 := GenS2{F0: "hello"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
