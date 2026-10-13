package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int; F2 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint64) uint64 {
    v7 := []int{-2147483648, 1, -2147483648, 0}
    p0 -= (p0 << ((255) & 63))
    var v8 float32 = (-2.25 - 0.1)
    var v9 float64 = 1.5
    fmt.Println(v8, v9)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := []int{-182, 9223372036854775807}
    v2 := []string{"\n\t", (("a" + ("foo bar" + "")) + "hello"), "", ((("hello" + "a") + (("a" + "foo bar") + ("a" + "\n\t"))) + (("hello" + "") + "a"))}
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 uint16 = 355
    var v6 uint32 = 2
    switch "\n\t" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 32768
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            switch fi10 {
            case -1:
                v5 = 255
            case 4294967295:
                if true {
                    var v11 int64 = 9223372036854775807
                    _ = v11
                }
            case 128:
                v5 = (v5 & v5)
            default:
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := GenS1{F0: -1, F1: 633}
                    v14 := GenS2{F0: 9223372036854775807, F1: -429}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        }
        _ = fi10
    }
    if (0.1 >= float64(0.1)) {
        var v15 uint = 4294967295
        var v16 uint8 = 2
        for fi17 := 0; fi17 < 8; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                switch (fi17 >> (((v15 << (((^(v15))) & 63))) & 63)) {
                case -2147483648:
                    if (!(-86 >= 42)) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if ((3.14159 <= 3.14159) || (-2.25 <= (1e10 * 0.0))) {
                            v19 := GenS2{F0: 1, F1: -775}
                            var v20 uint32 = (v6 % (((v6 % (((v6 ^ (v6 | v6))) | 1))) | 1))
                            if (261 != -846) {
                                v21 := GenS3{F0: 1}
                                v22 := GenS4{F0: true, F1: -9223372036854775808, F2: 42}
                                var v23 uint64 = 930
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            _ = v19
                            _ = v20
                        } else {
                            var v24 uint64 = 42
                            fmt.Println(fi17)
                            _ = v24
                        }
                        var v25 float32 = 0.1
                        _ = v25
                    }
                default:
                    v26 := GenS5{F0: (v6 % ((v6) | 1))}
                    _ = v26
                }
                var v27 bool = true
                v15 += (v15 / ((v15) | 1))
                fmt.Println(fi17, v15, v5)
                _ = fi18
                _ = v27
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                fmt.Println(v6, v6)
                v30 := GenS4{F0: true, F1: fi29, F2: 9223372036854775807}
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    }
    _ = len(v4)
    v31 := GenS6{F0: -1, F1: 127}
    var v32 uint16 = 388
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS4{F0: (-573 < (fi33 / (((^(fi33))) | 1))), F1: (fi33 / ((fi33) | 1)), F2: 42}
        _ = fi33
        _ = v34
    }
    fmt.Println(v5, v5)
    if true {
        _ = len(v2)
        v32 = v5
        v32 = (v32 % ((v32) | 1))
    }
    v5 = 582
    var v35 int64 = 1
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 16; fi37++ {
            v38 := GenS3{F0: 32767}
            switch (fi37 + 862) {
            case 2147483647:
                v6 += v6
            default:
                v39 := GenS1{F0: v35, F1: 416}
                _ = v39
            }
            switch 0 {
            case 713:
                _ = len(v3)
            case 42:
                if (!(!true)) {
                    fmt.Println(v35, v6)
                }
            case 2147483647:
                if (1e10 >= (float64(0.1) / (-2.25 - ((-2.25 - 1.5) + -2.25)))) {
                    var v40 int32 = 2147483647
                    _ = v40
                } else {
                    v41 := v3["a"]
                    v5 = v32
                    var v42 int32 = 1
                    _ = v41
                    _ = v42
                }
            default:
                var v43 int8 = -1
                _ = v43
            }
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    fmt.Println(v6, v35, v6)
    v44 := GenS3{F0: 1000}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v45 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
