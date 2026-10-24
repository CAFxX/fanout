package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int8 = -96
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 uint32 = 713
    var v5 int32 = -688
    switch 2147483647 {
    case 42:
        var v6 uint16 = 65535
        _ = v6
    case 636:
        v7 := GenS1{F0: 0, F1: int8(v4), F2: "\n\t"}
        _ = v7
    default:
        v4 = (v4 - v4)
    }
    if ((^(v1)) < (v1 ^ -62)) {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    v10 := GenS1{F0: 18446744073709551615, F1: v1, F2: "hello"}
                    fmt.Println(fi9, v4, v5)
                    _ = v10
                }
                _ = fi9
            }
            _ = fi8
        }
    } else {
        if false {
            v11 := GenS1{F0: 228, F1: -1, F2: "foo bar"}
            _ = v11
        } else {
            switch ((((("a" + ("\n\t" + "\n\t")) + "") + ("hello" + "hello")) + "\n\t") + "") {
            case "a":
                v5 += (v5 | v5)
            case "\n\t":
                v3[v4] = 127
            case "hello":
                v12 := v3[v4]
                _ = v12
            }
        }
    }
    var v13 float64 = float64(float32(1e10))
    if (v1 > (v1 * v1)) {
        switch "" {
        case "hello":
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 string = ("hello" + ("foo bar" + "hello"))
                    if (!(fi15 <= fi14)) {
                        v1 = v1
                        v17 := GenS2{F0: 255, F1: (v1 << ((882) & 63))}
                        _ = v17
                    } else {
                        delete(v3, 375)
                    }
                    v13 -= 1.5
                    delete(v2, v4)
                    var v18 int16 = -32768
                    _ = fi15
                    _ = v16
                    _ = v18
                }
                switch fi14 {
                case 2147483647:
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 int8 = -1
                        _ = len(v3)
                        _ = fi19
                        _ = v20
                    }
                default:
                    v21 := []int{fi14, (fi14 * (fi14 / ((-1) | 1)))}
                    _ = v21
                }
                if (2 > v4) {
                    v4 = 2147483647
                    v22 := GenS1{F0: 4294967295, F1: (v1 >> ((340) & 63)), F2: "\n\t"}
                    v13 = 0.0
                    _ = v22
                }
                _ = fi14
            }
        case "a":
            switch 4294967295 {
            case 127:
                var v23 int = 482
                _ = v23
            case 133:
                v5 = (v5 >> ((1) & 63))
            case 4294967295:
                if true {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := GenS2{F0: 2, F1: v1}
                        v4 -= 2
                        _ = fi24
                        _ = v25
                    }
                }
            }
        }
        for fi26 := 0; fi26 < 14; fi26++ {
            switch "\n\t" {
            case "hello":
                v1 = v1
            default:
                fmt.Println(v4, v1, fi26)
            }
            _ = fi26
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 bool = (true && false)
            _ = fi27
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 16; fi29++ {
            if (318 < 42) {
                v30 := GenS1{F0: 586, F1: (v1 ^ (^(v1))), F2: ""}
                _ = v30
            } else {
                var v31 int32 = -459
                _ = v31
            }
            var v32 int16 = -830
            if false {
                for fi33 := 0; fi33 < 8; fi33++ {
                    var v34 uint = 0
                    if ((false && true) || (1 <= 209)) {
                        var v35 uint64 = 42
                        v36 := []int{fi33, fi33, int(v32), (fi33 ^ (fi29 % ((fi33) | 1)))}
                        _ = v35
                        _ = v36
                    } else {
                        var v37 uint64 = 1
                        _ = v37
                    }
                    _ = fi33
                    _ = v34
                }
                var v38 string = ("" + ("a" + "foo bar"))
                _ = v38
            } else {
                var v39 int64 = 9223372036854775807
                v40 := GenS3{F0: v32}
                v41 := GenS4{F0: -40}
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi29
            _ = v32
        }
    }
    fmt.Println(v13, v4, v1)
    var v42 uint = 42
    v43 := GenS5{F0: uint32(v42), F1: 32767}
    var v44 int64 = 0
    fmt.Println(v1, v1)
    if true {
        _ = len(v3)
    } else {
        switch 582 {
        case 393:
            if (32767 >= 1000) {
                var v45 int8 = 0
                v2[v4] = ("foo bar" + "a")
                fmt.Println(v1, v5, v5)
                _ = v45
            } else {
                fmt.Println(v13)
                fmt.Println(v42, v42)
                var v46 int = 1
                _ = v46
            }
        case 32768:
            fmt.Println(v1, v4)
        }
        v47 := GenS4{F0: v1}
        if ((4294967295 <= 42) && false) {
            fmt.Println(v13, v1, v13)
            v48 := v2[(v4 & v4)]
            var v49 bool = (!true)
            _ = v48
            _ = v49
        } else {
            v50 := GenS2{F0: 2, F1: 1}
            _ = v50
        }
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
