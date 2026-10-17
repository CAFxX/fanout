package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0() uint8 {
    var v4 uint32 = 4294967295
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int8 = 121
        var v7 uint8 = 1
        v6 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 127 {
    case -9223372036854775808:
        switch 9223372036854775807 {
        case 1:
            if true {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v4 = v4
                    var v9 uint16 = 1
                    v4 = v4
                    _ = fi8
                    _ = v9
                }
            }
        case -773:
            v10 := []int{161, -2147483648, int(v4)}
            _ = v10
        case 128:
            v4 -= v4
        }
    default:
        var v11 uint16 = 899
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: 42, F2: 1}
    v13 := []int{-9223372036854775808, 9223372036854775807, 4294967295, 42}
    _ = v12
    _ = v13
    _ = v4
    return 2
}

func main() {
    var v1 bool = (100 >= 100)
    var v2 bool = (42 == 705)
    var v3 uint32 = 922
    if ("hello" < "hello") {
        var v14 int32 = 42
        fmt.Println(v1)
        _ = v14
    }
    fmt.Println(v3, v3, v1)
    var v15 uint64 = 575
    v16 := GenS2{F0: (v3 * v3), F1: 42, F2: 104}
    switch (("" + (("a" + "hello") + "a")) + "hello") {
    case "a":
        v17 := []int{32767}
        _ = v17
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            v15 -= v15
            var v19 uint = 255
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v3, v3)
    if v1 {
        var v20 int8 = 127
        switch "hello" {
        case "foo bar":
            switch 128 {
            case -2147483648:
                v21 := []int{-471, 32767}
                _ = v21
            }
        }
        var v22 float32 = (0.1 - -2.25)
        _ = v20
        _ = v22
    } else {
        switch (("a" + "a") + "\n\t") {
        case "foo bar":
            v1 = v2
        }
    }
    v23 := []int{127}
    fmt.Println(v15)
    var v24 uint = uint(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
