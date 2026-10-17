package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 uint; F2 string }
type GenS5 struct { F0 int8; F1 uint; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 string; F2 int16 }
func genFn0(p0 uint8, p1 bool, p2 uint16) uint32 {
    v6 := GenS1{F0: 0}
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int, p1 uint64, p2 uint64) bool {
    switch p0 {
    case -349:
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS2{F0: -9223372036854775808}
            _ = fi7
            _ = v8
        }
    }
    p2 -= p2
    if (((65535 != 325) || true) && (!(!(42 == 9223372036854775807)))) {
        switch "hello" {
        case "a":
            p2 = 0
        case "hello":
            p2 += (p2 + 1)
        case "\n\t":
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 string = "a"
                fmt.Println(p1, p2)
                _ = fi9
                _ = v10
            }
        default:
            v11 := GenS3{F0: p1, F1: 227, F2: 42}
            _ = v11
        }
        var v12 int64 = -309
        var v13 int32 = 83
        _ = v12
        _ = v13
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return (0 >= -1)
}

func main() {
    var v1 uint16 = 255
    var v2 int16 = 42
    var v3 uint32 = 65535
    v4 := []uint32{0, v3, 487}
    v5 := []string{"\n\t", "hello", ""}
    fmt.Println(v3, v3)
    if (("foo bar" + (("hello" + ("a" + ("" + ((("" + ("" + "foo bar")) + "hello") + (("hello" + ("" + "hello")) + "a"))))) + (("a" + "hello") + "\n\t"))) < "") {
        v1 = 225
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    }
    var v14 uint64 = 0
    var v15 uint = 255
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v2, v15, v1)
        v14 = (v14 << ((uint(v14)) & 63))
        fmt.Println(v14, v15)
        _ = fi16
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v3, v3)
    default:
        var v17 float64 = 1e10
        _ = v17
    }
    switch (("a" + "foo bar") + "") {
    case "hello":
        switch "hello" {
        case "hello":
            for fi18 := 0; fi18 < 2; fi18++ {
                v3 = v3
                _ = fi18
            }
        case "a":
            var v19 int32 = 0
            _ = v19
        default:
            v5 = append(v5, (("\n\t" + "") + "hello"))
        }
    case "a":
        fmt.Println(v1, v1)
    case "":
        v5 = append(v5, ("" + "\n\t"))
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            if true {
                var v21 uint = (^(v15))
                _ = v21
            }
            _ = fi20
        }
    }
    _ = len(v4)
    if ((v1 ^ 2) >= (v1 | (v1 & 32767))) {
        v22 := GenS3{F0: v14, F1: 2, F2: -434}
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(v3, v1)
        if (699 > 42) {
            v24 := GenS3{F0: (v14 * (^(v14))), F1: (v3 >> ((uint(v3)) & 63)), F2: (^(fi23))}
            fmt.Println(v14)
            v4 = append(v4, 0)
            _ = v24
        }
        v25 := GenS4{F0: fi23, F1: (v15 / ((v15) | 1)), F2: "a"}
        _ = fi23
        _ = v25
    }
    v26 := GenS2{F0: -9223372036854775808}
    switch "\n\t" {
    case "":
        v27 := GenS5{F0: 0, F1: v15, F2: 0}
        _ = v27
    case "foo bar":
        v15 -= (^(v15))
    default:
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := GenS6{F0: true}
            if false {
                if false {
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                }
                v30 := GenS3{F0: v14, F1: 0, F2: 32768}
                v31 := GenS7{F0: "foo bar"}
                _ = v30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
    }
    v4 = append(v4, v3)
    if true {
        v32 := GenS8{F0: 2, F1: "hello", F2: v2}
        var v33 int32 = -2147483648
        _ = v32
        _ = v33
    }
    var v34 uint16 = (^(v1))
    if ((^(v3)) <= genFn0(127, (-128 > int8(v14)), v1)) {
        v35 := GenS7{F0: ("hello" + (("foo bar" + "") + "a"))}
        for fi36 := 0; fi36 < 15; fi36++ {
            v2 -= int16(v15)
            _ = fi36
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            v15 += 18446744073709551615
            _ = fi37
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
