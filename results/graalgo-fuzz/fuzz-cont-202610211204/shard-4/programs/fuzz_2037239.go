package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint, p1 int32) uint {
    p0 -= 2
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 19; fi8++ {
        p0 = p0
        switch "hello" {
        case "\n\t":
            var v9 uint64 = 2
            _ = v9
        case "hello":
            switch ("hello" + "a") {
            case "a":
                if false {
                    p1 = p1
                    var v10 uint = (p0 >> ((p0) & 63))
                    var v11 uint = v10
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        switch "hello" {
                        case "a":
                            var v13 int = (fi12 << ((p0) & 63))
                            _ = v13
                        case "foo bar":
                            fmt.Println(fi8, p1)
                        default:
                            var v14 int64 = 91
                            _ = v14
                        }
                        var v15 bool = (!true)
                        _ = fi12
                        _ = v15
                    }
                    fmt.Println(p0, fi8, p1)
                    fmt.Println(p1, fi8, v7)
                }
            case "\n\t":
                fmt.Println(v7, v7, p0)
            case "hello":
                var v16 uint64 = 42
                _ = v16
            default:
                var v17 uint32 = 0
                _ = v17
            }
        case "":
            p1 += p1
        default:
            var v18 int64 = 9223372036854775807
            _ = v18
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return (^(p0))
}
func genFn1(p0 string, p1 int) int32 {
    if (-1 >= 42) {
        v19 := GenS1{F0: 820, F1: (p1 & 1)}
        var v20 int8 = 0
        fmt.Println(v20)
        _ = v19
        _ = v20
    }
    switch p0 {
    case "foo bar":
        for fi21 := 0; fi21 < 12; fi21++ {
            p1 = (fi21 * (fi21 + (fi21 / (((fi21 % ((-1) | 1))) | 1))))
            v22 := []int{32767}
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    return 366
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 42
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]string)
    var v5 float64 = 1e10
    v6 := make(map[uint32]uint32)
    v2 += v2
    var v23 string = (((("hello" + "hello") + "foo bar") + "hello") + ("hello" + "hello"))
    var v24 uint64 = (v3 >> ((0) & 63))
    switch "" {
    case "\n\t":
        var v25 bool = (!true)
        _ = v25
    case "foo bar":
        v2 += (v2 * 2)
    case "hello":
        v26 := v6[0]
        _ = v26
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 bool = true
        for fi29 := 0; fi29 < 5; fi29++ {
            v24 += v3
            if (0 < uint32(v24)) {
                v30 := GenS2{F0: 2, F1: v2}
                _ = v30
            }
            v2 -= (v2 % (((^(v2))) | 1))
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    if true {
        var v31 int8 = 33
        if true {
            fmt.Println(v3, v23, v2)
        } else {
            v32 := GenS3{F0: 32767, F1: ("" + v23), F2: -1}
            var v33 uint16 = uint16(v24)
            var v34 uint16 = 255
            _ = v32
            _ = v33
            _ = v34
        }
        _ = v31
    }
    if false {
        for fi35 := 0; fi35 < 7; fi35++ {
            _ = len(v4)
            if (2 == 159) {
                var v36 int32 = -617
                _ = v36
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 float32 = 0.1
                v39 := GenS3{F0: 172, F1: v23, F2: -761}
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi35
        }
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        switch int(v2) {
        case 127:
            v3 = 1
        case 1:
            v3 -= (v3 & 2)
        case -2147483648:
            v41 := GenS1{F0: 136, F1: (fi40 / ((fi40) | 1))}
            _ = v41
        default:
            if false {
                v3 = (v3 >> ((4294967295) & 63))
                var v42 int64 = -1
                v43 := GenS2{F0: 2, F1: 2}
                _ = v42
                _ = v43
            }
        }
        _ = fi40
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v24 -= v24
        _ = fi44
    }
    var v45 int64 = 0
    v24 -= v3
    v3 = 2
    if false {
        delete(v4, 1)
        v46 := []int{-331, 128}
        _ = v46
    } else {
        fmt.Println(v2, v45, v23)
        fmt.Println(v5, v5)
        switch int(v3) {
        case -254:
            v3 -= (v3 - 1)
        case 127:
            fmt.Println(v5)
        }
    }
    if false {
        if (0 >= 2) {
            v45 -= -110
        } else {
            for fi47 := 0; fi47 < 15; fi47++ {
                fmt.Println(v2, v24)
                switch fi47 {
                case 127:
                    var v48 bool = (18446744073709551615 == (v24 - v24))
                    _ = v48
                case -648:
                    fmt.Println(v5, v24)
                default:
                    v2 = 1
                }
                _ = fi47
            }
            v49 := []int{-943, -9223372036854775808, -759}
            _ = v49
        }
        v3 = v3
        v50 := GenS4{F0: 100}
        _ = v50
    }
    switch 4 {
    case 738:
        v51 := GenS5{F0: 676, F1: 1}
        _ = v51
    case 128:
        v3 = (v3 << ((42) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
