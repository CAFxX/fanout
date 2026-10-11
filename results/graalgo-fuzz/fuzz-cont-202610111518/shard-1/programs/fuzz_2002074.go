package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int64) int {
    v6 := []int{9223372036854775807}
    v7 := GenS1{F0: -32768}
    p0 -= p0
    _ = p0
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 bool) int8 {
    switch "foo bar" {
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 bool = false
            if ((255 > 255) && ((fi8 / ((genFn0(p0)) | 1)) > fi8)) {
                v9 = (-32768 > -1)
            }
            _ = fi8
            _ = v9
        }
    case "":
        fmt.Println(p0)
    case "foo bar":
        fmt.Println(p1)
    }
    v10 := []int{-756, -423, 128, genFn0(0)}
    var v11 uint16 = 42
    switch ("\n\t" + "hello") {
    case "foo bar":
        fmt.Println(p1, v11)
    case "\n\t":
        var v12 float64 = ((1.5 + (((1e10 + (3.14159 * 3.14159)) - float64(3.14159)) / (0.1 + 3.14159))) * 3.14159)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -1
}

func main() {
    var v1 int = 128
    var v2 int8 = -128
    v3 := []int{v1, v1, v1, v1}
    var v4 float32 = float32((3.14159 * -2.25))
    v5 := []int{v1}
    var v13 int16 = 1000
    v2 += (v2 - v2)
    var v14 string = "a"
    _ = len(v5)
    switch "a" {
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        if false {
            if false {
                if false {
                    if true {
                        var v15 uint8 = 127
                        _ = v15
                    }
                    var v16 string = ((("" + ("foo bar" + v14)) + (v14 + (("\n\t" + "") + (v14 + "a")))) + "\n\t")
                    _ = v16
                } else {
                    v17 := GenS1{F0: v13}
                    fmt.Println(v2, v4, v14)
                    _ = v17
                }
                v18 := GenS1{F0: 0}
                _ = v18
            } else {
                fmt.Println(v1, v4)
                for fi19 := 0; fi19 < 18; fi19++ {
                    v13 -= int16(v2)
                    fmt.Println(fi19, v4)
                    var v20 float32 = 3.14159
                    _ = fi19
                    _ = v20
                }
            }
        }
    default:
        v2 = genFn1(-1, (v1 != v1))
    }
    v21 := GenS1{F0: v13}
    v13 = (v13 >> ((42) & 63))
    var v22 int16 = v13
    v22 -= v13
    v23 := GenS1{F0: 1000}
    var v24 int32 = 1
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS1{F0: v13}
        _ = fi25
        _ = v26
    }
    var v27 int32 = v24
    for fi28 := 0; fi28 < 3; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int16 = v13
            var v31 int32 = (v24 % (((-(v24))) | 1))
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
