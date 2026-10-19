package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 int, p1 uint64) int16 {
    v6 := GenS1{F0: -1, F1: 2147483647, F2: (^(p1))}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{}
        var v9 int32 = 2147483647
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    if (-9223372036854775808 < p0) {
        v11 := []int{(p0 << ((42) & 63))}
        fmt.Println(v10)
        _ = v11
    } else {
        fmt.Println(p1, p0)
        var v12 int64 = -9223372036854775808
        v13 := GenS2{F0: v10, F1: 0}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 42
}

func main() {
    var v1 float32 = -2.25
    var v2 float64 = float64((-2.25 + (v1 / v1)))
    var v3 float32 = (float32(1.5) + float32(0.0))
    var v4 int = -394
    v5 := []int{(v4 + -180), 0}
    switch "foo bar" {
    case "":
        if (false || true) {
            switch ("a" + "\n\t") {
            case "":
                fmt.Println(v1, v3)
            case "\n\t":
                fmt.Println(v4, v3, v3)
            default:
                var v14 uint32 = 386
                _ = v14
            }
            v3 -= float32(v2)
        } else {
            v15 := GenS2{F0: -1, F1: 1}
            _ = v15
        }
    case "hello":
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS3{F0: -375, F1: 65535}
            _ = fi16
            _ = v17
        }
    case "foo bar":
        v1 -= v1
    }
    var v18 float64 = 0.0
    v4 += 9223372036854775807
    v18 = v2
    var v19 uint32 = 871
    switch v4 {
    case 128:
        v3 = v1
    case 96:
        switch "foo bar" {
        case "\n\t":
            v2 += (v18 / v18)
        case "foo bar":
            fmt.Println(v18, v1)
        case "a":
            var v20 int8 = 119
            _ = v20
        default:
            v21 := GenS1{F0: 0, F1: -1, F2: 42}
            _ = v21
        }
    case 2147483647:
        v4 = (v4 ^ v4)
    default:
        for fi22 := 0; fi22 < 3; fi22++ {
            if (!true) {
                fmt.Println(v1)
            } else {
                v19 -= v19
                if false {
                    if false {
                        for fi23 := 0; fi23 < 10; fi23++ {
                            if len(v5) > 0 {
                                v5[0] = fi22
                            }
                            var v24 uint8 = 163
                            _ = fi23
                            _ = v24
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        var v25 bool = (!false)
                        var v26 int64 = 1
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v28 := GenS4{F0: 2147483647, F1: 80, F2: 42}
                        _ = len(v5)
                        v29 := GenS4{F0: (^(v19)), F1: 255, F2: 640}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
            }
            var v30 uint32 = 931
            for fi31 := 0; fi31 < 17; fi31++ {
                if false {
                    v2 = 0.0
                } else {
                    v19 = (v19 * (v30 * (v30 >> ((0) & 63))))
                    var v32 uint8 = 0
                    v5 = append(v5, (v4 >> ((65535) & 63)))
                    fmt.Println(v1, fi31)
                    _ = v32
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v33 int64 = -1
                _ = fi31
                _ = v33
            }
            _ = fi22
            _ = v30
        }
    }
    if true {
        if false {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 uint16 = 1
                v18 -= 1e10
                var v36 bool = false
                _ = fi34
                _ = v35
                _ = v36
            }
            fmt.Println(v18, v19, v4)
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        var v37 uint16 = 2
        v18 += (v18 * (1.5 / 0.1))
        v2 += v2
        _ = v37
    }
    v2 = (((0.0 * float64(3.14159)) * 0.0) - 3.14159)
    v2 = v18
    var v38 uint32 = v19
    v5 = append(v5, v4)
    fmt.Println(v38, v18)
    var v39 int8 = 120
    fmt.Println(v1, v39, v19)
    v19 = v19
    var v40 int16 = 1
    switch (v4 >> ((42) & 63)) {
    case -776:
        v38 = 0
    case 2147483647:
        if (201 != 651) {
            fmt.Println(v18, v2)
            var v41 uint16 = 65535
            var v42 int8 = -128
            _ = v41
            _ = v42
        }
    case 1:
        if (false || true) {
            var v43 int32 = 42
            _ = v43
        }
    default:
        var v44 int8 = (^(v39))
        _ = v44
    }
    var v45 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v5
}
