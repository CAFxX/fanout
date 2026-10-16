package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -9223372036854775808
    var v3 uint = 609
    var v4 uint32 = 42
    v5 := make(map[string]int)
    v6 := make(map[string]bool)
    fmt.Println(v3, v2)
    v3 -= v3
    var v7 uint32 = (v4 & v4)
    v8 := []int{v2, v2, v2}
    fmt.Println(v4, v4, v4)
    if len(v8) > 0 {
        v8[0] = 9223372036854775807
    }
    v9 := GenS1{F0: (576 >= 67)}
    v2 = (v2 << ((4294967295) & 63))
    var v10 uint = 42
    if true {
        switch (v2 + -2147483648) {
        case 711:
            var v11 bool = (!(!(0 >= 641)))
            _ = v11
        case 128:
            fmt.Println(v4, v2)
        default:
            if (1 >= 42) {
                v12 := GenS2{F0: 603, F1: 1}
                v5[""] = 9223372036854775807
                switch (^(v2)) {
                case -9223372036854775808:
                    v3 -= (^(v10))
                case 4294967295:
                    var v13 int = v2
                    _ = v13
                }
                _ = v12
            } else {
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(fi14, v7)
                    if (112 >= 977) {
                        if true {
                            delete(v5, "")
                        }
                        for fi15 := 0; fi15 < 19; fi15++ {
                            var v16 float32 = 3.14159
                            if false {
                                switch (v2 % ((-1) | 1)) {
                                case 145:
                                    switch fi15 {
                                    case 32767:
                                        v5["\n\t"] = (fi15 ^ fi14)
                                    case -2147483648:
                                        v17 := GenS2{F0: 891, F1: 601}
                                        _ = v17
                                    }
                                default:
                                    if true {
                                        var v18 uint64 = 25
                                        var v19 int64 = -9223372036854775808
                                        var v20 string = ((((("hello" + "") + ("hello" + ("foo bar" + "hello"))) + "\n\t") + "") + "hello")
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := v6["\n\t"]
                                        _ = v21
                                    }
                                }
                                var v22 int16 = 42
                                var v23 int64 = 0
                                var v24 int8 = 67
                                v7 = 65535
                                var v25 uint16 = 255
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                var v26 float64 = (0.0 / 3.14159)
                                v27 := v6["foo bar"]
                                _ = v26
                                _ = v27
                            }
                            if (v3 >= v3) {
                                fmt.Println(fi14)
                                v28 := GenS3{F0: fi14}
                                v3 = v3
                                _ = v28
                            } else {
                                switch -809 {
                                case -1:
                                    switch 0 {
                                    case 4294967295:
                                        if true {
                                            switch fi15 {
                                            case -1:
                                                v6[""] = true
                                            }
                                            var v29 uint64 = 18446744073709551615
                                            _ = v29
                                        }
                                    case -2147483648:
                                        v30 := v5[("foo bar" + "foo bar")]
                                        _ = v30
                                    case -1:
                                        v3 += (v3 / ((v10) | 1))
                                    default:
                                        if true {
                                            fmt.Println(v16, v16, fi14)
                                        }
                                    }
                                case 127:
                                    if (0.1 == (0.0 - (3.14159 * 1.5))) {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        default:
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                        v8 = append(v8, 2147483647)
                                    }
                                }
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        for fi31 := 0; fi31 < 11; fi31++ {
                            if false {
                                if true {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v3, v2, fi14)
                                    case "foo bar":
                                        var v32 bool = (0 != 42)
                                        _ = v32
                                    default:
                                        switch ("\n\t" + (("hello" + "hello") + ("hello" + "\n\t"))) {
                                        case "hello":
                                            _ = len(v5)
                                        default:
                                            v33 := GenS4{F0: 1}
                                            _ = v33
                                        }
                                    }
                                    if ((float32(1e10) / 0.1) < (1.5 * 3.14159)) {
                                        var v34 int8 = 127
                                        v2 -= -9223372036854775808
                                        _ = v34
                                    } else {
                                        v3 += (v3 & v10)
                                        v35 := GenS1{F0: (0 >= -1)}
                                        _ = v35
                                    }
                                    switch "a" {
                                    case "":
                                        fmt.Println(v7)
                                    default:
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                var v36 int8 = -1
                                                _ = v36
                                            case "foo bar":
                                                var v37 uint16 = 0
                                                fmt.Println(fi14)
                                                _ = v37
                                            default:
                                                v38 := GenS3{F0: 42}
                                                _ = v38
                                            }
                                            if (true && true) {
                                                var v39 uint16 = 0
                                                switch (fi31 >> ((v3) & 63)) {
                                                case -971:
                                                    var v40 uint8 = 0
                                                    _ = v40
                                                case 9223372036854775807:
                                                    v41 := GenS5{F0: v7, F1: 1000}
                                                    _ = v41
                                                case 4294967295:
                                                    switch (("" + "hello") + (("\n\t" + (("a" + "\n\t") + "hello")) + ((((("" + ("foo bar" + "")) + "\n\t") + (("hello" + "") + "")) + "foo bar") + ("" + ("\n\t" + "\n\t"))))) {
                                                    case "a":
                                                        switch (fi14 ^ (fi31 & (v2 ^ (-(v2))))) {
                                                        case 0:
                                                            fmt.Println(v4, v10, v4)
                                                        case 4294967295:
                                                            var v42 uint16 = 32767
                                                            _ = v42
                                                        default:
                                                            var v43 int64 = int64(fi14)
                                                            _ = v43
                                                        }
                                                    case "\n\t":
                                                        switch 42 {
                                                        case 128:
                                                            _ = len(v8)
                                                        case 32768:
                                                            v44 := GenS6{F0: -19, F1: v4}
                                                            _ = v44
                                                        default:
                                                            fmt.Println(v39)
                                                        }
                                                    }
                                                    switch (v2 >> ((v3) & 63)) {
                                                    case -477:
                                                        v4 += (v7 + (^(v4)))
                                                        v45 := GenS2{F0: uint64(v10), F1: v39}
                                                        _ = v45
                                                    case 32768:
                                                        if (!(4294967295 < 2)) {
                                                            fmt.Println(v10, v10)
                                                            v2 = (^(fi14))
                                                        }
                                                    case -1:
                                                        var v46 float32 = (3.14159 / 0.1)
                                                        _ = v46
                                                    }
                                                    v47 := v1[(v7 & 246)]
                                                    _ = v47
                                                default:
                                                    var v48 int16 = -32768
                                                    _ = v48
                                                }
                                                _ = v39
                                            }
                                        } else {
                                            var v49 int8 = 64
                                            var v50 float64 = (0.1 - float64(3.14159))
                                            var v51 int16 = -32768
                                            switch (-(fi14)) {
                                            case 127:
                                                v3 = (v3 ^ (v10 | v3))
                                            case -1000:
                                                v50 = v50
                                            case 0:
                                                v52 := GenS1{F0: true}
                                                _ = v52
                                            default:
                                                if true {
                                                    v53 := v1[(v7 | 65535)]
                                                    if (255 >= 871) {
                                                        if len(v8) > 0 {
                                                            v8[0] = (fi14 / ((fi14) | 1))
                                                        }
                                                    } else {
                                                        fmt.Println(v49)
                                                    }
                                                    var v54 int = (fi14 << ((v10) & 63))
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    v55 := GenS2{F0: 42, F1: 42}
                                                    _ = v55
                                                }
                                            }
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                    }
                                } else {
                                    v8 = append(v8, -1)
                                }
                                _ = len(v8)
                            } else {
                                var v56 uint64 = 193
                                _ = v56
                            }
                            v57 := GenS1{F0: ((((1.5 < 3.14159) && (0.1 < (0.1 / -2.25))) || true) && (-9223372036854775808 != 42))}
                            _ = fi31
                            _ = v57
                        }
                        var v58 uint64 = 42
                        _ = v58
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi14
                }
            }
        }
    }
    var v59 float64 = float64(3.14159)
    var v60 string = (("hello" + "\n\t") + ("" + ("" + ("" + ("\n\t" + "foo bar")))))
    fmt.Println(v60, v59)
    fmt.Println(v2, v10)
    v2 = v2
    fmt.Println(v59)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v7
    _ = v8
    _ = v9
}
