package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint; F2 int }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
func genFn0() int64 {
    for fi4 := 0; fi4 < 11; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    v5 := GenS1{F0: 37, F1: 1, F2: 32}
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (3.14159 + -2.25)
    v3 := []int{423, 32768, 0, -2147483648}
    fmt.Println(v2)
    var v6 uint64 = 827
    switch 32767 {
    case -994:
        if len(v3) > 0 {
            v3[0] = 337
        }
    }
    var v7 float32 = v2
    v2 -= v2
    var v8 uint = 0
    switch 128 {
    case 1:
        var v9 int64 = 42
        _ = v9
    case -2147483648:
        var v10 int = -1
        _ = v10
    case 2147483647:
        for fi11 := 0; fi11 < 20; fi11++ {
            v7 = 0.1
            _ = fi11
        }
    default:
        var v12 float32 = (v7 / v7)
        _ = v12
    }
    var v13 bool = (genFn0() >= -344)
    switch ("hello" + "a") {
    case "a":
        var v14 uint = (^(v8))
        _ = v14
    case "hello":
        if v13 {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 int8 = 127
                var v17 uint8 = 2
                var v18 uint = (v8 >> ((4294967295) & 63))
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 float64 = 3.14159
            switch ("foo bar" + ("" + "foo bar")) {
            case "":
                switch 2147483647 {
                case 32767:
                    v20 := GenS2{F0: 4294967295, F1: 255, F2: 720}
                    _ = v20
                default:
                    if (true || (4294967295 > 2147483648)) {
                        v21 := GenS1{F0: 42, F1: 211, F2: 0}
                        _ = v21
                    }
                }
            case "hello":
                v22 := GenS3{F0: (v8 << (((^(v8))) & 63)), F1: (!false), F2: (0.1 <= (((((1.5 + v19) / -2.25) + (v19 + float64((v2 - ((v2 * (0.0 / 3.14159)) + v7))))) / (v19 - v19)) - 0.1))}
                _ = v22
            }
            _ = v19
        } else {
            var v23 int8 = -1
            _ = v23
        }
    }
    var v24 string = "a"
    _ = len(v3)
    var v25 uint = v8
    v7 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
