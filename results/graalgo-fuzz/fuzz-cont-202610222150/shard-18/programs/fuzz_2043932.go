package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string; F2 bool }
func genFn0() int8 {
    var v6 uint32 = 2147483648
    v7 := []int{9223372036854775807, -785, -691, 779}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return -112
}
func genFn1() int16 {
    var v10 int8 = 127
    for fi11 := 0; fi11 < 9; fi11++ {
        switch "a" {
        case "foo bar":
            switch 128 {
            case 4294967295:
                for fi12 := 0; fi12 < 4; fi12++ {
                    v10 = 1
                    v13 := GenS1{F0: -2147483648}
                    _ = fi12
                    _ = v13
                }
            case 128:
                var v14 uint32 = 134
                _ = v14
            }
        case "hello":
            v15 := []int{32768}
            _ = v15
        case "\n\t":
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := []int{(fi16 + 127)}
                fmt.Println(fi16, fi11, fi16)
                v10 += genFn0()
                _ = fi16
                _ = v17
            }
        }
        fmt.Println(fi11)
        _ = fi11
    }
    _ = v10
    return 884
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]string)
    var v3 uint16 = 0
    v4 := []int{}
    v5 := make(map[string]string)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v1 += 2
        if (564 != 2147483647) {
            v1 -= 166
        }
    }
    fmt.Println(v3)
    v18 := v2[42]
    v18 = v18
    v19 := GenS1{F0: 1}
    var v20 int8 = 100
    v21 := GenS2{F0: 1, F1: v18, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
