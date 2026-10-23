package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint64) int32 {
    if (4294967295 == 266) {
        if (-2.25 == 0.0) {
            var v7 int8 = 127
            v8 := GenS1{F0: (((0.1 * -2.25) / 1e10) <= 1e10)}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            if false {
                if ((1.5 != (-2.25 + 1e10)) && (0.0 <= 3.14159)) {
                    p1 += (p1 % ((p1) | 1))
                }
            } else {
                fmt.Println(p1)
            }
            _ = fi9
        }
        if ((p1 << ((18446744073709551615) & 63)) <= p1) {
            p1 = p1
            var v10 int64 = -1
            fmt.Println(p0)
            _ = v10
        } else {
            if (2147483647 <= -2147483648) {
                if (!((2147483648 != 65535) && (962 == 65535))) {
                    if false {
                        var v11 uint8 = 0
                        var v12 int8 = 1
                        _ = v11
                        _ = v12
                    } else {
                        var v13 float32 = 0.1
                        v14 := []int{-90}
                        _ = v13
                        _ = v14
                    }
                    var v15 uint8 = 42
                    var v16 uint64 = (p1 & (p1 + p1))
                    _ = v15
                    _ = v16
                }
            }
        }
    } else {
        if (50 >= p1) {
            for fi17 := 0; fi17 < 12; fi17++ {
                p0 = (255 != 0)
                var v18 int32 = 2147483647
                _ = fi17
                _ = v18
            }
            var v19 uint64 = (p1 ^ (p1 % ((p1) | 1)))
            _ = v19
        } else {
            v20 := GenS2{F0: false, F1: 1, F2: (p1 >> ((18446744073709551615) & 63))}
            p0 = p0
            _ = v20
        }
        p1 = (^(p1))
    }
    if (2 == 2147483647) {
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS3{F0: -258, F1: 1000}
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    } else {
        var v24 uint = 0
        var v25 int64 = -1
        _ = v24
        _ = v25
    }
    p0 = (!(0.1 < ((-2.25 / 1.0) - (0.1 * (-2.25 - 0.1)))))
    fmt.Println(p1)
    _ = p0
    _ = p1
    return int32(p1)
}
func genFn1(p0 int64, p1 uint64, p2 int16) string {
    if (false || (p0 != -9223372036854775808)) {
        var v26 uint16 = 0
        v27 := GenS2{F0: (!true), F1: 1, F2: p1}
        _ = v26
        _ = v27
    }
    fmt.Println(p1, p2, p1)
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint8 = 13
        v30 := []int{fi28, (fi28 | fi28)}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        p1 += p1
        var v32 float32 = 3.14159
        switch fi31 {
        case 9223372036854775807:
            var v33 uint32 = 2147483648
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(p0, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return (((("hello" + "") + "foo bar") + "a") + "hello")
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 130
    v3 := make(map[string]string)
    var v4 float64 = float64(1e10)
    v5 := []string{"\n\t", "\n\t", "", ("a" + "a")}
    v6 := []string{"a", ("\n\t" + ("" + "foo bar"))}
    var v34 float32 = 0.0
    if false {
        v35 := GenS4{F0: 703}
        _ = v35
    }
    v34 = (v34 / (v34 - v34))
    switch "foo bar" {
    case "hello":
        v36 := GenS3{F0: 538, F1: -32768}
        _ = v36
    case "":
        if false {
            v34 = ((v34 + 3.14159) * 0.0)
        }
    case "a":
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS3{F0: 1000, F1: 1}
            _ = fi37
            _ = v38
        }
    }
    switch 127 {
    case 957:
        v4 = ((1e10 + float64((v34 / float32(((v4 * ((((v4 - (1.5 + 1e10)) * (float64((0.1 - (v34 + 3.14159))) - (v4 - -2.25))) * v4) * v4)) + 1e10))))) * ((0.0 + v4) - 0.1))
    case 0:
        _ = len(v1)
    }
    _ = len(v1)
    for fi39 := 0; fi39 < 9; fi39++ {
        _ = len(v3)
        _ = fi39
    }
    var v40 float32 = ((float32(v4) - v34) + 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
