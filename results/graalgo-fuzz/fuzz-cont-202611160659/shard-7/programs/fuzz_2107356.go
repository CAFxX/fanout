package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0() uint8 {
    v5 := []int{-55, 32768, 127, 32768}
    var v6 string = ((("foo bar" + ("hello" + "foo bar")) + "foo bar") + "hello")
    var v7 int32 = 2147483647
    var v8 uint = 1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 uint32, p1 int64) uint8 {
    v9 := []int{0, 2147483647, 127}
    p1 = p1
    _ = p0
    _ = p1
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 65535
    var v3 bool = ((!false) || false)
    var v4 uint = 851
    var v10 int = 32767
    v10 = (v10 + -651)
    v1["a"] = uint32(v4)
    var v11 bool = v3
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{-610, (fi12 ^ (^(fi12))), fi12}
        v11 = false
        _ = len(v13)
        _ = fi12
        _ = v13
    }
    if (!(false || (!v3))) {
        fmt.Println(v11, v2)
    }
    v3 = (-2.25 >= 1.5)
    switch "hello" {
    case "a":
        fmt.Println(v4, v11, v11)
    case "foo bar":
        fmt.Println(v10, v10)
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := []int{fi14, fi14, 9223372036854775807}
        if (!v11) {
            v16 := GenS1{F0: 127}
            if (42 <= 9223372036854775807) {
                switch ("hello" + "a") {
                case "\n\t":
                    if (9223372036854775807 != 270) {
                        v2 = 717
                    } else {
                        if len(v15) > 0 {
                            v15[0] = -1
                        }
                        v4 = 0
                        v17 := GenS1{F0: int(v2)}
                        _ = v17
                    }
                case "foo bar":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v3 = v11
                        v11 = (843 >= 1)
                        _ = fi18
                    }
                case "":
                    fmt.Println(v2, v2, v2)
                default:
                    v19 := GenS1{F0: v10}
                    _ = v19
                }
                fmt.Println(v3, v3, v4)
            } else {
                for fi20 := 0; fi20 < 2; fi20++ {
                    _ = len(v1)
                    var v21 int16 = -1
                    _ = fi20
                    _ = v21
                }
                switch (fi14 ^ v10) {
                case -853:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        v23 := GenS2{F0: 0, F1: 2}
                        var v24 uint32 = 192
                        if (241 > 1) {
                            v25 := GenS2{F0: -266, F1: 65535}
                            v26 := v1["\n\t"]
                            v4 += (v4 * v4)
                            _ = v25
                            _ = v26
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    var v27 uint64 = 18446744073709551615
                    _ = v27
                }
                fmt.Println(v10, v11, v2)
            }
            var v28 string = (("hello" + "") + "hello")
            _ = v16
            _ = v28
        } else {
            v15 = append(v15, v10)
        }
        _ = fi14
        _ = v15
    }
    if (!(!((1 < -1) || (true || false)))) {
        for fi29 := 0; fi29 < 7; fi29++ {
            fmt.Println(v4, v10)
            _ = fi29
        }
        _ = len(v1)
    } else {
        v11 = v11
        if v11 {
            for fi30 := 0; fi30 < 17; fi30++ {
                switch 32767 {
                case 128:
                    switch "foo bar" {
                    case "hello":
                        v31 := v1["hello"]
                        _ = v31
                    case "a":
                        for fi32 := 0; fi32 < 14; fi32++ {
                            v33 := GenS3{F0: 873, F1: -119, F2: -32768}
                            _ = fi32
                            _ = v33
                        }
                    case "":
                        var v34 uint32 = (v2 * (v2 << ((408) & 63)))
                        _ = v34
                    }
                case 9223372036854775807:
                    if (!(65535 <= 1)) {
                        _ = len(v1)
                        var v35 float64 = float64((float32(((0.0 / 1.0) / 1.0)) * 1.5))
                        switch 32767 {
                        case 32768:
                            v3 = (fi30 >= (v10 & v10))
                        }
                        _ = v35
                    }
                case 4294967295:
                    var v36 int64 = 9223372036854775807
                    _ = v36
                default:
                    for fi37 := 0; fi37 < 7; fi37++ {
                        fmt.Println(v3)
                        _ = fi37
                    }
                }
                _ = fi30
            }
            fmt.Println(v3, v3)
        }
        v4 += v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
