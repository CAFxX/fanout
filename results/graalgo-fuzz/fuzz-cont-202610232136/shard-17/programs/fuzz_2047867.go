package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint32 }
func genFn0(p0 int64) int32 {
    var v4 uint8 = 127
    if true {
        var v5 uint = 0
        v6 := []int{-1, -2147483648, int(v5)}
        _ = v5
        _ = v6
    } else {
        if true {
            p0 = p0
            var v7 int16 = -32768
            v7 -= (-(v7))
            _ = v7
        }
        if false {
            for fi8 := 0; fi8 < 1; fi8++ {
                p0 = p0
                var v9 string = (("" + "a") + "hello")
                var v10 int8 = 42
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 float64 = ((0.0 / 1e10) + ((-2.25 / (1e10 * 1.5)) / -2.25))
            for fi12 := 0; fi12 < 10; fi12++ {
                if (v4 == v4) {
                    var v13 string = ""
                    _ = v13
                }
                _ = fi12
            }
            _ = v11
        }
        v4 = (v4 ^ v4)
    }
    _ = p0
    _ = v4
    return -950
}
func genFn1(p0 uint, p1 uint32) string {
    p1 = p1
    var v14 bool = false
    var v15 uint = (p0 - p0)
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint64 = 1
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 0
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v18
    return ("a" + "")
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    var v19 float64 = float64(((float32((0.0 / 0.1)) * 1.5) / (1e10 - 3.14159)))
    var v20 int16 = 1
    v19 = float64(0.0)
    for fi21 := 0; fi21 < 5; fi21++ {
        v3 += v3
        v22 := GenS1{F0: 19, F1: 9223372036854775807, F2: 169}
        delete(v1, ((((("foo bar" + ("foo bar" + "\n\t")) + "\n\t") + (("\n\t" + "") + ("a" + ((("hello" + "\n\t") + ((("hello" + ("hello" + "\n\t")) + "hello") + "foo bar")) + "a")))) + "foo bar") + "a"))
        _ = fi21
        _ = v22
    }
    v3 = v3
    var v23 uint8 = 127
    v19 = v19
    v24 := []int{792, -1}
    var v25 int16 = v20
    switch 128 {
    case 4294967295:
        v20 += 0
    case 42:
        var v26 uint64 = 989
        _ = v26
    }
    if false {
        if true {
            if (4294967295 <= 1) {
                if (false && true) {
                    var v27 int32 = 456
                    switch 32767 {
                    case 128:
                        _ = len(v1)
                    default:
                        var v28 float32 = 0.1
                        _ = v28
                    }
                    _ = v27
                }
            } else {
                v29 := GenS2{F0: -128, F1: 0}
                switch 128 {
                case 9223372036854775807:
                    v30 := GenS3{F0: 42}
                    _ = v30
                case 1:
                    switch "" {
                    case "\n\t":
                        fmt.Println(v3)
                    case "a":
                        for fi31 := 0; fi31 < 9; fi31++ {
                            var v32 string = "hello"
                            v33 := GenS3{F0: 1}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    default:
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    }
                default:
                    _ = len(v24)
                }
                _ = v29
            }
        } else {
            v3 -= (v3 ^ 255)
            var v34 uint16 = 2
            var v35 uint8 = 128
            _ = v34
            _ = v35
        }
        if len(v24) > 0 {
            v24[0] = -1
        }
        delete(v2, "hello")
    }
    var v36 uint16 = 1
    v3 = v3
    v36 = 477
    var v37 bool = true
    _ = len(v24)
    v38 := GenS4{F0: 2147483647, F1: (v23 ^ v23), F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
