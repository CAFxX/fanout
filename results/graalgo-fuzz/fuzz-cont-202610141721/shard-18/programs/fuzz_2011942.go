package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int32; F1 int64 }

func main() {
    var v1 int = 726
    var v2 uint16 = 65535
    v3 := make(map[string]bool)
    var v4 int16 = 1000
    v5 := []uint32{1, 2147483647, 756}
    var v6 string = "hello"
    v7 := GenS1{F0: 1, F1: (v1 & v1)}
    v2 = v2
    for fi8 := 0; fi8 < 15; fi8++ {
        if (0.0 != 1e10) {
            switch (v1 - (v1 - (fi8 << ((301) & 63)))) {
            case 42:
                var v9 int8 = 42
                _ = v9
            case 1:
                fmt.Println(fi8, v4, fi8)
            case -9223372036854775808:
                switch (("\n\t" + v6) + v6) {
                case "a":
                    v1 -= 4294967295
                case "\n\t":
                    v2 = 42
                default:
                    if ((float64((3.14159 - float32((1e10 * 1.5)))) - ((1.5 + 0.0) / float64(1.5))) > 0.1) {
                        v1 = 961
                        if true {
                            var v10 uint32 = 42
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 6; fi11++ {
                                fmt.Println(v4, fi8)
                                fmt.Println(fi11, v4)
                                var v12 uint = 2
                                _ = fi11
                                _ = v12
                            }
                        }
                    }
                }
            default:
                v2 = (^(v2))
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                fmt.Println(fi13, v4, v1)
                _ = len(v5)
                _ = fi13
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := v3[((((v6 + v6) + ((v6 + "\n\t") + "a")) + ("\n\t" + (((v6 + v6) + ("hello" + v6)) + v6))) + (("hello" + (v6 + (v6 + v6))) + v6))]
                _ = fi14
                _ = v15
            }
        } else {
            var v16 string = "\n\t"
            switch (((v6 + (v6 + v16)) + (v16 + ("hello" + v6))) + v6) {
            case "foo bar":
                if true {
                    if len(v5) > 0 {
                        v5[0] = 4294967295
                    }
                    switch "foo bar" {
                    case "hello":
                        v17 := GenS2{F0: v2, F1: 795, F2: (-(v4))}
                        _ = v17
                    }
                } else {
                    var v18 float32 = (0.1 / 1e10)
                    v1 += (^(v1))
                    _ = v18
                }
            default:
                v19 := v3[v6]
                _ = v19
            }
            v16 = (v16 + ("a" + ((v16 + v16) + "foo bar")))
            _ = v16
        }
        v20 := GenS3{F0: 42, F1: 127}
        _ = fi8
        _ = v20
    }
    v21 := GenS1{F0: 0, F1: v1}
    v22 := GenS4{F0: v2, F1: v4}
    if ((!false) || (false && false)) {
        fmt.Println(v6, v1)
        v23 := GenS5{F0: -2147483648, F1: 42}
        _ = v23
    } else {
        v24 := GenS5{F0: 42, F1: 339}
        _ = v24
    }
    var v25 float64 = 1.5
    var v26 uint = 4294967295
    for fi27 := 0; fi27 < 18; fi27++ {
        v1 = fi27
        var v28 uint32 = 42
        var v29 int64 = 809
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := v3[v6]
    v31 := GenS3{F0: 0, F1: 32767}
    v25 = float64(-2.25)
    for fi32 := 0; fi32 < 10; fi32++ {
        v33 := GenS5{F0: 2147483647, F1: -972}
        fmt.Println(v2)
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
