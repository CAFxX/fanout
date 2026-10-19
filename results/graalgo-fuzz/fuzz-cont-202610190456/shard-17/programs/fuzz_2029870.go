package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 uint }
func genFn0() uint64 {
    var v7 int32 = 2147483647
    v7 += (v7 >> ((726) & 63))
    var v8 int = 630
    if false {
        v8 = (v8 - v8)
        var v9 string = "hello"
        if (!(1.5 <= float32(1.5))) {
            var v10 string = ""
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: (v7 | v7), F1: 2}
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            if true {
                if (1 > 42) {
                    var v13 float32 = 1.5
                    var v14 uint32 = 0
                    var v15 float32 = ((1e10 + v13) - 1e10)
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            v16 := []int{(^(v8)), -9223372036854775808, v8, v8}
            fmt.Println(v9, v9)
            _ = v16
        }
        _ = v9
    } else {
        v8 -= (v8 * v8)
        fmt.Println(v7, v8)
    }
    if (true && (!true)) {
        var v17 uint8 = 159
        v18 := []int{(v8 * 127), (v8 ^ (v8 + v8)), v8, v8}
        v17 = v17
        _ = v17
        _ = v18
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 string = ("" + "a")
    v4 := make(map[int]string)
    v5 := []string{}
    var v6 int64 = 42
    var v19 string = ((v3 + v3) + v3)
    v6 += -79
    switch 32768 {
    case -9223372036854775808:
        if false {
            for fi20 := 0; fi20 < 1; fi20++ {
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 uint32 = 63
                    switch ("hello" + v3) {
                    case "a":
                        v23 := GenS2{F0: 166, F1: (v3 + v19)}
                        _ = v23
                    case "hello":
                        v5 = append(v5, ("hello" + "foo bar"))
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            v3 = ((v3 + "foo bar") + ((v3 + v19) + v19))
            var v24 uint8 = 232
            _ = v24
        }
    case 128:
        fmt.Println(v3)
    case 1:
        v3 = v19
    default:
        fmt.Println(v6, v3)
    }
    switch 127 {
    case -9223372036854775808:
        v6 -= 0
    case 32767:
        v6 = 238
    default:
        fmt.Println(v3, v19)
    }
    fmt.Println(v19, v3)
    v25 := GenS1{F0: 888, F1: 606}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v26 := GenS3{F0: 255, F1: 255}
    v19 = v3
    var v27 string = (v19 + v19)
    fmt.Println(v3, v3, v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
