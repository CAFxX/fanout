package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
func genFn0(p0 uint) int64 {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 string = ((("a" + "a") + "") + ("" + (("a" + (("\n\t" + "foo bar") + "hello")) + ("" + "hello"))))
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint, p2 string) uint {
    if false {
        p0 = p0
    } else {
        p1 += (p1 >> ((p1) & 63))
        var v6 uint64 = p0
        switch ((p2 + p2) + "foo bar") {
        case "foo bar":
            fmt.Println(p2)
        case "a":
            if (!((1 > 779) || (65535 == 32767))) {
                p2 = "hello"
                if (false || true) {
                    v7 := GenS1{F0: -894, F1: -2147483648, F2: -1}
                    _ = v7
                }
                v8 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                _ = v8
            } else {
                p0 = 538
            }
        }
        _ = v6
    }
    if true {
        v9 := []int{-705, 9223372036854775807}
        v10 := GenS2{F0: -1, F1: 833, F2: 0}
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint8 = 128
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 int = fi11
                var v16 float64 = 0.1
                v9 = append(v9, -9223372036854775808)
                var v17 int8 = 122
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            v18 := GenS2{F0: 0, F1: 2147483648, F2: -1}
            _ = fi11
            _ = v18
        }
        _ = v10
        _ = v9
    } else {
        v19 := []int{-2147483648, 904}
        for fi20 := 0; fi20 < 13; fi20++ {
            p0 = p0
            fmt.Println(p2, fi20, fi20)
            _ = fi20
        }
        _ = v19
    }
    if false {
        fmt.Println(p2, p0, p1)
    } else {
        var v21 int32 = 0
        v22 := GenS1{F0: v21, F1: 42, F2: 127}
        _ = v21
        _ = v22
    }
    if true {
        p2 = "a"
        var v23 uint16 = 818
        _ = v23
    } else {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS3{F0: (fi24 & fi24), F1: (0.1 != 1.5), F2: 148}
            var v26 int32 = 42
            p0 -= (^(p0))
            _ = fi24
            _ = v25
            _ = v26
        }
        p0 = p0
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return (p1 + (p1 % ((255) | 1)))
}

func main() {
    var v1 uint8 = 0
    var v2 int64 = 9223372036854775807
    var v3 bool = true
    v27 := []int{}
    v3 = true
    switch 2147483647 {
    case 0:
        for fi28 := 0; fi28 < 13; fi28++ {
            if (-112 <= 35) {
                v29 := GenS2{F0: -1, F1: 2, F2: 0}
                _ = v29
            }
            _ = fi28
        }
    case -985:
        v1 += v1
    case 153:
        var v30 int32 = -397
        _ = v30
    }
    v1 = 138
    v1 = v1
    fmt.Println(v2, v2, v3)
    for fi31 := 0; fi31 < 18; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 bool = (1 != 2147483647)
            _ = fi32
            _ = v33
        }
        var v34 uint16 = 65535
        _ = fi31
        _ = v34
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    switch (("hello" + (("foo bar" + "a") + "foo bar")) + ("\n\t" + "\n\t")) {
    case "\n\t":
        if ((-25 != -28) && true) {
            if (false || false) {
                v3 = v3
                var v35 float32 = 0.1
                fmt.Println(v2, v1, v1)
                _ = v35
            } else {
                var v36 int16 = 68
                var v37 bool = false
                _ = v36
                _ = v37
            }
            v27 = append(v27, 32768)
            if (-363 > 264) {
                v38 := GenS1{F0: -442, F1: 1, F2: 9223372036854775807}
                fmt.Println(v1)
                v2 -= v2
                _ = v38
            }
        }
    case "hello":
        v1 -= (v1 % ((128) | 1))
    default:
        var v39 uint16 = 0
        _ = v39
    }
    _ = len(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
