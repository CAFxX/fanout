package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0() uint {
    v6 := GenS1{F0: 42, F1: 2147483648, F2: ((-1 <= 9223372036854775807) || true)}
    if (1e10 != 0.0) {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{fi7, fi7}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS2{F0: 42}
    var v10 uint32 = 2147483648
    _ = v10
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 uint32) int16 {
    if (1 <= 127) {
        switch "" {
        case "":
            if (((((1e10 == (0.1 * 1.5)) && false) || (127 <= 42)) && (!true)) && false) {
                if true {
                    v11 := []int{}
                    _ = v11
                } else {
                    if (!((p0 < p0) && (0 != -1))) {
                        p0 = genFn0()
                    }
                    fmt.Println(p0, p0)
                    if (("\n\t" + "a") <= "hello") {
                        if ((int64(p1) >= 9223372036854775807) || ((-742 < -9223372036854775808) || (!true))) {
                            var v12 uint16 = 255
                            _ = v12
                        }
                        p1 = 4294967295
                        p1 = p1
                    }
                }
                p0 = p0
                var v13 int = -235
                _ = v13
            } else {
                if false {
                    var v14 bool = (!false)
                    var v15 uint16 = 65535
                    _ = v14
                    _ = v15
                }
                var v16 uint8 = 127
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 2; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    if ((2 <= 18446744073709551615) && true) {
                        p0 = genFn0()
                        v19 := []int{fi17, fi18, fi17}
                        fmt.Println(p1, fi18, fi18)
                        if true {
                            fmt.Println(p1)
                            var v20 uint64 = 2
                            p1 = p1
                            _ = v20
                        } else {
                            var v21 int64 = -161
                            p1 = (p1 + 243)
                            v22 := GenS3{F0: 245, F1: 2, F2: 2147483647}
                            _ = v21
                            _ = v22
                        }
                        p1 = p1
                        _ = v19
                    }
                    _ = fi18
                }
                var v23 int64 = 1
                _ = fi17
                _ = v23
            }
        }
        v24 := []int{127}
        v24 = append(v24, 0)
        _ = v24
    }
    switch 0 {
    case -9223372036854775808:
        p1 += p1
    case 9223372036854775807:
        p0 += 431
    case -78:
        p0 -= p0
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    var v1 int8 = 127
    var v2 string = "foo bar"
    v3 := []int{}
    var v4 int8 = 42
    v5 := make(map[string]int)
    fmt.Println(v1, v2)
    for fi25 := 0; fi25 < 20; fi25++ {
        switch (v2 + v2) {
        case "foo bar":
            v26 := GenS4{F0: 127}
            _ = v26
        }
        _ = fi25
    }
    v2 = "a"
    var v27 uint = 112
    for fi28 := 0; fi28 < 8; fi28++ {
        v3 = append(v3, fi28)
        _ = fi28
    }
    _ = len(v3)
    v1 += (-(v1))
    if (!(true && (false && (1 < 127)))) {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int8 = -1
            switch (fi29 | fi29) {
            case 9223372036854775807:
                for fi31 := 0; fi31 < 20; fi31++ {
                    var v32 float64 = (1.5 + (0.1 + -2.25))
                    fmt.Println(v1, v27)
                    var v33 bool = true
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case 127:
                v30 += (^(v30))
            case -854:
                _ = len(v3)
            }
            _ = fi29
            _ = v30
        }
        v34 := v5["foo bar"]
        if (18446744073709551615 > 4294967295) {
            v4 = v1
            var v35 int32 = 230
            _ = v35
        }
        _ = v34
    } else {
        v36 := GenS1{F0: 9223372036854775807, F1: 948, F2: (32767 > 42)}
        var v37 uint16 = 65535
        _ = v36
        _ = v37
    }
    if (0.0 == 0.1) {
        v38 := GenS2{F0: -1}
        var v39 float32 = ((1.5 + 0.0) / float32(0.1))
        _ = v38
        _ = v39
    }
    v40 := GenS4{F0: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
