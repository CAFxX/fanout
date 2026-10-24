package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
func genFn0() int64 {
    v6 := GenS1{F0: 2147483647}
    switch "hello" {
    case "foo bar":
        var v7 float32 = -2.25
        _ = v7
    case "a":
        var v8 int8 = -38
        _ = v8
    }
    _ = v6
    return -933
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{490, 32767, 0, 127}
    v3 := []string{"", "hello", (("foo bar" + "") + "")}
    v4 := make(map[string]bool)
    var v5 float64 = 1.5
    for fi9 := 0; fi9 < 11; fi9++ {
        if false {
            var v10 int = fi9
            _ = len(v3)
            _ = v10
        } else {
            var v11 string = "\n\t"
            var v12 int32 = -1
            _ = v11
            _ = v12
        }
        switch "hello" {
        case "\n\t":
            v5 += 1e10
        case "a":
            delete(v1, "")
        case "hello":
            var v13 uint8 = 128
            _ = v13
        }
        v5 -= (1e10 * v5)
        _ = fi9
    }
    if ("\n\t" >= "a") {
        var v14 uint8 = 42
        v15 := GenS1{F0: -1}
        v16 := v1["\n\t"]
        _ = v14
        _ = v15
        _ = v16
    }
    if (!("" <= "")) {
        v4[""] = true
        v5 += v5
    } else {
        v5 = v5
    }
    var v17 int16 = -1
    switch 2147483647 {
    case 127:
        for fi18 := 0; fi18 < 17; fi18++ {
            switch int(v17) {
            case 42:
                fmt.Println(v5, fi18)
            case 0:
                var v19 float32 = 1e10
                _ = v19
            case -1:
                for fi20 := 0; fi20 < 7; fi20++ {
                    v17 += (v17 & v17)
                    var v21 int64 = 42
                    delete(v4, "\n\t")
                    _ = fi20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v23 := GenS1{F0: fi22}
                var v24 int32 = -355
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(fi18, fi18, fi18)
            _ = fi18
        }
    case 261:
        v17 = v17
    default:
        var v25 uint64 = 0
        _ = v25
    }
    if false {
        _ = len(v1)
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 int16 = v17
        var v28 int16 = v17
        switch "\n\t" {
        case "\n\t":
            v3 = append(v3, "hello")
        case "foo bar":
            v29 := GenS2{F0: ((198 >= 242) && ((-(v17)) != (v28 & -1)))}
            _ = v29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v5, v5, v5)
        var v31 uint64 = 938
        for fi32 := 0; fi32 < 6; fi32++ {
            v33 := GenS2{F0: (65535 < 219)}
            var v34 bool = ((3.14159 + -2.25) > (1e10 / (0.1 + float32(float64(-2.25)))))
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    fmt.Println(v17)
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
