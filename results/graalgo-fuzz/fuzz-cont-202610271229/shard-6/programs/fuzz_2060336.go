package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int, p1 int32) int32 {
    switch -9223372036854775808 {
    case -118:
        v5 := []int{(p0 * (p0 << ((2) & 63))), (p0 << ((4294967295) & 63)), 128, p0}
        _ = v5
    case 1:
        v6 := GenS1{F0: 1}
        _ = v6
    case 95:
        p0 -= 32767
    }
    v7 := GenS2{F0: 1, F1: 26}
    var v8 int = (p0 >> ((42) & 63))
    var v9 uint16 = 134
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p1
}

func main() {
    var v1 uint32 = 2
    var v2 int32 = -48
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v1 = (v1 >> ((18446744073709551615) & 63))
    v10 := GenS1{F0: 50}
    fmt.Println(v2, v1)
    v11 := GenS1{F0: 255}
    var v12 uint32 = v1
    switch 502 {
    case 32767:
        if ((^(v12)) > 2147483648) {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 uint16 = 32767
                v15 := GenS2{F0: (v12 >> ((42) & 63)), F1: 65535}
                _ = fi13
                _ = v14
                _ = v15
            }
            if true {
                if true {
                    var v16 uint = 42
                    var v17 uint64 = 18446744073709551615
                    _ = v16
                    _ = v17
                }
                switch "hello" {
                case "":
                    v12 = 1
                case "a":
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v2 += v2
                        var v19 int64 = 1
                        _ = fi18
                        _ = v19
                    }
                case "\n\t":
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 int8 = 1
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 uint64 = 4294967295
                            _ = fi22
                            _ = v23
                        }
                        delete(v4, (fi20 * (fi20 << ((44) & 63))))
                        _ = fi20
                        _ = v21
                    }
                default:
                    if (0.0 >= 0.1) {
                        v1 = v1
                        v12 -= (v1 * (v1 / ((2) | 1)))
                    }
                }
                v24 := []int{32768, -932, 127, 2147483647}
                _ = v24
            }
            switch ((("\n\t" + "a") + (("\n\t" + "a") + "a")) + "") {
            case "a":
                v2 = (v2 ^ -1)
            case "hello":
                v25 := []int{9223372036854775807, 42, -41, 4294967295}
                _ = v25
            }
        }
    case 32768:
        if false {
            v26 := GenS1{F0: 676}
            v2 = 42
            switch -414 {
            case 831:
                var v27 uint32 = (v12 & v12)
                _ = v27
            case 1:
                fmt.Println(v1, v1, v12)
            }
            _ = v26
        }
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 uint64 = 675
        v30 := []int{fi28, (fi28 | (fi28 & (fi28 / ((fi28) | 1))))}
        v31 := v3[(fi28 | fi28)]
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    v4[32768] = (v12 & (v1 % ((65535) | 1)))
    v32 := []int{int(v1), -425}
    var v33 int32 = (v2 / ((v2) | 1))
    v34 := GenS1{F0: 1}
    var v35 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
