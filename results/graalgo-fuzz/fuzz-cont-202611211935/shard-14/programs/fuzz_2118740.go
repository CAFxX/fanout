package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 int; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
func genFn0() uint8 {
    v6 := GenS1{F0: 255}
    var v7 int = -2147483648
    v8 := GenS2{F0: 255, F1: 42, F2: 1}
    fmt.Println(v7)
    v7 = 32767
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 string, p1 int16, p2 int) uint {
    if ((1.5 != 1e10) || (2147483647 != 734)) {
        fmt.Println(p0, p0, p1)
    }
    var v9 int = p2
    v10 := GenS1{F0: 52}
    p1 = (p1 / ((671) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 124
}

func main() {
    v1 := []int{-9223372036854775808, 0, -821, 32767}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 int16 = 0
    v5 := []int{}
    for fi11 := 0; fi11 < 1; fi11++ {
        v4 = v4
        _ = fi11
    }
    fmt.Println(v4)
    var v12 uint32 = 851
    var v13 uint8 = 134
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint = 16
        v12 += (v12 >> ((2) & 63))
        if (42 <= 1) {
            v16 := GenS3{F0: (v4 * -451), F1: (fi14 * fi14), F2: int8(fi14)}
            var v17 string = ""
            var v18 uint32 = (v12 - (v12 ^ 2147483648))
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v19 := v3[(v12 & 2147483647)]
    v5 = append(v5, 1)
    if v19 {
        fmt.Println(v4, v13, v12)
        var v20 string = ""
        if (((1e10 * (-2.25 * float64(3.14159))) * 1.5) != float64(1.5)) {
            if false {
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 string = v20
                        if (127 >= 1) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v24 bool = v19
                            fmt.Println(fi22, v12, v20)
                            _ = v24
                        }
                        v12 = v12
                        fmt.Println(v13, v4)
                        v25 := GenS2{F0: 65535, F1: (v12 + (^(v12))), F2: 2}
                        _ = fi22
                        _ = v23
                        _ = v25
                    }
                    if (float64((1e10 + 1e10)) != (0.0 - 0.0)) {
                        var v26 int8 = -51
                        switch v20 {
                        case "\n\t":
                            v4 += v4
                        default:
                            v27 := GenS2{F0: 134, F1: v12, F2: 582}
                            _ = v27
                        }
                        v28 := GenS1{F0: v13}
                        _ = v26
                        _ = v28
                    } else {
                        if false {
                            if (4294967295 == 42) {
                                var v29 float64 = 0.0
                                var v30 int32 = 0
                                _ = v29
                                _ = v30
                            } else {
                                v13 -= v13
                                v13 = 213
                                v19 = true
                            }
                            for fi31 := 0; fi31 < 13; fi31++ {
                                v32 := GenS2{F0: 1, F1: (v12 - (v12 % (((v12 - (v12 ^ 2147483648))) | 1))), F2: 1}
                                _ = fi31
                                _ = v32
                            }
                        }
                        v12 = 443
                    }
                    _ = fi21
                }
                v12 = v12
                switch (v20 + (v20 + v20)) {
                case "\n\t":
                    var v33 uint8 = v13
                    _ = v33
                case "a":
                    var v34 int = -675
                    _ = v34
                case "foo bar":
                    v2[v20] = (v12 + 65535)
                }
            } else {
                if (!v19) {
                    v35 := GenS4{F0: ((v20 + v20) + ("" + "hello"))}
                    var v36 int64 = -9223372036854775808
                    fmt.Println(v19, v13)
                    _ = v35
                    _ = v36
                } else {
                    v37 := v3[(^(v12))]
                    var v38 uint64 = 2
                    v39 := v3[(^(v12))]
                    _ = v37
                    _ = v38
                    _ = v39
                }
            }
            var v40 int16 = v4
            if (58 >= 4294967295) {
                var v41 bool = true
                _ = v41
            }
            _ = v40
        }
        _ = v20
    } else {
        for fi42 := 0; fi42 < 17; fi42++ {
            v4 += v4
            _ = fi42
        }
        if ((-(v4)) != v4) {
            var v43 float32 = (0.1 - ((3.14159 / 1.5) + float32(0.0)))
            var v44 float32 = (v43 / (3.14159 + 1e10))
            _ = v43
            _ = v44
        } else {
            if (!(42 >= 0)) {
                for fi45 := 0; fi45 < 1; fi45++ {
                    v12 = 65535
                    _ = fi45
                }
                delete(v3, v12)
                v13 = 42
            } else {
                if v19 {
                    var v46 string = "\n\t"
                    v2[v46] = 504
                    v47 := GenS1{F0: (v13 + (v13 << ((984) & 63)))}
                    _ = v46
                    _ = v47
                } else {
                    var v48 uint = 4294967295
                    var v49 uint32 = (^(v12))
                    _ = v48
                    _ = v49
                }
                v2[(("foo bar" + "a") + "foo bar")] = (v12 + (v12 << ((2) & 63)))
            }
            fmt.Println(v12, v4)
        }
    }
    if ("a" >= "\n\t") {
        var v50 uint32 = v12
        _ = v50
    } else {
        for fi51 := 0; fi51 < 6; fi51++ {
            v1 = append(v1, 32767)
            _ = fi51
        }
        var v52 uint64 = 0
        _ = v52
    }
    v53 := GenS1{F0: genFn0()}
    v54 := GenS4{F0: "hello"}
    v55 := GenS5{F0: 1}
    v4 = 42
    v56 := GenS5{F0: 360}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
