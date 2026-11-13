package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64; F1 string }
func genFn0(p0 uint32, p1 uint16, p2 uint32) int16 {
    var v4 string = "a"
    for fi5 := 0; fi5 < 15; fi5++ {
        if false {
            v6 := GenS1{F0: 42, F1: 0, F2: 2147483647}
            v4 = ""
            _ = v6
        }
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 18446744073709551615, F1: 2, F2: 42}
            _ = fi7
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    var v1 int = -802
    v2 := make(map[string]string)
    var v3 uint64 = 42
    var v9 uint = 2
    delete(v2, ("\n\t" + ("foo bar" + "hello")))
    var v10 float32 = 3.14159
    v9 += 65535
    for fi11 := 0; fi11 < 17; fi11++ {
        v3 = (v3 * 831)
        var v12 int16 = genFn0(uint32(v3), 32767, 559)
        _ = fi11
        _ = v12
    }
    switch "foo bar" {
    case "a":
        if (((("foo bar" + "\n\t") + "\n\t") >= ("\n\t" + "")) && false) {
            v3 -= (^(v3))
        }
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: -9223372036854775808, F1: ""}
            var v15 int64 = 9223372036854775807
            var v16 bool = true
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
    }
    v10 -= (v10 * v10)
    v17 := GenS2{F0: 9223372036854775807, F1: ""}
    switch (v1 | (v1 % ((v1) | 1))) {
    case 32768:
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                fmt.Println(v1, fi19, v10)
                v1 = (fi19 >> ((255) & 63))
                _ = fi19
            }
            _ = fi18
        }
    default:
        v3 += v3
    }
    var v20 uint32 = 2147483647
    v21 := []int{v1, v1, 1}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v22 uint32 = (v20 + 695)
    _ = len(v21)
    for fi23 := 0; fi23 < 6; fi23++ {
        fmt.Println(v10, v22, v20)
        v20 = v20
        var v24 uint8 = 1
        _ = fi23
        _ = v24
    }
    var v25 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v9
}
