package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint16, p1 int32, p2 int16) uint64 {
    var v5 uint16 = p0
    v6 := GenS1{F0: -45}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint) int16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(p1, p0, p1)
        var v8 uint64 = 42
        var v9 bool = p0
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{127, 1, 127}
    var v11 uint16 = 242
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]bool)
    v3 := []string{}
    var v4 int = 0
    switch "\n\t" {
    case "":
        for fi12 := 0; fi12 < 14; fi12++ {
            if true {
                var v13 float32 = 0.1
                var v14 uint16 = v1
                v1 -= (^(v1))
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: 1, F1: 42, F2: fi12}
            fmt.Println(fi12, v1)
            _ = fi12
            _ = v15
        }
    case "hello":
        v16 := GenS3{F0: 1}
        _ = v16
    }
    if false {
        if true {
            v4 -= (v4 % ((-9223372036854775808) | 1))
            if true {
                if true {
                    fmt.Println(v4, v4)
                } else {
                    switch (("" + ("hello" + "foo bar")) + "a") {
                    case "foo bar":
                        var v17 int8 = -128
                        _ = v17
                    case "hello":
                        fmt.Println(v4)
                    default:
                        v18 := GenS4{F0: 305, F1: "foo bar"}
                        _ = v18
                    }
                    var v19 int32 = -1
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        if (true && true) {
                            var v20 uint32 = 2147483648
                            v4 = (v4 ^ (v4 / (((v4 + (^(v4)))) | 1)))
                            _ = v20
                        } else {
                            fmt.Println(v19, v4)
                        }
                    case "a":
                        v4 = (-(v4))
                    }
                    _ = v19
                }
            } else {
                fmt.Println(v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v21 float64 = 3.14159
                _ = v21
            }
        }
        v22 := GenS1{F0: 100}
        _ = v22
    }
    var v23 uint16 = v1
    v24 := GenS5{F0: 0}
    var v25 int16 = genFn1(false, 223)
    if true {
        if true {
            if len(v3) > 0 {
                v3[0] = ""
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float64 = float64(float32(((3.14159 / (0.0 + ((-2.25 * 1e10) + 3.14159))) * 1e10)))
                if false {
                    if len(v3) > 0 {
                        v3[0] = ("" + ("hello" + ("hello" + "a")))
                    }
                    v28 := GenS6{F0: (v25 | v25), F1: v1}
                    var v29 bool = ((235 != 67) && (true || true))
                    _ = v28
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
            var v30 int = 1
            _ = v30
        }
    } else {
        var v31 float64 = 1e10
        v31 -= ((v31 / 1.5) / 3.14159)
        _ = v31
    }
    var v32 bool = true
    var v33 float64 = -2.25
    fmt.Println(v23)
    v34 := GenS3{F0: 747}
    var v35 float32 = 0.0
    var v36 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
