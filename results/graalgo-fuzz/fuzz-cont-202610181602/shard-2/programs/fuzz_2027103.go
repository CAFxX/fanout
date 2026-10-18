package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 uint32, p2 uint) bool {
    if (!(42 >= 42)) {
        switch "\n\t" {
        case "a":
            for fi4 := 0; fi4 < 9; fi4++ {
                for fi5 := 0; fi5 < 8; fi5++ {
                    p1 = (p1 + (^(p1)))
                    _ = fi5
                }
                v6 := GenS1{F0: "a"}
                for fi7 := 0; fi7 < 7; fi7++ {
                    p0 += p0
                    p0 -= (p0 + p0)
                    _ = fi7
                }
                _ = fi4
                _ = v6
            }
        }
    } else {
        switch 0 {
        case -1:
            p2 = (p2 % ((p2) | 1))
        case 480:
            v8 := []int{2147483647, 0, 1}
            _ = v8
        default:
            var v9 int = 9223372036854775807
            _ = v9
        }
    }
    p1 = p1
    v10 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (42 == -1)
}

func main() {
    var v1 int = 0
    v2 := make(map[int]uint32)
    v3 := []string{}
    var v11 uint8 = 124
    if false {
        v11 -= 13
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 bool = true
            var v14 float64 = float64(1e10)
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(v11, v1)
        v11 = 2
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        if (65535 == 0) {
            v16 := GenS1{F0: "\n\t"}
            var v17 uint = 65535
            _ = v16
            _ = v17
        }
        _ = len(v2)
        var v18 uint = 4294967295
        _ = fi15
        _ = v18
    }
    if true {
        v1 += (v1 & (v1 >> ((693) & 63)))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if (2 < v11) {
            var v19 int16 = 1
            _ = v19
        }
    }
    v20 := GenS1{F0: "hello"}
    v1 = v1
    v21 := GenS2{F0: -9223372036854775808}
    switch v1 {
    case 32767:
        v11 = v11
    }
    v22 := GenS2{F0: 42}
    var v23 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
