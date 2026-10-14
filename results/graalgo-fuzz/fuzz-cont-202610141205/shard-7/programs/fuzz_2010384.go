package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
func genFn0(p0 int64, p1 int64) int16 {
    v4 := GenS1{F0: 65535, F1: 42, F2: (^(p1))}
    var v5 float32 = 0.1
    var v6 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 bool, p1 int) uint32 {
    var v7 uint8 = 255
    var v8 uint32 = 934
    v9 := GenS1{F0: 2147483647, F1: 65535, F2: 0}
    if (32767 < -844) {
        var v10 uint32 = (v8 ^ 0)
        v11 := GenS1{F0: v10, F1: 457, F2: 0}
        _ = v10
        _ = v11
    } else {
        v12 := []int{(p1 + 287)}
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint16 = 0
            v14 += v14
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: (p1 - p1), F1: (v7 >> ((722) & 63)), F2: -241}
        _ = v12
        _ = v15
    }
    if p0 {
        fmt.Println(v8)
        p1 = (p1 * (-(p1)))
    } else {
        var v16 uint16 = uint16(v7)
        p1 = p1
        var v17 bool = (0 < 794)
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v8 * (^(v8)))
}

func main() {
    v1 := []int{575, 4294967295, 9223372036854775807, -861}
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    switch ("\n\t" + ("hello" + "a")) {
    case "foo bar":
        var v18 uint8 = 255
        _ = v18
    default:
        switch "\n\t" {
        case "hello":
            if false {
                if true {
                    var v19 int64 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v20 float64 = (1.5 - 3.14159)
                    _ = v19
                    _ = v20
                } else {
                    var v21 uint8 = 103
                    _ = v21
                }
                var v22 float32 = 0.0
                _ = v22
            } else {
                for fi23 := 0; fi23 < 12; fi23++ {
                    var v24 int32 = -2147483648
                    _ = fi23
                    _ = v24
                }
            }
        }
    }
    var v25 bool = (((1e10 / -2.25) * 1e10) < -2.25)
    if v25 {
        _ = len(v1)
        switch 80 {
        case -116:
            var v26 int16 = 1000
            _ = v26
        case 32768:
            delete(v2, -707)
        }
        switch -529 {
        case 2147483647:
            var v27 int = 9223372036854775807
            _ = v27
        case -1:
            v25 = (2 <= 42)
        default:
            var v28 int32 = 71
            _ = v28
        }
    }
    v25 = (128 > 42)
    for fi29 := 0; fi29 < 14; fi29++ {
        v25 = false
        _ = fi29
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v25 = true
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 int16 = 1
        v32 := GenS1{F0: 658, F1: 543, F2: -1}
        var v33 uint16 = 2
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    v25 = (255 > 255)
    if v25 {
        var v34 int32 = 68
        var v35 string = "a"
        fmt.Println(v34, v35, v34)
        _ = v34
        _ = v35
    } else {
        for fi36 := 0; fi36 < 7; fi36++ {
            switch "a" {
            case "hello":
                fmt.Println(fi36, fi36)
            default:
                v25 = ((0.1 / float64((-2.25 / 1.5))) == (3.14159 * ((1e10 * -2.25) + 0.0)))
            }
            fmt.Println(v25)
            _ = fi36
        }
        fmt.Println(v25)
        for fi37 := 0; fi37 < 2; fi37++ {
            v25 = v25
            _ = fi37
        }
    }
    var v38 bool = (42 != 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v38
}
