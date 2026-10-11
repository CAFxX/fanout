package main

import (
    "fmt"
)

func genFn0(p0 int) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = (fi7 << ((4294967295) & 63))
        _ = fi7
    }
    var v8 int8 = 95
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{(fi9 << ((2) & 63)), fi9}
        v8 = 14
        fmt.Println(fi9, fi9)
        _ = fi9
        _ = v10
    }
    p0 = p0
    _ = p0
    _ = v8
    return 80
}

func main() {
    var v1 bool = (false || false)
    v2 := make(map[int]string)
    var v3 string = "\n\t"
    v4 := []string{"foo bar", "foo bar", v3, (((v3 + "hello") + (v3 + "foo bar")) + v3)}
    v5 := make(map[int]uint32)
    v6 := []uint32{1, 2147483647}
    var v11 uint16 = 42
    switch "hello" {
    case "a":
        if true {
            fmt.Println(v1)
            v3 = (v3 + v3)
        } else {
            var v12 uint8 = 188
            fmt.Println(v11, v3, v12)
            v12 = v12
            _ = v12
        }
    }
    var v13 uint8 = 127
    fmt.Println(v13, v1, v13)
    v1 = (uint16(v13) >= (v11 / ((v11) | 1)))
    if v1 {
        v1 = (-732 >= -9223372036854775808)
        var v14 float64 = float64(((1.5 * 1e10) / 3.14159))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v13)
        _ = fi15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v4 = append(v4, "a")
        _ = len(v5)
        v11 += (v11 << ((315) & 63))
        _ = fi16
    }
    if true {
        fmt.Println(v1, v3, v13)
    }
    v11 += genFn0(360)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
