package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) int64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float32 = ((1.5 + (0.0 - 1e10)) / -2.25)
        var v9 uint16 = 271
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2147483647
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: p0, F1: "foo bar", F2: "hello"}
            v14 := []int{32767, (fi12 + (fi11 & fi12)), 2147483647, -2147483648}
            var v15 int16 = -32768
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        switch "\n\t" {
        case "hello":
            for fi16 := 0; fi16 < 17; fi16++ {
                v10 -= v10
                _ = fi16
            }
        case "\n\t":
            if (float32(1e10) == 3.14159) {
                p0 -= (p0 / ((int16(v10)) | 1))
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 uint32 = 4294967295
                    _ = fi17
                    _ = v18
                }
                if false {
                    if (42 != 42) {
                        v19 := GenS2{F0: fi11, F1: fi11, F2: 0}
                        if true {
                            var v20 int64 = -9223372036854775808
                            var v21 uint8 = 119
                            if true {
                                v22 := GenS3{F0: v21}
                                _ = v22
                            } else {
                                v23 := []int{717, fi11, 81}
                                var v24 uint = 0
                                _ = v23
                                _ = v24
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    } else {
                        var v25 string = "foo bar"
                        v26 := []int{9223372036854775807, (fi11 | 32767), (fi11 / (((fi11 << ((371) & 63))) | 1))}
                        _ = v25
                        _ = v26
                    }
                    var v27 uint8 = 0
                    var v28 int = fi11
                    _ = v27
                    _ = v28
                } else {
                    v29 := []int{}
                    var v30 bool = true
                    if len(v29) > 0 {
                        v29[0] = -2147483648
                    }
                    _ = v29
                    _ = v30
                }
            }
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    return 535
}

func main() {
    v1 := []uint32{65535, 4294967295, 2}
    var v2 bool = false
    var v3 int64 = -1
    v4 := make(map[int]bool)
    v5 := []string{"hello", ""}
    var v6 int16 = 42
    v31 := GenS2{F0: 128, F1: 247, F2: 875}
    if v2 {
        var v32 uint8 = 127
        v6 = 42
        _ = v32
    }
    v4[-2147483648] = ((-2.25 * 0.1) > ((1.5 * -2.25) - -2.25))
    var v33 int16 = v6
    v34 := GenS3{F0: 42}
    var v35 bool = (((((-2.25 - ((0.0 / (0.1 / 1.0)) * (((float64(1.5) + (0.0 / (((-2.25 / float64(((1.5 / 0.1) * -2.25))) + 1.5) - 1e10))) * 0.1) - 1e10))) - 0.0) - 0.1) - 0.1) != (float64(-2.25) - (float64((3.14159 / float32((1e10 - (float64(((-2.25 / (1.5 + (-2.25 + 0.1))) / -2.25)) * -2.25))))) * -2.25)))
    var v36 string = "foo bar"
    _ = len(v5)
    fmt.Println(v2, v35)
    v37 := GenS1{F0: (v6 + (^(v6))), F1: v36, F2: v36}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
