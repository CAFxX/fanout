package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint }
func genFn0(p0 int32) int64 {
    if (!(false || false)) {
        v5 := []int{9223372036854775807, 2147483647, 127}
        var v6 int32 = (-(p0))
        _ = v5
        _ = v6
    } else {
        var v7 uint = 18446744073709551615
        v8 := []int{32767, 0, -408, 42}
        _ = v7
        _ = v8
    }
    if (!(true || (147 < 171))) {
        p0 -= 0
        var v9 uint8 = 42
        p0 += (p0 | (^(p0)))
        _ = v9
    } else {
        if (1 >= 0) {
            p0 = -864
            var v10 uint = 18446744073709551615
            _ = v10
        }
    }
    v11 := []int{}
    v11 = append(v11, 1)
    _ = p0
    _ = v11
    return 1
}

func main() {
    v1 := []uint32{2147483648, 65535}
    v2 := make(map[string]uint32)
    v3 := []string{}
    var v4 uint64 = 616
    var v12 float32 = (float32(-2.25) * 1e10)
    v3 = append(v3, (("a" + (("a" + ("hello" + "foo bar")) + ("foo bar" + "\n\t"))) + ("hello" + "\n\t")))
    fmt.Println(v4, v4)
    fmt.Println(v4)
    if true {
        fmt.Println(v4)
        var v13 int8 = -128
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 32767, F1: -1}
            var v16 string = ""
            if true {
                switch 127 {
                case 1:
                    switch "\n\t" {
                    case "":
                        v12 -= 0.1
                    case "hello":
                        var v17 float64 = (float64(0.1) + 3.14159)
                        _ = v17
                    default:
                        delete(v2, v16)
                    }
                case -958:
                    var v18 int32 = 0
                    _ = v18
                }
            } else {
                if false {
                    fmt.Println(v12)
                }
                v3 = append(v3, "\n\t")
                for fi19 := 0; fi19 < 10; fi19++ {
                    switch "hello" {
                    case "":
                        v4 = (v4 << ((0) & 63))
                    case "foo bar":
                        _ = len(v1)
                        switch (fi14 + fi19) {
                        case -383:
                            delete(v2, "hello")
                        case -1:
                            fmt.Println(v16)
                        default:
                            v20 := GenS2{F0: false, F1: 42}
                            _ = v20
                        }
                    }
                    fmt.Println(fi14)
                    v21 := GenS1{F0: 0, F1: 9223372036854775807}
                    var v22 uint16 = 2
                    v13 = v13
                    var v23 uint16 = (^(v22))
                    _ = fi19
                    _ = v21
                    _ = v22
                    _ = v23
                }
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v12 += -2.25
    v1 = append(v1, 299)
    var v24 int = 127
    v12 += 0.1
    v25 := GenS1{F0: -1, F1: 0}
    if (255 > 65535) {
        v4 = (v4 % ((v4) | 1))
        var v26 float32 = v12
        _ = v26
    }
    _ = len(v2)
    v12 = ((1e10 / float32(((0.1 / -2.25) + 0.0))) * (v12 + float32((((-2.25 * float64((float32((1.5 - float64(v12))) * v12))) * 3.14159) * float64((v12 - v12))))))
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 uint16 = 316
        if true {
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 int16 = 181
                var v31 int8 = 1
                var v32 float32 = (v12 - 3.14159)
                v33 := GenS1{F0: int16(v28), F1: 0}
                var v34 int16 = 42
                v24 = fi29
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            if false {
                var v35 float64 = 0.1
                if (v12 != 1e10) {
                    for fi36 := 0; fi36 < 4; fi36++ {
                        if true {
                            fmt.Println(v28)
                            if false {
                                switch (fi36 << ((65535) & 63)) {
                                case 1:
                                    v3 = append(v3, "")
                                case 270:
                                    v35 = v35
                                    switch (((("hello" + "foo bar") + ("a" + "foo bar")) + "hello") + "a") {
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "foo bar":
                                        fmt.Println(v28)
                                    case "":
                                        v37 := GenS3{F0: 42}
                                        _ = v37
                                    }
                                    v12 -= (v12 / 1.0)
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                default:
                                    fmt.Println(v4)
                                }
                                _ = len(v2)
                            }
                        }
                        var v38 float64 = (0.1 * 1.5)
                        _ = fi36
                        _ = v38
                    }
                }
                v2["hello"] = 73
                _ = v35
            } else {
                v24 = -1
                v39 := GenS4{F0: false, F1: 1, F2: 42}
                _ = v39
            }
        } else {
            var v40 uint8 = 0
            _ = len(v1)
            for fi41 := 0; fi41 < 8; fi41++ {
                var v42 string = ""
                _ = fi41
                _ = v42
            }
            _ = v40
        }
        v43 := v2[("" + "a")]
        _ = fi27
        _ = v28
        _ = v43
    }
    v12 = ((float32((1.5 + float64((1e10 - (((1.5 - v12) + (v12 - 1.5)) / (v12 * ((v12 / 1.0) / 3.14159))))))) / v12) + 1.5)
    var v44 uint32 = 1
    var v45 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
