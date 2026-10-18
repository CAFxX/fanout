package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8) int8 {
    v7 := GenS1{F0: -148, F1: 656, F2: -9223372036854775808}
    v8 := GenS1{F0: -823, F1: 2, F2: 32767}
    var v9 uint8 = 128
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 int = 750
    var v5 int32 = 1
    v6 := []int{}
    v5 += (v1 & 0)
    switch "foo bar" {
    case "foo bar":
        _ = len(v6)
    case "\n\t":
        var v10 uint = 1
        _ = v10
    case "hello":
        if (!((!(!false)) || (!(2 <= 1)))) {
            v2[v4] = 850
            v11 := GenS2{F0: true, F1: 1000, F2: 1}
            _ = v11
        }
    default:
        v4 -= (v4 - 9223372036854775807)
    }
    delete(v3, (v4 << ((1) & 63)))
    delete(v3, v4)
    switch "foo bar" {
    case "":
        var v12 uint = 65535
        _ = v12
    default:
        v4 = 32767
    }
    if (false || (!false)) {
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v5)
            v1 += v1
            _ = fi13
        }
        if (false || (false && (true || false))) {
            var v14 uint32 = 2147483648
            switch v4 {
            case 360:
                v15 := GenS3{F0: 641}
                _ = v15
            case 32768:
                for fi16 := 0; fi16 < 17; fi16++ {
                    if len(v6) > 0 {
                        v6[0] = (v4 % ((-9223372036854775808) | 1))
                    }
                    _ = fi16
                }
            case 42:
                _ = len(v2)
            }
            _ = v14
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS4{F0: 127, F1: 4294967295}
            _ = fi17
            _ = v18
        }
    }
    var v19 float64 = -2.25
    var v20 uint32 = 934
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
