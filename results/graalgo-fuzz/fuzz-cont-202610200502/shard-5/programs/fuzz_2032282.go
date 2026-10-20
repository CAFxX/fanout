package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int32) uint32 {
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 360
    var v3 float32 = 1.5
    v4 := []int{-285, 4294967295, 42}
    var v5 int32 = -35
    var v6 uint16 = 2
    v6 += v2
    v4 = append(v4, -1)
    for fi7 := 0; fi7 < 4; fi7++ {
        v1[(("" + ("\n\t" + ("a" + "hello"))) + "hello")] = 2147483648
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v6, v6, v5)
            v9 := GenS1{F0: 42, F1: 1, F2: v2}
            fmt.Println(v3)
            var v10 uint = 4294967295
            v6 = (v6 & uint16(v10))
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: -1, F1: 922, F2: v6}
    if false {
        if false {
            var v12 int32 = (v5 >> ((302) & 63))
            _ = v12
        } else {
            var v13 uint64 = 106
            var v14 int64 = 12
            v3 -= 0.0
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            if false {
                var v16 uint16 = (^(v6))
                _ = v16
            } else {
                var v17 uint8 = 255
                if false {
                    switch "foo bar" {
                    case "hello":
                        delete(v1, "foo bar")
                    case "foo bar":
                        var v18 uint8 = 0
                        _ = v18
                    case "\n\t":
                        var v19 int16 = -898
                        _ = v19
                    }
                    var v20 string = "\n\t"
                    var v21 uint16 = (v2 & 2)
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS1{F0: 9223372036854775807, F1: 703, F2: (v6 ^ 42)}
                    var v24 uint = 1
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v17
            }
            v25 := GenS1{F0: -864, F1: 18446744073709551615, F2: v2}
            _ = fi15
            _ = v25
        }
        v2 = v2
    } else {
        _ = len(v1)
        delete(v1, "foo bar")
    }
    if (!(42 < 42)) {
        var v26 int8 = 127
        _ = v26
    } else {
        fmt.Println(v5, v5, v3)
        var v27 int32 = 0
        _ = v27
    }
    var v28 uint64 = 6
    switch 39 {
    case 2147483647:
        v5 = 0
    }
    switch -27 {
    case -373:
        v2 += (v2 & v6)
    case -1:
        if (false && false) {
            var v29 int16 = -32768
            _ = v29
        } else {
            var v30 uint16 = 32767
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v31 int8 = 127
            _ = v30
            _ = v31
        }
    default:
        var v32 uint8 = 53
        _ = v32
    }
    _ = len(v4)
    var v33 float32 = ((v3 - 1e10) * v3)
    switch -537 {
    case 32768:
        v34 := GenS2{F0: 500, F1: (v2 % (((v2 << ((402) & 63))) | 1))}
        _ = v34
    case 445:
        v6 -= 0
    default:
        v4 = append(v4, int(v5))
    }
    switch -2147483648 {
    case 378:
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 uint32 = 561
            fmt.Println(v3, v33, v28)
            _ = fi35
            _ = v36
        }
    case -204:
        fmt.Println(v28)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
