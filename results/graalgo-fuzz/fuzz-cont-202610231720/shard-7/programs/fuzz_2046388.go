package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 int32 }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 string) uint16 {
    v7 := []int{0}
    v8 := GenS1{F0: 42, F1: 87, F2: 0}
    switch -893 {
    case 32768:
        v9 := GenS2{F0: true}
        _ = v9
    case 1:
        p0 = (p0 + ("a" + p0))
    case -646:
        var v10 string = p0
        _ = v10
    default:
        v11 := GenS2{F0: (!true)}
        _ = v11
    }
    v7 = append(v7, 0)
    switch p0 {
    case "hello":
        v12 := GenS1{F0: 127, F1: 2, F2: 42}
        _ = v12
    case "\n\t":
        fmt.Println(p0)
    case "a":
        p0 = "foo bar"
    default:
        if ((!false) && (-2147483648 == -2147483648)) {
            var v13 int = -1
            _ = v13
        } else {
            p0 = ""
            var v14 float64 = ((float64(3.14159) + (0.0 - -2.25)) + float64((float32(3.14159) + (0.0 / 1.0))))
            _ = v14
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint8) uint8 {
    if true {
        v15 := GenS2{F0: false}
        var v16 int32 = 524
        for fi17 := 0; fi17 < 16; fi17++ {
            if true {
                p0 = p0
            }
            var v18 int = (fi17 + 32768)
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    } else {
        var v19 string = (("foo bar" + "") + ("" + ("foo bar" + "foo bar")))
        p0 -= p0
        var v20 uint16 = 321
        _ = v19
        _ = v20
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        v21 := GenS2{F0: true}
        _ = v21
    default:
        fmt.Println(p0)
    }
    v22 := []int{9223372036854775807, 822}
    v23 := GenS2{F0: false}
    v24 := GenS3{F0: 0, F1: 341, F2: 2147483647}
    _ = p0
    _ = v22
    _ = v23
    _ = v24
    return (p0 << ((1) & 63))
}

func main() {
    var v1 int16 = 777
    v2 := make(map[string]int)
    var v3 uint16 = 32767
    v4 := make(map[int]string)
    v5 := []uint32{172, 0, 26}
    var v6 uint64 = 518
    for fi25 := 0; fi25 < 15; fi25++ {
        v1 += -371
        _ = fi25
    }
    switch "\n\t" {
    case "foo bar":
        var v26 string = "foo bar"
        _ = v26
    case "a":
        var v27 float32 = ((float32(0.1) * (1.5 + 3.14159)) - (-2.25 + float32(1e10)))
        _ = v27
    }
    switch -511 {
    case 32767:
        if false {
            v28 := v4[127]
            _ = v28
        } else {
            v3 = genFn0(("hello" + ("a" + "\n\t")))
            switch 127 {
            case 2147483647:
                v1 -= v1
            case -9223372036854775808:
                var v29 float32 = 0.1
                _ = v29
            case 0:
                v3 = v3
            default:
                fmt.Println(v3, v6, v6)
            }
        }
    case -932:
        if (42 < 0) {
            v6 += (v6 | v6)
            v6 = (v6 / ((v6) | 1))
            if false {
                fmt.Println(v3, v3, v1)
                fmt.Println(v6, v1)
            } else {
                v5 = append(v5, 2147483648)
            }
        } else {
            v30 := GenS3{F0: v6, F1: 510, F2: 42}
            _ = v30
        }
    default:
        v4[9223372036854775807] = ""
    }
    v6 = (v6 * 0)
    v31 := GenS4{F0: (^(v3)), F1: 2147483647, F2: 1}
    fmt.Println(v3, v6)
    v5 = append(v5, 470)
    var v32 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
