package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 uint8, p1 uint8) uint16 {
    v6 := GenS1{F0: 8, F1: -1, F2: 2}
    v7 := []int{32768, 32767, 128}
    v7 = append(v7, 4294967295)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 1000
    v3 := make(map[uint32]string)
    var v4 float64 = -2.25
    var v5 float32 = float32((0.1 * 3.14159))
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = -1
        v5 += (1.5 * 0.1)
        v9 -= -32768
        _ = fi8
        _ = v9
    }
    var v10 uint32 = 2
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 string = ""
        v10 = (v10 / ((v10) | 1))
        _ = fi11
        _ = v12
    }
    if true {
        v13 := GenS1{F0: 9223372036854775807, F1: -64, F2: 255}
        v14 := []int{32768, 4294967295}
        _ = v13
        _ = v14
    } else {
        v1 = ((v1 + ((v1 + "hello") + v1)) + (v1 + ("" + "hello")))
    }
    v4 = v4
    if (-91 < 100) {
        if false {
            var v15 uint16 = 418
            _ = v15
        } else {
            fmt.Println(v4, v5)
            v16 := GenS1{F0: 669, F1: -122, F2: 0}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int32 = 2147483647
        fmt.Println(v2)
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 uint = 255
            v21 := v3[(v10 >> ((v20) & 63))]
            if true {
                if false {
                    v22 := GenS1{F0: fi19, F1: -1, F2: 496}
                    _ = v22
                } else {
                    var v23 uint8 = 153
                    var v24 bool = false
                    _ = v23
                    _ = v24
                }
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    var v25 bool = ((!(!false)) && true)
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 int8 = 127
        v27 -= v27
        _ = fi26
        _ = v27
    }
    var v28 uint16 = genFn0(20, 128)
    v25 = ((0 <= -9223372036854775808) || (127 > -128))
    for fi29 := 0; fi29 < 3; fi29++ {
        v4 = 3.14159
        _ = fi29
    }
    v5 -= -2.25
    if true {
        var v30 int = 872
        v2 += v2
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 19; fi32++ {
            fmt.Println(v25)
            if ((!v25) || (1 <= -51)) {
                switch "\n\t" {
                case "\n\t":
                    v4 = (v4 / (((v4 * (float64(float32(v4)) + 0.0)) - v4) / float64(v5)))
                case "":
                    v5 -= v5
                case "foo bar":
                    v28 = 255
                }
            }
            v28 = (v28 << ((uint(v10)) & 63))
            _ = fi32
        }
        if false {
            for fi33 := 0; fi33 < 12; fi33++ {
                v10 = v10
                _ = fi33
            }
        } else {
            v34 := []int{fi31, 42}
            v35 := GenS2{F0: -1, F1: 15}
            _ = v34
            _ = v35
        }
        delete(v3, 2)
        _ = fi31
    }
    switch v1 {
    case "":
        for fi36 := 0; fi36 < 4; fi36++ {
            if true {
                var v37 float64 = v4
                for fi38 := 0; fi38 < 2; fi38++ {
                    v1 = "hello"
                    var v39 int8 = -71
                    if (!((!((904 == 42) && (!v25))) && (((v39 % ((v39) | 1)) >= int8(fi36)) && false))) {
                        if (9223372036854775807 == int64(fi36)) {
                            v40 := []int{int(v28), fi38, fi38, fi36}
                            _ = v40
                        }
                        var v41 uint64 = 95
                        _ = v41
                    } else {
                        v2 = -449
                        var v42 int = (fi38 >> ((1) & 63))
                        if v25 {
                            v43 := GenS2{F0: -2147483648, F1: 156}
                            _ = v43
                        }
                        _ = v42
                    }
                    var v44 int16 = (v2 * v2)
                    var v45 string = (v1 + "")
                    _ = fi38
                    _ = v39
                    _ = v44
                    _ = v45
                }
                var v46 string = ((v1 + ((("\n\t" + v1) + ("hello" + v1)) + v1)) + ("foo bar" + "hello"))
                _ = v37
                _ = v46
            } else {
                if v25 {
                    for fi47 := 0; fi47 < 7; fi47++ {
                        v48 := v3[v10]
                        v25 = (genFn0(171, 0) >= v28)
                        _ = fi47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 11; fi49++ {
                        fmt.Println(v10, fi36)
                        _ = fi49
                    }
                    switch "" {
                    case "a":
                        v50 := []int{(fi36 ^ (fi36 + (fi36 + (fi36 | -1))))}
                        _ = v50
                    case "":
                        v4 += float64(float32((v4 / 1e10)))
                    default:
                        if (!v25) {
                            var v51 uint64 = 291
                            for fi52 := 0; fi52 < 1; fi52++ {
                                delete(v3, v10)
                                v53 := GenS3{F0: (fi52 * -541), F1: "", F2: 0}
                                var v54 int = fi36
                                var v55 int = (fi36 ^ fi36)
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                            _ = v51
                        }
                    }
                } else {
                    v1 = v1
                    v2 -= -1
                }
                fmt.Println(v4, v28, v2)
            }
            switch "hello" {
            case "":
                v56 := v3[(v10 - (v10 & (v10 / ((858) | 1))))]
                _ = v56
            case "foo bar":
                v2 -= v2
            default:
                fmt.Println(v25, v28, v2)
            }
            _ = fi36
        }
    case "a":
        v57 := GenS3{F0: 128, F1: "foo bar", F2: 0}
        _ = v57
    case "\n\t":
        v2 = (v2 >> ((0) & 63))
    default:
        v10 = (v10 | (v10 + (^(v10))))
    }
    if (int32(v2) > 0) {
        var v58 int = 2147483647
        _ = v58
    } else {
        switch "hello" {
        case "foo bar":
            v25 = false
        case "\n\t":
            v59 := GenS2{F0: 2147483647, F1: 0}
            _ = v59
        }
        var v60 float64 = ((v4 * float64((3.14159 * -2.25))) - float64(v5))
        var v61 uint = 255
        _ = v60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
