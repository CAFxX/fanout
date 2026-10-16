package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ("\n\t" + "hello")
    var v3 string = "a"
    v4 := make(map[string]int)
    var v5 int = 257
    var v6 int32 = -883
    v7 := GenS1{F0: 2, F1: 9223372036854775807}
    if true {
        var v8 string = "hello"
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                fmt.Println(fi9)
                v11 := GenS2{F0: true, F1: 32767, F2: fi10}
                fmt.Println(v3, v6, v2)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    }
    var v12 bool = false
    v1[v5] = v5
    switch -1 {
    case -227:
        v3 = v3
    case -1:
        v13 := []int{(v5 % ((v5) | 1))}
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        switch 42 {
        case 42:
            v1[(fi14 / ((0) | 1))] = -394
        default:
            for fi15 := 0; fi15 < 9; fi15++ {
                v12 = v12
                var v16 int = 4294967295
                v17 := GenS3{F0: v6}
                v5 = (v5 / ((v5) | 1))
                fmt.Println(v16, v3)
                switch -1 {
                case -9223372036854775808:
                    if (-233 >= (v16 % ((v16) | 1))) {
                        v18 := []int{32768, (v16 % ((v5) | 1))}
                        _ = v18
                    }
                default:
                    v3 = (v2 + "a")
                }
                v19 := GenS2{F0: (100 == 26), F1: 65535, F2: (fi15 + fi15)}
                _ = fi15
                _ = v16
                _ = v17
                _ = v19
            }
        }
        _ = fi14
    }
    v20 := []int{v5, v5}
    var v21 float32 = 1e10
    var v22 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
