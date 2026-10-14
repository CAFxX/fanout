package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 int }

func main() {
    v1 := []uint32{}
    var v2 bool = true
    var v3 string = "\n\t"
    _ = len(v1)
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 8; fi5++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v6 uint32 = 42
        v6 = 1
        _ = fi5
        _ = v6
    }
    var v7 string = "\n\t"
    v8 := GenS1{F0: 2, F1: 2}
    var v9 uint8 = 117
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        fmt.Println(v2, v4, v7)
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "foo bar":
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: 85, F1: 18446744073709551615}
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint8 = (v9 >> ((4294967295) & 63))
                var v16 int = fi12
                fmt.Println(v3)
                v16 = v16
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    case "":
        var v17 int32 = 0
        _ = v17
    case "a":
        if false {
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS3{F0: 65535, F1: (fi18 + (fi18 & (fi18 * 127))), F2: (v9 - (^(v9)))}
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 bool = false
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            v4 = (v4 + v4)
        } else {
            switch 353 {
            case 32767:
                for fi22 := 0; fi22 < 6; fi22++ {
                    if v2 {
                        var v23 int32 = -2147483648
                        switch v3 {
                        case "\n\t":
                            if false {
                                v24 := GenS2{F0: -128, F1: 381}
                                if true {
                                    if (4294967295 < uint(v4)) {
                                        v3 = v7
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        v7 = (v3 + v3)
                                        v25 := GenS4{F0: (v7 + ""), F1: ("\n\t" + ("hello" + ("a" + ("hello" + ("" + (((v3 + (v7 + "foo bar")) + (v3 + (((("\n\t" + ((v7 + ("hello" + v7)) + "a")) + ("foo bar" + "a")) + "") + ("" + ("foo bar" + v3))))) + v3))))))}
                                        _ = v25
                                    }
                                } else {
                                    for fi26 := 0; fi26 < 11; fi26++ {
                                        v9 -= 2
                                        _ = fi26
                                    }
                                }
                                var v27 uint8 = 138
                                _ = v24
                                _ = v27
                            }
                        case "a":
                            v23 = v23
                        default:
                            v7 = v3
                        }
                        v28 := GenS4{F0: "a", F1: "\n\t"}
                        _ = v23
                        _ = v28
                    }
                    if (v2 || v2) {
                        var v29 float64 = ((0.1 - -2.25) / 1.5)
                        var v30 int32 = 1
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v2, v29)
                        default:
                            fmt.Println(v7, v3, fi22)
                        }
                        _ = v29
                        _ = v30
                    }
                    _ = len(v1)
                    _ = fi22
                }
            case 32768:
                for fi31 := 0; fi31 < 8; fi31++ {
                    if true {
                        var v32 int16 = 0
                        var v33 string = (v3 + v7)
                        v2 = ((v9 & 13) != (v9 << ((172) & 63)))
                        _ = v32
                        _ = v33
                    } else {
                        if (v9 <= 255) {
                            v34 := GenS5{F0: v4, F1: (v9 * 251)}
                            _ = v34
                        } else {
                            var v35 bool = (!v2)
                            _ = v35
                        }
                    }
                    v36 := GenS6{F0: 122, F1: v3}
                    _ = fi31
                    _ = v36
                }
            case -2147483648:
                fmt.Println(v2, v9, v2)
            }
            var v37 bool = v2
            _ = v37
        }
    default:
        v1 = append(v1, 1)
    }
    var v38 uint32 = 0
    _ = len(v1)
    fmt.Println(v4)
    v3 = v3
    var v39 uint16 = (v4 ^ (v4 % ((v4) | 1)))
    fmt.Println(v7, v39)
    _ = len(v1)
    for fi40 := 0; fi40 < 16; fi40++ {
        if v2 {
            for fi41 := 0; fi41 < 19; fi41++ {
                var v42 uint16 = v39
                if v2 {
                    v7 = "hello"
                    var v43 uint16 = (v39 * 2)
                    _ = len(v1)
                    switch (fi41 | 128) {
                    case 24:
                        var v44 int32 = -1
                        _ = v44
                    case 127:
                        v43 = v42
                    default:
                        var v45 int16 = 42
                        _ = v45
                    }
                    v43 = v39
                    _ = v43
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v46 int = fi40
                    _ = v46
                }
                if (!((1e10 >= 0.1) && false)) {
                    v47 := GenS7{F0: fi40}
                    var v48 uint16 = 310
                    _ = v47
                    _ = v48
                }
                _ = fi41
                _ = v42
            }
            switch v7 {
            case "\n\t":
                v3 = (v7 + v7)
            }
        }
        var v49 string = ((v7 + v7) + "\n\t")
        _ = fi40
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
