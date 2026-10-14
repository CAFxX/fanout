package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }

func main() {
    v1 := []string{("\n\t" + "")}
    var v2 uint64 = 4294967295
    var v3 int64 = -111
    v4 := []int{-2147483648}
    v5 := []uint32{1}
    v3 += -9223372036854775808
    if len(v5) > 0 {
        v5[0] = 921
    }
    v3 = (v3 ^ v3)
    v2 -= v2
    for fi6 := 0; fi6 < 9; fi6++ {
        if (!true) {
            for fi7 := 0; fi7 < 1; fi7++ {
                v3 += int64(fi6)
                _ = fi7
            }
            if true {
                v3 += (-(v3))
                if ((v3 ^ -1) != v3) {
                    if (2 < 1) {
                        v3 += v3
                    }
                    switch fi6 {
                    case 42:
                        var v8 uint32 = 4294967295
                        _ = v8
                    default:
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                    }
                }
            } else {
                switch 127 {
                case -191:
                    var v9 uint32 = 0
                    _ = v9
                case 4294967295:
                    var v10 uint64 = (v2 + (v2 ^ 62))
                    _ = v10
                default:
                    var v11 int = fi6
                    _ = v11
                }
                var v12 uint = 42
                _ = v12
            }
            v13 := GenS1{F0: (v2 % (((^(v2))) | 1))}
            _ = v13
        } else {
            v14 := GenS2{F0: 0, F1: -128}
            if false {
                if true {
                    v3 = 42
                } else {
                    var v15 int16 = 724
                    _ = v15
                }
                var v16 uint = 0
                _ = v16
            } else {
                _ = len(v1)
                v4 = append(v4, -412)
                var v17 bool = (!(fi6 <= (fi6 / ((127) | 1))))
                _ = v17
            }
            var v18 bool = false
            _ = v14
            _ = v18
        }
        v19 := GenS3{F0: "foo bar", F1: 556, F2: v2}
        switch "\n\t" {
        case "":
            v3 = (v3 / ((v3) | 1))
        case "foo bar":
            var v20 int8 = -128
            _ = v20
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi6
        _ = v19
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        if true {
            v22 := GenS1{F0: (^(v2))}
            v23 := GenS1{F0: v2}
            _ = v22
            _ = v23
        } else {
            v24 := GenS3{F0: ("hello" + ("foo bar" + "foo bar")), F1: -32768, F2: (v2 * (v2 * v2))}
            var v25 float64 = (-2.25 - -2.25)
            _ = v24
            _ = v25
        }
        var v26 uint8 = 255
        var v27 bool = (true && true)
        _ = fi21
        _ = v26
        _ = v27
    }
    switch "\n\t" {
    case "a":
        v28 := GenS2{F0: 65535, F1: 42}
        _ = v28
    case "\n\t":
        var v29 uint8 = 127
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        v2 += 230
        _ = fi30
    }
    v2 = 968
    v31 := GenS1{F0: (v2 % ((18446744073709551615) | 1))}
    if (!(true && true)) {
        switch "hello" {
        case "hello":
            v32 := GenS2{F0: 42, F1: 127}
            _ = v32
        case "foo bar":
            if true {
                var v33 int = -565
                _ = v33
            } else {
                v34 := GenS1{F0: v2}
                v2 += 0
                _ = v34
            }
        }
        var v35 int64 = v3
        for fi36 := 0; fi36 < 7; fi36++ {
            _ = len(v1)
            v35 = (v3 + v35)
            var v37 int8 = 17
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    switch 127 {
    case -9223372036854775808:
        if len(v4) > 0 {
            v4[0] = -9223372036854775808
        }
    case 9223372036854775807:
        if (0 > -107) {
            v3 = v3
        }
    case 141:
        for fi38 := 0; fi38 < 14; fi38++ {
            var v39 float64 = 3.14159
            var v40 int16 = -444
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
