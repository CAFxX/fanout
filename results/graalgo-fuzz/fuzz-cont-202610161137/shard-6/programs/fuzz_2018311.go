package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 int8; F2 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint16 = 832
    var v2 uint16 = (v1 % (((^(v1))) | 1))
    var v3 int64 = 9223372036854775807
    var v4 uint = 18446744073709551615
    v5 := make(map[string]int)
    var v6 int64 = (^(v3))
    for fi7 := 0; fi7 < 4; fi7++ {
        v6 = 0
        switch fi7 {
        case 42:
            for fi8 := 0; fi8 < 13; fi8++ {
                v2 = v2
                if ("\n\t" > ("" + "\n\t")) {
                    v5["foo bar"] = -2147483648
                    v4 = 65535
                }
                v9 := GenS1{F0: v4, F1: 1000, F2: (true && ((v4 / ((v4) | 1)) != v4))}
                _ = fi8
                _ = v9
            }
        case 9223372036854775807:
            v10 := GenS1{F0: (v4 & v4), F1: -1, F2: true}
            _ = v10
        case 246:
            var v11 int8 = int8(v6)
            _ = v11
        }
        fmt.Println(v6, v2)
        _ = fi7
    }
    v12 := GenS2{F0: -2147483648, F1: 127, F2: (v4 ^ (v4 + 0))}
    switch -598 {
    case 680:
        v4 = 42
    case -2147483648:
        if true {
            for fi13 := 0; fi13 < 16; fi13++ {
                for fi14 := 0; fi14 < 1; fi14++ {
                    if true {
                        if false {
                            var v15 int32 = 1
                            fmt.Println(v6, v15)
                            v15 = 2147483647
                            _ = v15
                        }
                        var v16 uint = v4
                        _ = v16
                    } else {
                        var v17 int8 = -56
                        switch fi13 {
                        case -1:
                            var v18 int = (fi13 ^ fi14)
                            _ = v18
                        default:
                            v1 = 32767
                        }
                        var v19 uint32 = 62
                        v20 := GenS1{F0: (v4 % ((v4) | 1)), F1: -32768, F2: false}
                        _ = v17
                        _ = v19
                        _ = v20
                    }
                    v21 := v5["\n\t"]
                    fmt.Println(v6)
                    _ = fi14
                    _ = v21
                }
                var v22 uint = (v4 & 2)
                for fi23 := 0; fi23 < 2; fi23++ {
                    if false {
                        _ = len(v5)
                    } else {
                        v3 = v6
                        v24 := []int{(fi23 ^ fi23), fi23, fi13, fi13}
                        var v25 bool = false
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(v22, v22)
                    switch fi13 {
                    case -430:
                        var v26 uint64 = 0
                        if false {
                            v27 := GenS2{F0: fi23, F1: -63, F2: (v22 + (v4 / (((^(v4))) | 1)))}
                            _ = v27
                        } else {
                            fmt.Println(v22)
                            v28 := GenS3{F0: v2}
                            v29 := GenS4{F0: 334}
                            _ = v28
                            _ = v29
                        }
                        var v30 bool = false
                        _ = v26
                        _ = v30
                    default:
                        var v31 uint32 = 896
                        _ = v31
                    }
                    _ = fi23
                }
                _ = fi13
                _ = v22
            }
        } else {
            for fi32 := 0; fi32 < 6; fi32++ {
                var v33 float32 = (1e10 + 3.14159)
                _ = fi32
                _ = v33
            }
            v34 := GenS1{F0: (v4 & v4), F1: -1, F2: false}
            _ = v34
        }
    case -1:
        if (false && false) {
            var v35 uint8 = 255
            var v36 int16 = 32767
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint32 = 65535
                fmt.Println(v2, v35, fi37)
                switch fi37 {
                case 590:
                    switch (fi37 - (fi37 % (((fi37 & 1)) | 1))) {
                    case -475:
                        v39 := []int{fi37, (^(fi37)), (-(fi37))}
                        _ = v39
                    case 9223372036854775807:
                        v3 -= 9223372036854775807
                    case 0:
                        var v40 float64 = ((1e10 - (float64(3.14159) - float64(1.5))) / (1.5 / 3.14159))
                        _ = v40
                    default:
                        var v41 uint16 = v2
                        _ = v41
                    }
                }
                _ = fi37
                _ = v38
            }
            _ = v35
            _ = v36
        }
    default:
        v42 := GenS4{F0: 18446744073709551615}
        _ = v42
    }
    v43 := []int{0}
    v4 -= (v4 << (((v4 / ((v4) | 1))) & 63))
    v44 := GenS5{F0: 583}
    var v45 int64 = v6
    if true {
        switch ("hello" + ("foo bar" + (("\n\t" + ("\n\t" + "")) + (("hello" + ("\n\t" + "\n\t")) + "")))) {
        case "hello":
            if false {
                v46 := GenS2{F0: int(v1), F1: 24, F2: v4}
                var v47 float64 = (1.5 / -2.25)
                v4 = v4
                _ = v46
                _ = v47
            }
        case "foo bar":
            var v48 uint8 = 177
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 16; fi49++ {
        fmt.Println(v45, v4)
        _ = fi49
    }
    var v50 uint64 = 4294967295
    for fi51 := 0; fi51 < 9; fi51++ {
        switch "" {
        case "foo bar":
            if (!true) {
                v52 := GenS2{F0: fi51, F1: 1, F2: (v4 >> ((65535) & 63))}
                var v53 int32 = -1
                if true {
                    v50 -= (v50 << (((v4 % ((18446744073709551615) | 1))) & 63))
                } else {
                    fmt.Println(v45, fi51, v2)
                }
                _ = v52
                _ = v53
            }
        case "hello":
            if (false || false) {
                if (644 == 0) {
                    switch (-(fi51)) {
                    case 1:
                        for fi54 := 0; fi54 < 19; fi54++ {
                            fmt.Println(v4, v6)
                            v55 := GenS5{F0: 2147483647}
                            _ = fi54
                            _ = v55
                        }
                    }
                } else {
                    fmt.Println(v4, v2, fi51)
                }
                v56 := GenS4{F0: v50}
                v57 := GenS2{F0: fi51, F1: 42, F2: 2}
                _ = v56
                _ = v57
            } else {
                for fi58 := 0; fi58 < 10; fi58++ {
                    var v59 int32 = 42
                    if (("hello" + "a") >= "\n\t") {
                        var v60 int64 = (v3 * 42)
                        fmt.Println(v60, v59)
                        fmt.Println(v4, fi58)
                        _ = v60
                    }
                    var v61 int8 = -112
                    _ = fi58
                    _ = v59
                    _ = v61
                }
                if len(v43) > 0 {
                    v43[0] = (fi51 >> ((750) & 63))
                }
                if false {
                    v50 = v50
                    for fi62 := 0; fi62 < 7; fi62++ {
                        var v63 uint8 = 200
                        fmt.Println(v1, v2, v3)
                        _ = len(v5)
                        _ = fi62
                        _ = v63
                    }
                }
            }
        default:
            switch fi51 {
            case -94:
                var v64 bool = true
                _ = v64
            case 32767:
                if true {
                    for fi65 := 0; fi65 < 8; fi65++ {
                        v45 += v3
                        switch 0 {
                        case -2147483648:
                            if (((v6 & v3) > v45) && false) {
                                v45 = -1
                                _ = len(v43)
                                var v66 int64 = -9223372036854775808
                                _ = v66
                            }
                        }
                        var v67 int = (^(fi65))
                        _ = fi65
                        _ = v67
                    }
                    var v68 float64 = 0.1
                    _ = v68
                } else {
                    for fi69 := 0; fi69 < 6; fi69++ {
                        var v70 string = ("\n\t" + "hello")
                        _ = fi69
                        _ = v70
                    }
                    fmt.Println(v3)
                }
            default:
                var v71 uint = v4
                _ = v71
            }
        }
        fmt.Println(v45, v1)
        _ = fi51
    }
    fmt.Println(v3, v6, v3)
    v6 -= (^(v3))
    for fi72 := 0; fi72 < 12; fi72++ {
        v4 = (v4 % ((v4) | 1))
        if true {
            fmt.Println(v1, v2)
        } else {
            var v73 int16 = 1
            switch fi72 {
            case -9223372036854775808:
                v5[((("\n\t" + "") + "a") + "\n\t")] = (fi72 >> ((v4) & 63))
            case 32:
                if true {
                    var v74 int32 = -808
                    v75 := GenS4{F0: v50}
                    _ = v74
                    _ = v75
                } else {
                    v6 -= int64(v73)
                }
            default:
                for fi76 := 0; fi76 < 3; fi76++ {
                    v77 := v5[""]
                    _ = fi76
                    _ = v77
                }
            }
            _ = v73
        }
        for fi78 := 0; fi78 < 4; fi78++ {
            v79 := GenS5{F0: -627}
            delete(v5, "")
            _ = fi78
            _ = v79
        }
        _ = fi72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v50
    _ = v6
}
