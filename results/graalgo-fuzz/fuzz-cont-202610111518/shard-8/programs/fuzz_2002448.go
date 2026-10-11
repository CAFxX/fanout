package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
func genFn0(p0 int8, p1 uint32, p2 int8) int32 {
    v7 := GenS1{F0: 2, F1: 9223372036854775807}
    v8 := GenS1{F0: 35, F1: 4294967295}
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -633
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := []int{}
    var v4 int8 = 100
    var v5 int8 = 1
    var v6 int8 = (v4 & (v5 / ((v5) | 1)))
    if true {
        v4 = v6
    } else {
        v6 = -40
        for fi10 := 0; fi10 < 11; fi10++ {
            fmt.Println(v5)
            v4 = 127
            v5 -= v4
            _ = fi10
        }
    }
    switch "a" {
    case "foo bar":
        v6 -= 0
    case "":
        for fi11 := 0; fi11 < 2; fi11++ {
            if (515 >= 2) {
                v12 := GenS2{F0: (fi11 - (fi11 >> ((4294967295) & 63))), F1: 366, F2: 2}
                if (genFn0(v4, 0, -67) != 1) {
                    v3 = append(v3, 398)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v13 int64 = 9223372036854775807
                    _ = v13
                }
                if len(v3) > 0 {
                    v3[0] = 611
                }
                _ = v12
            } else {
                v14 := GenS2{F0: fi11, F1: 701, F2: 170}
                switch 32768 {
                case 4294967295:
                    var v15 int8 = (v5 | 56)
                    _ = v15
                case -1:
                    fmt.Println(v5)
                case 128:
                    v16 := GenS1{F0: 2, F1: 488}
                    _ = v16
                }
                _ = v14
            }
            fmt.Println(v5)
            _ = fi11
        }
    case "hello":
        fmt.Println(v6, v6, v5)
    }
    v17 := GenS1{F0: 638, F1: -2147483648}
    if false {
        switch 0 {
        case 1:
            v18 := GenS1{F0: 254, F1: 51}
            _ = v18
        case 32767:
            var v19 uint16 = 0
            _ = v19
        default:
            var v20 int16 = 0
            _ = v20
        }
    }
    fmt.Println(v4, v5, v5)
    if ((v5 > v4) || false) {
        v5 += v4
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS2{F0: fi21, F1: -298, F2: 255}
            v4 = v4
            _ = fi21
            _ = v22
        }
    }
    v5 -= v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "" {
    case "":
        fmt.Println(v4, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
