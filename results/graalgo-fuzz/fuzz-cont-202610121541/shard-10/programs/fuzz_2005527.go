package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0(p0 string) uint64 {
    p0 = "hello"
    v6 := []int{-657, -9223372036854775808}
    p0 = ("\n\t" + (p0 + "\n\t"))
    var v7 uint = 2
    _ = p0
    _ = v6
    _ = v7
    return uint64(v7)
}
func genFn1(p0 uint8, p1 uint16, p2 int8) uint {
    fmt.Println(p2, p1)
    v8 := GenS1{F0: 1, F1: (p2 / ((-128) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    var v1 int64 = 0
    var v2 float32 = ((float32(((3.14159 + -2.25) - 0.0)) - 1e10) - float32(-2.25))
    v3 := []string{"hello", ((("a" + ("hello" + "hello")) + (("" + ("\n\t" + "hello")) + "hello")) + "hello"), "foo bar"}
    v4 := make(map[uint32]int)
    var v5 uint32 = 1
    v2 = float32(1.5)
    switch 42 {
    case 127:
        switch "hello" {
        case "\n\t":
            v9 := v4[(v5 | 1)]
            _ = v9
        default:
            if true {
                v10 := v4[(v5 % ((v5) | 1))]
                _ = v10
            } else {
                var v11 int32 = 1
                _ = v11
            }
        }
    case 1:
        if true {
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := GenS1{F0: 127, F1: 80}
                _ = fi12
                _ = v13
            }
            v14 := v4[(v5 >> ((838) & 63))]
            _ = v14
        } else {
            for fi15 := 0; fi15 < 10; fi15++ {
                v1 = v1
                var v16 uint16 = 2
                switch 630 {
                case 128:
                    v17 := GenS1{F0: fi15, F1: -71}
                    _ = v17
                default:
                    var v18 int32 = -1
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
    }
    var v19 uint = 65535
    if true {
        if ((v5 * 42) <= (v5 | 4294967295)) {
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    if false {
                        switch fi21 {
                        case -57:
                            var v22 float32 = v2
                            _ = v22
                        }
                    }
                    _ = fi21
                }
                v5 = (v5 | v5)
                var v23 int64 = -566
                _ = fi20
                _ = v23
            }
            v19 = (v19 / ((v19) | 1))
            v1 = 590
        } else {
            if false {
                var v24 uint32 = v5
                _ = v24
            }
        }
    }
    if false {
        v1 += v1
        v25 := GenS1{F0: 128, F1: 127}
        var v26 uint = 65535
        _ = v25
        _ = v26
    }
    if (false && true) {
        if false {
            if (255 >= 1) {
                switch -9223372036854775808 {
                case 103:
                    if (float32(float64(1.5)) <= float32(3.14159)) {
                        fmt.Println(v19, v2)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    }
                case -9223372036854775808:
                    v1 = 427
                }
                for fi27 := 0; fi27 < 4; fi27++ {
                    v1 -= 1
                    var v28 int32 = 42
                    _ = fi27
                    _ = v28
                }
                var v29 uint64 = 914
                _ = v29
            } else {
                if true {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v31 := GenS2{F0: "\n\t"}
                        _ = fi30
                        _ = v31
                    }
                    v32 := GenS2{F0: ("foo bar" + "\n\t")}
                    _ = v32
                } else {
                    if (genFn1(33, 0, 0) == 42) {
                        fmt.Println(v5, v5)
                    }
                    v4[(v5 >> ((genFn1(0, 0, 127)) & 63))] = -9223372036854775808
                    v3 = append(v3, "foo bar")
                }
                _ = len(v3)
            }
            var v33 float64 = 0.1
            _ = v33
        } else {
            v5 = v5
            var v34 int8 = 9
            v35 := GenS3{F0: (v5 & (^(v5)))}
            _ = v34
            _ = v35
        }
        fmt.Println(v19)
    }
    v5 -= 2
    v1 = (-(v1))
    v36 := GenS3{F0: v5}
    for fi37 := 0; fi37 < 3; fi37++ {
        v1 += (v1 << (((v19 / (((v19 * (^(v19)))) | 1))) & 63))
        _ = fi37
    }
    if false {
        v5 = 42
    } else {
        var v38 uint = 42
        _ = v38
    }
    v3 = append(v3, "foo bar")
    fmt.Println(v5)
    v39 := GenS4{F0: v5, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
}
