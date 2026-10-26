package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint; F2 bool }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 int32 = 2147483647
    v4 := []int{0, 127, 32768, 127}
    var v5 int16 = -32768
    var v6 int32 = -1
    if false {
        if (4294967295 > 4294967295) {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 string = ("" + ("hello" + "foo bar"))
                v9 := GenS1{F0: 65535, F1: 42, F2: 42}
                v10 := v1[v8]
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
        v11 := GenS1{F0: 255, F1: 18446744073709551615, F2: 42}
        for fi12 := 0; fi12 < 19; fi12++ {
            v13 := v1["hello"]
            var v14 int64 = -240
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    var v15 int = -9223372036854775808
    v5 -= (v5 << ((65535) & 63))
    var v16 bool = (((1.5 + 0.0) < -2.25) && false)
    _ = len(v4)
    var v17 uint16 = 42
    var v18 float32 = (3.14159 / 1e10)
    v19 := GenS1{F0: 670, F1: 0, F2: 150}
    if v16 {
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS2{F0: (v17 * v17), F1: v2}
            if false {
                v16 = v16
                var v22 int32 = 0
                _ = v22
            }
            _ = fi20
            _ = v21
        }
    }
    if (!(!("" >= (("foo bar" + ("\n\t" + ((((("a" + "foo bar") + (((("foo bar" + "\n\t") + "") + "\n\t") + "foo bar")) + "") + "foo bar") + ""))) + "a")))) {
        for fi23 := 0; fi23 < 17; fi23++ {
            if v16 {
                var v24 uint8 = 14
                var v25 uint32 = 42
                _ = v24
                _ = v25
            } else {
                switch ((("a" + "\n\t") + "hello") + "") {
                case "\n\t":
                    delete(v1, "a")
                case "hello":
                    switch ("hello" + "hello") {
                    case "hello":
                        if false {
                            if (42 >= 42) {
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    var v27 int16 = 1000
                                    v27 += v5
                                    _ = fi26
                                    _ = v27
                                }
                                var v28 uint16 = (v17 >> ((573) & 63))
                                if (!v16) {
                                    v1["\n\t"] = 2147483648
                                    var v29 uint = 972
                                    if (v3 != 42) {
                                        if len(v4) > 0 {
                                            v4[0] = v15
                                        }
                                        v6 -= 0
                                    } else {
                                        if (!((-2.25 - 1e10) > 1e10)) {
                                            _ = len(v1)
                                            if v16 {
                                                var v30 uint64 = 167
                                                delete(v1, ("a" + "a"))
                                                switch 42 {
                                                case 357:
                                                    if v16 {
                                                        v18 -= (v18 + v18)
                                                        var v31 float64 = 1.5
                                                        fmt.Println(v2, v3, v29)
                                                        _ = v31
                                                    }
                                                default:
                                                    var v32 uint = v29
                                                    _ = v32
                                                }
                                                _ = v30
                                            }
                                        } else {
                                            var v33 uint = 42
                                            if (false && false) {
                                                fmt.Println(v28)
                                                switch "" {
                                                case "foo bar":
                                                    v34 := GenS1{F0: (v29 ^ (^(v33))), F1: v29, F2: 854}
                                                    _ = v34
                                                }
                                                for fi35 := 0; fi35 < 9; fi35++ {
                                                    var v36 uint16 = (v28 | v28)
                                                    v37 := GenS3{F0: 4294967295}
                                                    _ = fi35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            } else {
                                                var v38 int64 = 0
                                                var v39 int8 = 1
                                                if (!true) {
                                                    if v16 {
                                                        v2 += (v2 + (v2 * 42))
                                                        v1[("hello" + "a")] = 2147483647
                                                    } else {
                                                        v40 := GenS3{F0: 2}
                                                        v16 = v16
                                                        v41 := GenS4{F0: (v6 / ((v2) | 1)), F1: uint64(v33), F2: 42}
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    var v42 uint32 = 616
                                                    if false {
                                                        for fi43 := 0; fi43 < 1; fi43++ {
                                                            v44 := GenS5{F0: (^(v5)), F1: uint(v42), F2: v16}
                                                            if true {
                                                                v2 = (^(v6))
                                                                v3 -= v6
                                                                fmt.Println(v6, v3, v38)
                                                            }
                                                            _ = fi43
                                                            _ = v44
                                                        }
                                                        if v16 {
                                                            switch "a" {
                                                            case "hello":
                                                                if (0 < 178) {
                                                                    v16 = v16
                                                                    for fi45 := 0; fi45 < 12; fi45++ {
                                                                        switch 4294967295 {
                                                                        case 32767:
                                                                            v46 := GenS6{F0: v5, F1: "a"}
                                                                            _ = v46
                                                                        }
                                                                        v6 = v6
                                                                        v5 = (v5 - (v5 + v5))
                                                                        _ = fi45
                                                                    }
                                                                    v47 := v1[(("a" + "") + ("hello" + "a"))]
                                                                    _ = v47
                                                                }
                                                            case "\n\t":
                                                                if v16 {
                                                                    for fi48 := 0; fi48 < 19; fi48++ {
                                                                        switch (("foo bar" + ("a" + "")) + ("\n\t" + (("" + "a") + ("" + "foo bar")))) {
                                                                        case "\n\t":
                                                                            delete(v1, "a")
                                                                        case "":
                                                                            if (uint(v42) <= (v29 << ((18446744073709551615) & 63))) {
                                                                                delete(v1, (("hello" + "foo bar") + "hello"))
                                                                                var v49 uint = (v29 & v29)
                                                                                _ = v49
                                                                            } else {
                                                                                var v50 int16 = 42
                                                                                v15 = (fi23 & fi23)
                                                                                v2 = v6
                                                                                _ = v50
                                                                            }
                                                                        case "foo bar":
                                                                            v15 += -9223372036854775808
                                                                            var v51 uint64 = 2
                                                                            fmt.Println(v28)
                                                                            _ = v51
                                                                        default:
                                                                            var v52 int32 = (v2 + (-(v3)))
                                                                            _ = v52
                                                                        }
                                                                        v53 := GenS4{F0: (v6 - (v6 << (((v29 + (^(v29)))) & 63))), F1: 42, F2: (^(v42))}
                                                                        delete(v1, "")
                                                                        _ = fi48
                                                                        _ = v53
                                                                    }
                                                                    v54 := GenS7{F0: (-(v15)), F1: v28}
                                                                    _ = v54
                                                                }
                                                            case "":
                                                                fmt.Println(v15, v18)
                                                            default:
                                                                if (v16 || (v42 != 42)) {
                                                                    v55 := GenS6{F0: (v5 | (v5 >> (((^(v33))) & 63))), F1: "a"}
                                                                    _ = v55
                                                                }
                                                            }
                                                            v56 := GenS8{F0: v17, F1: 230}
                                                            _ = v56
                                                        }
                                                    } else {
                                                        if v16 {
                                                            var v57 int8 = (v39 << ((466) & 63))
                                                            _ = v57
                                                        }
                                                    }
                                                    _ = v42
                                                }
                                                _ = v38
                                                _ = v39
                                            }
                                            var v58 int64 = -1
                                            _ = v33
                                            _ = v58
                                        }
                                        _ = len(v4)
                                    }
                                    _ = v29
                                }
                                _ = v28
                            } else {
                                var v59 uint8 = 0
                                _ = v59
                            }
                            var v60 float64 = 3.14159
                            _ = v60
                        }
                    case "\n\t":
                        var v61 bool = v16
                        _ = v61
                    case "foo bar":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                default:
                    fmt.Println(v16, v5)
                }
                _ = len(v4)
            }
            var v62 uint32 = 4294967295
            switch 127 {
            case 128:
                for fi63 := 0; fi63 < 14; fi63++ {
                    delete(v1, "hello")
                    v17 -= v17
                    v62 = v62
                    _ = fi63
                }
            case -2147483648:
                var v64 int32 = (^(v2))
                _ = v64
            }
            _ = fi23
            _ = v62
        }
    } else {
        var v65 int8 = 1
        v66 := GenS1{F0: 171, F1: 909, F2: 4294967295}
        _ = v65
        _ = v66
    }
    v1["a"] = 53
    for fi67 := 0; fi67 < 5; fi67++ {
        v15 = (^(v15))
        v16 = ((0.0 + (float64((v18 - 0.0)) * 0.1)) <= (0.0 - (((float64(v18) + 3.14159) + 0.1) / (1e10 / -2.25))))
        _ = fi67
    }
    var v68 bool = true
    v6 = -257
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
