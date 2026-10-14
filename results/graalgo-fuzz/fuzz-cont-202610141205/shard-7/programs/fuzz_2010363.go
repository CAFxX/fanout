package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 int8 = 127
    var v4 string = "a"
    var v5 string = v4
    var v6 int = int(v1)
    switch v5 {
    case "":
        fmt.Println(v1)
    default:
        var v7 int = (v6 >> ((65535) & 63))
        _ = v7
    }
    delete(v2, 768)
    var v8 float64 = 3.14159
    v4 = "a"
    for fi9 := 0; fi9 < 10; fi9++ {
        if true {
            v6 += v6
        }
        v10 := []int{-726, (fi9 & (^(v6)))}
        for fi11 := 0; fi11 < 1; fi11++ {
            if ((float32((0.1 / ((v8 / ((3.14159 + 0.1) / v8)) + 1e10))) != (0.0 / float32((v8 + v8)))) || (2 <= 948)) {
                v8 = float64(0.0)
                var v12 bool = false
                _ = v12
            } else {
                switch v4 {
                case "hello":
                    v13 := GenS1{F0: 0}
                    _ = v13
                case "":
                    v14 := GenS2{F0: fi11}
                    var v15 uint16 = 209
                    if true {
                        var v16 uint64 = 278
                        v4 = v5
                        v16 = v16
                        _ = v16
                    } else {
                        if (false && (2147483648 == 0)) {
                            v3 = (v3 ^ v3)
                            v1 = -93
                        }
                        var v17 uint = 18446744073709551615
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        _ = v17
                    }
                    _ = v14
                    _ = v15
                case "a":
                    var v18 bool = (true || (!false))
                    _ = v18
                default:
                    var v19 bool = false
                    _ = v19
                }
                fmt.Println(v5, v3, fi11)
                fmt.Println(v1)
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v20 := v2[42]
    v4 = v4
    switch v5 {
    case "hello":
        if true {
            var v21 uint = 758
            v22 := GenS2{F0: v20}
            var v23 string = (v5 + ("\n\t" + v4))
            _ = v21
            _ = v22
            _ = v23
        } else {
            var v24 bool = (42 > 36)
            v1 = (v3 % ((127) | 1))
            _ = v24
        }
    case "foo bar":
        var v25 uint = 1
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
