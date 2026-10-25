package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []int{}
    var v2 int16 = -32768
    v3 := make(map[uint32]bool)
    if (127 == 127) {
        for fi4 := 0; fi4 < 4; fi4++ {
            switch (-(fi4)) {
            case 32767:
                v2 = (v2 - v2)
            }
            var v5 uint32 = 954
            var v6 uint16 = 602
            _ = fi4
            _ = v5
            _ = v6
        }
        var v7 int16 = v2
        _ = v7
    }
    v8 := GenS1{F0: -2147483648, F1: 56, F2: 42}
    var v9 float32 = 3.14159
    if false {
        var v10 float64 = (-2.25 + 0.1)
        v11 := GenS2{F0: 2}
        var v12 int = -2147483648
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float64 = ((0.1 * (1e10 - 3.14159)) / (3.14159 * (0.1 * float64(v9))))
    if false {
        fmt.Println(v13)
        for fi14 := 0; fi14 < 8; fi14++ {
            switch "hello" {
            case "foo bar":
                v15 := GenS1{F0: 1, F1: 2, F2: -497}
                _ = v15
            }
            _ = fi14
        }
        if false {
            if true {
                v16 := GenS1{F0: 42, F1: 0, F2: 2147483647}
                var v17 uint32 = 2
                _ = v16
                _ = v17
            } else {
                var v18 int32 = 42
                _ = v18
            }
        }
    }
    fmt.Println(v2, v9, v13)
    v13 -= v13
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(fi19, fi19)
        fmt.Println(fi19, v13, fi19)
        for fi20 := 0; fi20 < 8; fi20++ {
            var v21 int8 = 46
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v22 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
