package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
func genFn0(p0 int) int64 {
    v5 := []int{(p0 | p0), (p0 % ((32767) | 1))}
    v6 := GenS1{F0: p0, F1: -2147483648}
    for fi7 := 0; fi7 < 10; fi7++ {
        switch ("" + ("foo bar" + ("foo bar" + "foo bar"))) {
        case "foo bar":
            p0 = (p0 & (fi7 & p0))
        }
        p0 += (fi7 << ((2) & 63))
        _ = fi7
    }
    if (false && ((32767 == 42) && (!(423 < 42)))) {
        if ((((0.1 + float32(1e10)) / 1e10) + 1e10) != float32(float64((-2.25 - ((3.14159 * 0.1) * 0.0))))) {
            for fi8 := 0; fi8 < 6; fi8++ {
                switch ("foo bar" + "hello") {
                case "":
                    var v9 uint64 = 4294967295
                    _ = v9
                default:
                    var v10 int32 = 918
                    _ = v10
                }
                _ = fi8
            }
            v11 := GenS1{F0: (p0 / (((p0 & (p0 << ((774) & 63)))) | 1)), F1: 718}
            _ = len(v5)
            _ = v11
        } else {
            var v12 int64 = 9223372036854775807
            var v13 bool = false
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            p0 += p0
            p0 = 1
            _ = fi14
        }
    }
    v15 := GenS2{F0: 566, F1: 569}
    _ = p0
    _ = v15
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 uint32, p2 uint16) string {
    switch "foo bar" {
    case "hello":
        for fi16 := 0; fi16 < 10; fi16++ {
            p1 += p1
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int64 = genFn0(fi17)
                v18 = v18
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    case "foo bar":
        var v19 int = -992
        _ = v19
    }
    p0 = (p0 && (!(((0.1 - 0.1) * -2.25) == -2.25)))
    for fi20 := 0; fi20 < 13; fi20++ {
        switch "a" {
        case "a":
            var v21 uint32 = (p1 + (p1 + (p1 >> ((698) & 63))))
            _ = v21
        case "":
            var v22 uint8 = 216
            _ = v22
        case "\n\t":
            p0 = true
        default:
            fmt.Println(p2)
        }
        fmt.Println(p2, fi20, p0)
        _ = fi20
    }
    if false {
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := []int{fi23, fi23, fi23, -780}
            for fi25 := 0; fi25 < 11; fi25++ {
                if (42 != 2147483647) {
                    _ = len(v24)
                }
                p2 = 42
                p0 = (311 >= -1)
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        v26 := GenS2{F0: 4294967295, F1: 279}
        _ = v26
    } else {
        p0 = (!false)
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint8 = 157
    var v3 uint64 = 42
    var v4 int64 = 42
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint = 440
        _ = fi27
        _ = v28
    }
    v29 := GenS1{F0: 42, F1: -2147483648}
    if (!((((true && true) && true) || false) || ((1 >= 42) && ((!true) || false)))) {
        if (42 < (v3 % ((v3) | 1))) {
            if true {
                var v30 int = 1
                _ = v30
            } else {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS3{F0: 544}
                    _ = fi31
                    _ = v32
                }
            }
            v33 := GenS3{F0: 42}
            for fi34 := 0; fi34 < 12; fi34++ {
                fmt.Println(v4, v4)
                _ = fi34
            }
            _ = v33
        }
        _ = len(v1)
        switch "\n\t" {
        case "hello":
            v3 = (^(v3))
        case "foo bar":
            var v35 bool = true
            _ = v35
        case "\n\t":
            switch -2147483648 {
            case 32767:
                var v36 float64 = 1e10
                _ = v36
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        var v37 float64 = 0.1
        v37 += 0.0
        _ = v37
    }
    v38 := GenS1{F0: 995, F1: -88}
    _ = len(v1)
    var v39 int32 = -2147483648
    v39 += v39
    fmt.Println(v39)
    for fi40 := 0; fi40 < 15; fi40++ {
        if false {
            v41 := GenS4{F0: -107}
            _ = v41
        }
        v3 -= 4294967295
        fmt.Println(v3, v2, v3)
        _ = fi40
    }
    if false {
        fmt.Println(v4, v39, v3)
        v4 -= 42
        v1 = append(v1, 65535)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
