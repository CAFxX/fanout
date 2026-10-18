package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 int16) int16 {
    switch "a" {
    case "hello":
        var v6 int8 = 1
        _ = v6
    case "":
        fmt.Println(p1, p1)
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 string, p1 uint64, p2 uint32) uint {
    v7 := []int{42}
    p2 = p2
    p1 = 935
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 333
}

func main() {
    var v1 int32 = 42
    v2 := []string{("\n\t" + "\n\t")}
    v3 := []int{2147483647, 0, 127}
    var v4 uint8 = 146
    var v5 uint32 = 65535
    if true {
        v1 = v1
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                v1 = v1
                var v10 uint = 42
                v4 += (v4 / (((v4 / ((128) | 1))) | 1))
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    } else {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        v5 -= (v5 + 65535)
    }
    v1 = v1
    v1 -= -2147483648
    var v11 uint32 = (v5 << ((genFn1("", 4294967295, 408)) & 63))
    var v12 int64 = 0
    v1 += v1
    var v13 uint16 = 399
    fmt.Println(v12, v4, v1)
    var v14 uint16 = 255
    if len(v3) > 0 {
        v3[0] = int(v11)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 uint8 = (v4 * 69)
    var v16 int64 = v12
    if false {
        fmt.Println(v1)
    }
    var v17 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
