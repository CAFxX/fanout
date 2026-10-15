package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0(p0 int64, p1 uint64) int16 {
    switch -9223372036854775808 {
    case -127:
        v4 := []int{128, 127, -150}
        _ = v4
    }
    p1 = 18446744073709551615
    var v5 uint32 = 1
    fmt.Println(p1)
    v6 := []int{413, 127, 540}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return int16(v5)
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]int)
    var v3 float64 = v1
    v1 = v1
    var v7 bool = (!((("" + "a") + "hello") > "foo bar"))
    v1 += 1.5
    var v8 uint = 255
    for fi9 := 0; fi9 < 4; fi9++ {
        v3 -= ((3.14159 / (v3 / -2.25)) / v3)
        _ = fi9
    }
    var v10 uint16 = 65535
    v11 := []int{1, 128}
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint16 = 32767
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v8, v3, v1)
        v3 = (v1 / (v3 + (v1 + ((float64(0.0) + (float64((1e10 * (1.5 + 0.1))) + ((v3 + v1) * v1))) / v1))))
        _ = fi14
    }
    if (!false) {
        v15 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v8 = v8
        _ = fi16
    }
    switch 4294967295 {
    case 32768:
        v8 = (^(v8))
    case -791:
        v7 = (903 <= 814)
    }
    fmt.Println(v1)
    switch (("\n\t" + ("foo bar" + ("hello" + "hello"))) + "foo bar") {
    case "a":
        var v17 int = 886
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 17; fi18++ {
            v7 = (-993 != 1)
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(v1, v1)
                switch "\n\t" {
                case "":
                    var v20 bool = (((2 < 1) || false) && (!v7))
                    _ = v20
                default:
                    var v21 string = "foo bar"
                    _ = v21
                }
                v1 += (v1 / 1.0)
                _ = fi19
            }
            if (!true) {
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int = fi18
                    var v24 int = (fi18 >> ((65535) & 63))
                    if (true || (v7 || true)) {
                        v25 := GenS2{F0: uint8(v8), F1: 1}
                        if (!v7) {
                            v3 = ((v3 - v1) * 0.0)
                        } else {
                            if v7 {
                                var v26 int = v23
                                var v27 bool = (!v7)
                                var v28 float64 = 0.1
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                if false {
                                    v29 := GenS2{F0: 42, F1: -604}
                                    v23 = fi18
                                    _ = v29
                                }
                                var v30 int64 = 0
                                v31 := v2[(fi18 ^ (v23 | fi18))]
                                _ = v30
                                _ = v31
                            }
                        }
                        if v7 {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            v10 = v10
                        }
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = fi18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
