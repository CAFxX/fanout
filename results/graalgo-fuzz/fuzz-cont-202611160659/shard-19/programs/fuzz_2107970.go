package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 bool, p2 uint8) uint8 {
    var v7 uint64 = 0
    var v8 int64 = -9223372036854775808
    if (113 >= 42) {
        fmt.Println(v7)
        v8 = v8
    }
    var v9 float32 = 1.5
    switch 127 {
    case 128:
        var v10 string = p0
        _ = v10
    case 653:
        var v11 float64 = 1.5
        _ = v11
    default:
        var v12 string = p0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 int64 = -307
    var v2 int32 = -664
    var v3 int = 899
    v4 := make(map[string]bool)
    var v5 int = v3
    var v6 bool = (2 > 255)
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := []int{1, (v5 * 4294967295)}
            _ = fi13
            _ = v14
        }
    }
    var v15 uint8 = 2
    if ("a" <= ((((("a" + "hello") + ((("" + "foo bar") + "hello") + ("\n\t" + ("\n\t" + "")))) + "\n\t") + ("a" + ("a" + ("a" + "")))) + "hello")) {
        if false {
            v16 := []int{}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            fmt.Println(v15, v6, v6)
            _ = v16
        }
    } else {
        if v6 {
            v17 := GenS1{F0: 83}
            _ = v17
        }
        fmt.Println(v1, v6, v5)
    }
    v3 = v5
    v1 += -925
    var v18 int16 = 0
    var v19 uint32 = 335
    fmt.Println(v3)
    switch (v5 + -9223372036854775808) {
    case 127:
        var v20 float64 = ((0.1 - 1.5) + (1e10 / (float64((((float32((0.0 / 3.14159)) + (0.0 - (0.0 - ((float32(1.5) - (0.0 - -2.25)) - float32((-2.25 * (float64(-2.25) + 0.1))))))) + (0.1 * (float32(1.5) / (1.5 * float32(0.1))))) - 3.14159)) / 1.0)))
        _ = v20
    case 1:
        var v21 uint = 42
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        if false {
            if (!(int64(v3) != (v1 + v1))) {
                v19 = v19
                var v23 bool = v6
                _ = len(v4)
                _ = v23
            } else {
                for fi24 := 0; fi24 < 12; fi24++ {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v15, v3)
                    }
                    _ = fi24
                }
                switch ("" + "") {
                case "foo bar":
                    v4[("foo bar" + "\n\t")] = v6
                case "":
                    var v25 float64 = (-2.25 - (float64((0.1 + 3.14159)) - (3.14159 * (-2.25 * (1e10 - -2.25)))))
                    _ = v25
                default:
                    v26 := GenS2{F0: ("" == "\n\t"), F1: v1}
                    _ = v26
                }
                switch (v3 >> ((18446744073709551615) & 63)) {
                case 32767:
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 uint16 = 42
                        _ = fi27
                        _ = v28
                    }
                }
            }
        } else {
            fmt.Println(v3, v3)
            v29 := GenS3{F0: 255}
            _ = v29
        }
        _ = fi22
    }
    v15 = v15
    var v30 uint = 0
    var v31 uint32 = v19
    v32 := []int{32768, (v5 - v5), v5, (v5 ^ (v5 + 127))}
    var v33 float32 = (float32(0.0) - ((0.1 * (1.5 / float32(1.5))) + 0.1))
    for fi34 := 0; fi34 < 4; fi34++ {
        if len(v32) > 0 {
            v32[0] = v3
        }
        var v35 string = "hello"
        _ = fi34
        _ = v35
    }
    v32 = append(v32, 4294967295)
    var v36 int16 = (^(v18))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
