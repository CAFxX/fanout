package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int8, p2 int32) int {
    p1 = 0
    var v7 int8 = p1
    switch int(p2) {
    case 1:
        var v8 int16 = 42
        _ = v8
    default:
        if (p2 != 1) {
            switch "hello" {
            case "a":
                v9 := []int{}
                _ = v9
            case "hello":
                fmt.Println(v7)
            default:
                if (((p2 / (((-(p2))) | 1)) <= p2) || (false && (223 == 2))) {
                    var v10 float32 = (3.14159 / 1e10)
                    v11 := []int{}
                    for fi12 := 0; fi12 < 6; fi12++ {
                        p1 = p1
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v14 := GenS1{F0: 65535, F1: 18446744073709551615}
                            p1 = (v7 << ((4294967295) & 63))
                            var v15 int8 = -128
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        _ = fi12
                    }
                    _ = v10
                    _ = v11
                } else {
                    v16 := GenS1{F0: 4294967295, F1: 18446744073709551615}
                    var v17 uint = 2
                    switch "\n\t" {
                    case "a":
                        v18 := GenS2{F0: uint8(v17)}
                        _ = v18
                    case "hello":
                        v19 := GenS2{F0: 32}
                        _ = v19
                    default:
                        v20 := []int{}
                        _ = v20
                    }
                    _ = v16
                    _ = v17
                }
            }
            v21 := []int{966, 32768}
            _ = v21
        } else {
            switch ("foo bar" + "\n\t") {
            case "":
                p2 = 0
            }
        }
    }
    p0 -= (p0 + (p0 << ((473) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 200
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (0.0 + (0.0 + 3.14159))
    var v4 uint32 = 2147483647
    var v5 float64 = 0.0
    v6 := []string{}
    fmt.Println(v3, v3)
    v4 = (v4 >> ((0) & 63))
    var v22 int16 = 32767
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v22 = (v22 << ((0) & 63))
    var v23 int8 = 127
    v24 := GenS1{F0: 4294967295, F1: 0}
    switch 0 {
    case -253:
        for fi25 := 0; fi25 < 18; fi25++ {
            fmt.Println(fi25)
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 uint32 = 65535
                _ = fi26
                _ = v27
            }
            v28 := GenS1{F0: 18446744073709551615, F1: 141}
            _ = fi25
            _ = v28
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
