package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 string }
type GenS4 struct { F0 bool; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 uint8; F2 int8 }

func main() {
    var v1 uint64 = 146
    var v2 uint8 = 42
    var v3 uint = 167
    v4 := make(map[int]int)
    var v5 uint = v3
    v6 := []int{}
    v7 := GenS1{F0: (v5 & v5), F1: -502, F2: 0}
    fmt.Println(v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 := GenS2{F0: true, F1: ("" + "hello"), F2: 2}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 127 {
    case 42:
        v5 = v3
    case 32768:
        v9 := GenS1{F0: 65535, F1: 32768, F2: 1}
        _ = v9
    }
    if true {
        if true {
            fmt.Println(v3, v2, v5)
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS3{F0: -123, F1: v2, F2: "\n\t"}
                _ = fi10
                _ = v11
            }
            if false {
                if false {
                    v12 := v4[-9223372036854775808]
                    v5 -= (v5 / (((^(v5))) | 1))
                    _ = v12
                } else {
                    switch -450 {
                    case -699:
                        if (v1 < (v1 - 2)) {
                            v13 := GenS2{F0: true, F1: (("a" + "a") + (("hello" + "foo bar") + "a")), F2: 334}
                            _ = v13
                        }
                    case 92:
                        fmt.Println(v2)
                    default:
                        v14 := GenS4{F0: (v1 != (^(v1))), F1: 796, F2: 601}
                        _ = v14
                    }
                    fmt.Println(v2, v2, v3)
                }
            } else {
                switch 688 {
                case -1:
                    v15 := GenS5{F0: false, F1: (v2 >> ((v5) & 63)), F2: 111}
                    _ = v15
                case -9223372036854775808:
                    v2 = (^(v2))
                case 9223372036854775807:
                    var v16 int64 = 0
                    _ = v16
                default:
                    var v17 float64 = 3.14159
                    _ = v17
                }
            }
        } else {
            var v18 int64 = 42
            fmt.Println(v1)
            _ = v18
        }
    }
    v19 := GenS2{F0: false, F1: "foo bar", F2: 115}
    v6 = append(v6, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
