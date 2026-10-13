package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string; F1 string; F2 int32 }
func genFn0(p0 uint8) int8 {
    v5 := GenS1{F0: ("hello" > ((("hello" + ("\n\t" + ("foo bar" + ("a" + "")))) + "foo bar") + "hello")), F1: 975}
    var v6 uint = 65535
    if (v6 != (v6 << ((v6) & 63))) {
        var v7 uint8 = (^(p0))
        _ = v7
    } else {
        v6 += (v6 ^ (v6 + 255))
    }
    switch -1 {
    case 127:
        fmt.Println(v6, v6)
    case -1:
        v8 := GenS2{F0: false, F1: 42, F2: 0}
        _ = v8
    case 78:
        switch -516 {
        case 42:
            var v9 int16 = -456
            _ = v9
        case 32768:
            var v10 string = "foo bar"
            _ = v10
        case 32767:
            switch "\n\t" {
            case "a":
                for fi11 := 0; fi11 < 16; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        fmt.Println(fi12, fi12)
                        v13 := GenS1{F0: (!true), F1: 2147483648}
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            default:
                for fi14 := 0; fi14 < 3; fi14++ {
                    if (p0 < 255) {
                        v15 := []int{0, -652, (fi14 - fi14)}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        v16 := GenS2{F0: (true || (false && true)), F1: 361, F2: 738}
                        _ = v15
                        _ = v16
                    }
                    p0 = 2
                    var v17 int8 = 42
                    _ = fi14
                    _ = v17
                }
            }
        default:
            v18 := GenS1{F0: false, F1: 4294967295}
            _ = v18
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := []uint32{2}
    v2 := []int{4294967295}
    var v3 string = ("hello" + "")
    var v4 uint = 2
    for fi19 := 0; fi19 < 13; fi19++ {
        switch v3 {
        case "\n\t":
            var v20 string = v3
            _ = v20
        case "hello":
            var v21 bool = (1e10 > 1e10)
            _ = v21
        default:
            var v22 uint32 = 1
            _ = v22
        }
        v3 = v3
        _ = fi19
    }
    fmt.Println(v3, v3)
    var v23 string = v3
    _ = len(v2)
    var v24 int8 = -23
    if (((1e10 / 1e10) / -2.25) <= 3.14159) {
        v3 = (((("hello" + "") + "\n\t") + v3) + v3)
        v3 = "\n\t"
        if true {
            if (true && (2147483648 >= 2147483648)) {
                v3 = "a"
                switch ((("" + (("foo bar" + v23) + ("\n\t" + "a"))) + "foo bar") + v23) {
                case "":
                    var v25 float32 = (0.0 * float32(-2.25))
                    _ = v25
                }
            }
        } else {
            if (949 > v4) {
                v26 := GenS2{F0: false, F1: 255, F2: 789}
                if (!(18446744073709551615 <= 1)) {
                    v27 := GenS1{F0: ((v24 - (v24 % ((-18) | 1))) < (v24 % ((v24) | 1))), F1: 4294967295}
                    var v28 uint = 1
                    _ = v27
                    _ = v28
                }
                _ = len(v2)
                _ = v26
            } else {
                v2 = append(v2, 32767)
                v23 = (("a" + "hello") + v23)
            }
            v29 := GenS3{F0: "hello", F1: v3, F2: 57}
            _ = v29
        }
    }
    switch -936 {
    case -9223372036854775808:
        v4 -= (v4 + 65535)
    case 648:
        var v30 uint64 = 1
        _ = v30
    default:
        if false {
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 string = "foo bar"
                v33 := GenS3{F0: "hello", F1: (v23 + "a"), F2: -2147483648}
                switch fi31 {
                case 127:
                    v34 := GenS3{F0: ("a" + ("" + v32)), F1: v32, F2: 213}
                    _ = v34
                case 421:
                    v23 = v32
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
    }
    v35 := GenS2{F0: true, F1: 32767, F2: 2147483647}
    if (-193 == -806) {
        switch -1 {
        case 403:
            v4 = v4
        default:
            var v36 int32 = 0
            _ = v36
        }
        fmt.Println(v3, v3)
        var v37 uint16 = 864
        _ = v37
    } else {
        _ = len(v2)
    }
    v3 = v23
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 float64 = -2.25
        _ = fi38
        _ = v39
    }
    v3 = v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
}
