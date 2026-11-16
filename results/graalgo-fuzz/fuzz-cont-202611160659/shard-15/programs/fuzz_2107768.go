package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 bool; F1 int }

func main() {
    var v1 uint16 = 2
    var v2 string = ""
    var v3 int32 = 42
    v4 := []uint32{371, 2}
    v5 := make(map[uint32]uint32)
    var v6 int = -384
    switch v6 {
    case 2147483647:
        for fi7 := 0; fi7 < 8; fi7++ {
            if true {
                v6 = (v6 + -2147483648)
                var v8 int64 = 42
                for fi9 := 0; fi9 < 6; fi9++ {
                    v1 = 2
                    var v10 int8 = 42
                    if ((!(!(1.5 <= ((float32((float64(1e10) - 1e10)) / 1e10) / ((1e10 + float32(-2.25)) + ((3.14159 * (1.5 - 1.5)) / (-2.25 - 3.14159))))))) && true) {
                        switch v6 {
                        case -1:
                            var v11 uint8 = 48
                            _ = v11
                        default:
                            v12 := GenS1{F0: (v10 | 0), F1: uint8(v1), F2: v10}
                            _ = v12
                        }
                        v1 -= (v1 / ((v1) | 1))
                        if (true && true) {
                            fmt.Println(fi9)
                            fmt.Println(v2)
                        } else {
                            var v13 int16 = 0
                            _ = v13
                        }
                        if false {
                            var v14 uint16 = 32767
                            v2 = v2
                            v15 := GenS1{F0: (v10 * -1), F1: 42, F2: v10}
                            _ = v14
                            _ = v15
                        }
                    } else {
                        v1 -= v1
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
            _ = fi7
        }
    case 42:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -12:
        v3 -= (v3 >> ((882) & 63))
    }
    v1 += (v1 % ((v1) | 1))
    var v16 float32 = 0.1
    var v17 int = -412
    var v18 uint8 = 0
    if false {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 uint16 = v1
            if true {
                if (36 <= 1) {
                    switch v17 {
                    case 2147483647:
                        fmt.Println(v1)
                    case -396:
                        var v21 float32 = v16
                        _ = v21
                    case 32768:
                        v2 = (v2 + "")
                    }
                    v2 = (v2 + v2)
                }
            }
            var v22 int8 = 1
            _ = fi19
            _ = v20
            _ = v22
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            v3 += (v3 << ((255) & 63))
            switch 9223372036854775807 {
            case 857:
                var v24 bool = false
                _ = v24
            default:
                if false {
                    v25 := v5[29]
                    v6 -= (-(v17))
                    _ = v25
                } else {
                    v17 -= (^(fi23))
                }
            }
            fmt.Println(fi23)
            _ = fi23
        }
    } else {
        _ = len(v4)
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v6, v18, v16)
            _ = fi26
        }
    }
    fmt.Println(v1)
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS2{F0: v18, F1: (^(fi27)), F2: false}
        var v29 int32 = 419
        v6 += -9223372036854775808
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint8 = v18
    v6 = -1
    var v31 uint8 = v30
    v32 := GenS1{F0: -128, F1: (v31 ^ 1), F2: -54}
    var v33 int32 = (v3 ^ v3)
    fmt.Println(v16, v30, v33)
    v34 := GenS2{F0: (v31 >> ((42) & 63)), F1: -9223372036854775808, F2: (-2.25 > float64(v16))}
    if (18446744073709551615 <= 18446744073709551615) {
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int8 = 21
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi35
            _ = v36
        }
    }
    if true {
        v1 -= v1
        fmt.Println(v16, v30, v1)
        switch (v6 ^ (-(v6))) {
        case 574:
            switch (("foo bar" + (v2 + ((("" + (v2 + v2)) + "a") + v2))) + v2) {
            case "a":
                if true {
                    delete(v5, 1)
                    delete(v5, 65535)
                } else {
                    v37 := GenS3{F0: (18446744073709551615 <= 4294967295), F1: v6}
                    _ = v37
                }
            case "hello":
                switch v2 {
                case "foo bar":
                    var v38 bool = false
                    _ = v38
                }
            }
        case -662:
            v17 = v17
        case -598:
            var v39 uint32 = 57
            _ = v39
        }
    } else {
        _ = len(v5)
        fmt.Println(v16)
        switch v17 {
        case 720:
            switch v6 {
            case -1:
                if len(v4) > 0 {
                    v4[0] = 65535
                }
            default:
                delete(v5, 519)
            }
        case -2147483648:
            var v40 float32 = float32(1e10)
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
