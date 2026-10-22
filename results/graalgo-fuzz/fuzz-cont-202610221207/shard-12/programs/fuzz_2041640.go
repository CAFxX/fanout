package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
func genFn0(p0 uint64, p1 bool, p2 int32) uint32 {
    fmt.Println(p1, p2, p1)
    p2 += (p2 >> ((875) & 63))
    for fi6 := 0; fi6 < 10; fi6++ {
        if p1 {
            p1 = (!(1 < 0))
            p2 -= 2147483647
        } else {
            v7 := []int{(^(fi6))}
            _ = v7
        }
        _ = fi6
    }
    var v8 uint = 0
    var v9 uint = v8
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 931
}
func genFn1(p0 uint8, p1 int, p2 uint) int8 {
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(p1, fi11, p0)
            _ = fi11
        }
        _ = fi10
    }
    if (65535 >= 461) {
        var v12 uint8 = (p0 & 255)
        v13 := []int{(p1 ^ (^(p1))), p1, 2147483647}
        _ = v12
        _ = v13
    } else {
        var v14 float64 = 3.14159
        fmt.Println(p1, p0)
        _ = v14
    }
    v15 := []int{}
    switch (p1 % (((p1 << ((757) & 63))) | 1)) {
    case 523:
        if true {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 uint = 2
                fmt.Println(p1, p2)
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS1{F0: p0, F1: fi16}
                    _ = fi18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            fmt.Println(p1, p1, p2)
        }
    }
    p2 += (p2 * (p2 * p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -107
}

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]bool)
    var v3 float32 = (1.5 * 1.5)
    var v4 uint8 = 255
    v5 := make(map[string]int)
    var v20 uint64 = 2
    var v21 float32 = v3
    fmt.Println(v21)
    var v22 int = -9223372036854775808
    var v23 int16 = 485
    var v24 uint16 = 1
    var v25 uint16 = v24
    if true {
        if (true || ((false && ("" >= ((("" + "") + "\n\t") + "foo bar"))) && (!true))) {
            v3 = 0.0
        }
    } else {
        fmt.Println(v21)
    }
    var v26 uint64 = v20
    var v27 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
