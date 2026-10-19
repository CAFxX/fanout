package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 uint8, p2 int) bool {
    switch 4294967295 {
    case 32767:
        var v5 float32 = (-2.25 / (0.1 + 1e10))
        _ = v5
    case -2147483648:
        var v6 uint64 = 621
        _ = v6
    case -471:
        var v7 string = "foo bar"
        _ = v7
    default:
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p0)
            switch "" {
            case "hello":
                switch (fi8 & 42) {
                case -362:
                    for fi9 := 0; fi9 < 9; fi9++ {
                        p1 = (p1 << ((2) & 63))
                        fmt.Println(fi8, fi9)
                        switch "\n\t" {
                        case "a":
                            if (!true) {
                                var v10 string = "\n\t"
                                _ = v10
                            }
                        }
                        _ = fi9
                    }
                }
            case "a":
                fmt.Println(fi8)
            }
            _ = fi8
        }
    }
    var v11 uint8 = (p1 % ((p1) | 1))
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint32 = 42
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 808}
    var v15 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]bool)
    var v3 uint = 2
    v4 := make(map[string]string)
    var v16 int64 = 1
    v17 := []int{996}
    v18 := GenS1{F0: 841}
    v16 -= (v16 << ((v3) & 63))
    var v19 int16 = 32767
    switch 4294967295 {
    case 881:
        var v20 int64 = v16
        _ = v20
    case 263:
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint16 = 789
            _ = fi21
            _ = v22
        }
    }
    var v23 uint64 = 743
    v24 := GenS2{F0: v3}
    var v25 int32 = v1
    var v26 int8 = 100
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            v25 += -467
            _ = fi28
        }
        v29 := GenS1{F0: 177}
        _ = fi27
        _ = v29
    }
    _ = len(v4)
    var v30 float32 = (float32(((1.5 / ((((0.0 / 0.1) - ((((0.1 * 0.0) + -2.25) - -2.25) + 1.5)) - 1.5) + 1e10)) * (0.1 - (1e10 * -2.25)))) / 3.14159)
    _ = len(v17)
    v1 = -636
    switch "foo bar" {
    case "a":
        for fi31 := 0; fi31 < 7; fi31++ {
            v4["a"] = ""
            _ = fi31
        }
    case "foo bar":
        v32 := v2[9223372036854775807]
        _ = v32
    case "\n\t":
        var v33 int32 = (v1 ^ v25)
        _ = v33
    default:
        var v34 uint64 = 18446744073709551615
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
}
