package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 bool; F2 uint64 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 uint = 392
    var v2 int32 = 938
    var v3 int8 = 54
    var v4 int8 = v3
    v5 := []int{-452, 32767}
    v5 = append(v5, -945)
    v6 := GenS1{F0: (v1 & v1)}
    v4 -= 100
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint = (^(v1))
            v9 += (v9 & 2)
            switch fi7 {
            case 127:
                switch (fi8 * fi7) {
                case -924:
                    v9 = (v1 + v1)
                case 282:
                    switch 42 {
                    case -197:
                        var v10 int16 = 1
                        _ = v10
                    case -984:
                        if (true || (0.1 <= (float32((1.5 * (1.5 - -2.25))) - 0.0))) {
                            if (!false) {
                                v5 = append(v5, fi7)
                                v11 := GenS2{F0: 859, F1: 4294967295, F2: 0}
                                _ = v11
                            } else {
                                fmt.Println(fi7, v4)
                                v12 := GenS1{F0: (v9 | v1)}
                                var v13 int16 = 0
                                _ = v12
                                _ = v13
                            }
                            v9 -= 2
                            fmt.Println(v1, v4)
                        }
                    default:
                        var v14 uint32 = 1
                        _ = v14
                    }
                }
            case -9223372036854775808:
                v15 := GenS1{F0: v9}
                _ = v15
            case -833:
                v4 = (v4 ^ (v4 + 1))
            }
            _ = fi8
            _ = v9
        }
        if len(v5) > 0 {
            v5[0] = (fi7 << ((v1) & 63))
        }
        _ = fi7
    }
    v5 = append(v5, 9223372036854775807)
    if (0.0 <= ((-2.25 / 0.1) - float64(1.5))) {
        v2 = (^(v2))
    }
    var v16 int64 = 9223372036854775807
    var v17 uint = (v1 >> ((1) & 63))
    _ = len(v5)
    for fi18 := 0; fi18 < 4; fi18++ {
        v5 = append(v5, fi18)
        v1 = v17
        fmt.Println(v16, fi18, v2)
        _ = fi18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v5 = append(v5, (^(fi19)))
        if (!(!false)) {
            fmt.Println(v16)
        } else {
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := GenS2{F0: (v17 * 2), F1: 391, F2: 724}
                if ((v2 % ((1) | 1)) != 1) {
                    if (false || (2147483648 >= 4294967295)) {
                        switch fi20 {
                        case -9223372036854775808:
                            var v22 uint = v1
                            _ = v22
                        case 4294967295:
                            v1 += v17
                        case 9223372036854775807:
                            fmt.Println(v4, fi19)
                        }
                        var v23 uint16 = 42
                        if (0.1 != float32((1e10 / 1.5))) {
                            var v24 int = (fi20 | fi19)
                            fmt.Println(v24, v24, v2)
                            _ = v24
                        } else {
                            if true {
                                if (!(true || true)) {
                                    v25 := GenS1{F0: v1}
                                    if false {
                                        if true {
                                            v1 -= (v17 & (v17 - (v17 - 146)))
                                            if false {
                                                if ((-(v2)) < v2) {
                                                    fmt.Println(v4, v16, v16)
                                                    switch (("" + "") + ("\n\t" + (("hello" + "") + ("hello" + "a")))) {
                                                    case "":
                                                        v1 -= v1
                                                    case "a":
                                                        switch (("" + "hello") + "foo bar") {
                                                        case "a":
                                                            var v26 uint32 = 2147483647
                                                            _ = v26
                                                        case "foo bar":
                                                            if false {
                                                                v27 := GenS2{F0: 1, F1: 270, F2: 1}
                                                                _ = len(v5)
                                                                _ = v27
                                                            } else {
                                                                switch int(v3) {
                                                                case 32768:
                                                                    if (627 >= 0) {
                                                                        v5 = append(v5, int(v3))
                                                                        var v28 float32 = 3.14159
                                                                        v23 -= v23
                                                                        v2 -= v2
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        _ = v28
                                                                    }
                                                                }
                                                                var v29 uint64 = 1
                                                                v1 = (v1 - 171)
                                                                _ = v29
                                                            }
                                                        }
                                                    default:
                                                        var v30 int32 = 647
                                                        _ = v30
                                                    }
                                                } else {
                                                    var v31 int16 = -689
                                                    v23 -= 32767
                                                    _ = v31
                                                }
                                            } else {
                                                v32 := GenS3{F0: 0}
                                                _ = v32
                                            }
                                            var v33 int16 = int16(fi19)
                                            _ = v33
                                        } else {
                                            v3 += -33
                                        }
                                        switch ("" + ("foo bar" + "")) {
                                        case "hello":
                                            fmt.Println(v4, v1)
                                        case "\n\t":
                                            v5 = append(v5, (-(fi19)))
                                        }
                                        v34 := GenS1{F0: 4294967295}
                                        _ = v34
                                    } else {
                                        var v35 int = fi20
                                        v3 = 0
                                        var v36 string = ""
                                        if true {
                                            v3 = v4
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            switch ("hello" + "foo bar") {
                                            case "\n\t":
                                                var v37 float32 = float32(-2.25)
                                                _ = v37
                                            case "a":
                                                v38 := GenS4{F0: int8(v23), F1: true, F2: 622}
                                                _ = v38
                                            }
                                        }
                                        _ = v35
                                        _ = v36
                                    }
                                    var v39 int16 = 1
                                    v40 := GenS3{F0: v39}
                                    v2 = v2
                                    v23 -= v23
                                    var v41 uint = (v17 >> ((v17) & 63))
                                    _ = v25
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                } else {
                                    v17 += (v1 + v1)
                                }
                                if len(v5) > 0 {
                                    v5[0] = (fi20 << ((v1) & 63))
                                }
                                if false {
                                    switch (-(fi20)) {
                                    case -800:
                                        var v42 uint64 = 249
                                        _ = v42
                                    case 42:
                                        fmt.Println(v1, v1)
                                    }
                                    if ((1.5 / 1.5) != ((1.5 * float64(3.14159)) + (1e10 / 1.5))) {
                                        _ = len(v5)
                                    } else {
                                        v4 = v3
                                        var v43 uint32 = 170
                                        var v44 string = ""
                                        _ = v43
                                        _ = v44
                                    }
                                    v17 += (v17 << ((42) & 63))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            } else {
                                v2 += -341
                                v5 = append(v5, (fi20 + (fi19 * fi19)))
                            }
                            var v45 uint32 = 0
                            _ = v45
                        }
                        v46 := GenS1{F0: v17}
                        var v47 uint8 = 157
                        _ = v23
                        _ = v46
                        _ = v47
                    }
                }
                if false {
                    switch "a" {
                    case "foo bar":
                        v48 := GenS5{F0: "foo bar", F1: v3}
                        _ = v48
                    default:
                        var v49 bool = false
                        _ = v49
                    }
                    switch fi19 {
                    case 4294967295:
                        var v50 uint16 = 284
                        _ = v50
                    }
                    fmt.Println(v3, v2)
                    fmt.Println(v16)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v51 := GenS4{F0: v3, F1: true, F2: 18446744073709551615}
                    switch ("foo bar" + "foo bar") {
                    case "foo bar":
                        var v52 int64 = v16
                        _ = v52
                    case "hello":
                        if false {
                            var v53 uint = 540
                            _ = v53
                        } else {
                            v3 -= v3
                            var v54 uint32 = 828
                            _ = v54
                        }
                    default:
                        v55 := GenS3{F0: 1000}
                        _ = v55
                    }
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        if (((v2 + 66) >= 2147483647) || false) {
                            var v56 int8 = v3
                            v4 = v4
                            _ = v56
                        } else {
                            if ((!false) || false) {
                                v17 -= (^(v1))
                                if (2 < 2) {
                                    v57 := GenS3{F0: -32768}
                                    _ = v57
                                } else {
                                    v16 += v16
                                    var v58 bool = ((255 >= 32767) && true)
                                    _ = v58
                                }
                            }
                            fmt.Println(v4, v2)
                        }
                    case "hello":
                        v3 += v4
                    }
                    _ = v51
                }
                var v59 int32 = (v2 >> ((743) & 63))
                var v60 string = "\n\t"
                _ = fi20
                _ = v21
                _ = v59
                _ = v60
            }
            v61 := GenS6{F0: -1}
            _ = v61
        }
        _ = fi19
    }
    for fi62 := 0; fi62 < 9; fi62++ {
        if len(v5) > 0 {
            v5[0] = fi62
        }
        _ = fi62
    }
    var v63 int = -740
    v5 = append(v5, v63)
    if (v3 != (v4 << ((v1) & 63))) {
        var v64 float32 = (0.1 - (1e10 - (0.1 * (0.1 * 0.1))))
        if false {
            _ = len(v5)
        } else {
            var v65 uint64 = 42
            var v66 int = (v63 * v63)
            _ = v65
            _ = v66
        }
        v17 = (v1 + (v1 ^ (v17 + (v17 & (^(v1))))))
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
}
