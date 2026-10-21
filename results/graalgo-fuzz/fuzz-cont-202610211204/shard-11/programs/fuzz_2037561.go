package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32; F1 int16; F2 uint }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    var v1 int16 = 0
    var v2 float64 = 1.5
    var v3 int64 = 447
    var v4 string = "foo bar"
    var v5 uint16 = 1
    v6 := make(map[string]bool)
    switch (v4 + v4) {
    case "":
        if true {
            fmt.Println(v2)
            var v7 string = v4
            _ = v7
        }
    case "a":
        v2 = 0.0
    case "foo bar":
        if false {
            fmt.Println(v5, v3)
        } else {
            v2 = float64(float32(((float64(1.5) / (v2 * 3.14159)) / ((v2 / v2) / float64(3.14159)))))
            if (false && false) {
                if (0.0 <= float32((float64(1e10) * 3.14159))) {
                    v4 = (v4 + "a")
                    var v8 int8 = 127
                    switch 2147483647 {
                    case 1:
                        v8 -= (v8 ^ v8)
                    case 32767:
                        if false {
                            var v9 int64 = -1
                            v5 = v5
                            _ = v9
                        }
                    case 904:
                        v10 := []int{9223372036854775807, 32767, int(v3)}
                        _ = v10
                    }
                    _ = v8
                } else {
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 int = fi11
                        var v13 int = (fi11 >> ((308) & 63))
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    if (!true) {
                        v2 -= (v2 * v2)
                        var v14 float64 = (float64((float32(float64((0.1 + float32(((1e10 + v2) + ((v2 + (v2 / (0.0 + float64(float32(((float64((0.0 - 1e10)) * v2) / v2)))))) * (3.14159 / (-2.25 / v2)))))))) - float32(0.1))) - v2)
                        v3 -= v3
                        _ = v14
                    } else {
                        var v15 int32 = -2147483648
                        if (!false) {
                            for fi16 := 0; fi16 < 11; fi16++ {
                                var v17 uint8 = 116
                                if false {
                                    var v18 uint64 = 18446744073709551615
                                    v19 := GenS1{F0: 127}
                                    _ = v18
                                    _ = v19
                                }
                                _ = fi16
                                _ = v17
                            }
                            if (!((v4 + "a") <= "foo bar")) {
                                var v20 int8 = 1
                                _ = v20
                            } else {
                                var v21 int = 9223372036854775807
                                _ = v21
                            }
                        }
                        _ = v15
                    }
                }
                fmt.Println(v2)
            } else {
                var v22 int16 = -1
                _ = v22
            }
        }
    }
    var v23 uint32 = 65535
    if false {
        fmt.Println(v1, v23)
    }
    if (false && false) {
        fmt.Println(v1)
        var v24 uint8 = 42
        if true {
            v25 := GenS1{F0: 127}
            _ = v25
        }
        _ = v24
    }
    if ((177 <= 234) || false) {
        fmt.Println(v1, v23)
    }
    v26 := GenS1{F0: 70}
    v27 := []int{630, -132, 4294967295, -1}
    switch "" {
    case "\n\t":
        if (-128 < 100) {
            if ((false && ((!(!true)) || true)) || (255 != 2)) {
                if true {
                    _ = len(v27)
                }
                v6[v4] = (127 > 136)
                switch ("" + (v4 + v4)) {
                case "":
                    v5 -= (^(v5))
                case "a":
                    var v28 uint16 = (^(v5))
                    _ = v28
                case "hello":
                    if false {
                        v1 += (v1 + -32768)
                        v4 = v4
                        if false {
                            var v29 uint64 = 103
                            _ = len(v27)
                            _ = v29
                        }
                    } else {
                        fmt.Println(v2)
                        var v30 int8 = 127
                        v4 = "\n\t"
                        _ = v30
                    }
                default:
                    fmt.Println(v3)
                }
            }
        }
    case "a":
        v3 += v3
    default:
        v31 := GenS1{F0: 127}
        _ = v31
    }
    switch -9223372036854775808 {
    case 32767:
        if len(v27) > 0 {
            v27[0] = -638
        }
    case 2147483647:
        v5 -= 32767
    default:
        v23 -= (v23 >> ((398) & 63))
    }
    v32 := GenS2{F0: false, F1: -863}
    var v33 uint32 = 2147483647
    v34 := v6[(v4 + v4)]
    v35 := GenS3{F0: -1, F1: v1, F2: 2}
    v34 = (v33 == (^(v33)))
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 int32 = int32(v3)
        if (((v4 + (v4 + v4)) + ((v4 + "hello") + v4)) > "hello") {
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            v38 := GenS4{F0: (v37 | (v37 | -2147483648)), F1: (v37 % (((-(v37))) | 1))}
            v37 = -1
            _ = v38
        }
        v37 -= 607
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
