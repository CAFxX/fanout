package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0(p0 int, p1 uint8, p2 int64) int8 {
    var v7 float64 = 1.5
    switch ("a" + ("a" + (("foo bar" + "") + "hello"))) {
    case "foo bar":
        var v8 int8 = 1
        _ = v8
    case "hello":
        v9 := []int{}
        _ = v9
    case "":
        fmt.Println(p2, p1, p1)
    }
    var v10 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 40
}
func genFn1() int {
    var v11 int16 = -32768
    if false {
        v12 := GenS1{F0: 0, F1: 128}
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 int8 = 1
            switch "foo bar" {
            case "a":
                if false {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        fmt.Println(fi13)
                        var v16 uint64 = 543
                        _ = fi15
                        _ = v16
                    }
                } else {
                    var v17 uint64 = 739
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS2{F0: (v17 & v17), F1: 271}
                        if false {
                            switch (fi18 << ((18446744073709551615) & 63)) {
                            case -654:
                                var v20 uint16 = 65535
                                _ = v20
                            default:
                                v11 = -32768
                            }
                        }
                        var v21 int = -2147483648
                        _ = fi18
                        _ = v19
                        _ = v21
                    }
                    v11 += v11
                    _ = v17
                }
            case "foo bar":
                for fi22 := 0; fi22 < 16; fi22++ {
                    v23 := GenS1{F0: 2, F1: -300}
                    v14 += 100
                    fmt.Println(fi13, v11, v14)
                    _ = fi22
                    _ = v23
                }
            case "\n\t":
                switch fi13 {
                case 32768:
                    for fi24 := 0; fi24 < 13; fi24++ {
                        var v25 uint32 = 900
                        _ = fi24
                        _ = v25
                    }
                case -378:
                    switch "foo bar" {
                    case "a":
                        v11 = int16(fi13)
                    }
                case 2147483647:
                    v14 += genFn0(fi13, 128, -1)
                default:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        fmt.Println(fi13)
                        v27 := GenS2{F0: 2, F1: 794}
                        _ = fi26
                        _ = v27
                    }
                }
            default:
                v28 := []int{(^(fi13))}
                _ = v28
            }
            switch "foo bar" {
            case "hello":
                switch "" {
                case "hello":
                    var v29 int16 = (v11 & 32767)
                    _ = v29
                case "":
                    for fi30 := 0; fi30 < 18; fi30++ {
                        var v31 int = int(v14)
                        var v32 uint64 = 639
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            case "foo bar":
                v14 = -1
            case "\n\t":
                for fi33 := 0; fi33 < 13; fi33++ {
                    v11 = v11
                    if ((560 == (^(v11))) || (false || (0 != 4294967295))) {
                        fmt.Println(v11, v14)
                        v34 := []int{}
                        _ = v34
                    } else {
                        v35 := GenS2{F0: 870, F1: 0}
                        v36 := GenS3{F0: 1, F1: 0, F2: v11}
                        v37 := GenS4{F0: (fi33 + fi13), F1: 128, F2: true}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    v38 := GenS5{F0: 634, F1: "a"}
                    var v39 uint16 = 2
                    v39 = 255
                    _ = fi33
                    _ = v38
                    _ = v39
                }
            }
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        v11 = (v11 / (((v11 - (^(v11)))) | 1))
        for fi41 := 0; fi41 < 2; fi41++ {
            switch (^(fi41)) {
            case -2147483648:
                fmt.Println(fi41, fi41, fi41)
                fmt.Println(v11, fi41)
                var v42 uint = 0
                switch "a" {
                case "a":
                    switch fi41 {
                    case -1:
                        v42 = v42
                    }
                }
                if (477 > 2147483648) {
                    var v43 uint8 = 128
                    var v44 int8 = 127
                    _ = v43
                    _ = v44
                }
                _ = v42
            }
            v45 := GenS6{F0: (fi41 >> ((1) & 63))}
            v11 = 1000
            _ = fi41
            _ = v45
        }
        _ = fi40
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        v47 := GenS7{F0: 531, F1: 128, F2: 0}
        var v48 int = 0
        v49 := []int{fi46}
        _ = fi46
        _ = v47
        _ = v48
        _ = v49
    }
    fmt.Println(v11)
    _ = v11
    return 45
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int64 = 0
    v3 := make(map[int]bool)
    var v4 string = ""
    var v5 int8 = 1
    var v6 float64 = float64((float32(0.0) - -2.25))
    var v50 uint64 = 757
    for fi51 := 0; fi51 < 3; fi51++ {
        if false {
            fmt.Println(v1, v1, v5)
        } else {
            v1 = (v1 / (((-(fi51))) | 1))
            v4 = v4
        }
        _ = fi51
    }
    v5 = 42
    var v52 int32 = 906
    v52 = v52
    v53 := GenS4{F0: -641, F1: 42, F2: true}
    v5 += 42
    for fi54 := 0; fi54 < 14; fi54++ {
        if (1 >= 0) {
            v55 := GenS8{F0: 257}
            if false {
                switch (fi54 ^ genFn1()) {
                case 185:
                    v1 = int(v5)
                case -1:
                    for fi56 := 0; fi56 < 2; fi56++ {
                        fmt.Println(v1, v50, v4)
                        _ = fi56
                    }
                }
                var v57 int = (fi54 % (((fi54 - fi54)) | 1))
                if false {
                    v58 := []int{(v1 + (v1 << ((255) & 63))), (v1 << ((18446744073709551615) & 63)), (fi54 + 32767), 0}
                    var v59 float64 = float64(float32((float64((3.14159 * (0.0 - 3.14159))) + v6)))
                    _ = v58
                    _ = v59
                } else {
                    var v60 bool = true
                    v61 := []int{}
                    v62 := GenS9{F0: 42}
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = v57
            } else {
                var v63 int16 = 1000
                fmt.Println(v50)
                _ = v63
            }
            for fi64 := 0; fi64 < 15; fi64++ {
                switch v4 {
                case "\n\t":
                    if true {
                        v1 = (v1 % ((32767) | 1))
                        v5 = v5
                        v6 -= (0.1 / (0.1 / ((0.0 * v6) / 3.14159)))
                    }
                }
                _ = fi64
            }
            _ = v55
        }
        _ = fi54
    }
    v2 -= 1
    v65 := GenS7{F0: (v2 >> ((4294967295) & 63)), F1: 255, F2: (v52 * 1)}
    v5 -= genFn0((v1 << ((255) & 63)), 186, v2)
    if (729 < 65535) {
        var v66 uint = 785
        fmt.Println(v4)
        _ = v66
    } else {
        v6 = v6
        for fi67 := 0; fi67 < 12; fi67++ {
            v68 := GenS10{F0: v2, F1: 682, F2: 42}
            var v69 uint = 42
            var v70 bool = false
            _ = fi67
            _ = v68
            _ = v69
            _ = v70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v52
    _ = v53
    _ = v6
    _ = v65
}
