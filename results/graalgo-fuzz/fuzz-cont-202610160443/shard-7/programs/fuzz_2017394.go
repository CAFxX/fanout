package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint16 }
func genFn0(p0 string, p1 int32, p2 uint8) string {
    var v7 int8 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        if (!true) {
            v7 += v7
        }
        var v9 uint8 = 2
        p1 += (p1 / ((p1) | 1))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: p0}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            p2 += (p2 | (p2 % ((p2) | 1)))
            var v13 uint16 = 454
            _ = fi12
            _ = v13
        }
        var v14 uint8 = (^(p2))
        v15 := []int{(^(fi11)), -716, 128}
        _ = fi11
        _ = v14
        _ = v15
    }
    v16 := []int{-48, -2147483648, -877, 820}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v16
    _ = v7
    return (p0 + "hello")
}
func genFn1(p0 uint8, p1 uint16, p2 string) int {
    if ((0 == -341) || (!(p1 == p1))) {
        fmt.Println(p1, p0)
        if (false && false) {
            p1 += p1
            var v17 int16 = 1
            _ = v17
        } else {
            p0 = (p0 & p0)
            var v18 string = (((p2 + p2) + p2) + "hello")
            _ = v18
        }
        v19 := []int{42}
        _ = v19
    } else {
        fmt.Println(p1, p2, p2)
        v20 := []int{0, 871, 0}
        p0 -= 173
        _ = v20
    }
    switch p2 {
    case "":
        switch "foo bar" {
        case "":
            p0 += p0
        case "\n\t":
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS2{F0: (false || true), F1: ("" + p2)}
                p2 = p2
                _ = fi21
                _ = v22
            }
        default:
            var v23 uint8 = (p0 % ((255) | 1))
            _ = v23
        }
    case "hello":
        if (0 >= 0) {
            var v24 float64 = ((1.5 / 3.14159) + (-2.25 / 1.0))
            v24 += 0.1
            _ = v24
        }
    default:
        var v25 uint64 = uint64(p1)
        _ = v25
    }
    fmt.Println(p0, p2, p0)
    var v26 int64 = -541
    v27 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 396
}

func main() {
    var v1 int8 = 127
    v2 := []int{291, -1, 9223372036854775807, -1}
    v3 := []int{128, 1, 849}
    var v4 string = "foo bar"
    v5 := []uint32{0, 0}
    v6 := make(map[string]bool)
    v4 = (v4 + v4)
    if true {
        if (-2147483648 > 888) {
            v28 := GenS3{F0: 42}
            switch ((v4 + ((v4 + v4) + v4)) + v4) {
            case "a":
                var v29 uint = 658
                _ = v29
            default:
                switch (("" + v4) + v4) {
                case "a":
                    var v30 float64 = -2.25
                    _ = v30
                case "":
                    var v31 int8 = -128
                    _ = v31
                }
            }
            _ = v28
        } else {
            var v32 float32 = 0.1
            v1 += (v1 >> ((65535) & 63))
            _ = v32
        }
    }
    fmt.Println(v1)
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint64 = 18446744073709551615
        _ = fi33
        _ = v34
    }
    delete(v6, "\n\t")
    v1 += -3
    v6[v4] = false
    if false {
        v35 := GenS3{F0: 32767}
        _ = len(v2)
        _ = v35
    } else {
        var v36 bool = true
        if true {
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := v6[v4]
                _ = fi37
                _ = v38
            }
        }
        _ = v36
    }
    if (v4 < v4) {
        for fi39 := 0; fi39 < 16; fi39++ {
            fmt.Println(v4)
            var v40 int64 = 9223372036854775807
            _ = fi39
            _ = v40
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(v4, fi41, fi41)
            case "":
                if (42 != 2) {
                    var v42 int8 = v1
                    _ = v42
                }
            case "hello":
                switch (fi41 - fi41) {
                case -904:
                    fmt.Println(fi41)
                case 32768:
                    fmt.Println(fi41, v1)
                }
            default:
                var v43 int8 = (v1 % ((v1) | 1))
                _ = v43
            }
            var v44 int16 = 0
            v44 -= (v44 ^ 42)
            _ = fi41
            _ = v44
        }
    } else {
        switch 652 {
        case 32768:
            if ((!(0 >= 0)) && true) {
                var v45 int8 = (-(v1))
                _ = v45
            }
        }
    }
    if (!false) {
        _ = len(v2)
    }
    if false {
        v46 := GenS2{F0: true, F1: "\n\t"}
        _ = v46
    } else {
        for fi47 := 0; fi47 < 8; fi47++ {
            fmt.Println(v1, fi47)
            var v48 int16 = 390
            _ = fi47
            _ = v48
        }
        var v49 int16 = 32767
        _ = v49
    }
    v6["a"] = (v1 == (v1 % (((v1 / ((0) | 1))) | 1)))
    for fi50 := 0; fi50 < 1; fi50++ {
        switch 32768 {
        case 128:
            for fi51 := 0; fi51 < 8; fi51++ {
                var v52 int16 = 0
                _ = fi51
                _ = v52
            }
        case 127:
            var v53 uint16 = 2
            _ = v53
        case -550:
            v4 = v4
        default:
            v54 := GenS3{F0: 341}
            _ = v54
        }
        _ = fi50
    }
    var v55 float32 = (1e10 / (3.14159 + (1.5 * (-2.25 - 1.5))))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
