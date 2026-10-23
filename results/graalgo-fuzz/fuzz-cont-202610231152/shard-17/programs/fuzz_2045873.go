package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 int16 }
func genFn0() uint {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float64 = 3.14159
        if ("" <= ((("a" + "\n\t") + "foo bar") + "foo bar")) {
            v6 := GenS1{F0: 42, F1: ("\n\t" + ("a" + ("\n\t" + ("a" + "a")))), F2: true}
            _ = v6
        } else {
            v5 -= (3.14159 + 0.0)
            v5 += v5
        }
        _ = fi4
        _ = v5
    }
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: (v7 ^ (^(v7))), F1: "hello", F2: false}
        _ = fi8
        _ = v9
    }
    _ = v7
    return 4294967295
}
func genFn1(p0 int32, p1 int32, p2 uint16) uint {
    for fi10 := 0; fi10 < 12; fi10++ {
        switch (fi10 << ((845) & 63)) {
        case 127:
            v11 := []int{fi10, (fi10 % ((782) | 1)), 4294967295, (fi10 + int(p2))}
            _ = v11
        default:
            var v12 uint8 = 127
            _ = v12
        }
        var v13 float64 = -2.25
        _ = fi10
        _ = v13
    }
    if ((p0 != p0) || false) {
        var v14 float64 = 0.0
        _ = v14
    }
    p1 = p0
    switch 42 {
    case -751:
        var v15 bool = true
        _ = v15
    case -922:
        p2 += 32767
    case 127:
        v16 := []int{}
        _ = v16
    default:
        if false {
            var v17 float32 = 0.1
            var v18 float64 = (((((3.14159 / (float64((float32(1.5) * (v17 - ((v17 + v17) / v17)))) * -2.25)) * 1.5) - float64(((v17 - (v17 - float32((float64((v17 + ((v17 / v17) + (0.1 + ((float32(1.5) / v17) * ((v17 / -2.25) - 1.5)))))) - 0.0)))) / 1e10))) - -2.25) / 1.0)
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return genFn0()
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 1
    var v3 uint32 = 6
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := []int{fi19, fi19, fi19, 231}
        fmt.Println(v1, v2, v3)
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v2)
    fmt.Println(v2, v3, v2)
    v1 -= (^(v1))
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int64 = 9223372036854775807
        switch ("foo bar" + "foo bar") {
        case "a":
            if (255 >= genFn1((v1 - 1), v1, 0)) {
                fmt.Println(fi21, v22)
                var v23 float32 = 1.5
                v24 := GenS1{F0: 102, F1: ("" + "\n\t"), F2: false}
                _ = v23
                _ = v24
            } else {
                switch fi21 {
                case 32768:
                    v22 += v22
                default:
                    v25 := []int{fi21, (fi21 << ((0) & 63))}
                    _ = v25
                }
            }
        default:
            var v26 int = (-(fi21))
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    var v27 uint32 = 417
    var v28 float64 = ((float64((((1.5 / 0.1) + 1.5) / 1.0)) * 0.1) + 1.5)
    if ((true && true) || (0 < 42)) {
        v3 += v3
        var v29 int64 = -1
        fmt.Println(v27, v28, v1)
        _ = v29
    }
    fmt.Println(v2)
    fmt.Println(v2, v3, v3)
    var v30 uint8 = 2
    v31 := []int{-2147483648, 1, 480}
    if false {
        var v32 float64 = (0.1 - v28)
        _ = v32
    }
    if false {
        v2 -= (v2 - (v2 >> ((159) & 63)))
        if ((^(v30)) == v30) {
            var v33 int = -205
            if true {
                v33 = 859
                v34 := GenS2{F0: int(v3), F1: 0}
                _ = v34
            } else {
                v28 = v28
                v27 = (v3 >> ((4294967295) & 63))
                if (0 <= 32767) {
                    v28 += v28
                    _ = len(v31)
                    v35 := GenS1{F0: (v30 & (v30 % (((v30 + v30)) | 1))), F1: "foo bar", F2: (v28 <= v28)}
                    _ = v35
                }
            }
            var v36 float32 = ((3.14159 + -2.25) - 0.0)
            _ = v33
            _ = v36
        }
        fmt.Println(v1, v1, v28)
    }
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
}
