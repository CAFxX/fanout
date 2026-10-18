package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int16; F2 int16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 string }
func genFn0(p0 int) int {
    p0 = -1
    v4 := []int{-668, p0, (-(p0)), 4294967295}
    p0 = (p0 * p0)
    p0 = -277
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 bool, p1 int32) uint8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{(fi5 >> ((18446744073709551615) & 63)), genFn0(210), -161}
        _ = fi5
        _ = v6
    }
    var v7 bool = ((18446744073709551615 > 783) || (!false))
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := []string{(("foo bar" + "") + "foo bar")}
    var v2 int8 = 0
    var v3 uint8 = 27
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 42, F1: 1}
        fmt.Println(fi8, v2, v3)
        _ = fi8
        _ = v9
    }
    v3 = 1
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(v3, fi11)
            if false {
                v2 += (v2 - 127)
                _ = len(v1)
            } else {
                v3 = 146
                switch 9223372036854775807 {
                case 1:
                    var v12 int64 = -9223372036854775808
                    _ = v12
                case -1:
                    fmt.Println(v2, v3)
                }
                v3 = genFn1((!true), -1)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi11
        }
        var v13 string = "hello"
        _ = fi10
        _ = v13
    }
    var v14 int = -506
    _ = len(v1)
    for fi15 := 0; fi15 < 5; fi15++ {
        v14 += v14
        var v16 string = "\n\t"
        for fi17 := 0; fi17 < 1; fi17++ {
            if true {
                if (78 <= v3) {
                    v3 -= (v3 / (((v3 % (((v3 ^ genFn1(false, -2147483648))) | 1))) | 1))
                    var v18 string = v16
                    fmt.Println(v18, v16)
                    _ = v18
                }
            } else {
                if false {
                    switch (fi17 << ((4294967295) & 63)) {
                    case 247:
                        var v19 uint16 = 65535
                        _ = v19
                    default:
                        v20 := GenS2{F0: v16, F1: 32767, F2: genFn0(1)}
                        _ = v20
                    }
                    var v21 bool = (-116 >= -917)
                    if len(v1) > 0 {
                        v1[0] = (v16 + v16)
                    }
                    _ = v21
                }
                switch v16 {
                case "a":
                    if true {
                        fmt.Println(v2)
                        v22 := GenS2{F0: "\n\t", F1: 32767, F2: genFn0(fi17)}
                        _ = v22
                    }
                case "hello":
                    v23 := GenS3{F0: 1, F1: 64}
                    _ = v23
                case "foo bar":
                    v24 := GenS4{F0: 65535}
                    _ = v24
                default:
                    if false {
                        if ((true && (!true)) && (2 == 1)) {
                            v3 = v3
                        }
                        if (0 != -480) {
                            fmt.Println(v16, v16, fi17)
                            v25 := GenS5{F0: v2, F1: 210, F2: -65}
                            var v26 float32 = (((float32(float64(1.5)) + 3.14159) - 1e10) * 3.14159)
                            _ = v25
                            _ = v26
                        }
                    }
                }
                var v27 int16 = 32767
                _ = v27
            }
            v28 := GenS3{F0: 4294967295, F1: 1}
            v29 := GenS6{F0: 4294967295, F1: 0, F2: "hello"}
            _ = fi17
            _ = v28
            _ = v29
        }
        _ = fi15
        _ = v16
    }
    var v30 uint64 = 4294967295
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int32 = 0
        _ = fi31
        _ = v32
    }
    v30 = v30
    v30 -= v30
    if true {
        fmt.Println(v3, v14, v3)
        var v33 uint16 = 0
        switch v14 {
        case 0:
            v30 -= v30
        default:
            for fi34 := 0; fi34 < 4; fi34++ {
                v33 += 32767
                for fi35 := 0; fi35 < 11; fi35++ {
                    switch ("hello" + (("hello" + ("hello" + "foo bar")) + ("" + ("foo bar" + "\n\t")))) {
                    case "hello":
                        var v36 float64 = -2.25
                        _ = v36
                    default:
                        var v37 int64 = 42
                        _ = v37
                    }
                    var v38 int64 = -385
                    _ = fi35
                    _ = v38
                }
                _ = fi34
            }
        }
        _ = v33
    }
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v39 uint16 = 108
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v39
}
