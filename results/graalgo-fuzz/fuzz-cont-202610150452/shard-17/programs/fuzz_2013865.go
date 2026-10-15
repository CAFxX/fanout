package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 int }

func main() {
    var v1 int64 = -1
    var v2 uint32 = 0
    var v3 int = 4294967295
    if false {
        v4 := []int{(v3 | (-(v3))), 128}
        v3 = (v3 >> ((285) & 63))
        _ = v4
    }
    v2 -= 4294967295
    if ((-2.25 != -2.25) || false) {
        if true {
            v2 = v2
            var v5 float64 = 3.14159
            v6 := GenS1{F0: 0, F1: 0}
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := []int{}
                v2 -= (^(v2))
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi7
                _ = v8
            }
        }
    } else {
        for fi9 := 0; fi9 < 13; fi9++ {
            switch (^(v3)) {
            case 0:
                v10 := GenS1{F0: 32767, F1: 1}
                _ = v10
            case 1:
                fmt.Println(fi9, v2)
            default:
                v11 := GenS2{F0: 42, F1: ("a" + "hello"), F2: 32768}
                _ = v11
            }
            _ = fi9
        }
    }
    var v12 string = "\n\t"
    v3 = int(v1)
    var v13 int = v3
    if false {
        v14 := GenS2{F0: 425, F1: ((("foo bar" + v12) + v12) + ""), F2: 127}
        _ = v14
    }
    var v15 int64 = v1
    v1 = 0
    var v16 uint8 = 147
    v17 := GenS2{F0: 0, F1: v12, F2: v3}
    var v18 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
