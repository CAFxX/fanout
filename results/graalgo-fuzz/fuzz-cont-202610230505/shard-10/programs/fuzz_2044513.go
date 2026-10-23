package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 int32; F1 uint16; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 717
    v3 := []uint32{747}
    v4 := []string{"hello", ("" + "a"), "a", "foo bar"}
    v5 := make(map[string]string)
    v6 := GenS1{F0: 9223372036854775807, F1: 1000, F2: ("hello" + "\n\t")}
    if (true && (0.0 > 1.5)) {
        v2 = -734
    } else {
        v2 = (v2 / ((v2) | 1))
    }
    var v7 uint8 = 227
    var v8 bool = (!false)
    v8 = (!((1.5 - 1e10) != 1.5))
    switch "a" {
    case "a":
        switch -9223372036854775808 {
        case 1:
            v7 = 2
        case 32767:
            v9 := GenS2{F0: 2147483647, F1: 2, F2: "hello"}
            _ = v9
        case 0:
            switch (v2 * 42) {
            case -789:
                for fi10 := 0; fi10 < 18; fi10++ {
                    if (1000 > 1) {
                        var v11 int64 = 42
                        var v12 int8 = 38
                        _ = v11
                        _ = v12
                    }
                    v7 -= v7
                    var v13 float32 = 1.5
                    _ = fi10
                    _ = v13
                }
            case 0:
                var v14 uint16 = 42
                _ = v14
            }
        }
    }
    var v15 int64 = 1
    fmt.Println(v15, v7, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
