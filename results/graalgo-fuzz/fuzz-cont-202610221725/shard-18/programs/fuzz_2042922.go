package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{718, 1, 32768}
    v3 := make(map[int]uint32)
    var v4 int8 = 1
    v4 += (v4 & v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: -9223372036854775808, F1: -333, F2: 911}
        _ = fi5
        _ = v6
    }
    if (!false) {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 42, F1: 374, F2: 255}
            _ = fi7
            _ = v8
        }
        v4 = v4
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v4 -= v4
        v4 = (v4 % (((v4 << ((2) & 63))) | 1))
    }
    v2 = append(v2, 128)
    _ = len(v2)
    if true {
        fmt.Println(v4)
        var v9 uint = 529
        _ = v9
    } else {
        switch 0 {
        case -335:
            switch -534 {
            case -670:
                for fi10 := 0; fi10 < 19; fi10++ {
                    v4 = 127
                    switch (^(fi10)) {
                    case 2147483647:
                        if false {
                            fmt.Println(v4, fi10)
                            v4 = (^(v4))
                            _ = len(v2)
                        } else {
                            var v11 uint32 = 2
                            switch ("foo bar" + ("a" + ("a" + "\n\t"))) {
                            case "a":
                                delete(v1, (v11 + v11))
                            default:
                                _ = len(v2)
                            }
                            var v12 int32 = -2147483648
                            _ = v11
                            _ = v12
                        }
                    }
                    if (255 != 181) {
                        v4 = v4
                        v13 := GenS1{F0: 9223372036854775807, F1: 1, F2: 1}
                        _ = v13
                    }
                    _ = fi10
                }
            case -64:
                switch "foo bar" {
                case "":
                    var v14 uint = 4294967295
                    _ = v14
                }
            case 1:
                var v15 float64 = (((1e10 - 1e10) / (1e10 - (float64(0.0) / 1e10))) + 0.1)
                _ = v15
            default:
                v16 := v3[-563]
                _ = v16
            }
        case 0:
            for fi17 := 0; fi17 < 17; fi17++ {
                fmt.Println(v4)
                _ = fi17
            }
        default:
            for fi18 := 0; fi18 < 20; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v19 := GenS1{F0: 0, F1: -1, F2: 2}
                _ = fi18
                _ = v19
            }
        }
    }
    v4 += (v4 % ((v4) | 1))
    for fi20 := 0; fi20 < 4; fi20++ {
        if ("a" <= "") {
            var v21 uint32 = 513
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v21
        }
        var v22 uint = 989
        delete(v3, fi20)
        _ = fi20
        _ = v22
    }
    v4 -= (v4 - (v4 % ((-128) | 1)))
    for fi23 := 0; fi23 < 6; fi23++ {
        if (true && false) {
            v24 := v1[65535]
            var v25 int64 = -9223372036854775808
            if true {
                v24 -= 2
                v4 += v4
            }
            _ = v24
            _ = v25
        }
        v4 = v4
        _ = fi23
    }
    switch -2147483648 {
    case -52:
        v26 := GenS1{F0: 68, F1: 0, F2: 42}
        _ = v26
    case 32768:
        var v27 int8 = -128
        _ = v27
    case 4294967295:
        v4 = v4
    }
    if true {
        v4 -= v4
        v28 := GenS1{F0: 1, F1: -1, F2: 32767}
        for fi29 := 0; fi29 < 18; fi29++ {
            fmt.Println(v4, fi29)
            _ = fi29
        }
        _ = v28
    }
    v4 += (v4 + (-(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
