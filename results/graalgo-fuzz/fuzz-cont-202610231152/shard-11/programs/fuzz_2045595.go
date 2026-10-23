package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool) uint64 {
    v7 := []int{4294967295, 127, 128}
    for fi8 := 0; fi8 < 17; fi8++ {
        p0 = (0 == 65535)
        if (p0 && (((!p0) || (2 > 18446744073709551615)) && (2 != 0))) {
            var v9 uint64 = 0
            switch (fi8 - fi8) {
            case 707:
                var v10 int16 = 986
                _ = v10
            case 127:
                var v11 int8 = -128
                _ = v11
            }
            _ = v9
        } else {
            v7 = append(v7, fi8)
            p0 = p0
        }
        p0 = p0
        _ = fi8
    }
    var v12 uint64 = 631
    _ = p0
    _ = v12
    _ = v7
    return (v12 << ((42) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 402
    var v3 uint = 42
    var v4 int32 = (v2 | -1)
    v5 := make(map[string]int)
    v6 := []uint32{2147483648, 698, 4294967295, 2}
    switch -293 {
    case 32767:
        fmt.Println(v3)
    case -396:
        var v13 float32 = (1e10 + 1e10)
        _ = v13
    default:
        var v14 int32 = (^(v2))
        _ = v14
    }
    v4 = v2
    v15 := v5["foo bar"]
    var v16 int32 = v4
    var v17 uint64 = 18446744073709551615
    var v18 string = "hello"
    v2 = -2147483648
    v6 = append(v6, 2147483647)
    var v19 uint32 = 2147483648
    v20 := GenS1{F0: -490}
    if (2 == 137) {
        switch v18 {
        case "hello":
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 int = (fi21 / ((v15) | 1))
                _ = fi21
                _ = v22
            }
        case "foo bar":
            var v23 uint64 = 42
            _ = v23
        }
        if len(v6) > 0 {
            v6[0] = (^(v19))
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint64 = v17
            v19 = (^(v19))
            _ = fi24
            _ = v25
        }
    } else {
        var v26 float32 = -2.25
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 20; fi28++ {
                switch (fi27 | fi27) {
                case 32767:
                    var v29 uint16 = 32767
                    _ = v29
                }
                _ = fi28
            }
            var v30 int = (fi27 >> ((v3) & 63))
            v3 = (^(v3))
            _ = fi27
            _ = v30
        }
        v2 += v16
        _ = v26
    }
    fmt.Println(v4, v17)
    if (((-2.25 == (3.14159 / -2.25)) && (!(int16(v3) <= 0))) && false) {
        var v31 uint = v3
        var v32 uint32 = (v19 * v19)
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
