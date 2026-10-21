package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
func genFn0(p0 string, p1 string) bool {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(fi7 & fi7), (fi7 % (((fi7 ^ fi7)) | 1))}
        var v9 int16 = -1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float32 = (-2.25 - ((0.1 - 1.5) * 0.1))
    var v11 uint16 = 0
    switch p0 {
    case "hello":
        var v12 bool = true
        _ = v12
    case "a":
        fmt.Println(v10, p0)
    case "":
        var v13 int16 = 1
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return false
}
func genFn1() bool {
    var v14 float32 = -2.25
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint16 = 65535
            var v18 uint8 = 99
            v14 = (v14 + v14)
            _ = fi16
            _ = v17
            _ = v18
        }
        v14 -= (v14 / v14)
        _ = fi15
    }
    _ = v14
    return false
}

func main() {
    v1 := []string{("" + (("a" + ("a" + "\n\t")) + "")), "hello", "hello"}
    var v2 float64 = float64(0.0)
    v3 := []string{}
    v4 := make(map[uint32]int)
    var v5 bool = true
    v6 := []uint32{1, 1}
    v5 = true
    var v19 uint32 = 2147483648
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS1{F0: -867, F1: 32767}
        v19 += 4294967295
        var v22 int64 = 42
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := v4[v19]
    var v24 float64 = v2
    v24 += v2
    v19 = (v19 >> ((2) & 63))
    fmt.Println(v23, v19)
    fmt.Println(v24, v5, v19)
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v5)
        _ = fi25
    }
    var v26 int = 32767
    var v27 int8 = -119
    if v5 {
        if (!((65535 != 65535) && (!v5))) {
            if v5 {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v29 := GenS1{F0: 32767, F1: 1000}
                    var v30 int16 = 0
                    v3 = append(v3, (("a" + "foo bar") + ""))
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v24 = (-2.25 / (v2 - v2))
                var v31 uint = 520
                _ = v31
            }
            var v32 float64 = ((v2 + 3.14159) + v24)
            var v33 int64 = -9223372036854775808
            _ = v32
            _ = v33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
