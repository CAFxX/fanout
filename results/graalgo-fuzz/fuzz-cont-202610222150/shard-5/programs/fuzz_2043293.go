package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }
func genFn0() uint32 {
    var v4 uint64 = 18446744073709551615
    var v5 float64 = (0.0 * 1e10)
    fmt.Println(v5, v4)
    _ = v4
    _ = v5
    return 926
}

func main() {
    v1 := []int{}
    var v2 float64 = (-2.25 / ((0.1 + -2.25) + (1e10 + float64((float32(1.5) - 0.0)))))
    v3 := make(map[string]string)
    v3["\n\t"] = "hello"
    for fi6 := 0; fi6 < 2; fi6++ {
        switch ("hello" + "a") {
        case "a":
            var v7 bool = (1.5 != (1e10 / (float32((float64(0.0) / v2)) / float32(float64((1e10 * 1.5))))))
            _ = v7
        case "foo bar":
            v2 += 1.5
        case "\n\t":
            v2 = v2
        default:
            if false {
                fmt.Println(v2)
            }
        }
        var v8 uint16 = 42
        var v9 int16 = 42
        _ = fi6
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    switch 32767 {
    case -1:
        if false {
            if (!(42 > 32767)) {
                v2 = 0.0
                if (0 != 796) {
                    var v10 uint = 65535
                    fmt.Println(v10, v2)
                    _ = v10
                }
                var v11 uint32 = 65535
                _ = v11
            }
            if (42 == 100) {
                v12 := GenS1{F0: 1}
                _ = v12
            } else {
                v13 := GenS1{F0: -554}
                _ = v13
            }
            if (true || ((((0.0 * v2) + v2) + v2) == (3.14159 + ((float64(1e10) + ((v2 / 1e10) / v2)) * (v2 * (float64(-2.25) + float64(1e10))))))) {
                fmt.Println(v2)
                v14 := GenS1{F0: 42}
                _ = v14
            } else {
                v2 = v2
                if (18446744073709551615 <= 65535) {
                    if (true || ((1 < 42) && (!false))) {
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v2 += float64(3.14159)
                            v2 -= ((0.1 * (float64((float32(v2) / 3.14159)) - ((v2 * v2) / (-2.25 / v2)))) + (v2 / v2))
                            v16 := GenS2{F0: 2147483647}
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        if false {
                            v17 := GenS1{F0: -374}
                            _ = v17
                        } else {
                            var v18 int8 = 127
                            v19 := GenS3{F0: 726, F1: 9223372036854775807, F2: "\n\t"}
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS3{F0: 32768, F1: 423, F2: "a"}
                        var v21 uint32 = 408
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if (0 < 100) {
                            v23 := GenS3{F0: 32767, F1: 0, F2: "a"}
                            _ = v23
                        } else {
                            var v24 int32 = 42
                            var v25 int32 = v24
                            _ = v24
                            _ = v25
                        }
                        _ = fi22
                    }
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 uint64 = 18446744073709551615
                        if (true && (!false)) {
                            v2 = v2
                        }
                        if false {
                            var v28 int16 = 0
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                } else {
                    switch (("" + "a") + "a") {
                    case "a":
                        v2 += (v2 + 3.14159)
                    case "foo bar":
                        v2 -= (v2 + ((1.5 - (-2.25 - (v2 / 3.14159))) + 3.14159))
                    case "hello":
                        fmt.Println(v2)
                    }
                }
                v2 = v2
            }
        }
    }
    var v29 uint32 = 4294967295
    v30 := GenS4{F0: 2, F1: 0, F2: 2}
    _ = len(v3)
    if false {
        var v31 uint64 = 458
        v31 = v31
        var v32 uint64 = v31
        _ = v31
        _ = v32
    } else {
        v3[(("" + ("a" + ("foo bar" + "foo bar"))) + "a")] = ""
    }
    var v33 int8 = -128
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v1)
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := v3["a"]
            var v37 float64 = (v2 / (v2 * v2))
            fmt.Println(v36, v36)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    switch 511 {
    case 42:
        _ = len(v3)
    case 128:
        var v38 uint = 0
        _ = v38
    default:
        for fi39 := 0; fi39 < 18; fi39++ {
            if len(v1) > 0 {
                v1[0] = int(v33)
            }
            v33 = 127
            _ = fi39
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
}
