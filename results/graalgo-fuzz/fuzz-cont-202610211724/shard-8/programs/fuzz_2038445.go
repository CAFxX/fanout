package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 bool }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 bool }
func genFn0(p0 uint16, p1 int64, p2 int32) int32 {
    var v5 uint8 = uint8(p0)
    p2 -= 847
    switch ("hello" + "foo bar") {
    case "hello":
        v6 := []int{4294967295, 128}
        _ = v6
    case "\n\t":
        p0 = p0
    default:
        v7 := []int{-447}
        _ = v7
    }
    v8 := []int{-1, 32767, 128, 263}
    if (1 < -72) {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch "" {
        case "":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        default:
            var v9 uint32 = 65535
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return (p2 / (((p2 >> ((2) & 63))) | 1))
}
func genFn1(p0 uint8, p1 int8, p2 int) int16 {
    if ((^(p1)) >= -1) {
        switch (^(p2)) {
        case 920:
            v10 := []int{32768, p2}
            _ = v10
        case 1:
            v11 := []int{int(p0), p2, (p2 - p2)}
            _ = v11
        case 42:
            if false {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 bool = false
                    fmt.Println(v13, p0)
                    _ = fi12
                    _ = v13
                }
            }
        default:
            v14 := GenS1{F0: 0, F1: p2, F2: 170}
            _ = v14
        }
    }
    switch "" {
    case "foo bar":
        v15 := GenS1{F0: 18446744073709551615, F1: p2, F2: uint16(p0)}
        _ = v15
    }
    v16 := GenS2{F0: 784, F1: p2, F2: (710 != 9223372036854775807)}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -914
}

func main() {
    var v1 uint64 = 316
    v2 := []uint32{598, 924}
    v3 := []string{"a"}
    v4 := []string{("foo bar" + ("foo bar" + (("\n\t" + "foo bar") + (("" + "foo bar") + "foo bar")))), ("\n\t" + ("a" + (("\n\t" + "") + "a"))), "foo bar"}
    if false {
        switch 4294967295 {
        case -632:
            v1 = v1
        case 127:
            var v17 int = 32768
            _ = v17
        default:
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS2{F0: 1, F1: (-(fi18)), F2: (-32768 >= 457)}
                if false {
                    var v20 uint64 = v1
                    _ = v20
                }
                var v21 int8 = 42
                _ = fi18
                _ = v19
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            if len(v4) > 0 {
                v4[0] = ("\n\t" + "\n\t")
            }
            v23 := GenS3{F0: v1, F1: -1, F2: 449}
            _ = fi22
            _ = v23
        }
        if (1 >= genFn1(127, 121, -223)) {
            for fi24 := 0; fi24 < 9; fi24++ {
                switch ("a" + (("" + ("foo bar" + "")) + "hello")) {
                case "":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        switch "a" {
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if (4294967295 == 0) {
                            v1 -= v1
                            v1 += (v1 ^ 629)
                            var v26 uint = 2
                            _ = v26
                        } else {
                            var v27 int16 = -540
                            _ = v27
                        }
                        v28 := GenS4{F0: true, F1: 0}
                        v29 := GenS1{F0: 202, F1: fi24, F2: 342}
                        _ = fi25
                        _ = v28
                        _ = v29
                    }
                case "\n\t":
                    switch (fi24 + 200) {
                    case 127:
                        var v30 int8 = 42
                        _ = v30
                    case 0:
                        v1 -= 18446744073709551615
                    case 9223372036854775807:
                        v1 = v1
                    }
                }
                _ = fi24
            }
            fmt.Println(v1)
            var v31 string = ""
            _ = v31
        }
    }
    if (429 != 1) {
        fmt.Println(v1)
    } else {
        v1 = v1
        v1 -= (v1 >> ((1) & 63))
        v32 := GenS3{F0: (^(v1)), F1: 42, F2: 704}
        _ = v32
    }
    v33 := GenS5{F0: 2147483647}
    var v34 bool = (42 < 42)
    v34 = ((104 == v1) && v34)
    fmt.Println(v34)
    var v35 float32 = 0.1
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS6{F0: 4294967295, F1: 2147483647, F2: 2147483647}
        _ = fi36
        _ = v37
    }
    v38 := GenS7{F0: 4294967295, F1: 960}
    v34 = true
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 float64 = (float64(1e10) * (1.5 / (1e10 + (((-2.25 + (float64(v35) - (1e10 * float64(v35)))) + 3.14159) / 0.1))))
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v41 := GenS8{F0: (v34 && true)}
        v34 = v34
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 8; fi42++ {
        if v34 {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (false && (!true)) {
                v1 -= (v1 & v1)
                var v43 int = fi42
                _ = v43
            } else {
                switch ("foo bar" + "hello") {
                case "":
                    var v44 int32 = -2147483648
                    _ = v44
                case "hello":
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v4 = append(v4, "")
                        _ = len(v3)
                        _ = fi45
                    }
                case "\n\t":
                    var v46 int = fi42
                    _ = v46
                }
                v35 = float32(3.14159)
            }
            for fi47 := 0; fi47 < 6; fi47++ {
                v39 = ((((1.5 / v39) - (float64(-2.25) / ((v39 / (float64(v35) - v39)) - 1.5))) / 3.14159) - 0.1)
                v1 += (v1 - 1)
                _ = fi47
            }
        }
        var v48 int64 = -31
        _ = fi42
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
}
