package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint, p1 int64) uint {
    if true {
        if true {
            var v4 int64 = (p1 ^ -614)
            _ = v4
        } else {
            var v5 int32 = 0
            _ = v5
        }
    }
    p1 += p1
    var v6 int16 = 0
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 int64 = -865
    var v3 uint64 = 2
    switch -2147483648 {
    case -22:
        var v7 uint = 4294967295
        _ = v7
    case 22:
        v8 := []int{-219, 0, 1}
        _ = v8
    case 42:
        for fi9 := 0; fi9 < 18; fi9++ {
            if false {
                var v10 uint32 = 42
                switch fi9 {
                case 1:
                    if (!false) {
                        var v11 float64 = 3.14159
                        _ = v11
                    }
                }
                v3 += v3
                _ = v10
            } else {
                v2 = v2
            }
            _ = fi9
        }
    }
    fmt.Println(v3, v2, v1)
    for fi12 := 0; fi12 < 14; fi12++ {
        v3 += v3
        _ = fi12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v3 += 248
        for fi14 := 0; fi14 < 15; fi14++ {
            fmt.Println(v1, fi13, fi13)
            fmt.Println(v3)
            _ = fi14
        }
        _ = fi13
    }
    v15 := []int{1}
    var v16 float64 = 0.1
    if false {
        fmt.Println(v1)
        v2 += v2
        if true {
            switch -15 {
            case 2147483647:
                v17 := GenS1{F0: 127}
                _ = v17
            case 32767:
                for fi18 := 0; fi18 < 17; fi18++ {
                    v1 = v1
                    _ = fi18
                }
            case -880:
                var v19 float64 = float64(-2.25)
                _ = v19
            default:
                var v20 float32 = 0.1
                _ = v20
            }
            v1 = v1
        } else {
            _ = len(v15)
            var v21 uint32 = 2
            for fi22 := 0; fi22 < 17; fi22++ {
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = fi22
            }
            _ = v21
        }
    }
    var v23 float64 = 0.1
    var v24 uint16 = 147
    if (false && (!(v24 > (v24 / (((v24 % (((v24 * (v24 & v24))) | 1))) | 1))))) {
        var v25 uint8 = 113
        v1 = v1
        var v26 float64 = (v23 * v16)
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
