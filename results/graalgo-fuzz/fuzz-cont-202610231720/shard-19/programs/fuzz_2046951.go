package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 uint32, p1 int64) int {
    var v5 float32 = 1.5
    var v6 int64 = -1
    v7 := GenS1{F0: -627, F1: "a"}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -861
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 bool = (!false)
    v3 := make(map[int]uint32)
    var v4 uint8 = 26
    if v2 {
        var v8 uint32 = 565
        _ = v8
    } else {
        var v9 uint16 = 65535
        v2 = (!v2)
        var v10 int64 = 9223372036854775807
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int64 = 0
        fmt.Println(v2, fi11, fi11)
        var v13 int16 = -32768
        _ = fi11
        _ = v12
        _ = v13
    }
    switch "" {
    case "hello":
        v14 := []int{}
        _ = v14
    case "":
        if true {
            v1 -= (((((float64(0.1) / v1) - 3.14159) - 3.14159) * float64(0.0)) + float64(((1.5 / float32((float64(float32(float64(3.14159))) / float64(1e10)))) - (((1e10 - (float32((-2.25 * v1)) * float32(v1))) / 0.1) - -2.25))))
            var v15 uint8 = 2
            _ = v15
        } else {
            var v16 uint64 = 18446744073709551615
            var v17 int8 = 127
            var v18 float64 = 1e10
            _ = v16
            _ = v17
            _ = v18
        }
    }
    var v19 int8 = 42
    delete(v3, 9223372036854775807)
    v20 := []int{4294967295, 128, 684, -9223372036854775808}
    if (-1 >= 1000) {
        var v21 uint64 = 18446744073709551615
        switch 2147483647 {
        case -448:
            _ = len(v20)
        case -471:
            fmt.Println(v19, v4)
        }
        var v22 int32 = -2147483648
        _ = v21
        _ = v22
    } else {
        v23 := v3[486]
        fmt.Println(v4)
        _ = v23
    }
    if (42 < 1) {
        var v24 uint8 = v4
        for fi25 := 0; fi25 < 4; fi25++ {
            v19 -= -128
            if (65535 != 424) {
                fmt.Println(v2)
                switch "\n\t" {
                case "\n\t":
                    if (("foo bar" + "hello") > "foo bar") {
                        v26 := v3[fi25]
                        _ = v26
                    } else {
                        switch fi25 {
                        case -130:
                            var v27 uint32 = 42
                            _ = v27
                        default:
                            v28 := GenS2{F0: v4}
                            _ = v28
                        }
                        v19 = -81
                    }
                case "foo bar":
                    var v29 int8 = (v19 >> ((891) & 63))
                    _ = v29
                }
            } else {
                if v2 {
                    var v30 uint16 = 2
                    if true {
                        fmt.Println(v19, v2, fi25)
                    }
                    _ = len(v20)
                    _ = v30
                } else {
                    switch ("" + "a") {
                    case "":
                        v31 := GenS3{F0: ""}
                        _ = v31
                    default:
                        var v32 int64 = -1
                        _ = v32
                    }
                    if (("hello" + "foo bar") >= "hello") {
                        switch genFn0(65535, 9223372036854775807) {
                        case -9223372036854775808:
                            if v2 {
                                v1 = v1
                                var v33 uint8 = (v24 | (v4 - v24))
                                fmt.Println(v24, v19, v1)
                                _ = v33
                            }
                        case 128:
                            v34 := v3[genFn0(uint32(v24), 725)]
                            _ = v34
                        }
                        var v35 uint = 18446744073709551615
                        _ = v35
                    } else {
                        delete(v3, (fi25 - genFn0(175, 42)))
                        var v36 bool = true
                        _ = v36
                    }
                    v37 := GenS2{F0: v24}
                    _ = v37
                }
                var v38 uint16 = 658
                var v39 uint16 = v38
                _ = v38
                _ = v39
            }
            _ = len(v20)
            _ = fi25
        }
        _ = v24
    }
    var v40 uint8 = 255
    _ = len(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
}
