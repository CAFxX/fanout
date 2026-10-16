package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32) int64 {
    if true {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    v5 := GenS1{F0: 32767}
    if true {
        v6 := []int{466, int(p0)}
        if true {
            fmt.Println(p0)
            v7 := GenS2{F0: 2, F1: 2, F2: 236}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 6; fi8++ {
                p0 -= p0
                _ = fi8
            }
            var v9 int64 = 42
            switch "\n\t" {
            case "foo bar":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = v9
        }
        v10 := GenS3{F0: 1, F1: 976, F2: 127}
        _ = v10
        _ = v6
    }
    p0 -= (p0 + (p0 >> ((18446744073709551615) & 63)))
    _ = p0
    _ = v5
    return -9223372036854775808
}
func genFn1() int {
    switch 738 {
    case 0:
        v11 := GenS4{F0: 42}
        _ = v11
    case -537:
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 bool = true
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := []int{fi14, 188}
                fmt.Println(fi12)
                _ = fi14
                _ = v15
            }
            v13 = (0 != -43)
            _ = fi12
            _ = v13
        }
    default:
        var v16 bool = (!true)
        _ = v16
    }
    switch 32768 {
    case -2147483648:
        v17 := []int{-949, -1}
        _ = v17
    case 917:
        var v18 float64 = float64(-2.25)
        _ = v18
    case 9223372036854775807:
        v19 := []int{96, -1, 42}
        _ = v19
    }
    var v20 bool = true
    var v21 uint16 = 462
    var v22 uint32 = 2147483648
    _ = v20
    _ = v21
    _ = v22
    return 776
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 int8 = -128
    var v4 float64 = -2.25
    v3 = 1
    switch 32767 {
    case 127:
        var v23 uint = 342
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        switch genFn1() {
        case 2147483647:
            v1 = true
        default:
            var v25 uint8 = 0
            _ = v25
        }
        v26 := GenS5{F0: 2147483647, F1: 255, F2: 495}
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint32 = 4294967295
            switch "" {
            case "foo bar":
                v3 -= v3
            case "\n\t":
                fmt.Println(fi27, v4, fi27)
            }
            _ = fi27
            _ = v28
        }
        _ = fi24
        _ = v26
    }
    v3 = (v3 >> ((0) & 63))
    var v29 uint64 = 4294967295
    switch "a" {
    case "a":
        v3 = -128
    case "foo bar":
        var v30 float64 = 1e10
        _ = v30
    case "hello":
        v1 = v1
    default:
        if (true && (4294967295 != 4294967295)) {
            switch "\n\t" {
            case "\n\t":
                v1 = v1
            case "foo bar":
                v4 -= (float64(0.1) / (-2.25 * (v4 + (1.5 + (float64(((1.5 + 1e10) - 1e10)) * 0.1)))))
            case "hello":
                var v31 int8 = (^(v3))
                _ = v31
            }
            v29 -= (v29 & v29)
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 uint8 = 127
                _ = fi32
                _ = v33
            }
        }
    }
    v4 = v4
    fmt.Println(v29, v4)
    if ((((0.0 * -2.25) + float32(v4)) + (-2.25 + (0.0 / float32(3.14159)))) >= float32(v4)) {
        v1 = v1
        if v1 {
            var v34 int16 = int16(v29)
            _ = v34
        } else {
            var v35 int16 = 42
            v36 := []int{-76, -1}
            v37 := GenS1{F0: 348}
            _ = v35
            _ = v36
            _ = v37
        }
    } else {
        fmt.Println(v4, v3, v1)
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        if ((!v1) && ("hello" >= (("a" + "foo bar") + "foo bar"))) {
            if (float32(-2.25) >= (1e10 - -2.25)) {
                v39 := GenS5{F0: fi38, F1: 255, F2: -301}
                _ = v39
            }
        } else {
            v40 := []int{}
            switch ("\n\t" + "hello") {
            case "":
                _ = len(v2)
            case "foo bar":
                for fi41 := 0; fi41 < 6; fi41++ {
                    fmt.Println(v3, v3, v1)
                    _ = fi41
                }
            }
            _ = v40
        }
        v1 = v1
        _ = fi38
    }
    v42 := GenS6{F0: 845}
    fmt.Println(v3, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v42
}
