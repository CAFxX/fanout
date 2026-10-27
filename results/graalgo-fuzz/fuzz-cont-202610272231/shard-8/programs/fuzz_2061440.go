package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int64) uint32 {
    for fi6 := 0; fi6 < 6; fi6++ {
        switch fi6 {
        case 0:
            var v7 int16 = 1
            _ = v7
        case 32767:
            p0 = p0
        case 9223372036854775807:
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int8 = 51
                v10 := GenS1{F0: 894}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
    }
    v11 := []int{-183}
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            var v13 int = fi12
            if true {
                if true {
                    switch v13 {
                    case 875:
                        p0 = p0
                    case -9223372036854775808:
                        if (65535 > 18446744073709551615) {
                            var v14 int64 = 0
                            var v15 int = v13
                            _ = v14
                            _ = v15
                        }
                    case 32767:
                        var v16 uint16 = 255
                        _ = v16
                    default:
                        v13 = (fi12 + (v13 * -355))
                    }
                }
                v17 := GenS2{F0: (p0 | 0)}
                if (42 >= 18446744073709551615) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        fmt.Println(p0)
                        _ = fi18
                    }
                    v19 := GenS2{F0: 9223372036854775807}
                    var v20 uint8 = 1
                    _ = v19
                    _ = v20
                } else {
                    p0 = (p0 / ((p0) | 1))
                    if (!true) {
                        v13 += v13
                        v21 := GenS1{F0: 1}
                        if (true || true) {
                            var v22 float64 = (float64(1.5) / 3.14159)
                            for fi23 := 0; fi23 < 3; fi23++ {
                                v24 := GenS3{F0: 32767}
                                p0 = (p0 & p0)
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                v22 = (v22 + 3.14159)
                                _ = fi23
                                _ = v24
                            }
                            _ = v22
                        }
                        _ = v21
                    } else {
                        v13 += fi12
                        var v25 float32 = 3.14159
                        _ = v25
                    }
                }
                _ = v17
            } else {
                var v26 uint8 = 142
                _ = v26
            }
            if (false || true) {
                v27 := GenS3{F0: 42}
                _ = v27
            }
            _ = v13
        } else {
            var v28 uint32 = 2147483647
            v29 := GenS1{F0: (v28 - v28)}
            if false {
                var v30 bool = (!((fi12 & fi12) < (fi12 << ((90) & 63))))
                _ = v30
            } else {
                v31 := GenS3{F0: -1}
                v28 -= 42
                var v32 int64 = (p0 / ((p0) | 1))
                _ = v31
                _ = v32
            }
            _ = v28
            _ = v29
        }
        var v33 int32 = 1
        v11 = append(v11, -497)
        _ = fi12
        _ = v33
    }
    _ = p0
    _ = v11
    return 2
}

func main() {
    v1 := []int{0, 128, -2147483648, 1}
    v2 := []int{222, 4294967295}
    var v3 bool = (1 <= -2147483648)
    var v4 uint32 = 2147483648
    var v5 uint32 = v4
    var v34 uint32 = v4
    v5 -= (v5 << ((42) & 63))
    if ("hello" < "") {
        var v35 uint8 = 23
        var v36 uint = 468
        _ = v35
        _ = v36
    }
    if v3 {
        var v37 uint32 = v34
        _ = v37
    }
    var v38 float32 = (3.14159 * 0.1)
    for fi39 := 0; fi39 < 19; fi39++ {
        var v40 float32 = (v38 + float32(1.5))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v41 int16 = 1
        _ = fi39
        _ = v40
        _ = v41
    }
    v42 := GenS3{F0: 0}
    for fi43 := 0; fi43 < 19; fi43++ {
        switch "\n\t" {
        case "\n\t":
            v4 -= (v5 - (v34 ^ 1))
        case "hello":
            fmt.Println(v34)
        }
        switch "" {
        case "\n\t":
            v44 := GenS2{F0: 9223372036854775807}
            _ = v44
        }
        var v45 string = ""
        _ = fi43
        _ = v45
    }
    v38 = (-2.25 - ((v38 / v38) + 1e10))
    v4 = v4
    v3 = v3
    v34 += v4
    v34 -= v4
    v46 := GenS3{F0: -1}
    v38 = (((v38 + (v38 + v38)) - ((v38 + 3.14159) * 0.1)) / v38)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v42
    _ = v46
    _ = v5
}
