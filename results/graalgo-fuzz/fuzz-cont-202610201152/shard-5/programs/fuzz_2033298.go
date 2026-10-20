package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int64 }
func genFn0(p0 int32, p1 uint16) int8 {
    if ("a" != ("" + "foo bar")) {
        var v4 uint16 = (p1 ^ 2)
        if (((67 > 2) || (true && (895 < 688))) || ((^(v4)) <= (^(v4)))) {
            for fi5 := 0; fi5 < 3; fi5++ {
                switch "\n\t" {
                case "a":
                    for fi6 := 0; fi6 < 6; fi6++ {
                        v4 = p1
                        var v7 string = "\n\t"
                        _ = fi6
                        _ = v7
                    }
                case "\n\t":
                    for fi8 := 0; fi8 < 14; fi8++ {
                        if (!(!(42 == 42))) {
                            if true {
                                if true {
                                    v9 := GenS1{F0: 255}
                                    _ = v9
                                } else {
                                    var v10 uint32 = 4294967295
                                    _ = v10
                                }
                                v11 := GenS2{F0: 1, F1: 194, F2: 255}
                                v4 += v4
                                _ = v11
                            }
                            fmt.Println(p0)
                            p1 = (^(v4))
                        } else {
                            v12 := GenS1{F0: (p1 << ((65535) & 63))}
                            _ = v12
                        }
                        v4 -= (v4 / ((255) | 1))
                        _ = fi8
                    }
                default:
                    var v13 int64 = -719
                    _ = v13
                }
                var v14 bool = (-9223372036854775808 <= 0)
                _ = fi5
                _ = v14
            }
            var v15 uint32 = 65535
            _ = v15
        } else {
            v16 := GenS3{F0: 16, F1: "hello"}
            v17 := []int{22}
            var v18 uint32 = 1
            _ = v16
            _ = v17
            _ = v18
        }
        _ = v4
    } else {
        switch 66 {
        case 32768:
            v19 := []int{}
            _ = v19
        case 9223372036854775807:
            p0 = -1
        default:
            fmt.Println(p0)
        }
    }
    var v20 int = -1
    var v21 int16 = 417
    v22 := GenS4{F0: 1}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return 127
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]bool)
    var v3 string = "foo bar"
    v3 = "\n\t"
    var v23 int32 = 42
    var v24 uint8 = 224
    var v25 uint8 = v24
    if (!true) {
        if true {
            v3 = v3
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 uint8 = (v24 >> ((331) & 63))
                var v28 int = fi26
                v23 -= v23
                _ = fi26
                _ = v27
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                v30 := GenS1{F0: 65535}
                v31 := []int{4294967295}
                var v32 float32 = 3.14159
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
        delete(v2, 0)
    }
    fmt.Println(v1, v3)
    fmt.Println(v1, v3)
    for fi33 := 0; fi33 < 7; fi33++ {
        switch (v3 + "") {
        case "a":
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS2{F0: 1, F1: -9223372036854775808, F2: 65535}
                v36 := GenS3{F0: (v25 / (((v24 / (((v24 & v25)) | 1))) | 1)), F1: ("hello" + ("\n\t" + "hello"))}
                switch v3 {
                case "":
                    v25 = v24
                case "hello":
                    var v37 float32 = float32(0.1)
                    _ = v37
                default:
                    v38 := GenS2{F0: 255, F1: -560, F2: 2}
                    _ = v38
                }
                _ = fi34
                _ = v35
                _ = v36
            }
        case "foo bar":
            fmt.Println(v1)
        case "hello":
            fmt.Println(v3, v24, v23)
        default:
            switch (-(fi33)) {
            case 4294967295:
                for fi39 := 0; fi39 < 19; fi39++ {
                    var v40 float64 = 0.0
                    delete(v2, 2147483647)
                    var v41 string = ((v3 + "hello") + "\n\t")
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            case 207:
                switch 128 {
                case 32767:
                    var v42 int = fi33
                    _ = v42
                case 678:
                    var v43 string = v3
                    _ = v43
                }
            case 679:
                v44 := []int{fi33, (fi33 ^ fi33), fi33, 32767}
                _ = v44
            default:
                v45 := []int{fi33}
                _ = v45
            }
        }
        _ = fi33
    }
    var v46 uint8 = uint8(v1)
    v47 := []int{173, -9223372036854775808}
    for fi48 := 0; fi48 < 14; fi48++ {
        var v49 bool = false
        _ = fi48
        _ = v49
    }
    if (false && (true && false)) {
        var v50 int16 = -150
        v51 := GenS5{F0: int32(v50), F1: 9223372036854775807}
        _ = v50
        _ = v51
    } else {
        if (-2.25 > -2.25) {
            var v52 uint64 = (v1 * (v1 - (v1 + 2)))
            _ = v52
        }
    }
    var v53 float64 = 3.14159
    if len(v47) >= 2 {
        v47 = v47[1:len(v47)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v46
    _ = v47
    _ = v53
}
