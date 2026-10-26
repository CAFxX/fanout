package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 uint16 }
func genFn0(p0 bool, p1 int) int64 {
    v6 := GenS1{F0: 4294967295, F1: ""}
    if (true || (686 > 4294967295)) {
        if ((!false) || (0 != 796)) {
            var v7 uint64 = 0
            p1 += (p1 | p1)
            _ = v7
        } else {
            fmt.Println(p0, p1)
            for fi8 := 0; fi8 < 11; fi8++ {
                if (!p0) {
                    if (1.5 == 1e10) {
                        var v9 string = "foo bar"
                        _ = v9
                    }
                    v10 := []int{}
                    if true {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            _ = len(v10)
                            _ = fi11
                        }
                        for fi12 := 0; fi12 < 2; fi12++ {
                            p0 = p0
                            p1 = p1
                            switch fi8 {
                            case 42:
                                var v13 uint8 = 67
                                _ = v13
                            case -396:
                                var v14 uint8 = 128
                                _ = v14
                            case 9223372036854775807:
                                p1 += 0
                            default:
                                var v15 int32 = 1
                                _ = v15
                            }
                            _ = fi12
                        }
                    } else {
                        var v16 int16 = 1
                        _ = v16
                    }
                    _ = v10
                }
                _ = fi8
            }
        }
        if (-1 > 259) {
            var v17 int32 = 2147483647
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 float64 = float64(float32(1.5))
                _ = fi18
                _ = v19
            }
            _ = v17
        } else {
            p0 = true
            var v20 bool = (110 == -1)
            _ = v20
        }
    } else {
        p0 = p0
        var v21 uint32 = 597
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := GenS1{F0: 4294967295, F1: "a"}
            var v24 uint = 0
            v25 := GenS2{F0: 79, F1: 608, F2: 127}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v21
    }
    var v26 int16 = -32768
    var v27 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v6
    return (v27 / (((v27 << ((4294967295) & 63))) | 1))
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = -1
    v3 := make(map[string]uint32)
    v4 := make(map[string]string)
    var v5 uint32 = 2147483648
    v5 = v5
    switch 128 {
    case 32767:
        if true {
            for fi28 := 0; fi28 < 20; fi28++ {
                if ((100 != -128) && true) {
                    v29 := []int{(fi28 * (fi28 % (((^(fi28))) | 1)))}
                    _ = v29
                }
                for fi30 := 0; fi30 < 17; fi30++ {
                    var v31 int64 = 9223372036854775807
                    v1 = 191
                    _ = fi30
                    _ = v31
                }
                _ = fi28
            }
        }
    }
    if (1 > 152) {
        v1 -= (v1 - (v1 ^ 399))
        var v32 int = -2147483648
        _ = v32
    }
    v33 := []int{}
    fmt.Println(v1, v5, v1)
    var v34 bool = true
    var v35 uint = 65535
    v36 := GenS3{F0: v34}
    fmt.Println(v5)
    var v37 int8 = -106
    if ((v2 - 1) <= (-(v2))) {
        if (!v34) {
            v38 := GenS4{F0: genFn0(((v37 >= v37) && v34), 2147483647), F1: "\n\t"}
            for fi39 := 0; fi39 < 13; fi39++ {
                delete(v3, "hello")
                _ = fi39
            }
            v40 := GenS5{F0: uint16(v5)}
            _ = v38
            _ = v40
        } else {
            var v41 bool = (!false)
            var v42 uint32 = uint32(v35)
            v43 := GenS4{F0: -1, F1: (("a" + "foo bar") + "")}
            _ = v41
            _ = v42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
