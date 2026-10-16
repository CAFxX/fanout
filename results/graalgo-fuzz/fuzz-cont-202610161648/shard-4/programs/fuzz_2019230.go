package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 int32) string {
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 = (p0 / ((-128) | 1))
        var v6 uint32 = 2147483648
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 65535
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    var v1 int8 = 0
    var v2 int32 = 42
    var v3 uint16 = 42
    var v4 int8 = v1
    v2 = (v2 % ((-421) | 1))
    if (!(float64(3.14159) >= 1.5)) {
        switch 127 {
        case 42:
            v4 = (-(v1))
        case 2147483647:
            v1 -= (^(v1))
        case -453:
            v3 = v3
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                v1 = -1
                _ = fi8
            }
        }
    }
    var v9 uint8 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{fi10, (fi10 + (^(fi10)))}
        v2 = (v2 >> ((42) & 63))
        _ = fi10
        _ = v11
    }
    fmt.Println(v4, v1, v3)
    var v12 bool = (false && true)
    fmt.Println(v12, v9, v4)
    v13 := GenS1{F0: v9}
    fmt.Println(v12, v3)
    v14 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
