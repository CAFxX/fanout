package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int) uint8 {
    v7 := []int{(p0 - p0), p0}
    var v8 uint16 = 32767
    var v9 uint8 = 42
    fmt.Println(v8, v9, v8)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 int16) int {
    p0 -= (-(p0))
    var v10 string = "hello"
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := GenS1{F0: 825, F1: 36}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    return 128
}

func main() {
    var v1 uint8 = 0
    var v2 uint8 = (^(v1))
    var v3 uint = 65535
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v6 int16 = int16(v2)
    var v14 int8 = 42
    v6 = (v6 / (((v6 << ((v3) & 63))) | 1))
    if true {
        var v15 uint32 = 1
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 int32 = 1
            switch fi16 {
            case -2147483648:
                v6 -= v6
            case 198:
                v18 := GenS1{F0: -1, F1: v14}
                _ = v18
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS2{F0: 42}
                _ = fi19
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    var v21 bool = false
    v22 := []int{genFn1(v6)}
    var v23 string = "\n\t"
    v24 := GenS1{F0: 23, F1: -113}
    var v25 uint16 = 1
    fmt.Println(v23, v2, v21)
    var v26 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
