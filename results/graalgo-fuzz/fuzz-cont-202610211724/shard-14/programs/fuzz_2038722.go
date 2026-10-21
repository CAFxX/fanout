package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 uint16, p1 string, p2 bool) uint64 {
    var v7 uint16 = 32767
    fmt.Println(v7, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{42, 233, -9223372036854775808, -9223372036854775808}
    var v3 int = 32768
    var v4 int64 = 0
    var v5 int32 = v1
    var v6 int64 = -534
    v6 = (v4 & -174)
    if false {
        v8 := GenS1{F0: 42, F1: 2}
        v5 = (^(v1))
        var v9 string = "foo bar"
        _ = v8
        _ = v9
    }
    if (127 < 42) {
        v10 := GenS1{F0: 127, F1: 255}
        var v11 float64 = ((3.14159 - 3.14159) / (1.5 * -2.25))
        _ = v10
        _ = v11
    } else {
        if (!(!false)) {
            v1 = v1
        } else {
            v12 := GenS2{F0: 1, F1: 255}
            v1 += v1
            _ = v12
        }
        v13 := GenS3{F0: v5}
        v4 += v4
        _ = v13
    }
    v5 += (v5 >> ((2) & 63))
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: 127, F1: 64}
        var v16 uint8 = 0
        v4 -= (v4 - (v6 - (v4 | -9223372036854775808)))
        _ = fi14
        _ = v15
        _ = v16
    }
    v6 = 307
    fmt.Println(v4, v1, v1)
    v17 := GenS4{F0: 65535, F1: 4294967295, F2: -52}
    for fi18 := 0; fi18 < 15; fi18++ {
        v6 = -241
        v1 = 1
        if (false && (true || true)) {
            v3 += v3
            fmt.Println(v5)
            if (!(!("foo bar" != "foo bar"))) {
                v19 := GenS3{F0: v1}
                _ = v19
            }
        } else {
            for fi20 := 0; fi20 < 20; fi20++ {
                if true {
                    if (!(203 != 157)) {
                        fmt.Println(fi20)
                    }
                    v1 = -453
                    if ("" <= "\n\t") {
                        switch ("a" + "a") {
                        case "a":
                            var v21 uint8 = 127
                            _ = v21
                        case "foo bar":
                            fmt.Println(v4)
                        default:
                            _ = len(v2)
                        }
                    } else {
                        v4 -= -1
                        if false {
                            fmt.Println(v5, fi18, v5)
                        }
                    }
                } else {
                    var v22 uint = 65535
                    v5 += v1
                    fmt.Println(v1, v4, v3)
                    var v23 int16 = 1000
                    var v24 uint64 = 18446744073709551615
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi20
            }
            v6 = v6
            v1 += (v1 ^ v5)
        }
        _ = fi18
    }
    if (!(!false)) {
        var v25 uint = 18446744073709551615
        v1 -= v5
        _ = v25
    } else {
        var v26 int64 = (v4 % (((v6 & v4)) | 1))
        _ = v26
    }
    switch (v3 * v3) {
    case -970:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 893:
        if false {
            var v27 uint8 = 1
            var v28 int64 = (v6 - (v4 * (v6 + v4)))
            _ = v27
            _ = v28
        }
    case 42:
        v1 -= v1
    default:
        v4 = (v4 & v4)
    }
    if true {
        v4 = v6
        var v29 int8 = -1
        if false {
            var v30 uint32 = 1
            _ = v30
        } else {
            if len(v2) > 0 {
                v2[0] = (^(v3))
            }
            _ = len(v2)
        }
        _ = v29
    }
    if (568 > 0) {
        if true {
            v2 = append(v2, (^(v3)))
            var v31 uint8 = 128
            var v32 int32 = v1
            _ = v31
            _ = v32
        }
        if false {
            var v33 int16 = 136
            var v34 bool = false
            _ = v33
            _ = v34
        } else {
            fmt.Println(v1, v4)
            var v35 uint64 = genFn0(799, "foo bar", ((127 >= 42) && true))
            var v36 float32 = (float32(0.1) - 0.0)
            _ = v35
            _ = v36
        }
    } else {
        _ = len(v2)
    }
    fmt.Println(v4)
    if true {
        if true {
            if (42 < 46) {
                v37 := GenS2{F0: 255, F1: 2}
                if (2147483647 != 4294967295) {
                    var v38 uint64 = 42
                    var v39 bool = ((!("a" != ("a" + "hello"))) || false)
                    v1 -= int32(v4)
                    _ = v38
                    _ = v39
                }
                _ = v37
            }
            var v40 int32 = v5
            if true {
                _ = len(v2)
                _ = len(v2)
                v4 += v4
            } else {
                var v41 int8 = -1
                _ = v41
            }
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
