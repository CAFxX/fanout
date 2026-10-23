package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
func genFn0(p0 uint16, p1 int) uint8 {
    p0 -= (^(p0))
    for fi5 := 0; fi5 < 17; fi5++ {
        switch p1 {
        case 799:
            for fi6 := 0; fi6 < 12; fi6++ {
                v7 := GenS1{F0: 0, F1: true}
                switch "foo bar" {
                case "":
                    p0 += 65535
                }
                p1 = 32767
                _ = fi6
                _ = v7
            }
        }
        _ = fi5
    }
    var v8 int32 = -1
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(p0)
        v10 := GenS1{F0: 42, F1: false}
        switch (fi9 + (^(fi9))) {
        case 32768:
            v11 := GenS1{F0: 4294967295, F1: false}
            _ = v11
        case 547:
            if false {
                var v12 uint16 = (p0 >> ((18446744073709551615) & 63))
                _ = v12
            } else {
                var v13 float64 = float64(1e10)
                _ = v13
            }
        case 0:
            var v14 string = "foo bar"
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v8, p1, p1)
    _ = p0
    _ = p1
    _ = v8
    return 214
}
func genFn1() int32 {
    v15 := GenS1{F0: 2, F1: (-32768 > 102)}
    var v16 int64 = 42
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 bool = false
    var v2 uint16 = 2
    var v3 int = -2147483648
    var v4 int32 = 1
    if true {
        v17 := GenS1{F0: 18446744073709551615, F1: (0.0 <= 1e10)}
        switch (v3 % (((^(v3))) | 1)) {
        case 636:
            v2 = 255
        case -9223372036854775808:
            v3 = v3
        default:
            var v18 uint8 = 255
            _ = v18
        }
        _ = v17
    }
    var v19 int64 = 42
    v20 := GenS1{F0: 42, F1: (!(!v1))}
    var v21 int8 = 127
    if (!(128 != 127)) {
        for fi22 := 0; fi22 < 7; fi22++ {
            v1 = false
            var v23 int8 = 42
            _ = fi22
            _ = v23
        }
    }
    v21 = int8(v19)
    v1 = v1
    var v24 uint = 255
    v25 := []int{v3, 415, v3, v3}
    v2 -= v2
    if v1 {
        var v26 uint32 = 2147483648
        _ = v26
    }
    _ = len(v25)
    v24 -= v24
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
