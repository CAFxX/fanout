package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0() int {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint16 = 0
        fmt.Println(v6, fi5)
        _ = fi5
        _ = v6
    }
    v7 := []int{373, -9223372036854775808}
    var v8 uint16 = 42
    v9 := GenS1{F0: 854}
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 int32, p1 uint, p2 bool) uint {
    v10 := []int{1, 371, 671, 217}
    v11 := GenS2{F0: (p0 << ((p1) & 63)), F1: 2}
    if (140 == -9223372036854775808) {
        var v12 int32 = -139
        switch 127 {
        case 42:
            if p2 {
                var v13 uint32 = 953
                _ = v13
            } else {
                var v14 string = ("a" + "\n\t")
                p2 = true
                if (!(756 > -9223372036854775808)) {
                    var v15 bool = true
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = v15
                } else {
                    p1 += 129
                }
                _ = v14
            }
        case 1:
            switch "a" {
            case "hello":
                var v16 float64 = 3.14159
                _ = v16
            }
        default:
            fmt.Println(p2, p2, p1)
        }
        p1 += p1
        _ = v12
    }
    if false {
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p1 % (((p1 - p1)) | 1))
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]bool)
    var v3 uint = 4294967295
    var v4 int16 = 1000
    v3 += (v3 / (((^(v3))) | 1))
    var v17 bool = (32767 > genFn0())
    var v18 uint16 = 255
    v18 -= v18
    v19 := []int{}
    if true {
        var v20 uint = (^(v3))
        var v21 uint = v3
        _ = v20
        _ = v21
    }
    v1 = v1
    fmt.Println(v4)
    var v22 string = ("\n\t" + ("a" + ""))
    var v23 int32 = -2147483648
    v1 += v1
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
