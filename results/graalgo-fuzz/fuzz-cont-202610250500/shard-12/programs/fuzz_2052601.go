package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 string) int64 {
    v4 := []int{-411, 4294967295, -287, 1}
    var v5 uint16 = 94
    _ = p0
    _ = v4
    _ = v5
    return 502
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := make(map[int]int)
    var v6 string = "a"
    v6 = "a"
    for fi7 := 0; fi7 < 1; fi7++ {
        if (-17 == -659) {
            v6 = v6
        }
        _ = fi7
    }
    v2 = v2
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: (fi9 << ((2) & 63))}
            _ = fi9
            _ = v10
        }
        fmt.Println(v6)
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: (^(fi11))}
            var v13 uint16 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    if (v2 && v2) {
        if (1 < 42) {
            v2 = true
            _ = len(v3)
            fmt.Println(v6, v2)
        } else {
            var v14 int64 = 0
            _ = v14
        }
    }
    var v15 uint32 = 586
    var v16 string = v6
    switch "hello" {
    case "a":
        if (-1 <= 0) {
            fmt.Println(v15)
        } else {
            v1[-2147483648] = (v15 * 2)
            var v17 uint16 = 1
            _ = v17
        }
    case "\n\t":
        v18 := v3[-710]
        _ = v18
    case "foo bar":
        if v2 {
            var v19 bool = v2
            var v20 bool = v19
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: -319}
            if (int8(v15) < -127) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(v15)
                    var v23 int = (fi22 >> ((362) & 63))
                    fmt.Println(v6, v23, v16)
                    _ = fi22
                    _ = v23
                }
            } else {
                for fi24 := 0; fi24 < 4; fi24++ {
                    v6 = ((("hello" + "hello") + "foo bar") + "\n\t")
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v26 := GenS1{F0: fi25}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
            }
            _ = v21
        }
    }
    v27 := []int{2147483647}
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int64 = -9223372036854775808
        v30 := GenS1{F0: fi28}
        var v31 bool = (!((86 > -99) && v2))
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v6
}
