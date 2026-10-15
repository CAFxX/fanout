package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 int64, p1 int16) uint16 {
    switch (("hello" + "foo bar") + ("a" + ("a" + ""))) {
    case "foo bar":
        var v6 bool = true
        _ = v6
    default:
        if (1 <= (p1 + (p1 | p1))) {
            for fi7 := 0; fi7 < 1; fi7++ {
                if true {
                    fmt.Println(p1)
                }
                var v8 uint16 = 32767
                v9 := []int{}
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            v10 := []int{}
            fmt.Println(p1)
            _ = v10
        }
    }
    var v11 uint16 = 65535
    _ = p0
    _ = p1
    _ = v11
    return 739
}
func genFn1() int8 {
    switch -304 {
    case 1:
        v12 := []int{128}
        _ = v12
    case 4294967295:
        var v13 uint64 = 373
        _ = v13
    case -1:
        if (0.1 >= 3.14159) {
            if true {
                var v14 float32 = (float32(0.0) - 0.0)
                v15 := GenS1{F0: 32767, F1: 2, F2: 871}
                _ = v14
                _ = v15
            }
            if (0 < 198) {
                if ((1e10 / 3.14159) <= float64(0.0)) {
                    var v16 int8 = 1
                    _ = v16
                }
            } else {
                var v17 uint8 = 128
                if (!(!(2 > 2))) {
                    v18 := []int{0, 9223372036854775807, 32768, 32767}
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 int32 = 0
                        v18 = append(v18, (fi19 ^ (-(fi19))))
                        fmt.Println(fi19, v17)
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                } else {
                    if (-375 >= 2147483647) {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            v22 := []int{}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v17 += (v17 | 40)
                                if ((v17 & v17) >= (v17 ^ (v17 % ((255) | 1)))) {
                                    v24 := GenS1{F0: 2, F1: 4294967295, F2: -6}
                                    _ = len(v22)
                                    _ = v24
                                } else {
                                    if false {
                                        var v25 float32 = 1e10
                                        _ = v25
                                    }
                                    v26 := GenS2{F0: 556, F1: 1000}
                                    _ = v26
                                }
                                var v27 int16 = 42
                                fmt.Println(v27, fi21, v17)
                                v22 = append(v22, -9223372036854775808)
                                _ = fi23
                                _ = v27
                            }
                            _ = fi21
                            _ = v22
                        }
                        var v28 float32 = 0.0
                        var v29 uint16 = 586
                        _ = v28
                        _ = v29
                    }
                }
                _ = v17
            }
            v30 := GenS2{F0: -357, F1: 2147483648}
            _ = v30
        }
    }
    var v31 bool = true
    v32 := []int{-1, 32768, -1, 785}
    var v33 int16 = -1
    _ = v31
    _ = v32
    _ = v33
    return 96
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[string]int)
    var v3 int = -838
    v4 := []string{"a"}
    v5 := make(map[int]bool)
    var v34 uint = 2
    var v35 int = (v3 | (v3 >> ((1) & 63)))
    var v36 bool = (42 > 4294967295)
    switch (("\n\t" + "hello") + "foo bar") {
    case "":
        v3 += v3
    case "\n\t":
        fmt.Println(v35)
    case "hello":
        v37 := v5[v3]
        _ = v37
    }
    if (2147483647 > 0) {
        var v38 int32 = -1
        v35 += -996
        _ = v38
    }
    var v39 int8 = 100
    v40 := GenS1{F0: 255, F1: 18446744073709551615, F2: 0}
    v4 = append(v4, "")
    var v41 bool = (("hello" + (("\n\t" + ("foo bar" + "a")) + "\n\t")) > ("hello" + "hello"))
    var v42 uint = (v34 ^ 0)
    v36 = ((float32(float64(3.14159)) - 3.14159) != float32(((((0.1 / (1.5 + -2.25)) + 0.0) * 0.0) * (float64(3.14159) / 1e10))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
