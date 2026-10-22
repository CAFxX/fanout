package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8, p2 int32) int8 {
    var v7 int = 2147483647
    var v8 uint16 = 0
    v7 += -2147483648
    p1 = (p0 / ((100) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 - (p1 & (p1 & -97)))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    var v3 uint16 = 505
    var v4 string = ""
    v5 := make(map[string]bool)
    var v6 bool = true
    if (-1 < -59) {
        v4 = v4
    } else {
        v6 = v6
        v4 = ("a" + v4)
    }
    delete(v2, 536)
    v9 := GenS1{F0: (v3 << ((2) & 63)), F1: -32768, F2: 842}
    fmt.Println(v6)
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v3)
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    var v12 bool = v6
    switch 218 {
    case 9223372036854775807:
        v13 := []int{32768, 32767, 32767}
        _ = v13
    }
    fmt.Println(v4, v12)
    switch v4 {
    case "a":
        var v14 int8 = 42
        _ = v14
    }
    if (!v12) {
        v15 := []int{-9223372036854775808, 32768}
        var v16 int32 = 0
        _ = v15
        _ = v16
    }
    v3 += (v3 * (v3 + (v3 + 255)))
    if (0 > 50) {
        for fi17 := 0; fi17 < 13; fi17++ {
            if (((fi17 | fi17) > fi17) && v12) {
                v18 := GenS1{F0: 969, F1: -32768, F2: (v3 | v3)}
                _ = v18
            } else {
                switch v4 {
                case "a":
                    var v19 int32 = -68
                    _ = v19
                case "hello":
                    var v20 uint32 = 0
                    _ = v20
                case "\n\t":
                    switch 0 {
                    case -708:
                        var v21 int16 = 0
                        _ = v21
                    default:
                        var v22 int32 = -2147483648
                        _ = v22
                    }
                default:
                    v6 = v12
                }
            }
            if true {
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v6)
                    _ = fi23
                }
                v4 = ((v4 + v4) + v4)
                var v24 uint32 = 1
                _ = v24
            } else {
                v3 = (v3 << ((42) & 63))
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := []int{}
                    var v27 uint32 = 2
                    switch fi25 {
                    case 4294967295:
                        v12 = true
                    case -1:
                        if (127 <= 42) {
                            v3 += (^(v3))
                            v28 := GenS1{F0: 255, F1: int16(v27), F2: v3}
                            _ = v28
                        } else {
                            var v29 int16 = 32767
                            if false {
                                var v30 uint64 = 18446744073709551615
                                var v31 int16 = int16(v30)
                                var v32 int16 = (^(v31))
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            _ = v29
                        }
                    }
                    v6 = (true && (1 >= 1))
                    var v33 float32 = 0.0
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v33
                }
                if true {
                    v34 := GenS1{F0: v3, F1: 0, F2: (v3 ^ v3)}
                    _ = v34
                }
            }
            v4 = v4
            _ = fi17
        }
        v6 = ((!(!(1.5 > 0.0))) && v6)
    } else {
        v12 = v12
        for fi35 := 0; fi35 < 8; fi35++ {
            for fi36 := 0; fi36 < 13; fi36++ {
                switch 9223372036854775807 {
                case 4294967295:
                    fmt.Println(fi36)
                case -604:
                    v2[fi35] = v4
                default:
                    if v6 {
                        v12 = v12
                    }
                }
                fmt.Println(fi36, fi36, v3)
                v12 = v12
                _ = fi36
            }
            _ = fi35
        }
    }
    var v37 int = 9223372036854775807
    v38 := GenS2{F0: 251}
    v12 = (!(true && true))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
