package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0() uint8 {
    var v4 int32 = -16
    v5 := []int{1, -466}
    var v6 uint8 = 99
    var v7 int8 = -1
    fmt.Println(v4, v7, v7)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v6 * (v6 & 1))
}
func genFn1(p0 uint64, p1 uint32) uint8 {
    for fi8 := 0; fi8 < 6; fi8++ {
        p0 = (p0 ^ p0)
        v9 := GenS1{F0: (p0 * p0), F1: 32767, F2: true}
        v10 := GenS1{F0: p0, F1: -567, F2: true}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 65535, F1: -2147483648}
    _ = p0
    _ = p1
    _ = v11
    return genFn0()
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 int = 2147483647
    v3 = v3
    for fi12 := 0; fi12 < 7; fi12++ {
        delete(v1, 448)
        var v13 int16 = 0
        if false {
            fmt.Println(fi12)
            switch v3 {
            case 2147483647:
                var v14 float32 = (((0.1 / -2.25) + 3.14159) + float32(3.14159))
                _ = v14
            default:
                v3 -= fi12
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: 1, F1: 2147483647}
                _ = fi15
                _ = v16
            }
        }
        _ = fi12
        _ = v13
    }
    v17 := GenS1{F0: 4294967295, F1: 791, F2: true}
    fmt.Println(v3)
    v18 := GenS1{F0: 2, F1: 42, F2: (65535 != 1)}
    v19 := GenS1{F0: 18446744073709551615, F1: 42, F2: true}
    if (0.0 <= 0.0) {
        switch (v3 - 9223372036854775807) {
        case 32767:
            for fi20 := 0; fi20 < 19; fi20++ {
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := []int{}
                    v23 := GenS1{F0: 2, F1: 621, F2: true}
                    var v24 float32 = (((-2.25 * (-2.25 / 3.14159)) - -2.25) / 0.1)
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v25 := []int{}
                switch v3 {
                case -2147483648:
                    fmt.Println(fi20, v3)
                case 863:
                    v3 -= 163
                case -501:
                    var v26 float64 = 0.0
                    _ = v26
                }
                _ = fi20
                _ = v25
            }
        case 0:
            var v27 int = (v3 + 32768)
            _ = v27
        }
    } else {
        if (!(1 <= 0)) {
            for fi28 := 0; fi28 < 2; fi28++ {
                for fi29 := 0; fi29 < 18; fi29++ {
                    v30 := GenS1{F0: 1, F1: 1000, F2: true}
                    v3 = (fi28 % ((fi28) | 1))
                    _ = fi29
                    _ = v30
                }
                v3 = v3
                for fi31 := 0; fi31 < 10; fi31++ {
                    v3 = 103
                    _ = fi31
                }
                _ = fi28
            }
        } else {
            v32 := []int{}
            switch ("hello" + "a") {
            case "hello":
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 float32 = -2.25
                    var v35 float32 = (v34 - 1.5)
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case "":
                _ = len(v32)
            }
            _ = v32
        }
        v3 = v3
    }
    switch v3 {
    case -630:
        for fi36 := 0; fi36 < 7; fi36++ {
            v3 -= v3
            _ = fi36
        }
    case -794:
        _ = len(v2)
    }
    v37 := GenS1{F0: 18446744073709551615, F1: 32767, F2: (-547 >= 1)}
    var v38 uint = 0
    v3 += -1
    var v39 uint16 = 0
    if ((0 > -168) && (false && (4294967295 > 18446744073709551615))) {
        switch (v3 & (v3 >> ((v38) & 63))) {
        case 71:
            v40 := GenS1{F0: 4294967295, F1: -32768, F2: false}
            _ = v40
        case -1:
            if false {
                v41 := GenS2{F0: 42, F1: 0}
                _ = v41
            } else {
                if true {
                    fmt.Println(v3)
                    for fi42 := 0; fi42 < 18; fi42++ {
                        var v43 int32 = -778
                        _ = fi42
                        _ = v43
                    }
                } else {
                    var v44 uint16 = 32767
                    v38 -= 4294967295
                    var v45 int16 = 72
                    _ = v44
                    _ = v45
                }
                v2[670] = "hello"
            }
        case 4294967295:
            v46 := []int{v3, (v3 % ((-715) | 1))}
            _ = v46
        }
        if true {
            v3 = (v3 ^ v3)
        } else {
            switch "foo bar" {
            case "":
                v38 += (v38 >> ((42) & 63))
            case "hello":
                if false {
                    v47 := GenS1{F0: 42, F1: -1, F2: (54 == 67)}
                    _ = v47
                } else {
                    var v48 uint32 = 248
                    for fi49 := 0; fi49 < 7; fi49++ {
                        v50 := []int{(fi49 + (fi49 >> ((65535) & 63))), (v3 * 789), int(v39), fi49}
                        _ = fi49
                        _ = v50
                    }
                    v51 := []int{v3, v3, v3}
                    _ = v48
                    _ = v51
                }
            default:
                v3 = (v3 - -191)
            }
            for fi52 := 0; fi52 < 15; fi52++ {
                var v53 uint32 = 862
                _ = fi52
                _ = v53
            }
            v38 += (v38 >> ((v38) & 63))
        }
        switch "\n\t" {
        case "\n\t":
            if (v39 > v39) {
                if true {
                    var v54 float32 = float32(0.0)
                    for fi55 := 0; fi55 < 4; fi55++ {
                        v39 -= (^(v39))
                        _ = fi55
                    }
                    _ = v54
                } else {
                    fmt.Println(v38)
                    fmt.Println(v38, v38, v39)
                }
                fmt.Println(v3, v38)
                for fi56 := 0; fi56 < 11; fi56++ {
                    v57 := GenS2{F0: 2147483647, F1: 1}
                    _ = fi56
                    _ = v57
                }
            } else {
                v38 -= (v38 % (((v38 % ((18446744073709551615) | 1))) | 1))
            }
        case "a":
            if true {
                if true {
                    for fi58 := 0; fi58 < 17; fi58++ {
                        v39 = (^(v39))
                        v59 := GenS3{F0: 796, F1: 121}
                        _ = fi58
                        _ = v59
                    }
                } else {
                    for fi60 := 0; fi60 < 8; fi60++ {
                        v3 -= fi60
                        v39 -= 32767
                        _ = fi60
                    }
                }
                if (42 >= -2147483648) {
                    v61 := v1[v3]
                    _ = v61
                }
            } else {
                v38 -= v38
                v62 := []int{}
                delete(v2, 1)
                _ = v62
            }
        case "hello":
            if (1 < 127) {
                v39 = v39
            } else {
                v63 := []int{v3, (v3 ^ v3), v3, v3}
                for fi64 := 0; fi64 < 1; fi64++ {
                    v38 -= v38
                    _ = fi64
                }
                switch "" {
                case "a":
                    v39 += v39
                default:
                    var v65 int32 = 1
                    _ = v65
                }
                _ = v63
            }
        default:
            if false {
                _ = len(v2)
                var v66 int16 = 1000
                var v67 uint32 = 2147483648
                _ = v66
                _ = v67
            } else {
                switch "" {
                case "\n\t":
                    v1[1] = (!false)
                case "":
                    var v68 int16 = 0
                    _ = v68
                case "hello":
                    for fi69 := 0; fi69 < 5; fi69++ {
                        v70 := GenS1{F0: 2, F1: 1000, F2: true}
                        var v71 float32 = ((3.14159 * 3.14159) - 1e10)
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                }
                if (42 <= 100) {
                    if false {
                        v72 := []int{v3, v3, -1, v3}
                        v38 = (v38 >> ((v38) & 63))
                        _ = v72
                    }
                    fmt.Println(v38, v3, v39)
                    v73 := GenS1{F0: 4294967295, F1: 0, F2: true}
                    _ = v73
                } else {
                    fmt.Println(v39, v3)
                    v3 -= 32767
                    delete(v1, 42)
                }
                fmt.Println(v3)
            }
        }
    } else {
        var v74 int32 = 2147483647
        v74 = v74
        _ = v74
    }
    if (!false) {
        v75 := GenS4{F0: 142, F1: int32(v38), F2: 1}
        v39 -= v39
        _ = v75
    } else {
        v39 = 42
    }
    v76 := GenS5{F0: 611, F1: 4294967295}
    switch "foo bar" {
    case "hello":
        var v77 float32 = float32(float64(float32(1e10)))
        _ = v77
    case "foo bar":
        var v78 float64 = 0.1
        _ = v78
    case "\n\t":
        var v79 bool = (!((((0.1 != 3.14159) || (0.1 == 3.14159)) && true) && true))
        _ = v79
    }
    v1[0] = (v39 >= (^(v39)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v76
}
