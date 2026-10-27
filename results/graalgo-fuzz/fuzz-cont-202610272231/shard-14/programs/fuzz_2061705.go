package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 int }
func genFn0() uint16 {
    v5 := []int{-1, 42}
    var v6 float32 = 3.14159
    v6 += 3.14159
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 2147483647
    v3 := make(map[string]int)
    var v4 uint64 = 18446744073709551615
    v2 = v2
    if (1 == 1) {
        fmt.Println(v4)
    } else {
        var v7 uint16 = 32767
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    var v9 uint64 = v4
    var v10 int32 = v2
    if false {
        v10 += (v10 * v2)
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: fi11, F1: (^(v9)), F2: "foo bar"}
                switch "hello" {
                case "":
                    if (true && (-1 < -128)) {
                        var v14 uint64 = v4
                        v15 := GenS1{F0: -9223372036854775808, F1: v14, F2: ""}
                        v16 := []int{fi11, (^(fi11)), -1, -2147483648}
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint32 = 65535
                        var v18 string = "hello"
                        delete(v1, (v17 - v17))
                        _ = v17
                        _ = v18
                    }
                }
                _ = fi12
                _ = v13
            }
            v9 = v4
            v19 := GenS2{F0: -2147483648}
            _ = fi11
            _ = v19
        }
    }
    var v20 uint8 = 127
    if true {
        fmt.Println(v4)
        v21 := []int{}
        var v22 int16 = -32768
        _ = v21
        _ = v22
    } else {
        v20 -= (v20 | 42)
        fmt.Println(v9, v9)
        fmt.Println(v20)
    }
    v10 = int32(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
