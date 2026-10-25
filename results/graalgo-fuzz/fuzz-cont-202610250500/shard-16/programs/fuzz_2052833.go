package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0() uint64 {
    var v5 float64 = 0.1
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = ((v5 * v5) + 0.0)
        fmt.Println(v5, v5)
        for fi7 := 0; fi7 < 19; fi7++ {
            if (1e10 <= 0.1) {
                var v8 string = ("\n\t" + "foo bar")
                if true {
                    if false {
                        var v9 uint64 = 18446744073709551615
                        if (((!true) || (!true)) || (1 != -32768)) {
                            var v10 uint32 = 0
                            if true {
                                fmt.Println(fi6, v9, v9)
                                var v11 int8 = -11
                                _ = v11
                            }
                            _ = v10
                        } else {
                            var v12 uint16 = 32767
                            _ = v12
                        }
                        var v13 uint = 626
                        _ = v13
                        _ = v9
                    } else {
                        v8 = v8
                        v14 := GenS1{F0: 255, F1: 0}
                        _ = v14
                    }
                    fmt.Println(fi6, fi7, v5)
                    if (2 >= 65535) {
                        v8 = "hello"
                        v15 := GenS1{F0: 18446744073709551615, F1: 463}
                        var v16 uint16 = 2
                        if (54 > 20) {
                            v8 = (v8 + "a")
                        }
                        v5 += 1e10
                        _ = v15
                        _ = v16
                    }
                } else {
                    var v17 uint8 = 128
                    var v18 uint16 = 255
                    v19 := GenS1{F0: 228, F1: 18446744073709551615}
                    _ = v17
                    _ = v18
                    _ = v19
                }
                switch -329 {
                case 32767:
                    var v20 int32 = 105
                    _ = v20
                case 9223372036854775807:
                    v21 := GenS1{F0: 54, F1: 1}
                    _ = v21
                case -1:
                    if false {
                        v22 := GenS1{F0: 42, F1: 42}
                        _ = v22
                    }
                }
                _ = v8
            } else {
                var v23 int32 = 0
                switch ((("" + "a") + (((((("hello" + "foo bar") + "hello") + "foo bar") + ("a" + "a")) + "a") + "\n\t")) + "hello") {
                case "":
                    v24 := []int{(fi6 >> ((178) & 63)), 2147483647}
                    _ = v24
                default:
                    if (2 != 18446744073709551615) {
                        v25 := []int{fi6}
                        fmt.Println(fi6, v23)
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        _ = v25
                    }
                }
                switch fi6 {
                case 128:
                    v26 := []int{}
                    _ = v26
                case -141:
                    v23 -= v23
                    var v27 float64 = ((0.1 - 0.1) / v5)
                    v28 := GenS1{F0: 0, F1: 340}
                    _ = v27
                    _ = v28
                case 32768:
                    v5 = float64(float32(3.14159))
                default:
                    v29 := []int{127, (fi7 << ((42) & 63))}
                    _ = v29
                }
                _ = v23
            }
            _ = fi7
        }
        _ = fi6
    }
    _ = v5
    return 42
}
func genFn1(p0 int, p1 int32) int32 {
    v30 := []int{(^(p0)), -2147483648, p0, (p0 - p0)}
    switch (p0 * (p0 & (p0 % ((-9223372036854775808) | 1)))) {
    case -964:
        var v31 uint16 = 255
        _ = v31
    default:
        fmt.Println(p1, p0)
    }
    v30 = append(v30, p0)
    _ = p0
    _ = p1
    _ = v30
    return (p1 * p1)
}

func main() {
    var v1 uint64 = 610
    var v2 int16 = -1
    v3 := []uint32{420}
    var v4 float64 = 0.1
    var v32 uint = 2
    var v33 uint8 = 128
    switch -1 {
    case 0:
        if false {
            v4 = (v4 * ((1.5 + ((v4 + 1e10) / ((-2.25 * (v4 - v4)) - v4))) + v4))
        } else {
            var v34 uint64 = (v1 >> ((255) & 63))
            _ = len(v3)
            _ = v34
        }
    case 32767:
        var v35 bool = (!true)
        _ = v35
    case 128:
        _ = len(v3)
    }
    v1 -= genFn0()
    v32 = 65535
    var v36 float32 = 1.5
    if ((v33 & v33) == v33) {
        for fi37 := 0; fi37 < 3; fi37++ {
            if false {
                if true {
                    v38 := GenS1{F0: (v32 ^ 255), F1: v1}
                    _ = v38
                }
            }
            _ = fi37
        }
    }
    var v39 float64 = float64(float32(v4))
    v40 := GenS1{F0: 407, F1: v1}
    var v41 int32 = genFn1(0, 2147483647)
    fmt.Println(v41)
    _ = len(v3)
    v42 := GenS1{F0: 0, F1: v1}
    switch 1 {
    case 9223372036854775807:
        var v43 string = "foo bar"
        _ = v43
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= (v2 << (((v32 & v32)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
