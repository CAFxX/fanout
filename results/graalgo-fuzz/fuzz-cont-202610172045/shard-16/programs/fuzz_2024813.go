package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 uint8 = 128
    v4 := make(map[string]bool)
    v5 := []string{""}
    v6 := GenS1{F0: 2, F1: 2}
    v7 := v2[2]
    fmt.Println(v3, v3)
    v3 = (v3 & v3)
    if false {
        var v8 string = v7
        v7 = "hello"
        _ = v8
    } else {
        if true {
            var v9 uint32 = 294
            v10 := GenS2{F0: "hello", F1: 32767}
            v7 = (v7 + "")
            _ = v10
            _ = v9
        }
        v3 = v3
        fmt.Println(v7, v7)
    }
    v7 = ""
    delete(v4, v7)
    for fi11 := 0; fi11 < 9; fi11++ {
        if true {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 string = v7
                fmt.Println(v3)
                fmt.Println(v7, v13)
                v14 := GenS2{F0: "hello", F1: 244}
                switch (fi12 * fi11) {
                case 32767:
                    var v15 uint8 = 255
                    _ = v15
                case 42:
                    switch v13 {
                    case "foo bar":
                        var v16 string = (v7 + (((((v13 + v13) + ("foo bar" + "")) + "") + "hello") + v7))
                        _ = v16
                    case "hello":
                        v13 = ((v13 + "hello") + "\n\t")
                    case "\n\t":
                        fmt.Println(v3, fi12, v7)
                    default:
                        fmt.Println(v3, fi11, fi12)
                    }
                default:
                    var v17 uint8 = (v3 ^ 2)
                    _ = v17
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 float32 = float32(3.14159)
                v20 := GenS2{F0: v7, F1: 1000}
                var v21 int64 = -1
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        _ = fi11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
