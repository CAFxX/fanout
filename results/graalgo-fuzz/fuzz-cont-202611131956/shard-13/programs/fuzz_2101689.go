package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0() uint16 {
    switch "hello" {
    case "":
        var v5 float64 = 3.14159
        _ = v5
    case "\n\t":
        var v6 int16 = 1
        _ = v6
    case "foo bar":
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    var v8 uint16 = 875
    _ = v8
    return (^(v8))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    var v4 int8 = -128
    fmt.Println(v2, v4)
    v9 := []int{-1, 863}
    v10 := GenS1{F0: 951, F1: 42}
    v9 = append(v9, 128)
    fmt.Println(v4)
    var v11 uint = 2
    var v12 int = -2147483648
    var v13 uint8 = 255
    for fi14 := 0; fi14 < 20; fi14++ {
        if true {
            if (!(v11 <= (v11 - (^(v11))))) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    var v16 string = ("hello" + "a")
                    fmt.Println(v4, v11, v4)
                    v17 := GenS1{F0: (v11 << (((^(v11))) & 63)), F1: 1}
                    var v18 uint = (^(v11))
                    v4 += v4
                    v19 := GenS1{F0: (v18 ^ 106), F1: -1}
                    var v20 uint64 = 1
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = len(v3)
            } else {
                var v21 float32 = 1.5
                var v22 string = "foo bar"
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int8 = v4
                    _ = fi23
                    _ = v24
                }
                _ = v21
                _ = v22
            }
            _ = len(v1)
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            fmt.Println(v12, v11, v11)
            var v25 int16 = 32767
            _ = v25
        }
        _ = fi14
    }
    if (!(false || true)) {
        v26 := GenS1{F0: v11, F1: -1}
        _ = v26
    } else {
        v27 := GenS1{F0: v11, F1: 42}
        var v28 string = ""
        _ = v27
        _ = v28
    }
    v4 = (v4 & -128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
