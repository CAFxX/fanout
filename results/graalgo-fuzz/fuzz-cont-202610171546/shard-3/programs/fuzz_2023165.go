package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 int, p1 uint64, p2 uint64) int16 {
    fmt.Println(p1, p1, p1)
    if true {
        v5 := []int{}
        _ = v5
    }
    var v6 uint8 = 255
    var v7 uint = 1
    v7 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 622
}
func genFn1(p0 int) uint32 {
    if false {
        var v8 uint = 2
        p0 = p0
        _ = v8
    } else {
        if ((783 >= 342) && false) {
            switch (p0 | p0) {
            case 667:
                v9 := []int{(-(p0)), p0, p0}
                _ = v9
            default:
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS1{F0: -9223372036854775808, F1: 0}
                    _ = fi10
                    _ = v11
                }
            }
            var v12 int8 = -108
            _ = v12
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                p0 -= -9223372036854775808
                fmt.Println(p0, fi13)
                _ = fi13
            }
            fmt.Println(p0)
        }
        if (((0.0 - (-2.25 - (1.5 / 1.0))) == -2.25) || (false || true)) {
            var v14 uint = 0
            v15 := []int{}
            if len(v15) > 0 {
                v15[0] = (p0 >> ((v14) & 63))
            }
            _ = v14
            _ = v15
        }
    }
    v16 := GenS1{F0: -9223372036854775808, F1: 100}
    if true {
        var v17 float32 = 1.5
        v18 := []int{p0, (^(p0)), p0, (-(p0))}
        switch "" {
        case "foo bar":
            var v19 uint16 = 704
            _ = v19
        case "":
            v20 := GenS1{F0: -1, F1: 127}
            _ = v20
        case "hello":
            v21 := GenS1{F0: -1, F1: -121}
            _ = v21
        }
        _ = v17
        _ = v18
    } else {
        p0 -= p0
        v22 := GenS1{F0: 42, F1: 100}
        _ = v22
    }
    var v23 uint = 18446744073709551615
    var v24 int16 = 32767
    _ = p0
    _ = v16
    _ = v23
    _ = v24
    return 2147483648
}

func main() {
    v1 := []int{127, 1}
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 uint64 = 537
    var v25 float64 = 3.14159
    v26 := GenS2{F0: 666, F1: 903, F2: 4294967295}
    _ = len(v2)
    var v27 int = 127
    v28 := GenS1{F0: -793, F1: 1}
    v25 -= (-2.25 + v25)
    v29 := GenS3{F0: 0, F1: ("a" + "foo bar")}
    v25 += v25
    v30 := GenS1{F0: -9223372036854775808, F1: -128}
    switch ("foo bar" + ("a" + ("\n\t" + (("a" + ((("" + "\n\t") + "\n\t") + "")) + "\n\t")))) {
    case "foo bar":
        _ = len(v3)
    case "a":
        for fi31 := 0; fi31 < 9; fi31++ {
            v4 = v4
            delete(v3, ("" + ((("foo bar" + "") + (("\n\t" + "\n\t") + "a")) + "foo bar")))
            v3["hello"] = 4294967295
            _ = fi31
        }
    case "":
        v4 -= (v4 ^ 4294967295)
    }
    if (175 < 128) {
        for fi32 := 0; fi32 < 20; fi32++ {
            v2[("a" + "foo bar")] = v27
            var v33 float32 = 3.14159
            _ = fi32
            _ = v33
        }
        delete(v3, "hello")
        for fi34 := 0; fi34 < 17; fi34++ {
            v25 = (((1e10 * ((-2.25 - (v25 + -2.25)) * v25)) + (((((v25 - (float64((1e10 - 3.14159)) + v25)) + v25) / v25) + (v25 - ((float64(((1.5 / 1.0) / 1e10)) * v25) - v25))) / v25)) * ((3.14159 * v25) - v25))
            fmt.Println(v25)
            if (true || (!false)) {
                var v35 float32 = 0.0
                var v36 string = ("a" + ("foo bar" + "a"))
                v4 -= (v4 ^ v4)
                _ = v35
                _ = v36
            }
            _ = fi34
        }
    }
    v4 = v4
    var v37 uint32 = 2
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v4
}
