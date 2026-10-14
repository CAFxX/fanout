package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 int64, p2 int) string {
    var v5 bool = false
    var v6 float32 = (float32((-2.25 - 0.0)) * -2.25)
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 4294967295}
        var v9 string = ("foo bar" + ((((("a" + "") + "hello") + "a") + ("hello" + "foo bar")) + ""))
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = (!v5)
    switch (^(p2)) {
    case -9223372036854775808:
        var v11 float32 = v6
        _ = v11
    case 216:
        v5 = false
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]int)
    var v3 int8 = (v1 ^ -128)
    var v4 bool = (!false)
    if (886 >= 2147483647) {
        var v12 int32 = -598
        var v13 float32 = float32((3.14159 + 3.14159))
        _ = v12
        _ = v13
    } else {
        var v14 int32 = 985
        v4 = (1 >= 32767)
        _ = v14
    }
    v4 = v4
    v15 := []int{9223372036854775807, 4294967295, -588}
    if (-2147483648 != 1) {
        v16 := GenS1{F0: 161}
        _ = v16
    } else {
        var v17 bool = v4
        var v18 int32 = -2147483648
        _ = v17
        _ = v18
    }
    fmt.Println(v3, v1, v3)
    v19 := GenS1{F0: 65535}
    v20 := GenS1{F0: 0}
    v15 = append(v15, 4294967295)
    fmt.Println(v3)
    var v21 int8 = v1
    switch "" {
    case "foo bar":
        switch 128 {
        case 613:
            v22 := GenS1{F0: 0}
            _ = v22
        case 198:
            v4 = (37 < 2)
        case 920:
            _ = len(v2)
        }
    default:
        v23 := GenS2{F0: 42}
        _ = v23
    }
    v24 := GenS3{F0: 255, F1: "hello", F2: 255}
    switch "hello" {
    case "":
        if (0 == 307) {
            if (-1 != 345) {
                var v25 uint = 18446744073709551615
                var v26 string = ""
                _ = v25
                _ = v26
            }
            fmt.Println(v1, v21, v4)
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 int64 = 9223372036854775807
                v28 -= 42
                _ = fi27
                _ = v28
            }
        } else {
            fmt.Println(v1, v1)
            v21 = v21
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := GenS2{F0: 1000}
                fmt.Println(v21)
                switch ("a" + "foo bar") {
                case "":
                    if len(v15) > 0 {
                        v15[0] = fi29
                    }
                case "a":
                    if v4 {
                        var v31 int8 = (v1 << ((1) & 63))
                        _ = v31
                    } else {
                        v3 -= v21
                        if v4 {
                            v4 = (0.0 <= (1.5 * 3.14159))
                            var v32 int64 = -9223372036854775808
                            _ = v32
                        } else {
                            v33 := GenS2{F0: 32767}
                            _ = v33
                        }
                        if (!false) {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                switch "a" {
                                case "hello":
                                    v35 := GenS3{F0: 0, F1: "foo bar", F2: 255}
                                    _ = v35
                                case "\n\t":
                                    fmt.Println(fi29)
                                case "foo bar":
                                    var v36 int = (fi34 - (fi29 ^ 42))
                                    _ = v36
                                }
                                var v37 float64 = -2.25
                                _ = fi34
                                _ = v37
                            }
                        } else {
                            for fi38 := 0; fi38 < 19; fi38++ {
                                _ = len(v15)
                                fmt.Println(fi38)
                                _ = fi38
                            }
                            var v39 float32 = ((1e10 * (float32(float64(0.1)) * (((0.1 + 3.14159) / (1e10 + (1e10 / float32(1e10)))) * -2.25))) / 1e10)
                            _ = v39
                        }
                    }
                case "hello":
                    v1 += (v3 << ((1) & 63))
                default:
                    var v40 uint = 1
                    _ = v40
                }
                _ = fi29
                _ = v30
            }
        }
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 uint8 = 127
        _ = fi41
        _ = v42
    }
    var v43 int8 = (v21 - v21)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v43
}
