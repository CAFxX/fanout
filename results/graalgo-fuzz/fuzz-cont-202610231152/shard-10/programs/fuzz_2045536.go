package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint16) uint {
    v4 := GenS1{F0: ((false && (!true)) && true), F1: 65535}
    p0 -= p0
    v5 := GenS1{F0: (!(("a" + ("\n\t" + "hello")) >= "hello")), F1: 65535}
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 string, p1 int64, p2 uint64) string {
    p1 = (p1 + p1)
    if false {
        v6 := []int{32767, -9223372036854775808, 128, -2147483648}
        var v7 float64 = float64(-2.25)
        if len(v6) > 0 {
            v6[0] = -9223372036854775808
        }
        _ = v6
        _ = v7
    } else {
        if ((((1e10 - 1.5) + (((1.5 + (float32(1e10) / ((1e10 + float32((1.5 + -2.25))) / 1.5))) + (1e10 + 1e10)) + 1.5)) * (1.5 * (0.0 * 0.1))) > ((3.14159 - ((-2.25 / 1e10) - (-2.25 / 1.0))) / float32(0.1))) {
            var v8 int32 = -1
            _ = v8
        } else {
            var v9 uint = genFn0(65535)
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if false {
                v11 := []int{-1, fi10, (fi10 - fi10)}
                fmt.Println(fi10, p2)
                _ = v11
            } else {
                v12 := GenS1{F0: true, F1: genFn0(32767)}
                _ = v12
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS2{F0: 65535, F1: -445}
                v15 := []int{fi10, (fi13 + (fi13 % (((fi13 & fi10)) | 1))), (fi10 >> ((18446744073709551615) & 63)), fi13}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                fmt.Println(fi13)
                p1 -= p1
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi10
        }
        switch -2147483648 {
        case 2147483647:
            p0 = p0
        default:
            p2 = (p2 * p2)
        }
    }
    switch p0 {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        v16 := GenS3{F0: (!true)}
        _ = v16
    case "":
        var v17 bool = ((float64(float32(((-2.25 - (0.1 + float64((float32(1e10) + 1e10)))) * (0.1 / ((0.1 / 1.0) / 1.5))))) - 0.0) >= 1.5)
        _ = v17
    }
    if (2147483647 < -2147483648) {
        switch (("a" + ("foo bar" + (p0 + "a"))) + ("hello" + p0)) {
        case "\n\t":
            switch "foo bar" {
            case "":
                var v18 uint = 18446744073709551615
                _ = v18
            }
        case "hello":
            p1 -= p1
        default:
            fmt.Println(p0)
        }
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []int{9223372036854775807, 4294967295, 4294967295}
    var v2 int = 32767
    var v3 int8 = 1
    if (!false) {
        for fi19 := 0; fi19 < 1; fi19++ {
            if false {
                v3 = (v3 >> ((genFn0(1)) & 63))
                v3 = 1
            }
            _ = fi19
        }
    } else {
        fmt.Println(v3, v2)
        if len(v1) > 0 {
            v1[0] = (v2 / ((v2) | 1))
        }
        if false {
            if ((42 <= 32767) || true) {
                v20 := GenS3{F0: true}
                _ = v20
            }
            switch "foo bar" {
            case "hello":
                var v21 int64 = -1
                _ = v21
            case "":
                switch 128 {
                case 940:
                    var v22 float32 = 1.5
                    _ = v22
                case 149:
                    fmt.Println(v2, v2)
                }
            }
            v23 := GenS1{F0: (!false), F1: 112}
            _ = v23
        } else {
            v2 = v2
            var v24 int8 = v3
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint64 = 1
            fmt.Println(fi26, fi25, fi25)
            fmt.Println(v3, v2, v27)
            _ = fi26
            _ = v27
        }
        v3 -= (v3 / (((v3 ^ (v3 << ((0) & 63)))) | 1))
        _ = fi25
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        v3 = -1
        v2 += 951
        v1 = append(v1, -2147483648)
        _ = fi28
    }
    fmt.Println(v2, v2)
    for fi29 := 0; fi29 < 2; fi29++ {
        fmt.Println(v2, v2, v2)
        _ = fi29
    }
    fmt.Println(v2)
    var v30 bool = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = (v2 - 4294967295)
    }
    v1 = append(v1, (v2 | v2))
    var v31 int32 = -2147483648
    v32 := GenS2{F0: 174, F1: 355}
    v33 := GenS3{F0: (v30 && ((((float64(0.0) + -2.25) + 0.0) >= 1e10) || v30))}
    var v34 int8 = v3
    if (true && false) {
        switch (v2 >> ((239) & 63)) {
        case 265:
            switch v2 {
            case -605:
                for fi35 := 0; fi35 < 9; fi35++ {
                    v31 += v31
                    _ = fi35
                }
            default:
                v2 = 2147483647
            }
        case 32768:
            var v36 int = (v2 ^ (^(v2)))
            _ = v36
        case 857:
            v31 = (v31 + -2147483648)
        }
    } else {
        for fi37 := 0; fi37 < 14; fi37++ {
            v1 = append(v1, fi37)
            v31 = 788
            _ = fi37
        }
    }
    if len(v1) > 0 {
        v1[0] = (v2 & (v2 * 4294967295))
    }
    var v38 float64 = 1.5
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
}
