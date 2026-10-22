package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0(p0 uint32, p1 bool) int8 {
    v5 := GenS1{F0: p1, F1: true, F2: p1}
    var v6 uint = 4294967295
    v7 := []int{18, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint64) bool {
    p0 = 18446744073709551615
    p0 = (p0 ^ (p0 + p0))
    p0 = 18446744073709551615
    v8 := []int{94, 863, 736}
    var v9 int8 = 42
    _ = p0
    _ = v8
    _ = v9
    return ((0.1 / -2.25) >= (-2.25 + ((1.5 - -2.25) + 1e10)))
}

func main() {
    var v1 uint64 = 229
    var v2 uint64 = v1
    v3 := make(map[string]bool)
    var v4 uint32 = 4294967295
    if true {
        v10 := GenS2{F0: 128, F1: 319, F2: 18446744073709551615}
        _ = v10
    }
    var v11 int64 = -682
    var v12 int16 = 1
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 10; fi14++ {
        v1 -= v2
        v11 -= (v11 >> ((255) & 63))
        if false {
            v15 := []int{(fi14 % ((int(v11)) | 1))}
            for fi16 := 0; fi16 < 18; fi16++ {
                v3["hello"] = false
                v15 = append(v15, (fi16 >> ((0) & 63)))
                var v17 float32 = (float32(0.0) * 0.0)
                v12 = int16(v11)
                _ = fi16
                _ = v17
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = v15
        }
        _ = fi14
    }
    var v18 uint8 = 255
    var v19 uint16 = 3
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: 42, F1: (^(v4)), F2: 141}
        _ = fi20
        _ = v21
    }
    switch 32767 {
    case 32768:
        for fi22 := 0; fi22 < 7; fi22++ {
            v19 = 622
            _ = fi22
        }
    case -847:
        v12 += (-(v12))
    default:
        v3[v13] = (-490 <= 1)
    }
    _ = len(v3)
    switch int(v2) {
    case 32767:
        v3[((v13 + v13) + (v13 + ("" + ("\n\t" + v13))))] = true
    case 4294967295:
        fmt.Println(v12, v4, v2)
    case 9223372036854775807:
        v19 = (v19 & 42)
    }
    var v23 float64 = float64((0.0 - 3.14159))
    switch ("foo bar" + (((v13 + v13) + "a") + v13)) {
    case "foo bar":
        switch "\n\t" {
        case "a":
            if false {
                v3[("hello" + v13)] = (2 != 1)
                var v24 uint = 255
                v23 = v23
                _ = v24
            } else {
                v25 := GenS1{F0: true, F1: (-2147483648 >= 0), F2: false}
                switch -2147483648 {
                case 743:
                    v1 = v2
                default:
                    if (((v12 == (v12 & (v12 >> ((42) & 63)))) && (!(!(-906 < 128)))) && true) {
                        var v26 uint = 2
                        _ = v26
                    }
                }
                _ = v25
            }
        }
    case "hello":
        v27 := GenS1{F0: true, F1: ((v18 & 2) != 196), F2: (2 <= 1)}
        _ = v27
    case "\n\t":
        v18 = v18
    }
    v23 = v23
    if (!(false && true)) {
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 float32 = 0.1
            fmt.Println(v29, v12)
            _ = fi28
            _ = v29
        }
        switch (v13 + ("a" + "\n\t")) {
        case "a":
            v18 += 255
        case "\n\t":
            v1 += 802
        default:
            v2 -= (v1 >> ((65535) & 63))
        }
        fmt.Println(v23, v18)
    }
    v30 := GenS2{F0: (v18 >> ((0) & 63)), F1: 169, F2: 1}
    fmt.Println(v2, v23, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
