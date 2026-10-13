package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0() uint32 {
    switch 32768 {
    case 571:
        var v6 float64 = 1.5
        _ = v6
    case 474:
        var v7 int32 = -552
        _ = v7
    case 4294967295:
        var v8 uint8 = 194
        _ = v8
    default:
        if (1 != 920) {
            var v9 int = -48
            _ = v9
        } else {
            if ((3.14159 >= (0.0 + ((1e10 - 0.1) / float64(0.1)))) && true) {
                switch "a" {
                case "hello":
                    if (42 == 42) {
                        v10 := GenS1{F0: 42, F1: 0, F2: 194}
                        switch (("\n\t" + ("hello" + "foo bar")) + "\n\t") {
                        case "":
                            if false {
                                for fi11 := 0; fi11 < 19; fi11++ {
                                    var v12 float64 = (1e10 - 1.5)
                                    if (0 == 132) {
                                        v13 := []int{fi11, fi11}
                                        _ = v13
                                    } else {
                                        v14 := []int{(fi11 % ((fi11) | 1)), (fi11 ^ fi11), (fi11 + -2147483648)}
                                        var v15 uint16 = 1
                                        var v16 float64 = v12
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = fi11
                                    _ = v12
                                }
                            } else {
                                var v17 string = "foo bar"
                                for fi18 := 0; fi18 < 18; fi18++ {
                                    if (((true || ("a" >= v17)) || (-2.25 < 1.5)) || true) {
                                        v19 := []int{0, (fi18 / ((-725) | 1)), 42}
                                        _ = v19
                                    }
                                    v20 := []int{(fi18 | fi18), 32767, fi18}
                                    for fi21 := 0; fi21 < 19; fi21++ {
                                        v17 = v17
                                        _ = fi21
                                    }
                                    _ = fi18
                                    _ = v20
                                }
                                v17 = ""
                                _ = v17
                            }
                        case "a":
                            var v22 int64 = 1
                            _ = v22
                        case "hello":
                            v23 := []int{42}
                            _ = v23
                        }
                        _ = v10
                    }
                case "\n\t":
                    v24 := []int{0, 128, 32768}
                    _ = v24
                case "foo bar":
                    v25 := []int{128}
                    _ = v25
                }
                v26 := []int{9223372036854775807, 127, 42, -1}
                var v27 int = -85
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint = 1
                _ = fi28
                _ = v29
            }
        }
    }
    if (!(true && true)) {
        var v30 int64 = 384
        var v31 int64 = (v30 ^ v30)
        _ = v30
        _ = v31
    }
    var v32 uint = 539
    _ = v32
    return 2147483647
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 804
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]uint32)
    v33 := v3[(v2 % ((v2) | 1))]
    v33 -= v33
    var v34 int16 = 285
    fmt.Println(v2, v1)
    switch "" {
    case "":
        v35 := GenS1{F0: int16(v2), F1: 2, F2: 4294967295}
        _ = v35
    case "\n\t":
        var v36 int32 = 792
        _ = v36
    default:
        var v37 int32 = -2147483648
        _ = v37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := GenS2{F0: -649, F1: 18, F2: v1}
        v40 := v4["\n\t"]
        v2 += (v2 % (((^(v2))) | 1))
        _ = fi38
        _ = v39
        _ = v40
    }
    v33 += 422
    fmt.Println(v2)
    for fi41 := 0; fi41 < 8; fi41++ {
        var v42 uint = 748
        var v43 uint8 = 29
        fmt.Println(v2, v34, v34)
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 int32 = 0
    v45 := GenS3{F0: 507, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
