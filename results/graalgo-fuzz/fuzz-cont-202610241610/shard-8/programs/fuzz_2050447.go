package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 int32 }
func genFn0(p0 uint, p1 int64) uint8 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{(fi5 * 42)}
        var v7 bool = true
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint = 4294967295
        var v10 int16 = 162
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 float64 = (-2.25 - 1.5)
    _ = p0
    _ = p1
    _ = v11
    return 2
}
func genFn1(p0 string) int {
    var v12 uint16 = 65535
    var v13 string = (p0 + p0)
    _ = p0
    _ = v12
    _ = v13
    return 32768
}

func main() {
    var v1 uint16 = 876
    v2 := []string{"foo bar", "hello"}
    v3 := []int{}
    var v4 string = ("\n\t" + "\n\t")
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 += v1
        _ = fi14
    }
    var v15 int64 = -9223372036854775808
    v16 := GenS1{F0: -891, F1: "foo bar", F2: ((("foo bar" + v4) + "hello") + "hello")}
    var v17 string = v4
    v1 -= 42
    if ((v4 + "a") == v17) {
        v18 := GenS1{F0: -462, F1: v17, F2: ("a" + ((v4 + v4) + ""))}
        _ = v18
    } else {
        v3 = append(v3, 891)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS2{F0: 1}
        switch 9223372036854775807 {
        case -1:
            v21 := GenS3{F0: genFn0(65535, -1), F1: 681}
            _ = v21
        default:
            v22 := GenS4{F0: v17, F1: -820, F2: 42}
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v1)
    switch (v4 + (v4 + v4)) {
    case "a":
        v3 = append(v3, 9223372036854775807)
    case "hello":
        v17 = ((v4 + v17) + v4)
    case "foo bar":
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS4{F0: v4, F1: 42, F2: -906}
            var v25 int16 = 128
            _ = len(v2)
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 uint32 = 0
    v1 -= (v1 / ((2) | 1))
    v2 = append(v2, "foo bar")
    v27 := GenS3{F0: 2, F1: (v1 & 32767)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
