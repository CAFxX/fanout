package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint32 = 2147483647
        _ = fi7
        _ = v8
    }
    var v9 int64 = 247
    v10 := []int{32768, 4294967295, 32768, 463}
    _ = p0
    _ = v10
    _ = v9
    return p0
}
func genFn1(p0 int8) uint8 {
    var v11 int32 = -2147483648
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int = fi12
        p0 = (p0 + (p0 - -128))
        _ = fi12
        _ = v13
    }
    var v14 int = 4294967295
    v15 := []int{-243}
    _ = p0
    _ = v11
    _ = v14
    _ = v15
    return 255
}

func main() {
    var v1 int16 = 219
    var v2 uint16 = 2
    var v3 int32 = 2147483647
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 49
    var v6 uint64 = uint64(v2)
    var v16 uint16 = (^(v2))
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            fmt.Println(v3, v3, v5)
            _ = fi18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            if false {
                if false {
                    v20 := GenS1{F0: (v16 % ((v2) | 1)), F1: v5, F2: 18446744073709551615}
                    var v21 int32 = 42
                    _ = v20
                    _ = v21
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            var v22 int64 = -790
                            var v23 bool = (3.14159 < -2.25)
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                fmt.Println(v5, v23, v23)
                                v24 := GenS1{F0: 2, F1: v5, F2: 2}
                                _ = v24
                            case "foo bar":
                                v25 := []int{(fi17 / ((fi17) | 1))}
                                _ = v25
                            case "hello":
                                var v26 bool = (v23 && v23)
                                _ = v26
                            }
                            var v27 float64 = float64(1e10)
                            _ = v22
                            _ = v23
                            _ = v27
                        } else {
                            v5 = 1
                            switch ("a" + "") {
                            case "\n\t":
                                v1 -= 790
                            }
                        }
                    case "hello":
                        var v28 int32 = (v3 << ((0) & 63))
                        _ = v28
                    }
                    if true {
                        var v29 uint64 = v6
                        _ = v29
                    } else {
                        _ = len(v4)
                    }
                }
                if true {
                    v30 := GenS2{F0: (v6 - 438), F1: v5}
                    fmt.Println(v5)
                    v2 = (^(v2))
                    _ = v30
                }
            }
            _ = fi19
        }
        var v31 int8 = 1
        _ = fi17
        _ = v31
    }
    v5 += genFn1(127)
    for fi32 := 0; fi32 < 11; fi32++ {
        switch (fi32 << ((18446744073709551615) & 63)) {
        case -2147483648:
            v6 += (v6 << ((57) & 63))
        case 42:
            switch "\n\t" {
            case "a":
                fmt.Println(v5)
            case "\n\t":
                v33 := []int{int(v6), fi32, (fi32 >> ((872) & 63)), fi32}
                _ = v33
            }
        }
        _ = fi32
    }
    switch "" {
    case "":
        switch 127 {
        case -2147483648:
            var v34 uint = 4294967295
            _ = v34
        case -134:
            fmt.Println(v3, v2, v1)
        case 1:
            if ((false && false) || true) {
                fmt.Println(v2, v3, v1)
            } else {
                fmt.Println(v16, v16, v3)
            }
        default:
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 bool = false
                if ((("\n\t" < "") || ((!(-232 != 9223372036854775807)) || (false && ((false || false) && (1.5 == 1.5))))) || (v6 > v6)) {
                    for fi37 := 0; fi37 < 11; fi37++ {
                        v38 := GenS3{F0: -1, F1: ("hello" + "hello"), F2: uint32(v5)}
                        v4[65535] = 1
                        fmt.Println(v2, fi35, v3)
                        _ = fi37
                        _ = v38
                    }
                }
                _ = fi35
                _ = v36
            }
        }
    default:
        var v39 int = 42
        _ = v39
    }
    v40 := []int{128, 128, 783}
    switch 9223372036854775807 {
    case -2147483648:
        v5 -= (v5 >> ((0) & 63))
    case 127:
        var v41 uint64 = (v6 % (((v6 + 42)) | 1))
        _ = v41
    default:
        v3 -= v3
    }
    switch "a" {
    case "a":
        if (!true) {
            if ((((v3 < (v3 / ((v3) | 1))) || false) && ((0.1 == 0.1) || false)) || (int(v1) < 42)) {
                var v42 float64 = (1.5 - ((-2.25 * -2.25) * (1e10 / 1e10)))
                var v43 int = int(v5)
                _ = v42
                _ = v43
            }
            delete(v4, 351)
        }
    case "hello":
        fmt.Println(v1, v1, v5)
    }
    if false {
        fmt.Println(v16, v1)
        var v44 int = 128
        v3 -= 1
        _ = v44
    }
    var v45 bool = true
    var v46 uint64 = (v6 << ((639) & 63))
    v1 -= (v1 ^ -32768)
    if (59 < 796) {
        var v47 int = 412
        v48 := v4[0]
        v49 := GenS2{F0: (v6 ^ 18446744073709551615), F1: (v5 + (v5 ^ genFn1(100)))}
        _ = v47
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
