package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
type GenS4 struct { F0 int32; F1 uint; F2 string }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() int32 {
    switch ((("a" + "hello") + ("hello" + "foo bar")) + ("hello" + "hello")) {
    case "\n\t":
        var v7 uint64 = 555
        _ = v7
    case "a":
        var v8 int32 = -665
        _ = v8
    default:
        var v9 float32 = float32((0.0 + (float64((1e10 * 0.0)) * 1e10)))
        _ = v9
    }
    if true {
        var v10 int16 = 130
        v11 := GenS1{F0: 127, F1: 519}
        var v12 float64 = ((1.5 * (float64((0.0 * ((1.5 - 3.14159) - 3.14159))) / (-2.25 / -2.25))) - 1e10)
        _ = v10
        _ = v11
        _ = v12
    }
    switch -1 {
    case 32767:
        v13 := GenS2{F0: 100}
        _ = v13
    case 0:
        var v14 bool = (255 > 2)
        _ = v14
    }
    var v15 int32 = 2147483647
    v15 -= (v15 & (v15 >> ((1) & 63)))
    _ = v15
    return (^(v15))
}
func genFn1(p0 int32) uint64 {
    var v16 uint64 = 0
    var v17 bool = true
    v18 := GenS3{F0: (((471 >= 514) && ((1e10 != (0.0 - 3.14159)) && (-105 == -126))) || (("" + ("hello" + "")) == "hello")), F1: 94, F2: -1}
    fmt.Println(v17, p0)
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return (v16 & (v16 >> ((4294967295) & 63)))
}

func main() {
    v1 := []uint32{}
    var v2 bool = true
    var v3 string = "a"
    var v4 int8 = 100
    var v5 bool = v2
    var v6 int8 = (^(v4))
    switch 9223372036854775807 {
    case 32768:
        for fi19 := 0; fi19 < 5; fi19++ {
            if ((386 <= 65535) && v2) {
                if (-1 <= 0) {
                    v6 -= 1
                    v2 = (0 > 128)
                }
            }
            _ = fi19
        }
    default:
        v20 := GenS4{F0: 0, F1: 65535, F2: "hello"}
        _ = v20
    }
    v4 = (-(v4))
    var v21 uint16 = 32767
    _ = len(v1)
    v2 = v2
    switch v3 {
    case "":
        v1 = append(v1, 65535)
    }
    if ((false && false) && (v2 && v2)) {
        v22 := GenS5{F0: -185, F1: (((42 < 127) && v5) && (!(!(((v2 || false) && (127 <= 0)) && ("foo bar" <= v3)))))}
        if (42 <= -296) {
            _ = len(v1)
            v23 := GenS5{F0: -1, F1: true}
            var v24 uint8 = 255
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 5; fi25++ {
                _ = len(v1)
                v26 := GenS2{F0: v6}
                _ = fi25
                _ = v26
            }
        }
        _ = v22
    }
    v5 = v5
    var v27 bool = (0.0 > 0.0)
    switch 187 {
    case 305:
        v4 = (v4 - (v4 ^ (-(v4))))
    case -1:
        var v28 int8 = v4
        _ = v28
    case 31:
        v6 = v6
    }
    v6 += 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
