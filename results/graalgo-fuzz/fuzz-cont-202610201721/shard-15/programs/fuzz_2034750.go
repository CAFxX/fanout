package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int16 }
func genFn0(p0 int16) int32 {
    var v4 int = -2147483648
    var v5 int16 = p0
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 string = ""
    var v6 uint = 4294967295
    if v1 {
        v7 := GenS1{F0: 423}
        v6 = (v6 + 0)
        _ = v7
    } else {
        v1 = v1
        if (((281 > 587) || false) || v1) {
            if v1 {
                switch "hello" {
                case "":
                    switch v3 {
                    case "foo bar":
                        v8 := []int{-798, 128}
                        _ = v8
                    case "\n\t":
                        fmt.Println(v1)
                    default:
                        for fi9 := 0; fi9 < 9; fi9++ {
                            fmt.Println(fi9, fi9, v1)
                            _ = fi9
                        }
                    }
                case "a":
                    var v10 bool = v1
                    _ = v10
                case "hello":
                    var v11 int16 = 1000
                    _ = v11
                }
                for fi12 := 0; fi12 < 17; fi12++ {
                    fmt.Println(v1)
                    _ = fi12
                }
            } else {
                fmt.Println(v3, v6)
            }
            v2[42] = 4294967295
        }
    }
    var v13 int8 = -1
    fmt.Println(v3, v6)
    var v14 int = 2147483647
    v15 := GenS2{F0: 1}
    v16 := GenS3{F0: 42, F1: 35, F2: 32767}
    var v17 uint8 = 5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
