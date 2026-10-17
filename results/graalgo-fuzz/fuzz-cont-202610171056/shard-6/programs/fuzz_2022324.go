package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int64; F1 uint32; F2 int }

func main() {
    var v1 uint = 131
    v2 := make(map[int]uint32)
    var v3 uint = 0
    var v4 uint = 65535
    v5 := make(map[string]string)
    v1 -= 843
    switch 1 {
    case 790:
        v6 := GenS1{F0: 42, F1: true}
        _ = v6
    case 127:
        v1 = (v4 - v4)
    default:
        switch "\n\t" {
        case "hello":
            var v7 bool = (!false)
            _ = v7
        case "\n\t":
            var v8 string = "\n\t"
            _ = v8
        default:
            v9 := GenS2{F0: 0}
            _ = v9
        }
    }
    v4 = v4
    if (true && (18446744073709551615 == 1)) {
        v10 := GenS3{F0: -1}
        switch (("hello" + ("hello" + "\n\t")) + (("hello" + "") + ("foo bar" + ""))) {
        case "\n\t":
            var v11 int16 = 0
            _ = v11
        case "a":
            var v12 int32 = -2147483648
            _ = v12
        case "":
            v4 = (v4 | v1)
        default:
            delete(v2, 42)
        }
        _ = v10
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if true {
            v3 += v3
            if (false || false) {
                switch -1 {
                case -460:
                    var v14 uint8 = 1
                    _ = v14
                default:
                    v3 += v3
                }
            } else {
                var v15 int64 = 1
                _ = v15
            }
        }
        if true {
            for fi16 := 0; fi16 < 13; fi16++ {
                v3 = (v3 >> (((v4 + 599)) & 63))
                var v17 bool = (true || (true || (202 < (v1 & v1))))
                _ = fi16
                _ = v17
            }
            var v18 int16 = -1
            _ = v18
        } else {
            if false {
                v19 := GenS4{F0: 42}
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 uint32 = 37
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        }
        v1 -= v1
        _ = fi13
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if false {
            if ((fi22 >> (((v4 ^ v4)) & 63)) == (fi22 % ((0) | 1))) {
                var v23 bool = true
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 int8 = 100
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                if (!true) {
                    if (((false && false) || (1 != -2147483648)) && false) {
                        for fi26 := 0; fi26 < 18; fi26++ {
                            v1 -= (v4 | (v4 % ((18446744073709551615) | 1)))
                            v27 := []int{(-(fi26)), 0, (fi26 ^ (^(fi22))), (fi26 >> ((v3) & 63))}
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 bool = true
                        var v29 int = fi22
                        _ = v28
                        _ = v29
                    }
                    v30 := []int{-724}
                    _ = v30
                }
                v31 := GenS3{F0: 124}
                if (!(!(true || false))) {
                    var v32 uint8 = 0
                    fmt.Println(v32, v32)
                    _ = v32
                }
                _ = v31
            }
        } else {
            _ = len(v2)
            v33 := GenS3{F0: 127}
            _ = v33
        }
        v34 := []int{(fi22 ^ fi22), fi22}
        _ = fi22
        _ = v34
    }
    var v35 int8 = 127
    v1 += (v3 / ((792) | 1))
    var v36 bool = (!false)
    fmt.Println(v3)
    var v37 uint32 = 2147483648
    if (v36 && (-1 >= 9223372036854775807)) {
        v38 := GenS4{F0: v37}
        var v39 bool = (32767 < 2)
        if (!false) {
            fmt.Println(v35)
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int16 = 0
                v42 := GenS5{F0: 584, F1: 1}
                _ = fi40
                _ = v41
                _ = v42
            }
            var v43 float64 = ((3.14159 - (1e10 * float64(1e10))) / (0.0 - 0.1))
            _ = v43
        } else {
            v36 = ((v35 % ((0) | 1)) == v35)
            for fi44 := 0; fi44 < 6; fi44++ {
                v35 += 100
                _ = fi44
            }
            fmt.Println(v1)
        }
        _ = v38
        _ = v39
    } else {
        if (-1 > -9223372036854775808) {
            v45 := GenS5{F0: 18446744073709551615, F1: 2}
            v37 -= v37
            v36 = (!true)
            _ = v45
        }
        switch 32768 {
        case -1:
            var v46 string = "hello"
            _ = v46
        case 350:
            var v47 int32 = -830
            _ = v47
        case 9223372036854775807:
            if ((float64((0.0 + ((0.0 + 1.5) * float32(3.14159)))) * -2.25) >= (1e10 - 1e10)) {
                var v48 float64 = 0.1
                var v49 int8 = v35
                v35 = v35
                _ = v48
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 17; fi50++ {
            v51 := GenS6{F0: 1}
            fmt.Println(v1)
            _ = fi50
            _ = v51
        }
    }
    if v36 {
        if (1e10 < 3.14159) {
            for fi52 := 0; fi52 < 9; fi52++ {
                v53 := GenS7{F0: v37, F1: 42, F2: 1}
                v36 = v36
                switch "a" {
                case "hello":
                    v54 := GenS3{F0: 42}
                    _ = v54
                case "foo bar":
                    v55 := []int{-9223372036854775808, (^(fi52)), (fi52 % (((fi52 * (fi52 & fi52))) | 1)), fi52}
                    _ = v55
                case "":
                    v1 = 42
                }
                _ = fi52
                _ = v53
            }
        } else {
            v4 -= v4
            for fi56 := 0; fi56 < 14; fi56++ {
                v3 = (v4 | 1)
                _ = fi56
            }
        }
        if v36 {
            v57 := GenS8{F0: -163}
            fmt.Println(v36, v37, v36)
            _ = v57
        }
        switch 0 {
        case 0:
            for fi58 := 0; fi58 < 14; fi58++ {
                fmt.Println(v4)
                delete(v2, (fi58 * 0))
                _ = fi58
            }
        }
    }
    if ((v1 << ((66) & 63)) > v3) {
        if v36 {
            fmt.Println(v4, v1)
            var v59 float32 = float32((float64(1.5) + float64((1.5 * float32(1e10)))))
            v60 := []int{-616}
            _ = v59
            _ = v60
        }
        var v61 int64 = -519
        var v62 int64 = 83
        _ = v61
        _ = v62
    }
    for fi63 := 0; fi63 < 11; fi63++ {
        fmt.Println(fi63, v4, v3)
        _ = fi63
    }
    if false {
        if ((("hello" + "") + ("hello" + "a")) < ("a" + "foo bar")) {
            v3 = v1
            for fi64 := 0; fi64 < 20; fi64++ {
                v65 := GenS4{F0: v37}
                for fi66 := 0; fi66 < 3; fi66++ {
                    v36 = (127 <= 127)
                    delete(v5, ("hello" + ""))
                    if v36 {
                        v1 = v1
                        v36 = (((42 < 887) || ((fi64 >= (^(fi64))) && (661 < 32767))) && v36)
                        var v67 uint = (^(v4))
                        _ = v67
                    }
                    if v36 {
                        var v68 int32 = 0
                        _ = v68
                    }
                    _ = fi66
                }
                v69 := GenS9{F0: -1, F1: (v37 >> ((4294967295) & 63)), F2: fi64}
                _ = fi64
                _ = v65
                _ = v69
            }
            var v70 int = 539
            _ = v70
        }
    } else {
        var v71 int8 = v35
        var v72 string = "hello"
        _ = v71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
