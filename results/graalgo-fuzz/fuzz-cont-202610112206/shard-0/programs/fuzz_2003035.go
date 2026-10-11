package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{0, 4294967295}
    v2 := make(map[int]int)
    var v3 bool = true
    if (!v3) {
        switch ("\n\t" + "\n\t") {
        case "hello":
            v4 := GenS1{F0: 65535, F1: 127}
            _ = v4
        default:
            if v3 {
                fmt.Println(v3)
                var v5 uint16 = 836
                _ = v5
            }
        }
        _ = len(v2)
        v6 := GenS1{F0: 585, F1: -1}
        _ = v6
    }
    v3 = (!(522 >= 0))
    var v7 int64 = -792
    fmt.Println(v7)
    if ((1e10 <= 1.5) && false) {
        v8 := GenS1{F0: 542, F1: 100}
        _ = v8
    } else {
        var v9 float64 = 0.1
        _ = len(v1)
        v9 += (v9 - float64(3.14159))
        _ = v9
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 string = ("hello" + ("foo bar" + "hello"))
        switch (fi10 | fi10) {
        case 0:
            v3 = v3
        }
        var v12 bool = (1 == 42)
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 116, F1: 100}
    for fi14 := 0; fi14 < 1; fi14++ {
        v3 = ((^(fi14)) > fi14)
        v7 += -1
        var v15 uint64 = 315
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: -96}
    var v17 float64 = (0.1 + float64((0.0 * float32(-2.25))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
