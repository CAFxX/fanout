package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 bool, p1 string) int16 {
    fmt.Println(p1, p1)
    var v4 int8 = 114
    if false {
        if (0 <= 4294967295) {
            fmt.Println(p1, p0, p1)
        } else {
            v5 := []int{-701, -1}
            _ = v5
        }
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        var v8 uint16 = 376
        _ = v8
    } else {
        v9 := GenS1{F0: 127}
        if false {
            switch 32768 {
            case 993:
                v10 := []int{32768, 958, -718}
                _ = v10
            case -1:
                var v11 uint = 2
                _ = v11
            }
        }
        fmt.Println(p0)
        _ = v9
    }
    fmt.Println(v4, p1, p1)
    _ = p0
    _ = p1
    _ = v4
    return 1000
}
func genFn1(p0 int16, p1 int) string {
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            p0 = p0
            _ = fi13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := []int{}
            v16 := GenS1{F0: fi12}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 float64 = (0.0 * 0.1)
        _ = fi12
        _ = v17
    }
    if false {
        p1 = (p1 * (-(p1)))
        if false {
            if (!true) {
                var v18 uint16 = 2
                _ = v18
            }
            v19 := []int{p1, (p1 / ((42) | 1)), (p1 << ((1) & 63)), (p1 - p1)}
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint16 = 255
                if true {
                    var v22 int16 = (^(p0))
                    _ = v22
                } else {
                    var v23 int = fi20
                    v24 := GenS1{F0: (-(fi20))}
                    var v25 string = "hello"
                    _ = v23
                    _ = v24
                    _ = v25
                }
                if len(v19) > 0 {
                    v19[0] = (^(p1))
                }
                _ = fi20
                _ = v21
            }
            _ = v19
        }
    }
    fmt.Println(p0, p0)
    var v26 int = -157
    _ = p0
    _ = p1
    _ = v26
    return "\n\t"
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 0.0
    v3 := []string{"hello", ((("\n\t" + "foo bar") + ("a" + "foo bar")) + "foo bar")}
    v3 = append(v3, "\n\t")
    v3 = append(v3, ("foo bar" + "\n\t"))
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 int = (fi28 & fi28)
        v27 -= (v27 & 213)
        switch fi28 {
        case 42:
            for fi30 := 0; fi30 < 16; fi30++ {
                _ = len(v3)
                _ = fi30
            }
        case -357:
            fmt.Println(v29)
        case 1:
            v2 -= v2
        }
        _ = fi28
        _ = v29
    }
    v2 = -2.25
    var v31 string = "hello"
    fmt.Println(v27, v27, v31)
    if len(v3) > 0 {
        v3[0] = ("hello" + "foo bar")
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v27 -= v27
        if (v2 <= v2) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) > 0 {
                v3[0] = (v31 + "")
            }
        } else {
            if (82 < 48) {
                v33 := GenS1{F0: (-(fi32))}
                _ = v33
            }
            fmt.Println(v2, v31)
            if (!false) {
                v27 = (v27 + (v27 ^ 2147483648))
            } else {
                fmt.Println(fi32, v27)
                v1[(fi32 * (^(fi32)))] = v31
                v34 := GenS1{F0: (-(fi32))}
                _ = v34
            }
        }
        _ = fi32
    }
    var v35 uint8 = 2
    fmt.Println(v2)
    switch 732 {
    case 32768:
        v27 += (^(v27))
    default:
        if ((!(-329 < 2147483647)) || false) {
            switch -405 {
            case 42:
                var v36 uint = 18446744073709551615
                _ = v36
            }
        } else {
            v35 -= (v35 ^ 2)
            fmt.Println(v35, v31, v27)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v35
}
