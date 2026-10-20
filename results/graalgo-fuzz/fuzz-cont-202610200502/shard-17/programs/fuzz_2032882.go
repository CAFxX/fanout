package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int64; F2 uint32 }
func genFn0(p0 string) string {
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 = p0
        v6 := []int{(fi5 / ((fi5) | 1)), (fi5 * 1), (fi5 | (-(fi5)))}
        var v7 uint32 = 2147483648
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint64 = 4294967295
        v9 = (v9 - (v9 - v9))
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        p0 = (p0 + "a")
        for fi12 := 0; fi12 < 5; fi12++ {
            if false {
                var v13 int32 = 0
                _ = v13
            }
            p0 = "hello"
            var v14 int32 = -1
            _ = fi12
            _ = v14
        }
        _ = fi11
    }
    _ = p0
    return (((p0 + "") + ("foo bar" + "")) + ("a" + p0))
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 uint64 = 42
    var v4 string = "hello"
    v15 := []int{2147483647}
    v4 = v4
    delete(v1, 65535)
    v16 := GenS1{F0: uint8(v2), F1: 128}
    var v17 uint8 = 148
    if false {
        fmt.Println(v3)
    } else {
        switch 42 {
        case 32768:
            switch "hello" {
            case "":
                var v18 int16 = -1
                _ = v18
            case "hello":
                if true {
                    _ = len(v1)
                } else {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v4, v2, v3)
                    }
                    _ = len(v15)
                }
            }
        case 32767:
            var v19 int64 = 1
            _ = v19
        default:
            v3 = (v3 + v3)
        }
        v20 := GenS2{F0: v17}
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS1{F0: v17, F1: (v17 << ((18446744073709551615) & 63))}
        _ = fi21
        _ = v22
    }
    if true {
        v17 = v17
    } else {
        v1[2147483647] = "\n\t"
        var v23 float32 = (((float32(0.1) + (1.5 * (((1e10 + ((1e10 - -2.25) - 1e10)) - -2.25) - (0.0 * 0.1)))) + 1e10) * ((1e10 + 1.5) + (0.1 * (0.0 - 0.1))))
        _ = v23
    }
    switch 9223372036854775807 {
    case -1:
        var v24 bool = (2147483647 >= 42)
        _ = v24
    default:
        if true {
            if false {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v17 -= v17
                    v3 -= v3
                    v26 := v1[584]
                    _ = fi25
                    _ = v26
                }
                v17 = v17
            } else {
                var v27 int64 = -481
                _ = v27
            }
            switch 9223372036854775807 {
            case -394:
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch v4 {
                    case "a":
                        v15 = append(v15, (fi28 % ((fi28) | 1)))
                    default:
                        var v29 int64 = int64(v17)
                        _ = v29
                    }
                    v30 := GenS3{F0: (fi28 % ((2147483647) | 1)), F1: 0, F2: 2147483647}
                    if ((v3 + 2) >= v3) {
                        fmt.Println(v2)
                    } else {
                        var v31 uint8 = 41
                        var v32 int = 32767
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                    _ = v30
                }
            default:
                if (true && (true && true)) {
                    var v33 int = 32768
                    for fi34 := 0; fi34 < 6; fi34++ {
                        if len(v15) > 0 {
                            v15[0] = v33
                        }
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v36 := GenS2{F0: (^(v17))}
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v17, v17, v33)
                        _ = fi34
                    }
                    _ = v33
                } else {
                    var v37 bool = true
                    for fi38 := 0; fi38 < 6; fi38++ {
                        for fi39 := 0; fi39 < 7; fi39++ {
                            if (255 != 2) {
                                v40 := GenS1{F0: (v17 * 0), F1: 0}
                                if v37 {
                                    v17 += 128
                                    var v41 int64 = -240
                                    _ = v41
                                }
                                v2 -= (v2 + v2)
                                v2 += v2
                                if v37 {
                                    var v42 float64 = (0.1 / (((-2.25 * 1e10) - 3.14159) / ((3.14159 * 1.5) / 1.0)))
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    _ = v42
                                }
                                _ = v40
                            }
                            _ = fi39
                        }
                        _ = fi38
                    }
                    _ = v37
                }
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        } else {
            switch "hello" {
            case "\n\t":
                var v43 int16 = 0
                _ = v43
            }
        }
    }
    v4 = v4
    switch 293 {
    case 1:
        delete(v1, 42)
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        v2 = (v2 / ((65535) | 1))
        for fi45 := 0; fi45 < 10; fi45++ {
            v3 += v3
            _ = fi45
        }
        if (255 >= (v17 & 42)) {
            switch v4 {
            case "":
                v4 = (("a" + v4) + "")
            case "\n\t":
                var v46 int = fi44
                _ = v46
            case "foo bar":
                for fi47 := 0; fi47 < 2; fi47++ {
                    _ = len(v15)
                    var v48 uint16 = 42
                    _ = fi47
                    _ = v48
                }
            }
            fmt.Println(v2)
        }
        _ = fi44
    }
    var v49 int32 = 1
    for fi50 := 0; fi50 < 13; fi50++ {
        var v51 uint = 4294967295
        v3 = (^(v3))
        var v52 string = ""
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v49
}
