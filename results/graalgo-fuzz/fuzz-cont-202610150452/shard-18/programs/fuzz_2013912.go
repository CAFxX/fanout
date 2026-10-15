package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int8; F1 string }
func genFn0(p0 uint64, p1 uint32, p2 bool) string {
    switch "foo bar" {
    case "\n\t":
        v5 := []int{4294967295, 805, 42, 9223372036854775807}
        _ = v5
    default:
        v6 := []int{-378, 0, -324, -2147483648}
        _ = v6
    }
    v7 := []int{128, -46, 4294967295, 32767}
    v8 := GenS1{F0: (1.5 < 1.5), F1: 274, F2: 141}
    for fi9 := 0; fi9 < 16; fi9++ {
        p0 = p0
        v7 = append(v7, -9223372036854775808)
        _ = fi9
    }
    var v10 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 bool) int {
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11, p0)
        p0 = (true || (p0 || true))
        _ = fi11
    }
    switch -556 {
    case -1:
        for fi12 := 0; fi12 < 5; fi12++ {
            if (((3.14159 + (3.14159 * 0.0)) + (0.1 + 3.14159)) > (0.1 - 0.1)) {
                switch (fi12 << ((255) & 63)) {
                case 32767:
                    fmt.Println(p0, fi12)
                case 32768:
                    switch "foo bar" {
                    case "a":
                        v13 := GenS1{F0: ((("hello" + "a") > ("hello" + (("hello" + "") + ""))) && true), F1: 740, F2: 23}
                        _ = v13
                    default:
                        switch ("a" + "") {
                        case "hello":
                            v14 := []int{fi12, fi12, 2147483647}
                            _ = v14
                        case "":
                            var v15 float32 = float32(-2.25)
                            _ = v15
                        case "\n\t":
                            v16 := GenS2{F0: -34}
                            _ = v16
                        default:
                            fmt.Println(p0)
                        }
                    }
                }
                var v17 int = fi12
                _ = v17
            }
            v18 := GenS2{F0: -128}
            _ = fi12
            _ = v18
        }
    case 42:
        v19 := GenS3{F0: 0, F1: 947, F2: 217}
        _ = v19
    case 128:
        v20 := []int{-359}
        _ = v20
    default:
        if (("a" != "foo bar") || (127 > -1)) {
            var v21 uint64 = 512
            v22 := []int{128, 32768}
            v21 = v21
            _ = v21
            _ = v22
        }
    }
    if (939 > 248) {
        fmt.Println(p0)
    }
    v23 := []int{9223372036854775807}
    _ = p0
    _ = v23
    return -9223372036854775808
}

func main() {
    var v1 float64 = (1.5 * float64(float32(3.14159)))
    var v2 int = -60
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    if false {
        var v24 float32 = (((0.0 * (float32((v1 + 0.1)) + float32(v1))) / 1.0) - 1.5)
        _ = v24
    } else {
        switch genFn1(("hello" == "")) {
        case -1:
            for fi25 := 0; fi25 < 19; fi25++ {
                v26 := GenS4{F0: 255, F1: 2, F2: 1}
                _ = fi25
                _ = v26
            }
        default:
            fmt.Println(v1, v2)
        }
        if (283 == 1) {
            switch (v2 * v2) {
            case 42:
                fmt.Println(v2)
            case -720:
                var v27 string = "a"
                _ = v27
            case 9223372036854775807:
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := v4["foo bar"]
                    v30 := GenS5{F0: (v2 ^ (fi28 / (((v2 << ((42) & 63))) | 1)))}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            default:
                v31 := GenS1{F0: (1 <= 1), F1: 0, F2: 194}
                _ = v31
            }
            for fi32 := 0; fi32 < 1; fi32++ {
                delete(v3, 1)
                _ = fi32
            }
            v33 := GenS4{F0: 255, F1: 640, F2: 509}
            _ = v33
        } else {
            v1 = v1
            v2 += 9223372036854775807
        }
        v34 := GenS3{F0: 2147483647, F1: 42, F2: 42}
        _ = v34
    }
    v1 += 3.14159
    v35 := GenS2{F0: 1}
    for fi36 := 0; fi36 < 18; fi36++ {
        delete(v4, "foo bar")
        if (-128 != 1) {
            v1 -= (0.0 / -2.25)
        } else {
            v2 = (v2 * 856)
        }
        var v37 int = (fi36 + fi36)
        _ = fi36
        _ = v37
    }
    var v38 float32 = 0.1
    var v39 int64 = 9223372036854775807
    for fi40 := 0; fi40 < 2; fi40++ {
        v41 := []int{42}
        v42 := GenS6{F0: 255, F1: 128}
        v43 := GenS6{F0: 0, F1: 121}
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    v44 := GenS7{F0: -2147483648, F1: (v39 * v39)}
    var v45 int32 = 334
    v1 = (((v1 + ((float64((v38 - 0.1)) - -2.25) - v1)) + (v1 * (v1 * 0.0))) + 0.1)
    fmt.Println(v2, v1, v38)
    switch (v2 | (v2 + 4294967295)) {
    case 32:
        v2 = (v2 | genFn1(((!(0 == 72)) || true)))
    default:
        _ = len(v4)
    }
    switch (v2 << ((65535) & 63)) {
    case 0:
        var v46 bool = false
        _ = v46
    case 1:
        var v47 int32 = (v45 | 1)
        _ = v47
    case 9223372036854775807:
        v48 := []int{}
        _ = v48
    default:
        if (!(int32(v39) >= (^(v45)))) {
            var v49 bool = (438 != int64(v2))
            v39 = 1
            _ = v49
        } else {
            v50 := GenS8{F0: 1, F1: "\n\t"}
            v2 += v2
            _ = v50
        }
    }
    v39 -= (^(v39))
    var v51 int8 = 110
    if false {
        var v52 float64 = (0.0 + (v1 + v1))
        _ = v52
    } else {
        v53 := []int{v2, (v2 >> ((748) & 63)), 42}
        _ = v53
    }
    if (42 >= 671) {
        fmt.Println(v45)
    } else {
        var v54 int8 = v51
        _ = v54
    }
    v55 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v51
    _ = v55
}
