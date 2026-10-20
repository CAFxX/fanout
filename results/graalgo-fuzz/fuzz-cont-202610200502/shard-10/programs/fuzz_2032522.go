package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 int64, p1 int8, p2 int16) string {
    v6 := []int{1}
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(p2, p0, fi7)
        _ = fi7
    }
    fmt.Println(p2, p2)
    p1 -= 42
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return (v8 + v8)
}
func genFn1(p0 uint64, p1 uint16, p2 int) int32 {
    p0 = 1
    var v9 uint64 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 741
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 0
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    v5 := make(map[int]string)
    v10 := GenS1{F0: 255, F1: 780, F2: 735}
    var v11 uint = 432
    if true {
        v11 += (v11 ^ (v11 % ((4294967295) | 1)))
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                v2 -= v2
                if ((0.1 / 1.0) != 0.0) {
                    v4 = -1
                    switch fi12 {
                    case -9223372036854775808:
                        if (2147483647 == 1) {
                            v11 = (v11 ^ (v11 - (v11 % ((v11) | 1))))
                            var v14 uint16 = 255
                            _ = v14
                        } else {
                            var v15 float32 = 3.14159
                            var v16 int32 = 0
                            _ = v15
                            _ = v16
                        }
                    }
                } else {
                    if false {
                        if (((2 == 0) || (2 > 945)) && true) {
                            if (!false) {
                                var v17 int = (fi13 >> ((v11) & 63))
                                v17 += fi13
                                var v18 string = "a"
                                if (42 > 42) {
                                    var v19 int32 = 0
                                    _ = v19
                                } else {
                                    var v20 uint = (^(v11))
                                    v21 := GenS1{F0: v11, F1: v20, F2: 4294967295}
                                    _ = v20
                                    _ = v21
                                }
                                switch v18 {
                                case "foo bar":
                                    v22 := []int{}
                                    v2 -= v2
                                    switch (fi13 ^ 127) {
                                    case 127:
                                        v23 := GenS2{F0: v18, F1: (v4 & 1)}
                                        _ = v23
                                    default:
                                        if (!(true || (2 == 42))) {
                                            v22 = append(v22, int(v11))
                                            v2 = (v2 & v2)
                                            v24 := GenS1{F0: (v11 & v11), F1: (v11 ^ 2), F2: (v11 * (v11 + (v11 % (((v11 << ((v11) & 63))) | 1))))}
                                            _ = v24
                                        } else {
                                            v22 = append(v22, 985)
                                        }
                                    }
                                    _ = v22
                                case "\n\t":
                                    var v25 uint32 = 65535
                                    _ = v25
                                }
                                _ = v17
                                _ = v18
                            } else {
                                fmt.Println(v4, fi12)
                            }
                            v26 := v5[(fi12 ^ fi12)]
                            _ = v26
                        } else {
                            v11 = 42
                            var v27 uint32 = 2147483648
                            _ = v27
                        }
                        switch -1 {
                        case 127:
                            if false {
                                v2 = (v2 ^ v2)
                                fmt.Println(v2, v2, v11)
                            } else {
                                _ = len(v5)
                                var v28 uint32 = 2147483647
                                v2 -= -1
                                _ = v28
                            }
                        case -2147483648:
                            _ = len(v5)
                        default:
                            if (18446744073709551615 != 4294967295) {
                                delete(v5, 4294967295)
                            }
                        }
                    }
                    v11 += v11
                }
                v29 := GenS2{F0: "\n\t", F1: 42}
                v2 = -32768
                _ = fi13
                _ = v29
            }
            switch 0 {
            case 128:
                v30 := GenS2{F0: "", F1: v4}
                _ = v30
            case -1:
                var v31 bool = (1 <= 1)
                _ = v31
            case 0:
                var v32 uint16 = 1
                _ = v32
            default:
                v4 = v4
            }
            _ = fi12
        }
    }
    var v33 int = 610
    v34 := []int{v33, v33, v33, (v33 ^ (v33 - -2147483648))}
    v35 := GenS1{F0: 733, F1: (^(v11)), F2: 2}
    v36 := v5[v33]
    v36 = v36
    fmt.Println(v4)
    v11 = v11
    v11 = (v11 | 65535)
    var v37 uint8 = 42
    fmt.Println(v33, v11, v36)
    v37 = (v37 | (v37 - (v37 * 171)))
    var v38 int = v33
    v11 = v11
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
