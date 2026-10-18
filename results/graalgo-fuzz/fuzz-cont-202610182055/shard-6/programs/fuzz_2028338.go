package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 2
    v3 := make(map[int]int)
    var v4 bool = (true || false)
    v5 := []int{891}
    var v6 uint32 = 600
    fmt.Println(v2)
    v2 = v2
    var v7 uint = uint(v1)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint64 = 42
        var v10 int32 = -366
        v2 -= v2
        _ = fi8
        _ = v10
        _ = v9
    }
    if (v4 && (v2 == v2)) {
        switch ("foo bar" + "\n\t") {
        case "hello":
            var v11 uint8 = v1
            _ = v11
        case "a":
            var v12 string = (("hello" + "") + "a")
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 int32 = -1
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: 1}
                _ = fi13
                _ = v16
            }
        }
    }
    var v17 uint8 = 85
    fmt.Println(v17, v7)
    fmt.Println(v1, v2, v1)
    for fi18 := 0; fi18 < 1; fi18++ {
        switch fi18 {
        case 4294967295:
            if (-128 != -1) {
                var v19 uint16 = 2
                v4 = true
                _ = v19
            }
        case 127:
            if (42 > -234) {
                delete(v3, fi18)
                var v20 float64 = -2.25
                var v21 uint = 42
                _ = v20
                _ = v21
            } else {
                v22 := GenS1{F0: v1}
                if ((v17 + 255) == v17) {
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -9223372036854775808:
                        v2 = uint64(v6)
                    case -25:
                        fmt.Println(v6, v17)
                    default:
                        if (!(1 >= 154)) {
                            v23 := GenS2{F0: v7, F1: -18, F2: (((v4 || true) || true) || true)}
                            if (18446744073709551615 < uint64(v17)) {
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v2, v7, v1)
                                    case "\n\t":
                                        if v4 {
                                            v17 -= 0
                                        }
                                        if (1 < 510) {
                                            if v4 {
                                                var v25 int = fi18
                                                var v26 uint8 = 254
                                                if false {
                                                    if v4 {
                                                        var v27 float32 = float32(3.14159)
                                                        fmt.Println(v26, v7, v25)
                                                        _ = v27
                                                    }
                                                    var v28 uint16 = 32767
                                                    _ = v28
                                                }
                                                _ = v25
                                                _ = v26
                                            }
                                            v6 = (v6 * v6)
                                            if false {
                                                switch ("foo bar" + ("hello" + ("a" + "foo bar"))) {
                                                case "foo bar":
                                                    var v29 bool = true
                                                    _ = v29
                                                case "\n\t":
                                                    var v30 float64 = (-2.25 / 3.14159)
                                                    _ = v30
                                                }
                                            }
                                            switch (fi24 << ((1) & 63)) {
                                            case 32768:
                                                var v31 int8 = 127
                                                _ = v31
                                            default:
                                                v5 = append(v5, fi24)
                                            }
                                            v5 = append(v5, (fi18 >> ((v7) & 63)))
                                        }
                                        v32 := GenS2{F0: (v7 - (v7 - v7)), F1: 127, F2: v4}
                                        _ = v32
                                    case "foo bar":
                                        var v33 uint64 = 2
                                        _ = v33
                                    }
                                    v4 = false
                                    _ = fi24
                                }
                                switch "hello" {
                                case "":
                                    _ = len(v5)
                                default:
                                    var v34 uint = v7
                                    _ = v34
                                }
                                if (!(v4 || v4)) {
                                    var v35 float32 = 0.1
                                    var v36 float64 = 3.14159
                                    if (v4 && (705 >= 32767)) {
                                        v37 := v3[(fi18 & fi18)]
                                        if v4 {
                                            switch "hello" {
                                            case "\n\t":
                                                v36 = v36
                                            }
                                        } else {
                                            for fi38 := 0; fi38 < 12; fi38++ {
                                                v39 := GenS1{F0: v1}
                                                v40 := GenS2{F0: uint(v17), F1: 1, F2: (!(0 <= 34))}
                                                _ = len(v5)
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        var v41 uint32 = (^(v6))
                                        _ = v37
                                        _ = v41
                                    } else {
                                        switch -2147483648 {
                                        case 919:
                                            var v42 int32 = -1
                                            _ = v42
                                        default:
                                            _ = len(v3)
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                } else {
                                    v4 = (fi18 == fi18)
                                }
                            }
                            _ = v23
                        } else {
                            switch (fi18 - int(v1)) {
                            case 9223372036854775807:
                                v1 = v1
                            case 32767:
                                fmt.Println(fi18, v6, v4)
                            case 28:
                                v43 := GenS3{F0: v2}
                                _ = v43
                            }
                            v1 -= 181
                            for fi44 := 0; fi44 < 9; fi44++ {
                                var v45 int64 = 788
                                v7 += 255
                                _ = fi44
                                _ = v45
                            }
                        }
                    }
                    var v46 float32 = 1.5
                    _ = v46
                } else {
                    var v47 float32 = (float32(((3.14159 + 1e10) * 1e10)) - 3.14159)
                    _ = v47
                }
                _ = v22
            }
        case 1:
            _ = len(v3)
        default:
            for fi48 := 0; fi48 < 6; fi48++ {
                switch (((("hello" + (("a" + ("a" + "")) + "a")) + "foo bar") + "\n\t") + ("\n\t" + ("hello" + "hello"))) {
                case "a":
                    var v49 float32 = 0.0
                    _ = v49
                case "":
                    v50 := GenS3{F0: v2}
                    _ = v50
                default:
                    v4 = ((1e10 >= 0.1) && v4)
                }
                v51 := GenS4{F0: (v2 + (v2 << (((v7 * uint(v2))) & 63)))}
                _ = fi48
                _ = v51
            }
        }
        v52 := GenS3{F0: (v2 & v2)}
        v2 -= (^(v2))
        _ = fi18
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
