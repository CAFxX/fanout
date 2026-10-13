package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint64 }

func main() {
    var v1 uint = 65535
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    v4 := GenS1{F0: 128}
    fmt.Println(v1)
    v5 := GenS1{F0: 239}
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 string = "foo bar"
                if false {
                    var v9 int = fi7
                    _ = v9
                }
                switch "foo bar" {
                case "\n\t":
                    var v10 int16 = 63
                    _ = v10
                case "a":
                    v8 = v8
                case "foo bar":
                    v11 := GenS2{F0: 0}
                    _ = v11
                }
                v1 = 18446744073709551615
                _ = fi7
                _ = v8
            }
            if ((0.0 - (-2.25 / 1e10)) != 3.14159) {
                if false {
                    v12 := GenS2{F0: 42}
                    var v13 uint32 = 988
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(v1, fi6)
                    fmt.Println(v1, fi6)
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v1 -= v1
                        _ = len(v2)
                        if (-1 >= -32768) {
                            if (false || true) {
                                fmt.Println(v1, v1)
                                var v15 uint = v1
                                _ = v15
                            }
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v16 := GenS3{F0: 324, F1: 0, F2: -804}
                                delete(v3, 2147483648)
                                var v17 int32 = -2147483648
                                _ = v16
                                _ = v17
                            } else {
                                v1 = v1
                            }
                            var v18 float64 = -2.25
                            _ = v18
                        } else {
                            v19 := GenS3{F0: 42, F1: 1, F2: 1}
                            if false {
                                v1 += (v1 + (^(v1)))
                                var v20 uint16 = 337
                                v21 := v3[2]
                                _ = v20
                                _ = v21
                            }
                            var v22 uint8 = 1
                            _ = v19
                            _ = v22
                        }
                        _ = fi14
                    }
                }
                v1 -= 18446744073709551615
                switch (fi6 | fi6) {
                case 9223372036854775807:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case -58:
                    var v23 uint16 = 2
                    _ = v23
                }
            }
        }
        v24 := GenS3{F0: 762, F1: 0, F2: -1}
        v25 := GenS2{F0: -748}
        _ = fi6
        _ = v24
        _ = v25
    }
    if (73 == 255) {
        v26 := GenS1{F0: 127}
        _ = v26
    } else {
        if (32767 <= 593) {
            if false {
                if true {
                    if (!false) {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    } else {
                        if ("foo bar" < "\n\t") {
                            v27 := GenS2{F0: -111}
                            for fi28 := 0; fi28 < 13; fi28++ {
                                fmt.Println(fi28)
                                switch (fi28 & fi28) {
                                case 0:
                                    delete(v3, 615)
                                case 127:
                                    var v29 uint = (v1 & 18446744073709551615)
                                    _ = v29
                                }
                                var v30 int32 = 42
                                _ = fi28
                                _ = v30
                            }
                            _ = v27
                        } else {
                            switch "foo bar" {
                            case "":
                                if (-1 > int16(v1)) {
                                    v31 := GenS1{F0: 2}
                                    _ = v31
                                } else {
                                    var v32 uint64 = 18446744073709551615
                                    _ = v32
                                }
                            case "a":
                                v1 += (v1 * (v1 - v1))
                            case "\n\t":
                                if false {
                                    v1 = (v1 | (v1 << (((v1 / ((v1) | 1))) & 63)))
                                } else {
                                    if false {
                                        var v33 uint32 = 0
                                        var v34 bool = true
                                        _ = v33
                                        _ = v34
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = -956
                                        }
                                    }
                                    switch ("foo bar" + ("\n\t" + "")) {
                                    case "hello":
                                        if (false && (v1 < 4294967295)) {
                                            v2 = append(v2, 42)
                                            fmt.Println(v1)
                                        }
                                    }
                                    var v35 uint32 = 65535
                                    _ = v35
                                }
                            }
                            v1 = 0
                        }
                    }
                    v1 += (v1 | (v1 & 809))
                }
                v36 := GenS1{F0: 254}
                _ = v36
            } else {
                switch -399 {
                case 160:
                    v1 = (v1 * v1)
                }
                switch 127 {
                case 1:
                    v37 := GenS4{F0: 1, F1: 21, F2: 1}
                    _ = v37
                default:
                    var v38 int8 = 1
                    _ = v38
                }
                for fi39 := 0; fi39 < 6; fi39++ {
                    v1 += 4294967295
                    var v40 uint32 = 1000
                    v41 := GenS2{F0: int16(v40)}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            }
        } else {
            for fi42 := 0; fi42 < 12; fi42++ {
                v1 += 42
                if (!(42 <= -9223372036854775808)) {
                    v1 = v1
                } else {
                    v43 := v3[487]
                    for fi44 := 0; fi44 < 14; fi44++ {
                        if ((-104 < -128) || (!(true || false))) {
                            v1 = v1
                        }
                        _ = fi44
                    }
                    var v45 uint32 = v43
                    _ = v43
                    _ = v45
                }
                v1 -= v1
                _ = fi42
            }
        }
        var v46 uint8 = 127
        var v47 uint32 = 0
        _ = v46
        _ = v47
    }
    var v48 uint16 = 966
    fmt.Println(v1, v1)
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v49 bool = (2147483647 != 32767)
    v48 = (v48 | v48)
    v48 -= (^(v48))
    for fi50 := 0; fi50 < 19; fi50++ {
        var v51 float64 = 1.5
        fmt.Println(v48, v51, v51)
        switch fi50 {
        case 4294967295:
            v1 += v1
        case 32767:
            var v52 float64 = v51
            _ = v52
        default:
            var v53 uint = (v1 << (((v1 >> (((v1 & v1)) & 63))) & 63))
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    v54 := GenS1{F0: 2}
    if false {
        fmt.Println(v1)
        fmt.Println(v1, v48, v1)
        v49 = true
    }
    var v55 uint16 = (v48 | v48)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v55
}
