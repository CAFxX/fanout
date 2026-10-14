package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 string; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64; F1 int8 }
func genFn0(p0 uint32) uint16 {
    var v5 bool = (-896 <= -1)
    v6 := []int{1}
    for fi7 := 0; fi7 < 4; fi7++ {
        if (v5 || v5) {
            var v8 int = fi7
            _ = v8
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v6 = append(v6, (fi7 / ((-2147483648) | 1)))
            v9 := GenS1{F0: ""}
            _ = v9
        }
        switch fi7 {
        case 32768:
            v10 := GenS2{F0: 42, F1: fi7}
            _ = v10
        default:
            v6 = append(v6, (fi7 >> ((43) & 63)))
        }
        _ = fi7
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        p0 = (p0 % ((p0) | 1))
    case "hello":
        var v11 uint8 = 127
        _ = v11
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            v5 = v5
            _ = fi12
        }
    }
    v13 := GenS3{F0: ("foo bar" + "\n\t"), F1: 555, F2: 255}
    _ = p0
    _ = v13
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    var v3 uint16 = 32767
    var v4 uint8 = 255
    v14 := []int{32767, -1, 127}
    fmt.Println(v4)
    for fi15 := 0; fi15 < 11; fi15++ {
        v3 += (v3 / (((v3 << ((65535) & 63))) | 1))
        switch v2 {
        case "\n\t":
            if false {
                v4 += 127
                for fi16 := 0; fi16 < 19; fi16++ {
                    _ = len(v14)
                    var v17 string = (("\n\t" + (v2 + "")) + ((v2 + (v2 + (v2 + v2))) + v2))
                    _ = fi16
                    _ = v17
                }
            }
        case "foo bar":
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 float64 = float64((0.0 / 1.0))
                v20 := GenS4{F0: false, F1: 18446744073709551615, F2: 42}
                var v21 float64 = 3.14159
                v21 = 3.14159
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        _ = fi15
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 uint32 = 42
        v3 -= (v3 * (v3 | 65535))
        v3 += (v3 + v3)
        _ = fi22
        _ = v23
    }
    fmt.Println(v2, v3, v3)
    if false {
        fmt.Println(v2)
        var v24 uint64 = 0
        _ = v24
    }
    v25 := GenS2{F0: -1, F1: -1}
    if true {
        if true {
            if false {
                var v26 float32 = ((float32(1.5) - float32(((1.5 - -2.25) + -2.25))) - 3.14159)
                for fi27 := 0; fi27 < 4; fi27++ {
                    var v28 int32 = -1
                    _ = fi27
                    _ = v28
                }
                _ = v26
            } else {
                var v29 string = v2
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 float64 = (1e10 + 1.5)
                    switch (fi30 + fi30) {
                    case 9223372036854775807:
                        v3 = (v3 + (v3 >> ((255) & 63)))
                    default:
                        var v32 int64 = -573
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 7; fi33++ {
                        fmt.Println(v4, v29)
                        v34 := GenS5{F0: 42, F1: 564}
                        fmt.Println(v3, v3, fi33)
                        if (false || (-452 == 42)) {
                            var v35 uint8 = (v4 * v4)
                            _ = v35
                        } else {
                            var v36 float64 = v31
                            v37 := GenS6{F0: 571, F1: 1}
                            var v38 float32 = (-2.25 + 3.14159)
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        _ = fi33
                        _ = v34
                    }
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
        }
        v4 += v4
    } else {
        switch ("hello" + v2) {
        case "":
            v14 = append(v14, -9223372036854775808)
        }
    }
    var v39 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
}
