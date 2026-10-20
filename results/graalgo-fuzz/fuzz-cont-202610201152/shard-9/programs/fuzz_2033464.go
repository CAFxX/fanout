package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 int32; F2 int }
type GenS4 struct { F0 uint; F1 uint16; F2 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 int; F1 int }
func genFn0(p0 int64, p1 int) string {
    v4 := []int{(p1 - (-(p1))), (p1 >> ((42) & 63))}
    fmt.Println(p1, p1)
    var v5 int8 = 103
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 uint32) bool {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int32 = -2147483648
        _ = fi6
        _ = v7
    }
    p0 += (p0 * (p0 + p0))
    _ = p0
    return (true || (1 != 1))
}

func main() {
    var v1 int = 127
    v2 := []int{(v1 >> ((2) & 63)), 160, v1, (v1 + -2147483648)}
    var v3 int64 = -1
    if false {
        if false {
            var v8 int = (v1 - (v1 * (v1 % ((v1) | 1))))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v3 = v3
            _ = v8
        } else {
            var v9 int16 = -1
            switch (("foo bar" + ("hello" + "\n\t")) + "foo bar") {
            case "hello":
                v1 = (v1 & v1)
            case "":
                v1 += (v1 ^ -510)
            default:
                v10 := GenS1{F0: ((("foo bar" + "foo bar") >= "a") || true)}
                _ = v10
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v9
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                v3 += v3
                v13 := GenS2{F0: fi12, F1: 2}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    } else {
        v14 := GenS3{F0: true, F1: 2147483647, F2: (v1 << ((4294967295) & 63))}
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = v14
    }
    if ((2 > 4294967295) || (false || false)) {
        fmt.Println(v3, v1)
        v3 += 0
        v3 += v3
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 float32 = 3.14159
            _ = fi15
            _ = v16
        }
        v2 = append(v2, (v1 % ((v1) | 1)))
    }
    v1 = (v1 % (((v1 | v1)) | 1))
    v3 = -9223372036854775808
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (0.0 > (0.1 - 1.5)) {
        v17 := GenS4{F0: 733, F1: 790, F2: (v1 << ((88) & 63))}
        var v18 int64 = v3
        _ = v17
        _ = v18
    } else {
        v19 := GenS5{F0: 1}
        switch ((("" + "hello") + "") + "") {
        case "foo bar":
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int = -809
                _ = fi20
                _ = v21
            }
        case "\n\t":
            for fi22 := 0; fi22 < 4; fi22++ {
                v23 := GenS6{F0: 0, F1: 142}
                _ = fi22
                _ = v23
            }
        case "":
            v3 += v3
        default:
            for fi24 := 0; fi24 < 3; fi24++ {
                v1 = -2147483648
                _ = fi24
            }
        }
        _ = v19
    }
    var v25 bool = (!false)
    fmt.Println(v25, v3)
    v3 -= int64(v1)
    if len(v2) > 0 {
        v2[0] = (v1 - 400)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS7{F0: -357, F1: v1}
    v2 = append(v2, (v1 / ((v1) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
