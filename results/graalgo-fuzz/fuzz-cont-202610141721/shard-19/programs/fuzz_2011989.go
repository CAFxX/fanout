package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 bool) uint8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = p0
        var v6 uint32 = 1
        _ = fi5
        _ = v6
    }
    var v7 int16 = 1
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 128
}
func genFn1() uint8 {
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint8 = 128
        fmt.Println(fi9, v10)
        var v11 int8 = -9
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 bool = (true && (9223372036854775807 < -1))
    switch (("foo bar" + "\n\t") + "") {
    case "":
        v13 := GenS1{F0: -2147483648}
        _ = v13
    case "hello":
        var v14 int16 = 32767
        _ = v14
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    var v3 float64 = ((1e10 / 1e10) - (1.5 * 1e10))
    v4 := []uint32{0, 2147483647}
    switch "" {
    case "a":
        v2 = v2
    case "":
        var v15 int32 = -41
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v2, v3)
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    v1[(-(fi17))] = (v2 < v2)
                    _ = fi17
                }
            }
            _ = fi16
        }
    }
    var v18 float64 = v3
    v19 := v1[702]
    fmt.Println(v2, v19)
    v18 = float64((0.0 * 3.14159))
    if len(v4) > 0 {
        v4[0] = 65535
    }
    switch (("\n\t" + v2) + v2) {
    case "\n\t":
        var v20 int32 = 42
        _ = v20
    default:
        var v21 uint = 0
        _ = v21
    }
    v22 := GenS1{F0: 155}
    var v23 float32 = (1e10 - ((0.1 * 1.5) / ((((float32(v3) * (0.1 * -2.25)) / (3.14159 - (float32(v3) * 3.14159))) / (-2.25 / float32(v3))) * (-2.25 * 0.1))))
    v3 -= ((float64((float32(0.0) / (v23 - v23))) + (v18 - v18)) + v18)
    var v24 int32 = -1
    for fi25 := 0; fi25 < 20; fi25++ {
        v26 := GenS1{F0: v24}
        if (v19 || ((4294967295 == 755) || v19)) {
            if (false && (!(65535 == 0))) {
                if (42 < 42) {
                    if v19 {
                        if v19 {
                            for fi27 := 0; fi27 < 5; fi27++ {
                                v28 := GenS1{F0: (v24 ^ (^(v24)))}
                                var v29 uint64 = 0
                                v29 = (v29 / ((4294967295) | 1))
                                v30 := GenS2{F0: ((((v2 + v2) + (v2 + (v2 + v2))) + ((v2 + v2) + (((v2 + "") + ((v2 + v2) + v2)) + v2))) + "a")}
                                if true {
                                    var v31 uint64 = 42
                                    _ = v31
                                }
                                _ = fi27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        }
                        v4 = append(v4, 42)
                    } else {
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v24 -= int32(fi32)
                            var v33 float32 = v23
                            fmt.Println(fi32)
                            if (42 == 42) {
                                var v34 uint8 = 1
                                v35 := GenS3{F0: (^(v34)), F1: fi32}
                                _ = v34
                                _ = v35
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                    v23 = v23
                    var v36 uint16 = 456
                    _ = v36
                } else {
                    var v37 int64 = 42
                    _ = v37
                }
                v2 = ("hello" + "hello")
            }
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
