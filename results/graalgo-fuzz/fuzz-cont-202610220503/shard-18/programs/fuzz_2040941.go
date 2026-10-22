package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 bool; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 uint32; F2 int64 }
type GenS10 struct { F0 uint; F1 int }
type GenS11 struct { F0 uint64 }

func main() {
    var v1 int64 = 42
    var v2 string = "\n\t"
    var v3 bool = (!((("\n\t" + "hello") > "a") || (1 != 18446744073709551615)))
    var v4 bool = true
    v5 := []int{256, 42}
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 = -1
        for fi7 := 0; fi7 < 4; fi7++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi7
        }
        if len(v5) > 0 {
            v5[0] = -280
        }
        _ = fi6
    }
    v8 := GenS1{F0: 32768}
    fmt.Println(v3, v2)
    v9 := GenS1{F0: 42}
    v2 = ((v2 + ("hello" + v2)) + v2)
    v2 = "foo bar"
    var v10 float64 = 1.5
    fmt.Println(v1, v1, v3)
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint64 = 786
        v1 += -9223372036854775808
        _ = fi11
        _ = v12
    }
    if (true || (0 < 1)) {
        v4 = true
    } else {
        if (0.0 >= -2.25) {
            for fi13 := 0; fi13 < 13; fi13++ {
                v10 += v10
                for fi14 := 0; fi14 < 9; fi14++ {
                    if (!v3) {
                        v15 := GenS2{F0: 18446744073709551615, F1: 255, F2: -103}
                        v16 := GenS3{F0: (fi13 | fi13), F1: 255, F2: "hello"}
                        _ = v15
                        _ = v16
                    }
                    _ = fi14
                }
                _ = fi13
            }
            switch v2 {
            case "\n\t":
                if v3 {
                    var v17 uint16 = 994
                    _ = v17
                } else {
                    var v18 uint16 = 0
                    v3 = (v3 && v4)
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v4 = ((v3 && (919 < 42)) && v3)
                        var v20 uint = 618
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
            case "foo bar":
                v21 := GenS4{F0: 65535, F1: ((v1 & 42) < v1), F2: ((v2 + "hello") + "")}
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            v10 -= (v10 / 1.0)
            if (235 < 155) {
                if v4 {
                    fmt.Println(v1, v2, fi22)
                    v4 = v3
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v2 = v2
                    if ((1 < -1) || (255 < 2)) {
                        var v23 bool = false
                        _ = v23
                    }
                    var v24 uint = 0
                    _ = v24
                }
                for fi25 := 0; fi25 < 15; fi25++ {
                    v3 = (true || v3)
                    fmt.Println(v2, fi25, v2)
                    v26 := GenS5{F0: fi25}
                    _ = fi25
                    _ = v26
                }
            } else {
                switch fi22 {
                case 2147483647:
                    if (-32768 >= -32768) {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 uint8 = 0
                            fmt.Println(v1)
                            if (v1 <= (v1 | 177)) {
                                v29 := GenS6{F0: 0}
                                v4 = (42 < -1)
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                        var v30 string = "\n\t"
                        _ = v30
                    }
                case 128:
                    v2 = (v2 + "a")
                }
                if v4 {
                    var v31 bool = false
                    if true {
                        switch fi22 {
                        case 1:
                            switch ("a" + v2) {
                            case "foo bar":
                                _ = len(v5)
                            case "\n\t":
                                switch (fi22 >> ((2) & 63)) {
                                case 2147483647:
                                    switch fi22 {
                                    case 415:
                                        v2 = v2
                                    case 2147483647:
                                        if v4 {
                                            var v32 int64 = 0
                                            v4 = (!v31)
                                            v2 = (v2 + v2)
                                            _ = v32
                                        } else {
                                            for fi33 := 0; fi33 < 13; fi33++ {
                                                if (!v31) {
                                                    var v34 uint16 = 6
                                                    if len(v5) > 0 {
                                                        v5[0] = fi22
                                                    }
                                                    var v35 int = 903
                                                    var v36 uint8 = 0
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                }
                                                if (fi33 < (fi33 * 292)) {
                                                    v1 -= 1
                                                }
                                                _ = fi33
                                            }
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v37 uint8 = 176
                                            _ = v37
                                        }
                                    case -9223372036854775808:
                                        if (139 == 1) {
                                            switch (fi22 / ((fi22) | 1)) {
                                            case 2147483647:
                                                var v38 uint = 65535
                                                _ = v38
                                            default:
                                                if v31 {
                                                    for fi39 := 0; fi39 < 5; fi39++ {
                                                        fmt.Println(v4, v1)
                                                        var v40 float64 = (0.0 / (v10 + float64(float32(v10))))
                                                        _ = fi39
                                                        _ = v40
                                                    }
                                                } else {
                                                    for fi41 := 0; fi41 < 11; fi41++ {
                                                        switch v2 {
                                                        case "":
                                                            var v42 string = "foo bar"
                                                            _ = v42
                                                        default:
                                                            v31 = v3
                                                        }
                                                        if v4 {
                                                            if false {
                                                                v43 := GenS7{F0: (!((1 >= 459) || (v31 && v31)))}
                                                                _ = v43
                                                            }
                                                            switch (fi41 - (fi41 * (fi41 + -9223372036854775808))) {
                                                            case 128:
                                                                v1 = v1
                                                            case 2147483647:
                                                                switch ((v2 + "") + "a") {
                                                                case "\n\t":
                                                                    v44 := GenS3{F0: (fi41 << ((18446744073709551615) & 63)), F1: 1, F2: (("foo bar" + ("a" + ("" + (v2 + ("foo bar" + v2))))) + ("a" + v2))}
                                                                    _ = v44
                                                                case "foo bar":
                                                                    _ = len(v5)
                                                                }
                                                            case 127:
                                                                if (v2 == "") {
                                                                    v45 := GenS8{F0: 127}
                                                                    _ = v45
                                                                }
                                                            }
                                                        }
                                                        var v46 float32 = -2.25
                                                        _ = fi41
                                                        _ = v46
                                                    }
                                                    if (fi22 == -1) {
                                                        var v47 uint64 = 1
                                                        v5 = append(v5, (-(fi22)))
                                                        if true {
                                                            fmt.Println(v4, v1)
                                                            var v48 string = ("\n\t" + ("\n\t" + "\n\t"))
                                                            _ = v48
                                                        } else {
                                                            if (32767 > 2) {
                                                                v10 = (((v10 - 0.1) / 1.0) - (v10 / -2.25))
                                                                _ = len(v5)
                                                                for fi49 := 0; fi49 < 1; fi49++ {
                                                                    v50 := GenS6{F0: (v47 - v47)}
                                                                    _ = len(v5)
                                                                    _ = fi49
                                                                    _ = v50
                                                                }
                                                            } else {
                                                                switch (fi22 << ((42) & 63)) {
                                                                case -2147483648:
                                                                    v3 = (true || ((v1 ^ v1) <= v1))
                                                                case -128:
                                                                    fmt.Println(v4, v2)
                                                                }
                                                            }
                                                            _ = len(v5)
                                                        }
                                                        _ = v47
                                                    }
                                                }
                                            }
                                            var v51 float64 = v10
                                            for fi52 := 0; fi52 < 6; fi52++ {
                                                v1 = (-(v1))
                                                if v31 {
                                                    _ = len(v5)
                                                    v4 = (v3 && true)
                                                    v10 += 3.14159
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi22 | fi52)
                                                    }
                                                    if (-128 > 1) {
                                                        switch v2 {
                                                        case "a":
                                                            var v53 int16 = int16(fi52)
                                                            _ = v53
                                                        case "foo bar":
                                                            _ = len(v5)
                                                        case "\n\t":
                                                            var v54 int8 = 127
                                                            _ = v54
                                                        default:
                                                            var v55 uint32 = 268
                                                            v51 -= v10
                                                            _ = v55
                                                        }
                                                        v56 := GenS9{F0: 65535, F1: 4294967295, F2: (v1 + 42)}
                                                        _ = v56
                                                    }
                                                }
                                                _ = len(v5)
                                                _ = fi52
                                            }
                                            _ = v51
                                        }
                                    }
                                case -2147483648:
                                    v31 = (0 <= 4294967295)
                                case 127:
                                    v57 := GenS1{F0: (fi22 & -532)}
                                    _ = v57
                                }
                            case "a":
                                var v58 float32 = -2.25
                                _ = v58
                            }
                        case 42:
                            v5 = append(v5, fi22)
                        }
                        var v59 uint16 = 0
                        _ = v59
                    }
                    var v60 uint = 0
                    _ = v31
                    _ = v60
                } else {
                    if v4 {
                        v1 = v1
                    } else {
                        var v61 float32 = 3.14159
                        _ = v61
                    }
                    var v62 uint64 = 18446744073709551615
                    switch v2 {
                    case "hello":
                        v63 := GenS3{F0: (fi22 >> ((65535) & 63)), F1: 763, F2: ("a" + "a")}
                        _ = v63
                    case "":
                        var v64 bool = false
                        _ = v64
                    case "a":
                        for fi65 := 0; fi65 < 8; fi65++ {
                            v66 := GenS10{F0: 0, F1: fi65}
                            _ = fi65
                            _ = v66
                        }
                    }
                    _ = v62
                }
                var v67 uint8 = 255
                _ = v67
            }
            _ = fi22
        }
        v1 += v1
    }
    v1 -= (v1 / (((^(v1))) | 1))
    v68 := GenS11{F0: 118}
    v69 := GenS9{F0: 42, F1: 1, F2: (v1 + (v1 + v1))}
    var v70 float64 = 0.1
    v2 = ("\n\t" + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v70
    _ = v8
    _ = v9
}
