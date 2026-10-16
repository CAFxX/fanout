package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 bool = (-766 != -493)
    var v3 int = -144
    v4 := make(map[uint32]int)
    var v5 string = ("hello" + v1)
    v6 := []uint32{1, 396}
    if v2 {
        v1 = v1
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            v4[2147483648] = -2147483648
            if (!v2) {
                v3 = (v3 - v3)
            }
            switch (fi7 & -9223372036854775808) {
            case -1:
                var v8 int32 = -2147483648
                _ = v8
            default:
                var v9 bool = (745 <= 65535)
                _ = v9
            }
            _ = fi7
        }
        v5 = "foo bar"
        _ = len(v6)
    }
    if ((3.14159 - 0.0) == 1e10) {
        switch "a" {
        case "hello":
            v10 := GenS1{F0: 32767, F1: 0}
            _ = v10
        case "\n\t":
            v11 := GenS1{F0: 58, F1: 42}
            _ = v11
        default:
            v4[2147483647] = v3
        }
    } else {
        var v12 uint = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        _ = len(v4)
        if false {
            v1 = v1
            v14 := GenS1{F0: 305, F1: 42}
            _ = v14
        }
        _ = fi13
    }
    var v15 int16 = 32767
    switch "a" {
    case "a":
        var v16 float32 = (float32(3.14159) - 3.14159)
        _ = v16
    case "\n\t":
        var v17 float32 = (3.14159 / -2.25)
        _ = v17
    }
    v2 = (v2 || v2)
    if (true || true) {
        var v18 uint = 0
        _ = v18
    } else {
        v19 := GenS2{F0: 127}
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v1 = (v5 + (v1 + "\n\t"))
        v2 = v2
        _ = fi20
    }
    fmt.Println(v15, v3)
    var v21 uint = 42
    delete(v4, 1)
    for fi22 := 0; fi22 < 4; fi22++ {
        delete(v4, 2)
        v1 = (("\n\t" + (v1 + v1)) + v5)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
