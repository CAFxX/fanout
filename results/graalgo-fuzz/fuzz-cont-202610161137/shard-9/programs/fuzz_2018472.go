package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{-128, -17, 948, 128}
    v3 := []int{}
    v4 := []string{"\n\t", "foo bar"}
    var v5 uint8 = 0
    var v6 int16 = 1
    v2 = append(v2, 1)
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = true
        var v9 uint64 = 0
        v2 = append(v2, (fi7 + fi7))
        _ = fi7
        _ = v8
        _ = v9
    }
    v3 = append(v3, -413)
    v6 = (v6 + v6)
    delete(v1, "")
    if (true || true) {
        var v10 int8 = 63
        if false {
            var v11 uint8 = (v5 * 0)
            var v12 float64 = float64(1e10)
            for fi13 := 0; fi13 < 6; fi13++ {
                v11 += 121
                v12 -= v12
                _ = fi13
            }
            _ = v11
            _ = v12
        } else {
            switch "a" {
            case "foo bar":
                for fi14 := 0; fi14 < 11; fi14++ {
                    v6 += -1
                    v1["hello"] = false
                    _ = fi14
                }
            case "a":
                var v15 int16 = 1
                _ = v15
            case "":
                v10 = v10
            }
            switch -1 {
            case 4294967295:
                _ = len(v3)
            case -2147483648:
                if len(v3) > 0 {
                    v3[0] = 42
                }
            default:
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
        }
        _ = v10
    } else {
        if true {
            if true {
                v6 += (v6 | 212)
                v16 := GenS1{F0: false, F1: 0}
                _ = v16
            }
        }
    }
    var v17 uint = 759
    var v18 uint64 = 2
    v17 += (v17 | 595)
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS1{F0: true, F1: -297}
            v22 := GenS1{F0: true, F1: 231}
            _ = fi20
            _ = v21
            _ = v22
        }
        v5 = 127
        v6 = (-(v6))
        _ = fi19
    }
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
