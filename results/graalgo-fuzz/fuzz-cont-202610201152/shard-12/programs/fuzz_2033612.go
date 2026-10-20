package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []int{-1, -2147483648}
    var v2 uint = 42
    var v3 uint32 = 4294967295
    var v4 float32 = (1e10 + 0.1)
    var v5 bool = false
    if (true && v5) {
        for fi6 := 0; fi6 < 18; fi6++ {
            if (("hello" + "\n\t") > "") {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v8 := GenS1{F0: "foo bar", F1: 2}
                    v9 := GenS2{F0: -128, F1: (v3 << ((v2) & 63))}
                    v3 -= 310
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                var v10 uint32 = v3
                _ = v10
            }
            v5 = v5
            _ = fi6
        }
        v5 = ("\n\t" <= "hello")
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 = (!v5)
            v12 := GenS2{F0: 1, F1: v3}
            _ = fi11
            _ = v12
        }
    }
    switch 0 {
    case 1:
        for fi13 := 0; fi13 < 12; fi13++ {
            v4 -= v4
            _ = fi13
        }
    case 32767:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -2147483648:
        switch 127 {
        case 4294967295:
            var v14 int64 = -9223372036854775808
            _ = v14
        case 128:
            v15 := GenS2{F0: 100, F1: (v3 | v3)}
            _ = v15
        case -2147483648:
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
        default:
            if false {
                v5 = (-1 != 1)
                v4 += v4
            }
        }
    default:
        v4 += 1.5
    }
    v16 := GenS1{F0: "a", F1: 36}
    switch 1 {
    case -338:
        v2 = 2
    case 32768:
        switch "a" {
        case "foo bar":
            v2 = v2
        default:
            var v17 int8 = 1
            _ = v17
        }
    }
    if (!(1 >= -498)) {
        var v18 int = 42
        _ = v18
    } else {
        fmt.Println(v3, v4, v4)
        var v19 uint8 = 255
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v22 := GenS3{F0: 567}
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        _ = v19
    }
    _ = len(v1)
    var v23 int64 = -631
    var v24 int16 = -1
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        if (255 != (v2 * v2)) {
            if (!(((("a" + "") + "\n\t") + ("foo bar" + "a")) == "a")) {
                v26 := GenS3{F0: 32767}
                var v27 int = -1
                if v5 {
                    for fi28 := 0; fi28 < 5; fi28++ {
                        fmt.Println(fi25)
                        var v29 int8 = 0
                        v3 += (^(v3))
                        _ = fi28
                        _ = v29
                    }
                    var v30 int32 = 1
                    fmt.Println(v27, v30, v23)
                    _ = v30
                }
                _ = v26
                _ = v27
            }
            var v31 int = (fi25 * fi25)
            v23 += 9223372036854775807
            _ = v31
        }
        var v32 int64 = (v23 >> (((v2 ^ v2)) & 63))
        _ = fi25
        _ = v32
    }
    switch -2147483648 {
    case 42:
        if ((((v4 + 1e10) + float32(float64((((1.5 + 1e10) / v4) + 0.0)))) + (v4 / (float32(float64(0.1)) * -2.25))) == 1e10) {
            switch 922 {
            case -872:
                v5 = (((v24 << (((^(v2))) & 63)) != v24) && (!v5))
            case 128:
                for fi33 := 0; fi33 < 20; fi33++ {
                    if true {
                        var v34 int8 = 9
                        v34 += v34
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v2 = 1
                            var v36 int16 = int16(v34)
                            if false {
                                var v37 uint8 = 1
                                _ = v37
                            } else {
                                var v38 bool = ((v4 != v4) && v5)
                                _ = v38
                            }
                            _ = fi35
                            _ = v36
                        }
                        _ = v34
                    }
                    fmt.Println(v4)
                    for fi39 := 0; fi39 < 15; fi39++ {
                        var v40 int64 = -9223372036854775808
                        _ = fi39
                        _ = v40
                    }
                    _ = fi33
                }
            default:
                v1 = append(v1, -1)
            }
            var v41 int32 = 91
            _ = v41
        } else {
            v42 := GenS4{F0: -681}
            _ = v42
        }
    case 1:
        v5 = true
    default:
        v24 = -1
    }
    v24 -= (v24 << ((v2) & 63))
    v43 := GenS5{F0: 0, F1: -5}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
