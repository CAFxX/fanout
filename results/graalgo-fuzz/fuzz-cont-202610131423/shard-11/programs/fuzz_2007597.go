package main

import (
    "fmt"
)

func genFn0(p0 uint, p1 int16) int {
    v5 := []int{0, -9223372036854775808, 891, -984}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 657
    var v3 uint8 = 51
    var v4 string = (("hello" + "a") + "")
    var v6 bool = false
    if v6 {
        if (0 < 32767) {
            fmt.Println(v2, v2)
        }
        switch 128 {
        case 1:
            fmt.Println(v3, v3, v6)
        }
    }
    v6 = false
    var v7 uint32 = 2147483647
    for fi8 := 0; fi8 < 1; fi8++ {
        v1[(fi8 >> ((4294967295) & 63))] = (fi8 + 7)
        _ = fi8
    }
    v1[genFn0(65535, 0)] = genFn0(293, 1000)
    v6 = v6
    v4 = ("" + "\n\t")
    var v9 uint16 = 65535
    v6 = (v4 <= v4)
    v9 += v9
    var v10 float32 = -2.25
    v9 = 560
    v3 += 42
    if true {
        if false {
            v11 := []int{32768, 42}
            fmt.Println(v10)
            v3 = (^(v3))
            _ = v11
        }
        fmt.Println(v9, v9, v2)
        v12 := []int{-2147483648, 42}
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
