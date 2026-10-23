package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 65535
    var v3 int64 = -1
    var v4 float32 = float32(3.14159)
    var v5 int64 = -195
    v6 := []uint32{0}
    switch (("a" + "") + "\n\t") {
    case "a":
        var v7 string = "foo bar"
        _ = v7
    case "foo bar":
        v2 = v2
    default:
        v8 := GenS1{F0: 224, F1: v1}
        _ = v8
    }
    if false {
        if (("\n\t" >= "a") || false) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    } else {
        switch -9223372036854775808 {
        case 128:
            var v9 string = ""
            _ = v9
        }
        v3 = v5
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if (!false) {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint = 509
                var v13 uint64 = 4294967295
                v2 += v2
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: fi14}
            if true {
                v1 -= (fi10 + (fi14 & (^(fi10))))
                if ((v3 ^ (v3 % (((v3 & (v3 * v5))) | 1))) >= v5) {
                    fmt.Println(v5, v4, v4)
                    var v16 uint32 = 0
                    _ = v16
                } else {
                    v1 = (fi10 + (v1 % (((v1 * 0)) | 1)))
                }
                v3 = (v5 >> ((4294967295) & 63))
            }
            _ = fi14
            _ = v15
        }
        _ = fi10
    }
    var v17 uint16 = v2
    v5 += v5
    switch 32768 {
    case 9223372036854775807:
        v18 := GenS1{F0: 124, F1: (v1 + (v1 << ((18446744073709551615) & 63)))}
        _ = v18
    case -786:
        v19 := GenS3{F0: (v17 >> ((364) & 63)), F1: 443}
        _ = v19
    default:
        if false {
            if false {
                if false {
                    v6 = append(v6, 0)
                    if ((v17 >> ((255) & 63)) >= (v17 / ((v17) | 1))) {
                        var v20 float32 = v4
                        var v21 bool = (!false)
                        _ = v20
                        _ = v21
                    }
                } else {
                    if true {
                        if false {
                            var v22 int32 = 0
                            if (9223372036854775807 < 0) {
                                if false {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                var v23 uint8 = 255
                                _ = v23
                            } else {
                                v17 += v17
                                v24 := GenS4{F0: 42}
                                for fi25 := 0; fi25 < 20; fi25++ {
                                    var v26 uint8 = 8
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                            v27 := GenS1{F0: 127, F1: (v1 ^ 32767)}
                            _ = v22
                            _ = v27
                        }
                        fmt.Println(v5, v1)
                    }
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 float32 = 1.5
                        var v30 int32 = 1
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    if ((true && true) || (!(77 <= 2))) {
                        switch 128 {
                        case -813:
                            var v31 uint64 = 270
                            _ = v31
                        case 32767:
                            var v32 uint16 = v17
                            _ = v32
                        case -63:
                            var v33 float32 = 1.5
                            _ = v33
                        }
                        var v34 uint8 = 127
                        _ = v34
                    } else {
                        var v35 int32 = -634
                        _ = v35
                    }
                }
            }
            v2 += v2
        }
    }
    v1 += v1
    if (v5 != (v5 | v5)) {
        if len(v6) > 0 {
            v6[0] = 140
        }
        v6 = append(v6, 42)
    }
    switch v1 {
    case 127:
        fmt.Println(v5, v4)
    case 128:
        var v36 uint16 = 392
        _ = v36
    }
    if (true && false) {
        if len(v6) > 0 {
            v6[0] = 65535
        }
    } else {
        if (3.14159 < float64(((v4 + (v4 + (1e10 - ((v4 * 0.1) - v4)))) - 1e10))) {
            if (true && false) {
                var v37 int = v1
                var v38 int16 = 118
                for fi39 := 0; fi39 < 4; fi39++ {
                    v5 += -9223372036854775808
                    _ = len(v6)
                    _ = fi39
                }
                _ = v37
                _ = v38
            }
        } else {
            switch "\n\t" {
            case "hello":
                switch (v1 & v1) {
                case 595:
                    v2 -= v2
                case 128:
                    fmt.Println(v1, v1, v1)
                case 0:
                    v17 += v2
                }
            case "\n\t":
                for fi40 := 0; fi40 < 3; fi40++ {
                    if false {
                        var v41 float32 = (float32(1e10) * v4)
                        _ = v41
                    }
                    _ = fi40
                }
            default:
                var v42 string = "foo bar"
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 2; fi43++ {
            v17 += v17
            v44 := GenS4{F0: 18446744073709551615}
            _ = fi43
            _ = v44
        }
    }
    switch ("" + (("foo bar" + "\n\t") + "a")) {
    case "\n\t":
        v4 = 3.14159
    case "":
        var v45 string = "\n\t"
        _ = v45
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if ((v17 | v17) <= 348) {
        v4 += 0.1
        fmt.Println(v3, v17, v1)
        var v46 uint8 = 117
        _ = v46
    }
    fmt.Println(v5)
    if true {
        v4 = -2.25
    }
    var v47 int16 = 1000
    if false {
        if false {
            _ = len(v6)
        }
        v47 = (v47 | v47)
        switch 128 {
        case -368:
            fmt.Println(v3, v5)
        case -9223372036854775808:
            var v48 int16 = 607
            _ = v48
        case 997:
            if false {
                fmt.Println(v5, v4)
                var v49 float64 = ((-2.25 / 1e10) + (0.1 / (-2.25 - 1.5)))
                v50 := GenS3{F0: 255, F1: v47}
                _ = v49
                _ = v50
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
