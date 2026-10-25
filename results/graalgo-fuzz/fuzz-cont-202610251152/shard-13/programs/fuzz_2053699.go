package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 int32) uint64 {
    var v4 bool = true
    var v5 uint16 = 1
    var v6 string = "a"
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, 9223372036854775807, -247}
        p0 += p0
        v9 := GenS1{F0: 1}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{(-(p0))}
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 string, p1 int8, p2 int) uint64 {
    var v11 uint = 65535
    v11 = (v11 - (^(v11)))
    v12 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return genFn0((p2 >> ((v11) & 63)), 42)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 255
    var v3 uint8 = 0
    var v13 int64 = 42
    var v14 int = 32768
    if (v2 > (v2 ^ v2)) {
        v15 := []int{(v14 >> ((uint(v1)) & 63))}
        _ = v15
    } else {
        var v16 uint64 = genFn0(32768, 1)
        _ = v16
    }
    fmt.Println(v1)
    var v17 string = "foo bar"
    var v18 int32 = -2147483648
    v1 += 0
    var v19 uint = 42
    var v20 int16 = 626
    for fi21 := 0; fi21 < 10; fi21++ {
        switch "hello" {
        case "\n\t":
            v22 := []int{fi21, 716, (fi21 | v14)}
            _ = v22
        default:
            v23 := GenS2{F0: int32(v13)}
            _ = v23
        }
        _ = fi21
    }
    v14 += v14
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 float32 = 1.5
        if false {
            switch v17 {
            case "hello":
                v19 = (v19 << (((v19 / ((v19) | 1))) & 63))
            case "foo bar":
                v26 := []int{fi24, (fi24 >> (((v19 << (((v19 ^ 17)) & 63))) & 63)), (-(v14)), (fi24 ^ (v14 - (v14 - (-(v14)))))}
                _ = v26
            case "\n\t":
                var v27 uint = v19
                _ = v27
            default:
                v28 := []int{fi24}
                _ = v28
            }
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
