package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint, p1 bool) int8 {
    v4 := GenS1{F0: 2}
    v5 := []int{32767}
    v5 = append(v5, -1)
    var v6 int64 = 0
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 127
}
func genFn1() string {
    if (((!false) && (1000 <= -649)) && (0.1 < 3.14159)) {
        var v7 bool = true
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 float64 = 0.0
            var v10 string = "\n\t"
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        var v11 uint16 = 65535
        var v12 uint8 = 255
        v13 := []int{128}
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 string = "a"
    v15 := []int{127, -963, -9223372036854775808}
    var v16 float32 = 1.5
    _ = v14
    _ = v15
    _ = v16
    return ("\n\t" + (v14 + (v14 + v14)))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(("hello" + "hello") + "hello"), "", "\n\t", "a"}
    var v3 int16 = 1000
    var v17 uint8 = 127
    fmt.Println(v3)
    fmt.Println(v17)
    v18 := v1[1]
    v18 = v18
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float64 = 0.1
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v3, v17)
    var v21 int16 = (v3 & -190)
    for fi22 := 0; fi22 < 17; fi22++ {
        v17 = (v17 << ((2) & 63))
        _ = fi22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        switch ("hello" + (((("hello" + "\n\t") + "hello") + "a") + "hello")) {
        case "hello":
            fmt.Println(v3, v18)
        case "foo bar":
            fmt.Println(v17, fi23, v21)
        case "a":
            var v24 uint16 = 565
            _ = v24
        default:
            var v25 uint32 = v18
            _ = v25
        }
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "")
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
