package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 uint8, p2 string) int16 {
    v5 := []int{343, 128, 128, 42}
    p2 = (p2 + p2)
    fmt.Println(p1, p2, p1)
    p2 = p2
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = "hello"
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -32768
}

func main() {
    var v1 float64 = 0.0
    var v2 uint16 = 2
    v3 := []string{"\n\t", "\n\t", (("" + ("a" + "\n\t")) + "\n\t"), ((("a" + ((("foo bar" + "\n\t") + ("foo bar" + ("hello" + "a"))) + ("a" + "\n\t"))) + "hello") + ("a" + ("foo bar" + "")))}
    var v4 string = (("" + "a") + "")
    for fi8 := 0; fi8 < 16; fi8++ {
        if (172 >= 42) {
            if false {
                v4 = (v4 + (v4 + "\n\t"))
            }
            var v9 int = -1
            v10 := GenS1{F0: 127, F1: v9}
            _ = v10
            _ = v9
        } else {
            var v11 int64 = 42
            var v12 int = int(v11)
            if false {
                v11 -= (^(v11))
                for fi13 := 0; fi13 < 14; fi13++ {
                    if len(v3) > 0 {
                        v3[0] = v4
                    }
                    _ = fi13
                }
            }
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    var v14 int = -1
    if (!(-27 <= -128)) {
        if false {
            v15 := GenS1{F0: 128, F1: (v14 - (-(v14)))}
            _ = v15
        }
    }
    _ = len(v3)
    if ("\n\t" < v4) {
        var v16 int = -196
        if ((v1 == (v1 * (float64(float32(v1)) + float64(((float32(v1) - float32((1.5 - v1))) - (-2.25 * (3.14159 / 1.5))))))) && (true || true)) {
            var v17 float32 = 0.0
            _ = v17
        }
        _ = v16
    } else {
        var v18 uint = 1
        fmt.Println(v4, v18)
        if false {
            var v19 uint16 = 65535
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v19
        } else {
            v20 := GenS1{F0: 128, F1: (v14 << ((v18) & 63))}
            _ = v20
        }
        _ = v18
    }
    switch v14 {
    case 127:
        fmt.Println(v4, v14)
    case 863:
        if ((!((float64(((float32(float64(0.0)) - 0.1) / 1.0)) - (v1 / float64(1e10))) < 0.0)) && true) {
            v2 = 32767
            if (true || true) {
                v14 = v14
                var v21 bool = true
                _ = v21
            }
            v3 = append(v3, v4)
        } else {
            fmt.Println(v1, v14)
            switch v14 {
            case 127:
                v22 := GenS1{F0: 2, F1: v14}
                _ = v22
            case 4294967295:
                fmt.Println(v1, v1, v14)
            }
        }
    case 128:
        v4 = ("" + (v4 + "\n\t"))
    }
    v23 := GenS1{F0: 2, F1: (v14 + v14)}
    if (false || true) {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 int16 = 42
            v26 := GenS1{F0: 127, F1: (fi24 & v14)}
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    v27 := GenS2{F0: 79}
    v28 := GenS2{F0: 42}
    fmt.Println(v2)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v29 uint8 = 2
        _ = v29
    } else {
        v30 := GenS3{F0: (v4 + v4), F1: 700}
        _ = v30
    }
    var v31 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
