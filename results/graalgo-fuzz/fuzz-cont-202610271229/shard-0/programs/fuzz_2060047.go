package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0(p0 bool) int16 {
    var v6 uint32 = 0
    fmt.Println(p0, v6)
    var v7 int16 = 0
    _ = p0
    _ = v6
    _ = v7
    return (v7 ^ 32767)
}

func main() {
    v1 := []uint32{930, 198, 42, 4294967295}
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    if ((1.5 <= 1.5) || false) {
        var v8 float32 = 3.14159
        if false {
            if (false || (42 >= -588)) {
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(fi9)
                    delete(v2, (fi9 - (^(fi9))))
                    _ = fi9
                }
                var v10 float32 = (((-2.25 * ((v8 / 3.14159) + 1e10)) + 1.5) / ((-2.25 + v8) + 0.1))
                _ = v10
            }
            v8 -= 0.0
        } else {
            fmt.Println(v8)
        }
        var v11 int = 127
        _ = v11
        _ = v8
    } else {
        var v12 int = 0
        if (3.14159 >= 3.14159) {
            for fi13 := 0; fi13 < 11; fi13++ {
                _ = len(v4)
                _ = fi13
            }
            if true {
                fmt.Println(v12)
            } else {
                _ = len(v1)
                switch (v12 - (v12 - 4294967295)) {
                case -51:
                    if ((false && true) && (-128 != -128)) {
                        switch ("" + ("hello" + ("" + ""))) {
                        case "foo bar":
                            var v14 int64 = 1
                            _ = v14
                        case "a":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v12)
                            case "hello":
                                var v15 uint64 = 42
                                _ = v15
                            }
                        case "\n\t":
                            var v16 uint32 = 2147483648
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 11; fi17++ {
                                if (100 != -93) {
                                    v18 := GenS1{F0: -59, F1: 4294967295, F2: false}
                                    _ = len(v1)
                                    for fi19 := 0; fi19 < 1; fi19++ {
                                        switch (v12 << ((2) & 63)) {
                                        case 9223372036854775807:
                                            var v20 int16 = 914
                                            _ = v20
                                        default:
                                            v21 := v3[1]
                                            _ = v21
                                        }
                                        _ = fi19
                                    }
                                    _ = v18
                                } else {
                                    var v22 uint = 255
                                    _ = v22
                                }
                                fmt.Println(fi17)
                                _ = fi17
                            }
                        }
                        v23 := GenS1{F0: 100, F1: 431, F2: true}
                        _ = v23
                    } else {
                        var v24 uint32 = 884
                        v25 := GenS2{F0: 255, F1: 817, F2: -1}
                        _ = v24
                        _ = v25
                    }
                case 42:
                    v12 -= (v12 + v12)
                case -9223372036854775808:
                    v26 := GenS3{F0: 255, F1: (((("foo bar" + "foo bar") + "") + ("\n\t" + "foo bar")) + "\n\t"), F2: 1}
                    _ = v26
                }
            }
            fmt.Println(v12)
        }
        if ((true && (!(true || (!(2147483647 != 2147483647))))) && false) {
            var v27 uint = 1
            v27 = (v27 & 2)
            var v28 int64 = -9223372036854775808
            _ = v27
            _ = v28
        }
        _ = v12
    }
    if true {
        v29 := GenS3{F0: 42, F1: "hello", F2: 65535}
        var v30 uint32 = 42
        fmt.Println(v30)
        _ = v29
        _ = v30
    } else {
        var v31 int = 42
        _ = v31
    }
    v1 = append(v1, 42)
    switch 686 {
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = 2
        }
    default:
        v1 = append(v1, 826)
    }
    var v32 int8 = 100
    v33 := GenS1{F0: (v32 ^ (-(v32))), F1: 4294967295, F2: (!(v32 >= v32))}
    v32 = (v32 | (-(v32)))
    v32 -= v32
    v4[2] = (false && (!false))
    v32 = v32
    var v34 int64 = 286
    var v35 int16 = -305
    switch "" {
    case "":
        if false {
            var v36 uint32 = 42
            var v37 int64 = v34
            _ = v36
            _ = v37
        }
    case "hello":
        v35 -= (v35 ^ 32767)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
