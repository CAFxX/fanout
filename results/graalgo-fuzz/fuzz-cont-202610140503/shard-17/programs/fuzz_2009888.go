package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []int{127, 127}
    v2 := []uint32{743, 428, 1, 648}
    var v3 uint32 = 2147483647
    var v4 float32 = float32(3.14159)
    var v5 bool = ("hello" <= ("" + ("" + "\n\t")))
    switch 127 {
    case 128:
        fmt.Println(v5, v3)
    case 62:
        if false {
            var v6 int = 4294967295
            var v7 bool = false
            if (!false) {
                v6 -= (v6 + -1)
            } else {
                if true {
                    var v8 int32 = 584
                    _ = v8
                } else {
                    var v9 uint8 = 81
                    _ = v9
                }
            }
            _ = v6
            _ = v7
        }
    case 4294967295:
        var v10 int64 = -9223372036854775808
        _ = v10
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int32 = 2147483647
            _ = fi11
            _ = v12
        }
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        v1 = append(v1, 9223372036854775807)
    default:
        switch -1 {
        case 32768:
            if (1.5 <= (((-2.25 * 0.1) + (((float64(float32(0.1)) / float64(1e10)) / (3.14159 * ((-2.25 / 0.1) - float64(0.0)))) - (0.0 + 0.1))) + 1e10)) {
                v13 := GenS1{F0: -2147483648}
                if (!true) {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            if (v5 && (v5 || (fi14 > int(v3)))) {
                                v3 = (v3 >> ((18446744073709551615) & 63))
                            }
                            v16 := GenS1{F0: 4294967295}
                            _ = fi15
                            _ = v16
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi14
                    }
                    var v17 int16 = -1
                    _ = v17
                }
                v3 += 650
                _ = v13
            }
        case 127:
            for fi18 := 0; fi18 < 20; fi18++ {
                if len(v2) > 0 {
                    v2[0] = v3
                }
                var v19 bool = v5
                _ = fi18
                _ = v19
            }
        case 2147483647:
            switch 105 {
            case 824:
                var v20 uint32 = v3
                _ = v20
            case -252:
                if (-2147483648 < -191) {
                    v4 = v4
                    v21 := GenS1{F0: int(v3)}
                    _ = v21
                }
            }
        default:
            var v22 int32 = 1
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v5, fi23, v5)
        v24 := GenS1{F0: (fi23 ^ 127)}
        _ = fi23
        _ = v24
    }
    switch 9223372036854775807 {
    case -286:
        v2 = append(v2, 2147483647)
    case -643:
        switch 128 {
        case -9223372036854775808:
            v4 -= v4
        case 32767:
            switch "foo bar" {
            case "hello":
                if ((!true) || v5) {
                    for fi25 := 0; fi25 < 13; fi25++ {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            var v27 int64 = 0
                            var v28 uint8 = 255
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v4 = v4
                        if v5 {
                            fmt.Println(v5, v3, v3)
                            if true {
                                var v29 string = "hello"
                                _ = v29
                            } else {
                                v4 = v4
                            }
                            v30 := GenS2{F0: -1}
                            _ = v30
                        } else {
                            v4 = 3.14159
                            switch 42 {
                            case 1:
                                var v31 int64 = int64(v3)
                                _ = v31
                            default:
                                v32 := GenS2{F0: 1}
                                _ = v32
                            }
                            v33 := GenS2{F0: 42}
                            _ = v33
                        }
                        _ = fi25
                    }
                }
            case "foo bar":
                var v34 bool = (!v5)
                _ = v34
            case "a":
                if ((1 != 42) || v5) {
                    var v35 uint16 = 255
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v3, v5)
                    default:
                        v3 = (v3 >> ((2) & 63))
                    }
                    _ = v35
                }
            default:
                v1 = append(v1, -9223372036854775808)
            }
        case 2147483647:
            for fi36 := 0; fi36 < 16; fi36++ {
                if len(v2) > 0 {
                    v2[0] = (^(v3))
                }
                v3 += v3
                _ = fi36
            }
        }
    default:
        v37 := GenS3{F0: 42}
        _ = v37
    }
    switch ("" + "hello") {
    case "":
        v4 += ((0.1 / ((v4 - v4) * v4)) * v4)
    case "\n\t":
        v3 = v3
    case "hello":
        var v38 string = ("" + ("\n\t" + "a"))
        _ = v38
    default:
        var v39 int = 189
        _ = v39
    }
    v40 := GenS4{F0: -2147483648, F1: ""}
    if (3.14159 < 1e10) {
        var v41 uint64 = 879
        v42 := GenS5{F0: 42}
        switch -9223372036854775808 {
        case 127:
            var v43 int64 = 154
            _ = v43
        }
        _ = v41
        _ = v42
    } else {
        v5 = v5
        v3 -= 0
    }
    var v44 int64 = -793
    for fi45 := 0; fi45 < 16; fi45++ {
        var v46 int16 = 32767
        v4 += float32(float64(((float32(3.14159) / 1.5) - -2.25)))
        _ = fi45
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v5
}
