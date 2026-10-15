package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 string }
type GenS4 struct { F0 uint16; F1 string; F2 int16 }
type GenS5 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    fmt.Println(p0)
    var v6 uint = 782
    var v7 uint64 = 18446744073709551615
    _ = p0
    _ = v6
    _ = v7
    return uint(v7)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]int)
    var v3 string = (("a" + (v1 + v1)) + "\n\t")
    v4 := make(map[int]string)
    var v5 uint = 0
    fmt.Println(v5)
    v8 := GenS1{F0: -9223372036854775808}
    v9 := []int{712, -922, 9223372036854775807, 128}
    v5 -= (v5 >> ((v5) & 63))
    v10 := GenS2{F0: v1, F1: 1, F2: 956}
    v3 = (("a" + v3) + ("foo bar" + v3))
    var v11 int = 786
    v12 := GenS3{F0: 255, F1: 886, F2: v1}
    fmt.Println(v3, v5, v3)
    v1 = v3
    if false {
        var v13 float32 = 1e10
        _ = v13
    } else {
        if (v1 > v3) {
            var v14 float32 = 0.0
            _ = v14
        } else {
            if ((3.14159 / 1e10) >= -2.25) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 2; fi16++ {
                        v17 := v4[32768]
                        _ = fi16
                        _ = v17
                    }
                    v18 := GenS4{F0: 707, F1: "foo bar", F2: 1000}
                    if (!(3.14159 == 0.0)) {
                        var v19 string = v3
                        for fi20 := 0; fi20 < 6; fi20++ {
                            fmt.Println(v3)
                            var v21 uint8 = 128
                            if len(v9) > 0 {
                                v9[0] = int(v21)
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                    _ = fi15
                    _ = v18
                }
                v5 += genFn0(34)
                var v22 int16 = 0
                _ = v22
            } else {
                fmt.Println(v11)
                v23 := GenS3{F0: 65535, F1: 577, F2: (("\n\t" + v3) + v3)}
                _ = v23
            }
            fmt.Println(v5, v1, v3)
        }
    }
    v1 = ((("foo bar" + ("a" + v3)) + (v3 + "hello")) + "a")
    for fi24 := 0; fi24 < 2; fi24++ {
        switch fi24 {
        case 1:
            if false {
                delete(v4, v11)
                for fi25 := 0; fi25 < 17; fi25++ {
                    v9 = append(v9, fi24)
                    var v26 int = (fi25 >> ((v5) & 63))
                    var v27 bool = (int32(v26) > -2147483648)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 uint8 = 42
                    v30 := GenS4{F0: 255, F1: v1, F2: 42}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if false {
                        var v31 int = -9223372036854775808
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                v32 := GenS2{F0: "hello", F1: -698, F2: int32(v5)}
                _ = v32
            }
        case 2147483647:
            v1 = (v1 + v1)
        default:
            if (-266 == (^(fi24))) {
                var v33 int = fi24
                for fi34 := 0; fi34 < 1; fi34++ {
                    if false {
                        var v35 uint = (v5 >> ((v5) & 63))
                        _ = v35
                    }
                    _ = fi34
                }
                switch v11 {
                case 7:
                    switch "hello" {
                    case "foo bar":
                        for fi36 := 0; fi36 < 4; fi36++ {
                            var v37 float32 = ((3.14159 - -2.25) * 3.14159)
                            v38 := GenS5{F0: 4294967295, F1: 42, F2: 849}
                            fmt.Println(fi24, v37)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    case "\n\t":
                        v5 = (v5 + (^(v5)))
                    case "a":
                        for fi39 := 0; fi39 < 19; fi39++ {
                            var v40 int64 = 9223372036854775807
                            v11 += -2147483648
                            _ = fi39
                            _ = v40
                        }
                    }
                case 9223372036854775807:
                    var v41 uint8 = 51
                    _ = v41
                default:
                    var v42 int16 = 32767
                    _ = v42
                }
                _ = v33
            } else {
                v43 := GenS1{F0: 2147483647}
                _ = v43
            }
        }
        if (1 >= -766) {
            fmt.Println(v11, v3)
        } else {
            if len(v9) > 0 {
                v9[0] = v11
            }
            v1 = "\n\t"
            var v44 uint = v5
            _ = v44
        }
        var v45 float64 = 1.5
        _ = fi24
        _ = v45
    }
    v11 = (v11 & (v11 >> ((v5) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
