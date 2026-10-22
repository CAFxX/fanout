package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 bool, p1 uint) bool {
    var v7 float32 = (-2.25 / ((0.1 + ((3.14159 * (3.14159 / 1.0)) / ((-2.25 / 1.0) + 3.14159))) - (-2.25 / 1.5)))
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int8 = 117
        _ = fi8
        _ = v9
    }
    if (false && ((962 < 719) && false)) {
        v7 -= v7
    } else {
        var v10 int64 = -284
        _ = v10
    }
    var v11 uint8 = 22
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return p0
}

func main() {
    var v1 int = 32768
    v2 := make(map[uint32]int)
    v3 := []int{v1, v1, (v1 + (v1 & v1)), v1}
    v4 := make(map[uint32]string)
    var v5 int16 = 32767
    v6 := make(map[uint32]bool)
    if false {
        v5 += 32767
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    v13 := GenS1{F0: "a", F1: false, F2: v1}
    v14 := GenS1{F0: ("foo bar" + ""), F1: (v1 >= 32768), F2: v1}
    switch v1 {
    case 127:
        fmt.Println(v5, v5)
    case 30:
        switch "hello" {
        case "a":
            var v15 uint32 = 2147483648
            _ = v15
        case "hello":
            for fi16 := 0; fi16 < 7; fi16++ {
                v5 += (v5 / (((v5 - v5)) | 1))
                switch v1 {
                case -2147483648:
                    v1 += 545
                default:
                    if true {
                        v5 = 32767
                        delete(v2, 4294967295)
                        fmt.Println(fi16)
                    }
                }
                _ = len(v3)
                _ = fi16
            }
        default:
            v2[65535] = 127
        }
    case 32767:
        v1 = int(v5)
    }
    if (!true) {
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                v1 = (v1 + (fi17 & (v1 * (-(fi17)))))
                _ = len(v3)
            } else {
                v18 := GenS1{F0: ("\n\t" + "foo bar"), F1: ((1.5 * 0.1) == (1.5 + 0.1)), F2: 685}
                var v19 float64 = 1e10
                _ = v18
                _ = v19
            }
            v1 -= 32767
            _ = fi17
        }
        if (!((-2147483648 <= int32(v1)) && (127 <= 0))) {
            v20 := GenS1{F0: "hello", F1: (!true), F2: 32768}
            _ = v20
        } else {
            v1 += v1
            fmt.Println(v1)
        }
        v4[308] = "a"
    } else {
        var v21 string = ("" + "\n\t")
        if ((true || false) && ((2147483648 >= 2147483647) || (127 != -45))) {
            v22 := GenS2{F0: 2}
            v21 = ((v21 + v21) + "")
            if len(v3) > 0 {
                v3[0] = 0
            }
            _ = v22
        }
        fmt.Println(v1, v1)
        _ = v21
    }
    fmt.Println(v5)
    v23 := GenS1{F0: ("\n\t" + (("a" + (("foo bar" + "hello") + ("a" + "hello"))) + "foo bar")), F1: false, F2: v1}
    if (!true) {
        if false {
            var v24 int = v1
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 uint32 = 4294967295
                if ((false && true) && ((v5 * v5) != 783)) {
                    v27 := GenS1{F0: ("" + ("\n\t" + ((("hello" + "a") + "") + ("hello" + "")))), F1: (!true), F2: 4294967295}
                    _ = v27
                } else {
                    fmt.Println(v26)
                }
                _ = fi25
                _ = v26
            }
            if len(v3) > 0 {
                v3[0] = (v1 / ((v1) | 1))
            }
            _ = v24
        } else {
            var v28 uint = 4294967295
            v28 = v28
            if true {
                v28 = v28
                switch v1 {
                case -829:
                    var v29 uint8 = 0
                    _ = v29
                case 196:
                    fmt.Println(v28, v1, v5)
                default:
                    var v30 bool = true
                    _ = v30
                }
                var v31 float32 = 0.1
                _ = v31
            } else {
                v5 -= v5
                v32 := GenS3{F0: -409, F1: 380}
                fmt.Println(v28, v5, v5)
                _ = v32
            }
            _ = v28
        }
    } else {
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS3{F0: -370, F1: 0}
            _ = fi33
            _ = v34
        }
    }
    var v35 uint32 = 100
    var v36 uint16 = 806
    v37 := GenS4{F0: 81, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
