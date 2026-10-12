package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
func genFn0(p0 string, p1 uint16, p2 uint) uint16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        p2 += p2
        _ = fi4
    }
    switch 32768 {
    case -348:
        v5 := []int{42, 370, -9223372036854775808}
        _ = v5
    }
    if (0 != 9223372036854775807) {
        p0 = "a"
        v6 := GenS1{F0: 42}
        _ = v6
    } else {
        switch (p0 + "a") {
        case "hello":
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 float32 = 1.5
                fmt.Println(p0, p1)
                p1 = p1
                _ = fi7
                _ = v8
            }
        }
        v9 := GenS2{F0: 0, F1: -957, F2: 4294967295}
        p2 -= 760
        _ = v9
    }
    if true {
        p1 -= (p1 % ((p1) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]string)
    var v3 uint8 = 128
    if ((!(!true)) || (2147483647 <= 4294967295)) {
        var v10 int8 = -128
        v11 := GenS2{F0: 70, F1: 42, F2: -2147483648}
        _ = v10
        _ = v11
    } else {
        v1 -= (v1 & (v1 / ((v1) | 1)))
        switch 5 {
        case 0:
            var v12 uint8 = (v3 << ((255) & 63))
            _ = v12
        default:
            v1 -= 1
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: 606}
        v1 += (v1 * v1)
        switch "" {
        case "":
            var v17 string = ""
            _ = v17
        case "a":
            v3 = 244
        case "\n\t":
            fmt.Println(fi15, v3, v3)
        default:
            for fi18 := 0; fi18 < 2; fi18++ {
                fmt.Println(fi18, fi18)
                _ = fi18
            }
        }
        _ = fi15
        _ = v16
    }
    v1 = v1
    v1 -= 0
    v19 := []int{}
    switch ("foo bar" + "a") {
    case "\n\t":
        var v20 int = -1
        _ = v20
    default:
        v3 = 128
    }
    v3 = v3
    var v21 string = "\n\t"
    fmt.Println(v21)
    fmt.Println(v21, v21)
    var v22 float32 = (0.1 - 0.0)
    var v23 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
