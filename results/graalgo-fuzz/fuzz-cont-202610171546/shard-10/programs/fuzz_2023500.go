package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int, p1 int64) string {
    v5 := GenS1{F0: -2147483648, F1: 42}
    var v6 uint = 65535
    v6 = (v6 >> ((v6) & 63))
    v7 := []int{(p0 * p0), p0, (p0 << ((42) & 63))}
    for fi8 := 0; fi8 < 10; fi8++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v9 int32 = 964
        v10 := GenS1{F0: 2147483647, F1: -128}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 bool, p1 int, p2 string) uint16 {
    p2 = ("hello" + ("a" + p2))
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch (fi12 ^ 32768) {
            case 38:
                v13 := GenS2{F0: 42}
                _ = v13
            case 4294967295:
                v14 := []int{fi11}
                _ = v14
            case 42:
                v15 := []int{}
                _ = v15
            }
            _ = fi12
        }
        _ = fi11
    }
    var v16 int64 = 0
    var v17 bool = p0
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 255
}

func main() {
    var v1 int16 = 809
    var v2 int = 42
    var v3 uint32 = 2147483647
    v4 := make(map[int]string)
    var v18 float32 = -2.25
    var v19 int = 32767
    if true {
        for fi20 := 0; fi20 < 2; fi20++ {
            v19 = (v2 + 0)
            v19 = (v19 | (v2 - v2))
            var v21 int8 = -1
            _ = fi20
            _ = v21
        }
    }
    v19 = 9223372036854775807
    v3 = (v3 | v3)
    v18 += v18
    fmt.Println(v18, v3, v2)
    if false {
        v1 -= v1
        v19 -= (v2 >> ((0) & 63))
        v3 += (v3 * (v3 + 4294967295))
    }
    var v22 int32 = 860
    fmt.Println(v18)
    if ((v19 << ((65535) & 63)) == 1) {
        fmt.Println(v3, v1)
        fmt.Println(v3)
    } else {
        v4[(v2 % ((128) | 1))] = ("hello" + "")
        fmt.Println(v19)
    }
    var v23 float32 = (v18 - 3.14159)
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: (v22 * (v22 ^ (v22 ^ (v22 % ((int32(fi24)) | 1))))), F1: -1}
            _ = fi25
            _ = v26
        }
        delete(v4, -696)
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 float64 = (1e10 + 0.0)
            v19 = (fi24 & fi24)
            _ = fi27
            _ = v28
        }
        _ = fi24
    }
    if (v22 <= v22) {
        v22 = -2147483648
        fmt.Println(v2, v18)
    } else {
        switch 128 {
        case -373:
            var v29 float64 = 0.0
            _ = v29
        default:
            var v30 float64 = 3.14159
            _ = v30
        }
        switch (("" + ("" + ("foo bar" + "foo bar"))) + "hello") {
        case "a":
            switch "" {
            case "\n\t":
                _ = len(v4)
            default:
                v31 := GenS2{F0: 42}
                _ = v31
            }
        default:
            if (false && true) {
                if false {
                    fmt.Println(v19)
                    var v32 int16 = (^(v1))
                    _ = v32
                } else {
                    if ((!(v18 >= (float32(3.14159) + v23))) && true) {
                        fmt.Println(v1, v2)
                    }
                }
                v33 := v4[0]
                _ = v33
            }
        }
    }
    v19 = v19
    switch "a" {
    case "":
        fmt.Println(v23)
    case "foo bar":
        v34 := GenS3{F0: 32767, F1: (v22 + (v22 / ((v22) | 1))), F2: 78}
        _ = v34
    default:
        v35 := GenS4{F0: (true && (1.5 > (1e10 - (1e10 - ((-2.25 / (0.0 + (0.1 / -2.25))) * ((-2.25 + 3.14159) + float64(v18))))))), F1: 42}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
