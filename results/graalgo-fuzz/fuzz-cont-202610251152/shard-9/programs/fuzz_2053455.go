package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0() int {
    var v5 int64 = -1
    v5 -= v5
    for fi6 := 0; fi6 < 9; fi6++ {
        if (!false) {
            v7 := GenS1{F0: 77, F1: 0}
            if false {
                v5 += 42
            } else {
                switch (("foo bar" + "a") + "hello") {
                case "\n\t":
                    v5 = v5
                case "a":
                    fmt.Println(fi6, fi6)
                case "hello":
                    v5 -= 9223372036854775807
                }
            }
            if true {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint32 = 1
                    v10 := []int{}
                    var v11 int8 = 127
                    v5 = v5
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
            } else {
                fmt.Println(v5, v5)
            }
            _ = v7
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{fi12, fi12, fi12, (fi12 / (((fi12 - 1)) | 1))}
        fmt.Println(v5)
        _ = fi12
        _ = v13
    }
    _ = v5
    return 387
}
func genFn1(p0 uint64, p1 uint16) int8 {
    var v14 uint32 = 2147483648
    v14 = 708
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 int = 32767
    var v3 string = ("a" + "foo bar")
    v4 := []int{532}
    if true {
        if (1.5 == 0.0) {
            var v15 bool = (!(-2.25 < float64(3.14159)))
            var v16 uint8 = 128
            _ = v15
            _ = v16
        }
        v1 += (v1 * -993)
        fmt.Println(v1, v3)
    } else {
        if len(v4) > 0 {
            v4[0] = (v2 + (v2 % (((-(v2))) | 1)))
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint8 = 240
            v18 = v18
            _ = fi17
            _ = v18
        }
    }
    v2 -= (v2 % ((127) | 1))
    _ = len(v4)
    for fi19 := 0; fi19 < 10; fi19++ {
        switch v3 {
        case "foo bar":
            _ = len(v4)
        default:
            if false {
                var v20 int64 = 9223372036854775807
                _ = v20
            }
        }
        if false {
            v4 = append(v4, (fi19 % (((v2 >> ((0) & 63))) | 1)))
            for fi21 := 0; fi21 < 4; fi21++ {
                if len(v4) > 0 {
                    v4[0] = genFn0()
                }
                fmt.Println(v2, v2)
                _ = len(v4)
                _ = fi21
            }
            v1 += 42
        } else {
            fmt.Println(fi19, v3)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        if (!true) {
            v4 = append(v4, v2)
        }
        var v23 int8 = genFn1(42, 627)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 float64 = 3.14159
            var v27 bool = true
            switch "a" {
            case "a":
                v27 = (!(42 >= 476))
            default:
                if v27 {
                    fmt.Println(v27, fi25)
                    switch genFn0() {
                    case -245:
                        var v28 uint64 = 1
                        _ = v28
                    default:
                        if len(v4) > 0 {
                            v4[0] = (v2 / ((-9223372036854775808) | 1))
                        }
                    }
                    var v29 string = ("" + v3)
                    _ = v29
                } else {
                    var v30 uint8 = 127
                    v30 = v30
                    if (((-9223372036854775808 >= 565) && (-2.25 == 1e10)) && (255 > 65535)) {
                        v30 += v30
                    } else {
                        var v31 string = (v3 + "")
                        v32 := GenS1{F0: 0, F1: 42}
                        _ = v31
                        _ = v32
                    }
                    _ = v30
                }
            }
            var v33 string = "\n\t"
            _ = fi25
            _ = v26
            _ = v27
            _ = v33
        }
        v2 = genFn0()
        _ = fi24
    }
    v2 = (-(v2))
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 int64 = 9223372036854775807
        v4 = append(v4, (-(v2)))
        v2 = -2147483648
        _ = fi34
        _ = v35
    }
    v3 = "foo bar"
    v1 -= (^(v1))
    var v36 string = v3
    v1 -= (v1 / ((v1) | 1))
    for fi37 := 0; fi37 < 11; fi37++ {
        if false {
            if (42 >= 42) {
                v38 := GenS2{F0: false, F1: -348, F2: 0}
                fmt.Println(v3, v3, v1)
                if ((v36 <= ("a" + (("hello" + v36) + "a"))) && (!(1 != genFn1(0, 32767)))) {
                    switch v36 {
                    case "":
                        var v39 int64 = 0
                        _ = v39
                    }
                    v3 = v3
                    var v40 int8 = genFn1(623, 65535)
                    _ = v40
                } else {
                    var v41 uint8 = 128
                    v4 = append(v4, fi37)
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    _ = v41
                }
                _ = v38
            }
        }
        _ = fi37
    }
    fmt.Println(v3, v36)
    for fi42 := 0; fi42 < 16; fi42++ {
        v43 := GenS1{F0: 2, F1: 1}
        for fi44 := 0; fi44 < 18; fi44++ {
            v3 = "\n\t"
            if (((!(!(-33 <= genFn1(816, 2)))) && false) && (!(42 < 2147483647))) {
                if ((-1 < -1) || (-128 < genFn1(2, 183))) {
                    var v45 float64 = 0.1
                    v46 := GenS3{F0: (fi44 | (fi42 + v2))}
                    var v47 int32 = -2147483648
                    _ = v45
                    _ = v46
                    _ = v47
                }
                var v48 uint64 = 0
                _ = v48
            } else {
                v49 := GenS1{F0: 2, F1: 953}
                _ = v49
            }
            _ = fi44
        }
        for fi50 := 0; fi50 < 8; fi50++ {
            v51 := GenS4{F0: 4294967295, F1: 128, F2: false}
            switch v2 {
            case -2147483648:
                var v52 int = (v2 >> ((18446744073709551615) & 63))
                var v53 uint32 = 270
                _ = v52
                _ = v53
            }
            _ = fi50
            _ = v51
        }
        _ = fi42
        _ = v43
    }
    if false {
        var v54 int32 = 1
        if len(v4) > 0 {
            v4[0] = (v2 >> ((625) & 63))
        }
        var v55 int32 = v54
        _ = v54
        _ = v55
    }
    fmt.Println(v36, v3)
    v3 = (("\n\t" + v3) + v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
