package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
func genFn0(p0 bool, p1 int) uint {
    p0 = ((-408 >= 42) || (-1 > 0))
    switch p1 {
    case 958:
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := GenS1{F0: 18446744073709551615}
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint = 838
                v9 := []int{673, -1, 234}
                v10 := GenS2{F0: 32767, F1: -234}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi5
            _ = v6
        }
    case 128:
        fmt.Println(p1)
    case -9223372036854775808:
        if (!true) {
            p1 = (p1 + 2147483647)
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 int16 = -1
                _ = fi11
                _ = v12
            }
            if ((false || (p0 || false)) && false) {
                switch p1 {
                case 127:
                    if true {
                        var v13 uint8 = 128
                        _ = v13
                    }
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    v15 := GenS1{F0: 417}
                    var v16 int32 = 1
                    for fi17 := 0; fi17 < 7; fi17++ {
                        var v18 int8 = -59
                        v19 := []int{1, p1, (fi14 << ((2) & 63))}
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi17, fi14)
                        case "a":
                            var v20 string = ("a" + "foo bar")
                            _ = v20
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                p1 = (p1 | p1)
            } else {
                v21 := GenS1{F0: 6}
                fmt.Println(p0, p1)
                v22 := []int{(p1 + (^(p1))), (p1 % ((p1) | 1)), -529, p1}
                _ = v21
                _ = v22
            }
        } else {
            var v23 uint32 = 2
            v23 += v23
            _ = v23
        }
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint8, p2 int64) int32 {
    var v24 int8 = -128
    if true {
        var v25 float32 = 1e10
        v26 := []int{42, 0, 32767, 2147483647}
        _ = v25
        _ = v26
    } else {
        v27 := []int{int(p1), -657}
        p2 = (-(p2))
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return -622
}

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{0, 0}
    var v3 uint32 = 2
    var v4 float64 = float64(0.1)
    for fi28 := 0; fi28 < 5; fi28++ {
        v1 = ""
        _ = fi28
    }
    v1 = v1
    var v29 int16 = -32768
    if true {
        var v30 int32 = 1
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS3{F0: 1, F1: 93}
            _ = fi31
            _ = v32
        }
        v30 = (v30 & v30)
        _ = v30
    }
    switch v1 {
    case "hello":
        if true {
            v3 = (v3 & 2)
        }
    case "foo bar":
        fmt.Println(v3)
    default:
        var v33 float64 = v4
        _ = v33
    }
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v34 bool = false
    if v34 {
        for fi35 := 0; fi35 < 3; fi35++ {
            if v34 {
                for fi36 := 0; fi36 < 13; fi36++ {
                    v2 = append(v2, v3)
                    var v37 float64 = 0.1
                    _ = fi36
                    _ = v37
                }
                var v38 string = v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v38
            }
            for fi39 := 0; fi39 < 7; fi39++ {
                v1 = ""
                v40 := GenS3{F0: 42, F1: 2}
                _ = fi39
                _ = v40
            }
            _ = fi35
        }
        switch (v1 + "\n\t") {
        case "foo bar":
            v2 = append(v2, (v3 - v3))
        case "hello":
            var v41 bool = ((((-2.25 * float32(-2.25)) / -2.25) - (-2.25 * 0.1)) <= 1e10)
            _ = v41
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            v43 := GenS4{F0: 65535, F1: (-(fi42))}
            for fi44 := 0; fi44 < 10; fi44++ {
                v45 := GenS5{F0: 504}
                v3 -= (v3 * (v3 % ((2) | 1)))
                v4 += (v4 * (float64(-2.25) - ((v4 + v4) + v4)))
                if true {
                    var v46 uint64 = 301
                    v47 := GenS1{F0: 255}
                    _ = v46
                    _ = v47
                }
                _ = fi44
                _ = v45
            }
            v1 = v1
            _ = fi42
            _ = v43
        }
    }
    var v48 uint = 65535
    var v49 int64 = 988
    v50 := GenS6{F0: (1.5 >= ((v4 / v4) + (0.1 / float64((3.14159 * 0.1)))))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
