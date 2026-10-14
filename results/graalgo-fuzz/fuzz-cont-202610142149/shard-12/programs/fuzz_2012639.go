package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 int32) uint8 {
    var v5 uint8 = 0
    v5 -= 128
    var v6 bool = true
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int64, p1 uint64) uint16 {
    p1 -= (p1 % (((p1 % ((p1) | 1))) | 1))
    if (563 == 42) {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "foo bar":
            p1 = 38
        }
    }
    var v7 int8 = 127
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint = 442
    var v2 int16 = -1
    v3 := make(map[uint32]string)
    var v4 float32 = 1e10
    v2 += (^(v2))
    v4 = float32(1e10)
    var v8 uint32 = 65535
    v1 = v1
    v3[966] = "foo bar"
    for fi9 := 0; fi9 < 5; fi9++ {
        if (!true) {
            var v10 int8 = 0
            v11 := GenS1{F0: 128, F1: 1}
            _ = v10
            _ = v11
        } else {
            v2 = -587
            v12 := GenS1{F0: genFn0(-9223372036854775808, 487), F1: 32767}
            _ = v12
        }
        _ = fi9
    }
    v4 -= v4
    fmt.Println(v2, v1, v4)
    v2 = v2
    if ((!("\n\t" >= (("foo bar" + "\n\t") + "a"))) && (42 <= 4294967295)) {
        var v13 int = -9223372036854775808
        _ = v13
    }
    var v14 float32 = v4
    if (1 <= 1) {
        if true {
            switch "foo bar" {
            case "":
                v2 += v2
            case "\n\t":
                fmt.Println(v4)
            }
            var v15 uint8 = 0
            _ = v15
        } else {
            var v16 uint64 = 18446744073709551615
            v17 := GenS1{F0: 255, F1: 231}
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS2{F0: 65535}
                _ = fi19
                _ = v20
            }
            switch "a" {
            case "a":
                for fi21 := 0; fi21 < 11; fi21++ {
                    switch fi18 {
                    case 32767:
                        if false {
                            v22 := GenS1{F0: 122, F1: 1}
                            v4 -= (v14 * v14)
                            var v23 uint64 = 42
                            _ = v22
                            _ = v23
                        }
                    case -9223372036854775808:
                        switch int(v1) {
                        case -2147483648:
                            var v24 uint32 = 1
                            _ = v24
                        case 32768:
                            v25 := GenS1{F0: 0, F1: 901}
                            _ = v25
                        case 128:
                            if (v14 <= (v14 * v14)) {
                                switch fi18 {
                                case 128:
                                    if true {
                                        var v26 float32 = v4
                                        _ = v26
                                    }
                                case 9223372036854775807:
                                    switch fi21 {
                                    case 32767:
                                        var v27 int8 = -56
                                        _ = v27
                                    case 128:
                                        var v28 uint64 = 282
                                        _ = v28
                                    case 0:
                                        v14 += v4
                                    }
                                    var v29 uint16 = 548
                                    var v30 float32 = float32(1e10)
                                    _ = v29
                                    _ = v30
                                default:
                                    _ = len(v3)
                                }
                                if (!true) {
                                    var v31 int16 = v2
                                    v32 := v3[1]
                                    fmt.Println(fi18)
                                    var v33 int8 = -128
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                v8 -= v8
                            }
                        }
                    case 9223372036854775807:
                        v14 += 1e10
                    }
                    _ = fi21
                }
            case "":
                var v34 string = ("\n\t" + ("\n\t" + ""))
                _ = v34
            case "foo bar":
                if ((fi18 / ((fi18) | 1)) < fi18) {
                    v2 -= (v2 % ((v2) | 1))
                    var v35 uint64 = 393
                    var v36 float32 = -2.25
                    _ = v35
                    _ = v36
                }
            default:
                fmt.Println(v1, v2, v2)
            }
            if true {
                _ = len(v3)
                v14 += 0.1
                var v37 string = "\n\t"
                _ = v37
            } else {
                v3[v8] = ""
                fmt.Println(v2, v4)
                v2 = (v2 ^ (v2 + (v2 >> ((v1) & 63))))
            }
            _ = fi18
        }
    }
    var v38 int32 = -869
    v1 = (v1 % ((42) | 1))
    var v39 int = 9223372036854775807
    for fi40 := 0; fi40 < 6; fi40++ {
        fmt.Println(v39, v8, v4)
        _ = fi40
    }
    v38 = (v38 / ((v38) | 1))
    v14 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v8
}
