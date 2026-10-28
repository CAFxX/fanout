package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 string; F2 string }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 int; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 string }
type GenS7 struct { F0 bool }
func genFn0(p0 bool, p1 int16) string {
    v5 := GenS1{F0: 42, F1: 119, F2: 42}
    fmt.Println(p1, p1)
    fmt.Println(p1, p1)
    switch 1 {
    case -481:
        var v6 uint = 394
        _ = v6
    default:
        if p0 {
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := []int{2147483647, (fi7 + fi7)}
                var v9 int64 = -9223372036854775808
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 string = "foo bar"
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return ("hello" + ("" + ((("" + "a") + ("hello" + "foo bar")) + "")))
}
func genFn1(p0 int64, p1 string, p2 int16) uint {
    var v11 float32 = (3.14159 - 1e10)
    if true {
        v11 = v11
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{}
            var v14 bool = true
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        if false {
            var v15 uint32 = 1
            _ = v15
        }
    }
    p2 = p2
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1e10
    var v3 int32 = 42
    v4 := []int{1, -1}
    var v16 int = 215
    var v17 uint8 = 255
    if len(v4) > 0 {
        v4[0] = 32767
    }
    fmt.Println(v16, v17)
    for fi18 := 0; fi18 < 19; fi18++ {
        if true {
            delete(v1, "a")
            if true {
                switch "foo bar" {
                case "a":
                    var v19 uint = 1
                    _ = v19
                case "\n\t":
                    var v20 uint64 = 896
                    _ = v20
                case "":
                    v16 -= 2147483647
                }
                v1[((("hello" + ("\n\t" + "foo bar")) + "\n\t") + (("hello" + ("foo bar" + "a")) + "\n\t"))] = 65535
            } else {
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 float64 = v2
                    var v23 uint16 = 0
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            var v24 uint8 = v17
            _ = v24
        }
        var v25 bool = true
        _ = fi18
        _ = v25
    }
    var v26 int16 = 32767
    var v27 float32 = (-2.25 + (0.0 + (-2.25 / (3.14159 - 0.0))))
    if (v17 >= (v17 & v17)) {
        v16 -= (v16 ^ 32768)
        if ((!(42 >= 255)) && (((true && false) && false) && true)) {
            for fi28 := 0; fi28 < 10; fi28++ {
                for fi29 := 0; fi29 < 8; fi29++ {
                    switch "\n\t" {
                    case "":
                        v26 = v26
                        v26 -= v26
                        v26 = (v26 & (v26 - v26))
                    }
                    v17 = v17
                    _ = fi29
                }
                _ = fi28
            }
            if false {
                if true {
                    if true {
                        if ((v2 * v2) != (((v2 * (v2 * (float64(v27) / (v2 - float64(0.0))))) + (v2 * float64((v27 + float32(-2.25))))) / v2)) {
                            fmt.Println(v27)
                            var v30 int32 = -1
                            fmt.Println(v3)
                            _ = v30
                        }
                    } else {
                        fmt.Println(v2)
                    }
                }
            }
            v31 := GenS2{F0: 2147483648, F1: "\n\t", F2: "a"}
            _ = v31
        }
    } else {
        var v32 int = (v16 % (((v16 >> ((401) & 63))) | 1))
        for fi33 := 0; fi33 < 16; fi33++ {
            v34 := GenS2{F0: 462, F1: "hello", F2: ("hello" + "hello")}
            v35 := v1[""]
            _ = fi33
            _ = v34
            _ = v35
        }
        v17 -= (v17 >> ((78) & 63))
        _ = v32
    }
    if false {
        fmt.Println(v2, v16, v2)
        for fi36 := 0; fi36 < 3; fi36++ {
            v17 = 1
            _ = fi36
        }
        v37 := v1[("\n\t" + "hello")]
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            _ = len(v1)
            var v40 int = (v16 << ((genFn1(1, "hello", -32768)) & 63))
            _ = fi39
            _ = v40
        }
        if (false || true) {
            if len(v4) > 0 {
                v4[0] = v16
            }
            v26 -= v26
            v26 = (v26 / ((v26) | 1))
        }
        if (v2 <= v2) {
            var v41 uint8 = v17
            if true {
                var v42 int32 = v3
                fmt.Println(fi38)
                _ = v42
            }
            _ = v41
        } else {
            v43 := v1["foo bar"]
            v43 += 42
            _ = v43
        }
        _ = fi38
    }
    v44 := GenS3{F0: 1, F1: -1}
    for fi45 := 0; fi45 < 5; fi45++ {
        if ((((v27 - v27) - (((((v27 * (v27 - (((v27 - 0.0) * 3.14159) + 0.1))) - (float32(1e10) * float32(v2))) / 3.14159) + v27) / 1.5)) * (0.0 * (((float32((v2 + 0.1)) / v27) + (float32((v2 * v2)) + v27)) - v27))) <= 1e10) {
            v46 := GenS4{F0: 4294967295, F1: (v17 + 110)}
            _ = v46
        }
        _ = fi45
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        if (!true) {
            var v48 float64 = v2
            v27 -= v27
            _ = v48
        }
        var v49 uint16 = 32767
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 bool = true
            _ = fi50
            _ = v51
        }
        _ = fi47
        _ = v49
    }
    fmt.Println(v3, v27)
    if (204 == -310) {
        v52 := GenS5{F0: v16, F1: 4294967295, F2: (v17 << ((0) & 63))}
        v53 := GenS6{F0: 127, F1: 815, F2: ("a" + ("a" + "a"))}
        _ = v52
        _ = v53
    }
    var v54 float64 = (((v2 + v2) - (v2 / v2)) + (1.5 * v2))
    fmt.Println(v17)
    v55 := GenS7{F0: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v54
    _ = v55
}
