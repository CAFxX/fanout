package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint8, p1 uint64) int32 {
    v7 := []int{}
    if false {
        for fi8 := 0; fi8 < 4; fi8++ {
            p0 = (p0 + (^(p0)))
            fmt.Println(p0, p0)
            _ = fi8
        }
        v9 := GenS1{F0: 0}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch ("hello" + "") {
            case "foo bar":
                switch fi10 {
                case -921:
                    if (1 == 2) {
                        var v11 uint64 = p1
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 uint = 18446744073709551615
                            _ = len(v7)
                            v11 -= (p1 << ((931) & 63))
                            var v14 int8 = 45
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = v11
                    }
                default:
                    v15 := GenS2{F0: 42, F1: -59}
                    _ = v15
                }
            case "a":
                switch fi10 {
                case 668:
                    _ = len(v7)
                case 32767:
                    var v16 uint = 251
                    _ = v16
                }
            case "hello":
                p0 -= 0
            default:
                fmt.Println(p1, p0, p0)
            }
            var v17 int64 = -477
            _ = fi10
            _ = v17
        }
        fmt.Println(p0, p1)
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 uint = 255
    v2 := []string{"\n\t", "\n\t"}
    var v3 int32 = -2147483648
    var v4 uint32 = 2
    var v5 float32 = (((float32(0.1) / 1.0) + -2.25) + (-2.25 / 3.14159))
    var v6 string = ("" + "a")
    var v18 int16 = 50
    v4 += v4
    var v19 float64 = 0.0
    v20 := GenS2{F0: genFn0(0, uint64(v1)), F1: -83}
    var v21 uint8 = 42
    v3 += (^(v3))
    var v22 int64 = 908
    v3 = v3
    v18 -= v18
    fmt.Println(v22, v5, v1)
    if false {
        _ = len(v2)
        fmt.Println(v21, v3)
    } else {
        v23 := GenS1{F0: v22}
        switch 485 {
        case -121:
            for fi24 := 0; fi24 < 8; fi24++ {
                v21 += 127
                v3 = int32(fi24)
                for fi25 := 0; fi25 < 5; fi25++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v21 = uint8(v4)
                    _ = fi25
                }
                _ = fi24
            }
        case 9223372036854775807:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v5 = 1e10
        }
        _ = v23
    }
    fmt.Println(v22)
    v26 := GenS1{F0: (v22 << (((v1 + (v1 << (((v1 >> ((18446744073709551615) & 63))) & 63)))) & 63))}
    if true {
        v27 := GenS3{F0: 701}
        v4 += v4
        _ = v27
    }
    v22 += (v22 * v22)
    v21 = 255
    if false {
        fmt.Println(v5, v18)
        for fi28 := 0; fi28 < 2; fi28++ {
            if (!((^(v4)) > 65535)) {
                var v29 uint16 = 893
                var v30 bool = false
                _ = v29
                _ = v30
            } else {
                var v31 float64 = v19
                fmt.Println(v4)
                if true {
                    v32 := GenS2{F0: (v3 * (v3 ^ v3)), F1: -1}
                    _ = v32
                }
                _ = v31
            }
            _ = fi28
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v33 := GenS2{F0: v3, F1: 100}
        _ = v33
    }
    v34 := GenS4{F0: 32767, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
