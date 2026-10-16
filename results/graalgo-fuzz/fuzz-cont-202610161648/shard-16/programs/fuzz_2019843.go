package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() bool {
    var v6 int32 = 1
    if (false && (!false)) {
        v7 := []int{-181, -613}
        v6 = (^(v6))
        _ = v7
    } else {
        fmt.Println(v6)
    }
    fmt.Println(v6)
    _ = v6
    return true
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    v3 := []uint32{2147483647, 0, 1}
    v4 := make(map[int]string)
    var v5 int = 0
    if (!false) {
        v2 = (v2 + (v2 + v2))
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v5, v5)
            if (388 >= 255) {
                v5 += (fi8 % (((-(fi8))) | 1))
                delete(v4, 0)
                var v9 int8 = 127
                _ = v9
            }
            _ = fi8
        }
        fmt.Println(v5)
    }
    var v10 float64 = (0.1 / float64(1e10))
    var v11 float64 = (0.1 * 0.1)
    delete(v1, v5)
    for fi12 := 0; fi12 < 18; fi12++ {
        v11 = float64(1e10)
        v2 -= v2
        _ = fi12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        switch "\n\t" {
        case "foo bar":
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi13)
                if true {
                    v15 := GenS1{F0: 2147483647}
                    v10 = float64(1.5)
                    var v16 string = ""
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
        case "hello":
            for fi17 := 0; fi17 < 4; fi17++ {
                v11 = v11
                var v18 uint = 4294967295
                _ = fi17
                _ = v18
            }
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = len(v1)
        fmt.Println(v5)
        _ = fi13
    }
    var v19 int8 = 42
    var v20 int16 = 1
    var v21 bool = false
    v22 := GenS1{F0: 737}
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            fmt.Println(v11)
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
