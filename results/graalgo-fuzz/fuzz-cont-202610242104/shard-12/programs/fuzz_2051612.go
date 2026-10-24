package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "hello" {
        case "foo bar":
            fmt.Println(fi7)
        default:
            v8 := []int{387, (fi7 - (fi7 ^ fi7)), fi7, 974}
            _ = v8
        }
        var v9 bool = (-2147483648 <= (-(fi7)))
        _ = fi7
        _ = v9
    }
    var v10 int64 = 42
    _ = v10
    return 1
}

func main() {
    var v1 uint16 = 371
    var v2 uint16 = v1
    var v3 string = "foo bar"
    v4 := []uint32{66, 42, 42, 266}
    v5 := make(map[string]bool)
    v6 := make(map[int]int)
    fmt.Println(v2, v1)
    var v11 int8 = -1
    var v12 float32 = 0.0
    fmt.Println(v2, v11)
    v5[v3] = (v1 <= (^(v1)))
    var v13 uint8 = genFn0()
    switch "foo bar" {
    case "foo bar":
        var v14 float64 = -2.25
        _ = v14
    case "":
        var v15 int32 = 983
        _ = v15
    }
    switch (v3 + "hello") {
    case "a":
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: v3, F1: 0, F2: ((v3 + v3) + "\n\t")}
            _ = fi16
            _ = v17
        }
    case "":
        v13 += (v13 ^ 0)
    case "hello":
        for fi18 := 0; fi18 < 5; fi18++ {
            v13 = 1
            fmt.Println(v11, v13)
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint32 = 323
                fmt.Println(v20, v1, v20)
                if (!(-2147483648 >= 1)) {
                    var v21 uint = 658
                    v1 = uint16(v13)
                    _ = v21
                } else {
                    v6[(fi18 << ((34) & 63))] = fi18
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    delete(v6, -1)
    v4 = append(v4, 850)
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 uint64 = 2
        var v24 uint8 = (v13 ^ uint8(v1))
        switch fi22 {
        case -1:
            var v25 int = 32768
            _ = v25
        case -2147483648:
            _ = len(v4)
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    switch ((((v3 + "\n\t") + v3) + ((v3 + "\n\t") + v3)) + "") {
    case "hello":
        v3 = "hello"
    }
    switch 127 {
    case 598:
        v5["\n\t"] = true
    case -1:
        var v26 int8 = (v11 - v11)
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint = 337
        v4 = append(v4, 532)
        _ = fi27
        _ = v28
    }
    v29 := v6[32768]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
