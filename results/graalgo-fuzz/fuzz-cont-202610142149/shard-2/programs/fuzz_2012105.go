package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint16; F2 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string; F1 string; F2 uint }
type GenS11 struct { F0 uint8 }

func main() {
    v1 := []string{("hello" + (("" + ("foo bar" + "a")) + "")), "foo bar", ""}
    var v2 string = ((("\n\t" + "") + (("a" + "a") + ("\n\t" + ("" + "")))) + "\n\t")
    var v3 string = ("foo bar" + v2)
    var v4 int8 = 0
    v5 := []uint32{216}
    fmt.Println(v3)
    v2 = v3
    v4 += 42
    var v6 bool = true
    v4 += v4
    v5 = append(v5, 42)
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 132, F1: -1, F2: 1}
        v4 = v4
        _ = fi7
        _ = v8
    }
    if v6 {
        for fi9 := 0; fi9 < 10; fi9++ {
            v6 = v6
            for fi10 := 0; fi10 < 12; fi10++ {
                var v11 bool = ((!(10 != 42)) && (0 < 1))
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        if (308 < -1) {
            v12 := GenS1{F0: 1, F1: 9223372036854775807, F2: 2}
            v2 = v2
            for fi13 := 0; fi13 < 7; fi13++ {
                switch (-(fi13)) {
                case -2147483648:
                    var v14 uint64 = 1
                    _ = v14
                default:
                    var v15 uint8 = 2
                    _ = v15
                }
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v3, v3)
                case "":
                    fmt.Println(v3)
                case "a":
                    var v16 int = (fi13 << ((890) & 63))
                    _ = v16
                default:
                    v3 = v3
                }
                _ = fi13
            }
            _ = v12
        } else {
            v4 = (v4 >> ((784) & 63))
            var v17 int64 = 0
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 string = (v2 + v2)
                if ((v17 % (((v17 * (v17 ^ v17))) | 1)) <= (-(v17))) {
                    _ = len(v5)
                    v6 = (!(105 <= 18446744073709551615))
                    if v6 {
                        if false {
                            v20 := GenS2{F0: 255, F1: 434, F2: 0}
                            v17 = v17
                            _ = v20
                        } else {
                            v21 := GenS1{F0: 65535, F1: v17, F2: 1}
                            if (1 <= 1) {
                                v22 := GenS3{F0: 32767, F1: -217, F2: 1}
                                v4 += -128
                                if (2 > 647) {
                                    if true {
                                        for fi23 := 0; fi23 < 3; fi23++ {
                                            fmt.Println(v19, v2)
                                            _ = fi23
                                        }
                                    }
                                    v24 := GenS2{F0: 212, F1: 2147483647, F2: 2147483647}
                                    v4 -= 1
                                    _ = v24
                                } else {
                                    v25 := GenS4{F0: (v3 + v19)}
                                    var v26 uint64 = 42
                                    v19 = ""
                                    _ = v25
                                    _ = v26
                                }
                                _ = v22
                            } else {
                                v27 := GenS3{F0: 0, F1: 1000, F2: 0}
                                v28 := GenS5{F0: 1}
                                v19 = ((v19 + v19) + "\n\t")
                                _ = v27
                                _ = v28
                            }
                            switch -9223372036854775808 {
                            case 9223372036854775807:
                                if v6 {
                                    fmt.Println(v17)
                                    for fi29 := 0; fi29 < 2; fi29++ {
                                        switch v2 {
                                        case "foo bar":
                                            fmt.Println(v6, fi18)
                                        default:
                                            v19 = "\n\t"
                                            v2 = ""
                                            switch ("hello" + "\n\t") {
                                            case "":
                                                v5 = append(v5, 2147483647)
                                            case "foo bar":
                                                switch -2147483648 {
                                                case 42:
                                                    v30 := GenS6{F0: 2}
                                                    _ = v30
                                                case -9223372036854775808:
                                                    var v31 float64 = -2.25
                                                    if v6 {
                                                        v3 = "\n\t"
                                                        var v32 int = fi29
                                                        v33 := GenS7{F0: v2, F1: 2, F2: v4}
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    if v6 {
                                                        var v34 int16 = 1000
                                                        if len(v5) > 0 {
                                                            v5[0] = 348
                                                        }
                                                        var v35 int = fi18
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        v36 := GenS8{F0: 176}
                                                        _ = v36
                                                    }
                                                    v37 := GenS9{F0: 886}
                                                    _ = v31
                                                    _ = v37
                                                case -476:
                                                    v38 := GenS8{F0: 127}
                                                    _ = v38
                                                }
                                            case "a":
                                                if ((1 != 2) && (v17 == v17)) {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v3 = ("a" + ("hello" + (v2 + "foo bar")))
                                                    case "a":
                                                        fmt.Println(v19)
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                } else {
                                                    if (691 >= -2147483648) {
                                                        var v39 uint = 4294967295
                                                        var v40 uint32 = 209
                                                        v41 := GenS10{F0: v3, F1: ("foo bar" + "foo bar"), F2: (^(v39))}
                                                        _ = v39
                                                        _ = v40
                                                        _ = v41
                                                    } else {
                                                        _ = len(v1)
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v42 uint16 = 2
                                                        _ = v42
                                                    }
                                                    if (127 <= 255) {
                                                        switch (fi18 * -9223372036854775808) {
                                                        case 625:
                                                            v43 := GenS11{F0: 0}
                                                            _ = v43
                                                        case -507:
                                                            var v44 uint = 204
                                                            _ = v44
                                                        case 4294967295:
                                                            v4 -= v4
                                                        }
                                                    } else {
                                                        _ = len(v1)
                                                        _ = len(v1)
                                                    }
                                                    var v45 uint32 = 716
                                                    _ = v45
                                                }
                                            }
                                            var v46 float32 = 0.1
                                            _ = v46
                                        }
                                        _ = fi29
                                    }
                                    v6 = (845 == 421)
                                } else {
                                    fmt.Println(v4, v17)
                                    v47 := GenS7{F0: (v3 + ("foo bar" + ((v19 + "a") + (("hello" + ("a" + "foo bar")) + "foo bar")))), F1: 200, F2: v4}
                                    v3 = "\n\t"
                                    _ = v47
                                }
                            case 127:
                                fmt.Println(fi18, v3)
                            case -233:
                                var v48 float32 = 0.1
                                _ = v48
                            default:
                                var v49 int16 = 32767
                                _ = v49
                            }
                            _ = v21
                        }
                    } else {
                        for fi50 := 0; fi50 < 4; fi50++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            v6 = (42 == 32767)
                            _ = fi50
                        }
                    }
                } else {
                    var v51 int = 42
                    _ = v51
                }
                for fi52 := 0; fi52 < 18; fi52++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v6 = v6
                    v2 = v2
                    var v53 uint16 = 2
                    v5 = append(v5, 0)
                    _ = fi52
                    _ = v53
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
