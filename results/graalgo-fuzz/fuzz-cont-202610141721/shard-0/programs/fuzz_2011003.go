package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]string)
    var v4 float32 = 1.5
    var v5 bool = (true && false)
    var v6 int64 = 114
    v7 := GenS1{F0: (42 >= 32767)}
    v8 := GenS1{F0: (-32768 <= -193)}
    v9 := GenS1{F0: false}
    fmt.Println(v4, v1)
    var v10 uint64 = 597
    v2 = (v2 & v10)
    if (!v5) {
        fmt.Println(v1, v4)
        v3["\n\t"] = "\n\t"
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := []int{}
            _ = fi11
            _ = v12
        }
    }
    switch "" {
    case "":
        if true {
            if ((1 > 0) && v5) {
                v13 := GenS1{F0: true}
                switch "" {
                case "hello":
                    var v14 uint8 = (^(v1))
                    _ = v14
                case "a":
                    var v15 uint16 = 65535
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 int8 = 100
                        _ = fi16
                        _ = v17
                    }
                }
                switch "" {
                case "hello":
                    delete(v3, "foo bar")
                case "foo bar":
                    var v18 string = "\n\t"
                    _ = v18
                default:
                    switch -9223372036854775808 {
                    case -974:
                        v19 := []int{-2147483648, -250}
                        _ = v19
                    case 2147483647:
                        var v20 int64 = (v6 - (v6 % (((v6 >> ((388) & 63))) | 1)))
                        _ = v20
                    default:
                        v21 := []int{}
                        _ = v21
                    }
                }
                _ = v13
            }
            var v22 uint8 = v1
            _ = v22
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
