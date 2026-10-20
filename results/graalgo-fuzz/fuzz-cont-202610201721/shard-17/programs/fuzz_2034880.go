package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
func genFn0(p0 int64, p1 uint16) bool {
    v6 := GenS1{F0: 529, F1: "hello", F2: (p0 & (p0 * p0))}
    for fi7 := 0; fi7 < 3; fi7++ {
        p1 = (p1 - p1)
        v8 := []int{fi7, 2147483647, fi7, fi7}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return (true && true)
}
func genFn1(p0 int64, p1 int8, p2 uint) uint64 {
    v9 := GenS1{F0: 42, F1: (("\n\t" + "\n\t") + ("hello" + "a")), F2: 42}
    var v10 float32 = 0.0
    p2 = (p2 << ((p2) & 63))
    v11 := []int{32768}
    p2 += (p2 - 255)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 754
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    v3 := []uint32{4294967295, 4294967295, 633}
    v4 := make(map[int]int)
    var v5 int32 = -1
    var v12 uint64 = 591
    v5 = v5
    var v13 float64 = 1e10
    v1 = ""
    if true {
        if (-195 <= -600) {
            switch 32768 {
            case 1:
                v5 -= (^(v5))
            default:
                v14 := GenS1{F0: -32768, F1: (v1 + ("a" + (v1 + v1))), F2: -9223372036854775808}
                _ = v14
            }
            switch 0 {
            case 32768:
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 int64 = -334
                    fmt.Println(v12)
                    _ = fi15
                    _ = v16
                }
            }
        }
        fmt.Println(v5)
    }
    if (-2147483648 >= 128) {
        var v17 int16 = -570
        _ = v17
    }
    v1 = v1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int8 = 0
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 32767, F1: (v1 + (v1 + v1)), F2: -1}
    _ = len(v2)
    var v21 string = "\n\t"
    var v22 float64 = 3.14159
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 int16 = -32768
        _ = fi23
        _ = v24
    }
    switch v1 {
    case "a":
        var v25 float32 = 0.1
        _ = v25
    case "hello":
        var v26 uint64 = 2
        _ = v26
    default:
        var v27 int32 = -2147483648
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
