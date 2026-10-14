package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint, p2 int8) uint {
    p2 = 100
    var v4 float64 = (1.5 / 1.5)
    v5 := []int{42}
    switch -2147483648 {
    case 894:
        switch "foo bar" {
        case "":
            var v6 bool = false
            _ = v6
        }
    default:
        v5 = append(v5, 1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}

func main() {
    v1 := []string{"\n\t", ("foo bar" + ("\n\t" + ""))}
    var v2 bool = (1e10 <= 0.1)
    v3 := make(map[string]uint32)
    if v2 {
        var v7 uint32 = 42
        _ = v7
    } else {
        switch "hello" {
        case "\n\t":
            fmt.Println(v2)
        case "a":
            if true {
                var v8 int = -799
                _ = v8
            } else {
                var v9 float64 = (1e10 + 0.1)
                var v10 uint32 = 231
                _ = v10
                _ = v9
            }
        case "hello":
            v2 = (-1 <= 0)
        default:
            var v11 uint32 = 1
            _ = v11
        }
        v12 := GenS1{F0: 663}
        v2 = ((25 > 127) && false)
        _ = v12
    }
    var v13 uint8 = 42
    v14 := GenS1{F0: -2147483648}
    if false {
        var v15 uint32 = 789
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int16 = 1
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        var v18 int8 = -63
        v13 += v13
        _ = v18
    }
    v2 = (0 > -5)
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(fi19, v2)
        _ = len(v1)
        v20 := GenS2{F0: ""}
        _ = fi19
        _ = v20
    }
    v2 = v2
    if (!(!((float64((-2.25 * 0.1)) == 1.5) || v2))) {
        var v21 float64 = (0.1 + 0.1)
        var v22 float64 = float64(float32((((3.14159 - (v21 / (((v21 * ((((0.0 + (0.1 + 3.14159)) - ((1.5 + v21) * v21)) * (v21 + v21)) + (v21 * float64(1.5)))) - ((v21 / 1e10) * float64((0.1 / (-2.25 + 1.5))))) - v21))) - v21) * (v21 * v21))))
        _ = v21
        _ = v22
    } else {
        if (!false) {
            if v2 {
                for fi23 := 0; fi23 < 15; fi23++ {
                    fmt.Println(v13, v2)
                    _ = fi23
                }
                fmt.Println(v13)
                _ = len(v1)
            } else {
                v13 += 255
                v13 = 221
                var v24 int16 = 0
                _ = v24
            }
            fmt.Println(v2, v13)
            var v25 string = ("a" + (("hello" + ("hello" + "foo bar")) + "a"))
            _ = v25
        }
    }
    v13 = v13
    var v26 float64 = (3.14159 / (0.1 + 0.1))
    v26 += (v26 + (v26 + float64(-2.25)))
    var v27 float32 = (0.0 - (1.5 / ((((1.5 / 3.14159) * 0.1) + 1e10) / 1.0)))
    v27 -= v27
    v3[""] = 413
    fmt.Println(v13, v26)
    var v28 int64 = 9223372036854775807
    v29 := GenS3{F0: 42}
    for fi30 := 0; fi30 < 8; fi30++ {
        v26 = v26
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 int64 = v28
            var v33 int64 = v32
            v3[""] = 4294967295
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
