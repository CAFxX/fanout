package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float32 = ((-2.25 + 3.14159) / 3.14159)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483647
    if true {
        for fi4 := 0; fi4 < 17; fi4++ {
            if ((!false) && (!true)) {
                switch (("hello" + "") + ("foo bar" + "")) {
                case "foo bar":
                    for fi5 := 0; fi5 < 3; fi5++ {
                        v6 := []int{4294967295, fi4, (fi5 ^ fi4)}
                        _ = fi5
                        _ = v6
                    }
                case "\n\t":
                    v1 = (-2.25 / v1)
                default:
                    v7 := GenS1{F0: 42, F1: -1, F2: 2147483648}
                    _ = v7
                }
                switch "foo bar" {
                case "a":
                    switch "a" {
                    case "hello":
                        switch "\n\t" {
                        case "":
                            v1 = 1.5
                        case "a":
                            fmt.Println(v3, v3, v3)
                        default:
                            var v8 float64 = 0.1
                            _ = v8
                        }
                    }
                case "foo bar":
                    fmt.Println(v1, fi4, v1)
                case "":
                    fmt.Println(fi4, fi4)
                default:
                    fmt.Println(v1, fi4)
                }
            } else {
                if true {
                    fmt.Println(fi4)
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v10 := GenS1{F0: 1, F1: -1, F2: 2147483648}
                        _ = fi9
                        _ = v10
                    }
                    v1 -= ((((v1 + v1) - (((v1 * float32(0.0)) + v1) - v1)) * v1) - 0.1)
                } else {
                    var v11 uint = 255
                    for fi12 := 0; fi12 < 9; fi12++ {
                        if (false || ((fi12 & 4294967295) < (-(fi12)))) {
                            v1 += v1
                            var v13 float32 = 1.5
                            _ = v13
                        }
                        _ = fi12
                    }
                    _ = v11
                }
            }
            _ = fi4
        }
        v1 -= v1
        for fi14 := 0; fi14 < 16; fi14++ {
            v3 = (v3 + (v3 >> ((18446744073709551615) & 63)))
            switch "hello" {
            case "foo bar":
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 int64 = 42
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi14
        }
    } else {
        delete(v2, (v3 - 42))
        v1 -= (v1 * float32(0.0))
        switch -9223372036854775808 {
        case -742:
            v17 := []int{32767}
            _ = v17
        case 32767:
            if true {
                v18 := GenS1{F0: int16(v3), F1: 42, F2: (v3 + v3)}
                for fi19 := 0; fi19 < 12; fi19++ {
                    fmt.Println(fi19, fi19)
                    if false {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            fmt.Println(fi19, fi20, v1)
                            fmt.Println(v3)
                            _ = fi20
                        }
                    }
                    var v21 bool = (2 != 18446744073709551615)
                    _ = fi19
                    _ = v21
                }
                v3 = v3
                _ = v18
            }
        default:
            for fi22 := 0; fi22 < 11; fi22++ {
                v3 = v3
                var v23 uint8 = 59
                var v24 int16 = 0
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        if (false && (42 != 1000)) {
            v26 := GenS2{F0: -1}
            if (fi25 != fi25) {
                v27 := GenS3{F0: 127}
                for fi28 := 0; fi28 < 3; fi28++ {
                    v29 := GenS1{F0: 1, F1: 815, F2: (v3 ^ (^(v3)))}
                    var v30 bool = (v1 <= float32(((float64(float32(1.5)) * 1.5) / (0.1 / (((float64(0.1) / 1.0) - (0.0 * 1.5)) - 1.5)))))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v3 = (v3 ^ 4294967295)
                _ = v27
            } else {
                v3 += (v3 ^ v3)
                v1 -= v1
            }
            _ = v26
        }
        var v31 string = ("\n\t" + ("\n\t" + "hello"))
        v32 := []int{(fi25 >> ((1) & 63)), fi25, fi25, fi25}
        _ = fi25
        _ = v31
        _ = v32
    }
    fmt.Println(v3)
    var v33 uint8 = uint8(v3)
    v33 += v33
    v34 := v2[v3]
    fmt.Println(v3, v1, v34)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v34 = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        fmt.Println(v1)
        _ = fi35
    }
    var v36 int32 = 2147483647
    for fi37 := 0; fi37 < 14; fi37++ {
        v38 := GenS3{F0: v33}
        var v39 bool = (250 >= 924)
        _ = fi37
        _ = v38
        _ = v39
    }
    v34 = ((false && (((v36 & (v36 % ((-1) | 1))) < 0) && (v34 || true))) && v34)
    v40 := GenS2{F0: 32767}
    v41 := []int{98, 2147483647, -120}
    v3 = (^(v3))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v40
    _ = v41
}
