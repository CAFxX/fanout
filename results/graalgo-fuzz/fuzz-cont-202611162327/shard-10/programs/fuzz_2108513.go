package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 int16, p1 uint) string {
    if (!(-1 != -298)) {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch fi7 {
            case 4294967295:
                v8 := []int{(fi7 / ((9223372036854775807) | 1)), -611, fi7, fi7}
                _ = v8
            default:
                var v9 float64 = (1e10 - 1e10)
                _ = v9
            }
            fmt.Println(p0)
            p0 = p0
            _ = fi7
        }
    }
    if (true && true) {
        for fi10 := 0; fi10 < 15; fi10++ {
            switch "\n\t" {
            case "\n\t":
                var v11 int64 = 42
                _ = v11
            default:
                var v12 int16 = 712
                _ = v12
            }
            v13 := []int{fi10}
            var v14 uint = (p1 + 366)
            _ = fi10
            _ = v13
            _ = v14
        }
    }
    p1 = (^(p1))
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint16 = 553
        switch "\n\t" {
        case "foo bar":
            v17 := []int{(fi15 % ((42) | 1)), fi15}
            _ = v17
        case "":
            var v18 uint = (^(p1))
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(p1, fi15)
                p1 += (^(p1))
                v16 += (v16 - v16)
                _ = fi19
            }
        }
        _ = fi15
        _ = v16
    }
    if (p0 < 695) {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint8 = 92
            var v22 int8 = 127
            if ((p1 * uint(v21)) < (^(p1))) {
                var v23 uint32 = 4294967295
                _ = v23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        p0 += p0
        v24 := []int{128, 42}
        _ = v24
    }
    _ = p0
    _ = p1
    return ("foo bar" + "")
}

func main() {
    var v1 uint = 0
    var v2 int = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 int16 = 42
    var v5 bool = false
    var v6 int = v2
    var v25 int32 = -307
    if v5 {
        var v26 uint32 = 0
        var v27 bool = (v5 || (!(42 > 426)))
        var v28 uint8 = 241
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 string = ""
    _ = len(v3)
    if (-1 >= 42) {
        var v30 float64 = (1e10 + 1e10)
        _ = v30
    }
    v2 += (-(v2))
    fmt.Println(v29)
    var v31 string = v29
    if v5 {
        if ((-1 != 0) && (0 < 0)) {
            v6 = (v2 & 4294967295)
            if (-973 < 378) {
                v32 := GenS1{F0: ("a" + (v31 + "")), F1: (v1 | (v1 << (((^(v1))) & 63))), F2: 127}
                _ = v32
            }
        }
        if (926 >= 1) {
            var v33 int32 = -988
            var v34 int64 = 468
            _ = v33
            _ = v34
        }
    } else {
        fmt.Println(v1, v5, v2)
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        v1 = (v1 << (((v1 * v1)) & 63))
        _ = fi35
    }
    fmt.Println(v25)
    v36 := GenS2{F0: false, F1: 42, F2: 9223372036854775807}
    for fi37 := 0; fi37 < 10; fi37++ {
        delete(v3, 42)
        _ = fi37
    }
    var v38 int16 = v4
    _ = len(v3)
    switch v6 {
    case 0:
        if v5 {
            _ = len(v3)
            var v39 int8 = -1
            _ = v39
        }
    case 850:
        var v40 int64 = -9223372036854775808
        _ = v40
    default:
        v29 = v31
    }
    fmt.Println(v5, v38)
    v41 := GenS3{F0: 127, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v36
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
