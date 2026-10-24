package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16; F1 bool; F2 bool }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    v1 := []string{}
    var v2 int8 = -121
    v3 := []string{"hello", ("a" + ""), ("a" + "")}
    v4 := []uint32{1, 2, 2147483647, 114}
    v2 = -1
    v2 += 12
    switch "\n\t" {
    case "":
        var v5 uint8 = 1
        _ = v5
    case "\n\t":
        var v6 int64 = 85
        _ = v6
    default:
        v2 = -74
    }
    var v7 uint = 2
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v2)
        fmt.Println(v2, v7)
        var v9 uint8 = 2
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "hello" {
        case "foo bar":
            v2 -= (v2 << (((v7 >> ((161) & 63))) & 63))
        }
        _ = fi10
    }
    if true {
        v11 := GenS1{F0: (v7 + (v7 >> (((^(v7))) & 63)))}
        for fi12 := 0; fi12 < 4; fi12++ {
            if (255 != 0) {
                v13 := GenS1{F0: (v7 << ((689) & 63))}
                switch "foo bar" {
                case "\n\t":
                    var v14 int32 = 626
                    _ = v14
                case "hello":
                    var v15 uint = 890
                    _ = v15
                }
                if true {
                    var v16 int8 = v2
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v16
                } else {
                    var v17 uint = 1
                    _ = v17
                }
                _ = v13
            }
            _ = fi12
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 int16 = -683
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 = v2
                v21 := GenS1{F0: (v7 % ((v7) | 1))}
                switch "hello" {
                case "":
                    if false {
                        var v22 int = 0
                        v2 = 1
                        if len(v4) > 0 {
                            v4[0] = 2147483648
                        }
                        var v23 uint64 = 18446744073709551615
                        _ = v22
                        _ = v23
                    } else {
                        v7 = 0
                        if true {
                            var v24 uint64 = 149
                            v25 := GenS2{F0: "\n\t", F1: 9223372036854775807}
                            v4 = append(v4, 42)
                            _ = v24
                            _ = v25
                        } else {
                            _ = len(v3)
                        }
                    }
                }
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
        _ = v11
    }
    var v26 bool = false
    fmt.Println(v26)
    if true {
        if len(v4) > 0 {
            v4[0] = 2
        }
        v7 += (^(v7))
        switch ("foo bar" + (("\n\t" + "\n\t") + "\n\t")) {
        case "foo bar":
            switch ("" + ("a" + "foo bar")) {
            case "foo bar":
                var v27 bool = true
                _ = v27
            case "\n\t":
                var v28 int8 = (v2 + (^(v2)))
                _ = v28
            case "hello":
                v2 -= v2
            default:
                if v26 {
                    fmt.Println(v26, v26)
                    for fi29 := 0; fi29 < 9; fi29++ {
                        var v30 uint64 = 2
                        var v31 int16 = 1000
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    switch (("foo bar" + "foo bar") + "\n\t") {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = 2147483648
                        }
                    }
                } else {
                    var v32 int32 = 636
                    var v33 uint32 = 872
                    if true {
                        var v34 uint16 = 894
                        if v26 {
                            var v35 uint8 = 127
                            _ = v35
                        }
                        _ = v34
                    }
                    _ = v32
                    _ = v33
                }
            }
        default:
            for fi36 := 0; fi36 < 7; fi36++ {
                fmt.Println(v7)
                if v26 {
                    fmt.Println(v2, v26)
                    v37 := GenS3{F0: (fi36 * (fi36 | fi36))}
                    _ = v37
                } else {
                    v26 = v26
                    if true {
                        for fi38 := 0; fi38 < 5; fi38++ {
                            switch (fi38 % ((fi36) | 1)) {
                            case 39:
                                if false {
                                    if (2 != 65535) {
                                        v26 = false
                                        var v39 uint32 = 42
                                        _ = v39
                                    } else {
                                        v40 := GenS4{F0: 2147483647, F1: 698}
                                        fmt.Println(v7, v26, fi38)
                                        _ = v40
                                    }
                                    v41 := GenS5{F0: -880, F1: ((int16(fi36) >= 32767) || v26), F2: v26}
                                    _ = v41
                                }
                            }
                            _ = fi38
                        }
                        for fi42 := 0; fi42 < 18; fi42++ {
                            fmt.Println(v2, v26)
                            if v26 {
                                if false {
                                    var v43 int32 = 42
                                    v1 = append(v1, "hello")
                                    _ = v43
                                }
                            }
                            v44 := GenS6{F0: 32767, F1: v26}
                            _ = fi42
                            _ = v44
                        }
                    } else {
                        if (1 >= 42) {
                            var v45 int16 = 1
                            v46 := GenS7{F0: -520, F1: (v7 >> (((v7 | 40)) & 63))}
                            v2 -= v2
                            _ = v45
                            _ = v46
                        } else {
                            var v47 uint8 = 222
                            _ = v47
                        }
                    }
                    var v48 uint64 = 0
                    _ = v48
                }
                v26 = v26
                _ = fi36
            }
        }
    } else {
        if v26 {
            v7 = (v7 * v7)
            var v49 uint8 = 127
            v50 := GenS5{F0: -178, F1: v26, F2: false}
            _ = v49
            _ = v50
        } else {
            v7 += (v7 * (v7 << ((v7) & 63)))
            fmt.Println(v7, v26)
            if v26 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v51 uint32 = 514
                var v52 bool = v26
                _ = v51
                _ = v52
            }
        }
        _ = len(v1)
    }
    fmt.Println(v2, v2, v2)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v7
}
