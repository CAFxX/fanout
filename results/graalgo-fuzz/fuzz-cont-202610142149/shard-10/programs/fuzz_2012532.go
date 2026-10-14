package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 string, p2 string) uint32 {
    switch "" {
    case "a":
        fmt.Println(p2, p0, p2)
    }
    v6 := []int{p0, 42, 111, (p0 >> ((0) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 390
}

func main() {
    var v1 bool = false
    v2 := []int{4294967295, 0}
    var v3 int16 = -270
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    v1 = v1
    v7 := GenS1{F0: genFn0(0, ("foo bar" + "hello"), ""), F1: 306}
    v1 = (4294967295 != 1)
    fmt.Println(v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 string = ("a" + "\n\t")
        _ = fi8
        _ = v9
    }
    if (0 > 0) {
        delete(v4, 32768)
        var v10 int16 = (v3 << ((292) & 63))
        _ = v10
    }
    if false {
        _ = len(v5)
        var v11 float32 = 0.0
        v3 -= v3
        _ = v11
    }
    if (false && (!(0.1 != -2.25))) {
        v3 -= (v3 / ((32767) | 1))
    } else {
        v12 := GenS1{F0: 1, F1: 2}
        v3 = v3
        _ = v12
    }
    v13 := GenS1{F0: 2, F1: 2}
    if (!(57 != 128)) {
        for fi14 := 0; fi14 < 18; fi14++ {
            switch "hello" {
            case "\n\t":
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 string = "hello"
                    switch ("foo bar" + v16) {
                    case "a":
                        v16 = v16
                    }
                    v3 = v3
                    v3 = v3
                    _ = fi15
                    _ = v16
                }
            default:
                v3 += 1000
            }
            _ = fi14
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int64 = -9223372036854775808
            _ = fi17
            _ = v18
        }
        switch "foo bar" {
        case "a":
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint16 = 1
                _ = fi19
                _ = v20
            }
        case "":
            if false {
                fmt.Println(v3, v3)
            }
        }
    }
    v21 := GenS1{F0: 2, F1: 146}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
