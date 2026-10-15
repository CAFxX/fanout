package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0(p0 bool) uint {
    fmt.Println(p0)
    var v6 uint32 = 2
    var v7 int16 = -921
    p0 = (1e10 <= (0.1 / 1e10))
    v7 += (v7 | (v7 | -32768))
    _ = p0
    _ = v6
    _ = v7
    return 733
}
func genFn1(p0 int8, p1 int) bool {
    for fi8 := 0; fi8 < 4; fi8++ {
        switch -1 {
        case 127:
            p0 = p0
        case -419:
            var v9 int = p1
            _ = v9
        case 128:
            if ((true && (2147483647 > 659)) || false) {
                p0 += 42
                var v10 int32 = 1
                _ = v10
            } else {
                p1 = (-(fi8))
            }
        default:
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 string = "foo bar"
                var v13 float32 = float32(float64((float32((float64(0.0) / (0.0 + 0.1))) - (0.0 / 3.14159))))
                switch ((("a" + "") + v12) + "\n\t") {
                case "hello":
                    v12 = v12
                case "a":
                    p0 = (p0 * (p0 - p0))
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi8
    }
    var v14 uint16 = 2
    for fi15 := 0; fi15 < 8; fi15++ {
        switch p1 {
        case -2147483648:
            p1 = 56
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v14
    return false
}

func main() {
    var v1 int32 = -675
    var v2 uint16 = 255
    v3 := make(map[uint32]int)
    var v4 int8 = 127
    v5 := []string{"hello", (("hello" + (("foo bar" + "") + "")) + ("\n\t" + "\n\t")), ("hello" + "")}
    v1 -= v1
    var v16 int32 = v1
    if (-2.25 > ((0.1 * (((((0.0 + 0.0) + (1e10 - 1.5)) * 1.5) - (float64(float32((-2.25 + 0.0))) * (float64(float32(1e10)) / 1.0))) / -2.25)) / 1.0)) {
        switch -665 {
        case 4294967295:
            fmt.Println(v1, v2)
        case -9223372036854775808:
            v17 := GenS1{F0: 760, F1: 1, F2: 1}
            _ = v17
        default:
            if (!false) {
                v16 = (v16 / (((v1 * 0)) | 1))
                var v18 int16 = 1
                if (false || (float32(3.14159) < (float32(-2.25) / 1.5))) {
                    fmt.Println(v18, v2)
                } else {
                    v16 += v16
                    if (!(2 > 843)) {
                        fmt.Println(v16)
                        switch "\n\t" {
                        case "hello":
                            if (true && true) {
                                fmt.Println(v2, v4)
                                if false {
                                    var v19 int32 = v1
                                    _ = v19
                                }
                                v1 -= v16
                            }
                        case "a":
                            v16 -= -2147483648
                        case "foo bar":
                            var v20 uint32 = 42
                            _ = v20
                        }
                    }
                }
                _ = v18
            } else {
                v21 := GenS1{F0: 2, F1: 42, F2: 18446744073709551615}
                var v22 uint16 = v2
                _ = v21
                _ = v22
            }
        }
        v2 -= 82
        v23 := GenS2{F0: -925}
        _ = v23
    } else {
        fmt.Println(v2, v4, v2)
        switch 750 {
        case -915:
            fmt.Println(v2, v16)
        case 4294967295:
            if true {
                var v24 uint8 = 2
                v25 := GenS2{F0: -201}
                _ = v24
                _ = v25
            }
        case -2147483648:
            var v26 float64 = (((0.0 * -2.25) / ((0.0 * 3.14159) + 0.1)) / (1e10 + 1.5))
            _ = v26
        default:
            v27 := GenS1{F0: 255, F1: 166, F2: 1}
            _ = v27
        }
    }
    var v28 float32 = float32((3.14159 + (0.0 - (0.1 * 0.0))))
    var v29 string = "\n\t"
    if true {
        var v30 string = v29
        var v31 float64 = 1e10
        _ = v30
        _ = v31
    } else {
        fmt.Println(v1, v16)
    }
    fmt.Println(v1)
    var v32 float32 = v28
    v1 -= v16
    v1 += v1
    v33 := GenS3{F0: (v2 & (v2 - 2)), F1: 22, F2: 2}
    var v34 int8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
