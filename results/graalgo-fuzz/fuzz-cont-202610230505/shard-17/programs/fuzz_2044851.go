package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 uint32) int64 {
    fmt.Println(p1)
    v6 := GenS1{F0: 761, F1: 0, F2: false}
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint8) string {
    p0 = 128
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int32 = 1
        v9 := []int{fi7, (fi7 | fi7)}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            if true {
                v12 := GenS1{F0: 0, F1: genFn0(true, 354), F2: false}
                v13 := []int{}
                v14 := GenS2{F0: 1, F1: -1}
                _ = v12
                _ = v13
                _ = v14
            }
            v15 := []int{(-(fi11)), int(p0), (^(fi10)), (-(fi10))}
            _ = fi11
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 string = "hello"
            if ((-2.25 + 3.14159) >= (float64(3.14159) - 0.1)) {
                var v18 int16 = 1
                fmt.Println(p0)
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        _ = fi10
    }
    if (32767 != 977) {
        var v19 int16 = -32768
        var v20 string = "\n\t"
        for fi21 := 0; fi21 < 11; fi21++ {
            if false {
                if true {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        var v23 string = (v20 + (v20 + (v20 + v20)))
                        v24 := GenS2{F0: 100, F1: (v19 * int16(p0))}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v26 := []int{fi25}
                        p0 += (p0 << ((1) & 63))
                        v27 := GenS2{F0: -65, F1: v19}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v28 := []int{(fi21 / (((^(fi21))) | 1)), (fi21 | fi21)}
                    _ = v28
                }
                switch ("foo bar" + (v20 + ((v20 + v20) + ((v20 + v20) + ("foo bar" + "hello"))))) {
                case "hello":
                    var v29 uint8 = (p0 + (p0 / (((p0 << ((84) & 63))) | 1)))
                    _ = v29
                case "a":
                    if false {
                        v30 := GenS2{F0: -121, F1: (v19 + 42)}
                        v19 = 1
                        var v31 uint8 = 0
                        _ = v30
                        _ = v31
                    } else {
                        switch (v20 + v20) {
                        case "":
                            v20 = v20
                        }
                        if (1e10 != 0.1) {
                            p0 = (^(p0))
                            v19 = 1000
                            var v32 int = (fi21 | int(v19))
                            _ = v32
                        } else {
                            v33 := GenS2{F0: 63, F1: (v19 << ((2) & 63))}
                            var v34 uint8 = 208
                            _ = v33
                            _ = v34
                        }
                    }
                default:
                    fmt.Println(fi21, v19)
                }
            }
            v35 := GenS2{F0: 96, F1: 917}
            _ = fi21
            _ = v35
        }
        _ = v19
        _ = v20
    } else {
        v36 := []int{}
        p0 = p0
        p0 += (p0 >> ((2) & 63))
        _ = v36
    }
    fmt.Println(p0)
    _ = p0
    return (("\n\t" + "foo bar") + "a")
}

func main() {
    var v1 uint = 0
    var v2 float32 = (3.14159 / 0.1)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 uint16 = 42
    if true {
        var v37 int64 = 42
        if (1 > -65) {
            var v38 int64 = (v37 << ((1) & 63))
            if (!false) {
                if false {
                    for fi39 := 0; fi39 < 11; fi39++ {
                        for fi40 := 0; fi40 < 14; fi40++ {
                            var v41 uint64 = 4294967295
                            v3[627] = "\n\t"
                            _ = fi40
                            _ = v41
                        }
                        v42 := GenS2{F0: 0, F1: 0}
                        for fi43 := 0; fi43 < 10; fi43++ {
                            var v44 string = "\n\t"
                            _ = fi43
                            _ = v44
                        }
                        _ = fi39
                        _ = v42
                    }
                    var v45 uint16 = v5
                    switch "foo bar" {
                    case "":
                        switch 1 {
                        case 42:
                            var v46 uint32 = 2147483647
                            _ = v46
                        default:
                            v47 := GenS2{F0: 52, F1: -894}
                            _ = v47
                        }
                    case "a":
                        switch 128 {
                        case -9223372036854775808:
                            var v48 bool = (2147483648 > 65535)
                            _ = v48
                        case 746:
                            fmt.Println(v37)
                        case -1:
                            if false {
                                fmt.Println(v2)
                            }
                        }
                    case "hello":
                        if false {
                            fmt.Println(v38, v38)
                            var v49 int16 = -458
                            _ = v49
                        }
                    default:
                        v50 := []int{}
                        _ = v50
                    }
                    _ = v45
                }
            } else {
                var v51 int64 = v37
                _ = v51
            }
            _ = v38
        }
        fmt.Println(v5)
        _ = v37
    } else {
        for fi52 := 0; fi52 < 15; fi52++ {
            for fi53 := 0; fi53 < 11; fi53++ {
                v5 = (v5 % ((v5) | 1))
                _ = fi53
            }
            var v54 string = ("foo bar" + "a")
            _ = fi52
            _ = v54
        }
    }
    fmt.Println(v1, v1)
    var v55 uint16 = (v5 ^ v5)
    for fi56 := 0; fi56 < 20; fi56++ {
        v2 += v2
        var v57 int16 = -340
        _ = fi56
        _ = v57
    }
    fmt.Println(v5, v1)
    for fi58 := 0; fi58 < 10; fi58++ {
        v59 := []int{}
        fmt.Println(fi58)
        _ = fi58
        _ = v59
    }
    var v60 bool = false
    var v61 float64 = 3.14159
    v62 := GenS2{F0: -1, F1: 42}
    v63 := GenS1{F0: (v5 - 2), F1: 0, F2: v60}
    fmt.Println(v5, v55)
    v55 = (v5 - (v55 % ((v5) | 1)))
    v5 = (^(v55))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
