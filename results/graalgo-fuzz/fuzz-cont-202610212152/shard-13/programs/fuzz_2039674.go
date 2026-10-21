package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int8; F1 uint; F2 uint64 }
func genFn0() uint16 {
    var v4 uint64 = 1
    fmt.Println(v4)
    v4 = (v4 + (^(v4)))
    if (false && true) {
        if (65535 >= 742) {
            fmt.Println(v4)
        } else {
            var v5 int32 = 2147483647
            v5 = -778
            var v6 string = "hello"
            _ = v5
            _ = v6
        }
    }
    v7 := []int{50, 42, -386}
    _ = v4
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 uint64) int8 {
    v8 := []int{489}
    switch ((((((("\n\t" + "") + "hello") + "hello") + "a") + "") + "a") + ("hello" + (((("" + (("a" + "hello") + ("hello" + ((("hello" + "a") + (("a" + "hello") + (("a" + "foo bar") + ("hello" + (("hello" + "a") + "hello"))))) + "hello")))) + ("a" + ("\n\t" + ("\n\t" + "a")))) + "a") + "foo bar"))) {
    case "hello":
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 string = ""
                var v11 uint16 = 2
                if (false || false) {
                    v12 := GenS1{F0: (^(p1)), F1: 65535, F2: 2147483648}
                    fmt.Println(p1, p1, v11)
                    v8 = append(v8, 127)
                    _ = v12
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            var v13 int16 = p0
            var v14 uint32 = 1
            _ = v13
            _ = v14
        }
    default:
        v15 := GenS2{F0: 56}
        _ = v15
    }
    if len(v8) > 0 {
        v8[0] = -1
    }
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    var v1 int = -766
    v2 := []int{(v1 | 9223372036854775807), v1, -69, (v1 << ((94) & 63))}
    var v3 uint = 42
    if true {
        var v16 int64 = -9223372036854775808
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 float32 = (1e10 + (-2.25 + (3.14159 * (-2.25 / 1.5))))
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println(v1, v3)
    if ((true || (1e10 > ((0.1 + (3.14159 + -2.25)) * -2.25))) || (-13 != 85)) {
        v3 += 674
        _ = len(v2)
    } else {
        v2 = append(v2, v1)
        v20 := GenS3{F0: v1}
        _ = v20
    }
    v3 = (v3 / ((v3) | 1))
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v21 float64 = float64(3.14159)
    var v22 uint16 = 32767
    v23 := GenS4{F0: 65535, F1: true}
    v22 -= v22
    switch ("foo bar" + "foo bar") {
    case "hello":
        if (!(false || false)) {
            var v24 uint64 = 18446744073709551615
            _ = v24
        }
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        v25 := GenS5{F0: genFn1(42, 366), F1: 255, F2: 16}
        _ = v25
    }
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
