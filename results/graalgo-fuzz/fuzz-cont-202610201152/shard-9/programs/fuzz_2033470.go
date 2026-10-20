package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint; F1 uint; F2 uint8 }
func genFn0(p0 int64, p1 uint8) bool {
    var v5 int32 = 42
    p1 = (p1 << ((65535) & 63))
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 int8) int32 {
    if false {
        var v9 uint32 = 397
        var v10 int = 1
        _ = v10
        _ = v9
    }
    var v11 int16 = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            p0 = (p0 - p0)
            _ = fi13
        }
        if (4294967295 <= 243) {
            switch "hello" {
            case "foo bar":
                var v14 string = ("hello" + "hello")
                _ = v14
            case "hello":
                for fi15 := 0; fi15 < 20; fi15++ {
                    p0 = -48
                    var v16 uint8 = 128
                    _ = fi15
                    _ = v16
                }
            case "\n\t":
                if ((0.1 - 0.0) > 0.1) {
                    switch 128 {
                    case 0:
                        var v17 int64 = 995
                        _ = v17
                    case 32768:
                        var v18 int16 = (-(v11))
                        _ = v18
                    }
                    var v19 uint8 = 0
                    switch (("" + (("\n\t" + "a") + "foo bar")) + ("\n\t" + "hello")) {
                    case "\n\t":
                        var v20 int8 = (p0 + p0)
                        _ = v20
                    case "foo bar":
                        v21 := GenS1{F0: (p0 / (((p0 >> ((0) & 63))) | 1))}
                        _ = v21
                    case "a":
                        if false {
                            v19 = v19
                            switch fi12 {
                            case -9223372036854775808:
                                var v22 float64 = 0.0
                                _ = v22
                            case 2147483647:
                                v23 := GenS1{F0: (p0 ^ p0)}
                                _ = v23
                            default:
                                v24 := GenS2{F0: (v11 * v11), F1: 9223372036854775807}
                                _ = v24
                            }
                        }
                    default:
                        v11 = (v11 % ((v11) | 1))
                    }
                    _ = v19
                }
            }
            v11 += v11
        } else {
            for fi25 := 0; fi25 < 18; fi25++ {
                var v26 bool = true
                _ = fi25
                _ = v26
            }
            var v27 uint = 0
            _ = v27
        }
        _ = fi12
    }
    _ = p0
    _ = v11
    return -2147483648
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    v1 = (-(v1))
    var v28 float64 = 3.14159
    v29 := v4[-9223372036854775808]
    v30 := GenS1{F0: 100}
    if true {
        if (127 != 172) {
            for fi31 := 0; fi31 < 11; fi31++ {
                v29 += fi31
                _ = fi31
            }
            v3 = v3
        }
        v32 := []int{(v29 & v29)}
        v33 := GenS3{F0: 955, F1: uint(v3), F2: 2}
        _ = v32
        _ = v33
    } else {
        var v34 uint32 = 65535
        _ = v34
    }
    var v35 float32 = -2.25
    var v36 float32 = v35
    v36 += v36
    var v37 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
