package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint16) int8 {
    var v6 uint = 540
    v6 += (v6 & v6)
    _ = p0
    _ = v6
    return 107
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("a" + "foo bar")
    v3 := []uint32{1, 2147483647}
    var v4 uint32 = 622
    v5 := make(map[int]string)
    var v7 string = "hello"
    var v8 uint64 = 2
    var v9 float32 = (-2.25 / 3.14159)
    fmt.Println(v2, v4)
    var v10 uint = 1
    if (!false) {
        for fi11 := 0; fi11 < 11; fi11++ {
            v4 += v4
            _ = fi11
        }
        if false {
            v3 = append(v3, (v4 | 2147483647))
            var v12 int32 = 42
            _ = v12
        }
        if false {
            switch -999 {
            case -9223372036854775808:
                for fi13 := 0; fi13 < 12; fi13++ {
                    if (false && (0 < 128)) {
                        if (-9223372036854775808 != 1) {
                            var v14 uint64 = v8
                            v2 = "hello"
                            _ = v14
                        } else {
                            v5[-971] = ("\n\t" + (v7 + v7))
                        }
                        if (127 == (fi13 / ((-9223372036854775808) | 1))) {
                            v15 := GenS1{F0: (fi13 + (fi13 & (fi13 >> (((v10 ^ v10)) & 63)))), F1: false, F2: -9223372036854775808}
                            v8 = (v8 % ((v8) | 1))
                            switch 42 {
                            case 4294967295:
                                fmt.Println(v4, fi13)
                            case -412:
                                if (!((genFn0(701) > -93) && ((2 >= 152) || true))) {
                                    v9 -= v9
                                    for fi16 := 0; fi16 < 4; fi16++ {
                                        var v17 string = v7
                                        v10 = v10
                                        var v18 uint32 = (v4 ^ v4)
                                        _ = fi16
                                        _ = v17
                                        _ = v18
                                    }
                                    v19 := GenS2{F0: -463, F1: 0}
                                    _ = v19
                                }
                            case -1:
                                if (v9 != (v9 - float32(1e10))) {
                                    var v20 uint16 = 0
                                    var v21 string = v7
                                    _ = v20
                                    _ = v21
                                }
                            }
                            _ = v15
                        } else {
                            var v22 int64 = 9223372036854775807
                            switch (v2 + v2) {
                            case "":
                                v23 := v5[(fi13 << ((v10) & 63))]
                                _ = v23
                            case "hello":
                                if (v2 != v2) {
                                    if (false || (-138 >= 1000)) {
                                        var v24 int8 = -128
                                        v25 := GenS2{F0: -1, F1: 2}
                                        if false {
                                            var v26 uint16 = 478
                                            _ = v26
                                        } else {
                                            if true {
                                                for fi27 := 0; fi27 < 7; fi27++ {
                                                    switch "a" {
                                                    case "":
                                                        v3 = append(v3, 0)
                                                    default:
                                                        var v28 int32 = 42
                                                        _ = v28
                                                    }
                                                    var v29 uint = v10
                                                    var v30 uint8 = 128
                                                    _ = fi27
                                                    _ = v29
                                                    _ = v30
                                                }
                                                fmt.Println(v10)
                                                v22 = (v22 | -9223372036854775808)
                                            } else {
                                                for fi31 := 0; fi31 < 5; fi31++ {
                                                    var v32 int8 = 1
                                                    _ = fi31
                                                    _ = v32
                                                }
                                                if false {
                                                    fmt.Println(v4)
                                                    fmt.Println(v7, v7, v9)
                                                }
                                                v24 += (v24 % (((v24 + 63)) | 1))
                                            }
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        for fi33 := 0; fi33 < 8; fi33++ {
                                            v34 := GenS1{F0: (fi13 << ((v10) & 63)), F1: (int(v8) <= (fi13 % ((fi33) | 1))), F2: v22}
                                            _ = fi33
                                            _ = v34
                                        }
                                        var v35 float64 = 0.0
                                        _ = v35
                                    }
                                }
                            }
                            _ = v22
                        }
                        var v36 string = v7
                        _ = v36
                    }
                    var v37 uint16 = 1
                    _ = fi13
                    _ = v37
                }
            case 128:
                var v38 float64 = -2.25
                _ = v38
            default:
                v10 -= (^(v10))
            }
            for fi39 := 0; fi39 < 1; fi39++ {
                switch v2 {
                case "\n\t":
                    v40 := v5[fi39]
                    _ = v40
                }
                v41 := GenS3{F0: 491, F1: -798, F2: 0}
                _ = fi39
                _ = v41
            }
            v42 := GenS4{F0: 2, F1: 42}
            _ = v42
        } else {
            v2 = "\n\t"
            fmt.Println(v10)
        }
    }
    v7 = ("foo bar" + "a")
    v43 := GenS3{F0: 1, F1: -378, F2: -487}
    var v44 float64 = (1e10 * 1.5)
    if len(v3) > 0 {
        v3[0] = v4
    }
    v5[282] = v7
    var v45 int64 = -9223372036854775808
    switch (v2 + "") {
    case "":
        for fi46 := 0; fi46 < 4; fi46++ {
            fmt.Println(v44, v44, v7)
            for fi47 := 0; fi47 < 5; fi47++ {
                switch int(v10) {
                case 32768:
                    var v48 float64 = v44
                    _ = v48
                case 4294967295:
                    fmt.Println(v4)
                }
                _ = fi47
            }
            v45 = (v45 ^ (^(v45)))
            _ = fi46
        }
    }
    fmt.Println(v4)
    var v49 uint = v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v49
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
