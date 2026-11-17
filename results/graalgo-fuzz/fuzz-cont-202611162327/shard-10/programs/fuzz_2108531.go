package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint, p1 int16, p2 int64) uint32 {
    var v6 int = -960
    switch (("foo bar" + ("a" + "a")) + "") {
    case "\n\t":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 bool = (p0 != (p0 / ((1) | 1)))
            var v9 uint16 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        var v10 uint64 = 18446744073709551615
        _ = v10
    default:
        v11 := GenS1{F0: "", F1: (p2 | (^(p2)))}
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: (p2 & (p2 & (p2 << ((p0) & 63))))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 65535
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    var v4 int32 = -2147483648
    v5 := make(map[int]uint32)
    v4 = (v4 + 2147483647)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int = -1
        v3 = (^(v3))
        _ = fi13
        _ = v14
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v3, v4)
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS2{F0: -88, F1: 1}
        v1 = append(v1, (fi15 | (fi15 + fi15)))
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: v3, F1: (2147483647 == -9223372036854775808), F2: 255}
    fmt.Println(v4)
    if (true && (false && ((42 <= 1) || true))) {
        fmt.Println(v4)
        for fi18 := 0; fi18 < 9; fi18++ {
            if (!((fi18 << ((18446744073709551615) & 63)) == (fi18 | (^(fi18))))) {
                v3 = v3
            } else {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if (-128 != 127) {
                        var v20 int16 = -32768
                        var v21 string = ("\n\t" + "a")
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v3, fi19)
                    _ = fi19
                }
            }
            _ = len(v5)
            _ = fi18
        }
        v3 = 820
    }
    v4 += v4
    delete(v2, 2147483647)
    var v22 int = 4294967295
    v3 -= 2
    var v23 string = ((("foo bar" + (("\n\t" + "foo bar") + "hello")) + "hello") + ("" + "a"))
    v24 := GenS3{F0: (^(v3)), F1: true, F2: (v3 + v3)}
    if false {
        v23 = (v23 + "hello")
        for fi25 := 0; fi25 < 16; fi25++ {
            switch ((((("" + ("\n\t" + v23)) + "hello") + "a") + v23) + ("hello" + ("foo bar" + "hello"))) {
            case "":
                var v26 uint64 = 0
                _ = v26
            case "a":
                fmt.Println(v4)
            case "foo bar":
                if (v4 >= -1) {
                    for fi27 := 0; fi27 < 17; fi27++ {
                        delete(v5, -2147483648)
                        var v28 uint64 = 42
                        fmt.Println(v22)
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v29 := GenS3{F0: (v3 >> ((652) & 63)), F1: (!(542 == 42)), F2: v3}
                    _ = v29
                }
            }
            switch ((v23 + (v23 + "")) + v23) {
            case "\n\t":
                if true {
                    switch (v23 + v23) {
                    case "hello":
                        switch v23 {
                        case "foo bar":
                            var v30 uint16 = (^(v3))
                            _ = v30
                        case "hello":
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v32 := GenS4{F0: 42, F1: 0}
                                var v33 int64 = 0
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        default:
                            var v34 int32 = (v4 * (v4 + (v4 - -366)))
                            _ = v34
                        }
                    }
                    v22 = (-(fi25))
                    if (float32(1.5) > ((((-2.25 / (3.14159 - 0.1)) / ((0.1 + float32(((1e10 * 1e10) + ((float64(float32(0.1)) / -2.25) * -2.25)))) / 1.0)) / (((0.1 / (1.5 - 3.14159)) - ((1.5 + (0.0 / float32((-2.25 / float64(0.1))))) / 3.14159)) + 3.14159)) / 1.0)) {
                        v35 := GenS5{F0: 107, F1: ((genFn0(42, -32768, -9223372036854775808) <= 1) || (!true)), F2: 2}
                        _ = v35
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (false || (v23 == v23)) {
                            v36 := GenS6{F0: 562}
                            v22 = (v22 & v22)
                            _ = v36
                        } else {
                            v37 := GenS2{F0: 56, F1: -9223372036854775808}
                            _ = v37
                        }
                    }
                } else {
                    v38 := GenS6{F0: 1}
                    v22 -= 2147483647
                    _ = v38
                }
            }
            delete(v5, v22)
            _ = fi25
        }
        switch "\n\t" {
        case "a":
            var v39 int64 = 0
            _ = v39
        case "\n\t":
            var v40 int64 = -9223372036854775808
            _ = v40
        default:
            v22 = v22
        }
    } else {
        v3 += v3
        var v41 uint64 = uint64(v3)
        switch v23 {
        case "foo bar":
            var v42 int16 = 1
            _ = v42
        case "a":
            fmt.Println(v3)
        default:
            var v43 float32 = 0.0
            _ = v43
        }
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
