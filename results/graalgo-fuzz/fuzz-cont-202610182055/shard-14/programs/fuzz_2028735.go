package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 uint; F1 int8 }
type GenS11 struct { F0 uint; F1 bool; F2 bool }
type GenS12 struct { F0 bool; F1 int64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int64 }
func genFn0(p0 int, p1 int32, p2 int) uint64 {
    v5 := GenS1{F0: 255, F1: 189}
    for fi6 := 0; fi6 < 6; fi6++ {
        switch p0 {
        case 1:
            if true {
                p1 = -1
                for fi7 := 0; fi7 < 10; fi7++ {
                    if true {
                        if true {
                            var v8 int8 = 0
                            switch (("foo bar" + "") + "hello") {
                            case "foo bar":
                                v9 := GenS2{F0: -621, F1: -9223372036854775808}
                                v10 := []int{142, 27, p2, p0}
                                fmt.Println(p1, p2, p0)
                                _ = v10
                                _ = v9
                            }
                            fmt.Println(fi7, fi7, p1)
                            v11 := []int{fi6, fi6, (fi6 >> ((375) & 63)), 1}
                            _ = v11
                            _ = v8
                        } else {
                            v12 := []int{(fi7 & p2), (fi7 / ((fi7) | 1)), (p0 ^ 0), 4294967295}
                            _ = v12
                        }
                        switch 1 {
                        case 9223372036854775807:
                            fmt.Println(p2, p1, p1)
                            p2 = fi7
                            var v13 string = "a"
                            _ = v13
                        case 42:
                            if true {
                                var v14 float32 = 0.0
                                v15 := GenS2{F0: -9223372036854775808, F1: (p2 ^ p2)}
                                var v16 uint16 = 59
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                        case -789:
                            p0 -= p2
                        }
                    }
                    v17 := []int{p2}
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi7
                    _ = v17
                }
                v18 := []int{}
                _ = v18
            } else {
                p2 -= -9223372036854775808
                for fi19 := 0; fi19 < 17; fi19++ {
                    if true {
                        v20 := GenS2{F0: 0, F1: -954}
                        fmt.Println(fi6, p2)
                        _ = v20
                    } else {
                        var v21 uint = 42
                        fmt.Println(v21, fi19, fi19)
                        _ = v21
                    }
                    var v22 uint32 = 65535
                    _ = fi19
                    _ = v22
                }
            }
        case 0:
            v23 := GenS3{F0: 255, F1: 2147483647, F2: 2147483648}
            _ = v23
        }
        v24 := GenS4{F0: 84, F1: "hello", F2: (("hello" + "foo bar") + ("a" + "a"))}
        _ = fi6
        _ = v24
    }
    v25 := GenS5{F0: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v5
    return 391
}
func genFn1(p0 int16) int16 {
    var v26 uint8 = 127
    for fi27 := 0; fi27 < 19; fi27++ {
        p0 = (-(p0))
        _ = fi27
    }
    p0 += (-(p0))
    _ = p0
    _ = v26
    return p0
}

func main() {
    var v1 int = 2147483647
    v2 := []string{"hello", ("\n\t" + "foo bar"), ("" + ("a" + "")), ("hello" + (("a" + "hello") + "\n\t"))}
    var v3 int64 = 42
    var v4 uint32 = 2
    fmt.Println(v1, v1)
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 int32 = -2147483648
        v4 = (v4 + 870)
        var v30 int = v1
        _ = fi28
        _ = v29
        _ = v30
    }
    v4 = v4
    fmt.Println(v3)
    v1 = -1
    if (2 < 104) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (true || (42 > (v3 / ((0) | 1)))) {
            v1 += int(v3)
            v4 = (v4 << ((2) & 63))
            v1 = v1
        }
    }
    switch v1 {
    case -9223372036854775808:
        switch v1 {
        case -397:
            var v31 uint64 = 721
            _ = v31
        case 128:
            v32 := GenS6{F0: 100}
            _ = v32
        case 9223372036854775807:
            var v33 uint = 4294967295
            _ = v33
        default:
            switch "\n\t" {
            case "hello":
                v34 := GenS4{F0: -2147483648, F1: ("\n\t" + ("foo bar" + "")), F2: "foo bar"}
                _ = v34
            default:
                fmt.Println(v4, v3)
            }
        }
    }
    v35 := GenS7{F0: 0}
    switch "a" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 1; fi37++ {
                if true {
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v38 := GenS8{F0: -379, F1: 42}
                    _ = v38
                } else {
                    var v39 uint8 = 61
                    _ = v39
                }
                _ = fi37
            }
            _ = fi36
        }
    case "foo bar":
        v1 += 0
    }
    if true {
        fmt.Println(v4, v1)
    } else {
        v40 := GenS5{F0: 0}
        _ = v40
    }
    if true {
        v1 = (v1 ^ (v1 / (((v1 / ((-1) | 1))) | 1)))
    } else {
        v1 = (v1 * (v1 % (((v1 / ((-215) | 1))) | 1)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi41 := 0; fi41 < 16; fi41++ {
            switch (v1 + 32767) {
            case 32767:
                v42 := GenS8{F0: 128, F1: 2}
                _ = v42
            case -264:
                v2 = append(v2, ("foo bar" + "hello"))
            case 32768:
                for fi43 := 0; fi43 < 14; fi43++ {
                    v4 += 2147483647
                    switch ("\n\t" + ("hello" + ("hello" + "hello"))) {
                    case "":
                        var v44 bool = true
                        _ = v44
                    case "\n\t":
                        v1 = -9223372036854775808
                    }
                    var v45 uint16 = 32767
                    _ = fi43
                    _ = v45
                }
            }
            v4 = v4
            _ = fi41
        }
    }
    var v46 uint32 = 630
    for fi47 := 0; fi47 < 1; fi47++ {
        switch (-(v1)) {
        case 42:
            v48 := GenS9{F0: 42, F1: 796}
            _ = v48
        case 127:
            if (2147483647 >= 570) {
                v2 = append(v2, ("\n\t" + "hello"))
                var v49 bool = true
                var v50 string = "hello"
                _ = v49
                _ = v50
            }
        default:
            if true {
                if false {
                    var v51 uint8 = 128
                    switch v1 {
                    case 1:
                        for fi52 := 0; fi52 < 6; fi52++ {
                            v53 := GenS10{F0: 2, F1: -93}
                            if false {
                                var v54 bool = false
                                _ = v54
                            } else {
                                if ((^(v3)) > (v3 >> ((2) & 63))) {
                                    if (true || true) {
                                        v46 = v46
                                        v55 := GenS6{F0: -59}
                                        switch (v1 - (fi52 + 128)) {
                                        case -980:
                                            fmt.Println(v1, v3)
                                        default:
                                            v1 += (v1 >> ((42) & 63))
                                        }
                                        _ = v55
                                    } else {
                                        fmt.Println(v46, v51, v1)
                                        v56 := GenS11{F0: 206, F1: true, F2: (!false)}
                                        _ = v56
                                    }
                                    v57 := GenS1{F0: v51, F1: 749}
                                    switch "" {
                                    case "":
                                        v46 = (v4 >> ((255) & 63))
                                    }
                                    _ = v57
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    fmt.Println(fi47)
                                }
                            }
                            var v58 string = "a"
                            _ = fi52
                            _ = v53
                            _ = v58
                        }
                    case -1:
                        v46 -= v4
                    }
                    _ = v51
                } else {
                    v59 := GenS12{F0: ((v4 >> ((255) & 63)) <= v4), F1: 9223372036854775807}
                    var v60 int8 = 42
                    if ((0.0 * (3.14159 - 1.5)) >= (float64(((1.5 / float32((1.5 - (((0.1 * (-2.25 * 0.0)) + 1.5) / 1e10)))) - -2.25)) - 3.14159)) {
                        var v61 int16 = 1
                        _ = v61
                    }
                    _ = v59
                    _ = v60
                }
                v1 += (v1 + (fi47 / ((fi47) | 1)))
                switch "a" {
                case "a":
                    var v62 int32 = 279
                    _ = v62
                default:
                    v3 -= (v3 * v3)
                }
            }
        }
        var v63 bool = (true && true)
        if false {
            if (-2147483648 >= -2147483648) {
                for fi64 := 0; fi64 < 1; fi64++ {
                    var v65 uint64 = 18446744073709551615
                    fmt.Println(v65)
                    if len(v2) > 0 {
                        v2[0] = ("a" + "foo bar")
                    }
                    _ = fi64
                    _ = v65
                }
                var v66 bool = false
                _ = v66
            }
            for fi67 := 0; fi67 < 13; fi67++ {
                switch (fi47 >> ((42) & 63)) {
                case -875:
                    var v68 uint8 = 1
                    _ = v68
                }
                var v69 uint = 1
                v70 := GenS1{F0: 2, F1: 42}
                var v71 float32 = 0.0
                _ = fi67
                _ = v69
                _ = v70
                _ = v71
            }
            for fi72 := 0; fi72 < 18; fi72++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v73 uint16 = 388
                if (!false) {
                    v74 := GenS1{F0: 17, F1: 65535}
                    _ = v74
                } else {
                    v75 := GenS13{F0: 746}
                    _ = v75
                }
                _ = fi72
                _ = v73
            }
        } else {
            for fi76 := 0; fi76 < 17; fi76++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
                var v77 uint64 = 42
                _ = fi76
                _ = v77
            }
            var v78 int16 = 32767
            switch v1 {
            case 2147483647:
                v4 += (v46 * (v46 + (v4 & 2)))
            case 32768:
                v79 := GenS14{F0: v3}
                _ = v79
            case -651:
                var v80 uint64 = 0
                _ = v80
            default:
                v63 = v63
            }
            _ = v78
        }
        _ = fi47
        _ = v63
    }
    fmt.Println(v3)
    _ = len(v2)
    fmt.Println(v4)
    var v81 int8 = 23
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v81
}
