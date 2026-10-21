package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 int32, p1 uint64) int8 {
    fmt.Println(p0, p0)
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int64, p1 uint) bool {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int64 = p0
        _ = fi6
        _ = v7
    }
    v8 := []int{2147483647, 539}
    for fi9 := 0; fi9 < 17; fi9++ {
        if len(v8) > 0 {
            v8[0] = (fi9 + (fi9 % ((fi9) | 1)))
        }
        _ = fi9
    }
    if true {
        var v10 int = 4294967295
        _ = v10
    } else {
        v11 := GenS1{F0: 4294967295}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    return true
}

func main() {
    v1 := []string{"a", "foo bar", "\n\t", "hello"}
    v2 := make(map[int]uint32)
    v3 := []string{"a", "", "\n\t", "hello"}
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    switch -1 {
    case 4294967295:
        var v12 float32 = 0.1
        _ = v12
    case 0:
        switch (((("\n\t" + "foo bar") + "\n\t") + "hello") + "a") {
        case "foo bar":
            for fi13 := 0; fi13 < 8; fi13++ {
                v5[("hello" + ("a" + ("\n\t" + ("\n\t" + "foo bar"))))] = "hello"
                v14 := GenS1{F0: 80}
                var v15 uint64 = 4294967295
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 int64 = -924
    switch 649 {
    case 0:
        v16 = v16
    case 32768:
        v17 := GenS1{F0: 4294967295}
        _ = v17
    case 4294967295:
        if false {
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 int32 = -2147483648
                var v20 uint32 = 2147483648
                var v21 string = (("\n\t" + "\n\t") + ("foo bar" + ""))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            v16 = (-(v16))
            v16 += v16
        } else {
            var v22 bool = false
            _ = v22
        }
    default:
        var v23 float64 = 3.14159
        _ = v23
    }
    v3 = append(v3, "")
    var v24 int64 = (v16 - v16)
    switch "a" {
    case "":
        switch 2147483647 {
        case -670:
            v24 = (v24 - v24)
        default:
            v25 := GenS1{F0: 4294967295}
            _ = v25
        }
    }
    fmt.Println(v24)
    var v26 float32 = (0.0 / 1.5)
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 uint16 = 0
        v24 = v24
        _ = fi27
        _ = v28
    }
    var v29 int = -255
    if true {
        var v30 string = "\n\t"
        var v31 uint = 0
        v4[127] = 42
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 19; fi32++ {
            v26 = (((((1e10 - ((v26 / (1.5 - (3.14159 / (float32((1e10 / 3.14159)) + v26)))) - 0.0)) + v26) - float32(1e10)) * v26) - (float32(0.1) - float32(0.1)))
            switch "hello" {
            case "hello":
                v26 = -2.25
            }
            _ = fi32
        }
        v29 += (v29 | v29)
        for fi33 := 0; fi33 < 8; fi33++ {
            v34 := GenS2{F0: "a", F1: v24}
            var v35 int = (fi33 << ((27) & 63))
            v29 = 128
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v36 := GenS1{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
