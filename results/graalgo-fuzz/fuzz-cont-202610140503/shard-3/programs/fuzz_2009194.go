package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (0 < 450)
    var v3 int16 = 871
    v4 := GenS1{F0: 2, F1: 990, F2: 2}
    var v5 uint = 0
    v6 := GenS1{F0: 862, F1: 65535, F2: 0}
    switch int(v3) {
    case -764:
        var v7 uint32 = 919
        _ = v7
    case 1:
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 bool = ((!(-125 < 1)) && false)
            v10 := v1[uint32(v5)]
            for fi11 := 0; fi11 < 19; fi11++ {
                _ = len(v1)
                _ = fi11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    case 42:
        fmt.Println(v3, v3)
    default:
        v2 = (9223372036854775807 == -821)
    }
    fmt.Println(v3, v2)
    switch "hello" {
    case "":
        v12 := []int{}
        _ = v12
    case "a":
        switch ("foo bar" + (("" + "a") + "\n\t")) {
        case "foo bar":
            var v13 uint8 = 121
            _ = v13
        case "":
            switch "foo bar" {
            case "":
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 float32 = 0.1
                    v16 := v1[2147483647]
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "\n\t":
                var v17 uint8 = 32
                _ = v17
            case "foo bar":
                v18 := GenS1{F0: 2, F1: 255, F2: 2147483647}
                _ = v18
            default:
                if (v2 && true) {
                    v19 := GenS2{F0: 2, F1: v3}
                    _ = v19
                } else {
                    var v20 bool = false
                    _ = v20
                }
            }
        case "hello":
            v2 = (1 != 42)
        }
    case "hello":
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint16 = 32767
            _ = fi21
            _ = v22
        }
    }
    v2 = (true || (127 <= 78))
    var v23 float32 = 3.14159
    v24 := GenS2{F0: 4294967295, F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
