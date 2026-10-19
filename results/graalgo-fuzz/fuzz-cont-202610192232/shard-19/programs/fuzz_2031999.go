package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 int16; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint; F1 uint16; F2 int32 }
type GenS10 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint8) int8 {
    v7 := GenS1{F0: -801, F1: 4294967295, F2: 1}
    fmt.Println(p0)
    v8 := []int{32768, -1}
    p0 -= p0
    v9 := GenS1{F0: -9223372036854775808, F1: 350, F2: 2147483647}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -128
}

func main() {
    var v1 int16 = -1
    var v2 uint32 = 0
    var v3 int32 = 0
    v4 := make(map[uint32]bool)
    var v5 uint8 = 0
    v6 := []uint32{2147483648, v2, 2}
    switch ("hello" + ((("\n\t" + "a") + "\n\t") + ("hello" + "hello"))) {
    case "a":
        switch ("" + "a") {
        case "":
            for fi10 := 0; fi10 < 2; fi10++ {
                v2 -= (v2 + (^(v2)))
                fmt.Println(v1, v3, fi10)
                for fi11 := 0; fi11 < 8; fi11++ {
                    _ = len(v6)
                    v4[v2] = true
                    if (true || false) {
                        v3 = v3
                    }
                    _ = fi11
                }
                _ = fi10
            }
        case "a":
            var v12 int64 = 42
            _ = v12
        }
    default:
        v1 = (v1 << ((65535) & 63))
    }
    var v13 float32 = 1e10
    v14 := GenS2{F0: 312}
    switch "hello" {
    case "foo bar":
        switch 9223372036854775807 {
        case 42:
            _ = len(v4)
        case 9223372036854775807:
            _ = len(v4)
        default:
            if ((v2 + (^(v2))) == 2) {
                if false {
                    switch 915 {
                    case 451:
                        if (true && true) {
                            v15 := GenS1{F0: 42, F1: 1, F2: v3}
                            var v16 int32 = 363
                            _ = v15
                            _ = v16
                        } else {
                            v13 = v13
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v17 int16 = (v1 + v1)
                            _ = v17
                        }
                    case -2147483648:
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v19 := GenS3{F0: (fi18 ^ 32768), F1: v2, F2: v5}
                            v20 := GenS1{F0: 1, F1: 255, F2: v3}
                            delete(v4, (v2 * (v2 | 0)))
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    default:
                        var v21 uint = 65535
                        _ = v21
                    }
                }
            }
        }
    case "a":
        var v22 float32 = v13
        _ = v22
    }
    var v23 int64 = 1
    if true {
        v24 := GenS4{F0: (v23 | 0)}
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 7; fi26++ {
                _ = len(v4)
                switch 2147483647 {
                case 127:
                    v2 -= (v2 + 2)
                case 128:
                    switch 32767 {
                    case -305:
                        v6 = append(v6, (v2 >> ((754) & 63)))
                    case 577:
                        fmt.Println(fi26)
                    }
                default:
                    v23 = (^(v23))
                    var v27 int8 = 127
                    fmt.Println(v27, v1, v3)
                    _ = v27
                }
                if false {
                    v28 := GenS5{F0: v2, F1: 42}
                    _ = v28
                }
                _ = fi26
            }
            _ = fi25
        }
        _ = len(v6)
        _ = v24
    } else {
        var v29 uint32 = v2
        switch 4294967295 {
        case 32768:
            v1 -= (^(v1))
        case 2147483647:
            var v30 uint8 = (v5 ^ (v5 << ((369) & 63)))
            _ = v30
        default:
            for fi31 := 0; fi31 < 14; fi31++ {
                v13 -= ((v13 + v13) - v13)
                var v32 int = (fi31 - -129)
                _ = fi31
                _ = v32
            }
        }
        _ = v29
    }
    v33 := GenS6{F0: 23}
    v34 := GenS7{F0: 177, F1: (^(v1)), F2: false}
    if true {
        delete(v4, v2)
        v35 := GenS8{F0: (v5 % ((v5) | 1))}
        fmt.Println(v23, v5)
        _ = v35
    }
    if false {
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint16 = 79
            fmt.Println(v37)
            _ = fi36
            _ = v37
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 1 {
    case 32768:
        for fi38 := 0; fi38 < 6; fi38++ {
            for fi39 := 0; fi39 < 7; fi39++ {
                v3 += 15
                _ = fi39
            }
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint16 = 995
                if false {
                    fmt.Println(v1)
                    if true {
                        v23 = (^(v23))
                    } else {
                        v13 -= (1e10 * 1.5)
                        v42 := GenS9{F0: 0, F1: v41, F2: 1}
                        _ = v42
                    }
                }
                _ = fi40
                _ = v41
            }
            _ = fi38
        }
    default:
        if (444 < v2) {
            for fi43 := 0; fi43 < 14; fi43++ {
                fmt.Println(v3, v3, fi43)
                _ = fi43
            }
            if false {
                v5 += (v5 << ((915) & 63))
            }
            for fi44 := 0; fi44 < 3; fi44++ {
                v4[v2] = (!false)
                for fi45 := 0; fi45 < 3; fi45++ {
                    v46 := GenS10{F0: 0, F1: 12}
                    fmt.Println(v3)
                    var v47 int = -900
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                switch "hello" {
                case "\n\t":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "":
                    v23 -= v23
                default:
                    if (true || true) {
                        var v48 int8 = 40
                        v2 += v2
                        _ = v48
                    } else {
                        fmt.Println(v2)
                    }
                }
                _ = fi44
            }
        }
    }
    var v49 int8 = genFn0(1)
    var v50 int8 = int8(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
