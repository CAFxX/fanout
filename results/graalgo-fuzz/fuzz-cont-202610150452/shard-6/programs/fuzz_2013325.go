package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int32; F1 uint64 }
func genFn0(p0 int) bool {
    p0 = p0
    if (((1 <= 9223372036854775807) || false) || false) {
        switch "foo bar" {
        case "":
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 uint64 = 192
                p0 = (fi5 % ((fi5) | 1))
                _ = fi5
                _ = v6
            }
        case "hello":
            for fi7 := 0; fi7 < 2; fi7++ {
                if (!((^(p0)) <= 4294967295)) {
                    fmt.Println(p0, fi7)
                    v8 := GenS1{F0: 18446744073709551615}
                    v9 := []int{(p0 + 178), p0, (fi7 >> ((18446744073709551615) & 63)), p0}
                    _ = v8
                    _ = v9
                }
                v10 := []int{fi7}
                _ = fi7
                _ = v10
            }
        case "foo bar":
            p0 = -660
        default:
            switch -1 {
            case 128:
                var v11 int = (p0 >> ((65535) & 63))
                _ = v11
            case 4294967295:
                fmt.Println(p0)
            }
        }
        p0 += (-(p0))
        p0 = p0
    }
    switch -1 {
    case 42:
        switch "foo bar" {
        case "hello":
            v12 := GenS2{F0: "\n\t", F1: -862}
            _ = v12
        default:
            var v13 uint8 = 38
            _ = v13
        }
    case -9223372036854775808:
        fmt.Println(p0)
    }
    var v14 uint64 = 4294967295
    v15 := []int{(p0 * (p0 >> ((0) & 63))), 574}
    _ = p0
    _ = v14
    _ = v15
    return ((0.0 != (((0.0 - 3.14159) + 0.1) * (-2.25 - 1.5))) && ("a" >= (("hello" + "foo bar") + ("a" + "hello"))))
}
func genFn1() int32 {
    var v16 int32 = 42
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := []int{(fi17 + -2147483648)}
        var v19 uint64 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = v16
    return (v16 & -1)
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    v4 := []uint32{321, 2, 42, 4294967295}
    for fi20 := 0; fi20 < 12; fi20++ {
        if len(v4) > 0 {
            v4[0] = 2147483648
        }
        v21 := GenS2{F0: "foo bar", F1: (v1 << ((4294967295) & 63))}
        fmt.Println(fi20)
        _ = fi20
        _ = v21
    }
    fmt.Println(v1)
    switch "a" {
    case "":
        switch 895 {
        case -30:
            fmt.Println(v1)
        }
    default:
        switch -9223372036854775808 {
        case 4294967295:
            for fi22 := 0; fi22 < 1; fi22++ {
                v1 -= v1
                var v23 uint = 65535
                _ = fi22
                _ = v23
            }
        case -1:
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
    }
    if false {
        for fi24 := 0; fi24 < 10; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                switch "" {
                case "hello":
                    v1 -= v1
                case "a":
                    v1 = v1
                case "":
                    var v26 int8 = 42
                    _ = v26
                }
                v27 := GenS2{F0: "", F1: (v1 - v1)}
                v1 = (v1 - (v1 % ((994) | 1)))
                _ = fi25
                _ = v27
            }
            _ = len(v2)
            _ = fi24
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1)
    v28 := GenS2{F0: "\n\t", F1: (v1 * v1)}
    for fi29 := 0; fi29 < 11; fi29++ {
        if true {
            v1 += (v1 % ((v1) | 1))
            v4 = append(v4, 2147483647)
        } else {
            v30 := v3["a"]
            var v31 int = 9223372036854775807
            _ = v30
            _ = v31
        }
        _ = fi29
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            v1 -= (^(v1))
            _ = fi33
        }
        switch "foo bar" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            v1 -= 0
        case "\n\t":
            v34 := GenS1{F0: 2}
            _ = v34
        default:
            v1 -= v1
        }
        var v35 uint = 2
        _ = fi32
        _ = v35
    }
    fmt.Println(v1)
    var v36 float32 = float32(1e10)
    var v37 uint8 = 0
    v38 := GenS3{F0: 42, F1: 1}
    delete(v3, "")
    var v39 bool = (4294967295 != 0)
    for fi40 := 0; fi40 < 19; fi40++ {
        var v41 float64 = float64(v36)
        var v42 float64 = v41
        var v43 bool = (false || (-2.25 == v36))
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
