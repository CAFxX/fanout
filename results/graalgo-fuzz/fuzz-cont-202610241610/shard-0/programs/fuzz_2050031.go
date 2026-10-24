package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 uint32, p1 bool) int8 {
    v4 := GenS1{F0: 945, F1: p1, F2: "\n\t"}
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(p0, p0, fi5)
        p0 = p0
        var v6 uint32 = p0
        _ = fi5
        _ = v6
    }
    var v7 int64 = 590
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 92
}

func main() {
    var v1 uint16 = 2
    v2 := []string{"hello", "", ""}
    var v3 float32 = 0.1
    if (-1 <= -474) {
        fmt.Println(v1, v1)
        if (!false) {
            v3 = (v3 + float32(((-2.25 + ((0.0 + 0.1) / (1.5 * (3.14159 - (1.5 / 1.0))))) / (1.5 / float64(v3)))))
            var v8 int8 = 127
            v3 = (3.14159 * v3)
            _ = v8
        } else {
            v2 = append(v2, "foo bar")
            if (true && false) {
                var v9 uint8 = 2
                switch 4294967295 {
                case 32767:
                    var v10 float64 = -2.25
                    _ = v10
                }
                var v11 bool = true
                _ = v11
                _ = v9
            }
            if (false && true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v1 = 399
            } else {
                v1 += (v1 / (((^(v1))) | 1))
            }
        }
        fmt.Println(v3)
    }
    var v12 bool = (((((("foo bar" + "") + "\n\t") + ("\n\t" + "\n\t")) + "") + "") == "foo bar")
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS2{F0: "\n\t", F1: 1000}
        _ = len(v2)
        _ = fi13
        _ = v14
    }
    if v12 {
        var v15 int64 = 42
        var v16 int16 = 32767
        fmt.Println(v15)
        _ = v15
        _ = v16
    } else {
        _ = len(v2)
        if v12 {
            if (!((4294967295 == 1) && v12)) {
                v3 = ((1e10 * v3) / (v3 + ((v3 / v3) / ((v3 + v3) / -2.25))))
                fmt.Println(v3)
            }
        }
        v3 = (((float32(float64((1e10 - v3))) * (v3 + v3)) * (v3 - v3)) - 0.1)
    }
    v3 -= float32(0.1)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v1, v1, v12)
    var v17 float64 = (-2.25 - 0.0)
    var v18 int64 = 789
    var v19 uint8 = 1
    v18 = v18
    switch 32767 {
    case 128:
        var v20 int64 = v18
        _ = v20
    case -570:
        v17 -= v17
    }
    var v21 uint16 = v1
    fmt.Println(v1, v18, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
