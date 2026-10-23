package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 int) uint8 {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 1, F1: -104}
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 int32, p1 int8, p2 int8) int32 {
    v8 := []int{0, 32768, 128, -843}
    v9 := GenS2{F0: 0}
    if (-2.25 != -2.25) {
        var v10 uint8 = 2
        _ = v10
    } else {
        p0 += 2147483647
        p2 = 27
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        _ = len(v8)
        if (false && false) {
            for fi12 := 0; fi12 < 1; fi12++ {
                p1 = (p2 * -48)
                _ = fi12
            }
            switch 960 {
            case -2147483648:
                v8 = append(v8, (fi11 | fi11))
            case -9223372036854775808:
                _ = len(v8)
            default:
                if (("hello" < "a") && true) {
                    var v13 float32 = float32(3.14159)
                    _ = v13
                } else {
                    var v14 int64 = -104
                    _ = v14
                }
            }
            var v15 float32 = -2.25
            _ = v15
        }
        if len(v8) > 0 {
            v8[0] = fi11
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p0 & p0)
}

func main() {
    var v1 string = ""
    var v2 float64 = 0.1
    v3 := []int{-315, -582, 1, 314}
    var v4 int64 = 0
    var v5 uint64 = 1
    var v16 uint = 233
    for fi17 := 0; fi17 < 16; fi17++ {
        v5 += 4294967295
        var v18 string = (((v1 + (("a" + "\n\t") + ("foo bar" + ("foo bar" + v1)))) + v1) + (("foo bar" + v1) + v1))
        _ = fi17
        _ = v18
    }
    v4 = (-(v4))
    v1 = "hello"
    v16 = v16
    var v19 int64 = (v4 ^ v4)
    fmt.Println(v19, v5)
    v1 = "\n\t"
    var v20 uint64 = (^(v5))
    var v21 int = 32767
    var v22 bool = (127 < -1)
    fmt.Println(v20, v19, v16)
    var v23 int8 = int8(v21)
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int8 = (v23 + (v23 & 60))
        if (v22 && (v22 && v22)) {
            switch "hello" {
            case "":
                if (57 == 42) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v26 := GenS2{F0: (v4 >> (((^(v16))) & 63))}
                    v19 -= v19
                    _ = v26
                }
            case "foo bar":
                v23 = (^(v25))
            case "a":
                v3 = append(v3, (v21 & (^(v21))))
            }
        } else {
            for fi27 := 0; fi27 < 3; fi27++ {
                fmt.Println(v4, v2, v20)
                if v22 {
                    if ((v23 - -128) < (-(v25))) {
                        var v28 int64 = v19
                        v3 = append(v3, (fi27 * (v21 << (((v16 - 530)) & 63))))
                        _ = v28
                    }
                    switch (fi27 / ((-2147483648) | 1)) {
                    case 0:
                        var v29 uint32 = uint32(v5)
                        _ = v29
                    case -2147483648:
                        v25 -= (-(v25))
                    default:
                        v30 := GenS1{F0: v19, F1: v23}
                        _ = v30
                    }
                } else {
                    switch (v21 << (((v16 / (((^(v16))) | 1))) & 63)) {
                    case 42:
                        v20 += (v5 << ((v16) & 63))
                    default:
                        fmt.Println(v20, v2, v16)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v31 uint64 = 2
                        _ = v31
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi27
            }
            switch 1 {
            case 1:
                v32 := GenS1{F0: v4, F1: v23}
                _ = v32
            case 0:
                if false {
                    for fi33 := 0; fi33 < 20; fi33++ {
                        if true {
                            var v34 int16 = 32767
                            _ = v34
                        } else {
                            var v35 uint8 = 128
                            var v36 string = ""
                            var v37 uint16 = 439
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        v16 = (v16 & uint(v5))
                        var v38 int64 = (^(v19))
                        _ = fi33
                        _ = v38
                    }
                }
            }
        }
        v39 := GenS3{F0: 127}
        _ = fi24
        _ = v25
        _ = v39
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        fmt.Println(v4, fi40)
        _ = fi40
    }
    v41 := GenS2{F0: -9223372036854775808}
    var v42 int64 = (v19 & (v4 >> ((v16) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
