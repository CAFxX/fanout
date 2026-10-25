package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
func genFn0(p0 bool) uint8 {
    switch 1 {
    case 32768:
        p0 = p0
    default:
        for fi4 := 0; fi4 < 17; fi4++ {
            var v5 float64 = float64((3.14159 * (-2.25 - 1e10)))
            _ = fi4
            _ = v5
        }
    }
    v6 := []int{-882, -40, -48}
    var v7 int8 = 22
    fmt.Println(p0, v7)
    v6 = append(v6, 911)
    _ = p0
    _ = v6
    _ = v7
    return 154
}

func main() {
    v1 := []string{"hello"}
    var v2 uint = 2
    var v3 uint = (v2 + v2)
    var v8 uint32 = 0
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v1 = append(v1, "a")
    v9 := GenS1{F0: v2, F1: (v3 ^ v3)}
    var v10 int32 = -1
    var v11 int = 789
    v12 := GenS2{F0: 60, F1: true, F2: -111}
    for fi13 := 0; fi13 < 7; fi13++ {
        if (-397 != -1) {
            v2 = (v2 + 0)
            v3 = 847
            var v14 int32 = (^(v10))
            _ = v14
        } else {
            if len(v1) > 0 {
                v1[0] = (("a" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))) + "a")
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    v1 = append(v1, "foo bar")
    var v15 int8 = 100
    v10 = v10
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(v2, v3)
        _ = fi16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if true {
            var v18 uint16 = 65535
            _ = v18
        }
        _ = fi17
    }
    v15 = v15
    var v19 int32 = (^(v10))
    var v20 int8 = (v15 | v15)
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 int64 = -9223372036854775808
        _ = fi21
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
