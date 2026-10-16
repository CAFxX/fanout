package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 int; F1 int; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0(p0 string, p1 uint64) uint16 {
    var v6 string = p0
    p1 = p1
    if (v6 != (p0 + ("foo bar" + ("foo bar" + (v6 + p0))))) {
        fmt.Println(p0, v6, p1)
        p0 = (p0 + (v6 + "a"))
        v7 := []int{}
        _ = v7
    } else {
        v8 := GenS1{F0: false, F1: p1, F2: 255}
        switch 32768 {
        case 42:
            if (((!true) && (!true)) && ((!(0 >= 119)) || ((p1 - 4294967295) <= p1))) {
                var v9 int64 = -1
                _ = v9
            }
        }
        _ = v8
    }
    v10 := GenS2{F0: 166, F1: 2147483647, F2: 2}
    if false {
        v11 := GenS1{F0: (false || (float32(3.14159) > 3.14159)), F1: p1, F2: 65535}
        v6 = "foo bar"
        if (!true) {
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(p0, p1, p1)
                _ = fi12
            }
            v13 := []int{4294967295, 2147483647, 32768, 923}
            _ = v13
        }
        _ = v11
    } else {
        v6 = (p0 + (v6 + ("hello" + "\n\t")))
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 0
}
func genFn1(p0 string) int8 {
    if ((true && (2 != 2)) || true) {
        var v14 bool = ((true || (4294967295 == 2147483648)) || (0 > -31))
        var v15 uint8 = 2
        _ = v14
        _ = v15
    } else {
        var v16 string = ""
        v17 := GenS2{F0: 127, F1: 42, F2: 255}
        var v18 int = 629
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: 65535}
    switch 32767 {
    case 883:
        if ((-2.25 / 1.0) != 0.1) {
            v20 := []int{}
            var v21 float64 = ((3.14159 * 0.0) - float64((3.14159 * (3.14159 * -2.25))))
            _ = v20
            _ = v21
        } else {
            v22 := []int{}
            if true {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                var v23 float64 = 1e10
                v23 = 1.5
                _ = v23
            } else {
                if true {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        switch ("hello" + p0) {
                        case "hello":
                            fmt.Println(p0, p0)
                        case "a":
                            p0 = ""
                        default:
                            v25 := GenS1{F0: (false && true), F1: 4294967295, F2: 807}
                            _ = v25
                        }
                        _ = fi24
                    }
                    var v26 string = p0
                    _ = v26
                } else {
                    var v27 float64 = float64((1e10 * (1e10 * (-2.25 + -2.25))))
                    p0 = ("a" + (p0 + "\n\t"))
                    p0 = "\n\t"
                    _ = v27
                }
            }
            p0 = p0
            _ = v22
        }
    case 127:
        p0 = (("a" + "") + p0)
    case -9223372036854775808:
        fmt.Println(p0)
    default:
        p0 = p0
    }
    if true {
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := []int{1, 32768, fi28}
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
            _ = len(v29)
            _ = fi28
            _ = v29
        }
        if false {
            v30 := GenS4{F0: 0, F1: 18446744073709551615}
            _ = v30
        }
        p0 = ("hello" + p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = v19
    return -88
}

func main() {
    var v1 int16 = -755
    var v2 int64 = -889
    var v3 bool = (((0.0 > ((1.5 / (1.5 - (1e10 * (0.1 - 3.14159)))) / (((-2.25 + (1e10 / -2.25)) - 0.1) * 1.5))) && true) || true)
    v4 := make(map[int]bool)
    var v5 uint16 = 217
    v31 := GenS1{F0: (!v3), F1: 599, F2: 2}
    v2 = (v2 * v2)
    for fi32 := 0; fi32 < 18; fi32++ {
        v5 = 42
        switch "" {
        case "\n\t":
            var v33 bool = ((42 != 4294967295) || v3)
            _ = v33
        default:
            v34 := GenS1{F0: (1 < 0), F1: 51, F2: 104}
            _ = v34
        }
        _ = fi32
    }
    if v3 {
        var v35 int32 = 1
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 float32 = (0.1 / (1e10 + 1.5))
            var v38 uint64 = 4294967295
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = v35
    }
    fmt.Println(v1, v5)
    v3 = v3
    v2 -= -951
    v4[236] = (!v3)
    v5 += 32767
    for fi39 := 0; fi39 < 9; fi39++ {
        v5 = 146
        v1 += -225
        _ = fi39
    }
    var v40 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
}
