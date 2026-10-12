package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint }
func genFn0(p0 uint, p1 int8, p2 int64) uint64 {
    p2 += p2
    p0 = 1
    p0 = 910
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 int16 = -1
    var v2 bool = false
    var v3 uint32 = 905
    var v4 int32 = 1
    v5 := GenS1{F0: 32767, F1: 84, F2: 4294967295}
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = ((1.5 * -2.25) + 0.0)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{(fi8 & 32767), fi8}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{(fi10 | (^(fi10))), (fi10 >> ((828) & 63))}
        _ = fi10
        _ = v11
    }
    fmt.Println(v2, v1)
    v12 := []int{32767, 4294967295}
    if true {
        v1 += (v1 ^ -1)
        var v13 string = ""
        _ = v13
    }
    var v14 uint32 = v3
    var v15 int64 = -9223372036854775808
    var v16 int64 = v15
    if len(v12) > 0 {
        v12[0] = -318
    }
    var v17 bool = (16 != 127)
    for fi18 := 0; fi18 < 20; fi18++ {
        _ = len(v12)
        for fi19 := 0; fi19 < 20; fi19++ {
            if v17 {
                switch "foo bar" {
                case "foo bar":
                    if (("foo bar" + "\n\t") != "foo bar") {
                        v20 := GenS1{F0: 305, F1: 2, F2: 2}
                        var v21 int32 = v4
                        fmt.Println(v2, v16)
                        _ = v20
                        _ = v21
                    }
                default:
                    v2 = true
                }
            } else {
                var v22 uint16 = 65535
                var v23 uint32 = (v14 % ((v3) | 1))
                _ = v22
                _ = v23
            }
            var v24 int16 = (^(v1))
            v17 = v2
            v17 = false
            var v25 int16 = v1
            _ = fi19
            _ = v24
            _ = v25
        }
        v3 += (v3 / ((v14) | 1))
        _ = fi18
    }
    v26 := GenS1{F0: 102, F1: 65535, F2: 65535}
    v1 = v1
    if (!(v2 || v2)) {
        for fi27 := 0; fi27 < 14; fi27++ {
            fmt.Println(v17, v3, v1)
            _ = fi27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
