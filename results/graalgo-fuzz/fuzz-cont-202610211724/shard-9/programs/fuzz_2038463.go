package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 int16 = 1
    var v2 uint32 = 380
    v3 := []uint32{(v2 + (v2 & v2))}
    v1 = (v1 * (-(v1)))
    if (!((v1 << ((65535) & 63)) >= (v1 << ((384) & 63)))) {
        switch -2147483648 {
        case 2147483647:
            var v4 int32 = 42
            _ = v4
        }
    }
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 float64 = (0.1 - float64(3.14159))
        if false {
            var v7 int8 = 0
            for fi8 := 0; fi8 < 6; fi8++ {
                if false {
                    v9 := GenS1{F0: 9223372036854775807, F1: 42, F2: (1 <= 144)}
                    v10 := GenS2{F0: 578}
                    _ = v10
                    _ = v9
                }
                var v11 int64 = -503
                _ = fi8
                _ = v11
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                if len(v3) > 0 {
                    v3[0] = (^(v2))
                }
                _ = fi12
            }
            _ = v7
        }
        switch "hello" {
        case "hello":
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 float64 = float64(0.1)
                v6 += v6
                _ = fi13
                _ = v14
            }
        case "foo bar":
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS2{F0: 1}
                var v17 uint32 = v2
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = fi5
        _ = v6
    }
    if true {
        v1 = (v1 ^ 1000)
    }
    if false {
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(v2, fi18)
            for fi19 := 0; fi19 < 11; fi19++ {
                if (129 != 128) {
                    var v20 uint64 = 157
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v20
                } else {
                    _ = len(v3)
                    switch (^(fi19)) {
                    case 128:
                        v1 = v1
                    case -105:
                        v1 = v1
                    case 9223372036854775807:
                        v21 := GenS3{F0: 100}
                        _ = v21
                    default:
                        if ((!((!(122 < 100)) || (true && true))) && (0 != 42)) {
                            fmt.Println(fi18, v1)
                        }
                    }
                    v1 = (v1 & (-(v1)))
                }
                fmt.Println(v2, v1)
                var v22 uint = 138
                fmt.Println(fi19)
                _ = fi19
                _ = v22
            }
            _ = fi18
        }
        v3 = append(v3, (^(v2)))
        for fi23 := 0; fi23 < 7; fi23++ {
            _ = len(v3)
            fmt.Println(v2, v1, fi23)
            _ = fi23
        }
    } else {
        if (!(float32(3.14159) >= float32((((float64(0.1) - 0.1) + 1e10) * (float64((1.5 * 1e10)) * 0.1))))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if false {
            var v24 int32 = -2147483648
            _ = v24
        } else {
            fmt.Println(v1, v1)
            if len(v3) > 0 {
                v3[0] = (^(v2))
            }
        }
    }
    v25 := GenS2{F0: 281}
    if (((2 == 969) && false) && (1e10 >= (3.14159 / 1.0))) {
        if len(v3) > 0 {
            v3[0] = (v2 << ((898) & 63))
        }
        v1 = (^(v1))
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v26 int = 0
        _ = v26
    }
    v27 := GenS3{F0: 1}
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 uint = 42
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := GenS4{F0: 142, F1: (v1 * (v1 & 1000)), F2: -575}
        var v32 bool = true
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 bool = false
    var v34 uint8 = 221
    v35 := GenS5{F0: 0}
    if (false || v33) {
        if false {
            v34 += (^(v34))
            if (v34 < v34) {
                var v36 int = 140
                _ = v36
            }
            var v37 uint8 = v34
            _ = v37
        } else {
            switch 0 {
            case 32768:
                var v38 string = "foo bar"
                _ = v38
            default:
                var v39 int16 = v1
                _ = v39
            }
        }
        v2 -= 2147483647
        var v40 float32 = (float32((-2.25 * float64(0.0))) - (0.1 + float32(-2.25)))
        _ = v40
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        if (4294967295 != 42) {
            _ = len(v3)
            var v42 int8 = int8(fi41)
            _ = v42
        }
        for fi43 := 0; fi43 < 7; fi43++ {
            if false {
                switch fi41 {
                case -786:
                    if true {
                        v33 = (!v33)
                    }
                case 127:
                    v44 := GenS6{F0: 2, F1: 0}
                    _ = v44
                case 1:
                    v45 := GenS6{F0: 4294967295, F1: 32767}
                    _ = v45
                default:
                    v2 = v2
                }
                switch (fi41 | (fi41 / (((fi43 << ((18446744073709551615) & 63))) | 1))) {
                case 0:
                    var v46 int8 = -6
                    _ = v46
                }
            } else {
                fmt.Println(v1, v33)
                fmt.Println(v33, v34)
            }
            var v47 int16 = int16(fi43)
            v48 := GenS1{F0: int64(v47), F1: 66, F2: v33}
            _ = fi43
            _ = v47
            _ = v48
        }
        _ = fi41
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        v34 = v34
        switch "foo bar" {
        case "foo bar":
            v1 = (-(v1))
        case "hello":
            v34 = v34
        case "a":
            switch (-(fi49)) {
            case 0:
                v50 := GenS7{F0: uint64(v2)}
                _ = v50
            case 42:
                fmt.Println(v1, v1)
            case 2147483647:
                var v51 float64 = 0.1
                _ = v51
            default:
                if len(v3) > 0 {
                    v3[0] = (v2 % (((v2 << ((65535) & 63))) | 1))
                }
            }
        }
        v34 -= v34
        _ = fi49
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
