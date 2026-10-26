package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }

func main() {
    var v1 int64 = -711
    var v2 uint64 = 18446744073709551615
    var v3 int64 = (^(v1))
    v4 := []int{1, 127}
    v5 := GenS1{F0: 2, F1: -9223372036854775808, F2: ("\n\t" <= "foo bar")}
    _ = len(v4)
    if (0.0 != (float32(1.5) * 3.14159)) {
        if true {
            var v6 uint16 = 110
            _ = v6
        } else {
            fmt.Println(v1, v1)
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS2{F0: true, F1: fi7, F2: 0}
            switch ("hello" + "\n\t") {
            case "hello":
                v1 = 9223372036854775807
            case "":
                v2 = 1
            case "foo bar":
                v1 = (v3 >> ((34) & 63))
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint16 = 341
                if false {
                    var v11 string = ""
                    v12 := GenS2{F0: false, F1: -2147483648, F2: 2147483648}
                    switch fi9 {
                    case 2147483647:
                        v3 = v3
                        v13 := GenS3{F0: 42, F1: v1}
                        var v14 bool = true
                        _ = v13
                        _ = v14
                    case 32768:
                        switch (v11 + v11) {
                        case "foo bar":
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = (fi7 - fi7)
                                }
                                fmt.Println(fi9)
                            } else {
                                if (!(v2 >= v2)) {
                                    var v15 int32 = 2147483647
                                    _ = v15
                                }
                                fmt.Println(fi9, v3)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            if len(v4) > 0 {
                                v4[0] = (fi7 << ((424) & 63))
                            }
                        case "\n\t":
                            switch fi7 {
                            case -1:
                                var v16 int16 = 32767
                                _ = v16
                            default:
                                switch (fi7 << ((514) & 63)) {
                                case 2147483647:
                                    switch (fi9 >> ((1) & 63)) {
                                    case -2147483648:
                                        fmt.Println(fi9, v3, v3)
                                    case 1:
                                        v17 := GenS2{F0: (!(!false)), F1: (fi9 - (fi9 ^ fi9)), F2: 42}
                                        _ = v17
                                    case -506:
                                        v18 := GenS2{F0: true, F1: fi9, F2: 2147483648}
                                        _ = v18
                                    }
                                default:
                                    fmt.Println(v1, v1)
                                }
                            }
                        default:
                            v10 = v10
                        }
                    default:
                        if (v2 == 18446744073709551615) {
                            v11 = ((((v11 + "a") + v11) + v11) + "\n\t")
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v19 int64 = (v1 * v1)
                            if false {
                                v10 = 0
                                switch (fi7 + fi9) {
                                case 42:
                                    v20 := GenS1{F0: (^(v10)), F1: fi7, F2: false}
                                    _ = v20
                                case -263:
                                    switch 32768 {
                                    case -9223372036854775808:
                                        v19 = (v3 - (v19 + (v1 | (-(v3)))))
                                    case -137:
                                        fmt.Println(v10, fi7, v2)
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = fi9
                                        }
                                    }
                                default:
                                    v21 := GenS2{F0: (255 >= 0), F1: fi9, F2: 39}
                                    _ = v21
                                }
                                if len(v4) > 0 {
                                    v4[0] = (fi7 ^ (fi9 ^ fi7))
                                }
                            }
                            _ = v19
                        }
                    }
                    _ = v11
                    _ = v12
                }
                var v22 uint16 = uint16(v2)
                _ = fi9
                _ = v10
                _ = v22
            }
            _ = fi7
            _ = v8
        }
        if ((v3 >> ((65535) & 63)) <= (v3 / ((-41) | 1))) {
            v23 := GenS3{F0: 42, F1: 9223372036854775807}
            var v24 bool = (false || (!(42 >= 1)))
            _ = v23
            _ = v24
        }
    } else {
        switch "a" {
        case "foo bar":
            for fi25 := 0; fi25 < 17; fi25++ {
                v3 = v1
                _ = fi25
            }
        case "":
            v2 = v2
        }
    }
    var v26 uint8 = 204
    fmt.Println(v26)
    v2 -= v2
    switch -721 {
    case 128:
        v1 = 0
    default:
        var v27 int16 = -32768
        _ = v27
    }
    switch "foo bar" {
    case "\n\t":
        v28 := GenS3{F0: 0, F1: (v1 & v1)}
        _ = v28
    default:
        switch 32768 {
        case 9223372036854775807:
            v29 := GenS2{F0: ((("a" + "a") + "hello") != "\n\t"), F1: 9223372036854775807, F2: 620}
            _ = v29
        case -994:
            for fi30 := 0; fi30 < 3; fi30++ {
                v1 -= (v3 << ((271) & 63))
                v26 += v26
                v31 := GenS4{F0: 42, F1: "\n\t", F2: v3}
                _ = fi30
                _ = v31
            }
        case 2147483647:
            v4 = append(v4, 1)
        default:
            if (0 < -2147483648) {
                fmt.Println(v3, v2, v3)
            } else {
                switch int(v2) {
                case 128:
                    if ((^(v1)) != v1) {
                        v26 -= 0
                    }
                default:
                    v1 -= 0
                }
            }
        }
    }
    var v32 uint8 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
