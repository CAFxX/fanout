package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 int8 }
func genFn0() int {
    var v4 uint64 = 42
    v4 = 18446744073709551615
    fmt.Println(v4)
    _ = v4
    return -1
}
func genFn1(p0 uint16, p1 int64) uint32 {
    var v5 uint64 = 4294967295
    var v6 float32 = (1.5 - ((-2.25 * 3.14159) + (0.1 / -2.25)))
    p1 -= (p1 | (p1 - p1))
    switch ((("hello" + "\n\t") + ("foo bar" + ("hello" + ""))) + "hello") {
    case "foo bar":
        v7 := GenS1{F0: (v5 | v5)}
        _ = v7
    }
    switch ("foo bar" + (("" + "") + "")) {
    case "":
        var v8 int16 = 32767
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint8 = 42
    v3 := []uint32{}
    var v9 int16 = 1
    v2 += v2
    v9 -= (v9 | (v9 / (((v9 * v9)) | 1)))
    v1 = append(v1, ("\n\t" + "a"))
    if true {
        for fi10 := 0; fi10 < 6; fi10++ {
            v9 -= v9
            for fi11 := 0; fi11 < 9; fi11++ {
                if (false && true) {
                    v9 = 42
                }
                var v12 uint = 1
                fmt.Println(fi10, v12)
                v12 = 255
                _ = fi11
                _ = v12
            }
            var v13 int = fi10
            _ = fi10
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float32 = 3.14159
            v9 -= (v9 / ((v9) | 1))
            v15 = v15
            _ = fi14
            _ = v15
        }
    } else {
        switch ("hello" + "") {
        case "a":
            var v16 uint32 = 65535
            _ = v16
        case "hello":
            switch 9223372036854775807 {
            case -282:
                var v17 int = 380
                _ = v17
            }
        default:
            if len(v1) > 0 {
                v1[0] = ("a" + "")
            }
        }
        var v18 string = "\n\t"
        v1 = append(v1, (v18 + v18))
        _ = v18
    }
    if false {
        v19 := GenS2{F0: 127, F1: 42, F2: 1}
        _ = len(v1)
        var v20 float64 = -2.25
        _ = v19
        _ = v20
    }
    fmt.Println(v9, v2)
    var v21 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
