package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{128, 9223372036854775807}
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 int32 = -1
    var v5 uint8 = 1
    var v6 float32 = -2.25
    if false {
        var v7 int16 = -322
        var v8 int = 2147483647
        v9 := GenS1{F0: 0, F1: 620}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 int16 = 1
        _ = v10
    }
    switch 1 {
    case 128:
        fmt.Println(v4)
    case 42:
        for fi11 := 0; fi11 < 9; fi11++ {
            if (fi11 <= (^(fi11))) {
                v3["a"] = "\n\t"
                switch (fi11 + 127) {
                case -891:
                    for fi12 := 0; fi12 < 17; fi12++ {
                        var v13 bool = true
                        switch (fi12 + fi11) {
                        case 739:
                            var v14 uint = 691
                            _ = v14
                        case -9223372036854775808:
                            fmt.Println(v4, fi11)
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = fi11
        }
    default:
        if (v6 == (((3.14159 - v6) / 0.1) * (((v6 / ((v6 + v6) * v6)) + ((float32(float64(1.5)) + v6) - v6)) / float32(1e10)))) {
            if true {
                v15 := GenS1{F0: 42, F1: v4}
                _ = v15
            }
        } else {
            v6 = 1.5
            v16 := GenS1{F0: 63, F1: (^(v4))}
            v17 := GenS1{F0: 4294967295, F1: (v4 & v4)}
            _ = v16
            _ = v17
        }
    }
    if true {
        v5 += v5
        var v18 uint16 = 1
        v18 += (^(v18))
        _ = v18
    }
    v2 = append(v2, 397)
    var v19 float32 = v6
    fmt.Println(v6, v5, v4)
    v19 -= v6
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            if false {
                var v22 string = ("a" + "")
                _ = v22
            }
            _ = fi21
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 string = ("\n\t" + ("foo bar" + "a"))
            v2 = append(v2, fi20)
            _ = fi23
            _ = v24
        }
        _ = fi20
    }
    switch ("" + "foo bar") {
    case "hello":
        v4 = -1
    }
    switch "" {
    case "foo bar":
        switch 142 {
        case 128:
            v4 = -2147483648
        case -836:
            v19 -= (-2.25 + float32((3.14159 * 1.5)))
        case 42:
            v19 -= 0.1
        }
    case "hello":
        fmt.Println(v5)
    }
    var v25 uint64 = 42
    v2 = append(v2, -9223372036854775808)
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "foo bar" {
        case "hello":
            var v26 uint8 = (^(v5))
            _ = v26
        case "\n\t":
            v5 += 142
        case "a":
            var v27 float32 = 0.1
            _ = v27
        default:
            if false {
                if len(v1) > 0 {
                    v1[0] = 32767
                }
            } else {
                switch 4294967295 {
                case -9223372036854775808:
                    if true {
                        v28 := GenS1{F0: 733, F1: v4}
                        if (!true) {
                            var v29 uint = 327
                            v5 = (v5 + (v5 + v5))
                            _ = v29
                        } else {
                            switch 0 {
                            case 42:
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v19)
                                default:
                                    var v30 float32 = v6
                                    _ = v30
                                }
                            case 127:
                                delete(v3, ("" + "a"))
                            case 4294967295:
                                var v31 float64 = (-2.25 * (-2.25 * 1e10))
                                _ = v31
                            default:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        var v32 int = 42
                        _ = v28
                        _ = v32
                    } else {
                        switch (("a" + "\n\t") + (("a" + "hello") + ("a" + "a"))) {
                        case "":
                            var v33 int16 = 32767
                            _ = v33
                        case "\n\t":
                            var v34 int8 = 84
                            _ = v34
                        default:
                            v2 = append(v2, -231)
                        }
                        fmt.Println(v19)
                    }
                case 254:
                    v35 := GenS1{F0: 65535, F1: (^(v4))}
                    _ = v35
                case -30:
                    fmt.Println(v4, v25)
                default:
                    switch 0 {
                    case 559:
                        v36 := GenS2{F0: v4}
                        _ = v36
                    case 42:
                        v37 := GenS1{F0: uint(v25), F1: (v4 % ((-2147483648) | 1))}
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 uint64 = 0
                    _ = fi38
                    _ = v39
                }
                var v40 int8 = 106
                _ = v40
            }
        }
    } else {
        v6 += v6
    }
    fmt.Println(v19, v6, v25)
    if (true || (799 >= 18446744073709551615)) {
        v41 := GenS1{F0: 719, F1: (v4 << ((495) & 63))}
        if (false || ("hello" <= "hello")) {
            var v42 int = 42
            _ = v42
        }
        _ = v41
    } else {
        fmt.Println(v5)
        var v43 uint32 = 0
        _ = v43
    }
    if false {
        if (!(0 >= 255)) {
            fmt.Println(v4, v5, v25)
        } else {
            v6 = v6
        }
        fmt.Println(v19)
        fmt.Println(v25)
    } else {
        for fi44 := 0; fi44 < 15; fi44++ {
            switch (-(fi44)) {
            case 682:
                v2 = append(v2, 32768)
            case 1:
                var v45 float32 = v19
                _ = v45
            default:
                var v46 uint16 = 255
                _ = v46
            }
            _ = fi44
        }
    }
    v47 := GenS1{F0: 582, F1: v4}
    for fi48 := 0; fi48 < 20; fi48++ {
        var v49 bool = false
        fmt.Println(fi48, fi48)
        _ = fi48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
