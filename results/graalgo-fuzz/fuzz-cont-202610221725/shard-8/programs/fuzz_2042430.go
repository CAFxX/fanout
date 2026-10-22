package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 string) uint16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = 2
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2, F1: -113, F2: 213}
    _ = p0
    _ = v8
    return 0
}
func genFn1() uint8 {
    var v9 float32 = (-2.25 - float32(1.5))
    v10 := GenS1{F0: 113, F1: 93, F2: 0}
    if (0 != -901) {
        fmt.Println(v9)
        var v11 bool = ("\n\t" > "\n\t")
        _ = v11
    } else {
        if (2 == 65535) {
            v12 := []int{865, -9223372036854775808, -2147483648, -1}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                if (((("" + ("" + "\n\t")) + "a") + "a") <= "foo bar") {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        if (!true) {
                            var v15 int64 = 9223372036854775807
                            _ = v15
                        }
                        _ = fi14
                    }
                    v16 := []int{fi13, (fi13 * 9223372036854775807), fi13, (^(fi13))}
                    _ = v16
                }
                _ = fi13
            }
        }
    }
    v17 := GenS2{F0: false}
    _ = v10
    _ = v17
    _ = v9
    return 42
}

func main() {
    v1 := []int{-218}
    var v2 uint8 = 159
    v3 := make(map[string]uint32)
    var v4 uint32 = 223
    var v5 int16 = 0
    var v18 float64 = (0.0 / 1.0)
    var v19 int32 = -547
    if ((v18 != (v18 * 1e10)) || true) {
        if (!false) {
            v4 += v4
            if (32768 == -816) {
                if true {
                    var v20 float64 = 0.1
                    _ = v20
                } else {
                    fmt.Println(v2)
                    v21 := GenS3{F0: 32767, F1: 0}
                    _ = v21
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    v23 := v3["\n\t"]
                    _ = fi22
                    _ = v23
                }
            } else {
                v24 := GenS2{F0: true}
                _ = v24
            }
        }
    } else {
        v25 := GenS4{F0: 42}
        var v26 uint64 = 78
        _ = v25
        _ = v26
    }
    var v27 uint16 = 0
    v28 := GenS1{F0: v2, F1: 1, F2: 2}
    var v29 int = 939
    var v30 uint = 1
    v27 = 42
    if (3.14159 < v18) {
        fmt.Println(v2, v19)
    }
    var v31 uint8 = v2
    switch "hello" {
    case "foo bar":
        _ = len(v1)
    case "hello":
        var v32 uint64 = uint64(v27)
        _ = v32
    }
    v3["\n\t"] = 1
    var v33 int64 = 42
    v19 = -983
    v18 += (((-2.25 / (v18 / v18)) * (v18 / v18)) * (v18 / ((float64(0.1) / float64(1e10)) * (-2.25 + 1e10))))
    var v34 int16 = (v5 ^ (v5 / (((v5 + v5)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
