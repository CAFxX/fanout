package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint64; F1 int64 }
func genFn0(p0 int8) uint64 {
    v4 := []int{258, -94}
    var v5 uint = 1
    v5 += (^(v5))
    _ = p0
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 int = -97
    v3 := make(map[uint32]bool)
    v6 := []int{(v2 + (v2 % ((v2) | 1))), v2}
    v2 = (v2 % ((32768) | 1))
    var v7 int16 = 42
    var v8 uint64 = 2
    var v9 bool = ((!true) || false)
    v7 = (v7 / ((-32768) | 1))
    if (42 < 2) {
        if (!v9) {
            v10 := GenS1{F0: 0}
            switch (v2 ^ 127) {
            case 880:
                switch 4294967295 {
                case -1:
                    var v11 uint64 = v8
                    _ = v11
                }
            case -2147483648:
                if ((v1 / (3.14159 * (((0.1 / (((((v1 / v1) / v1) / v1) / (v1 + float64((((0.1 + 0.1) - 3.14159) / 3.14159)))) + 1.5)) * 1.5) + 3.14159))) <= (0.1 * v1)) {
                    v12 := GenS1{F0: 128}
                    var v13 int = v2
                    var v14 int = -910
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 128}
                    var v16 float32 = 0.1
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        var v17 int16 = 1000
                        _ = v17
                    }
                    _ = v15
                    _ = v16
                }
            default:
                fmt.Println(v1, v7, v1)
            }
            if v9 {
                v1 = (0.1 * (((v1 / (1e10 * (0.1 + 3.14159))) - 3.14159) + 0.1))
                var v18 uint = 255
                _ = v18
            } else {
                v7 += v7
                v9 = ((v7 % (((v7 << ((42) & 63))) | 1)) == (v7 % (((v7 + 32767)) | 1)))
                var v19 float32 = 1e10
                _ = v19
            }
            _ = v10
        } else {
            var v20 int32 = -1
            fmt.Println(v20)
            _ = v20
        }
        fmt.Println(v9, v7, v7)
    }
    if v9 {
        switch v2 {
        case 9223372036854775807:
            var v21 bool = (1 != 42)
            _ = v21
        case 4294967295:
            v1 += v1
        }
        var v22 bool = (!(v9 && v9))
        _ = v22
    } else {
        var v23 int16 = v7
        if (926 == 0) {
            v24 := GenS2{F0: 32767}
            for fi25 := 0; fi25 < 16; fi25++ {
                v8 = 18446744073709551615
                _ = fi25
            }
            _ = v24
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 float32 = 1.5
                for fi28 := 0; fi28 < 15; fi28++ {
                    v29 := GenS2{F0: 0}
                    v23 = v23
                    _ = fi28
                    _ = v29
                }
                fmt.Println(fi26, v2)
                _ = fi26
                _ = v27
            }
            v23 = (-(v7))
            var v30 int8 = 1
            _ = v30
        }
        _ = v23
    }
    if (2 <= 971) {
        var v31 uint16 = 2
        for fi32 := 0; fi32 < 20; fi32++ {
            for fi33 := 0; fi33 < 17; fi33++ {
                var v34 uint64 = genFn0(0)
                _ = fi33
                _ = v34
            }
            if ((v7 | (-(v7))) > v7) {
                var v35 float32 = (-2.25 * ((-2.25 / 1e10) / 3.14159))
                if (2147483647 >= 4294967295) {
                    if (v31 == 65535) {
                        if false {
                            v9 = (1 == (v7 ^ v7))
                            if v9 {
                                var v36 int8 = -1
                                _ = v36
                            } else {
                                fmt.Println(v2)
                                v9 = (!(!("foo bar" <= (("foo bar" + "hello") + ""))))
                            }
                        } else {
                            v9 = true
                            switch "a" {
                            case "\n\t":
                                switch (v2 - (v2 >> ((18446744073709551615) & 63))) {
                                case 817:
                                    v9 = v9
                                case 2147483647:
                                    switch v2 {
                                    case 1:
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            v37 := GenS2{F0: (v31 / ((v31) | 1))}
                                            _ = v37
                                        default:
                                            var v38 float32 = v35
                                            _ = v38
                                        }
                                    case -412:
                                        fmt.Println(fi32, v2)
                                    }
                                case 1:
                                    v39 := GenS1{F0: 127}
                                    _ = v39
                                default:
                                    var v40 uint32 = 65535
                                    _ = v40
                                }
                            case "":
                                var v41 uint16 = (^(v31))
                                _ = v41
                            case "foo bar":
                                for fi42 := 0; fi42 < 3; fi42++ {
                                    switch "\n\t" {
                                    case "a":
                                        if (v9 && v9) {
                                            var v43 uint16 = v31
                                            fmt.Println(v8, v9, v8)
                                            v44 := GenS3{F0: 13, F1: 42}
                                            if ((0 == 42) && (!v9)) {
                                                var v45 bool = v9
                                                _ = v45
                                            } else {
                                                v31 -= v43
                                            }
                                            if v9 {
                                                _ = len(v3)
                                                v8 = genFn0(127)
                                            } else {
                                                if ((^(v8)) >= v8) {
                                                    var v46 uint32 = 0
                                                    _ = v46
                                                } else {
                                                    v47 := v3[1]
                                                    v8 = genFn0(100)
                                                    _ = v47
                                                }
                                            }
                                            _ = v43
                                            _ = v44
                                        }
                                    default:
                                        v35 += v35
                                    }
                                    if v9 {
                                        if (42 != 127) {
                                            switch (^(v2)) {
                                            case 0:
                                                fmt.Println(v31)
                                            case 43:
                                                fmt.Println(v7)
                                            }
                                            _ = len(v6)
                                            var v48 int32 = -1
                                            v49 := GenS3{F0: 2, F1: -1}
                                            v2 += (fi42 >> ((1) & 63))
                                            _ = v48
                                            _ = v49
                                        }
                                        v31 = 0
                                    }
                                    _ = fi42
                                }
                            }
                            v3[2147483647] = (!((v9 && false) && (v8 >= (^(v8)))))
                        }
                    }
                    if ((true && (v8 > v8)) || v9) {
                        if ((v7 * (v7 & v7)) != (v7 * -606)) {
                            v50 := GenS4{F0: 2, F1: -497}
                            _ = v50
                        } else {
                            delete(v3, uint32(v31))
                        }
                        v3[2] = v9
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    v9 = (!v9)
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 int64 = 9223372036854775807
                        var v53 int32 = -1
                        v8 += (v8 / ((839) | 1))
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 uint8 = 1
                    _ = v54
                }
                var v55 float32 = v35
                _ = v35
                _ = v55
            }
            _ = fi32
        }
        _ = v31
    }
    var v56 string = "a"
    v1 = (float64((1.5 * -2.25)) + v1)
    var v57 uint32 = 2147483647
    var v58 uint64 = 877
    var v59 uint64 = (^(v58))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
