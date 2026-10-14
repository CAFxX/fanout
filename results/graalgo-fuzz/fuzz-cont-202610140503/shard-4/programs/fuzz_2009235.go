package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }
func genFn0(p0 string) uint64 {
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 = p0
        var v6 int8 = 127
        p0 = "a"
        _ = fi5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    return 4294967295
}
func genFn1(p0 uint16, p1 bool) uint16 {
    var v7 uint16 = p0
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(p0, p1, p1)
        fmt.Println(fi8)
        var v9 uint64 = 42
        _ = fi8
        _ = v9
    }
    if (p0 < 2) {
        var v10 uint8 = 255
        var v11 float64 = (0.0 - 1e10)
        _ = v10
        _ = v11
    } else {
        if ((p1 && false) || p1) {
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{(fi12 << ((255) & 63))}
                var v14 int64 = 593
                v14 = (v14 >> ((255) & 63))
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 uint8 = 128
                fmt.Println(p0, v7)
                _ = fi15
                _ = v16
            }
            var v17 bool = false
            _ = v17
        }
        var v18 int16 = 1
        for fi19 := 0; fi19 < 20; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                v21 := GenS1{F0: (^(p0)), F1: (p0 >> ((4294967295) & 63)), F2: (^(v18))}
                v18 = -313
                var v22 int32 = -893
                _ = fi20
                _ = v21
                _ = v22
            }
            v7 -= (v7 << ((18446744073709551615) & 63))
            _ = fi19
        }
        _ = v18
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 6; fi24++ {
            p0 += (^(v7))
            v25 := GenS2{F0: 0, F1: 127}
            _ = fi24
            _ = v25
        }
        fmt.Println(p1)
        var v26 int8 = -1
        _ = fi23
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v7
    return (v7 / ((v7) | 1))
}

func main() {
    var v1 bool = false
    var v2 int32 = 42
    var v3 int16 = -32768
    v4 := make(map[int]int)
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS3{F0: "a"}
        var v29 uint8 = 255
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint64 = 1
    for fi31 := 0; fi31 < 12; fi31++ {
        fmt.Println(v2)
        v32 := GenS4{F0: 171}
        _ = fi31
        _ = v32
    }
    if (((1.5 + ((float64(3.14159) + float64(((0.1 + 1e10) + 1.5))) + 0.0)) != 0.1) || false) {
        switch "" {
        case "a":
            _ = len(v4)
        case "hello":
            switch "foo bar" {
            case "a":
                v33 := GenS4{F0: 65535}
                _ = v33
            case "\n\t":
                fmt.Println(v30, v30, v2)
            case "hello":
                v1 = v1
            }
        }
    } else {
        v34 := []int{}
        for fi35 := 0; fi35 < 18; fi35++ {
            for fi36 := 0; fi36 < 16; fi36++ {
                v37 := GenS1{F0: 0, F1: 405, F2: 307}
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        switch 4294967295 {
        case 32768:
            v3 -= v3
        case 42:
            v30 += 64
        case 9223372036854775807:
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 float32 = (((0.0 - 0.1) - -2.25) + (float32(1e10) - ((0.0 - (-2.25 + 1e10)) - 0.1)))
                var v40 uint16 = 65535
                _ = fi38
                _ = v39
                _ = v40
            }
        default:
            var v41 uint = 476
            _ = v41
        }
        _ = v34
    }
    v1 = true
    for fi42 := 0; fi42 < 20; fi42++ {
        switch "a" {
        case "":
            var v43 float32 = (1.5 - float32(((1e10 / -2.25) / ((1.5 * (float64(1.5) - (float64(-2.25) + ((3.14159 - 1.5) - (-2.25 * float64(0.1)))))) * 3.14159))))
            _ = v43
        case "\n\t":
            switch fi42 {
            case -867:
                v2 += -1
            case 1:
                v44 := []int{}
                _ = v44
            case 42:
                v45 := GenS5{F0: ""}
                _ = v45
            default:
                var v46 int = (fi42 * fi42)
                _ = v46
            }
        case "a":
            var v47 int = 9223372036854775807
            _ = v47
        }
        if true {
            fmt.Println(v1)
            fmt.Println(v30, v2, fi42)
        } else {
            var v48 uint8 = 0
            fmt.Println(v30, fi42, v30)
            v30 -= genFn0("a")
            _ = v48
        }
        if (0 < 42) {
            var v49 uint = 65535
            _ = v49
        }
        _ = fi42
    }
    v30 += v30
    var v50 int64 = 9223372036854775807
    var v51 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v50
    _ = v51
}
