package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint }
func genFn0(p0 int16) uint8 {
    var v7 uint16 = 785
    var v8 uint64 = 1
    fmt.Println(v7, v7, p0)
    _ = p0
    _ = v7
    _ = v8
    return uint8(v8)
}
func genFn1(p0 uint, p1 int32, p2 int8) int16 {
    switch 9223372036854775807 {
    case 829:
        var v9 int16 = -911
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    var v4 string = "\n\t"
    v5 := []string{(("foo bar" + "hello") + ("" + v4)), ((v4 + (v4 + (v4 + v4))) + "a"), "hello"}
    v6 := make(map[string]uint32)
    if (2 >= 125) {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch v4 {
            case "":
                v11 := v1["hello"]
                _ = v11
            default:
                v12 := GenS1{F0: 594}
                _ = v12
            }
            _ = fi10
        }
        v13 := GenS1{F0: -1}
        switch 0 {
        case -824:
            v4 = ("" + (("" + ((((v4 + "a") + v4) + (v4 + (v4 + "\n\t"))) + v4)) + ""))
        }
        _ = v13
    } else {
        var v14 string = v4
        var v15 float64 = float64(0.0)
        _ = v14
        _ = v15
    }
    var v16 uint16 = 32767
    v4 = "foo bar"
    for fi17 := 0; fi17 < 4; fi17++ {
        if len(v5) > 0 {
            v5[0] = "a"
        }
        if (0.1 == 0.1) {
            var v18 uint8 = 53
            var v19 string = ""
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: 1}
            v2[v4] = ((!false) || false)
            _ = v20
        }
        v4 = "foo bar"
        _ = fi17
    }
    v4 = (v4 + v4)
    for fi21 := 0; fi21 < 18; fi21++ {
        v16 -= 1
        for fi22 := 0; fi22 < 3; fi22++ {
            if false {
                var v23 bool = (600 >= 308)
                _ = v23
            }
            _ = len(v5)
            _ = fi22
        }
        v24 := GenS2{F0: 255, F1: 546, F2: 248}
        _ = fi21
        _ = v24
    }
    v25 := v1["foo bar"]
    var v26 bool = (!(881 < 42))
    var v27 int32 = 147
    v28 := GenS2{F0: 2, F1: 925, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
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
