package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 int16) uint32 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 1; fi7++ {
        if (798 >= 0) {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: true, F1: -824, F2: fi7}
                var v10 uint8 = 150
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint16 = 1
            switch (fi7 & fi7) {
            case -2147483648:
                v11 += (^(v11))
            case 4294967295:
                v12 := []int{fi7}
                _ = v12
            case -798:
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint64 = 0
                    if (((v14 | (v14 & (v14 & v14))) == v14) && true) {
                        fmt.Println(v14)
                        v11 -= (^(v11))
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = v11
        }
        _ = fi7
    }
    v15 := GenS1{F0: false, F1: 2147483647, F2: 0}
    var v16 uint16 = 0
    if false {
        v17 := GenS1{F0: false, F1: 0, F2: 0}
        v18 := GenS2{F0: 42, F1: 127}
        var v19 bool = (!true)
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v15
    _ = v16
    return 4294967295
}

func main() {
    var v1 uint16 = 2
    var v2 uint = uint(v1)
    var v3 int = 590
    v4 := []int{32768}
    v5 := []int{v3}
    var v6 int32 = 2147483647
    var v20 uint = v2
    var v21 int = v3
    v2 = v2
    switch ("\n\t" + "\n\t") {
    case "hello":
        for fi22 := 0; fi22 < 13; fi22++ {
            _ = len(v5)
            v5 = append(v5, (fi22 << (((v2 - v2)) & 63)))
            v1 -= v1
            _ = fi22
        }
    default:
        var v23 uint32 = 337
        _ = v23
    }
    v2 = (v2 + (v2 & 256))
    v24 := GenS3{F0: -42, F1: (0.1 < (((3.14159 / (1.5 / 0.1)) * 0.0) * 0.1))}
    var v25 float64 = 0.0
    if false {
        var v26 uint32 = 983
        v20 += (v20 >> ((319) & 63))
        v4 = append(v4, (v21 << ((v2) & 63)))
        _ = v26
    } else {
        v27 := GenS3{F0: 42, F1: (0 > int(v1))}
        fmt.Println(v25, v2)
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS1{F0: (!(false && ((1.5 + -2.25) < (float32(3.14159) / (1e10 - (3.14159 / -2.25)))))), F1: v6, F2: (fi28 % ((fi28) | 1))}
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
