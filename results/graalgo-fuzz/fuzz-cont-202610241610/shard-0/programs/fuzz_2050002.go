package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 int, p1 int) uint8 {
    if (2147483647 > 42) {
        p0 = 32767
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int8 = 100
            _ = fi7
            _ = v8
        }
        switch "\n\t" {
        case "":
            switch "hello" {
            case "hello":
                if true {
                    var v9 int16 = -486
                    for fi10 := 0; fi10 < 2; fi10++ {
                        for fi11 := 0; fi11 < 19; fi11++ {
                            v12 := []int{fi11, fi11, fi11, (^(fi11))}
                            _ = fi11
                            _ = v12
                        }
                        v13 := []int{(p1 >> ((899) & 63))}
                        _ = fi10
                        _ = v13
                    }
                    _ = v9
                } else {
                    var v14 uint = 2
                    fmt.Println(p0, v14)
                    for fi15 := 0; fi15 < 11; fi15++ {
                        if (1 != 4294967295) {
                            switch 367 {
                            case 4294967295:
                                var v16 uint32 = 1
                                _ = v16
                            case 9223372036854775807:
                                for fi17 := 0; fi17 < 17; fi17++ {
                                    p1 = fi17
                                    v18 := GenS1{F0: 460, F1: 140}
                                    v14 -= 0
                                    _ = fi17
                                    _ = v18
                                }
                            case 42:
                                fmt.Println(fi15, p0, v14)
                            }
                            if (127 == 255) {
                                v14 = v14
                                var v19 int32 = 42
                                var v20 int = (p0 | -1)
                                _ = v19
                                _ = v20
                            }
                            var v21 int8 = 82
                            _ = v21
                        } else {
                            v22 := []int{p1}
                            _ = v22
                        }
                        p0 += (p0 - (p1 ^ (p1 - (^(p1)))))
                        _ = fi15
                    }
                    _ = v14
                }
            case "foo bar":
                p0 -= (-(p1))
            }
        case "a":
            var v23 float32 = ((float32(((float64(1e10) / 0.1) * 0.1)) + (-2.25 + 1.5)) - (0.1 / 1e10))
            _ = v23
        case "\n\t":
            for fi24 := 0; fi24 < 10; fi24++ {
                p0 = (fi24 % (((^(p1))) | 1))
                v25 := GenS1{F0: 909, F1: 1}
                _ = fi24
                _ = v25
            }
        }
    } else {
        p1 -= (p1 / (((p1 / (((p0 % ((p1) | 1))) | 1))) | 1))
        p1 = p0
    }
    var v26 uint8 = 255
    if true {
        var v27 uint16 = 270
        v28 := GenS1{F0: 678, F1: 1}
        p0 = (-(p0))
        _ = v27
        _ = v28
    }
    if (600 >= 1) {
        v29 := GenS1{F0: -362, F1: 1}
        _ = v29
    }
    var v30 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v26
    _ = v30
    return (v26 | v26)
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    var v4 int16 = 32767
    v5 := make(map[int]string)
    var v6 uint16 = 979
    v6 -= (v6 >> ((1) & 63))
    v1 = (v1 + ((v1 + v1) + v1))
    delete(v5, 960)
    var v31 bool = ((127 >= 127) || true)
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(v4)
        _ = fi32
    }
    v33 := GenS2{F0: -2147483648, F1: 255}
    var v34 bool = (!v31)
    var v35 int32 = 2147483647
    switch ("" + (v1 + v1)) {
    case "a":
        if (1 != 90) {
            v34 = true
            var v36 int32 = v35
            var v37 uint64 = 18446744073709551615
            _ = v36
            _ = v37
        } else {
            if v34 {
                v6 = 980
            } else {
                v38 := v5[0]
                var v39 int32 = (v35 + v35)
                _ = v38
                _ = v39
            }
        }
    default:
        var v40 int8 = 100
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
