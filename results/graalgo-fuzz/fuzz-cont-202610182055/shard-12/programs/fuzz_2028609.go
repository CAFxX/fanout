package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int; F2 string }

func main() {
    var v1 int64 = 0
    var v2 float64 = 0.1
    var v3 string = "foo bar"
    var v4 int = -2147483648
    var v5 uint = 2
    v6 := []int{2147483647, 2147483647, (v4 - v4)}
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            if true {
                var v9 bool = (false && (1 != 137))
                v10 := GenS1{F0: (v5 & v5)}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
        if (42 == 127) {
            if false {
                v4 -= fi7
                v4 = (fi7 % ((-2147483648) | 1))
                var v11 bool = ((42 > 459) || (false && false))
                _ = v11
            }
        } else {
            v6 = append(v6, (fi7 / ((fi7) | 1)))
            v4 = -2147483648
            switch fi7 {
            case 0:
                v3 = v3
            case 126:
                switch (fi7 ^ (fi7 >> ((v5) & 63))) {
                case 552:
                    v4 += (v4 % (((fi7 * -2147483648)) | 1))
                case 224:
                    v5 = v5
                case 975:
                    var v12 int16 = 1000
                    _ = v12
                }
            }
        }
        var v13 int16 = 32767
        _ = fi7
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v4 = (v4 << ((255) & 63))
        _ = fi14
    }
    if ((v3 + "") > (((((((v3 + v3) + ((v3 + ("hello" + (v3 + "a"))) + v3)) + (v3 + v3)) + v3) + ("a" + "foo bar")) + v3) + (v3 + ((v3 + ("foo bar" + v3)) + v3)))) {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) > 0 {
                v6[0] = v4
            }
            v15 := GenS1{F0: v5}
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v1 += v1
            _ = fi16
        }
    }
    v4 = (v4 * v4)
    _ = len(v6)
    if true {
        if true {
            fmt.Println(v5, v5)
        }
        v4 += (v4 ^ (v4 + (v4 & (-(v4)))))
        if (false || (1 >= 18446744073709551615)) {
            v17 := GenS2{F0: 126, F1: -2147483648, F2: 32767}
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint64 = 650
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch ((v3 + ("foo bar" + "\n\t")) + v3) {
                    case "foo bar":
                        if (32767 < (fi20 >> ((18446744073709551615) & 63))) {
                            fmt.Println(v2, v1)
                            if (!(fi20 != 2147483647)) {
                                if true {
                                    var v21 uint16 = 65535
                                    _ = v21
                                }
                            } else {
                                v4 -= (v4 * fi18)
                                if len(v6) > 0 {
                                    v6[0] = 0
                                }
                            }
                        } else {
                            fmt.Println(v5, v2, v2)
                            if (446 == 351) {
                                switch fi20 {
                                case 0:
                                    v3 = "\n\t"
                                case -2147483648:
                                    v6 = append(v6, fi20)
                                case -1:
                                    if false {
                                        v3 = (v3 + v3)
                                        if (!(!false)) {
                                            v19 += v19
                                            switch 127 {
                                            case -2147483648:
                                                v22 := GenS2{F0: 127, F1: -1, F2: 1}
                                                if len(v6) > 0 {
                                                    v6[0] = fi18
                                                }
                                                if false {
                                                    if (1000 <= 781) {
                                                        v23 := GenS1{F0: v5}
                                                        v2 = (v2 - (-2.25 / ((0.0 + v2) / 1e10)))
                                                        var v24 int8 = 127
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                    switch v3 {
                                                    case "foo bar":
                                                        fmt.Println(v3)
                                                    case "hello":
                                                        v25 := GenS3{F0: -780}
                                                        _ = v25
                                                    case "":
                                                        var v26 int32 = -1
                                                        _ = v26
                                                    }
                                                } else {
                                                    switch (fi18 ^ -44) {
                                                    case 918:
                                                        v27 := GenS3{F0: 42}
                                                        _ = v27
                                                    default:
                                                        switch (v3 + ((("a" + "") + "\n\t") + "\n\t")) {
                                                        case "foo bar":
                                                            v1 -= 148
                                                        }
                                                    }
                                                    v3 = (((((v3 + (((v3 + v3) + "hello") + v3)) + "a") + v3) + v3) + "hello")
                                                    v1 += v1
                                                }
                                                _ = v22
                                            case 2147483647:
                                                var v28 float32 = ((1e10 - ((-2.25 + (3.14159 / -2.25)) / 0.1)) / -2.25)
                                                var v29 uint64 = v19
                                                _ = v28
                                                _ = v29
                                            case 688:
                                                var v30 float64 = v2
                                                _ = v30
                                            }
                                            var v31 int64 = (v1 * (v1 % ((-9223372036854775808) | 1)))
                                            _ = v31
                                        }
                                    }
                                }
                            }
                            var v32 int8 = -1
                            _ = v32
                        }
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v33 int32 = -103
                            v34 := GenS4{F0: v33, F1: -556, F2: "hello"}
                            _ = v33
                            _ = v34
                        }
                        v35 := GenS3{F0: -2147483648}
                        _ = v35
                    case "\n\t":
                        v19 += 4294967295
                    }
                    v36 := GenS1{F0: v5}
                    var v37 float64 = v2
                    _ = fi20
                    _ = v36
                    _ = v37
                }
                _ = fi18
                _ = v19
            }
            var v38 uint32 = 2147483648
            _ = v17
            _ = v38
        } else {
            var v39 uint = v5
            _ = v39
        }
    } else {
        for fi40 := 0; fi40 < 9; fi40++ {
            v41 := GenS3{F0: -2147483648}
            v3 = v3
            v3 = "hello"
            _ = fi40
            _ = v41
        }
        v5 -= (^(v5))
        var v42 int16 = -781
        _ = v42
    }
    var v43 string = "a"
    v1 = v1
    var v44 int16 = int16(v5)
    v6 = append(v6, (v4 / ((v4) | 1)))
    switch (v4 | (v4 - (v4 & (^(v4))))) {
    case -2147483648:
        switch v4 {
        case 42:
            if (!false) {
                fmt.Println(v3, v4, v44)
                v45 := GenS1{F0: 42}
                if true {
                    fmt.Println(v1, v3, v4)
                    for fi46 := 0; fi46 < 1; fi46++ {
                        v47 := GenS3{F0: -2147483648}
                        _ = fi46
                        _ = v47
                    }
                }
                _ = v45
            }
        case -306:
            if false {
                var v48 float64 = 0.0
                _ = v48
            }
        default:
            var v49 float32 = 0.0
            _ = v49
        }
    case 32767:
        fmt.Println(v2, v1, v3)
    case 516:
        for fi50 := 0; fi50 < 14; fi50++ {
            v51 := GenS3{F0: -1}
            if (true && (1 > -67)) {
                for fi52 := 0; fi52 < 14; fi52++ {
                    if len(v6) > 0 {
                        v6[0] = int(v1)
                    }
                    _ = fi52
                }
            }
            _ = fi50
            _ = v51
        }
    }
    var v53 float32 = 1e10
    for fi54 := 0; fi54 < 19; fi54++ {
        switch v3 {
        case "a":
            v43 = v3
        case "foo bar":
            v3 = "\n\t"
        default:
            switch (fi54 & v4) {
            case 128:
                if (749 == 42) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                    var v55 uint16 = 65535
                    v4 += v4
                    _ = v55
                }
            case 550:
                v4 = (v4 + (fi54 % ((1) | 1)))
            case 595:
                var v56 string = v3
                _ = v56
            default:
                v4 -= (v4 & (^(v4)))
            }
        }
        v6 = append(v6, (^(v4)))
        _ = fi54
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        var v58 int64 = v1
        v59 := GenS2{F0: 128, F1: -1, F2: (v44 + v44)}
        fmt.Println(v43, v44, v43)
        _ = fi57
        _ = v58
        _ = v59
    }
    fmt.Println(v53, v43, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v53
    _ = v6
}
