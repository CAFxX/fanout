package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }
func genFn0(p0 bool) string {
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{127, -218, (fi7 | (fi8 << ((4294967295) & 63)))}
            _ = fi8
            _ = v9
        }
        p0 = true
        for fi10 := 0; fi10 < 4; fi10++ {
            p0 = (2 == 912)
            v11 := GenS1{F0: -2147483648, F1: (("hello" + "hello") + "\n\t"), F2: 0}
            var v12 uint8 = 251
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    p0 = (127 == -128)
    v13 := []int{127, -9223372036854775808, 4294967295}
    _ = len(v13)
    var v14 uint8 = 0
    _ = p0
    _ = v13
    _ = v14
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 2
    var v3 float64 = 0.0
    var v4 int64 = -747
    var v5 uint16 = 169
    v6 := make(map[uint32]int)
    v2 -= (v2 ^ v2)
    switch "hello" {
    case "hello":
        v4 -= v4
    default:
        delete(v6, 2147483647)
    }
    v15 := []int{953}
    var v16 string = "\n\t"
    var v17 float64 = 0.1
    v16 = "foo bar"
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int64 = (v4 / ((-1) | 1))
        v3 += (0.1 * float64(1e10))
        for fi20 := 0; fi20 < 3; fi20++ {
            if (32767 > 0) {
                _ = len(v6)
            } else {
                v2 = (v2 & 397)
                v3 = v3
            }
            v19 -= v4
            if (1 != 2) {
                var v21 uint16 = 855
                switch (fi20 + (fi18 | fi18)) {
                case 20:
                    v2 = uint64(v21)
                case 2147483647:
                    v22 := GenS2{F0: v16, F1: 0, F2: -123}
                    _ = v22
                default:
                    v5 += 0
                }
                _ = v21
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            switch v16 {
            case "a":
                v17 = v17
            case "hello":
                switch 2147483647 {
                case 42:
                    fmt.Println(fi24)
                case 4294967295:
                    var v25 int32 = 666
                    _ = v25
                }
            default:
                switch 2147483647 {
                case -1:
                    if false {
                        if (42 == 32767) {
                            var v26 uint = 483
                            v3 = v17
                            v27 := GenS2{F0: v16, F1: 2, F2: 42}
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS1{F0: 2147483647, F1: v16, F2: 0}
                            v15 = append(v15, (fi23 + (^(fi23))))
                            var v29 uint = 18446744073709551615
                            _ = v28
                            _ = v29
                        }
                        v3 = v3
                        v4 -= v4
                    }
                default:
                    if false {
                        switch ((v16 + "foo bar") + "") {
                        case "":
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        v4 = 9223372036854775807
                        if true {
                            v30 := v1[(v16 + v16)]
                            _ = v30
                        }
                    } else {
                        if false {
                            var v31 int64 = (^(v4))
                            if false {
                                if (1e10 != 1.5) {
                                    var v32 uint32 = 693
                                    _ = v32
                                }
                                _ = len(v15)
                                _ = len(v15)
                            } else {
                                if (18446744073709551615 < 255) {
                                    v2 = uint64(v5)
                                    var v33 int8 = 42
                                    _ = v33
                                }
                            }
                            switch (v16 + "") {
                            case "foo bar":
                                var v34 float64 = v17
                                _ = v34
                            case "":
                                v16 = (("a" + "foo bar") + ("foo bar" + ((v16 + v16) + v16)))
                                var v35 int8 = int8(v5)
                                _ = v35
                            case "hello":
                                v36 := GenS1{F0: 42, F1: v16, F2: v2}
                                _ = v36
                            }
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v37 int64 = (v31 - v4)
                            _ = v31
                            _ = v37
                        }
                    }
                }
            }
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println(v17)
    var v38 uint32 = 969
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
