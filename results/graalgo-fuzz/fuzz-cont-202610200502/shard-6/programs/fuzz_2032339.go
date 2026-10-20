package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0(p0 int16, p1 uint8, p2 uint32) int8 {
    fmt.Println(p1)
    var v6 int64 = int64(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 int16) uint8 {
    v7 := []int{32768}
    var v8 float64 = float64(1e10)
    v9 := GenS1{F0: 100, F1: 26}
    if true {
        for fi10 := 0; fi10 < 4; fi10++ {
            switch "a" {
            case "":
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 float32 = 1.5
                    if (4294967295 <= 165) {
                        v13 := GenS1{F0: 1, F1: 0}
                        var v14 uint16 = 679
                        _ = v13
                        _ = v14
                    }
                    _ = len(v7)
                    _ = fi11
                    _ = v12
                }
            case "hello":
                v8 = v8
            case "foo bar":
                if (90 <= p0) {
                    var v15 uint16 = 642
                    _ = v15
                } else {
                    fmt.Println(p1, v8)
                }
            }
            v8 += v8
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 int16 = (p1 << ((4294967295) & 63))
                if (true || (false && true)) {
                    if len(v7) > 0 {
                        v7[0] = 9223372036854775807
                    }
                } else {
                    fmt.Println(v8)
                }
                v7 = append(v7, fi10)
                _ = fi16
                _ = v17
            }
            _ = fi10
        }
    } else {
        p1 = p1
        switch 365 {
        case 42:
            p1 = (-(p1))
        case 4294967295:
            var v18 string = "hello"
            _ = v18
        }
    }
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 35
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]int)
    v3 := []uint32{65535, 2147483647}
    var v4 int16 = 1
    v5 := []int{372}
    fmt.Println(v1, v4)
    v4 = 32767
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint32 = 76
        switch 128 {
        case 32768:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 4294967295:
            v21 := GenS1{F0: 115, F1: 18446744073709551615}
            _ = v21
        case -73:
            var v22 bool = false
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v23 := GenS1{F0: 42, F1: 42}
    v24 := GenS1{F0: -109, F1: 2}
    fmt.Println(v4, v4)
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS2{F0: 42, F1: -1, F2: 255}
        var v27 int64 = 0
        v28 := GenS2{F0: 654, F1: v1, F2: 55}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v4 = -1
    var v29 uint = 0
    fmt.Println(v4, v4, v4)
    if false {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 uint = v29
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
