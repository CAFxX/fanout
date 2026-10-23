package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int64; F1 uint8; F2 int64 }

func main() {
    var v1 string = ""
    var v2 uint64 = 377
    var v3 int16 = 1
    v4 := []uint32{}
    var v5 uint = 845
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        _ = len(v6)
        var v8 string = ("hello" + "\n\t")
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v5)
    if (!false) {
        var v9 int16 = v3
        _ = v9
    } else {
        v5 = v5
    }
    v1 = v1
    v2 += 2
    var v10 int32 = 695
    var v11 float64 = 0.1
    if true {
        switch 2147483647 {
        case -912:
            switch (v1 + "hello") {
            case "a":
                fmt.Println(v11)
            case "\n\t":
                v11 -= ((v11 - v11) / float64(0.1))
            default:
                _ = len(v4)
            }
        case 42:
            var v12 float64 = (((((v11 - v11) + v11) + v11) / 1.0) / (v11 * -2.25))
            _ = v12
        default:
            v5 = (v5 ^ (v5 | (v5 / (((v5 & v5)) | 1))))
        }
        var v13 int16 = (v3 + -1)
        _ = v13
    } else {
        v2 = 1
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v11 = 0.0
        _ = fi14
    }
    switch (v1 + "") {
    case "":
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                v3 = v3
                switch (fi16 / ((fi15) | 1)) {
                case -9223372036854775808:
                    v11 = (1.5 - 1e10)
                case -2147483648:
                    v10 = (v10 / (((v10 % ((v10) | 1))) | 1))
                default:
                    if len(v4) > 0 {
                        v4[0] = 235
                    }
                }
                _ = fi16
            }
            v17 := GenS1{F0: 762, F1: 105, F2: v2}
            v18 := GenS1{F0: 112, F1: v2, F2: 4294967295}
            _ = fi15
            _ = v17
            _ = v18
        }
    default:
        switch ("\n\t" + "foo bar") {
        case "hello":
            switch (v1 + "") {
            case "a":
                v5 += v5
            case "\n\t":
                v10 = -106
            default:
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS1{F0: 385, F1: 1, F2: 0}
                    v3 -= (v3 << ((347) & 63))
                    var v21 uint = (v5 - 305)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        case "foo bar":
            switch 0 {
            case 32768:
                fmt.Println(v10, v10)
            case 316:
                switch -931 {
                case 4294967295:
                    if true {
                        v6[2147483647] = 0
                        v5 = 1
                        var v22 uint32 = uint32(v5)
                        _ = v22
                    } else {
                        var v23 uint = (^(v5))
                        v24 := GenS2{F0: 100, F1: 42, F2: (v10 | (-(v10)))}
                        v23 = v5
                        _ = v23
                        _ = v24
                    }
                case 0:
                    if true {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v5 = v5
                            _ = fi25
                        }
                        v10 -= v10
                    }
                default:
                    var v26 string = (v1 + v1)
                    _ = v26
                }
            case 1:
                v27 := GenS1{F0: 0, F1: v2, F2: 80}
                _ = v27
            }
        default:
            v28 := GenS3{F0: v5, F1: 98}
            _ = v28
        }
    }
    fmt.Println(v5)
    v3 = v3
    v10 -= v10
    if (338 < (v5 << ((v5) & 63))) {
        v29 := GenS1{F0: 65535, F1: (v2 << ((v5) & 63)), F2: 4294967295}
        _ = v29
    } else {
        delete(v6, 4294967295)
        v10 = v10
        v10 -= v10
    }
    var v30 int64 = 1
    if (2 < 65535) {
        v3 += v3
        var v31 uint = v5
        _ = v31
    } else {
        for fi32 := 0; fi32 < 2; fi32++ {
            fmt.Println(v3, v10)
            var v33 float32 = 0.0
            var v34 float32 = v33
            _ = fi32
            _ = v33
            _ = v34
        }
        if false {
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS4{F0: v10, F1: v3}
                for fi37 := 0; fi37 < 11; fi37++ {
                    v1 = "foo bar"
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                    var v38 float32 = (((((1.5 - ((float32(v11) / (1e10 * 1.5)) * -2.25)) * 0.0) + (1.5 / 1.5)) + (1.5 - 0.1)) + 1.5)
                    _ = fi37
                    _ = v38
                }
                if true {
                    v39 := GenS5{F0: (v30 + (v30 - int64(v10))), F1: 127, F2: -9223372036854775808}
                    v11 = -2.25
                    if (v11 <= v11) {
                        _ = len(v4)
                        var v40 int16 = int16(v10)
                        _ = v40
                    }
                    _ = v39
                }
                _ = fi35
                _ = v36
            }
            v3 -= -32768
        }
    }
    if true {
        var v41 uint16 = 830
        switch -933 {
        case -1:
            var v42 uint8 = 178
            _ = v42
        case 127:
            v43 := GenS1{F0: 4294967295, F1: (v2 / ((v2) | 1)), F2: (v2 - v2)}
            _ = v43
        }
        v10 -= 1
        _ = v41
    } else {
        switch ("foo bar" + v1) {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "":
            fmt.Println(v10)
        case "hello":
            var v44 uint = (v5 >> (((v5 >> ((v5) & 63))) & 63))
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
