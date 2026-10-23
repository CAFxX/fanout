package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0() uint8 {
    v4 := GenS1{F0: 65535, F1: 4294967295}
    if false {
        v5 := GenS1{F0: 255, F1: 42}
        var v6 float64 = 1e10
        _ = v5
        _ = v6
    } else {
        switch ("hello" + "a") {
        case "hello":
            if false {
                v7 := GenS1{F0: 255, F1: 1}
                var v8 int = -563
                _ = v7
                _ = v8
            }
        case "a":
            var v9 int = 127
            _ = v9
        default:
            var v10 uint32 = 42
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint16 = 0
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 bool = true
                v12 -= (^(v12))
                if v14 {
                    fmt.Println(v14, fi13, fi11)
                    v12 = 42
                    var v15 int = 128
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := []int{}
        if (0 <= 0) {
            fmt.Println(fi16)
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                v19 := GenS1{F0: 0, F1: 212}
                _ = len(v17)
                _ = fi18
                _ = v19
            }
        }
        v20 := GenS1{F0: 174, F1: 4294967295}
        _ = fi16
        _ = v17
        _ = v20
    }
    v21 := []int{128, 42, 4294967295, 2147483647}
    _ = v21
    _ = v4
    return 44
}

func main() {
    var v1 int64 = 0
    v2 := []int{42}
    v3 := make(map[uint32]uint32)
    fmt.Println(v1)
    if false {
        if ((false && true) || false) {
            for fi22 := 0; fi22 < 11; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 float64 = float64(1.5)
                    var v25 int64 = v1
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                fmt.Println(fi22, v1)
                _ = fi22
            }
            var v26 uint8 = 255
            _ = v26
        } else {
            v1 -= (v1 & (-(v1)))
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v1)
        case "":
            v3[680] = 0
        default:
            for fi27 := 0; fi27 < 13; fi27++ {
                var v28 int32 = 0
                _ = fi27
                _ = v28
            }
        }
    }
    var v29 int8 = 127
    for fi30 := 0; fi30 < 17; fi30++ {
        if false {
            var v31 int = fi30
            fmt.Println(v29, fi30, v29)
            if (true && false) {
                v3[42] = 42
                var v32 uint16 = 737
                v33 := GenS1{F0: v32, F1: 42}
                _ = v32
                _ = v33
            } else {
                v34 := GenS1{F0: 32767, F1: 42}
                var v35 int8 = 95
                _ = v34
                _ = v35
            }
            _ = v31
        }
        _ = fi30
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        fmt.Println(fi36, fi36)
        v37 := v3[360]
        _ = fi36
        _ = v37
    }
    v2 = append(v2, -29)
    v3[314] = 2
    v38 := v3[2147483648]
    v1 -= v1
    var v39 uint16 = uint16(v38)
    var v40 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
