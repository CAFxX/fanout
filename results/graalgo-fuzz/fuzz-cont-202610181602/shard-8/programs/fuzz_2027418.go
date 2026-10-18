package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int16; F2 bool }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint32 = 42
            v10 := GenS1{F0: (fi7 | fi8), F1: 4294967295}
            v9 = 273
            _ = fi8
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 float32 = 3.14159
            _ = fi11
            _ = v12
        }
        if (4294967295 != 2) {
            fmt.Println(fi7)
            if (true && false) {
                var v13 int8 = 1
                _ = v13
            }
        } else {
            var v14 string = (("foo bar" + ((("hello" + "\n\t") + (("\n\t" + "a") + ("a" + "a"))) + "\n\t")) + ("\n\t" + "a"))
            var v15 uint32 = 1
            v14 = "foo bar"
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    if true {
        var v16 float64 = 1.5
        _ = v16
    }
    v17 := []int{9223372036854775807, 127}
    var v18 int64 = -841
    _ = len(v17)
    _ = v17
    _ = v18
    return 65535
}
func genFn1(p0 uint8, p1 int8, p2 int32) uint32 {
    var v19 bool = (false || ("" == ("foo bar" + "")))
    for fi20 := 0; fi20 < 12; fi20++ {
        if (v19 && true) {
            v21 := GenS2{F0: 1, F1: 0, F2: int32(fi20)}
            _ = v21
        }
        if (!((!v19) && ("hello" != ("" + "hello")))) {
            p2 += -1
        }
        if ("a" == "foo bar") {
            if ("hello" == "a") {
                var v22 float64 = 1.5
                _ = v22
            }
        }
        _ = fi20
    }
    v23 := GenS1{F0: 9223372036854775807, F1: 700}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v23
    return uint32(p0)
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    var v3 int32 = -1
    v4 := make(map[int]bool)
    var v5 int = -202
    v6 := []string{"", ("a" + "a"), (("foo bar" + "hello") + (("foo bar" + "") + ""))}
    v24 := GenS3{F0: true}
    fmt.Println(v5, v5)
    v3 = 1
    switch "a" {
    case "hello":
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int = 0
            _ = fi25
            _ = v26
        }
    default:
        var v27 int8 = -1
        _ = v27
    }
    if (((((42 == 1) || false) || (((-160 >= -9223372036854775808) || false) || ((((false || true) && (((v5 << ((205) & 63)) == -139) || false)) && (false && true)) && false))) && (0.1 != (3.14159 + (((((0.0 + 0.0) / 1.0) * 0.0) / (1e10 - (float64(1e10) * ((1e10 * (-2.25 + 3.14159)) / 1.0)))) + (1e10 - 0.0))))) && (-650 >= -1)) {
        v5 -= 466
        for fi28 := 0; fi28 < 8; fi28++ {
            if (false || false) {
                var v29 bool = (!((-(v3)) != v3))
                fmt.Println(v3, v5)
                v30 := GenS2{F0: -9223372036854775808, F1: 2, F2: v3}
                _ = v29
                _ = v30
            }
            if len(v6) > 0 {
                v6[0] = ("hello" + "\n\t")
            }
            if false {
                v31 := GenS4{F0: -717}
                v32 := GenS3{F0: (128 >= 127)}
                fmt.Println(fi28, v5)
                _ = v31
                _ = v32
            }
            _ = fi28
        }
        var v33 uint16 = 0
        _ = v33
    } else {
        switch v5 {
        case -2147483648:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        default:
            v5 = int(v3)
        }
        fmt.Println(v5, v5)
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v5, v5, v5)
        switch v5 {
        case 728:
            v35 := GenS5{F0: 65535}
            _ = v35
        case 1:
            for fi36 := 0; fi36 < 9; fi36++ {
                v5 -= -132
                _ = fi36
            }
        case -814:
            v37 := GenS6{F0: (v5 >> ((932) & 63)), F1: 0, F2: false}
            _ = v37
        }
        _ = fi34
    }
    v38 := GenS2{F0: 0, F1: 42, F2: 156}
    var v39 int32 = (-(v3))
    v3 = -338
    v5 -= (v5 << ((0) & 63))
    switch "\n\t" {
    case "":
        v5 -= (v5 * 2147483647)
    case "hello":
        switch "foo bar" {
        case "":
            v5 = 32768
        case "hello":
            v40 := GenS2{F0: 1, F1: 42, F2: 2147483647}
            _ = v40
        }
    }
    v39 -= v3
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 int64 = 1
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
