package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS7 struct { F0 uint64 }
func genFn0() uint64 {
    switch 0 {
    case -2147483648:
        var v7 bool = (false && true)
        _ = v7
    case -884:
        switch ("a" + "foo bar") {
        case "foo bar":
            v8 := GenS1{F0: 9223372036854775807, F1: -452}
            _ = v8
        case "":
            var v9 bool = ((false && (2147483647 <= 0)) || (!(1000 != 745)))
            _ = v9
        }
    case -9223372036854775808:
        v10 := []int{}
        _ = v10
    }
    var v11 uint8 = 127
    for fi12 := 0; fi12 < 18; fi12++ {
        switch fi12 {
        case -2147483648:
            fmt.Println(v11, v11)
        case -825:
            var v13 uint = 535
            _ = v13
        default:
            v14 := []int{-1, -506}
            _ = v14
        }
        fmt.Println(fi12, fi12)
        switch (fi12 + (fi12 * fi12)) {
        case 0:
            var v15 int8 = 0
            _ = v15
        case 482:
            v11 -= (v11 - 2)
        default:
            v16 := []int{42, (fi12 << ((2) & 63)), 127}
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: 18446744073709551615, F1: 42}
    _ = v11
    _ = v17
    return 0
}
func genFn1(p0 int64, p1 int64) uint64 {
    v18 := []int{2147483647, 521, 0, 9223372036854775807}
    p0 -= -1
    p0 -= 9223372036854775807
    p0 -= -121
    _ = p0
    _ = p1
    _ = v18
    return 330
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 128
    v4 := make(map[uint32]uint32)
    var v5 float32 = 0.0
    v6 := make(map[string]int)
    var v19 int64 = 90
    fmt.Println(v19, v3)
    for fi20 := 0; fi20 < 14; fi20++ {
        if (true || ((91 <= 0) && ((v5 + (1e10 - 0.0)) == float32((1e10 / 3.14159))))) {
            var v21 float64 = (-2.25 / (-2.25 * 1e10))
            switch (fi20 - fi20) {
            case 32768:
                var v22 uint8 = v3
                _ = v22
            case 4294967295:
                v23 := GenS1{F0: (fi20 ^ (fi20 << ((937) & 63))), F1: 304}
                _ = v23
            case 0:
                _ = len(v1)
            }
            var v24 int8 = 0
            _ = v21
            _ = v24
        }
        switch ("foo bar" + "") {
        case "\n\t":
            switch (^(fi20)) {
            case -9223372036854775808:
                fmt.Println(fi20)
            case 1:
                var v25 int64 = int64(v3)
                _ = v25
            default:
                var v26 int = fi20
                _ = v26
            }
        default:
            v27 := v6["\n\t"]
            _ = v27
        }
        v5 = v5
        _ = fi20
    }
    if ((!false) && true) {
        v19 -= (v19 & v19)
        v19 = (v19 >> ((4294967295) & 63))
    } else {
        v19 += (v19 % (((v19 / ((-346) | 1))) | 1))
        fmt.Println(v5, v19)
    }
    switch -66 {
    case 9223372036854775807:
        var v28 uint = 65535
        _ = v28
    case 0:
        var v29 bool = false
        _ = v29
    }
    v30 := v2[1]
    switch "hello" {
    case "":
        v31 := GenS2{F0: 1, F1: -87}
        _ = v31
    case "\n\t":
        var v32 uint64 = 18446744073709551615
        _ = v32
    case "foo bar":
        v33 := GenS3{F0: 173, F1: 177}
        _ = v33
    default:
        if ((!(1 < -32768)) && (true && (32767 >= 9223372036854775807))) {
            for fi34 := 0; fi34 < 8; fi34++ {
                delete(v2, 65535)
                _ = fi34
            }
            v30 = (775 >= 348)
            var v35 float64 = 0.0
            _ = v35
        }
    }
    var v36 int32 = -32
    v37 := GenS4{F0: -2147483648}
    var v38 int16 = 1
    v30 = (v36 < (v36 & v36))
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint32 = 2147483647
        _ = fi39
        _ = v40
    }
    switch 42 {
    case 4294967295:
        v41 := v6["\n\t"]
        _ = v41
    case 2147483647:
        v3 += 2
    case 127:
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 bool = (127 == -42)
            v44 := GenS5{F0: "\n\t"}
            _ = fi42
            _ = v43
            _ = v44
        }
    }
    switch "" {
    case "a":
        if v30 {
            if (float32(3.14159) <= ((1.5 - v5) - (1e10 / v5))) {
                v38 = v38
                var v45 uint64 = 2
                _ = v45
            } else {
                v46 := GenS6{F0: 196, F1: -128, F2: v3}
                _ = v46
            }
        } else {
            if (v36 > v36) {
                v19 -= (v19 / ((v19) | 1))
                fmt.Println(v30)
            }
        }
    case "foo bar":
        for fi47 := 0; fi47 < 14; fi47++ {
            if v30 {
                v30 = v30
                v36 = (v36 >> ((65535) & 63))
                _ = len(v2)
            } else {
                for fi48 := 0; fi48 < 6; fi48++ {
                    var v49 uint = 2
                    fmt.Println(v19, v3)
                    _ = fi48
                    _ = v49
                }
            }
            var v50 float32 = v5
            _ = fi47
            _ = v50
        }
    case "":
        fmt.Println(v19, v5)
    }
    fmt.Println(v30, v5)
    v51 := GenS7{F0: 4294967295}
    _ = len(v1)
    var v52 string = ("a" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
