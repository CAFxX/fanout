package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 2
    var v3 float64 = 1.5
    var v4 int64 = -1
    if false {
        v5 := []int{128, -1, -155, 32768}
        if ((658 < 2) && false) {
            if false {
                if len(v5) > 0 {
                    v5[0] = 127
                }
                v6 := GenS1{F0: 2147483647, F1: 749, F2: "foo bar"}
                v3 = v3
                _ = v6
            } else {
                switch ("foo bar" + ("foo bar" + "a")) {
                case "hello":
                    if ((float32(((v3 * v3) / 1e10)) * ((float32(float64((1e10 + 0.1))) * (-2.25 * (-2.25 * 3.14159))) + (-2.25 * 1.5))) != float32(3.14159)) {
                        v1[2147483648] = 0
                    }
                }
            }
        }
        var v7 int8 = -59
        _ = v5
        _ = v7
    }
    v8 := []int{-628, 128, -9223372036854775808}
    v9 := GenS1{F0: 128, F1: 0, F2: "\n\t"}
    v10 := GenS1{F0: -453, F1: 113, F2: ("" + "foo bar")}
    if true {
        var v11 bool = false
        switch -9223372036854775808 {
        case 260:
            switch 2147483647 {
            case 466:
                v12 := GenS1{F0: 1, F1: 4294967295, F2: ("hello" + "")}
                _ = v12
            case 32767:
                v13 := GenS2{F0: 0, F1: 4294967295, F2: 4294967295}
                _ = v13
            case 323:
                v2 += 195
            }
        default:
            var v14 string = "a"
            _ = v14
        }
        _ = v11
    }
    v15 := GenS3{F0: -32768, F1: 739}
    var v16 uint64 = 756
    var v17 uint64 = (v16 - v16)
    delete(v1, 181)
    switch int(v4) {
    case 853:
        v3 = v3
    }
    var v18 int64 = (v4 + v4)
    var v19 uint64 = (v16 ^ (v16 & (^(v17))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
