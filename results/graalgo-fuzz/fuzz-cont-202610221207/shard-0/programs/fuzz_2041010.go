package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }
func genFn0() string {
    var v4 bool = false
    switch "\n\t" {
    case "\n\t":
        if v4 {
            var v5 string = "a"
            if v4 {
                if (-128 <= -16) {
                    v6 := GenS1{F0: ((v5 + v5) + v5), F1: 118, F2: -9223372036854775808}
                    _ = v6
                } else {
                    v5 = "foo bar"
                }
            }
            if false {
                v7 := GenS1{F0: ("a" + (("" + v5) + v5)), F1: 0, F2: 827}
                _ = v7
            } else {
                v8 := GenS1{F0: ((v5 + ("hello" + v5)) + v5), F1: -91, F2: -1}
                _ = v8
            }
            _ = v5
        }
    case "hello":
        v9 := GenS2{F0: 416, F1: 56}
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 498
        v12 := GenS3{F0: fi10, F1: 4294967295}
        if (1e10 >= 3.14159) {
            fmt.Println(v11)
        } else {
            v13 := GenS4{F0: fi10, F1: v11, F2: 425}
            _ = v13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v4
    return "hello"
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{"foo bar"}
    var v3 int16 = 0
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = -1
        _ = fi14
        _ = v15
    }
    var v16 int32 = -2147483648
    for fi17 := 0; fi17 < 1; fi17++ {
        v16 -= (-(v16))
        var v18 string = "foo bar"
        if false {
            _ = len(v2)
        } else {
            var v19 float64 = float64(0.0)
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    var v20 uint = 42
    v1 = 2
    v16 = v16
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint = (v20 - uint(v1))
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint64 = 42
        v20 = (v20 - v20)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v20, v16, fi25)
        fmt.Println(v16, v16)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
}
