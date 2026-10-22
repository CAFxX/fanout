package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 42
    v3 := []int{1, 866, -1, 127}
    v4 := []string{"", "a"}
    switch "" {
    case "":
        if (true || true) {
            v2 = 0
            v2 -= v2
            v2 += (v2 | (v2 - v2))
        }
    case "\n\t":
        if true {
            var v5 uint16 = 0
            switch "" {
            case "a":
                v6 := GenS1{F0: false, F1: 2147483647}
                _ = v6
            default:
                if (false && (true || (uint32(v5) < 65535))) {
                    for fi7 := 0; fi7 < 18; fi7++ {
                        var v8 float64 = 1.5
                        _ = fi7
                        _ = v8
                    }
                } else {
                    v5 -= (v5 << ((298) & 63))
                    if false {
                        if false {
                            if true {
                                var v9 int = -152
                                _ = v9
                            }
                            v2 += (v2 * (v2 << ((65535) & 63)))
                        } else {
                            v5 += v5
                            var v10 int8 = -1
                            _ = v10
                        }
                        var v11 bool = false
                        switch int(v2) {
                        case 32767:
                            for fi12 := 0; fi12 < 3; fi12++ {
                                var v13 int32 = 1
                                _ = fi12
                                _ = v13
                            }
                        default:
                            v14 := GenS1{F0: false, F1: 1}
                            _ = v14
                        }
                        _ = v11
                    }
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v2)
                    case "a":
                        if true {
                            v2 += 0
                            if len(v1) > 0 {
                                v1[0] = 542
                            }
                            v15 := GenS2{F0: 1}
                            _ = v15
                        } else {
                            var v16 bool = (!false)
                            _ = v16
                        }
                    case "\n\t":
                        if true {
                            for fi17 := 0; fi17 < 11; fi17++ {
                                v18 := GenS1{F0: false, F1: -2147483648}
                                v5 = v5
                                _ = fi17
                                _ = v18
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    default:
                        v19 := GenS1{F0: (18446744073709551615 >= 255), F1: 2147483647}
                        _ = v19
                    }
                }
            }
            fmt.Println(v2, v2)
            _ = v5
        }
    }
    v2 = (v2 + v2)
    for fi20 := 0; fi20 < 19; fi20++ {
        v2 -= v2
        v2 = (v2 - (v2 - (v2 >> ((255) & 63))))
        _ = fi20
    }
    var v21 bool = (2 < 0)
    v21 = (9223372036854775807 >= -599)
    fmt.Println(v21)
    if (18446744073709551615 > 865) {
        v22 := GenS3{F0: 1}
        _ = v22
    }
    if ((!(0 != 1)) && (2 > 0)) {
        fmt.Println(v21)
        if false {
            switch "" {
            case "a":
                if true {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        switch (("\n\t" + "") + "hello") {
                        case "hello":
                            var v24 uint = 1
                            _ = v24
                        default:
                            _ = len(v1)
                        }
                        _ = fi23
                    }
                }
            case "":
                for fi25 := 0; fi25 < 10; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        if len(v4) > 0 {
                            v4[0] = (("" + ("foo bar" + "a")) + "\n\t")
                        }
                        _ = fi26
                    }
                    var v27 uint8 = 33
                    var v28 string = ("" + "\n\t")
                    _ = fi25
                    _ = v27
                    _ = v28
                }
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = (("a" + "a") + "\n\t")
                }
            }
            var v29 uint32 = 65535
            for fi30 := 0; fi30 < 16; fi30++ {
                v21 = ((int8(fi30) < -1) || (v21 || (1e10 == -2.25)))
                v29 = v29
                var v31 int16 = 1
                _ = fi30
                _ = v31
            }
            _ = v29
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v32 string = "a"
        _ = v32
    }
    v33 := GenS4{F0: 2, F1: ((("a" + "foo bar") + "foo bar") + ("\n\t" + ("foo bar" + "\n\t"))), F2: 255}
    v34 := GenS3{F0: 174}
    v35 := GenS1{F0: ((!(v21 || (0 > 170))) || (!(true || v21))), F1: -838}
    v21 = ((v21 && v21) || (v21 || false))
    var v36 uint = 65535
    fmt.Println(v21, v36, v2)
    if false {
        var v37 int = -152
        v1 = append(v1, 1)
        v21 = false
        _ = v37
    } else {
        fmt.Println(v2, v36)
    }
    if v21 {
        var v38 uint64 = 138
        _ = v38
    }
    v39 := GenS1{F0: v21, F1: -1}
    switch 32767 {
    case -9223372036854775808:
        v1 = append(v1, 65535)
    case -1:
        var v40 int32 = -1
        _ = v40
    default:
        v41 := GenS3{F0: uint16(v36)}
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
