package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int32; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 1
    var v3 float64 = 1.5
    if ((-1 == -1) || (-671 > -32768)) {
        var v4 int32 = -1
        var v5 int8 = -104
        for fi6 := 0; fi6 < 11; fi6++ {
            v4 -= 825
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 float64 = v3
                var v9 float32 = (1.5 + 1.5)
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
        _ = v4
        _ = v5
    } else {
        v2 += v2
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{(fi10 + 32767), 0, (fi10 << ((65535) & 63))}
        switch "foo bar" {
        case "\n\t":
            var v12 int32 = -1
            _ = v12
        case "":
            var v13 uint64 = 941
            _ = v13
        default:
            var v14 string = "\n\t"
            _ = v14
        }
        var v15 uint16 = 573
        _ = fi10
        _ = v11
        _ = v15
    }
    if false {
        v2 -= v2
        var v16 uint8 = 98
        _ = v16
    }
    switch 9223372036854775807 {
    case 990:
        switch ("hello" + "hello") {
        case "\n\t":
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 uint8 = 132
                v19 := GenS1{F0: false, F1: 4294967295, F2: false}
                _ = fi17
                _ = v18
                _ = v19
            }
        default:
            switch "foo bar" {
            case "a":
                v20 := []int{128, 203, 32768}
                _ = v20
            case "\n\t":
                v21 := GenS1{F0: (!(0 > 0)), F1: 42, F2: false}
                _ = v21
            default:
                for fi22 := 0; fi22 < 8; fi22++ {
                    v2 += (v2 / ((v2) | 1))
                    _ = fi22
                }
            }
        }
    case -256:
        v23 := []int{1, 42}
        _ = v23
    case -9223372036854775808:
        v24 := GenS2{F0: 1, F1: (true && (550 == 397))}
        _ = v24
    }
    if true {
        for fi25 := 0; fi25 < 11; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 int64 = 9223372036854775807
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                if true {
                    switch ("a" + ("foo bar" + "\n\t")) {
                    case "hello":
                        v29 := []int{-133, 4294967295, (fi28 | int(v2))}
                        fmt.Println(v2)
                        _ = v29
                    case "":
                        if ((1e10 / 0.1) == 1.5) {
                            var v30 int16 = 0
                            v2 = (v2 & -1)
                            v31 := GenS2{F0: 985, F1: ((1.5 != (v3 * v3)) && (true || true))}
                            v32 := GenS2{F0: 120, F1: (false && true)}
                            v3 -= v3
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    case "a":
                        if true {
                            v2 = v2
                            if true {
                                var v33 uint = 18446744073709551615
                                fmt.Println(v3)
                                switch "foo bar" {
                                case "\n\t":
                                    delete(v1, ("a" + ("\n\t" + ("" + "\n\t"))))
                                case "":
                                    var v34 bool = (v33 <= v33)
                                    _ = v34
                                case "foo bar":
                                    v35 := GenS3{F0: (fi28 << (((v33 ^ v33)) & 63))}
                                    _ = v35
                                }
                                _ = v33
                            }
                            v36 := []int{fi25, (fi25 * (fi25 ^ (fi28 >> ((18446744073709551615) & 63)))), fi28}
                            _ = v36
                        }
                    }
                }
                v37 := GenS2{F0: 136, F1: true}
                _ = fi28
                _ = v37
            }
            var v38 uint16 = 42
            _ = fi25
            _ = v38
        }
        if (255 >= 417) {
            if (0 <= 176) {
                var v39 string = ""
                for fi40 := 0; fi40 < 12; fi40++ {
                    if (!true) {
                        v41 := GenS3{F0: fi40}
                        fmt.Println(v3, v2)
                        var v42 uint64 = 981
                        _ = v41
                        _ = v42
                    } else {
                        v43 := GenS2{F0: 18446744073709551615, F1: (506 < 65535)}
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 10; fi44++ {
                        v45 := GenS4{F0: 0, F1: 1}
                        v39 = ((v39 + ((("hello" + (("a" + "a") + v39)) + ("foo bar" + "a")) + v39)) + "a")
                        v2 -= v2
                        switch fi40 {
                        case -218:
                            if false {
                                var v46 int8 = 0
                                v47 := GenS2{F0: 42, F1: false}
                                var v48 bool = false
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                        case 298:
                            fmt.Println(fi40)
                        case -2147483648:
                            v39 = ("\n\t" + (v39 + "a"))
                        }
                        _ = fi44
                        _ = v45
                    }
                    _ = fi40
                }
                _ = v39
            } else {
                if false {
                    var v49 int16 = 447
                    v49 = 563
                    var v50 uint16 = 2
                    _ = v49
                    _ = v50
                } else {
                    v3 -= (v3 - (v3 / v3))
                }
                v51 := []int{9223372036854775807, 0, 892, -9223372036854775808}
                for fi52 := 0; fi52 < 2; fi52++ {
                    v53 := GenS4{F0: 2, F1: 0}
                    if (!false) {
                        v54 := GenS1{F0: true, F1: 342, F2: (!true)}
                        var v55 float32 = (0.0 * 3.14159)
                        _ = v54
                        _ = v55
                    } else {
                        if len(v51) > 0 {
                            v51[0] = (fi52 + -2147483648)
                        }
                    }
                    fmt.Println(v2, v2)
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
            v3 = v3
        } else {
            v2 -= (v2 << ((4294967295) & 63))
            for fi56 := 0; fi56 < 13; fi56++ {
                var v57 float64 = ((v3 - float64(((1e10 - 0.0) * float32(-2.25)))) * 3.14159)
                var v58 uint64 = 4294967295
                var v59 float64 = v57
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
            }
        }
        switch 72 {
        case -2147483648:
            var v60 int16 = 1
            _ = v60
        case 2147483647:
            v2 += (v2 << ((419) & 63))
        }
    }
    var v61 float64 = (((v3 + (v3 * 0.1)) + v3) - ((v3 + 0.1) - (float64(-2.25) / v3)))
    switch "hello" {
    case "":
        v1[("a" + "\n\t")] = -956
    case "foo bar":
        v62 := v1["hello"]
        _ = v62
    case "hello":
        fmt.Println(v2)
    default:
        fmt.Println(v2, v2, v3)
    }
    var v63 string = "\n\t"
    v61 = v3
    _ = len(v1)
    fmt.Println(v2)
    switch v63 {
    case "\n\t":
        v61 -= 0.1
    case "hello":
        v2 -= (v2 >> ((672) & 63))
    case "foo bar":
        if (2147483647 != 605) {
            for fi64 := 0; fi64 < 19; fi64++ {
                _ = len(v1)
                _ = fi64
            }
            v65 := GenS5{F0: 0, F1: true, F2: 0}
            _ = v65
        } else {
            if false {
                v61 += v61
            } else {
                switch 128 {
                case 1:
                    var v66 float64 = v3
                    _ = v66
                default:
                    var v67 uint = 0
                    _ = v67
                }
            }
        }
    }
    switch v63 {
    case "\n\t":
        switch -2147483648 {
        case 0:
            var v68 uint = 960
            _ = v68
        case -619:
            _ = len(v1)
        }
    case "hello":
        v69 := GenS2{F0: 357, F1: (899 == 4294967295)}
        _ = v69
    case "a":
        v70 := GenS6{F0: 42, F1: 2}
        _ = v70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v63
}
