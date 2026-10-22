package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 uint64, p1 uint16) int64 {
    var v5 uint64 = (p0 / (((p0 * p0)) | 1))
    p1 -= 0
    _ = p0
    _ = p1
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 bool) uint64 {
    switch 42 {
    case 504:
        fmt.Println(p0)
    }
    v6 := GenS1{F0: -12}
    var v7 float64 = 1e10
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := []int{32767, 1}
    if (!(2147483647 != -1)) {
        if true {
            var v8 uint = 851
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(fi9, fi9)
                _ = fi9
            }
            _ = v8
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 int64 = 161
                var v12 uint32 = 120
                fmt.Println(v11, fi10, v11)
                _ = fi10
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 228, F1: 2, F2: genFn1((103 != 127))}
            _ = v13
        }
        switch 932 {
        case 2147483647:
            switch "\n\t" {
            case "\n\t":
                var v14 int16 = -1
                _ = v14
            case "foo bar":
                for fi15 := 0; fi15 < 5; fi15++ {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        var v17 uint16 = 0
                        _ = fi16
                        _ = v17
                    }
                    _ = fi15
                }
            default:
                for fi18 := 0; fi18 < 16; fi18++ {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v20 int64 = 0
                        switch "foo bar" {
                        case "hello":
                            delete(v1, 736)
                        case "\n\t":
                            var v21 int64 = 1
                            _ = v21
                        case "foo bar":
                            var v22 int16 = 42
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                    v23 := GenS2{F0: 2147483647, F1: 42, F2: 727}
                    v4 = append(v4, fi18)
                    _ = fi18
                    _ = v23
                }
            }
        case 1:
            var v24 bool = (true || false)
            _ = v24
        case -2147483648:
            var v25 string = ("foo bar" + "a")
            _ = v25
        default:
            if len(v4) > 0 {
                v4[0] = 1
            }
        }
        v26 := v3[-863]
        _ = v26
    } else {
        var v27 string = "foo bar"
        v27 = ("\n\t" + "foo bar")
        _ = v27
    }
    if (!false) {
        var v28 uint64 = 0
        v3[4294967295] = (2147483647 > 65535)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(fi29)
            var v30 uint = 18446744073709551615
            _ = fi29
            _ = v30
        }
        v31 := GenS2{F0: 578, F1: 0, F2: 828}
        _ = v31
    }
    v32 := GenS2{F0: 65535, F1: 255, F2: genFn1((!((!false) || (-1 < -1))))}
    var v33 uint32 = 444
    v34 := v2[2147483647]
    v34 -= v34
    delete(v2, (v33 << ((42) & 63)))
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
