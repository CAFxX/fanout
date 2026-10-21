package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 int64) int8 {
    var v7 int = 42
    v8 := GenS1{F0: 1, F1: 0}
    if (((p1 % ((0) | 1)) != p0) && true) {
        p0 = (p1 % ((1) | 1))
        v7 += v7
        if false {
            var v9 uint = 2
            if false {
                if false {
                    var v10 uint8 = 1
                    _ = v10
                }
            }
            p0 = p1
            _ = v9
        }
    } else {
        v11 := GenS1{F0: 457, F1: 128}
        var v12 uint = 0
        var v13 float64 = ((((3.14159 - (((1.5 + 1e10) - 1.5) + 3.14159)) * ((0.0 - 1e10) * 3.14159)) / (3.14159 - (1e10 / 1.5))) + (-2.25 - 3.14159))
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 62, F1: 100}
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    _ = v8
    return -102
}

func main() {
    var v1 bool = false
    var v2 uint32 = 2
    var v3 uint8 = 234
    v4 := []string{"\n\t", "", "foo bar", ("" + ("\n\t" + ("a" + ("a" + ""))))}
    v5 := []int{4294967295}
    v6 := make(map[uint32]uint32)
    for fi15 := 0; fi15 < 10; fi15++ {
        v4 = append(v4, "foo bar")
        v2 = v2
        switch 9223372036854775807 {
        case 127:
            if (!(fi15 != (-(fi15)))) {
                v16 := GenS1{F0: 65535, F1: fi15}
                _ = v16
            }
        case 4294967295:
            var v17 float32 = 0.0
            _ = v17
        case 0:
            _ = len(v4)
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 = v2
        fmt.Println(v2, v3, v2)
        _ = fi18
    }
    v1 = (false && (("" + "foo bar") != "hello"))
    v3 = (v3 + 198)
    fmt.Println(v2, v2, v3)
    var v19 uint32 = v2
    if v1 {
        v1 = (2 > 1)
        switch "foo bar" {
        case "hello":
            if (156 != 2) {
                fmt.Println(v1, v19)
            }
        }
        if ("foo bar" <= "foo bar") {
            for fi20 := 0; fi20 < 17; fi20++ {
                v1 = (1 <= 0)
                if (((42 == 0) && v1) && false) {
                    v1 = false
                    for fi21 := 0; fi21 < 10; fi21++ {
                        if false {
                            if v1 {
                                v3 -= 1
                                v5 = append(v5, -9223372036854775808)
                                v19 -= 1
                            } else {
                                v22 := GenS3{F0: 596}
                                _ = v22
                            }
                            if (!v1) {
                                var v23 uint = 423
                                _ = v23
                            }
                        }
                        _ = fi21
                    }
                }
                var v24 uint64 = 18446744073709551615
                _ = fi20
                _ = v24
            }
            v6[v2] = v19
            v3 -= v3
        } else {
            v4 = append(v4, "foo bar")
            var v25 uint64 = 2
            v19 = v2
            _ = v25
        }
    }
    fmt.Println(v3, v3)
    if true {
        if (v2 <= (^(v2))) {
            if (-1 == 0) {
                var v26 uint16 = 32767
                switch ("" + "") {
                case "foo bar":
                    switch ("" + ("" + ("\n\t" + (("\n\t" + "foo bar") + "")))) {
                    case "\n\t":
                        var v27 int64 = -1
                        _ = v27
                    case "foo bar":
                        if (v1 || (255 >= v26)) {
                            var v28 uint8 = (v3 << ((1) & 63))
                            for fi29 := 0; fi29 < 3; fi29++ {
                                v2 = v19
                                _ = fi29
                            }
                            _ = v28
                        }
                    default:
                        v19 = (^(v19))
                    }
                }
                _ = v26
            }
        }
        v19 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
