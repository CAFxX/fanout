package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 uint16) uint64 {
    var v7 int64 = 0
    p0 = p0
    v8 := []int{}
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = -565
    var v3 int16 = 1000
    var v4 int16 = (v2 * v3)
    v5 := []string{(((v1 + ("a" + v1)) + "foo bar") + "\n\t"), v1, "hello", v1}
    var v6 int32 = 0
    fmt.Println(v1)
    if (!false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v5 = append(v5, "hello")
        if (255 != 155) {
            v9 := GenS1{F0: -1, F1: genFn0(255)}
            _ = v9
        } else {
            var v10 int = 0
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS2{F0: 0}
                var v13 float64 = (0.0 - (1e10 * 3.14159))
                v4 = (v4 - (v3 & (v4 * (v2 >> ((354) & 63)))))
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: -1}
            _ = v10
            _ = v14
        }
    } else {
        if (v1 == v1) {
            v15 := GenS3{F0: 42}
            _ = v15
        } else {
            if true {
                var v16 float32 = float32(0.1)
                _ = v16
            } else {
                v17 := GenS4{F0: 0}
                v18 := GenS2{F0: 0}
                _ = v17
                _ = v18
            }
        }
        v2 += -32768
    }
    if false {
        fmt.Println(v4)
        fmt.Println(v6, v3, v2)
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v2)
            v2 = v3
            v4 -= (v4 & 1)
            _ = fi19
        }
        if true {
            var v20 int32 = v6
            fmt.Println(v4)
            if (false || (65535 < 1)) {
                v5 = append(v5, v1)
                for fi21 := 0; fi21 < 3; fi21++ {
                    v22 := GenS2{F0: -1}
                    v2 = (^(v4))
                    v1 = ((v1 + ((v1 + v1) + v1)) + (v1 + (v1 + v1)))
                    _ = fi21
                    _ = v22
                }
            }
            _ = v20
        }
        v2 += -1
    }
    fmt.Println(v4, v1, v3)
    switch 1 {
    case 366:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -75:
        switch 32767 {
        case 821:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    default:
        for fi23 := 0; fi23 < 5; fi23++ {
            if (((v1 + v1) > v1) || true) {
                var v24 float32 = 0.0
                v25 := GenS3{F0: 100}
                var v26 int16 = v3
                _ = v24
                _ = v25
                _ = v26
            }
            if len(v5) > 0 {
                v5[0] = v1
            }
            _ = fi23
        }
    }
    v27 := GenS4{F0: 1}
    v1 = "\n\t"
    v1 = v1
    if false {
        v28 := GenS1{F0: 78, F1: 42}
        v1 = (v1 + v1)
        _ = v28
    } else {
        if len(v5) > 0 {
            v5[0] = v1
        }
        if (2147483647 <= 58) {
            v2 = -251
            var v29 uint = 1
            _ = v29
        } else {
            for fi30 := 0; fi30 < 19; fi30++ {
                fmt.Println(v6, v3, v4)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi30
            }
            var v31 int32 = (v6 >> ((45) & 63))
            _ = v31
        }
        switch 32767 {
        case 9223372036854775807:
            fmt.Println(v1)
        case 0:
            for fi32 := 0; fi32 < 18; fi32++ {
                for fi33 := 0; fi33 < 16; fi33++ {
                    if (4294967295 == 42) {
                        var v34 uint16 = 691
                        if len(v5) > 0 {
                            v5[0] = ((v1 + "a") + v1)
                        }
                        v35 := GenS5{F0: 9223372036854775807}
                        _ = v34
                        _ = v35
                    }
                    if (false || (!((18446744073709551615 >= 18446744073709551615) || (float64(0.1) <= 0.1)))) {
                        if len(v5) > 0 {
                            v5[0] = v1
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi33
                }
                _ = fi32
            }
        }
    }
    var v36 uint8 = 2
    if (42 < 0) {
        v36 -= (v36 ^ v36)
        var v37 bool = false
        var v38 uint16 = 2
        _ = v37
        _ = v38
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v39 := GenS4{F0: 674}
        v40 := GenS6{F0: false, F1: v36, F2: (v36 << ((873) & 63))}
        _ = v39
        _ = v40
    }
    var v41 uint = 0
    var v42 int32 = (v6 << ((18446744073709551615) & 63))
    switch 9223372036854775807 {
    case -806:
        fmt.Println(v36)
    }
    v42 -= 528
    var v43 int64 = -9223372036854775808
    var v44 int8 = 93
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
