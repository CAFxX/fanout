package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 bool = (1 != 65535)
    var v2 int8 = 0
    var v3 string = ""
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    var v6 uint16 = 0
    if (1 > 959) {
        if (!((!(true && ((v1 && (!(42 > 32767))) || true))) && true)) {
            switch -728 {
            case 1:
                switch -500 {
                case 127:
                    var v7 uint = 885
                    _ = v7
                case -573:
                    delete(v4, 128)
                default:
                    v8 := GenS1{F0: 2, F1: v1}
                    _ = v8
                }
            }
            var v9 uint = 836
            _ = v9
        } else {
            v3 = (((("" + "") + "hello") + "foo bar") + (v3 + v3))
        }
        fmt.Println(v3)
    } else {
        switch v3 {
        case "foo bar":
            v6 += (v6 & v6)
        case "hello":
            if ((!(!((v2 + v2) <= (v2 - 103)))) || (103 > v2)) {
                var v10 uint16 = 490
                v11 := []int{32767, 44, 32768}
                _ = v10
                _ = v11
            } else {
                v1 = (true || (1 <= 2))
            }
        case "":
            var v12 bool = true
            _ = v12
        default:
            switch "" {
            case "\n\t":
                if v1 {
                    var v13 string = v3
                    switch "hello" {
                    case "":
                        v14 := GenS1{F0: 2, F1: (!(!v1))}
                        _ = v14
                    case "a":
                        fmt.Println(v6)
                    case "foo bar":
                        v15 := v5[(v3 + v13)]
                        _ = v15
                    default:
                        if (2147483647 != 4294967295) {
                            for fi16 := 0; fi16 < 11; fi16++ {
                                var v17 uint32 = 2
                                _ = fi16
                                _ = v17
                            }
                        }
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := GenS2{F0: 898, F1: (v3 + v3), F2: 1}
                        fmt.Println(v13, v2, v2)
                        _ = fi18
                        _ = v19
                    }
                    _ = v13
                } else {
                    v6 -= v6
                }
            case "hello":
                v4[-9223372036854775808] = v3
            default:
                v3 = v3
            }
        }
    }
    var v20 uint = 0
    fmt.Println(v2, v2)
    if (-9223372036854775808 < -556) {
        v21 := GenS2{F0: 307, F1: (v3 + v3), F2: (v20 + (^(v20)))}
        if (!(v2 == (v2 - (v2 * (-(v2)))))) {
            if true {
                v22 := []int{9223372036854775807}
                fmt.Println(v1, v1)
                _ = v22
            } else {
                fmt.Println(v20, v6)
                v23 := GenS3{F0: 0, F1: 127}
                _ = v23
            }
            var v24 uint16 = (v6 * (v6 >> (((^(v20))) & 63)))
            _ = v24
        }
        _ = v21
    }
    v25 := GenS2{F0: 145, F1: v3, F2: 4294967295}
    for fi26 := 0; fi26 < 9; fi26++ {
        fmt.Println(fi26)
        _ = fi26
    }
    v27 := v5["a"]
    if v1 {
        var v28 int32 = 1
        v6 -= 65535
        switch (v3 + v3) {
        case "\n\t":
            v20 = (v20 & (v20 ^ v20))
        case "a":
            v29 := []int{32768, 570}
            _ = v29
        default:
            switch v3 {
            case "foo bar":
                if v1 {
                    if true {
                        fmt.Println(v1, v2, v3)
                        v30 := GenS2{F0: 1, F1: ("\n\t" + ""), F2: v20}
                        if (!true) {
                            v28 = v28
                        }
                        _ = v30
                    }
                    if v1 {
                        v6 = v6
                    }
                    for fi31 := 0; fi31 < 2; fi31++ {
                        switch 2147483647 {
                        case 4294967295:
                            v32 := []int{(fi31 << ((2) & 63)), 4294967295, fi31}
                            _ = v32
                        case -657:
                            switch v27 {
                            case "a":
                                _ = len(v4)
                            }
                        default:
                            var v33 uint = (v20 << ((693) & 63))
                            _ = v33
                        }
                        _ = fi31
                    }
                } else {
                    switch (v27 + ("foo bar" + v3)) {
                    case "a":
                        fmt.Println(v27, v6, v28)
                    case "hello":
                        var v34 uint32 = 660
                        _ = v34
                    }
                    if (!(!false)) {
                        if false {
                            if true {
                                v3 = v27
                                v35 := GenS2{F0: 0, F1: v27, F2: (^(v20))}
                                _ = v35
                            }
                            v6 += (v6 & v6)
                        }
                        var v36 int8 = 1
                        _ = v36
                    }
                }
            case "hello":
                var v37 bool = v1
                _ = v37
            case "\n\t":
                v6 = (v6 >> ((v20) & 63))
            }
        }
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
