package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint32; F2 int8 }

func main() {
    var v1 uint32 = 0
    v2 := []string{"\n\t", "", "\n\t", ""}
    var v3 uint32 = (v1 & (v1 / ((v1) | 1)))
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    if true {
        switch 42 {
        case -218:
            v3 = v1
        case -2147483648:
            v4[("\n\t" + "\n\t")] = ("hello" + "foo bar")
        case -1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            switch "hello" {
            case "hello":
                for fi6 := 0; fi6 < 3; fi6++ {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        var v8 int16 = int16(v3)
                        _ = fi7
                        _ = v8
                    }
                    _ = fi6
                }
            case "a":
                v9 := GenS1{F0: 1, F1: ("a" + "\n\t"), F2: 42}
                _ = v9
            case "foo bar":
                switch ("a" + "\n\t") {
                case "":
                    if false {
                        v10 := GenS1{F0: 1, F1: (("hello" + "") + "a"), F2: 127}
                        var v11 uint16 = 42
                        var v12 uint32 = v1
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        if ((v1 | (v1 | 238)) <= (^(v3))) {
                            fmt.Println(v3)
                        } else {
                            var v13 bool = (true && false)
                            v1 += 4294967295
                            fmt.Println(v1, v1, v1)
                            _ = v13
                        }
                        v3 = v3
                    }
                case "hello":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        var v15 int = (fi14 | 98)
                        _ = fi14
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 2; fi16++ {
                        _ = len(v2)
                        fmt.Println(v3, v1, fi16)
                        _ = fi16
                    }
                }
            }
        }
        if (false || true) {
            switch 4294967295 {
            case -2147483648:
                for fi17 := 0; fi17 < 12; fi17++ {
                    v18 := GenS2{F0: 990}
                    if false {
                        for fi19 := 0; fi19 < 12; fi19++ {
                            switch (fi19 * (fi19 >> ((4294967295) & 63))) {
                            case 32768:
                                fmt.Println(fi19, v3)
                            case 42:
                                delete(v5, fi19)
                            default:
                                v20 := GenS2{F0: 559}
                                _ = v20
                            }
                            switch (fi19 ^ fi19) {
                            case -538:
                                v21 := GenS2{F0: -1}
                                _ = v21
                            }
                            v22 := v5[4294967295]
                            _ = fi19
                            _ = v22
                        }
                        var v23 uint16 = 506
                        switch ("\n\t" + "\n\t") {
                        case "":
                            var v24 uint = 255
                            _ = v24
                        case "hello":
                            if false {
                                var v25 int16 = -1
                                _ = v25
                            }
                        }
                        _ = v23
                    } else {
                        var v26 uint32 = (v3 >> ((1) & 63))
                        switch ("" + ("hello" + "foo bar")) {
                        case "\n\t":
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 int16 = 297
                                if len(v2) > 0 {
                                    v2[0] = (("a" + "") + "")
                                }
                                _ = fi27
                                _ = v28
                            }
                        case "hello":
                            fmt.Println(v1, v3)
                        default:
                            v26 = (^(v1))
                        }
                        var v29 uint16 = 682
                        _ = v26
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 int8 = int8(fi17)
                        var v32 uint32 = 42
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = fi17
                    _ = v18
                }
            case 660:
                v3 = (v3 & v1)
            case 128:
                var v33 int32 = 1
                _ = v33
            default:
                v34 := GenS1{F0: 216, F1: (("\n\t" + "a") + ("\n\t" + ("\n\t" + "\n\t"))), F2: 128}
                _ = v34
            }
            delete(v4, (("\n\t" + "hello") + "\n\t"))
            for fi35 := 0; fi35 < 12; fi35++ {
                v36 := GenS1{F0: 708, F1: "hello", F2: 106}
                for fi37 := 0; fi37 < 14; fi37++ {
                    v1 += 814
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi37
                }
                _ = fi35
                _ = v36
            }
        } else {
            v38 := GenS1{F0: 896, F1: "a", F2: 127}
            v39 := GenS3{F0: 33}
            _ = v38
            _ = v39
        }
        v1 -= 644
    } else {
        v1 += v3
        fmt.Println(v3)
        v40 := GenS2{F0: 9223372036854775807}
        _ = v40
    }
    v3 = (v1 | (^(v1)))
    for fi41 := 0; fi41 < 2; fi41++ {
        switch "foo bar" {
        case "":
            v1 -= (v3 & (v1 | (v1 << ((1) & 63))))
        case "\n\t":
            fmt.Println(v1, v1, fi41)
        }
        fmt.Println(v1, fi41)
        _ = fi41
    }
    fmt.Println(v1, v1)
    var v42 float32 = 3.14159
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 42 {
    case 127:
        if true {
            v1 -= v3
            switch -2147483648 {
            case -9223372036854775808:
                for fi43 := 0; fi43 < 19; fi43++ {
                    v42 -= 3.14159
                    v42 -= -2.25
                    _ = fi43
                }
            default:
                if (158 != 2) {
                    delete(v4, "a")
                    fmt.Println(v42, v3)
                }
            }
        } else {
            v3 += v1
            v44 := GenS4{F0: false, F1: (v1 ^ v3), F2: 42}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v44
        }
    default:
        if false {
            var v45 float64 = (0.1 - 1e10)
            if false {
                var v46 int = 42
                v1 = v3
                _ = v46
            }
            _ = v45
        }
    }
    if (-128 >= -1) {
        fmt.Println(v42, v1, v42)
        v47 := v5[-60]
        _ = v47
    }
    v2 = append(v2, "\n\t")
    for fi48 := 0; fi48 < 5; fi48++ {
        v3 += v3
        v3 += v1
        _ = fi48
    }
    v49 := GenS4{F0: true, F1: v3, F2: 1}
    var v50 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v50
}
