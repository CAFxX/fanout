package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -99
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 uint8 = 42
    var v6 float64 = -2.25
    var v7 int64 = 1
    v8 := []int{}
    var v9 int64 = -425
    for fi10 := 0; fi10 < 9; fi10++ {
        v4[2147483647] = 1
        _ = fi10
    }
    var v11 int64 = -1
    _ = len(v4)
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            var v13 int16 = 624
            v13 = (^(v13))
            _ = v13
        } else {
            var v14 int32 = 816
            fmt.Println(v6, v9, v9)
            _ = v14
        }
        _ = fi12
    }
    _ = len(v8)
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS1{F0: (42 < 42), F1: (^(v5))}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            switch fi16 {
            case -1:
                var v20 bool = (false || false)
                _ = v20
            }
            var v21 uint16 = 2
            v22 := GenS2{F0: 42, F1: (v21 >> ((647) & 63)), F2: (v5 & (v5 + v5))}
            _ = fi19
            _ = v21
            _ = v22
        }
        _ = fi16
    }
    switch 9223372036854775807 {
    case -2147483648:
        v23 := GenS2{F0: 9223372036854775807, F1: 2, F2: (v5 - (v5 >> ((2) & 63)))}
        _ = v23
    default:
        v9 = v11
    }
    v24 := GenS1{F0: ("foo bar" <= "hello"), F1: 117}
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 bool = (1000 == -618)
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 int8 = int8(v5)
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    _ = len(v8)
    v6 += v6
    var v29 float64 = (3.14159 / -2.25)
    v8 = append(v8, 4294967295)
    v5 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
