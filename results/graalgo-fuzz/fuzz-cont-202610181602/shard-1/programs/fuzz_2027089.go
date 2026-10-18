package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 string, p2 int16) int16 {
    v6 := []int{32768}
    var v7 int = -326
    var v8 uint32 = 342
    fmt.Println(p0, p1)
    if len(v6) > 0 {
        v6[0] = 32768
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return int16(p0)
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]bool)
    var v3 uint16 = 65535
    v4 := []int{-197, 0, 399}
    var v5 uint16 = v3
    _ = len(v4)
    var v9 uint = 255
    if true {
        v3 += v3
        var v10 uint8 = 57
        _ = v10
    }
    v11 := GenS1{F0: 1}
    v12 := GenS2{F0: 234, F1: v3}
    if (true && true) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            if false {
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 bool = (741 == -897)
                    _ = fi14
                    _ = v15
                }
                var v16 uint64 = 730
                _ = v16
            } else {
                fmt.Println(v9, v3, fi13)
                v17 := GenS1{F0: 100}
                _ = v17
            }
            v3 += 2
            _ = fi13
        }
    }
    var v18 string = ("\n\t" + "foo bar")
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS1{F0: -24}
        if true {
            fmt.Println(v3)
            fmt.Println(fi19, fi19)
        } else {
            var v21 uint64 = 18446744073709551615
            v9 = (v9 + v9)
            switch 128 {
            case 0:
                v5 += 32767
            case -2147483648:
                var v22 uint16 = (^(v3))
                _ = v22
            }
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if false {
        var v23 uint8 = 131
        v9 = v9
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 bool = true
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        var v26 int8 = 127
        if true {
            if (42 <= 127) {
                v27 := GenS3{F0: 1}
                _ = v27
            }
            for fi28 := 0; fi28 < 11; fi28++ {
                if ((22 == 2147483647) && ("a" < (v18 + v18))) {
                    if (!false) {
                        switch (fi28 << ((v9) & 63)) {
                        case -297:
                            fmt.Println(v3, v5)
                        }
                    }
                    if ((1000 <= 1) || false) {
                        v1 -= 3.14159
                    }
                } else {
                    switch fi28 {
                    case 9223372036854775807:
                        fmt.Println(v3, fi28)
                    }
                    if (false || (v18 >= "foo bar")) {
                        if false {
                            for fi29 := 0; fi29 < 17; fi29++ {
                                var v30 float64 = -2.25
                                v30 = 1e10
                                v31 := GenS2{F0: 0, F1: (v5 - 65535)}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        } else {
                            v26 -= v26
                        }
                    } else {
                        v26 = v26
                        for fi32 := 0; fi32 < 12; fi32++ {
                            var v33 bool = (128 <= 0)
                            v34 := v2[894]
                            var v35 uint8 = 255
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                    fmt.Println(fi28)
                }
                _ = fi28
            }
            var v36 int8 = v26
            _ = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 uint16 = (v5 | (v3 & v5))
            if false {
                v38 = (v5 | (v5 ^ v3))
                for fi39 := 0; fi39 < 1; fi39++ {
                    v1 = 0.0
                    _ = fi39
                }
            } else {
                if (((v9 == 255) && (!(false && (0.0 >= -2.25)))) || true) {
                    if (!(!true)) {
                        for fi40 := 0; fi40 < 11; fi40++ {
                            var v41 int32 = 42
                            v42 := GenS1{F0: (-(v26))}
                            v38 += (v38 % ((v38) | 1))
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        v43 := GenS2{F0: 1, F1: 32767}
                        fmt.Println(fi37, v5)
                        _ = v43
                    }
                } else {
                    v26 -= (v26 >> ((v9) & 63))
                    for fi44 := 0; fi44 < 16; fi44++ {
                        if false {
                            v4 = append(v4, -208)
                            if ((!(-1 <= 221)) || false) {
                                if true {
                                    switch v18 {
                                    case "foo bar":
                                        v38 -= 32767
                                    }
                                    v9 -= (v9 & 255)
                                } else {
                                    v2[4294967295] = (!true)
                                }
                                if false {
                                    v38 += v5
                                } else {
                                    var v45 int32 = -2147483648
                                    _ = len(v2)
                                    v46 := GenS1{F0: v26}
                                    _ = v45
                                    _ = v46
                                }
                                var v47 int16 = -32768
                                _ = v47
                            }
                            switch ("hello" + v18) {
                            case "\n\t":
                                var v48 bool = (128 != 42)
                                _ = v48
                            }
                        }
                        switch v18 {
                        case "hello":
                            v38 -= v38
                        default:
                            v49 := GenS4{F0: (v9 << ((v9) & 63)), F1: 9223372036854775807}
                            _ = v49
                        }
                        switch (v18 + "foo bar") {
                        case "\n\t":
                            fmt.Println(v3)
                        case "":
                            v3 -= v3
                            if ((v9 + uint(v3)) == (v9 - (v9 % ((v9) | 1)))) {
                                var v50 uint32 = 685
                                if false {
                                    if (false && false) {
                                        v26 += v26
                                        var v51 uint32 = (^(v50))
                                        delete(v2, (v50 * v51))
                                        _ = v51
                                    }
                                    v1 += v1
                                    v18 = ("hello" + "a")
                                    switch -9223372036854775808 {
                                    case -1:
                                        delete(v2, 409)
                                    case 32768:
                                        v52 := GenS4{F0: v9, F1: fi37}
                                        _ = v52
                                    case 42:
                                        switch fi44 {
                                        case 578:
                                            v53 := v2[(v50 * v50)]
                                            _ = v53
                                        case -414:
                                            var v54 uint8 = uint8(v50)
                                            _ = v54
                                        }
                                    }
                                    delete(v2, 239)
                                    v50 = 65535
                                } else {
                                    fmt.Println(v5, v38, fi37)
                                    var v55 uint16 = (^(v38))
                                    _ = len(v2)
                                    var v56 uint64 = 0
                                    _ = v55
                                    _ = v56
                                }
                                fmt.Println(v18)
                                fmt.Println(v1)
                                var v57 string = "a"
                                v5 -= (v38 | 65535)
                                v57 = (v57 + (("a" + ("hello" + (v18 + "\n\t"))) + v18))
                                fmt.Println(v26)
                                fmt.Println(fi44)
                                if false {
                                    if ((^(v26)) == -1) {
                                        var v58 string = ((((v57 + v57) + v57) + v18) + ("hello" + "a"))
                                        switch fi44 {
                                        case -1:
                                            var v59 int32 = 2147483647
                                            _ = v59
                                        case -363:
                                            if true {
                                                v1 -= v1
                                                _ = len(v2)
                                                switch v18 {
                                                case "":
                                                    fmt.Println(v38)
                                                case "foo bar":
                                                    var v60 uint64 = 2
                                                    _ = v60
                                                }
                                            }
                                        case 2147483647:
                                            fmt.Println(v1, fi44, v1)
                                        }
                                        v5 = (^(v38))
                                        _ = v58
                                    } else {
                                        var v61 uint32 = uint32(v5)
                                        v50 = (v61 >> ((v9) & 63))
                                        if (!false) {
                                            v62 := GenS3{F0: 255}
                                            if false {
                                                if ((v26 * (-(v26))) > (^(v26))) {
                                                    var v63 int16 = 237
                                                    v4 = append(v4, (fi37 | fi37))
                                                    v64 := GenS2{F0: 57, F1: (v5 | (^(v3)))}
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    if ((((v1 - v1) * (float64(1.5) / (((v1 * (v1 / v1)) * float64(0.0)) / ((v1 - 0.0) - -2.25)))) == (-2.25 / (v1 + v1))) || (0 >= -1)) {
                                                        v38 += 1
                                                    }
                                                    var v65 uint8 = 0
                                                    _ = v65
                                                }
                                                v57 = "hello"
                                                switch v57 {
                                                case "a":
                                                    fmt.Println(v5, v9)
                                                case "\n\t":
                                                    v66 := GenS4{F0: (^(v9)), F1: (fi44 ^ int(v9))}
                                                    _ = v66
                                                case "foo bar":
                                                    v3 -= (v3 - 0)
                                                }
                                                v67 := GenS5{F0: "hello"}
                                                if len(v4) > 0 {
                                                    v4[0] = (fi37 * 572)
                                                }
                                                _ = v67
                                            } else {
                                                var v68 uint32 = v50
                                                _ = v68
                                            }
                                            _ = v62
                                        } else {
                                            fmt.Println(v26)
                                        }
                                        _ = v61
                                    }
                                }
                                _ = v50
                                _ = v57
                            }
                            var v69 uint16 = (v3 + 1)
                            _ = v69
                        }
                        _ = fi44
                    }
                    v26 += -1
                }
                if false {
                    delete(v2, 1)
                    v70 := GenS6{F0: uint8(v5), F1: 0}
                    v9 += 109
                    _ = v70
                }
            }
            _ = fi37
            _ = v38
        }
        _ = v26
    }
    var v71 float32 = ((1.5 - -2.25) * 0.0)
    v3 = v5
    _ = len(v4)
    var v72 string = (v18 + "")
    var v73 uint64 = 42
    v74 := GenS1{F0: -1}
    switch -1 {
    case 2:
        var v75 uint16 = v3
        _ = v75
    default:
        for fi76 := 0; fi76 < 19; fi76++ {
            for fi77 := 0; fi77 < 15; fi77++ {
                var v78 uint64 = v73
                _ = fi77
                _ = v78
            }
            if false {
                var v79 float64 = (v1 * v1)
                fmt.Println(v9, v3)
                _ = v79
            }
            for fi80 := 0; fi80 < 16; fi80++ {
                v3 = v5
                v9 -= 0
                fmt.Println(v73, v5, v1)
                _ = fi80
            }
            _ = fi76
        }
    }
    for fi81 := 0; fi81 < 12; fi81++ {
        for fi82 := 0; fi82 < 18; fi82++ {
            v2[0] = false
            v18 = ""
            _ = fi82
        }
        _ = fi81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v9
}
