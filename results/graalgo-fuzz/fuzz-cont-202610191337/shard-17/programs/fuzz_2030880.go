package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int8 = 48
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(v1)
        v1 = 127
        _ = fi5
    }
    var v6 string = ""
    switch 1 {
    case 32767:
        if ((2 == 140) || false) {
            var v7 bool = false
            for fi8 := 0; fi8 < 10; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    v6 = "hello"
                    _ = fi9
                }
                v10 := GenS1{F0: 155}
                _ = fi8
                _ = v10
            }
            for fi11 := 0; fi11 < 17; fi11++ {
                if false {
                    v12 := v3[467]
                    _ = v12
                } else {
                    fmt.Println(v6)
                    var v13 uint = 0
                    _ = v13
                }
                _ = fi11
            }
            _ = v7
        } else {
            var v14 uint8 = 71
            _ = v14
        }
    case 0:
        if (!(v6 == "hello")) {
            v1 -= (^(v1))
        } else {
            v15 := GenS1{F0: 2}
            for fi16 := 0; fi16 < 6; fi16++ {
                if (-1 <= 176) {
                    if (!(true || true)) {
                        var v17 int32 = -2147483648
                        _ = v17
                    } else {
                        if true {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                if false {
                                    if true {
                                        v19 := v2[0]
                                        _ = v19
                                    }
                                    v20 := GenS1{F0: 1}
                                    _ = v20
                                }
                                v1 += (v1 * (-(v1)))
                                _ = fi18
                            }
                            v21 := []int{(fi16 / ((-2147483648) | 1)), -1, (fi16 ^ (fi16 & (fi16 * -177)))}
                            _ = v21
                        } else {
                            var v22 int16 = -1
                            fmt.Println(fi16)
                            _ = v22
                        }
                        v23 := []int{(fi16 ^ 4294967295)}
                        var v24 uint8 = 160
                        _ = v23
                        _ = v24
                    }
                } else {
                    switch 4294967295 {
                    case 2147483647:
                        fmt.Println(v1, v6, v1)
                    default:
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 int16 = 0
                            v6 = "hello"
                            v27 := GenS1{F0: 1}
                            fmt.Println(v26)
                            v28 := GenS1{F0: 4294967295}
                            _ = fi25
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                    }
                    var v29 uint = 4294967295
                    _ = v29
                }
                delete(v4, fi16)
                _ = fi16
            }
            _ = v15
        }
    case -267:
        if false {
            v1 -= 61
            v6 = v6
        } else {
            switch v6 {
            case "hello":
                v6 = v6
            case "\n\t":
                v1 = 1
            case "":
                if (false || false) {
                    v1 -= v1
                    v6 = v6
                    var v30 bool = false
                    _ = v30
                } else {
                    var v31 uint32 = 457
                    for fi32 := 0; fi32 < 14; fi32++ {
                        delete(v3, v31)
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 uint16 = 0
                            v35 := GenS1{F0: 4294967295}
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS1{F0: 18446744073709551615}
                        _ = fi32
                        _ = v36
                    }
                    _ = v31
                }
            default:
                switch -49 {
                case 127:
                    switch v6 {
                    case "foo bar":
                        fmt.Println(v1, v6)
                    case "hello":
                        fmt.Println(v6)
                    case "":
                        v37 := []int{-582, 1, -1, 32768}
                        _ = v37
                    }
                case -7:
                    for fi38 := 0; fi38 < 12; fi38++ {
                        v1 += (v1 | v1)
                        _ = fi38
                    }
                case 32768:
                    v1 -= -1
                }
            }
            v1 += v1
            _ = len(v3)
        }
    }
    switch v6 {
    case "a":
        v39 := []int{2147483647, 32768, 610, 32768}
        _ = v39
    case "":
        if false {
            v1 -= v1
            for fi40 := 0; fi40 < 7; fi40++ {
                if true {
                    v6 = v6
                    if (!false) {
                        var v41 bool = true
                        v42 := GenS2{F0: 18446744073709551615, F1: v1}
                        v6 = ""
                        _ = v41
                        _ = v42
                    } else {
                        switch (fi40 ^ fi40) {
                        case 123:
                            var v43 float64 = 1e10
                            _ = v43
                        }
                    }
                } else {
                    v6 = v6
                }
                _ = fi40
            }
            switch "foo bar" {
            case "":
                v44 := GenS1{F0: 1}
                _ = v44
            case "\n\t":
                var v45 int8 = (v1 << ((65535) & 63))
                _ = v45
            }
        }
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 5; fi47++ {
            v1 = v1
            _ = fi47
        }
        _ = fi46
    }
    v48 := []int{}
    switch "\n\t" {
    case "foo bar":
        var v49 uint = 255
        _ = v49
    case "hello":
        v50 := GenS1{F0: 278}
        _ = v50
    case "":
        var v51 uint64 = 0
        _ = v51
    }
    v6 = "a"
    v1 -= v1
    v52 := GenS3{F0: -347}
    v6 = "foo bar"
    fmt.Println(v1)
    var v53 uint64 = 2
    for fi54 := 0; fi54 < 1; fi54++ {
        v55 := v2[2147483648]
        v56 := GenS4{F0: (v6 > "a")}
        _ = fi54
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v52
    _ = v53
    _ = v6
}
