package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint; F2 int16 }

func main() {
    var v1 uint8 = 42
    var v2 int16 = -32768
    v3 := make(map[int]bool)
    if false {
        fmt.Println(v1, v1)
        fmt.Println(v1, v2)
    } else {
        for fi4 := 0; fi4 < 11; fi4++ {
            v5 := GenS1{F0: "a", F1: 520}
            var v6 int64 = 1
            switch ("\n\t" + "hello") {
            case "foo bar":
                var v7 uint16 = 2
                _ = v7
            case "a":
                v8 := GenS2{F0: 2}
                _ = v8
            default:
                var v9 uint = 2
                _ = v9
            }
            _ = fi4
            _ = v5
            _ = v6
        }
        v1 = 1
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS2{F0: 2147483647}
            if ((true && (!(!(0.0 != (3.14159 - (3.14159 - -2.25)))))) && ((0.1 - 1.5) != 3.14159)) {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v2 -= v2
                    fmt.Println(fi12, fi10, v1)
                    _ = fi12
                }
                var v13 int8 = -128
                v1 = (v1 & v1)
                _ = v13
            } else {
                if false {
                    switch "a" {
                    case "foo bar":
                        v14 := GenS2{F0: 42}
                        _ = v14
                    case "a":
                        v15 := GenS2{F0: 0}
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 2; fi16++ {
                    if false {
                        var v17 int64 = -999
                        fmt.Println(v2)
                        v2 += v2
                        _ = v17
                    } else {
                        var v18 uint64 = 42
                        v2 = v2
                        fmt.Println(v18, v1, v1)
                        _ = v18
                    }
                    _ = fi16
                }
                v19 := GenS3{F0: 2147483647, F1: 255, F2: false}
                _ = v19
            }
            v20 := []int{(fi10 / ((fi10) | 1)), (fi10 / ((fi10) | 1))}
            _ = fi10
            _ = v11
            _ = v20
        }
    }
    v21 := GenS3{F0: 65535, F1: 10, F2: false}
    for fi22 := 0; fi22 < 20; fi22++ {
        switch "a" {
        case "foo bar":
            if (491 > 339) {
                switch "\n\t" {
                case "\n\t":
                    v2 += int16(v1)
                case "":
                    var v23 uint32 = 2147483648
                    _ = v23
                case "foo bar":
                    v24 := []int{fi22}
                    _ = v24
                }
                v2 = 1
            } else {
                v25 := []int{fi22, fi22, (-(fi22)), (fi22 + (fi22 * fi22))}
                v1 += (v1 + v1)
                _ = v25
            }
        default:
            v26 := GenS4{F0: ""}
            _ = v26
        }
        var v27 int16 = v2
        _ = fi22
        _ = v27
    }
    if ((!(true && ((!(true && ((v2 | 32767) != int16(v1)))) || ((("hello" + ("\n\t" + "foo bar")) + "\n\t") < "foo bar")))) && true) {
        for fi28 := 0; fi28 < 9; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                switch "\n\t" {
                case "\n\t":
                    v1 = v1
                default:
                    delete(v3, fi28)
                }
                var v30 int = (fi28 * (fi28 % (((fi29 | (^(fi28)))) | 1)))
                _ = fi29
                _ = v30
            }
            if (234 < 9223372036854775807) {
                for fi31 := 0; fi31 < 7; fi31++ {
                    if true {
                        v32 := []int{128, fi31, (fi31 - fi31)}
                        _ = v32
                    } else {
                        var v33 uint8 = (v1 / (((v1 >> ((4294967295) & 63))) | 1))
                        _ = v33
                    }
                    _ = fi31
                }
                fmt.Println(v2, v1, v1)
                switch ("foo bar" + "") {
                case "foo bar":
                    var v34 uint8 = (v1 >> ((692) & 63))
                    _ = v34
                default:
                    switch (fi28 << ((65535) & 63)) {
                    case -437:
                        for fi35 := 0; fi35 < 10; fi35++ {
                            v36 := []int{}
                            _ = fi35
                            _ = v36
                        }
                    case -2147483648:
                        if (!((!(v1 <= v1)) || (!false))) {
                            var v37 uint8 = (v1 << ((427) & 63))
                            var v38 int = (fi28 - 4294967295)
                            _ = v37
                            _ = v38
                        }
                    }
                }
            } else {
                var v39 bool = ((!(652 > -1)) && true)
                v1 += (v1 & 237)
                var v40 int16 = v2
                _ = v39
                _ = v40
            }
            _ = fi28
        }
        var v41 float64 = 1.5
        _ = v41
    } else {
        v42 := []int{32768}
        var v43 int32 = -1
        _ = v42
        _ = v43
    }
    var v44 bool = false
    if v44 {
        var v45 int8 = -128
        var v46 string = "a"
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 int32 = -2147483648
            var v49 int16 = (-(v2))
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v45
        _ = v46
    }
    v44 = ((495 == 0) && true)
    if v44 {
        fmt.Println(v1, v44)
    }
    var v50 int16 = v2
    var v51 int = 1
    if (0 >= -9223372036854775808) {
        if v44 {
            for fi52 := 0; fi52 < 13; fi52++ {
                v51 += (v51 / ((fi52) | 1))
                v53 := GenS1{F0: ("\n\t" + (("foo bar" + "") + "hello")), F1: -2147483648}
                v54 := GenS2{F0: 824}
                _ = fi52
                _ = v53
                _ = v54
            }
            v50 += (^(v50))
            for fi55 := 0; fi55 < 11; fi55++ {
                fmt.Println(v2, v2)
                _ = fi55
            }
        } else {
            v50 = 42
            fmt.Println(v51)
            v56 := []int{(v51 & v51), -467}
            _ = v56
        }
    }
    v50 = v50
    if v44 {
        if false {
            var v57 float64 = float64(1.5)
            _ = v57
        } else {
            if v44 {
                var v58 string = (("hello" + ("foo bar" + "hello")) + (("hello" + (("" + "") + "\n\t")) + "\n\t"))
                fmt.Println(v44, v44)
                var v59 int8 = 100
                _ = v58
                _ = v59
            } else {
                if v44 {
                    v60 := GenS5{F0: (false && true), F1: 184, F2: -1}
                    _ = v60
                }
                switch "\n\t" {
                case "":
                    v44 = v44
                case "a":
                    v61 := GenS4{F0: ("\n\t" + ("\n\t" + "\n\t"))}
                    _ = v61
                case "\n\t":
                    var v62 uint32 = 2147483647
                    _ = v62
                default:
                    v1 -= v1
                }
            }
            v63 := GenS6{F0: 462}
            var v64 bool = (true || v44)
            _ = v63
            _ = v64
        }
        v65 := GenS5{F0: (!v44), F1: 0, F2: -9223372036854775808}
        for fi66 := 0; fi66 < 16; fi66++ {
            for fi67 := 0; fi67 < 18; fi67++ {
                var v68 bool = (0 >= 42)
                if true {
                    var v69 int32 = -1
                    var v70 int32 = v69
                    v71 := GenS6{F0: 945}
                    v72 := GenS4{F0: "foo bar"}
                    _ = v69
                    _ = v70
                    _ = v71
                    _ = v72
                } else {
                    var v73 float64 = float64(0.1)
                    v74 := GenS2{F0: 2147483648}
                    _ = v73
                    _ = v74
                }
                _ = fi67
                _ = v68
            }
            var v75 bool = v44
            delete(v3, (v51 - v51))
            _ = fi66
            _ = v75
        }
        _ = v65
    }
    var v76 int = (v51 - (v51 | -323))
    v77 := GenS7{F0: 2, F1: uint(v1), F2: 370}
    var v78 uint = 489
    switch "a" {
    case "foo bar":
        var v79 int32 = 868
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v44
    _ = v50
    _ = v51
    _ = v76
    _ = v77
    _ = v78
}
