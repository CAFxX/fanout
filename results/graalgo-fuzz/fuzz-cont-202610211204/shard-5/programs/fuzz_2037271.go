package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
func genFn0(p0 int8, p1 int64, p2 uint) int32 {
    if false {
        if false {
            v6 := []int{9223372036854775807, 42, 718, -297}
            p1 -= 42
            _ = v6
        }
        v7 := GenS1{F0: 0}
        _ = v7
    }
    if false {
        v8 := GenS2{F0: 2}
        for fi9 := 0; fi9 < 6; fi9++ {
            switch ("foo bar" + "hello") {
            case "a":
                if false {
                    v10 := []int{(fi9 + fi9), (-(fi9)), (fi9 | (-(fi9))), (fi9 - (fi9 & (fi9 >> ((p2) & 63))))}
                    fmt.Println(p2, p0, p1)
                    _ = v10
                } else {
                    v11 := []int{}
                    _ = v11
                }
            case "\n\t":
                for fi12 := 0; fi12 < 20; fi12++ {
                    fmt.Println(p2)
                    _ = fi12
                }
            case "hello":
                switch "a" {
                case "\n\t":
                    p0 += (p0 / ((59) | 1))
                case "":
                    if false {
                        for fi13 := 0; fi13 < 2; fi13++ {
                            switch "" {
                            case "foo bar":
                                var v14 int16 = -32768
                                _ = v14
                            }
                            var v15 int8 = p0
                            _ = fi13
                            _ = v15
                        }
                        v16 := []int{(fi9 * (fi9 * (fi9 + fi9))), fi9, -741}
                        fmt.Println(fi9)
                        _ = v16
                    }
                }
            }
            if (1000 > 42) {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS3{F0: 255, F1: (^(fi17)), F2: (p0 / (((p0 - p0)) | 1))}
                    v19 := []int{}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint8 = 45
                _ = v20
            }
            fmt.Println(fi9)
            _ = fi9
        }
        _ = v8
    }
    switch int(p1) {
    case 32767:
        fmt.Println(p1, p0)
    case 9223372036854775807:
        p1 += p1
    default:
        p0 = 103
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    v5 := make(map[string]bool)
    for fi21 := 0; fi21 < 17; fi21++ {
        delete(v3, 829)
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 string = ""
            var v24 uint16 = 255
            _ = len(v1)
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi21
    }
    v25 := v1["a"]
    v26 := []int{127}
    v27 := GenS1{F0: 0}
    var v28 bool = (3.14159 > -2.25)
    v2 += v2
    _ = len(v26)
    switch 0 {
    case 127:
        var v29 uint32 = 887
        _ = v29
    case -787:
        fmt.Println(v2, v25)
    default:
        var v30 int32 = genFn0(12, 0, 196)
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if true {
            v32 := GenS3{F0: v2, F1: 349, F2: -128}
            v33 := GenS3{F0: v2, F1: fi31, F2: 4}
            _ = v32
            _ = v33
        } else {
            v34 := GenS2{F0: 0}
            v25 = v25
            _ = v34
        }
        v5["\n\t"] = (!false)
        v2 = v2
        _ = fi31
    }
    delete(v1, ("a" + "a"))
    fmt.Println(v25, v28, v25)
    var v35 float32 = 3.14159
    for fi36 := 0; fi36 < 7; fi36++ {
        v3[(v25 + v25)] = "\n\t"
        if v28 {
            var v37 int32 = -194
            _ = v37
        } else {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
        }
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := GenS3{F0: 1, F1: fi36, F2: -128}
            v2 -= 1
            _ = fi38
            _ = v39
        }
        _ = fi36
    }
    v35 = (0.0 - -2.25)
    _ = len(v3)
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 int = -9223372036854775808
        var v42 bool = false
        var v43 uint32 = (v25 - v25)
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    if len(v26) > 0 {
        v26[0] = 9223372036854775807
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
