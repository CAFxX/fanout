package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 int) int64 {
    if false {
        p1 = p1
        p1 = (p1 / (((^(p1))) | 1))
    }
    if ((42 <= 127) && true) {
        v7 := []int{-700, (p1 / ((p1) | 1)), (p1 / ((9223372036854775807) | 1))}
        p1 = (p1 % ((p1) | 1))
        _ = v7
    } else {
        var v8 int32 = 1
        v9 := GenS1{F0: "a"}
        var v10 float64 = 1e10
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := []int{}
    v12 := GenS1{F0: ("hello" + "hello")}
    v11 = append(v11, p1)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -29
    var v3 int32 = 2147483647
    var v4 uint16 = 1
    var v5 float64 = -2.25
    var v6 int32 = (v3 >> ((933) & 63))
    v4 -= (^(v4))
    v3 = (v3 << ((236) & 63))
    v6 -= (v6 | v3)
    var v13 uint16 = 1
    if true {
        fmt.Println(v13)
        var v14 uint32 = 65535
        _ = v14
    }
    var v15 uint8 = 228
    v16 := GenS1{F0: ("hello" + ("\n\t" + ("\n\t" + "\n\t")))}
    for fi17 := 0; fi17 < 13; fi17++ {
        fmt.Println(v4)
        if (34 < 42) {
            fmt.Println(v5, v4, v6)
            _ = len(v1)
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 uint8 = (v15 / ((v15) | 1))
                _ = fi18
                _ = v19
            }
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        fmt.Println(v3, v2)
        var v21 uint32 = 2147483647
        v5 += 1.5
        _ = fi20
        _ = v21
    }
    v15 = 210
    var v22 string = "hello"
    _ = len(v1)
    v4 -= 32767
    v15 += v15
    for fi23 := 0; fi23 < 19; fi23++ {
        v4 += 789
        _ = fi23
    }
    var v24 uint32 = 926
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
