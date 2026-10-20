package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint16, p2 int8) uint16 {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{}
        switch fi7 {
        case 2147483647:
            var v9 uint = 1
            _ = v9
        case 1:
            p0 = -1
        }
        if false {
            if ((true || (1e10 != -2.25)) && (0 > 128)) {
                v8 = append(v8, fi7)
                var v10 float64 = (float64(1.5) / 3.14159)
                _ = v10
            }
            p1 = p1
        } else {
            v11 := GenS1{F0: ((3.14159 <= (-2.25 - 3.14159)) && false), F1: (p2 ^ -128), F2: p1}
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS2{F0: p0}
    v13 := GenS3{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (p1 - (p1 - p1))
}

func main() {
    var v1 int32 = 42
    var v2 float64 = 1e10
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483648
    var v5 int16 = -932
    v6 := []string{("hello" + "\n\t"), "foo bar"}
    _ = len(v6)
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int8 = 127
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v4)
    for fi16 := 0; fi16 < 10; fi16++ {
        _ = len(v6)
        _ = fi16
    }
    switch 42 {
    case -467:
        var v17 float64 = v2
        _ = v17
    case 128:
        v2 = float64(0.1)
    default:
        v5 = v5
    }
    fmt.Println(v5, v5, v4)
    for fi18 := 0; fi18 < 12; fi18++ {
        if true {
            v5 += 0
            var v19 uint32 = (v4 << ((4294967295) & 63))
            fmt.Println(v2)
            _ = v19
        }
        _ = fi18
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        _ = len(v3)
        if (!true) {
            v1 = (v1 - 760)
        } else {
            if len(v6) > 0 {
                v6[0] = "hello"
            }
        }
        _ = fi20
    }
    v4 -= v4
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(501 <= 42)) {
        v2 += v2
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 uint = uint(v4)
                fmt.Println(fi21)
                _ = fi22
                _ = v23
            }
            var v24 string = "hello"
            _ = fi21
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint32 = 2147483647
            var v27 int32 = (v1 >> ((18446744073709551615) & 63))
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS2{F0: -1}
            v4 = v4
            v4 += 4294967295
            _ = fi28
            _ = v29
        }
    }
    if true {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
