package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32) int8 {
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = p1
    return 127
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 int64 = -890
    v5 := make(map[int]bool)
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        if (65535 > 881) {
            v8 := GenS1{F0: -82}
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int16 = 32767
        _ = fi9
        _ = v10
    }
    fmt.Println(v4)
    v4 = (^(v4))
    var v11 int32 = -1
    v12 := GenS1{F0: genFn0(127, v11)}
    if ("hello" != "hello") {
        if true {
            if true {
                v11 = (^(v11))
            }
            switch "a" {
            case "\n\t":
                v11 -= (v11 & v11)
            case "a":
                v4 += -1
            }
            if true {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS1{F0: 34}
                    _ = fi13
                    _ = v14
                }
                v15 := GenS1{F0: -1}
                _ = v15
            }
        }
    } else {
        fmt.Println(v4, v11)
        if (!false) {
            if true {
                var v16 uint8 = 2
                _ = v16
            } else {
                fmt.Println(v11, v11)
            }
        } else {
            fmt.Println(v4, v4)
            var v17 int16 = 924
            v17 += (^(v17))
            _ = v17
        }
    }
    fmt.Println(v4, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
