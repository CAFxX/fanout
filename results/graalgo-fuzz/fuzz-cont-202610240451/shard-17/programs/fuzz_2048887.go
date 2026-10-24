package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint64 }
func genFn0(p0 int32, p1 uint32) int64 {
    v7 := GenS1{F0: 42}
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (v8 * (v8 | -9223372036854775808))
}
func genFn1(p0 uint, p1 string) string {
    switch p1 {
    case "foo bar":
        if false {
            p0 = p0
            if (false || ((!true) && false)) {
                v9 := []int{601, 1}
                if true {
                    if false {
                        p1 = p1
                        var v10 uint8 = uint8(p0)
                        if true {
                            v11 := GenS1{F0: 100}
                            for fi12 := 0; fi12 < 13; fi12++ {
                                switch (fi12 | 32768) {
                                case 1:
                                    fmt.Println(p1, fi12)
                                }
                                var v13 uint64 = 94
                                var v14 float64 = (((((3.14159 / 1.0) + (0.0 / -2.25)) - (((3.14159 * 3.14159) * 3.14159) - -2.25)) + 1e10) * 0.0)
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                            _ = v11
                        }
                        _ = v10
                    } else {
                        fmt.Println(p0)
                        p1 = p1
                        p1 = p1
                    }
                    var v15 uint32 = 919
                    _ = v15
                }
                _ = v9
            }
            if true {
                var v16 uint64 = 2
                var v17 string = p1
                _ = v16
                _ = v17
            }
        }
    case "\n\t":
        fmt.Println(p0, p0)
    case "":
        if false {
            v18 := GenS1{F0: 1}
            v19 := GenS1{F0: 42}
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 7; fi20++ {
                p1 = "\n\t"
                _ = fi20
            }
            v21 := []int{32767, 42, 4294967295, -1}
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            p0 -= (p0 % ((1) | 1))
            for fi23 := 0; fi23 < 18; fi23++ {
                v24 := GenS2{F0: -1, F1: (p0 / ((197) | 1))}
                v25 := []int{(fi23 << ((2) & 63))}
                var v26 uint16 = 32767
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 uint64 = 1
                var v29 int8 = 127
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi22
        }
    }
    p0 = 18446744073709551615
    var v30 uint = (p0 * 0)
    _ = p0
    _ = p1
    _ = v30
    return p1
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 65535
    var v3 bool = (18446744073709551615 >= 1)
    var v4 uint32 = 42
    var v5 int8 = -50
    v6 := []string{"a"}
    if (228 > 255) {
        v31 := GenS3{F0: (v5 & (v5 % (((v5 | (-(v5)))) | 1))), F1: 42, F2: 4294967295}
        fmt.Println(v2)
        _ = v31
    } else {
        v1 += float64(((0.0 / 3.14159) * ((-2.25 / 1e10) - float32(((v1 - (((float64(3.14159) * v1) + 0.0) / 3.14159)) / (float64(((((-2.25 - float32((1.5 * v1))) * 3.14159) / (-2.25 - 3.14159)) - 0.0)) * (((v1 * 3.14159) - v1) - 0.1)))))))
        v4 -= 515
        v32 := GenS4{F0: 42, F1: false}
        _ = v32
    }
    if len(v6) > 0 {
        v6[0] = ("hello" + "foo bar")
    }
    v1 += ((v1 - v1) + ((v1 + (float64(-2.25) - 1e10)) - 1e10))
    if v3 {
        if len(v6) > 0 {
            v6[0] = "hello"
        }
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v2 = (^(v2))
        _ = fi33
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 int8 = (v5 ^ (v5 % ((-102) | 1)))
        _ = fi34
        _ = v35
    }
    if v3 {
        switch "\n\t" {
        case "hello":
            v4 -= (v4 - (^(v4)))
        case "":
            for fi36 := 0; fi36 < 19; fi36++ {
                for fi37 := 0; fi37 < 18; fi37++ {
                    switch "hello" {
                    case "\n\t":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if ((1 == 1) && (-2.25 == 0.1)) {
                            var v38 int32 = 772
                            v2 = v2
                            v39 := GenS1{F0: (v5 % (((v5 & v5)) | 1))}
                            v4 -= v4
                            fmt.Println(v4, v3, fi36)
                            _ = v38
                            _ = v39
                        }
                        if (float64(((0.1 * float32(v1)) / float32(v1))) != (float64(1e10) - v1)) {
                            v40 := GenS5{F0: "foo bar", F1: 0}
                            _ = v40
                        } else {
                            var v41 int = (fi37 % ((-1) | 1))
                            _ = v41
                        }
                    case "foo bar":
                        v42 := GenS6{F0: 1}
                        _ = v42
                    case "hello":
                        v4 -= 65535
                    default:
                        var v43 uint8 = 128
                        _ = v43
                    }
                    _ = len(v6)
                    v1 += v1
                    _ = fi37
                }
                _ = fi36
            }
        case "a":
            var v44 float32 = 3.14159
            _ = v44
        }
    } else {
        v5 -= v5
        v5 = v5
        v4 += v4
    }
    var v45 int8 = v5
    for fi46 := 0; fi46 < 8; fi46++ {
        switch fi46 {
        case 2147483647:
            v47 := GenS2{F0: (v45 % (((v45 << ((255) & 63))) | 1)), F1: 257}
            _ = v47
        default:
            fmt.Println(v5)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi46
    }
    for fi48 := 0; fi48 < 20; fi48++ {
        v49 := GenS7{F0: 0, F1: 32767, F2: 307}
        if len(v6) > 0 {
            v6[0] = ("\n\t" + "\n\t")
        }
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
