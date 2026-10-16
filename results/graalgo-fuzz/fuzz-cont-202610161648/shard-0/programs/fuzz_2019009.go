package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
func genFn0(p0 uint32, p1 int) uint8 {
    fmt.Println(p1)
    p0 = p0
    switch "foo bar" {
    case "foo bar":
        p1 -= -1
    case "a":
        v7 := []int{2147483647, 504}
        _ = v7
    case "hello":
        v8 := GenS1{F0: "hello", F1: 0}
        _ = v8
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 string, p1 uint16, p2 int8) int64 {
    switch p0 {
    case "":
        v9 := GenS2{F0: 968, F1: 155}
        _ = v9
    case "hello":
        if (true || (42 <= 127)) {
            if (1 == 966) {
                v10 := []int{-78, 32767, 32767, 921}
                switch 986 {
                case -2147483648:
                    for fi11 := 0; fi11 < 19; fi11++ {
                        v12 := GenS3{F0: -1, F1: (1000 < -825), F2: -728}
                        switch -9223372036854775808 {
                        case 4294967295:
                            _ = len(v10)
                        }
                        fmt.Println(fi11, p2)
                        _ = fi11
                        _ = v12
                    }
                case 9223372036854775807:
                    p2 += -128
                case -9223372036854775808:
                    var v13 int = 1
                    _ = v13
                }
                for fi14 := 0; fi14 < 10; fi14++ {
                    p0 = "hello"
                    v15 := GenS2{F0: 318, F1: -2147483648}
                    _ = fi14
                    _ = v15
                }
                _ = v10
            }
            if true {
                if false {
                    v16 := []int{}
                    p1 = p1
                    for fi17 := 0; fi17 < 7; fi17++ {
                        if len(v16) > 0 {
                            v16[0] = fi17
                        }
                        var v18 float64 = 1.5
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                switch (p0 + (p0 + "\n\t")) {
                case "\n\t":
                    if (0 > 0) {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 float32 = ((0.0 - 1.5) - 0.0)
                            _ = fi19
                            _ = v20
                        }
                        p2 = (-(p2))
                    } else {
                        var v21 float64 = (0.1 + ((0.0 - (1e10 / 1.0)) - 0.0))
                        p0 = "a"
                        p2 = (p2 & 127)
                        _ = v21
                    }
                case "a":
                    p2 = (p2 << ((808) & 63))
                default:
                    v22 := []int{}
                    _ = v22
                }
                v23 := GenS1{F0: p0, F1: 155}
                _ = v23
            }
            var v24 bool = false
            _ = v24
        }
    default:
        var v25 string = (((p0 + (p0 + "hello")) + (p0 + "hello")) + p0)
        _ = v25
    }
    var v26 bool = false
    p0 = p0
    switch "" {
    case "hello":
        p0 = "foo bar"
    case "a":
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 int32 = 857
    var v3 uint = 255
    v4 := make(map[string]uint32)
    v5 := []uint32{2147483648, 2, 2147483648, 65535}
    var v6 int16 = int16(v3)
    delete(v4, "")
    v5 = append(v5, 42)
    v1 = (v6 - v1)
    v3 += v3
    for fi27 := 0; fi27 < 1; fi27++ {
        v2 = -2147483648
        _ = fi27
    }
    var v28 string = ""
    v2 = (v2 % ((304) | 1))
    var v29 uint = v3
    var v30 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
