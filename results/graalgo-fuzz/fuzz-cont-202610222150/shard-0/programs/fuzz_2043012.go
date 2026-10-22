package main

import (
    "fmt"
)

func genFn0(p0 uint32, p1 uint) uint16 {
    var v5 uint8 = 127
    var v6 int = 32767
    v7 := []int{}
    p1 += (p1 | 2)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 int8 = 99
    var v2 uint8 = 252
    v3 := make(map[int]string)
    var v4 int = 32768
    for fi8 := 0; fi8 < 14; fi8++ {
        v1 = (v1 - (v1 & (v1 - 42)))
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(v2, v1)
        v1 += 0
        _ = fi9
    }
    v4 = 0
    switch "" {
    case "":
        fmt.Println(v2, v2)
    }
    switch "a" {
    case "\n\t":
        var v10 uint32 = 829
        _ = v10
    }
    var v11 string = "a"
    if false {
        if true {
            var v12 int32 = 798
            _ = v12
        } else {
            if (42 != (v1 * (^(v1)))) {
                var v13 int8 = v1
                v4 = -994
                _ = v13
            }
        }
    } else {
        v14 := []int{(v4 | v4), (v4 >> ((255) & 63))}
        _ = v14
    }
    switch -1 {
    case 127:
        var v15 int16 = 0
        _ = v15
    }
    if (false || (true || false)) {
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v4)
            _ = fi16
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if ((-(v1)) != v1) {
            for fi18 := 0; fi18 < 2; fi18++ {
                v11 = v11
                var v19 bool = (-336 != 1)
                var v20 int8 = v1
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            var v21 float64 = 1e10
            _ = v21
        }
        _ = fi17
    }
    fmt.Println(v2)
    var v22 uint64 = 42
    v22 = 567
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
