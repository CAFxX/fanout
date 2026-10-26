package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 uint64, p2 int) int64 {
    v7 := GenS1{F0: -1, F1: 467}
    v8 := []int{856, p2, (p0 >> ((42) & 63))}
    fmt.Println(p0, p2, p1)
    fmt.Println(p2, p1, p0)
    var v9 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 267
}

func main() {
    var v1 string = (("\n\t" + "a") + "\n\t")
    v2 := make(map[string]bool)
    v3 := []uint32{2147483648, 835, 1, 42}
    var v4 int64 = 1
    var v5 int = -9223372036854775808
    v6 := make(map[int]string)
    v10 := GenS2{F0: 126}
    if false {
        v1 = "hello"
        v4 += v4
        var v11 string = (("foo bar" + v1) + ("" + "\n\t"))
        _ = v11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS2{F0: 65535}
    if (((!(4294967295 <= 2)) || true) || (!true)) {
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := v2[v1]
            var v15 int8 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 uint16 = 2
    _ = len(v3)
    fmt.Println(v4, v16, v16)
    v16 = v16
    for fi17 := 0; fi17 < 2; fi17++ {
        v16 -= v16
        fmt.Println(fi17, v4)
        _ = fi17
    }
    switch ((v1 + v1) + v1) {
    case "":
        fmt.Println(v5, v4, v5)
    case "hello":
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                v1 = v1
                if false {
                    var v20 int8 = 100
                    _ = v20
                }
                v21 := GenS1{F0: -1, F1: 255}
                var v22 float64 = -2.25
                var v23 int32 = 541
                _ = fi19
                _ = v21
                _ = v22
                _ = v23
            }
            v3 = append(v3, 42)
            _ = fi18
        }
    default:
        if true {
            _ = len(v3)
            _ = len(v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch v1 {
            case "foo bar":
                v5 -= v5
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v24 float64 = float64(-2.25)
            v3 = append(v3, 1)
            _ = v24
        }
    }
    fmt.Println(v5, v5)
    var v25 int32 = 508
    var v26 int8 = -128
    var v27 int8 = (-(v26))
    var v28 uint = 732
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
