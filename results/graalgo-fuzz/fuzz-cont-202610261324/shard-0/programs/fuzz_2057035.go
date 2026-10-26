package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 42
    var v3 float64 = 3.14159
    var v4 float32 = ((3.14159 + (float32(float64(0.1)) - (1e10 / (0.0 + 0.1)))) - ((((float32(v3) - (-2.25 + 0.1)) * 1e10) * float32(float64(1e10))) - 0.1))
    var v5 int64 = -839
    if ((!false) && true) {
        var v6 int = -9223372036854775808
        _ = v6
    } else {
        v3 = v3
        if (v3 != (float64(0.1) / (v3 + (float64(v4) - v3)))) {
            var v7 uint32 = 2
            _ = v7
        } else {
            var v8 uint64 = 1
            for fi9 := 0; fi9 < 7; fi9++ {
                if false {
                    v4 = v4
                }
                v5 += -9223372036854775808
                var v10 float64 = (v3 * (0.0 + v3))
                _ = fi9
                _ = v10
            }
            if false {
                v11 := GenS1{F0: "hello", F1: 1}
                _ = v11
            }
            _ = v8
        }
    }
    if false {
        v12 := GenS1{F0: ((("a" + "") + ("a" + "foo bar")) + "hello"), F1: 1}
        _ = v12
    } else {
        switch -622 {
        case 0:
            v13 := []int{4294967295, -808}
            _ = v13
        case -2147483648:
            v4 -= v4
        }
        var v14 int64 = (v5 | v5)
        _ = v14
    }
    var v15 float64 = v3
    v2 = 9223372036854775807
    v16 := GenS1{F0: "a", F1: 4294967295}
    delete(v1, 2)
    var v17 float32 = v4
    fmt.Println(v4, v17, v2)
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := []int{(fi19 / ((fi19) | 1)), 2147483647, (fi18 + 127)}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    var v21 int64 = (v5 / ((-744) | 1))
    var v22 uint16 = 168
    if (207 != 2) {
        for fi23 := 0; fi23 < 9; fi23++ {
            if true {
                var v24 uint = 245
                _ = v24
            } else {
                v2 += 0
                fmt.Println(v21, v3)
            }
            var v25 uint64 = 4294967295
            v25 = (^(v25))
            _ = fi23
            _ = v25
        }
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v5)
        case "hello":
            var v26 bool = (-1 == -101)
            _ = v26
        case "":
            var v27 float64 = float64(0.1)
            _ = v27
        }
        v28 := []int{655, 127}
        _ = v28
    } else {
        var v29 bool = ((254 > 255) && true)
        _ = v29
    }
    if true {
        if (18 >= 2) {
            if true {
                if true {
                    v30 := GenS1{F0: ("" + "a"), F1: 4294967295}
                    if true {
                        var v31 int8 = -1
                        _ = v31
                    } else {
                        v4 = v17
                        var v32 int16 = 299
                        _ = v32
                    }
                    fmt.Println(v22, v2, v17)
                    _ = v30
                } else {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v15 += v3
                            var v35 uint = 0
                            var v36 string = "foo bar"
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                        var v37 float64 = (-2.25 / (v3 / v3))
                        _ = fi33
                        _ = v37
                    }
                    var v38 uint8 = 255
                    v39 := GenS2{F0: (^(v38)), F1: (v22 / (((v22 ^ v22)) | 1)), F2: 1}
                    _ = v38
                    _ = v39
                }
                v1[167] = "a"
            }
            var v40 float64 = v15
            if false {
                switch ((("" + "foo bar") + "foo bar") + ("" + "\n\t")) {
                case "hello":
                    v21 += (-(v2))
                case "":
                    fmt.Println(v40, v22)
                default:
                    v3 += v15
                }
                fmt.Println(v17)
                v41 := GenS3{F0: ("foo bar" + "a"), F1: 0, F2: v22}
                _ = v41
            }
            _ = v40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
