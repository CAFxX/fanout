package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16, p1 uint, p2 int16) int64 {
    var v6 int16 = (p2 >> (((p1 | 474)) & 63))
    if true {
        v7 := GenS1{F0: false, F1: p2, F2: "hello"}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -791
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (1 <= 65535)
    v3 := make(map[uint32]string)
    var v4 uint64 = 1
    var v5 uint32 = 677
    fmt.Println(v2)
    if false {
        v8 := GenS2{F0: -416}
        if ((1e10 / float64(float32((1e10 + 3.14159)))) != 1.5) {
            if false {
                fmt.Println(v2, v2, v2)
                fmt.Println(v4, v5, v2)
            }
        } else {
            for fi9 := 0; fi9 < 19; fi9++ {
                v10 := GenS2{F0: -9223372036854775808}
                var v11 int32 = 0
                if true {
                    switch "hello" {
                    case "a":
                        _ = len(v1)
                    case "":
                        v2 = ((v11 != v11) && ((((1e10 * ((1e10 / 1e10) / 1.0)) + -2.25) != (1e10 * 1.5)) && v2))
                    }
                    var v12 uint16 = 32767
                    _ = v12
                } else {
                    var v13 bool = (!true)
                    _ = v13
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            if true {
                if (1 < 2) {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 bool = v2
                        v16 := GenS2{F0: 9223372036854775807}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 18; fi17++ {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v19 := v1[v5]
                            switch (((("\n\t" + "") + (("" + "hello") + "\n\t")) + "") + (("hello" + ("\n\t" + "foo bar")) + "")) {
                            case "hello":
                                var v20 bool = true
                                _ = v20
                            case "foo bar":
                                var v21 uint16 = 32767
                                _ = v21
                            case "a":
                                v22 := GenS3{F0: fi17, F1: 127}
                                _ = v22
                            default:
                                var v23 bool = ((v5 | 1) >= 2147483647)
                                _ = v23
                            }
                            _ = fi18
                            _ = v19
                        }
                        fmt.Println(fi17, v4, v2)
                        switch fi17 {
                        case -1:
                            v4 = v4
                        case -9223372036854775808:
                            for fi24 := 0; fi24 < 6; fi24++ {
                                v4 -= v4
                                v25 := []int{-1, (fi17 * 42)}
                                _ = fi24
                                _ = v25
                            }
                        }
                        _ = fi17
                    }
                }
                var v26 uint16 = 255
                _ = v26
            }
        }
        var v27 string = ("a" + "foo bar")
        _ = v27
        _ = v8
    }
    if ((100 <= -128) && v2) {
        if v2 {
            if ((!(403 <= 0)) || false) {
                _ = len(v3)
                v28 := []int{32767}
                _ = v28
            } else {
                var v29 uint16 = 42
                v30 := GenS4{F0: 1, F1: v5}
                v31 := GenS3{F0: -2147483648, F1: 127}
                _ = v29
                _ = v30
                _ = v31
            }
        }
    }
    v32 := []int{}
    v4 -= 183
    _ = len(v1)
    v33 := v1[(v5 / (((v5 & v5)) | 1))]
    v33 = (v2 || v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
