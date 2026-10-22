package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int }
type GenS6 struct { F0 uint64; F1 string; F2 int16 }
type GenS7 struct { F0 int32; F1 int64 }
type GenS8 struct { F0 int32; F1 bool }
func genFn0() bool {
    if true {
        v7 := GenS1{F0: 255}
        _ = v7
    } else {
        var v8 int32 = 2147483647
        if false {
            v8 = v8
        }
        _ = v8
    }
    var v9 int = 32768
    v9 = (v9 & (-(v9)))
    v9 -= (v9 & v9)
    _ = v9
    return (!true)
}
func genFn1() uint8 {
    v10 := GenS1{F0: 2}
    var v11 bool = true
    switch "hello" {
    case "a":
        if true {
            if (0 > -100) {
                v12 := []int{4294967295, -9223372036854775808, 32767}
                var v13 uint8 = 2
                v11 = (3.14159 < 1.5)
                _ = v12
                _ = v13
            }
        } else {
            v14 := []int{2147483647, 746, -244, 2147483647}
            for fi15 := 0; fi15 < 3; fi15++ {
                v11 = true
                var v16 string = ("foo bar" + ("foo bar" + ""))
                for fi17 := 0; fi17 < 1; fi17++ {
                    switch v16 {
                    case "foo bar":
                        var v18 int16 = 42
                        _ = v18
                    }
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    }
    switch 0 {
    case 4294967295:
        for fi19 := 0; fi19 < 18; fi19++ {
            v11 = v11
            var v20 uint32 = 536
            v21 := GenS2{F0: 774}
            _ = fi19
            _ = v20
            _ = v21
        }
    case 9223372036854775807:
        switch "hello" {
        case "\n\t":
            v22 := GenS3{F0: (18446744073709551615 <= 2), F1: 1, F2: v11}
            _ = v22
        default:
            v11 = (127 < 368)
        }
    case -2147483648:
        if (v11 && v11) {
            v23 := GenS3{F0: false, F1: 42, F2: v11}
            _ = v23
        } else {
            var v24 int32 = -1
            v25 := []int{841, -9223372036854775808, 128, -1}
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        switch (fi26 - (fi26 / ((fi26) | 1))) {
        case 32767:
            if (0 >= 0) {
                v27 := GenS4{F0: 1, F1: 253}
                _ = v27
            } else {
                v11 = (!v11)
            }
        case 1:
            switch fi26 {
            case -1:
                if (!v11) {
                    if v11 {
                        switch -1 {
                        case -9223372036854775808:
                            if v11 {
                                v11 = (!(128 >= 2))
                                v11 = ((845 >= 660) && v11)
                                v11 = v11
                            } else {
                                if v11 {
                                    fmt.Println(v11, v11)
                                }
                                v28 := GenS4{F0: 1, F1: 128}
                                _ = v28
                            }
                        case 32767:
                            if false {
                                var v29 int32 = 1
                                for fi30 := 0; fi30 < 5; fi30++ {
                                    if (715 > 2) {
                                        v29 += 1
                                    } else {
                                        var v31 float32 = (1e10 / 1.0)
                                        var v32 bool = v11
                                        v33 := GenS5{F0: 100, F1: 745, F2: 9223372036854775807}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    v34 := []int{fi30, fi26, (fi30 * -161), (fi26 - fi26)}
                                    switch (fi26 - (fi26 & -1)) {
                                    case -301:
                                        switch ("a" + "\n\t") {
                                        case "a":
                                            if len(v34) >= 2 {
                                                v34 = v34[1:len(v34)]
                                            }
                                        case "\n\t":
                                            var v35 string = ("hello" + ("hello" + "foo bar"))
                                            _ = v35
                                        case "hello":
                                            v36 := GenS2{F0: (v29 - v29)}
                                            _ = v36
                                        default:
                                            v29 = (v29 << ((42) & 63))
                                        }
                                    default:
                                        if (18446744073709551615 == 1) {
                                            var v37 float32 = (1e10 - -2.25)
                                            _ = v37
                                        }
                                    }
                                    _ = fi30
                                    _ = v34
                                }
                                switch 127 {
                                case -909:
                                    if (0 != 255) {
                                        var v38 int64 = -305
                                        var v39 int32 = (v29 >> ((125) & 63))
                                        v40 := GenS2{F0: (v39 ^ v39)}
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v11 = v11
                                    }
                                case 127:
                                    var v41 int32 = v29
                                    _ = v41
                                case 2147483647:
                                    v29 += (-(v29))
                                default:
                                    v42 := []int{fi26, (^(fi26)), fi26, fi26}
                                    _ = v42
                                }
                                _ = v29
                            }
                        }
                        v43 := []int{}
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 2; fi44++ {
                        v45 := []int{42}
                        v11 = (!v11)
                        _ = fi44
                        _ = v45
                    }
                    if (!v11) {
                        switch "" {
                        case "\n\t":
                            v46 := GenS4{F0: -1, F1: 224}
                            _ = v46
                        default:
                            var v47 int16 = 32767
                            _ = v47
                        }
                        var v48 float64 = (1.5 / 1.5)
                        _ = v48
                    }
                } else {
                    for fi49 := 0; fi49 < 9; fi49++ {
                        switch "a" {
                        case "hello":
                            fmt.Println(fi26)
                        }
                        v50 := GenS6{F0: 0, F1: "foo bar", F2: -21}
                        var v51 int16 = 1000
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 14; fi52++ {
                        v11 = ((fi26 << ((255) & 63)) > fi52)
                        v53 := GenS2{F0: 1}
                        v11 = v11
                        _ = fi52
                        _ = v53
                    }
                }
            }
        case -667:
            var v54 uint64 = 4294967295
            _ = v54
        default:
            var v55 int32 = -1
            _ = v55
        }
        if (32767 >= 1) {
            v11 = (-2.25 <= float32(0.1))
            v56 := []int{0, (fi26 / ((fi26) | 1))}
            v11 = (!(v11 && (!(((3.14159 - (0.1 * (1.5 / -2.25))) - ((0.0 * (0.0 * -2.25)) / 1.0)) <= (0.1 - (float64(float32(0.1)) * 0.1))))))
            _ = v56
        } else {
            v57 := GenS7{F0: -2147483648, F1: 9223372036854775807}
            _ = v57
        }
        v11 = v11
        _ = fi26
    }
    _ = v10
    _ = v11
    return 215
}

func main() {
    var v1 uint64 = 2
    var v2 int32 = 500
    var v3 int64 = 0
    var v4 uint16 = 746
    var v5 uint = 2
    var v6 bool = ((1 <= (v2 | (v2 & v2))) && ((v3 ^ v3) < 9223372036854775807))
    var v58 float64 = 1e10
    var v59 int32 = (v2 << ((255) & 63))
    fmt.Println(v6, v5)
    v60 := GenS8{F0: (^(v59)), F1: true}
    v61 := []int{32768, 32768, 128, 2147483647}
    if v6 {
        for fi62 := 0; fi62 < 8; fi62++ {
            if false {
                switch (fi62 ^ fi62) {
                case 32767:
                    var v63 uint64 = 1
                    _ = v63
                }
            }
            var v64 int = (fi62 >> ((v5) & 63))
            _ = fi62
            _ = v64
        }
        if len(v61) > 0 {
            v61[0] = 32767
        }
        for fi65 := 0; fi65 < 13; fi65++ {
            var v66 bool = (408 <= 4294967295)
            v4 -= v4
            v59 -= int32(v3)
            _ = fi65
            _ = v66
        }
    }
    v2 = v59
    _ = len(v61)
    v4 -= (v4 + v4)
    if v6 {
        var v67 int32 = (v59 * v59)
        var v68 int32 = (v2 << ((4294967295) & 63))
        _ = v67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
