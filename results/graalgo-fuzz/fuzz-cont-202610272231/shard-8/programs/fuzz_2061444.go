package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 bool, p1 int16) int8 {
    var v6 int = 127
    var v7 bool = (748 > 200)
    if true {
        v8 := GenS1{F0: p1}
        var v9 int = (v6 ^ 532)
        _ = v8
        _ = v9
    } else {
        v10 := []int{(-(v6))}
        v10 = append(v10, (-(v6)))
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}
func genFn1() uint16 {
    var v11 float64 = 1e10
    v11 -= v11
    if (!(false && false)) {
        if true {
            var v12 uint32 = 4294967295
            _ = v12
        } else {
            var v13 uint32 = 4294967295
            _ = v13
        }
        v11 = v11
    } else {
        v14 := GenS1{F0: 42}
        v15 := []int{2147483647}
        for fi16 := 0; fi16 < 3; fi16++ {
            if (false && false) {
                var v17 int64 = 9223372036854775807
                v11 -= (1e10 - ((-2.25 + (1e10 * float64(1e10))) / v11))
                if (127 <= 42) {
                    var v18 bool = false
                    if v18 {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v11 -= v11
                            v20 := GenS1{F0: -32768}
                            fmt.Println(v11)
                            v11 -= ((v11 + v11) + v11)
                            v18 = true
                            _ = fi19
                            _ = v20
                        }
                        v21 := GenS1{F0: -1}
                        switch "a" {
                        case "foo bar":
                            v11 += (0.1 / float64(1e10))
                        case "":
                            v22 := GenS2{F0: 482, F1: -32768}
                            _ = v22
                        default:
                            switch (fi16 % (((fi16 - 9223372036854775807)) | 1)) {
                            case 32768:
                                var v23 float32 = -2.25
                                _ = v23
                            case 127:
                                var v24 int32 = 797
                                _ = v24
                            case 682:
                                fmt.Println(v17, v18, v17)
                            }
                        }
                        _ = v21
                    }
                    _ = v18
                } else {
                    v25 := GenS3{F0: ("" + ("" + ""))}
                    v26 := GenS2{F0: 255, F1: -140}
                    switch ("hello" + "foo bar") {
                    case "a":
                        v11 -= 1e10
                    case "":
                        _ = len(v15)
                    case "\n\t":
                        v27 := GenS1{F0: 32767}
                        _ = v27
                    }
                    _ = v25
                    _ = v26
                }
                _ = v17
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    }
    v28 := GenS1{F0: -623}
    _ = v11
    _ = v28
    return 948
}

func main() {
    v1 := []string{"a", "hello"}
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    var v4 int8 = 42
    v5 := make(map[int]uint32)
    v29 := GenS2{F0: 255, F1: -1}
    delete(v2, "")
    var v30 int64 = 0
    fmt.Println(v3, v4)
    if ((741 != 2) && (-402 <= -1)) {
        v3 -= v3
        v31 := GenS3{F0: ""}
        _ = v31
    } else {
        v4 = genFn0(false, int16(v30))
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS2{F0: uint16(v3), F1: 0}
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi32
        _ = v33
    }
    fmt.Println(v3, v30, v3)
    v30 += v30
    if (0.1 <= -2.25) {
        var v34 uint8 = 129
        fmt.Println(v30)
        _ = v34
    }
    v30 = v30
    var v35 uint = 4294967295
    var v36 uint = (v35 | v35)
    var v37 uint8 = 127
    for fi38 := 0; fi38 < 7; fi38++ {
        v37 = 217
        var v39 int = fi38
        _ = fi38
        _ = v39
    }
    v40 := GenS3{F0: (((("hello" + "a") + "foo bar") + "\n\t") + "hello")}
    v35 = uint(v37)
    v30 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
