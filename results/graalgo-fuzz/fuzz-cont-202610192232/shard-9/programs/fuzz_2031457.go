package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint }
type GenS2 struct { F0 int32; F1 int; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
func genFn0(p0 string) int32 {
    var v6 int16 = -32768
    if false {
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v7 uint8 = 127
            _ = v7
        case 9223372036854775807:
            if false {
                v8 := GenS1{F0: 402, F1: (v6 * (-(v6))), F2: 947}
                if false {
                    for fi9 := 0; fi9 < 11; fi9++ {
                        if (590 != 65535) {
                            fmt.Println(v6, v6, fi9)
                            if false {
                                var v10 uint8 = 127
                                var v11 float32 = 0.0
                                _ = v10
                                _ = v11
                            } else {
                                switch (p0 + "foo bar") {
                                case "":
                                    v12 := []int{fi9, fi9}
                                    _ = v12
                                case "\n\t":
                                    v6 = 47
                                case "a":
                                    var v13 int16 = v6
                                    _ = v13
                                }
                            }
                        }
                        _ = fi9
                    }
                    v14 := GenS1{F0: 42, F1: (v6 / ((v6) | 1)), F2: 146}
                    p0 = (("" + ((p0 + ((("foo bar" + "\n\t") + (p0 + "a")) + "\n\t")) + (("foo bar" + ("a" + "")) + p0))) + p0)
                    _ = v14
                } else {
                    v15 := GenS1{F0: 748, F1: 42, F2: 585}
                    if (42 >= 2) {
                        var v16 uint = 923
                        switch "hello" {
                        case "a":
                            v6 = -1
                        default:
                            for fi17 := 0; fi17 < 12; fi17++ {
                                switch p0 {
                                case "":
                                    v18 := GenS1{F0: 42, F1: (v6 << ((v16) & 63)), F2: 303}
                                    _ = v18
                                case "a":
                                    v19 := []int{(^(fi17)), fi17, fi17}
                                    _ = v19
                                case "hello":
                                    v6 = -565
                                }
                                v20 := []int{fi17, fi17, (fi17 << ((v16) & 63)), (fi17 - (fi17 * 127))}
                                _ = fi17
                                _ = v20
                            }
                        }
                        v6 = v6
                        _ = v16
                    }
                    v21 := []int{-1, -1, 128}
                    _ = v15
                    _ = v21
                }
                p0 = ("\n\t" + "\n\t")
                _ = v8
            }
        }
        if (-1 == -1) {
            v22 := GenS1{F0: 100, F1: v6, F2: 1}
            _ = v22
        }
    }
    var v23 float64 = (1e10 - 1e10)
    _ = p0
    _ = v23
    _ = v6
    return 2147483647
}

func main() {
    var v1 int16 = 63
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 uint8 = 92
    var v5 uint64 = 42
    for fi24 := 0; fi24 < 13; fi24++ {
        if true {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 float64 = 0.0
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                v4 += 1
                fmt.Println(v1, v5, v1)
                _ = fi27
            }
        } else {
            v28 := GenS2{F0: 878, F1: (fi24 ^ (fi24 - (^(fi24)))), F2: int(v5)}
            _ = v28
        }
        v4 -= v4
        _ = fi24
    }
    if false {
        v1 -= 1
    } else {
        v29 := GenS3{F0: 9223372036854775807}
        _ = v29
    }
    v1 = -836
    v5 = 1
    v1 = v1
    v30 := GenS1{F0: 32767, F1: v1, F2: 201}
    var v31 uint16 = 1
    v32 := GenS4{F0: v4, F1: v4, F2: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
