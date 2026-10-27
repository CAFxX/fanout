package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 string; F2 uint }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(fi6)
        var v7 int32 = 0
        var v8 uint16 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 255
    v10 := GenS1{F0: uint16(v9)}
    fmt.Println(v9)
    _ = v10
    _ = v9
    return 0
}
func genFn1(p0 uint64, p1 int16, p2 int32) int32 {
    fmt.Println(p0)
    if (!(true || ("" > "a"))) {
        var v11 uint = 2
        var v12 int32 = -537
        fmt.Println(p2, p2)
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 * p2)
}

func main() {
    v1 := []int{42, -9223372036854775808, 32768}
    v2 := make(map[int]bool)
    v3 := []string{"a"}
    var v4 int8 = 127
    v5 := make(map[string]uint32)
    _ = len(v5)
    v15 := GenS1{F0: 0}
    v4 -= 42
    fmt.Println(v4)
    v4 -= -128
    v4 += (-(v4))
    for fi16 := 0; fi16 < 4; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: 42}
            _ = fi17
            _ = v18
        }
        var v19 float32 = 1.5
        fmt.Println(fi16)
        _ = fi16
        _ = v19
    }
    var v20 uint = 4294967295
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int32 = -275
            var v24 uint = 2
            v25 := v5["foo bar"]
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        fmt.Println(v20)
        v26 := GenS3{F0: (^(v20))}
        _ = fi21
        _ = v26
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := GenS2{F0: 42}
        v20 = (v20 % ((v20) | 1))
        v29 := GenS4{F0: v4, F1: "foo bar", F2: v20}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint = v20
    switch (("" + "\n\t") + "foo bar") {
    case "":
        if true {
            var v31 int32 = 0
            _ = v31
        }
    }
    var v32 int8 = (v4 - (v4 - 0))
    fmt.Println(v20)
    _ = len(v5)
    v33 := GenS2{F0: -128}
    if (127 == 255) {
        if (-9223372036854775808 != 1) {
            var v34 uint = (^(v20))
            for fi35 := 0; fi35 < 7; fi35++ {
                switch 4294967295 {
                case 57:
                    var v36 float64 = 0.0
                    _ = v36
                case 125:
                    v37 := GenS2{F0: v4}
                    _ = v37
                case 630:
                    v38 := GenS3{F0: v30}
                    _ = v38
                }
                v34 = v34
                _ = fi35
            }
            var v39 string = "foo bar"
            _ = v34
            _ = v39
        } else {
            v3 = append(v3, "\n\t")
            switch 32768 {
            case 2147483647:
                if true {
                    if true {
                        var v40 uint32 = 92
                        _ = v40
                    }
                }
            default:
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v4, v30, v20)
                case "a":
                    fmt.Println(v30)
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
