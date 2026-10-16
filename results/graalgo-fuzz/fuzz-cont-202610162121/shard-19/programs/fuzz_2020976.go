package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 string; F2 int16 }
func genFn0(p0 int, p1 uint64) int8 {
    switch "" {
    case "":
        v7 := GenS1{F0: -1, F1: -60}
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 7; fi8++ {
            fmt.Println(p0, p1)
            p1 = (^(p1))
            _ = fi8
        }
    default:
        if (!(65535 <= 42)) {
            var v9 int16 = -1
            v10 := GenS1{F0: v9, F1: -9223372036854775808}
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 0, F1: -698}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 int8 = 127
        v14 := []int{}
        v15 := GenS2{F0: true}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(p1)
    if false {
        if (172 == (p1 | 18446744073709551615)) {
            var v16 int32 = -1
            _ = v16
        } else {
            v17 := []int{p0, 32767, -322}
            v18 := GenS1{F0: -32768, F1: 42}
            var v19 uint32 = 334
            _ = v17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            switch (fi20 + p0) {
            case 128:
                p0 = p0
            default:
                fmt.Println(fi20)
            }
            v21 := []int{p0}
            var v22 uint64 = 18446744073709551615
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 int64 = -1
        _ = v23
    } else {
        v24 := []int{p0, (p0 - (p0 + p0)), (p0 & p0)}
        p1 = p1
        v25 := GenS3{F0: -947, F1: 18446744073709551615, F2: -2147483648}
        _ = v24
        _ = v25
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 92
}

func main() {
    var v1 uint = 91
    var v2 float32 = ((0.1 - (-2.25 + ((float32(0.0) * 1e10) + 0.0))) / 1e10)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    var v5 float64 = 3.14159
    var v6 string = "\n\t"
    if true {
        fmt.Println(v5, v1)
        if (130 <= 0) {
            _ = len(v4)
        }
    }
    v26 := GenS4{F0: 0}
    for fi27 := 0; fi27 < 13; fi27++ {
        switch v6 {
        case "\n\t":
            v28 := GenS1{F0: 529, F1: -1}
            _ = v28
        case "foo bar":
            if (1 < 4294967295) {
                v1 = 1
            } else {
                v5 -= v5
            }
        default:
            v2 += (v2 * ((1e10 * v2) + (1e10 * v2)))
        }
        fmt.Println(v2, v2, fi27)
        if false {
            var v29 uint8 = uint8(v1)
            _ = v29
        } else {
            v6 = "a"
            v30 := GenS4{F0: (fi27 * (fi27 >> ((65535) & 63)))}
            var v31 uint64 = 1
            _ = v30
            _ = v31
        }
        _ = fi27
    }
    var v32 uint32 = 1
    v33 := GenS5{F0: 9223372036854775807}
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := []int{(fi34 >> (((^(v1))) & 63)), (fi34 * (-(fi34))), (fi34 ^ 1)}
        if false {
            v32 = (v32 | 4294967295)
            v36 := GenS2{F0: ((-1 >= 389) && true)}
            _ = v36
        } else {
            var v37 int = -9223372036854775808
            fmt.Println(v6, v32)
            _ = v37
        }
        if len(v35) > 0 {
            v35[0] = (fi34 ^ 481)
        }
        _ = fi34
        _ = v35
    }
    var v38 uint32 = 2
    switch 998 {
    case 32767:
        var v39 uint8 = 1
        _ = v39
    case 127:
        switch (("hello" + "foo bar") + (((v6 + ((v6 + v6) + v6)) + v6) + v6)) {
        case "foo bar":
            var v40 float64 = v5
            _ = v40
        case "\n\t":
            v41 := GenS6{F0: 42}
            _ = v41
        case "":
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := v3[(fi42 >> ((v1) & 63))]
                _ = fi42
                _ = v43
            }
        default:
            fmt.Println(v6, v6)
        }
    default:
        var v44 float64 = v5
        _ = v44
    }
    v45 := []int{}
    if false {
        v1 -= (v1 * v1)
        v6 = v6
    } else {
        v46 := GenS7{F0: v38, F1: (v6 + v6), F2: 293}
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 float64 = v5
            var v49 float64 = (0.0 + ((float64(v2) / v5) + v48))
            v49 = (v49 * 3.14159)
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v46
    }
    delete(v4, ("foo bar" + "foo bar"))
    switch (v6 + ("foo bar" + (v6 + v6))) {
    case "":
        v2 = 0.0
    case "foo bar":
        for fi50 := 0; fi50 < 16; fi50++ {
            fmt.Println(fi50, v38, fi50)
            _ = fi50
        }
    default:
        for fi51 := 0; fi51 < 1; fi51++ {
            for fi52 := 0; fi52 < 13; fi52++ {
                var v53 uint = 18446744073709551615
                var v54 int = (^(fi51))
                _ = fi52
                _ = v53
                _ = v54
            }
            v3[fi51] = v6
            _ = fi51
        }
    }
    switch (((v6 + "a") + v6) + v6) {
    case "\n\t":
        delete(v4, "")
    case "hello":
        if false {
            var v55 bool = true
            _ = v55
        }
    case "":
        if (-426 != -10) {
            for fi56 := 0; fi56 < 8; fi56++ {
                if true {
                    var v57 int16 = -1
                    _ = v57
                } else {
                    v38 -= 42
                    var v58 int16 = 160
                    v2 = 1.5
                    _ = v58
                }
                for fi59 := 0; fi59 < 16; fi59++ {
                    v2 += v2
                    _ = fi59
                }
                v32 -= 65535
                _ = fi56
            }
        }
    default:
        var v60 uint32 = 65535
        _ = v60
    }
    var v61 int16 = -293
    for fi62 := 0; fi62 < 1; fi62++ {
        var v63 int8 = -8
        var v64 uint = (v1 % (((v1 - v1)) | 1))
        var v65 int64 = -9223372036854775808
        _ = fi62
        _ = v63
        _ = v64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v61
}
