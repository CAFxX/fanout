package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int32; F2 int }
type GenS4 struct { F0 int; F1 string }

func main() {
    var v1 uint = 970
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 uint8 = uint8(v1)
    if v3 {
        var v5 uint16 = 42
        _ = v5
    }
    switch 260 {
    case -738:
        if (!(!(42 <= -9223372036854775808))) {
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 string = "hello"
                delete(v2, ((v7 + v7) + (v7 + v7)))
                if false {
                    switch v7 {
                    case "a":
                        v8 := []int{(fi6 + (fi6 >> ((v1) & 63))), fi6, fi6}
                        _ = v8
                    case "":
                        v9 := []int{fi6, 127, (fi6 & (-(fi6)))}
                        _ = v9
                    }
                }
                _ = fi6
                _ = v7
            }
            if (v3 || (42 > -1)) {
                v10 := GenS1{F0: 255, F1: 864, F2: 32767}
                v3 = false
                _ = v10
            }
        }
    default:
        delete(v2, "\n\t")
    }
    var v11 uint8 = v4
    v3 = (128 == (^(v11)))
    if (!((v4 - 234) != v11)) {
        var v12 int32 = -2147483648
        v4 = 10
        _ = v12
    } else {
        var v13 uint16 = 42
        v3 = v3
        _ = v13
    }
    if (((v3 || (v1 < (v1 - v1))) && false) || (795 != 624)) {
        if (1 < 323) {
            fmt.Println(v3, v3, v1)
        } else {
            v14 := GenS2{F0: "\n\t"}
            v4 -= v4
            _ = v14
        }
    } else {
        v15 := GenS2{F0: "a"}
        var v16 uint = (v1 | v1)
        delete(v2, "\n\t")
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: (("\n\t" + ("hello" + "")) + ("\n\t" + (("a" + "") + "a")))}
    var v18 float32 = (float32(((float64(1.5) / float64(((1e10 * 0.0) - ((0.0 + float32((1e10 - float64(0.1)))) / (float32(((-2.25 / 0.1) * (0.0 / 3.14159))) - 0.1))))) / 0.1)) * 0.0)
    v19 := v2[("foo bar" + "a")]
    var v20 float32 = v18
    v1 -= v1
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: 65535, F1: -705, F2: int16(v1)}
        var v23 int = (fi21 + (fi21 << ((0) & 63)))
        v24 := GenS3{F0: (v1 / ((4294967295) | 1)), F1: -1, F2: (fi21 >> ((v1) & 63))}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            if false {
                switch (fi25 << ((272) & 63)) {
                case 462:
                    switch fi25 {
                    case 2147483647:
                        if (2 == 65535) {
                            v27 := GenS4{F0: fi26, F1: "hello"}
                            _ = len(v2)
                            _ = v27
                        }
                    case 42:
                        fmt.Println(v11, v1, v19)
                    case 32767:
                        v28 := []int{-344, fi25}
                        _ = v28
                    }
                }
            }
            var v29 int64 = -406
            v4 = uint8(v19)
            _ = fi26
            _ = v29
        }
        _ = fi25
    }
    v30 := GenS3{F0: v1, F1: -23, F2: 9223372036854775807}
    var v31 int = 32768
    var v32 float64 = ((float64((-2.25 - v20)) * 1.5) * float64((v18 - v18)))
    if (117 < 106) {
        if (1.5 < v20) {
            v11 = (v4 << (((v1 | (v1 - v1))) & 63))
        }
        var v33 int32 = -2147483648
        var v34 float64 = v32
        _ = v33
        _ = v34
    }
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
