package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS6 struct { F0 int16; F1 uint32; F2 uint }
func genFn0(p0 int64, p1 int, p2 int64) uint16 {
    var v7 uint16 = 224
    p0 = p2
    v8 := GenS1{F0: (v7 >> ((2) & 63)), F1: (^(v7))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{}
    v3 := make(map[string]bool)
    var v4 uint8 = 127
    v5 := make(map[string]uint32)
    v6 := make(map[string]int)
    fmt.Println(v4, v4)
    var v9 int32 = v1
    fmt.Println(v9, v1)
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            _ = len(v2)
            _ = fi11
        }
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        v12 := GenS2{F0: "hello"}
        _ = fi10
        _ = v12
    }
    var v13 bool = true
    if (42 != genFn0(9223372036854775807, -1000, 42)) {
        delete(v6, "foo bar")
        fmt.Println(v9, v4, v13)
        var v14 int16 = -1
        _ = v14
    } else {
        fmt.Println(v9)
    }
    switch -667 {
    case 1:
        for fi15 := 0; fi15 < 11; fi15++ {
            v9 -= (v9 ^ v1)
            v16 := GenS2{F0: "hello"}
            switch (fi15 / ((648) | 1)) {
            case 4294967295:
                fmt.Println(v1, v1)
            case -852:
                v17 := GenS3{F0: 200}
                _ = v17
            }
            _ = fi15
            _ = v16
        }
    case 32768:
        for fi18 := 0; fi18 < 12; fi18++ {
            v13 = v13
            _ = fi18
        }
    default:
        switch -2147483648 {
        case -278:
            fmt.Println(v1)
        case -1:
            fmt.Println(v13)
        }
    }
    v13 = (true || (false || (-1 <= v1)))
    switch "a" {
    case "hello":
        v13 = (2147483648 < 798)
    case "":
        if true {
            var v19 float32 = -2.25
            var v20 uint8 = (v4 >> ((569) & 63))
            switch 42 {
            case 4294967295:
                for fi21 := 0; fi21 < 6; fi21++ {
                    v22 := GenS4{F0: (v1 / ((v1) | 1)), F1: 1}
                    v13 = v13
                    v20 = v20
                    _ = fi21
                    _ = v22
                }
            }
            _ = v19
            _ = v20
        }
    case "a":
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float32 = 0.0
            _ = fi23
            _ = v24
        }
    default:
        if true {
            if v13 {
                var v25 bool = ((((true || v13) || ((1.5 + -2.25) > 0.0)) && v13) || (4294967295 == 42))
                _ = v25
            } else {
                if (2 != 851) {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v9 = 42
                        v9 = (v1 | v9)
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v9 = v9
                            v28 := GenS2{F0: ("\n\t" + "")}
                            _ = fi27
                            _ = v28
                        }
                        _ = fi26
                    }
                } else {
                    v1 += 42
                    v9 += 0
                }
                var v29 float64 = 0.1
                _ = len(v2)
                _ = v29
            }
            for fi30 := 0; fi30 < 6; fi30++ {
                v31 := GenS3{F0: v4}
                _ = fi30
                _ = v31
            }
            var v32 int = 1
            _ = v32
        } else {
            v33 := GenS5{F0: -434, F1: 234, F2: uint64(v4)}
            if (32767 >= 42) {
                v34 := GenS6{F0: 723, F1: 0, F2: 42}
                _ = v34
            } else {
                v1 = (v9 ^ (v1 & (v1 - (v9 * (v1 & v9)))))
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v33
        }
    }
    v1 += v1
    switch 2147483647 {
    case 4294967295:
        v35 := GenS5{F0: -32768, F1: v4, F2: 42}
        _ = v35
    case 2147483647:
        if (2 != 32767) {
            for fi36 := 0; fi36 < 8; fi36++ {
                var v37 uint = 472
                switch (fi36 / ((fi36) | 1)) {
                case 128:
                    v9 = v1
                case 32767:
                    var v38 uint32 = 913
                    _ = v38
                }
                _ = fi36
                _ = v37
            }
            if v13 {
                var v39 bool = ((v4 << ((0) & 63)) > (v4 << ((2) & 63)))
                var v40 uint16 = 857
                _ = v39
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 2; fi41++ {
                var v42 string = ("hello" + "a")
                for fi43 := 0; fi43 < 16; fi43++ {
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    _ = fi43
                }
                _ = fi41
                _ = v42
            }
            fmt.Println(v13)
            for fi44 := 0; fi44 < 8; fi44++ {
                v3[""] = ((true || (((!(65 >= 1)) && v13) && (((v4 & v4) < v4) && (9 == 18446744073709551615)))) && false)
                _ = fi44
            }
        }
    default:
        for fi45 := 0; fi45 < 16; fi45++ {
            v1 += v9
            _ = fi45
        }
    }
    var v46 uint64 = 1
    fmt.Println(v9)
    if v13 {
        for fi47 := 0; fi47 < 10; fi47++ {
            if len(v2) > 0 {
                v2[0] = 65535
            }
            if ((float32(0.0) < -2.25) && (!true)) {
                switch "foo bar" {
                case "\n\t":
                    v9 = -2147483648
                }
            }
            v48 := GenS2{F0: "hello"}
            _ = fi47
            _ = v48
        }
    }
    v13 = false
    v1 = (v1 / (((-(v9))) | 1))
    if (123 > v4) {
        if (v1 == v1) {
            switch 0 {
            case 334:
                fmt.Println(v1, v13)
            case 42:
                v2 = append(v2, 2147483647)
            case -1:
                v49 := GenS4{F0: (-(v1)), F1: 9223372036854775807}
                _ = v49
            default:
                v50 := GenS6{F0: -330, F1: 860, F2: 401}
                _ = v50
            }
            v1 = -411
            var v51 int16 = 0
            _ = v51
        }
        v9 -= v1
    } else {
        var v52 int16 = -32768
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
