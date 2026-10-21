package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0() uint16 {
    v5 := []int{32768, 2147483647}
    var v6 string = ("" + ("hello" + (("a" + "hello") + ("\n\t" + "hello"))))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v7 int32 = -1
    var v8 uint32 = 932
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 696
}
func genFn1(p0 uint, p1 int16, p2 uint32) int8 {
    var v9 int8 = -128
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS1{F0: 0}
        var v12 float32 = (1e10 / float32(-2.25))
        if (848 < 65535) {
            p0 -= (p0 - (p0 / ((p0) | 1)))
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := []int{(^(fi13)), fi10, fi10}
                _ = fi13
                _ = v14
            }
            fmt.Println(v12, v12, fi10)
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (v9 - -60)
}

func main() {
    var v1 int32 = 892
    v2 := make(map[uint32]string)
    var v3 int16 = -921
    v4 := []uint32{}
    switch ("a" + "a") {
    case "\n\t":
        v2[0] = "a"
    case "a":
        v2[1] = ""
    case "":
        if (true && false) {
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS2{F0: 42}
                delete(v2, 0)
                v1 = v1
                _ = fi15
                _ = v16
            }
            _ = len(v4)
            if true {
                if true {
                    var v17 int8 = 1
                    v18 := GenS3{F0: 255, F1: 0}
                    _ = v17
                    _ = v18
                } else {
                    v1 = (v1 | -508)
                    v3 -= 637
                    if true {
                        var v19 int16 = (v3 - v3)
                        _ = len(v4)
                        _ = v19
                    } else {
                        _ = len(v4)
                        var v20 int32 = (v1 << ((4294967295) & 63))
                        _ = len(v4)
                        _ = v20
                    }
                }
                switch ("a" + ("a" + ("hello" + (("foo bar" + ("hello" + "hello")) + "\n\t")))) {
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    for fi21 := 0; fi21 < 4; fi21++ {
                        if (-128 < genFn1(316, v3, 65535)) {
                            v22 := GenS2{F0: 127}
                            _ = len(v2)
                            switch fi21 {
                            case -2147483648:
                                fmt.Println(fi21, v3)
                            case -596:
                                var v23 string = "foo bar"
                                _ = v23
                            case 663:
                                fmt.Println(v3, fi21)
                            }
                            _ = v22
                        }
                        _ = fi21
                    }
                case "":
                    _ = len(v4)
                }
            } else {
                fmt.Println(v1, v3)
            }
        }
    }
    if (67 <= v3) {
        var v24 int = 2147483647
        v25 := v2[0]
        _ = v24
        _ = v25
    }
    var v26 string = ((("\n\t" + "foo bar") + ("hello" + "a")) + (("a" + "a") + ("hello" + "a")))
    v27 := GenS1{F0: 9223372036854775807}
    fmt.Println(v3, v1)
    v1 += -147
    v4 = append(v4, 2)
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := v2[2147483648]
        _ = fi28
        _ = v29
    }
    fmt.Println(v1, v26, v1)
    v1 -= (v1 + -1)
    if (!true) {
        _ = len(v4)
        var v30 float64 = float64(float32(0.1))
        _ = v30
    }
    var v31 float64 = 1e10
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
}
