package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint32 = 1
    v2 := []uint32{v1, 2147483648, v1}
    var v3 uint16 = 1
    var v4 int32 = -1
    v5 := make(map[uint32]uint32)
    _ = len(v5)
    var v6 uint32 = (^(v1))
    var v7 bool = false
    fmt.Println(v6)
    var v8 uint64 = 42
    var v9 int8 = 42
    switch ((("foo bar" + "\n\t") + "hello") + "a") {
    case "":
        if v7 {
            v4 = int32(v8)
            v10 := GenS1{F0: 1, F1: 42}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                if ("hello" <= "\n\t") {
                    v5[65535] = (v6 & 933)
                    var v12 int8 = 127
                    v13 := GenS1{F0: (v9 + (v12 | v9)), F1: 42}
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        switch "\n\t" {
                        case "a":
                            var v15 int64 = 42
                            _ = v15
                        case "hello":
                            v8 += uint64(v3)
                        }
                        var v16 string = (("foo bar" + "\n\t") + ("" + ""))
                        _ = fi14
                        _ = v16
                    }
                    var v17 uint16 = v3
                    fmt.Println(v6, v17)
                    _ = v17
                }
                fmt.Println(v1)
                _ = fi11
            }
        }
    case "a":
        v2 = append(v2, (v6 + (v1 >> ((42) & 63))))
    case "hello":
        delete(v5, v6)
    default:
        v18 := GenS1{F0: -126, F1: 1}
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v7 = v7
        v9 = (v9 * v9)
        var v20 float32 = 0.1
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v22 uint32 = v6
        v2 = append(v2, 0)
        _ = fi21
        _ = v22
    }
    var v23 int64 = -1
    if (-1 > -1) {
        if v7 {
            fmt.Println(v8)
        } else {
            delete(v5, v1)
            if v7 {
                for fi24 := 0; fi24 < 19; fi24++ {
                    switch (-(fi24)) {
                    case 415:
                        switch (fi24 >> ((125) & 63)) {
                        case 0:
                            v25 := GenS1{F0: 0, F1: (v23 << ((802) & 63))}
                            _ = v25
                        case 127:
                            v26 := GenS1{F0: -1, F1: v23}
                            _ = v26
                        }
                    default:
                        for fi27 := 0; fi27 < 15; fi27++ {
                            var v28 uint64 = uint64(v3)
                            _ = fi27
                            _ = v28
                        }
                    }
                    _ = len(v2)
                    _ = fi24
                }
                v9 = 0
                v9 = -1
            }
        }
        var v29 uint16 = 1
        switch "foo bar" {
        case "hello":
            var v30 int8 = v9
            _ = v30
        default:
            if (true && (v8 < v8)) {
                var v31 int32 = -1
                _ = v31
            } else {
                var v32 uint16 = (v29 << ((0) & 63))
                _ = v32
            }
        }
        _ = v29
    } else {
        if v7 {
            if (549 <= 0) {
                v33 := GenS1{F0: (-(v9)), F1: -9223372036854775808}
                if v7 {
                    fmt.Println(v7)
                    fmt.Println(v8, v9)
                }
                v34 := GenS2{F0: false}
                _ = v33
                _ = v34
            }
        }
    }
    v1 += 65535
    if (true && (-716 > 32767)) {
        if (!((0 != (v8 & v8)) || v7)) {
            v35 := v5[uint32(v3)]
            _ = len(v5)
            for fi36 := 0; fi36 < 7; fi36++ {
                v8 = (v8 * (^(v8)))
                if false {
                    if (v7 || v7) {
                        switch fi36 {
                        case 1:
                            for fi37 := 0; fi37 < 6; fi37++ {
                                if (0 != 48) {
                                    v38 := v5[v35]
                                    fmt.Println(v38, fi37)
                                    _ = v38
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v39 int16 = 1000
                                    switch ("foo bar" + (("foo bar" + "hello") + ("hello" + ("a" + "hello")))) {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "a":
                                            var v40 uint16 = 65535
                                            _ = v40
                                        default:
                                            v6 -= 65535
                                        }
                                    }
                                    _ = v39
                                }
                                _ = fi37
                            }
                        case 330:
                            for fi41 := 0; fi41 < 8; fi41++ {
                                if true {
                                    _ = len(v2)
                                    var v42 int = fi36
                                    var v43 string = "foo bar"
                                    switch (v43 + ("" + ("a" + "\n\t"))) {
                                    case "hello":
                                        var v44 uint = 255
                                        _ = v44
                                    default:
                                        var v45 bool = (!v7)
                                        _ = v45
                                    }
                                    if (v7 && v7) {
                                        if true {
                                            v46 := GenS1{F0: (-(v9)), F1: v23}
                                            var v47 int64 = 9223372036854775807
                                            delete(v5, v35)
                                            _ = v46
                                            _ = v47
                                        }
                                        if v7 {
                                            if false {
                                                _ = len(v2)
                                                var v48 string = v43
                                                v5[v1] = (v6 >> ((255) & 63))
                                                _ = v48
                                            }
                                            switch v43 {
                                            case "hello":
                                                v9 = (v9 + v9)
                                            }
                                        } else {
                                            v49 := GenS2{F0: (false && v7)}
                                            switch v42 {
                                            case 0:
                                                v50 := GenS3{F0: (v8 << ((18446744073709551615) & 63)), F1: 944, F2: (v43 > (v43 + v43))}
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                    }
                                    _ = v42
                                    _ = v43
                                }
                                switch 32767 {
                                case 2147483647:
                                    v5[(v6 >> ((4294967295) & 63))] = v35
                                default:
                                    switch -867 {
                                    case -9223372036854775808:
                                        if v7 {
                                            fmt.Println(fi36, fi41)
                                            var v51 int = fi36
                                            switch "" {
                                            case "":
                                                v51 = (fi36 ^ v51)
                                                var v52 uint16 = (v3 + 0)
                                                if true {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v53 := GenS2{F0: false}
                                                        _ = v53
                                                    case "":
                                                        var v54 uint64 = v8
                                                        _ = v54
                                                    }
                                                    var v55 string = ("\n\t" + "")
                                                    var v56 uint16 = 0
                                                    if false {
                                                        v7 = (v7 || false)
                                                        v5[(v6 >> ((0) & 63))] = 0
                                                        fmt.Println(fi36, v56)
                                                        v57 := GenS4{F0: (v35 + 2147483647), F1: v4}
                                                        _ = v57
                                                    } else {
                                                        v58 := GenS5{F0: 128, F1: 0, F2: v4}
                                                        _ = v58
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                }
                                                _ = v52
                                            case "foo bar":
                                                delete(v5, v6)
                                            case "\n\t":
                                                var v59 float32 = (3.14159 * 1e10)
                                                _ = v59
                                            }
                                            _ = v51
                                        }
                                        var v60 uint = 319
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v60
                                    }
                                }
                                _ = fi41
                            }
                        case 888:
                            var v61 string = "hello"
                            _ = v61
                        }
                    }
                } else {
                    v62 := GenS6{F0: 0}
                    _ = v62
                }
                _ = fi36
            }
            _ = v35
        }
    } else {
        for fi63 := 0; fi63 < 10; fi63++ {
            if (int16(v4) > -622) {
                switch "a" {
                case "foo bar":
                    v6 += v6
                }
                if v7 {
                    v9 -= v9
                } else {
                    v64 := GenS1{F0: 1, F1: 42}
                    _ = v64
                }
                for fi65 := 0; fi65 < 15; fi65++ {
                    v9 = v9
                    v66 := GenS3{F0: (v8 - 4294967295), F1: 18446744073709551615, F2: (v7 && (float64((float32(3.14159) - 0.1)) > 0.0))}
                    fmt.Println(v9, v3)
                    _ = fi65
                    _ = v66
                }
            }
            _ = fi63
        }
        v67 := GenS6{F0: 42}
        var v68 uint64 = (v8 << ((0) & 63))
        _ = v67
        _ = v68
    }
    v4 = (v4 * 334)
    var v69 int64 = v23
    var v70 uint32 = v1
    v5[(v6 + (v70 + 125))] = v70
    v7 = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v70
    _ = v8
    _ = v9
}
