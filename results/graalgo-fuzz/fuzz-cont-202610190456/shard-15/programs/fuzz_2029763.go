package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 uint) int32 {
    var v6 uint64 = 42
    fmt.Println(v6)
    v7 := GenS1{F0: -815, F1: p1, F2: 464}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{2147483647, 65535, 1, 123}
    var v2 int16 = -32768
    var v3 int16 = v2
    var v4 float32 = -2.25
    var v5 uint32 = 0
    var v8 uint = 0
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v8, v4, v3)
    case "hello":
        switch -9223372036854775808 {
        case 892:
            for fi9 := 0; fi9 < 4; fi9++ {
                switch (fi9 << ((65535) & 63)) {
                case -2147483648:
                    v8 += (v8 ^ 1)
                }
                if false {
                    var v10 float64 = 1e10
                    _ = v10
                }
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(v5, v8)
                    switch -300 {
                    case 127:
                        fmt.Println(v5)
                    }
                    v8 -= 351
                    _ = fi11
                }
                _ = fi9
            }
        default:
            var v12 int8 = -1
            _ = v12
        }
    default:
        if false {
            for fi13 := 0; fi13 < 1; fi13++ {
                if ((v2 / ((v3) | 1)) == v2) {
                    if (true || (((v8 <= 1) && ((v4 != v4) || (0.0 == (float64(float32(0.1)) - float64(v4))))) && false)) {
                        v14 := GenS1{F0: genFn0(uint64(v5), v8), F1: v8, F2: (v5 / ((2147483647) | 1))}
                        v15 := GenS1{F0: -940, F1: v8, F2: 2147483647}
                        _ = v14
                        _ = v15
                    }
                    var v16 int16 = (v3 & (^(v3)))
                    _ = v16
                } else {
                    v5 = v5
                    if (uint16(v5) != 2) {
                        v4 += v4
                        var v17 uint8 = 83
                        var v18 string = ("a" + "foo bar")
                        _ = v17
                        _ = v18
                    }
                    var v19 bool = ((0.1 / 3.14159) <= 0.1)
                    _ = v19
                }
                _ = fi13
            }
        } else {
            if true {
                if ((-1 < 0) && (32767 != v2)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi20 := 0; fi20 < 8; fi20++ {
                        switch "a" {
                        case "hello":
                            var v21 int64 = 1
                            _ = v21
                        default:
                            fmt.Println(fi20, v3, v5)
                        }
                        if len(v1) > 0 {
                            v1[0] = uint32(v8)
                        }
                        _ = fi20
                    }
                }
                v22 := GenS1{F0: 2147483647, F1: v8, F2: (v5 * (^(v5)))}
                _ = v22
            }
            v23 := GenS2{F0: 128, F1: 65535, F2: 982}
            _ = v23
        }
    }
    var v24 int64 = -597
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v25 uint16 = 65535
    v26 := GenS2{F0: 254, F1: (v25 * v25), F2: 260}
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS3{F0: fi27, F1: -27}
        v4 += 1e10
        if true {
            v4 = 3.14159
            if false {
                switch "foo bar" {
                case "hello":
                    v24 += (v24 << ((uint(v25)) & 63))
                case "\n\t":
                    fmt.Println(v2, v24)
                case "foo bar":
                    v24 = (v24 + -9223372036854775808)
                }
                var v29 int = fi27
                _ = v29
            } else {
                v5 += (^(v5))
                var v30 uint64 = 503
                _ = v30
            }
            v31 := GenS3{F0: (fi27 | 32767), F1: -72}
            _ = v31
        } else {
            fmt.Println(v5, v4, v3)
            switch fi27 {
            case 32767:
                v25 += v25
            default:
                switch (("foo bar" + "") + "a") {
                case "foo bar":
                    v32 := GenS4{F0: "foo bar", F1: "a"}
                    _ = v32
                }
            }
        }
        _ = fi27
        _ = v28
    }
    switch 1 {
    case 127:
        var v33 int32 = genFn0(4294967295, uint(v5))
        _ = v33
    case 32768:
        _ = len(v1)
    case 925:
        if (32767 != 0) {
            fmt.Println(v5)
            var v34 int16 = (-(v3))
            _ = v34
        } else {
            var v35 bool = false
            _ = v35
        }
    default:
        v4 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
