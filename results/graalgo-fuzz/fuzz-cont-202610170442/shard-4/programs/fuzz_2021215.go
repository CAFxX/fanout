package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 string, p1 int64, p2 int16) uint32 {
    switch 2147483647 {
    case 1:
        p0 = ("" + ((("hello" + p0) + "") + ((p0 + p0) + "a")))
    case 32768:
        if (819 != 18446744073709551615) {
            for fi7 := 0; fi7 < 19; fi7++ {
                for fi8 := 0; fi8 < 7; fi8++ {
                    p2 += (p2 >> ((2) & 63))
                    if false {
                        p0 = "hello"
                    }
                    _ = fi8
                }
                var v9 uint8 = 183
                _ = fi7
                _ = v9
            }
            fmt.Println(p2)
        } else {
            var v10 int8 = 0
            var v11 int32 = 458
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := []int{(fi12 >> ((416) & 63)), (fi12 >> ((18446744073709551615) & 63))}
                var v14 int16 = -32768
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v10
            _ = v11
        }
    case 127:
        switch 1 {
        case 1:
            v15 := []int{}
            _ = v15
        case 2147483647:
            for fi16 := 0; fi16 < 11; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    p2 = (p2 + (p2 - (p2 << ((296) & 63))))
                    v18 := []int{}
                    fmt.Println(p0)
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
    default:
        v19 := []int{}
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := []int{fi20, (fi20 * (fi20 * (fi20 | (fi20 ^ fi20)))), fi20, -9223372036854775808}
        var v22 string = p0
        v23 := GenS1{F0: 127}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 2
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint = 793
    var v5 uint8 = 2
    var v6 uint32 = 0
    v25 := GenS1{F0: -128}
    fmt.Println(v5, v4)
    if (true || false) {
        fmt.Println(v1)
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            v5 = v5
            v27 := GenS1{F0: 127}
            _ = fi26
            _ = v27
        }
    }
    v5 = v5
    var v28 float32 = (v1 + v1)
    var v29 int8 = 84
    v6 = (v6 << (((v4 | 126)) & 63))
    if (false && false) {
        fmt.Println(v29, v28, v1)
        var v30 uint32 = 2
        _ = v30
    }
    _ = len(v2)
    var v31 uint64 = 1
    if ((1000 <= int16(v6)) && false) {
        if false {
            v32 := []int{860, -727}
            _ = v32
        }
        var v33 int16 = 436
        switch "" {
        case "":
            var v34 uint8 = 83
            _ = v34
        case "hello":
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int32 = 2147483647
                _ = fi35
                _ = v36
            }
        default:
            v37 := GenS2{F0: (((("foo bar" + "a") + ("" + "foo bar")) + (("hello" + ("\n\t" + ("" + "hello"))) + (("\n\t" + "a") + "\n\t"))) + ""), F1: (v31 - (v31 % ((0) | 1)))}
            _ = v37
        }
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
