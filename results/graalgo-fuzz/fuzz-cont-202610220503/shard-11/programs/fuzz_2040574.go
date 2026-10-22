package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int) uint8 {
    var v7 int32 = 0
    v7 += v7
    fmt.Println(p0, v7)
    v7 = (v7 & v7)
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 string, p2 uint) uint32 {
    for fi8 := 0; fi8 < 1; fi8++ {
        p1 = p1
        for fi9 := 0; fi9 < 19; fi9++ {
            if false {
                var v10 int = (fi9 >> (((p2 + (^(p2)))) & 63))
                v11 := []int{v10, (fi8 / ((-2147483648) | 1)), fi8}
                p2 = p2
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        var v12 float32 = (1e10 / ((-2.25 * (((1.5 / (3.14159 * 1e10)) * 0.1) * -2.25)) / 3.14159))
        _ = fi8
        _ = v12
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 324
}

func main() {
    var v1 bool = false
    var v2 bool = (!(((float64(float32(float64((0.0 * 1.5)))) >= 1.5) || (v1 && v1)) || true))
    var v3 int8 = -15
    var v4 int64 = -246
    var v5 int32 = -1
    var v6 int = 42
    var v13 uint8 = 0
    v14 := []int{(v6 | v6)}
    v13 = (v13 + 255)
    var v15 int8 = 0
    var v16 int = (v6 ^ -2147483648)
    var v17 string = ((("a" + ("a" + "\n\t")) + "\n\t") + "\n\t")
    var v18 uint = 42
    if v1 {
        v19 := GenS1{F0: v13, F1: 65535}
        _ = v19
    } else {
        fmt.Println(v3, v16, v3)
        var v20 uint = (v18 ^ v18)
        _ = v20
    }
    v2 = (v1 || ((v4 >= (^(v4))) && (0.0 < 1.5)))
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            if v2 {
                if (!true) {
                    var v23 string = "hello"
                    _ = v23
                } else {
                    v14 = append(v14, v6)
                    v24 := GenS1{F0: (v13 + 255), F1: genFn1((((((v17 + (("hello" + (v17 + v17)) + v17)) + v17) + "foo bar") + v17) > (("a" + v17) + "hello")), v17, (v18 ^ (v18 << ((42) & 63))))}
                    _ = v24
                }
                switch "" {
                case "\n\t":
                    var v25 bool = (v1 && false)
                    _ = v25
                }
            } else {
                var v26 uint32 = 0
                switch v6 {
                case 127:
                    switch (("hello" + v17) + v17) {
                    case "a":
                        v17 = v17
                    }
                case -978:
                    if len(v14) > 0 {
                        v14[0] = 4294967295
                    }
                case -1:
                    var v27 int8 = v3
                    _ = v27
                default:
                    if v2 {
                        _ = len(v14)
                        v28 := GenS1{F0: genFn0(fi22), F1: v26}
                        var v29 string = (v17 + (v17 + "foo bar"))
                        _ = v28
                        _ = v29
                    }
                }
                var v30 uint32 = v26
                _ = v26
                _ = v30
            }
            var v31 int8 = v3
            v32 := GenS1{F0: v13, F1: uint32(v13)}
            _ = fi22
            _ = v31
            _ = v32
        }
        var v33 uint16 = 42
        _ = fi21
        _ = v33
    }
    fmt.Println(v18, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
