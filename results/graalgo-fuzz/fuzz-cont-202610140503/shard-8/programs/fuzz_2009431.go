package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint64, p1 uint16, p2 int64) uint8 {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int = 4294967295
        p1 = (p1 - (p1 ^ (p1 << ((810) & 63))))
        var v8 uint16 = 770
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "":
        if (!(false || true)) {
            v9 := GenS1{F0: (p1 | (p1 ^ p1)), F1: ""}
            p2 += (p2 ^ p2)
            _ = v9
        }
    case "a":
        var v10 bool = true
        _ = v10
    case "foo bar":
        v11 := []int{0, -303, -1, 0}
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint32 = 2147483648
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(p2, p1, p2)
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    if false {
        var v15 uint16 = p1
        v16 := GenS1{F0: (^(v15)), F1: ""}
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 uint32, p1 uint16) uint16 {
    p0 = p0
    var v17 int = 127
    fmt.Println(v17, p1)
    if (((float64(3.14159) + (0.0 / 1e10)) != float64((((0.1 / (1.5 + -2.25)) + -2.25) * (-2.25 / 1.0)))) || false) {
        if (0.0 <= 1e10) {
            v18 := []int{v17}
            _ = v18
        } else {
            var v19 float64 = 3.14159
            _ = v19
        }
        fmt.Println(v17, p0, p0)
        var v20 uint8 = 42
        _ = v20
    } else {
        v21 := []int{v17}
        if (-2147483648 > 345) {
            v22 := GenS1{F0: p1, F1: "\n\t"}
            v17 = 32768
            _ = v22
        }
        _ = v21
    }
    if (true || (867 > 42)) {
        fmt.Println(p0, p1)
        if false {
            var v23 int16 = 101
            var v24 float64 = 3.14159
            if false {
                p1 += (p1 << ((523) & 63))
            }
            _ = v23
            _ = v24
        } else {
            var v25 int32 = -977
            _ = v25
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            for fi27 := 0; fi27 < 6; fi27++ {
                switch "foo bar" {
                case "hello":
                    v28 := []int{(fi26 & 127)}
                    _ = v28
                case "foo bar":
                    fmt.Println(fi27, v17)
                    var v29 int32 = 0
                    _ = v29
                case "":
                    p0 = (p0 - (p0 + uint32(p1)))
                }
                v30 := GenS1{F0: p1, F1: (("foo bar" + "hello") + "foo bar")}
                v31 := []int{(fi26 + fi27), (v17 ^ 0)}
                _ = fi27
                _ = v30
                _ = v31
            }
            p1 -= p1
            _ = fi26
        }
    } else {
        switch v17 {
        case 1:
            fmt.Println(v17, p0)
        case -346:
            fmt.Println(p0, v17)
        }
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = v17
    return p1
}

func main() {
    var v1 int = 134
    v2 := []int{v1, 9223372036854775807, 32767}
    v3 := make(map[uint32]string)
    v4 := []int{v1, (v1 / ((v1) | 1)), v1, (^(v1))}
    var v5 uint = 168
    fmt.Println(v1, v1)
    v1 -= (v1 + (v1 + v1))
    if true {
        _ = len(v2)
        _ = len(v3)
        for fi32 := 0; fi32 < 4; fi32++ {
            fmt.Println(v1, fi32, v5)
            _ = fi32
        }
    }
    if (!((-2.25 * 3.14159) != 3.14159)) {
        var v33 uint = (v5 % (((^(v5))) | 1))
        _ = v33
    }
    switch "hello" {
    case "hello":
        v34 := GenS1{F0: genFn1(2147483647, 32767), F1: "hello"}
        _ = v34
    default:
        fmt.Println(v5, v5)
    }
    v1 -= (v1 / ((v1) | 1))
    v1 -= -563
    v35 := GenS1{F0: genFn1(1, genFn1(618, genFn1(2147483647, 1))), F1: "a"}
    var v36 uint16 = 32767
    v5 = (v5 * 4294967295)
    var v37 int64 = 220
    v3[42] = ((("" + "a") + "a") + "hello")
    if false {
        v37 = (v37 / (((v37 ^ v37)) | 1))
    }
    v5 -= (v5 % ((v5) | 1))
    var v38 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
