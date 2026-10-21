package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32) uint32 {
    v5 := []int{9223372036854775807, 127}
    var v6 uint16 = 2
    var v7 string = "foo bar"
    switch "hello" {
    case "hello":
        if (((946 > 32767) || true) && (true && true)) {
            if (true && true) {
                v7 = (v7 + "hello")
                var v8 int = 2147483647
                _ = v8
            }
        }
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 230
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 0
    v3 := []int{-307, 554, -2147483648, 42}
    var v4 string = "\n\t"
    v9 := GenS1{F0: true}
    v2 = (v2 / (((v2 * v2)) | 1))
    v4 = ("a" + v4)
    v10 := GenS2{F0: 221, F1: 255, F2: 95}
    if (1 > 992) {
        var v11 bool = true
        switch (("\n\t" + v4) + (v4 + "hello")) {
        case "hello":
            v12 := GenS1{F0: v11}
            _ = v12
        case "":
            v13 := GenS1{F0: ((v4 >= v4) || (-2147483648 <= 42))}
            _ = v13
        }
        _ = v11
    } else {
        var v14 int16 = 832
        v15 := GenS2{F0: 45, F1: 42, F2: 0}
        v14 = 32767
        _ = v14
        _ = v15
    }
    _ = len(v1)
    var v16 int16 = -1
    v3 = append(v3, 4294967295)
    for fi17 := 0; fi17 < 18; fi17++ {
        v1[(v4 + "hello")] = v4
        _ = fi17
    }
    if true {
        var v18 bool = true
        _ = v18
    } else {
        if (128 <= 64) {
            var v19 uint64 = 1
            if (0 < 0) {
                v20 := GenS2{F0: 0, F1: 351, F2: 255}
                v21 := GenS1{F0: false}
                _ = v20
                _ = v21
            } else {
                fmt.Println(v4, v19, v2)
            }
            _ = v19
        } else {
            fmt.Println(v16, v16)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v22 := GenS3{F0: ((("a" + "hello") + (v4 + "")) + "a"), F1: -2147483648, F2: 1}
        _ = v22
    }
    v23 := GenS1{F0: true}
    var v24 int32 = 725
    v25 := GenS4{F0: false, F1: 18446744073709551615}
    var v26 float64 = float64(-2.25)
    v16 -= 1000
    var v27 bool = ((!(((v24 != v24) && true) || false)) || false)
    var v28 uint32 = (v2 % ((v2) | 1))
    var v29 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
