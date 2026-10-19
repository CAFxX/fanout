package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    var v1 int32 = 926
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 bool = ((!false) || false)
    v5 := make(map[uint32]string)
    if (-1 <= -174) {
        v6 := GenS1{F0: 466, F1: 128}
        var v7 float64 = float64(3.14159)
        v8 := GenS1{F0: 796, F1: 978}
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 float32 = 0.0
        v10 = (float32(0.1) - float32(float64(v10)))
        var v11 bool = (3.14159 <= 3.14159)
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v1, v4)
    case "a":
        var v12 string = ("foo bar" + "")
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if v4 {
            for fi14 := 0; fi14 < 1; fi14++ {
                v1 -= 873
                _ = fi14
            }
            var v15 int64 = 838
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := []int{}
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    fmt.Println(v1, v4)
    v4 = v4
    for fi18 := 0; fi18 < 1; fi18++ {
        v1 = 1
        _ = fi18
    }
    var v19 int = -173
    var v20 uint32 = 0
    var v21 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
