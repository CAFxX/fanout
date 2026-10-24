package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int; F1 int32 }
func genFn0() uint64 {
    var v6 int = 32768
    var v7 int = v6
    fmt.Println(v6)
    v8 := GenS1{F0: (36 != 2), F1: -9223372036854775808}
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint {
    v9 := []int{677}
    var v10 uint8 = 127
    var v11 bool = true
    var v12 int = 772
    for fi13 := 0; fi13 < 9; fi13++ {
        v9 = append(v9, (fi13 - (v12 * 32768)))
        fmt.Println(fi13, v11, v10)
        v14 := GenS1{F0: true, F1: 42}
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]bool)
    var v3 int64 = (v1 * -619)
    var v4 int16 = 1
    v5 := make(map[uint32]uint32)
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 int = (^(fi15))
        _ = fi15
        _ = v16
    }
    if false {
        var v17 int8 = -113
        _ = v17
    }
    v18 := GenS1{F0: true, F1: (v3 / ((v1) | 1))}
    for fi19 := 0; fi19 < 17; fi19++ {
        fmt.Println(v3, v3, v1)
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                v1 = (v3 + v1)
                v1 = v1
                _ = fi22
            }
            fmt.Println(v3, v4, v1)
            if (!false) {
                fmt.Println(v4)
            } else {
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 uint16 = 42
                    _ = fi23
                    _ = v24
                }
                if false {
                    if ((!false) && (0.0 < 0.1)) {
                        fmt.Println(v1, v3)
                        var v25 int16 = int16(fi21)
                        _ = v25
                    }
                } else {
                    v26 := []int{fi21}
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v1 -= v1
                        var v28 bool = false
                        var v29 float64 = float64(-2.25)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                }
            }
            _ = fi21
        }
    } else {
        if (false || ((v3 / ((v3) | 1)) < (v1 + v3))) {
            fmt.Println(v3, v3)
        }
    }
    var v30 uint64 = 562
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            switch (fi32 - int(v30)) {
            case -480:
                v1 -= v1
            case 42:
                v33 := GenS2{F0: fi31, F1: 2147483647}
                _ = v33
            default:
                var v34 uint64 = v30
                _ = v34
            }
            _ = fi32
        }
        fmt.Println(fi31, v30, fi31)
        v5[537] = 2147483647
        _ = fi31
    }
    v35 := GenS1{F0: ((((("\n\t" + "") + ("\n\t" + "hello")) + "\n\t") + ("a" + (("a" + ("\n\t" + "hello")) + (("a" + "") + ("a" + "a"))))) > ""), F1: v3}
    var v36 int64 = 860
    v3 = 42
    v2[1] = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
