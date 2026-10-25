package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint32 }
func genFn0() bool {
    var v7 uint16 = 0
    var v8 bool = false
    v9 := GenS1{F0: v7}
    v10 := GenS1{F0: (v7 * (v7 << ((214) & 63)))}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ((!v8) || (87 == 0))
}
func genFn1(p0 bool, p1 uint8, p2 int16) int16 {
    p1 -= p1
    v11 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]uint32)
    var v3 float64 = v1
    var v4 string = ((("a" + ("" + "foo bar")) + "") + "foo bar")
    var v5 float64 = v3
    var v6 int64 = 0
    fmt.Println(v5)
    v6 += -159
    v12 := []int{}
    v12 = append(v12, -21)
    switch 128 {
    case 1:
        if (!false) {
            v6 += v6
            v4 = ("" + (((v4 + (("hello" + v4) + v4)) + "hello") + ("a" + v4)))
            var v13 int64 = v6
            _ = v13
        }
    case -9223372036854775808:
        switch 42 {
        case 9223372036854775807:
            switch -9223372036854775808 {
            case 127:
                if false {
                    for fi14 := 0; fi14 < 11; fi14++ {
                        v15 := GenS1{F0: 16}
                        v16 := GenS1{F0: 255}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v17 string = ((v4 + "hello") + (v4 + ""))
                    _ = v17
                }
            case 128:
                if false {
                    v1 += 3.14159
                    switch 1 {
                    case 128:
                        if false {
                            fmt.Println(v5)
                            v6 += (-(v6))
                            v1 = v3
                        } else {
                            v3 = 0.1
                            for fi18 := 0; fi18 < 5; fi18++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v12 = append(v12, (fi18 - fi18))
                                case "a":
                                    v6 += (-(v6))
                                case "\n\t":
                                    var v19 uint8 = 1
                                    _ = v19
                                default:
                                    v20 := GenS2{F0: 2147483647, F1: v6}
                                    _ = v20
                                }
                                if true {
                                    v21 := GenS2{F0: 2147483648, F1: (v6 * (v6 * -1))}
                                    var v22 int32 = -2147483648
                                    fmt.Println(v22, v22)
                                    _ = v21
                                    _ = v22
                                }
                                var v23 int8 = 127
                                _ = fi18
                                _ = v23
                            }
                            switch (("" + ("hello" + v4)) + v4) {
                            case "hello":
                                var v24 uint8 = 205
                                _ = v24
                            }
                        }
                    case 0:
                        v25 := GenS2{F0: 2147483648, F1: 9223372036854775807}
                        _ = v25
                    }
                }
            case -1:
                v26 := GenS2{F0: 0, F1: v6}
                _ = v26
            }
        case 835:
            v27 := GenS1{F0: 803}
            _ = v27
        }
    case -1:
        if (1 >= 65535) {
            v6 = 52
        } else {
            if (-9223372036854775808 > (v6 << ((662) & 63))) {
                v3 = v1
                switch 196 {
                case 4294967295:
                    v5 = v3
                case -1:
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 float32 = 3.14159
                        var v30 uint32 = 65535
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                case 2147483647:
                    var v31 int8 = -1
                    _ = v31
                }
                v32 := GenS1{F0: 180}
                _ = v32
            } else {
                v1 = (v1 / (float64(1e10) - v5))
            }
            var v33 int8 = 1
            _ = v33
        }
    default:
        delete(v2, 4294967295)
    }
    v4 = v4
    var v34 int16 = -573
    _ = len(v12)
    for fi35 := 0; fi35 < 2; fi35++ {
        switch (fi35 % ((32767) | 1)) {
        case -406:
            var v36 float64 = ((float64((0.0 + (3.14159 / (0.1 - float32(float64((1.5 / 0.1))))))) + 0.1) / 3.14159)
            _ = v36
        case 595:
            v37 := v2[2]
            _ = v37
        case 49:
            var v38 float64 = ((((v1 / v5) + v1) - v5) - (v3 * 3.14159))
            _ = v38
        default:
            v6 += -754
        }
        v39 := GenS2{F0: 2147483647, F1: 1}
        var v40 float32 = ((1.5 + float32(3.14159)) - 0.0)
        _ = fi35
        _ = v39
        _ = v40
    }
    if (1 < 127) {
        delete(v2, 1)
        var v41 uint32 = 65535
        v42 := GenS3{F0: (v41 | (v41 * (v41 >> ((18446744073709551615) & 63))))}
        _ = v41
        _ = v42
    }
    var v43 uint8 = 42
    switch ((("a" + "a") + v4) + v4) {
    case "\n\t":
        v44 := GenS1{F0: 0}
        _ = v44
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        v5 -= 0.0
    }
    var v45 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v45
    _ = v5
    _ = v6
}
