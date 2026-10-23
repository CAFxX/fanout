package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int64; F1 int; F2 uint16 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint32 }
func genFn0() int16 {
    var v6 int = -2147483648
    v6 += v6
    for fi7 := 0; fi7 < 5; fi7++ {
        v6 -= fi7
        switch "a" {
        case "":
            var v8 bool = (!true)
            _ = v8
        }
        v9 := []int{(fi7 >> ((255) & 63))}
        _ = fi7
        _ = v9
    }
    _ = v6
    return 1000
}

func main() {
    v1 := []uint32{42, 2147483648}
    var v2 int16 = 42
    v3 := []int{}
    var v4 uint16 = 0
    v5 := make(map[int]uint32)
    fmt.Println(v4)
    v5[0] = 263
    v4 += 65535
    for fi10 := 0; fi10 < 14; fi10++ {
        switch 2147483647 {
        case 2147483647:
            var v11 uint16 = v4
            _ = v11
        case 32767:
            v4 = (v4 + v4)
        }
        switch "hello" {
        case "":
            v12 := GenS1{F0: 9223372036854775807}
            _ = v12
        case "foo bar":
            switch (fi10 - fi10) {
            case 2147483647:
                v4 = (^(v4))
            case 616:
                v4 = 32767
            default:
                if len(v1) > 0 {
                    v1[0] = 557
                }
            }
        case "\n\t":
            v13 := GenS2{F0: 357, F1: (fi10 ^ (fi10 + fi10))}
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS2{F0: 0, F1: -2147483648}
    if len(v3) > 0 {
        v3[0] = 461
    }
    var v15 uint8 = 2
    switch ("\n\t" + "hello") {
    case "hello":
        v16 := GenS3{F0: 814, F1: 1, F2: (^(v4))}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 7; fi17++ {
            _ = len(v5)
            _ = fi17
        }
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        switch "\n\t" {
        case "foo bar":
            var v18 int64 = 42
            _ = v18
        case "a":
            switch "\n\t" {
            case "\n\t":
                v15 = 162
            case "foo bar":
                fmt.Println(v15)
            case "a":
                for fi19 := 0; fi19 < 20; fi19++ {
                    v3 = append(v3, fi19)
                    var v20 string = "hello"
                    _ = fi19
                    _ = v20
                }
            default:
                if false {
                    switch 42 {
                    case 0:
                        v4 = v4
                    case 32768:
                        var v21 uint64 = 18446744073709551615
                        _ = v21
                    case 200:
                        var v22 bool = false
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 10; fi23++ {
                            v4 = v4
                            fmt.Println(fi23, v4, fi23)
                            if (-1 >= 10) {
                                v4 += v4
                                _ = len(v3)
                                var v24 uint = 42
                                _ = v24
                            } else {
                                var v25 string = ("foo bar" + ((("foo bar" + "a") + ("hello" + (("" + ("foo bar" + "foo bar")) + "hello"))) + "a"))
                                fmt.Println(v25, v25, v15)
                                for fi26 := 0; fi26 < 3; fi26++ {
                                    var v27 uint64 = uint64(v4)
                                    _ = len(v3)
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            _ = fi23
                        }
                    }
                }
            }
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v15 -= (v15 & (v15 % ((v15) | 1)))
    v28 := GenS3{F0: -9223372036854775808, F1: 883, F2: v4}
    v29 := GenS4{F0: (v15 <= v15), F1: 0}
    for fi30 := 0; fi30 < 1; fi30++ {
        if false {
            v4 -= (v4 | 1)
            var v31 float64 = (0.0 - (1e10 * float64((((((-2.25 / (-2.25 * (float32(1.5) / ((-2.25 + 1e10) - 1e10)))) - (float32(((float64((0.0 / 1e10)) * (1.5 / 1.0)) - (0.1 + -2.25))) / ((-2.25 * 1.5) + -2.25))) * 1.5) / 1.0) + 3.14159))))
            _ = v31
        }
        v15 = (v15 * 1)
        v15 = v15
        _ = fi30
    }
    if (124 != -2147483648) {
        switch 905 {
        case 9223372036854775807:
            v1 = append(v1, 65535)
        case -645:
            v32 := GenS1{F0: -284}
            _ = v32
        case 42:
            v4 = (v4 - 32767)
        }
        var v33 int32 = 141
        _ = v33
    } else {
        if false {
            fmt.Println(v2, v4, v15)
        }
        if ((886 >= uint32(v4)) && ((0.0 * (1.5 * float64(((float32(3.14159) / ((1.5 - 1e10) + -2.25)) / 0.1)))) < (((0.0 * ((((1e10 - float64(0.0)) / 1e10) * (0.1 * 3.14159)) + -2.25)) - float64(0.1)) - 0.1))) {
            fmt.Println(v15, v4)
            delete(v5, 32767)
            v15 = 255
        }
    }
    if (v15 == (v15 * v15)) {
        _ = len(v3)
        for fi34 := 0; fi34 < 5; fi34++ {
            v2 = -32768
            _ = fi34
        }
    }
    if ((!(false && (((1.5 / (((-2.25 / 1.0) + 1.5) + float32(3.14159))) + float32(-2.25)) != 1.5))) || true) {
        for fi35 := 0; fi35 < 1; fi35++ {
            v36 := GenS5{F0: 42}
            v4 -= (v4 * 1)
            _ = fi35
            _ = v36
        }
        var v37 uint8 = (v15 << ((4294967295) & 63))
        _ = v37
    }
    v4 = (v4 / ((v4) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
