package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0() int64 {
    var v4 uint = 290
    if ((3.14159 * 0.1) == 0.1) {
        for fi5 := 0; fi5 < 4; fi5++ {
            v4 = 903
            v4 = 2
            if ((true && true) || (-866 != 1)) {
                fmt.Println(v4, fi5)
            } else {
                fmt.Println(fi5, fi5)
                fmt.Println(v4)
                v6 := GenS1{F0: (v4 & (v4 >> ((v4) & 63))), F1: 255}
                _ = v6
            }
            _ = fi5
        }
    }
    _ = v4
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := []string{"\n\t", ("" + "")}
    v3 := make(map[string]int)
    var v7 string = ("foo bar" + "hello")
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if true {
                var v10 uint64 = 0
                v11 := GenS2{F0: 42, F1: 128, F2: -1}
                fmt.Println(fi9)
                _ = v10
                _ = v11
            } else {
                if len(v2) > 0 {
                    v2[0] = v7
                }
                var v12 float64 = float64((v1 / v1))
                _ = v12
            }
            var v13 uint = 61
            switch v7 {
            case "a":
                v7 = v7
            case "hello":
                v3[(v7 + "")] = (fi9 & (^(fi9)))
            case "foo bar":
                v14 := GenS1{F0: (v13 * v13), F1: 2}
                _ = v14
            }
            _ = fi9
            _ = v13
        }
        _ = fi8
    }
    v7 = v7
    v7 = v7
    v15 := GenS1{F0: 2, F1: 65535}
    v7 = "a"
    v16 := GenS1{F0: 65535, F1: 255}
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(v7, v7, v1)
        for fi18 := 0; fi18 < 12; fi18++ {
            if ((((1.5 / 1.5) + (-2.25 + 0.1)) - 0.1) >= (-2.25 + (((((1.5 * -2.25) - float64((v1 - v1))) / (1e10 * (float64(v1) + (1e10 * 0.0)))) / (3.14159 * -2.25)) - 0.1))) {
                v19 := GenS1{F0: 2, F1: 42}
                v7 = ""
                _ = v19
            }
            _ = len(v2)
            v20 := GenS3{F0: 327, F1: 0, F2: 255}
            _ = fi18
            _ = v20
        }
        _ = fi17
    }
    v1 = (0.1 - v1)
    if false {
        _ = len(v2)
        v21 := GenS3{F0: 454, F1: -1, F2: 45}
        _ = v21
    }
    var v22 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v7
}
