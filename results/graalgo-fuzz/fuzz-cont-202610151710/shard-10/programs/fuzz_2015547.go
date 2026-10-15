package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint; F1 int; F2 int16 }
func genFn0() uint64 {
    var v4 bool = false
    if v4 {
        v5 := []int{}
        switch 42 {
        case 32768:
            if v4 {
                var v6 uint32 = 65535
                v6 += v6
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                _ = v6
            }
        case 1:
            v7 := GenS1{F0: 65535, F1: 194, F2: 937}
            _ = v7
        }
        _ = v5
    }
    var v8 uint32 = 42
    if (!((((42 <= 128) || (!v4)) || (2147483647 >= -601)) || true)) {
        var v9 uint64 = 680
        if true {
            v10 := []int{-436, -233, 1}
            var v11 float32 = -2.25
            _ = v10
            _ = v11
        }
        _ = v9
    }
    switch ("a" + "a") {
    case "":
        v8 += 2147483647
    case "foo bar":
        var v12 uint = 42
        _ = v12
    case "a":
        var v13 int8 = 0
        _ = v13
    default:
        v14 := []int{0}
        _ = v14
    }
    _ = v4
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("a" + "a"), "a", ((("foo bar" + "\n\t") + "") + "a"), "hello"}
    v3 := []int{4294967295, 0, -426, -9223372036854775808}
    var v15 int16 = 1000
    if (-128 <= -128) {
        if (!(!false)) {
            _ = len(v3)
        } else {
            var v16 string = "hello"
            if (9223372036854775807 >= 42) {
                v16 = v16
                v17 := GenS1{F0: 32767, F1: 1, F2: 9223372036854775807}
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (42 >= 42) {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v20 := v1[2]
                            _ = fi19
                            _ = v20
                        }
                        var v21 int64 = 42
                        v16 = v16
                        _ = v21
                    } else {
                        if false {
                            for fi22 := 0; fi22 < 1; fi22++ {
                                fmt.Println(fi22, fi22, fi18)
                                v16 = v16
                                _ = fi22
                            }
                        } else {
                            v23 := GenS1{F0: 497, F1: 196, F2: -1}
                            fmt.Println(fi18, v16, v16)
                            v24 := GenS1{F0: 429, F1: 255, F2: -781}
                            _ = v23
                            _ = v24
                        }
                        v16 = ""
                    }
                    fmt.Println(v15, v16)
                    if (-126 < (fi18 | (fi18 | fi18))) {
                        fmt.Println(v15, v15)
                        var v25 int64 = 9223372036854775807
                        _ = v25
                    } else {
                        v15 = (v15 + v15)
                        var v26 int32 = -927
                        v27 := GenS2{F0: 65535, F1: ((32767 > -920) && (false && false))}
                        _ = v26
                        _ = v27
                    }
                    _ = fi18
                }
                _ = v17
            } else {
                var v28 int32 = 136
                _ = v28
            }
            v16 = (v16 + (("foo bar" + v16) + ((v16 + "foo bar") + "hello")))
            _ = v16
        }
        fmt.Println(v15)
    }
    v15 = 502
    if (!false) {
        v1[42] = 2147483647
    }
    v15 += (v15 | v15)
    var v29 float64 = (1.5 * (1e10 * 1.5))
    v30 := GenS3{F0: 4294967295, F1: 2, F2: 42}
    v15 = v15
    delete(v1, 818)
    for fi31 := 0; fi31 < 7; fi31++ {
        var v32 float64 = 1.5
        v32 = (v32 + -2.25)
        var v33 uint32 = 2147483647
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := v1[746]
    var v35 uint = 1
    fmt.Println(v34, v29, v29)
    var v36 int16 = (v15 / (((v15 - v15)) | 1))
    var v37 int16 = v36
    v1[470] = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
