package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 int64; F2 int8 }
func genFn0(p0 string, p1 bool) uint64 {
    p0 = (("foo bar" + ("a" + p0)) + "\n\t")
    var v5 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v5
    return (v5 * 4294967295)
}

func main() {
    var v1 float32 = -2.25
    var v2 uint8 = 128
    var v3 int32 = 1
    v4 := make(map[int]uint32)
    if false {
        for fi6 := 0; fi6 < 20; fi6++ {
            v7 := GenS1{F0: true, F1: 42, F2: fi6}
            for fi8 := 0; fi8 < 3; fi8++ {
                fmt.Println(fi8)
                v9 := []int{fi8, fi8, 4294967295, fi6}
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    v10 := GenS2{F0: v3, F1: 9223372036854775807, F2: 1}
    var v11 string = "foo bar"
    v12 := v4[2147483647]
    v2 = 1
    switch 924 {
    case 732:
        v13 := []int{}
        _ = v13
    case -857:
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 uint8 = (v2 & v2)
            _ = fi14
            _ = v15
        }
    case -704:
        var v16 int64 = 1
        _ = v16
    default:
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 uint = 2
            v4[9223372036854775807] = (v12 % (((^(v12))) | 1))
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v11, v11)
    v1 += float32((3.14159 - (3.14159 - 3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
