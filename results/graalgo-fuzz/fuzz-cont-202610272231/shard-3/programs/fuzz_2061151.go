package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint = 1
    var v2 uint32 = 125
    v3 := make(map[string]string)
    v4 := GenS1{F0: (-45 != 100)}
    if (false && true) {
        if (1 < -1) {
            if false {
                v5 := GenS2{F0: 2, F1: 1, F2: 255}
                switch 4294967295 {
                case 32767:
                    v1 = (v1 ^ 4294967295)
                }
                _ = v5
            }
            for fi6 := 0; fi6 < 5; fi6++ {
                v2 = v2
                v7 := GenS2{F0: uint8(v2), F1: 0, F2: 127}
                _ = fi6
                _ = v7
            }
        } else {
            if (810 != 1) {
                var v8 string = ("\n\t" + "\n\t")
                _ = v8
            } else {
                v1 -= (v1 - v1)
                v9 := GenS1{F0: false}
                switch (("foo bar" + (("hello" + "hello") + "")) + "") {
                case "\n\t":
                    v10 := []int{-2147483648, -2147483648, -30, -194}
                    _ = v10
                case "hello":
                    switch "" {
                    case "hello":
                        var v11 uint32 = (v2 >> (((v1 / (((^(v1))) | 1))) & 63))
                        _ = v11
                    case "a":
                        v12 := []int{84, 32768, -405}
                        _ = v12
                    }
                case "a":
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v1 = 1
                        _ = fi13
                    }
                }
                _ = v9
            }
            v2 = 0
            v2 -= v2
        }
        v2 -= v2
    } else {
        v14 := GenS2{F0: 2, F1: -432, F2: 1}
        _ = v14
    }
    if ((!(v2 == 1)) || true) {
        var v15 int = 70
        _ = len(v3)
        fmt.Println(v2, v15, v15)
        _ = v15
    } else {
        v2 = v2
        fmt.Println(v2)
    }
    var v16 uint = 69
    var v17 uint16 = 0
    v2 = (^(v2))
    v2 -= 4294967295
    if true {
        if false {
            v1 += (v1 / ((982) | 1))
            v2 -= (v2 ^ v2)
        } else {
            v16 = 1
        }
        var v18 int = 32767
        _ = v18
    }
    v16 += 578
    var v19 uint8 = 255
    fmt.Println(v17, v16, v16)
    if (!true) {
        if false {
            var v20 float64 = 3.14159
            v20 = float64(3.14159)
            v21 := GenS3{F0: v17}
            _ = v20
            _ = v21
        }
        var v22 uint32 = (v2 * 491)
        _ = v22
    }
    v2 -= 2
    v19 -= v19
    v23 := []int{-9223372036854775808, 127}
    delete(v3, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
