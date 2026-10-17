package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int, p1 int32) uint {
    fmt.Println(p0)
    fmt.Println(p0, p0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 767
}
func genFn1(p0 int64) uint {
    v8 := []int{0, 9223372036854775807, 127}
    p0 -= (p0 / (((p0 / ((p0) | 1))) | 1))
    switch "hello" {
    case "\n\t":
        for fi9 := 0; fi9 < 15; fi9++ {
            if true {
                _ = len(v8)
                var v10 uint16 = 121
                v8 = append(v8, fi9)
                _ = v10
            } else {
                switch 42 {
                case 4294967295:
                    for fi11 := 0; fi11 < 19; fi11++ {
                        fmt.Println(p0, fi11, fi9)
                        switch 127 {
                        case 32767:
                            _ = len(v8)
                        case 1:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        _ = fi11
                    }
                case -1:
                    v8 = append(v8, int(p0))
                case -735:
                    switch (fi9 & fi9) {
                    case 4294967295:
                        var v12 float64 = 0.0
                        _ = v12
                    case -718:
                        if (true || (2 != 1)) {
                            v13 := GenS1{F0: 137, F1: "\n\t"}
                            _ = v13
                        } else {
                            if len(v8) > 0 {
                                v8[0] = (-(fi9))
                            }
                        }
                    default:
                        fmt.Println(p0, fi9)
                    }
                }
            }
            p0 -= (p0 ^ p0)
            var v14 int = 32768
            _ = fi9
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS1{F0: 164, F1: ("foo bar" + (("hello" + "foo bar") + (("hello" + "a") + "hello")))}
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint8 = 2
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 255
    v3 := make(map[int]string)
    var v4 bool = (false && true)
    var v5 float64 = float64(float32(((3.14159 * 1.5) * 0.1)))
    var v6 int16 = 32767
    var v19 uint32 = 65535
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: 42, F1: (("a" + "a") + "foo bar")}
        _ = fi20
        _ = v21
    }
    if ((836 <= 343) && (100 >= -128)) {
        v22 := GenS1{F0: 0, F1: ""}
        if (0 != 42) {
            v23 := []int{4294967295}
            var v24 int32 = 1
            _ = v23
            _ = v24
        }
        _ = v22
    }
    v4 = false
    switch -961 {
    case 32768:
        for fi25 := 0; fi25 < 1; fi25++ {
            switch "" {
            case "a":
                v1 = v1
            default:
                v26 := GenS1{F0: 0, F1: "\n\t"}
                _ = v26
            }
            v4 = (-9223372036854775808 <= -9223372036854775808)
            _ = fi25
        }
    case 4294967295:
        v27 := GenS1{F0: uint8(v2), F1: ("\n\t" + "a")}
        _ = v27
    default:
        v2 += v2
    }
    delete(v3, 128)
    fmt.Println(v1, v19, v6)
    var v28 int8 = -1
    v29 := GenS2{F0: 0, F1: 42}
    v1 += (((((1e10 * v1) - ((float32((-2.25 + (v5 / 1.5))) * (float32(float64((1.5 - 0.1))) * (v1 - 1.5))) - (float32(v5) * v1))) * 1e10) * 1.5) / ((v1 + ((1e10 - v1) - 0.0)) - (v1 + v1)))
    var v30 int = -601
    switch (("a" + "a") + "foo bar") {
    case "hello":
        v31 := v3[(v30 >> ((genFn0(32768, 42)) & 63))]
        _ = v31
    case "a":
        switch int(v19) {
        case 9223372036854775807:
            fmt.Println(v1)
        }
    default:
        delete(v3, 307)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
