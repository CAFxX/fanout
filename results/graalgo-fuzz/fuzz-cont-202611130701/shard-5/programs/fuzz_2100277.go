package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
func genFn0() uint32 {
    v7 := []int{}
    v8 := GenS1{F0: 42, F1: 1, F2: false}
    var v9 int32 = -1
    _ = v7
    _ = v8
    _ = v9
    return 470
}
func genFn1(p0 uint8, p1 int8) string {
    v10 := []int{622, -180, -1, 936}
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = v10
    return "a"
}

func main() {
    var v1 string = "\n\t"
    var v2 uint = 42
    var v3 int8 = 42
    var v4 int64 = 1
    var v5 uint8 = 42
    v6 := make(map[int]int)
    var v11 int16 = 32767
    var v12 int = 32768
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 10; fi14++ {
        v5 -= uint8(v2)
        var v15 int8 = v3
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int64 = (^(v4))
            var v18 int64 = v17
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    var v19 uint32 = 4294967295
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 7; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 uint32 = genFn0()
            fmt.Println(fi22, v1)
            var v24 int = v12
            fmt.Println(v12, v2)
            fmt.Println(v3)
            _ = fi22
            _ = v23
            _ = v24
        }
        v5 -= (v5 / ((v5) | 1))
        switch v12 {
        case 128:
            var v25 int32 = 0
            _ = v25
        case 705:
            var v26 int8 = (v3 >> ((42) & 63))
            _ = v26
        }
        _ = fi21
    }
    var v27 string = "hello"
    switch ((v1 + (v27 + (v27 + v1))) + ((v1 + v27) + v1)) {
    case "\n\t":
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v2)
            v27 = "\n\t"
            if (!false) {
                v29 := GenS1{F0: v20, F1: (-(v11)), F2: (1.5 == (3.14159 / -2.25))}
                v30 := v6[int(v20)]
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(fi28)
                    _ = fi31
                }
                fmt.Println(v5, v20, v1)
                fmt.Println(v3)
            }
            _ = fi28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
