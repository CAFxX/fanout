package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int16, p1 uint) int32 {
    var v6 uint32 = 111
    p1 = (p1 / ((p1) | 1))
    p1 = (p1 % (((p1 * 65535)) | 1))
    var v7 uint = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string) int {
    v8 := GenS1{F0: 1}
    if true {
        var v9 uint16 = 263
        fmt.Println(v9)
        v9 = v9
        _ = v9
    } else {
        p0 = "hello"
    }
    _ = p0
    _ = v8
    return 128
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]string)
    var v3 uint16 = 279
    var v4 int32 = 1
    var v5 uint32 = 65535
    if true {
        var v10 int16 = 1000
        var v11 int64 = 709
        fmt.Println(v5, v4, v10)
        _ = v10
        _ = v11
    } else {
        var v12 uint = 1
        v1 -= (((v1 - 1.5) / v1) - ((0.0 - float64((0.1 * 1e10))) - (1.5 - float64(((((3.14159 / (0.0 + -2.25)) * 1.5) * 1e10) - 1.5)))))
        switch "a" {
        case "hello":
            fmt.Println(v12)
        case "foo bar":
            fmt.Println(v12, v12, v3)
        default:
            v4 = v4
        }
        _ = v12
    }
    switch 42 {
    case 0:
        fmt.Println(v4, v4, v5)
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS2{F0: v5}
        var v15 string = "a"
        fmt.Println(v15)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS3{F0: (!(v5 >= (v5 ^ v5))), F1: v5}
    fmt.Println(v3, v4)
    v17 := []int{int(v5), -977}
    v2[genFn1("\n\t")] = "a"
    v4 = v4
    for fi18 := 0; fi18 < 13; fi18++ {
        delete(v2, 128)
        v19 := GenS3{F0: (128 >= 127), F1: (v5 + uint32(v3))}
        v20 := GenS4{F0: ((!(-128 <= -89)) || ((v5 * v5) <= v5)), F1: 200}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS5{F0: -1}
    v2[int(v5)] = (("\n\t" + ("\n\t" + "foo bar")) + ("\n\t" + (("\n\t" + "a") + "hello")))
    v1 = v1
    if false {
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS1{F0: 1000}
            _ = fi22
            _ = v23
        }
        fmt.Println(v1, v4)
    }
    if (false || ((v3 / ((uint16(v5)) | 1)) < v3)) {
        delete(v2, 127)
        if false {
            v24 := GenS5{F0: -80}
            _ = v24
        } else {
            var v25 int = 128
            var v26 uint16 = (v3 - 1)
            for fi27 := 0; fi27 < 11; fi27++ {
                if true {
                    v3 += v26
                    switch "" {
                    case "\n\t":
                        v28 := GenS5{F0: int8(fi27)}
                        _ = v28
                    default:
                        v26 = (v3 / (((^(v26))) | 1))
                    }
                    v29 := GenS1{F0: 1000}
                    _ = v29
                } else {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    delete(v2, (v25 + fi27))
                    v4 = (v4 & v4)
                }
                var v30 uint = 323
                v4 = v4
                _ = fi27
                _ = v30
            }
            _ = v25
            _ = v26
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 uint16 = v3
        var v33 int = fi31
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 int8 = -77
    var v35 uint8 = 18
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
