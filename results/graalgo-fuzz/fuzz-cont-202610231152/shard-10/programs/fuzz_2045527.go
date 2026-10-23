package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int32, p1 uint16, p2 uint16) uint32 {
    p2 = (p2 + p2)
    var v4 uint16 = (p2 >> ((83) & 63))
    fmt.Println(p0)
    v5 := GenS1{F0: true, F1: 1, F2: 684}
    if (1 != 1) {
        switch 0 {
        case 32768:
            var v6 int8 = -59
            _ = v6
        case 127:
            var v7 int64 = 0
            _ = v7
        case -659:
            fmt.Println(p2, p1)
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                p0 += (p0 | (p0 * 2147483647))
            } else {
                var v9 uint8 = 0
                _ = v9
            }
            v10 := []int{}
            p1 = (p2 / ((v4) | 1))
            _ = fi8
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 uint8, p1 string) int8 {
    if false {
        p1 = p1
    } else {
        if true {
            p0 -= (p0 >> ((18446744073709551615) & 63))
        }
    }
    if false {
        var v11 int32 = -158
        v11 -= v11
        v12 := []int{}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            p0 += (p0 - p0)
            if false {
                switch ((p1 + ((p1 + (("a" + p1) + "\n\t")) + (p1 + (("\n\t" + p1) + p1)))) + (p1 + (p1 + "hello"))) {
                case "":
                    v14 := GenS2{F0: fi13, F1: 18446744073709551615, F2: 42}
                    _ = v14
                case "a":
                    var v15 float32 = 1.5
                    _ = v15
                }
                v16 := []int{}
                _ = v16
            }
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    return 100
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 uint = 4294967295
    v17 := GenS2{F0: 473, F1: 42, F2: 42}
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(fi18)
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 float64 = ((0.0 - ((0.1 * 0.0) / 1.0)) + float64(1.5))
            v2[(-(fi19))] = (-(fi18))
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    switch 223 {
    case -9223372036854775808:
        fmt.Println(v3)
    case 2147483647:
        var v21 int32 = -2147483648
        _ = v21
    case 32767:
        v22 := GenS1{F0: false, F1: -79, F2: v3}
        _ = v22
    default:
        v23 := GenS2{F0: 127, F1: 18446744073709551615, F2: 4294967295}
        _ = v23
    }
    if (0 < 0) {
        var v24 int8 = -53
        if (0 > -2147483648) {
            var v25 int32 = -572
            for fi26 := 0; fi26 < 2; fi26++ {
                var v27 uint = 2
                v28 := GenS1{F0: true, F1: -1, F2: (v3 ^ v3)}
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v25
        } else {
            var v29 float64 = (0.0 * (((1e10 - (((1.5 / 1.5) + -2.25) / (1e10 + (0.1 * 3.14159)))) - (float64(0.0) - 1.5)) / 3.14159))
            _ = v29
        }
        _ = v24
    } else {
        v30 := GenS3{F0: -1, F1: (("\n\t" + ("hello" + (("hello" + "a") + ""))) + "")}
        var v31 uint32 = 65535
        _ = v30
        _ = v31
    }
    var v32 uint8 = 42
    if (9223372036854775807 >= 32768) {
        fmt.Println(v32, v3)
        var v33 float64 = (0.0 / 1.0)
        _ = v33
    } else {
        var v34 uint = v3
        switch -9223372036854775808 {
        case 0:
            delete(v1, 32767)
        case 127:
            v32 -= v32
        case 4294967295:
            switch (("" + ("\n\t" + "a")) + "foo bar") {
            case "foo bar":
                var v35 int = 32768
                _ = v35
            }
        }
        var v36 float64 = (0.0 / 1e10)
        _ = v34
        _ = v36
    }
    v3 += v3
    fmt.Println(v32, v3)
    var v37 uint64 = 545
    var v38 uint16 = 611
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v38
}
