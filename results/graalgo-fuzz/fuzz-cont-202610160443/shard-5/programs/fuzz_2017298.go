package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16 }
func genFn0() int8 {
    if (0 < -1) {
        if true {
            if false {
                switch -9223372036854775808 {
                case 630:
                    var v6 uint16 = 42
                    _ = v6
                case 289:
                    if true {
                        var v7 float32 = 1.5
                        var v8 int8 = -1
                        switch 127 {
                        case -9223372036854775808:
                            var v9 uint64 = 1
                            _ = v9
                        default:
                            v8 += v8
                        }
                        _ = v7
                        _ = v8
                    } else {
                        v10 := GenS1{F0: 747}
                        var v11 int8 = 1
                        _ = v10
                        _ = v11
                    }
                case -9223372036854775808:
                    if false {
                        var v12 uint8 = 128
                        fmt.Println(v12)
                        _ = v12
                    }
                }
            }
            if (true && (false || (!false))) {
                var v13 uint64 = 0
                _ = v13
            }
        }
        if false {
            if false {
                var v14 uint16 = 304
                fmt.Println(v14)
                _ = v14
            } else {
                var v15 int16 = 1
                _ = v15
            }
        }
        var v16 int = -2147483648
        _ = v16
    } else {
        if (-2147483648 != 0) {
            var v17 float64 = 1.5
            _ = v17
        } else {
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := []int{fi18, (fi18 + (fi18 << ((536) & 63))), (fi18 / (((fi18 ^ (^(fi18)))) | 1))}
                var v20 int16 = -32768
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 int64 = -9223372036854775808
            fmt.Println(v21)
            _ = v21
        }
    }
    v22 := GenS1{F0: 264}
    var v23 int = 1
    if (32768 == (v23 ^ (v23 << ((42) & 63)))) {
        v23 = (v23 * v23)
    } else {
        var v24 uint32 = 2147483648
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        for fi26 := 0; fi26 < 10; fi26++ {
            v23 -= fi25
            v27 := []int{-388, fi25, fi25}
            _ = fi26
            _ = v27
        }
        fmt.Println(v23)
        _ = fi25
    }
    _ = v22
    _ = v23
    return 1
}

func main() {
    v1 := []int{128, 127}
    var v2 uint = 4294967295
    var v3 float64 = (1.5 + 1e10)
    v4 := []uint32{2147483648, 2, 4294967295, 961}
    var v5 float64 = (((v3 + v3) + 3.14159) * -2.25)
    v28 := GenS2{F0: (false && false), F1: true, F2: true}
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 uint16 = 1
        v30 -= v30
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    var v31 uint16 = 65535
    v5 = float64(0.1)
    v32 := GenS1{F0: 1}
    if true {
        for fi33 := 0; fi33 < 10; fi33++ {
            fmt.Println(fi33, v5)
            var v34 int64 = 9223372036854775807
            _ = fi33
            _ = v34
        }
        switch ("" + (("\n\t" + ("hello" + "\n\t")) + ("" + ""))) {
        case "foo bar":
            v2 += v2
        case "":
            for fi35 := 0; fi35 < 12; fi35++ {
                fmt.Println(fi35, v5)
                v5 -= (v5 + 0.1)
                _ = fi35
            }
        default:
            var v36 int8 = -75
            _ = v36
        }
    } else {
        if (true && ((false || true) || (-708 >= 1000))) {
            v37 := GenS3{F0: 127, F1: (^(v31)), F2: 0}
            if (1 != 848) {
                if (!true) {
                    v2 += v2
                    fmt.Println(v5)
                }
                for fi38 := 0; fi38 < 14; fi38++ {
                    v2 = (v2 >> ((v2) & 63))
                    for fi39 := 0; fi39 < 4; fi39++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if false {
                            v2 -= (v2 / (((v2 ^ v2)) | 1))
                        }
                        _ = fi39
                    }
                    var v40 float32 = ((3.14159 + ((1.5 + (-2.25 - 1e10)) - ((((0.1 * (((3.14159 - (((0.0 + (float32(3.14159) * (-2.25 - 1.5))) + 3.14159) - 0.0)) - 3.14159) * 0.0)) * -2.25) / float32(1.5)) / 3.14159))) + 1.5)
                    _ = fi38
                    _ = v40
                }
                if true {
                    fmt.Println(v3)
                }
            } else {
                v5 += v5
            }
            _ = v37
        } else {
            switch 731 {
            case -118:
                v2 = v2
            case 128:
                var v41 int64 = -1
                _ = v41
            }
            var v42 int64 = 42
            if (v42 < (v42 & (v42 / ((v42) | 1)))) {
                for fi43 := 0; fi43 < 6; fi43++ {
                    fmt.Println(v2, v2)
                    fmt.Println(v5)
                    v5 = v3
                    _ = fi43
                }
                fmt.Println(v5)
                v31 -= (v31 / ((376) | 1))
            }
            _ = v42
        }
    }
    v31 += (v31 << (((^(v2))) & 63))
    v44 := GenS3{F0: int(v31), F1: 581, F2: 1}
    var v45 int = 32767
    for fi46 := 0; fi46 < 16; fi46++ {
        if (!true) {
            for fi47 := 0; fi47 < 5; fi47++ {
                v48 := GenS3{F0: v45, F1: (v31 + 32767), F2: genFn0()}
                v3 = 3.14159
                v45 = v45
                _ = fi47
                _ = v48
            }
            var v49 string = ("a" + "hello")
            _ = v49
        } else {
            v4 = append(v4, 42)
            switch (v45 / (((v45 % ((v45) | 1))) | 1)) {
            case 128:
                for fi50 := 0; fi50 < 6; fi50++ {
                    switch (-(fi46)) {
                    case 1:
                        var v51 uint32 = uint32(v31)
                        _ = v51
                    case 128:
                        v52 := GenS3{F0: (v45 % ((4294967295) | 1)), F1: 114, F2: -1}
                        _ = v52
                    case 4294967295:
                        v53 := GenS4{F0: v31}
                        _ = v53
                    }
                    _ = fi50
                }
            case 423:
                if (2147483647 >= 42) {
                    switch (("" + "\n\t") + "hello") {
                    case "a":
                        switch "a" {
                        case "foo bar":
                            if (true && (1 >= 291)) {
                                v2 -= v2
                            } else {
                                fmt.Println(v31, v45)
                                v3 = 0.0
                                var v54 uint8 = 37
                                _ = v54
                            }
                        case "\n\t":
                            var v55 float64 = v5
                            _ = v55
                        }
                    default:
                        var v56 uint8 = 3
                        _ = v56
                    }
                } else {
                    var v57 int32 = -2147483648
                    var v58 uint32 = 1
                    var v59 bool = true
                    _ = v57
                    _ = v58
                    _ = v59
                }
            }
            var v60 int16 = -183
            _ = v60
        }
        var v61 int64 = -9223372036854775808
        _ = fi46
        _ = v61
    }
    if (255 <= 71) {
        v45 = 4294967295
        v2 = v2
    }
    if len(v1) > 0 {
        v1[0] = v45
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v62 uint8 = 6
    switch int(v62) {
    case 127:
        if false {
            for fi63 := 0; fi63 < 14; fi63++ {
                var v64 uint = 428
                _ = fi63
                _ = v64
            }
        }
    case 2147483647:
        var v65 uint8 = (^(v62))
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v62
}
