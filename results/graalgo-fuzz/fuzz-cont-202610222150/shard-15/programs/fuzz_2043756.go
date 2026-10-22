package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 uint64) uint8 {
    if (true && false) {
        v5 := []int{1, 4294967295, 328, 4294967295}
        v6 := GenS1{F0: 0, F1: 174}
        for fi7 := 0; fi7 < 2; fi7++ {
            if true {
                if (-1 < -60) {
                    fmt.Println(p0)
                }
            }
            p0 -= (p0 & -2147483648)
            _ = fi7
        }
        _ = v5
        _ = v6
    } else {
        v8 := []int{}
        if len(v8) > 0 {
            v8[0] = 2147483647
        }
        _ = v8
    }
    fmt.Println(p0, p1)
    var v9 bool = (168 > 1)
    var v10 int8 = 100
    v9 = (944 > 643)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int8 = 100
    var v2 string = "foo bar"
    var v3 string = "hello"
    var v4 float64 = 1e10
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint64 = 4294967295
            if true {
                v14 := GenS2{F0: -2147483648}
                if false {
                    switch fi11 {
                    case 2147483647:
                        var v15 bool = true
                        _ = v15
                    case -1:
                        v4 = (1.5 - 1.5)
                    case -639:
                        switch ("" + "\n\t") {
                        case "hello":
                            v16 := []int{-9223372036854775808, (fi11 / ((fi12) | 1))}
                            _ = v16
                        case "":
                            if true {
                                v17 := []int{(fi12 - fi11), fi11, 32768, fi12}
                                switch (v3 + v2) {
                                case "foo bar":
                                    var v18 uint = 0
                                    _ = v18
                                default:
                                    if len(v17) > 0 {
                                        v17[0] = fi12
                                    }
                                    v4 += ((((((v4 - float64(0.0)) - v4) + (3.14159 + v4)) + 1e10) / 1.5) * -2.25)
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v1 = v1
                                    case "foo bar":
                                        var v19 uint8 = 127
                                        _ = v19
                                    default:
                                        var v20 float32 = (-2.25 + 0.1)
                                        _ = v20
                                    }
                                }
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                if ((42 > -1) || (-1 <= 2147483647)) {
                                    v13 += 18446744073709551615
                                }
                                v13 = (^(v13))
                                _ = v17
                            } else {
                                fmt.Println(v4, v13)
                                var v21 uint = 1
                                var v22 int32 = -2147483648
                                fmt.Println(v2, fi12)
                                if true {
                                    var v23 float32 = 0.1
                                    v24 := []int{9223372036854775807, 128, fi12, (fi12 - fi11)}
                                    v4 = ((3.14159 / v4) + ((1e10 + -2.25) / 1.0))
                                    _ = v23
                                    _ = v24
                                }
                                _ = v21
                                _ = v22
                            }
                        default:
                            var v25 uint = 2
                            _ = v25
                        }
                    }
                } else {
                    v2 = ""
                    v13 += (v13 / ((4294967295) | 1))
                }
                _ = v14
            }
            var v26 uint64 = v13
            _ = fi12
            _ = v13
            _ = v26
        }
        var v27 uint32 = 2
        _ = fi11
        _ = v27
    }
    v28 := GenS1{F0: 32767, F1: 7}
    v4 -= v4
    var v29 float64 = v4
    var v30 float64 = (float64(1e10) / 1.5)
    var v31 int64 = 345
    if (0 == 342) {
        var v32 string = "a"
        v33 := []int{4294967295, 42, 4294967295, 127}
        var v34 uint16 = 639
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 int = -798
    var v36 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
