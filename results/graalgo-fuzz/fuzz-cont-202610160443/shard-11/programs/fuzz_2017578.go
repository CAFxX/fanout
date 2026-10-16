package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int8 = -1
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: 255}
    if (("a" + ("\n\t" + "foo bar")) >= "foo bar") {
        var v9 string = ("\n\t" + (("hello" + ("" + ("" + ""))) + "\n\t"))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 string = "hello"
            _ = fi10
            _ = v11
        }
        var v12 uint16 = 44
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 bool = false
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        _ = v12
    }
    if (!(0 <= 32767)) {
        if false {
            var v16 uint = 65535
            v16 -= v16
            _ = v16
        } else {
            var v17 int16 = -1
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := []int{fi18, (fi18 % (((fi18 / (((fi18 + (fi18 | fi18))) | 1))) | 1)), (fi18 + 127)}
                fmt.Println(v17, v17)
                _ = fi18
                _ = v19
            }
            v20 := GenS1{F0: v17, F1: 65535}
            _ = v17
            _ = v20
        }
        v21 := GenS2{F0: 527, F1: 1}
        _ = v21
    } else {
        var v22 bool = ("" >= (("" + "foo bar") + ""))
        _ = v22
    }
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{201, 735}
    var v3 int16 = 132
    var v4 float32 = (-2.25 - float32((float64(0.1) - (-2.25 - ((v1 / (v1 + 0.1)) * 3.14159)))))
    var v5 float32 = (0.1 + ((1.5 - v4) / v4))
    if true {
        v4 = 1.5
    }
    var v23 bool = (!(-9223372036854775808 >= -9223372036854775808))
    v3 -= (v3 + v3)
    var v24 int16 = v3
    if len(v2) > 0 {
        v2[0] = 651
    }
    if (42 > -293) {
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := GenS2{F0: -173, F1: 38}
            _ = fi25
            _ = v26
        }
    } else {
        var v27 int16 = v24
        v28 := GenS1{F0: (v27 % ((42) | 1)), F1: 385}
        if true {
            switch "\n\t" {
            case "a":
                var v29 uint32 = 2147483647
                _ = v29
            case "foo bar":
                var v30 int64 = 0
                _ = v30
            case "\n\t":
                switch (("foo bar" + "a") + "") {
                case "foo bar":
                    switch ("\n\t" + "") {
                    case "a":
                        var v31 uint64 = 1
                        _ = v31
                    case "":
                        fmt.Println(v5, v24, v23)
                    case "\n\t":
                        v32 := GenS2{F0: 0, F1: 48}
                        _ = v32
                    default:
                        fmt.Println(v3, v5)
                    }
                default:
                    for fi33 := 0; fi33 < 15; fi33++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if v23 {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                if (v23 && v23) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi34)
                                }
                                v27 = v24
                                _ = fi34
                            }
                        }
                        if true {
                            switch ("a" + "") {
                            case "\n\t":
                                var v35 uint16 = 2
                                _ = v35
                            case "hello":
                                switch "a" {
                                case "\n\t":
                                    v3 = 1
                                }
                            }
                            fmt.Println(v23, v1, v1)
                            if false {
                                v2 = append(v2, 1)
                            } else {
                                var v36 uint16 = 255
                                v37 := GenS2{F0: -9223372036854775808, F1: 255}
                                if false {
                                    var v38 float32 = (v5 * v4)
                                    _ = len(v2)
                                    _ = v38
                                }
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = fi33
                    }
                }
            }
        } else {
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int64 = 0
                _ = fi39
                _ = v40
            }
            v41 := GenS1{F0: (-(v27)), F1: 2}
            _ = v41
        }
        _ = v27
        _ = v28
    }
    fmt.Println(v5, v1, v5)
    var v42 string = (("\n\t" + ("" + "")) + "a")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -128 {
    case 108:
        switch v42 {
        case "a":
            v5 -= v4
        case "":
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 int8 = 127
                v4 += v5
                var v45 int64 = -1
                _ = fi43
                _ = v44
                _ = v45
            }
        case "hello":
            var v46 int32 = 0
            _ = v46
        }
    case -1:
        v4 += v4
    default:
        v23 = (0 == 42)
    }
    v3 = (v3 * 1)
    var v47 string = v42
    for fi48 := 0; fi48 < 18; fi48++ {
        _ = len(v2)
        v49 := GenS1{F0: v24, F1: 0}
        v5 = (float32(v1) - 0.1)
        _ = fi48
        _ = v49
    }
    if (406 > -2147483648) {
        v50 := GenS2{F0: -109, F1: 128}
        v51 := GenS1{F0: (v3 << ((65535) & 63)), F1: 42}
        _ = v50
        _ = v51
    }
    var v52 int8 = 42
    var v53 int16 = 0
    if (((853 <= 42) && v23) && v23) {
        if ((v24 | -643) != 1) {
            if true {
                var v54 float64 = 1e10
                if true {
                    var v55 float64 = 0.0
                    v2 = append(v2, 0)
                    _ = v55
                }
                _ = len(v2)
                _ = v54
            }
        } else {
            if v23 {
                fmt.Println(v23)
                for fi56 := 0; fi56 < 2; fi56++ {
                    v57 := GenS2{F0: 1, F1: 7}
                    _ = fi56
                    _ = v57
                }
            }
            v58 := GenS1{F0: v3, F1: 179}
            var v59 int16 = -413
            _ = v58
            _ = v59
        }
    } else {
        v60 := GenS1{F0: (v24 >> ((717) & 63)), F1: 42}
        v61 := GenS2{F0: 9223372036854775807, F1: 1}
        _ = v60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v52
    _ = v53
}
