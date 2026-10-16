package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint32; F1 bool; F2 uint }
type GenS7 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0() int {
    var v6 int32 = 2147483647
    var v7 uint8 = 42
    if (true && ((!false) || true)) {
        v7 -= (^(v7))
        switch 1 {
        case 4294967295:
            fmt.Println(v6, v7)
        case 576:
            v8 := GenS1{F0: 128, F1: 42}
            _ = v8
        default:
            switch -212 {
            case 317:
                if false {
                    var v9 int8 = -51
                    if (true && ((^(v9)) > v9)) {
                        v10 := []int{128}
                        v7 -= (^(v7))
                        _ = v10
                    }
                    _ = v9
                } else {
                    fmt.Println(v6, v7)
                }
            case 4294967295:
                for fi11 := 0; fi11 < 19; fi11++ {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v13 := GenS2{F0: 1, F1: 127}
                        var v14 int = (fi11 << ((549) & 63))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi11
                }
            case 32767:
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint16 = 255
                    v17 := []int{int(v7)}
                    fmt.Println(v7, fi15)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
        v6 -= v6
    } else {
        switch -2147483648 {
        case -653:
            v18 := GenS1{F0: v7, F1: 964}
            _ = v18
        case 2147483647:
            for fi19 := 0; fi19 < 20; fi19++ {
                v7 += (v7 | v7)
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch "" {
                    case "":
                        var v21 int16 = -1
                        _ = v21
                    default:
                        if ((fi19 << ((894) & 63)) >= -2147483648) {
                            if false {
                                v22 := GenS1{F0: v7, F1: 281}
                                var v23 string = (("\n\t" + ("a" + ("\n\t" + (("hello" + ("hello" + "hello")) + ("hello" + ""))))) + ("foo bar" + "hello"))
                                v24 := GenS3{F0: -1}
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                    var v25 int64 = 206
                    _ = fi20
                    _ = v25
                }
                v26 := GenS4{F0: "\n\t"}
                _ = fi19
                _ = v26
            }
        case 1:
            v27 := GenS4{F0: (("a" + "") + (("a" + "\n\t") + (("" + "") + "\n\t")))}
            _ = v27
        }
        switch int(v7) {
        case 9223372036854775807:
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 uint32 = 2147483647
                _ = fi28
                _ = v29
            }
        case -2147483648:
            switch 2147483647 {
            case -2147483648:
                v30 := []int{}
                _ = v30
            case 859:
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := []int{}
                    _ = fi31
                    _ = v32
                }
            }
        default:
            v7 = v7
        }
        if false {
            fmt.Println(v6)
            v6 += 903
            v6 += v6
        } else {
            for fi33 := 0; fi33 < 15; fi33++ {
                if (!(true && true)) {
                    if ((("a" == ("foo bar" + "hello")) || false) || (!false)) {
                        v34 := []int{0, int(v6)}
                        _ = len(v34)
                        _ = v34
                    }
                    v35 := []int{(fi33 | (fi33 - 0))}
                    v36 := GenS5{F0: fi33, F1: (v7 + v7), F2: 65535}
                    _ = v35
                    _ = v36
                } else {
                    v7 = (^(v7))
                }
                _ = fi33
            }
            switch 32768 {
            case -754:
                v6 += 1
            case -584:
                v37 := GenS5{F0: 127, F1: 1, F2: 0}
                _ = v37
            case 32768:
                v38 := GenS6{F0: 218, F1: ((v7 ^ 128) <= 255), F2: 65535}
                _ = v38
            default:
                v6 += v6
            }
            for fi39 := 0; fi39 < 8; fi39++ {
                v6 += v6
                _ = fi39
            }
        }
    }
    var v40 uint16 = 255
    _ = v40
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int64 = 18
    var v2 uint64 = 42
    var v3 uint = 0
    v4 := []string{(("\n\t" + (("foo bar" + "a") + "hello")) + "\n\t"), "\n\t", "a", ""}
    v5 := make(map[string]uint32)
    v2 -= v2
    fmt.Println(v3, v1, v3)
    v41 := GenS3{F0: v1}
    for fi42 := 0; fi42 < 15; fi42++ {
        v3 = (v3 / ((2) | 1))
        _ = fi42
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS7{F0: genFn0(), F1: (v2 / (((v2 >> ((v3) & 63))) | 1)), F2: 434}
        _ = fi43
        _ = v44
    }
    var v45 float64 = 3.14159
    if true {
        var v46 uint = (v3 / ((614) | 1))
        for fi47 := 0; fi47 < 14; fi47++ {
            fmt.Println(v3, v3, v2)
            if false {
                fmt.Println(v46, v2, v3)
                delete(v5, "foo bar")
                var v48 uint16 = 42
                _ = v48
            } else {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    var v49 uint16 = 679
                    _ = v49
                case "foo bar":
                    v50 := GenS1{F0: 255, F1: 258}
                    _ = v50
                default:
                    switch (fi47 >> ((v46) & 63)) {
                    case 676:
                        var v51 bool = (v45 > ((v45 + float64((((1e10 / -2.25) * 0.1) + 1e10))) - 0.0))
                        _ = v51
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
            _ = fi47
        }
        var v52 uint16 = 1
        _ = v46
        _ = v52
    }
    var v53 float64 = v45
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v5
    _ = v53
}
