package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int16; F2 bool }
type GenS7 struct { F0 int }
func genFn0(p0 uint) uint64 {
    v4 := GenS1{F0: -9223372036854775808}
    v5 := GenS2{F0: 9, F1: 42}
    _ = p0
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 int64) uint64 {
    p0 = (-(p0))
    var v6 string = ("hello" + "a")
    _ = p0
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 bool = true
    v3 = (42 == 42)
    v3 = (!(((-2.25 * (0.1 / 1.0)) + (1e10 + 1.5)) == (-2.25 - 1.5)))
    switch "\n\t" {
    case "":
        if (!(v3 || (!((1.5 * (float32(float64(((1.5 * 3.14159) * 1.5))) - (1e10 * -2.25))) == 0.0)))) {
            fmt.Println(v3)
        }
    case "foo bar":
        var v7 float64 = 3.14159
        _ = v7
    }
    fmt.Println(v3)
    v3 = false
    var v8 int16 = -32768
    v8 = -32768
    switch "a" {
    case "":
        v9 := GenS1{F0: -678}
        _ = v9
    case "hello":
        v10 := GenS3{F0: false, F1: 42}
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 16; fi11++ {
            if (!(!("hello" <= ((("foo bar" + "") + "a") + (((("foo bar" + "\n\t") + "foo bar") + "foo bar") + ("\n\t" + "\n\t")))))) {
                delete(v2, (fi11 - 0))
            } else {
                v3 = (v3 || v3)
                _ = len(v1)
                v3 = (false && ((!("" != ("" + "hello"))) && (224 == 247)))
            }
            v1[("hello" + "\n\t")] = (488 != 18446744073709551615)
            _ = fi11
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint = 293
            v13 = (v13 | (v13 - v13))
            v8 = -868
            _ = fi12
            _ = v13
        }
    }
    if (866 > 458) {
        if (!true) {
            if v3 {
                v14 := GenS3{F0: (!true), F1: 0}
                _ = v14
            } else {
                if false {
                    v15 := []int{2147483647, 42, -36, -664}
                    _ = v15
                }
                v8 = (v8 % ((v8) | 1))
            }
            var v16 uint64 = 42
            v17 := GenS4{F0: 9223372036854775807, F1: 255}
            _ = v16
            _ = v17
        } else {
            var v18 float32 = 0.0
            v3 = ((153 < 65535) && v3)
            _ = v18
        }
        for fi19 := 0; fi19 < 10; fi19++ {
            if (180 == 4294967295) {
                delete(v2, (fi19 / ((9223372036854775807) | 1)))
                v20 := GenS5{F0: 0}
                _ = v20
            } else {
                v21 := GenS5{F0: int64(fi19)}
                if ((878 != 444) || true) {
                    v3 = v3
                    if (1 >= -128) {
                        v3 = v3
                    }
                    var v22 int8 = int8(v8)
                    _ = v22
                }
                _ = v21
            }
            v23 := GenS6{F0: genFn0(1), F1: v8, F2: (false || v3)}
            _ = fi19
            _ = v23
        }
    } else {
        var v24 uint32 = 1
        _ = v24
    }
    fmt.Println(v3, v3)
    v3 = true
    v25 := GenS7{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v8
}
