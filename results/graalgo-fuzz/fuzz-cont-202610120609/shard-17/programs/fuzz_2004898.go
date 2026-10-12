package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8) int16 {
    var v7 float64 = (0.0 - 1.5)
    v8 := GenS1{F0: 255, F1: -884}
    for fi9 := 0; fi9 < 17; fi9++ {
        v7 -= 1e10
        v7 -= 1.5
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return -719
}
func genFn1() bool {
    var v10 string = (((("" + ("" + ("\n\t" + "foo bar"))) + "") + "hello") + "")
    v11 := []int{}
    fmt.Println(v10)
    fmt.Println(v10)
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 uint8 = 96
    v2 := make(map[string]string)
    v3 := []int{42, 0, -689}
    v4 := make(map[string]int)
    var v5 uint64 = 2
    v6 := make(map[int]uint32)
    v12 := GenS2{F0: v1}
    var v13 int32 = 0
    v5 += (v5 | 781)
    switch 42 {
    case 1:
        v14 := GenS1{F0: v1, F1: 775}
        _ = v14
    case 32767:
        if (!true) {
            switch -657 {
            case 127:
                for fi15 := 0; fi15 < 8; fi15++ {
                    v1 += v1
                    v16 := GenS1{F0: v1, F1: -9223372036854775808}
                    if false {
                        _ = len(v6)
                        v17 := GenS2{F0: v1}
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
            case -1:
                var v18 float32 = 1e10
                _ = v18
            case 1:
                switch "hello" {
                case "hello":
                    var v19 int = -2147483648
                    _ = v19
                case "":
                    _ = len(v3)
                case "a":
                    var v20 bool = (0 >= 65535)
                    _ = v20
                default:
                    var v21 int16 = -190
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int = -285
                v1 = v1
                _ = fi22
                _ = v23
            }
            var v24 string = "hello"
            _ = v24
        }
    case -406:
        v13 = 767
    }
    var v25 uint16 = 255
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 int64 = -878
        switch "hello" {
        case "":
            if false {
                var v28 float64 = float64((1.5 - (-2.25 / (-2.25 * float32((-2.25 - 1.5))))))
                var v29 int = -2147483648
                _ = v28
                _ = v29
            }
        case "a":
            v30 := GenS1{F0: (v1 ^ v1), F1: (v27 >> ((65535) & 63))}
            _ = v30
        case "foo bar":
            switch "foo bar" {
            case "hello":
                switch -575 {
                case 32767:
                    v13 -= v13
                }
            case "":
                for fi31 := 0; fi31 < 6; fi31++ {
                    switch "hello" {
                    case "":
                        if ((709 == 18446744073709551615) || true) {
                            var v32 int32 = v13
                            var v33 float64 = 0.0
                            v25 -= uint16(v1)
                            _ = v32
                            _ = v33
                        }
                    }
                    var v34 uint32 = 65535
                    _ = fi31
                    _ = v34
                }
            case "a":
                switch ("hello" + "") {
                case "foo bar":
                    v27 += (-(v27))
                case "":
                    _ = len(v6)
                case "\n\t":
                    v1 = 2
                default:
                    if (-24 < 42) {
                        var v35 float64 = float64(1.5)
                        v13 -= (v13 + (v13 * -2147483648))
                        if len(v3) > 0 {
                            v3[0] = fi26
                        }
                        _ = v35
                    }
                }
            default:
                v25 -= v25
            }
        }
        _ = fi26
        _ = v27
    }
    delete(v4, ("a" + "hello"))
    var v36 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
