package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint32; F1 int16; F2 uint8 }
func genFn0(p0 int32, p1 uint, p2 int64) uint8 {
    var v7 int16 = 225
    v8 := []int{42, -750, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 140
}

func main() {
    v1 := []string{"hello", "", "a", "\n\t"}
    var v2 uint64 = 4294967295
    v3 := []int{-9223372036854775808}
    var v4 float64 = 1e10
    var v5 float64 = v4
    var v6 uint8 = uint8(v2)
    fmt.Println(v2)
    if (!false) {
        fmt.Println(v5)
        if false {
            fmt.Println(v4, v2)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (false && (!false)) {
        if true {
            v5 += v5
            v9 := GenS1{F0: 4294967295, F1: 42}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi10
    }
    _ = len(v3)
    switch 4294967295 {
    case 128:
        v11 := GenS2{F0: 1}
        _ = v11
    case 9223372036854775807:
        var v12 bool = true
        _ = v12
    }
    v13 := GenS1{F0: -730, F1: 4294967295}
    v1 = append(v1, "")
    switch "hello" {
    case "hello":
        if false {
            v6 = (v6 % ((v6) | 1))
            switch "" {
            case "":
                var v14 float32 = ((0.1 / -2.25) - 3.14159)
                _ = v14
            case "hello":
                var v15 int = -1
                _ = v15
            }
        } else {
            if (false && (true || (2147483647 >= 306))) {
                if (!false) {
                    _ = len(v3)
                    switch 9223372036854775807 {
                    case 638:
                        v4 += 1e10
                    case 1:
                        var v16 uint = 4294967295
                        _ = v16
                    }
                }
                if true {
                    if false {
                        v6 -= v6
                        v17 := GenS3{F0: v6, F1: (v2 * v2), F2: 71}
                        _ = v17
                    } else {
                        var v18 uint32 = 403
                        _ = v18
                    }
                    v19 := GenS4{F0: 4294967295, F1: -1, F2: int64(v6)}
                    v6 = v6
                    _ = v19
                }
                if len(v3) > 0 {
                    v3[0] = 528
                }
            }
            var v20 int64 = 9223372036854775807
            var v21 uint = 1
            _ = v20
            _ = v21
        }
    case "\n\t":
        fmt.Println(v4, v4, v5)
    case "a":
        v5 = v4
    default:
        v4 -= v4
    }
    if false {
        v22 := GenS2{F0: 569}
        if false {
            for fi23 := 0; fi23 < 6; fi23++ {
                v2 = 0
                switch 4294967295 {
                case -1:
                    v4 += v4
                default:
                    fmt.Println(v6, v6, fi23)
                }
                _ = fi23
            }
            if (3.14159 > v4) {
                var v24 int = -2147483648
                v25 := GenS4{F0: 4294967295, F1: 127, F2: 1}
                _ = len(v3)
                _ = v24
                _ = v25
            }
            _ = len(v3)
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            if true {
                var v27 int8 = 100
                _ = v27
            } else {
                v2 -= v2
                if false {
                    v28 := GenS5{F0: 1, F1: 11}
                    var v29 uint16 = 964
                    _ = v28
                    _ = v29
                } else {
                    v2 += v2
                }
            }
            var v30 int64 = 1
            var v31 string = "hello"
            _ = fi26
            _ = v30
            _ = v31
        }
        _ = v22
    } else {
        if (-1 <= 100) {
            v6 += v6
            if (42 >= (v6 - (v6 << ((18446744073709551615) & 63)))) {
                var v32 int64 = 42
                _ = v32
            }
            _ = len(v1)
        } else {
            v5 += 3.14159
        }
        fmt.Println(v5, v6)
        v33 := GenS1{F0: -2147483648, F1: 18446744073709551615}
        _ = v33
    }
    switch ("" + "a") {
    case "foo bar":
        for fi34 := 0; fi34 < 15; fi34++ {
            v2 += (v2 << ((394) & 63))
            if ((!(117 > 127)) || true) {
                fmt.Println(v6, v5, v5)
                var v35 uint64 = (v2 << ((19) & 63))
                var v36 uint = 255
                _ = v35
                _ = v36
            }
            _ = fi34
        }
    case "":
        var v37 uint = 474
        _ = v37
    case "\n\t":
        var v38 int32 = 612
        _ = v38
    }
    var v39 float64 = v5
    var v40 uint64 = v2
    v39 -= (0.1 - ((float64(0.0) + v5) + 0.0))
    v41 := GenS6{F0: 2147483648, F1: -1, F2: (v6 >> ((uint(v40)) & 63))}
    var v42 float64 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
