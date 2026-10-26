package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 string, p1 int32, p2 uint) int {
    p1 -= (^(p1))
    var v4 uint = (^(p2))
    var v5 string = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 bool = (-101 != 42)
    var v2 uint32 = 901
    v3 := make(map[string]string)
    var v6 string = "a"
    v2 += v2
    v7 := GenS1{F0: (^(v2)), F1: (v6 + "a")}
    v6 = "\n\t"
    v3[""] = v6
    if (v1 || false) {
        if (100 != -58) {
            if (!v1) {
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 string = "foo bar"
                    var v10 uint64 = uint64(v2)
                    var v11 int64 = -212
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v2 -= v2
            }
        }
        switch "\n\t" {
        case "a":
            v6 = ("" + v6)
        }
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v13 := GenS1{F0: v2, F1: ("" + v6)}
            var v14 int64 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v2, v2)
    v1 = (9 != -77)
    v15 := v3[""]
    v16 := []int{-813, genFn0("", -2147483648, 0), -9223372036854775808, 127}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
