package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
func genFn0(p0 uint8, p1 int32) int8 {
    if false {
        var v7 float32 = 1e10
        v8 := GenS1{F0: 2, F1: 1}
        var v9 float64 = (1e10 + (1e10 + (1.5 / 3.14159)))
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = v10
    return 44
}
func genFn1(p0 int32, p1 int32, p2 int) int64 {
    var v11 float32 = float32((0.0 - float64(1e10)))
    var v12 int32 = 0
    var v13 float64 = 0.0
    var v14 bool = (!(true || false))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 252
}

func main() {
    var v1 string = ((("\n\t" + "\n\t") + "foo bar") + "foo bar")
    var v2 int64 = 292
    var v3 uint = 1
    var v4 int64 = v2
    var v5 int8 = int8(v3)
    var v6 int8 = -4
    v6 += (v5 << ((v3) & 63))
    v2 += v4
    v15 := []int{0, 822, 32768, 127}
    var v16 int32 = 0
    var v17 uint = (v3 << ((v3) & 63))
    if (42 >= 4294967295) {
        var v18 bool = false
        _ = v18
    }
    if false {
        _ = len(v15)
        for fi19 := 0; fi19 < 11; fi19++ {
            if true {
                if true {
                    var v20 int16 = -32768
                    if (((^(fi19)) >= fi19) && true) {
                        var v21 int64 = (v2 * int64(fi19))
                        v22 := GenS1{F0: 1, F1: v20}
                        var v23 float32 = (1.5 / ((-2.25 + -2.25) * -2.25))
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int32 = v16
                        switch fi19 {
                        case 1:
                            if len(v15) > 0 {
                                v15[0] = fi19
                            }
                        case 9223372036854775807:
                            v25 := GenS1{F0: 4294967295, F1: (v20 * v20)}
                            _ = v25
                        case 233:
                            var v26 bool = ((0 != uint32(v17)) && (!true))
                            _ = v26
                        default:
                            v27 := GenS2{F0: 2, F1: ("hello" + v1), F2: v16}
                            _ = v27
                        }
                        _ = v24
                    }
                    _ = v20
                }
                switch "a" {
                case "hello":
                    if len(v15) > 0 {
                        v15[0] = (^(fi19))
                    }
                case "\n\t":
                    var v28 uint16 = 266
                    _ = v28
                case "a":
                    var v29 int64 = -9223372036854775808
                    _ = v29
                }
                fmt.Println(v6)
            }
            var v30 string = ""
            var v31 uint64 = 42
            _ = fi19
            _ = v30
            _ = v31
        }
    } else {
        v15 = append(v15, 0)
    }
    v32 := GenS3{F0: (-(v6)), F1: 42, F2: v16}
    v2 -= -9223372036854775808
    if false {
        _ = len(v15)
        v4 = (v4 | 1)
        if (1e10 != (1.5 / 1e10)) {
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v15) > 0 {
                    v15[0] = int(v2)
                }
                _ = fi33
            }
        }
    }
    v1 = ((v1 + "") + v1)
    var v34 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
