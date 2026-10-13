package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 bool) uint64 {
    var v4 int8 = 96
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{-483, fi5}
        p0 = (!p0)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{fi7}
            _ = fi8
            _ = v9
        }
        v4 += (^(v4))
        var v10 float32 = (0.0 - float32(float64((3.14159 + 1.5))))
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: 42, F1: 1}
        var v13 float64 = float64(3.14159)
        _ = fi11
        _ = v12
        _ = v13
    }
    p0 = (42 >= 65535)
    _ = p0
    _ = v4
    return 4294967295
}

func main() {
    var v1 float32 = (1.5 / 3.14159)
    var v2 uint16 = 805
    var v3 uint = 758
    var v14 float32 = (float32(3.14159) + v1)
    var v15 int16 = -1
    var v16 uint = v3
    if true {
        switch "\n\t" {
        case "a":
            v2 -= 255
        default:
            var v17 bool = (v3 <= 2)
            _ = v17
        }
    }
    if false {
        v18 := []int{2147483647, -2147483648, 4294967295, 128}
        _ = v18
    }
    if false {
        if false {
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 string = "a"
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: -1, F1: v15}
            var v22 bool = true
            _ = v21
            _ = v22
        }
        v23 := []int{-352, 529, 4294967295, 32768}
        _ = v23
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2, v3, v2)
    case "hello":
        fmt.Println(v2, v2, v15)
    case "\n\t":
        var v24 float64 = 1e10
        _ = v24
    default:
        v1 = v1
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := []int{fi25, fi25, fi25, fi25}
        fmt.Println(fi25, v3, v14)
        fmt.Println(v2, v1, v14)
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 4294967295, F1: 2}
    v14 = float32(0.0)
    switch 147 {
    case -1:
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 float32 = (v14 - ((v14 + v1) + v14))
            _ = fi28
            _ = v29
        }
    case 922:
        v3 += 18446744073709551615
    }
    fmt.Println(v2, v3, v15)
    v30 := []int{42}
    if false {
        fmt.Println(v16, v14)
        v30 = append(v30, -9223372036854775808)
    }
    var v31 uint16 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
