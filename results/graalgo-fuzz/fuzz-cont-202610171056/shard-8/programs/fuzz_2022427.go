package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 string; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8) uint64 {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 255, F1: false}
    var v8 uint16 = 885
    v9 := []int{-9223372036854775808, 621, -1, -7}
    fmt.Println(v8, v6)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 int64) int8 {
    if (-2.25 == 1.5) {
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                p0 += p0
                p0 = (p0 * p0)
                var v12 uint32 = 65535
                _ = fi11
                _ = v12
            }
            switch (fi10 - (fi10 << ((42) & 63))) {
            case 13:
                v13 := GenS2{F0: 32767, F1: true}
                _ = v13
            case 602:
                v14 := []int{-690, (fi10 | 1), fi10, 42}
                _ = v14
            case 32768:
                var v15 float64 = 0.0
                _ = v15
            }
            _ = fi10
        }
        var v16 uint16 = 823
        if true {
            p1 -= p1
        }
        _ = v16
    }
    p0 -= p0
    p1 -= 9223372036854775807
    switch "\n\t" {
    case "a":
        for fi17 := 0; fi17 < 12; fi17++ {
            p1 = (p1 % ((p1) | 1))
            switch (fi17 / ((fi17) | 1)) {
            case -9223372036854775808:
                if true {
                    if true {
                        fmt.Println(fi17, p1, p1)
                    } else {
                        var v18 uint16 = 0
                        v19 := GenS3{F0: 42}
                        v18 = 992
                        _ = v18
                        _ = v19
                    }
                    var v20 int64 = (p1 / ((9223372036854775807) | 1))
                    _ = v20
                } else {
                    p0 -= (p0 & 173)
                    var v21 uint64 = 4294967295
                    v22 := GenS4{F0: false, F1: (("foo bar" + "") + ("" + "")), F2: genFn0(p0)}
                    _ = v21
                    _ = v22
                }
            case -940:
                v23 := []int{(fi17 % ((fi17) | 1)), -1, fi17}
                _ = v23
            default:
                fmt.Println(p0, fi17)
            }
            _ = fi17
        }
    case "":
        for fi24 := 0; fi24 < 12; fi24++ {
            switch (fi24 + 253) {
            case -2147483648:
                switch (fi24 ^ 9223372036854775807) {
                case -9223372036854775808:
                    v25 := []int{fi24, (^(fi24)), 9223372036854775807, 128}
                    _ = v25
                case -635:
                    if true {
                        p0 = (^(p0))
                    } else {
                        for fi26 := 0; fi26 < 16; fi26++ {
                            v27 := []int{}
                            p1 += (p1 & (p1 >> ((961) & 63)))
                            p1 += p1
                            p0 -= (p0 ^ (p0 >> ((18446744073709551615) & 63)))
                            _ = fi26
                            _ = v27
                        }
                        var v28 float32 = 0.0
                        _ = v28
                    }
                }
            case 42:
                if (839 < 834) {
                    if false {
                        var v29 uint32 = 1
                        v29 -= (v29 & 42)
                        _ = v29
                    }
                    p0 = p0
                    if true {
                        v30 := GenS5{F0: (p1 + (p1 & 839))}
                        _ = v30
                    }
                } else {
                    v31 := GenS3{F0: 212}
                    fmt.Println(fi24, fi24, fi24)
                    fmt.Println(p1)
                    _ = v31
                }
            }
            _ = fi24
        }
    case "\n\t":
        fmt.Println(p1)
    default:
        v32 := GenS6{F0: ((!true) && true), F1: 767}
        _ = v32
    }
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("\n\t" + ((("hello" + (("foo bar" + "\n\t") + "a")) + "hello") + "a")), ("a" + ("a" + "hello")), "a"}
    v3 := make(map[string]int)
    var v4 uint64 = 257
    v5 := make(map[uint32]uint32)
    v4 = 2
    var v33 uint16 = 0
    var v34 string = "foo bar"
    for fi35 := 0; fi35 < 11; fi35++ {
        v34 = (v34 + v34)
        fmt.Println(fi35, v4)
        _ = fi35
    }
    switch 913 {
    case 32767:
        var v36 int = -699
        _ = v36
    case 881:
        v33 -= 0
    }
    var v37 string = (v34 + "foo bar")
    v5[42] = 42
    delete(v1, 9223372036854775807)
    v4 = v4
    fmt.Println(v34, v37)
    v33 -= (v33 << ((65535) & 63))
    v38 := GenS1{F0: uint(v4), F1: ((((float32((1.5 * (1e10 * -2.25))) - 1e10) + 1.5) != ((0.1 + 3.14159) * (-2.25 - 1.5))) || (!false))}
    for fi39 := 0; fi39 < 12; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            if (42 != (v33 & (v33 / ((v33) | 1)))) {
                var v41 uint8 = 2
                _ = v41
            }
            var v42 bool = ((v4 / ((genFn0(128)) | 1)) == (v4 ^ v4))
            var v43 int64 = 9223372036854775807
            _ = fi40
            _ = v42
            _ = v43
        }
        fmt.Println(v37, fi39, fi39)
        var v44 float64 = float64(-2.25)
        _ = fi39
        _ = v44
    }
    _ = len(v2)
    v45 := GenS6{F0: (13 != -1), F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
}
