package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint }
func genFn0() int16 {
    switch 127 {
    case 4294967295:
        var v4 int8 = 127
        _ = v4
    default:
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint16 = 245
            v6 = (v6 / ((v6) | 1))
            v6 += v6
            _ = fi5
            _ = v6
        }
    }
    var v7 int8 = 76
    _ = v7
    return -758
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    var v3 string = ("\n\t" + "a")
    v3 = (v3 + "\n\t")
    _ = len(v1)
    fmt.Println(v3)
    v3 = "\n\t"
    if (42 >= -1) {
        switch v3 {
        case "":
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 string = ("" + (v3 + v3))
                var v10 int32 = 44
                v11 := GenS1{F0: (v3 + v3), F1: 9223372036854775807}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        case "foo bar":
            v12 := []int{32767, -2, -1}
            _ = v12
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS2{F0: 65535, F1: -215, F2: 648}
                for fi15 := 0; fi15 < 5; fi15++ {
                    fmt.Println(v3)
                    _ = fi15
                }
                v16 := []int{-940, (fi13 % ((fi13) | 1)), (fi13 | fi13)}
                _ = fi13
                _ = v14
                _ = v16
            }
        }
    }
    v3 = ""
    v3 = "foo bar"
    var v17 int16 = -796
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int16 = v17
        _ = fi18
        _ = v19
    }
    v20 := []int{166}
    fmt.Println(v3, v17)
    delete(v1, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
}
