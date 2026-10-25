package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 int64) uint16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(p0, p0)
        var v7 float32 = 1.5
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (p0 << ((65535) & 63)), F1: 0}
    _ = p0
    _ = v8
    return 255
}

func main() {
    var v1 float64 = (-2.25 + (0.1 + 0.0))
    var v2 float64 = (0.0 - v1)
    var v3 uint8 = 192
    v4 := []string{"", "foo bar", ""}
    v5 := make(map[uint32]uint32)
    switch "a" {
    case "\n\t":
        for fi9 := 0; fi9 < 11; fi9++ {
            if true {
                for fi10 := 0; fi10 < 4; fi10++ {
                    _ = len(v4)
                    v11 := GenS2{F0: 860, F1: (("foo bar" + ("a" + "")) + ("hello" + "\n\t"))}
                    _ = fi10
                    _ = v11
                }
                var v12 uint16 = 1
                var v13 int32 = -957
                _ = v12
                _ = v13
            }
            if (651 != 228) {
                var v14 int32 = 1
                if (69 == 4294967295) {
                    if ((-(fi9)) == 127) {
                        if false {
                            v1 += 1.5
                            v4 = append(v4, "hello")
                            var v15 int8 = -1
                            _ = v15
                        } else {
                            var v16 uint8 = 127
                            _ = v16
                        }
                        var v17 bool = true
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v17
                    } else {
                        var v18 int8 = 0
                        v19 := GenS1{F0: 9223372036854775807, F1: -921}
                        switch -812 {
                        case 241:
                            var v20 int8 = -47
                            _ = v20
                        default:
                            v21 := GenS2{F0: 42, F1: "foo bar"}
                            _ = v21
                        }
                        _ = v18
                        _ = v19
                    }
                } else {
                    var v22 uint8 = (^(v3))
                    _ = v22
                }
                _ = v14
            } else {
                var v23 int = fi9
                var v24 float32 = float32(v1)
                v3 = 1
                _ = v23
                _ = v24
            }
            _ = fi9
        }
    case "hello":
        delete(v5, 670)
    default:
        var v25 int8 = -50
        _ = v25
    }
    switch "foo bar" {
    case "\n\t":
        var v26 int8 = 127
        _ = v26
    case "hello":
        if ("a" < "a") {
            v1 += v1
            var v27 uint64 = 2
            _ = v27
        }
    }
    var v28 uint64 = 2
    fmt.Println(v2, v2)
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(fi29, v3)
        for fi30 := 0; fi30 < 17; fi30++ {
            v2 -= v1
            _ = fi30
        }
        _ = fi29
    }
    switch 952 {
    case 32768:
        fmt.Println(v1, v28, v1)
    case 0:
        for fi31 := 0; fi31 < 16; fi31++ {
            if ((((v28 & uint64(v3)) != (v28 >> ((712) & 63))) || (!false)) || (((65535 < 803) && true) && false)) {
                v32 := GenS2{F0: v28, F1: ""}
                for fi33 := 0; fi33 < 17; fi33++ {
                    var v34 int64 = 42
                    if true {
                        var v35 float32 = (0.1 - 1.5)
                        switch ("\n\t" + "") {
                        case "\n\t":
                            v1 += float64((v35 + 1.5))
                        case "foo bar":
                            v28 -= 0
                        case "":
                            fmt.Println(v2, v35, v2)
                        default:
                            var v36 int16 = 0
                            var v37 int = fi33
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS3{F0: uint16(v3), F1: 250, F2: 555}
                        var v39 uint = 2
                        _ = v35
                        _ = v38
                        _ = v39
                    } else {
                        v28 += v28
                        if true {
                            switch "hello" {
                            case "foo bar":
                                var v40 int8 = -1
                                _ = v40
                            }
                        } else {
                            var v41 int = fi33
                            _ = v41
                        }
                        _ = len(v4)
                        var v42 int32 = -1
                        _ = v42
                    }
                    fmt.Println(fi31, fi33, v28)
                    _ = fi33
                    _ = v34
                }
                fmt.Println(v2, v1, v2)
                _ = v32
            } else {
                for fi43 := 0; fi43 < 13; fi43++ {
                    v3 -= 246
                    _ = fi43
                }
                v44 := GenS3{F0: 1, F1: 2147483648, F2: 729}
                _ = v44
            }
            if true {
                v2 -= v1
                var v45 float64 = v2
                v1 += ((v2 + (0.0 + 1.5)) + v2)
                _ = v45
            } else {
                v2 = 0.0
                for fi46 := 0; fi46 < 16; fi46++ {
                    _ = len(v5)
                    var v47 string = ("hello" + "a")
                    if (!(0.1 == 1.5)) {
                        var v48 float64 = v1
                        var v49 bool = false
                        v1 += (1e10 + v2)
                        v5[42] = 476
                        _ = v48
                        _ = v49
                    }
                    var v50 float64 = v1
                    _ = fi46
                    _ = v47
                    _ = v50
                }
                v1 += -2.25
            }
            _ = fi31
        }
    }
    v5[4294967295] = 289
    fmt.Println(v2, v2)
    fmt.Println(v28, v1)
    for fi51 := 0; fi51 < 5; fi51++ {
        _ = len(v4)
        _ = fi51
    }
    v28 -= v28
    switch ("" + ("\n\t" + "\n\t")) {
    case "":
        switch "a" {
        case "hello":
            if (42 >= 1) {
                v28 = 0
                var v52 int16 = 1
                _ = v52
            } else {
                var v53 int8 = 42
                v1 = (v1 - v2)
                fmt.Println(v28, v1)
                _ = v53
            }
        }
    case "hello":
        for fi54 := 0; fi54 < 10; fi54++ {
            if false {
                _ = len(v4)
                switch 4294967295 {
                case -1:
                    v3 = (^(v3))
                case 32767:
                    if false {
                        var v55 uint16 = 2
                        v56 := GenS4{F0: (v55 >> ((14) & 63)), F1: 42}
                        _ = v55
                        _ = v56
                    } else {
                        v57 := GenS4{F0: 1, F1: (-(fi54))}
                        _ = v57
                    }
                default:
                    for fi58 := 0; fi58 < 17; fi58++ {
                        if false {
                            fmt.Println(v1, v1, fi58)
                            v1 += float64((1.5 - float32((((v1 * -2.25) - 3.14159) + v1))))
                        }
                        v28 -= 18446744073709551615
                        _ = fi58
                    }
                }
            }
            _ = fi54
        }
    case "a":
        fmt.Println(v2, v1)
    default:
        for fi59 := 0; fi59 < 11; fi59++ {
            fmt.Println(v2, v3, fi59)
            _ = fi59
        }
    }
    var v60 uint64 = v28
    fmt.Println(v3)
    if len(v4) > 0 {
        v4[0] = (("" + "") + ("hello" + ("hello" + "a")))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
