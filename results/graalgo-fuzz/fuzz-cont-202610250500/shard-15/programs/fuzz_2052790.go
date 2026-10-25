package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 127
    var v3 int64 = -9223372036854775808
    var v4 uint32 = 4294967295
    var v5 int16 = 58
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 520}
    v8 := GenS1{F0: 42}
    v9 := []int{1, (v2 >> ((250) & 63)), (v2 ^ 490)}
    for fi10 := 0; fi10 < 12; fi10++ {
        v3 -= v3
        _ = fi10
    }
    v3 = -9223372036854775808
    v2 += v2
    switch v2 {
    case 2147483647:
        delete(v6, 576)
    default:
        switch v2 {
        case -642:
            switch (v2 | v2) {
            case 128:
                _ = len(v1)
            case 42:
                switch 128 {
                case 32767:
                    var v11 uint16 = 1
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 uint = 2
                    var v14 bool = false
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case 9223372036854775807:
            var v15 float64 = -2.25
            _ = v15
        case 32767:
            v4 += 2
        }
    }
    var v16 int16 = 1000
    if false {
        v16 -= -972
        for fi17 := 0; fi17 < 15; fi17++ {
            v2 = v2
            var v18 string = ""
            v16 = int16(v3)
            _ = fi17
            _ = v18
        }
    }
    _ = len(v9)
    v5 -= (v16 | v5)
    v16 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
