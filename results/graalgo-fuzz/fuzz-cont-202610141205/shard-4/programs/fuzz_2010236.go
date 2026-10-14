package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    var v3 int = 127
    var v4 bool = ((127 <= -78) && false)
    v5 := make(map[uint32]int)
    var v6 int16 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        if v4 {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint16 = 42
                var v10 uint64 = 0
                v9 = v9
                _ = fi8
                _ = v10
                _ = v9
            }
            switch v2 {
            case "foo bar":
                var v11 uint32 = 4294967295
                _ = v11
            case "\n\t":
                switch v3 {
                case 0:
                    var v12 float32 = 1e10
                    _ = v12
                case 565:
                    v5[2] = v3
                case 32767:
                    switch -2147483648 {
                    case 32768:
                        var v13 int16 = -32768
                        _ = v13
                    case 128:
                        v14 := GenS1{F0: -601}
                        _ = v14
                    case 127:
                        var v15 int = (fi7 % ((2147483647) | 1))
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 2; fi16++ {
                            var v17 float64 = 1e10
                            v18 := GenS1{F0: -561}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                }
            }
            v4 = (v3 > fi7)
        }
        _ = fi7
    }
    v19 := []int{v3, (v3 % ((v3) | 1))}
    if v4 {
        v2 = v2
        v20 := GenS1{F0: 100}
        _ = v20
    }
    fmt.Println(v2, v6, v4)
    v2 = (("\n\t" + v2) + ("" + ("hello" + v2)))
    v21 := GenS1{F0: 0}
    for fi22 := 0; fi22 < 17; fi22++ {
        v3 -= (v3 << ((65535) & 63))
        var v23 uint64 = 42
        _ = fi22
        _ = v23
    }
    if (v4 && v4) {
        v6 = (v6 + 42)
        _ = len(v19)
        v19 = append(v19, -2147483648)
    } else {
        switch v2 {
        case "\n\t":
            if ((!(!false)) && v4) {
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int = 128
                    var v26 string = ("\n\t" + v2)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 string = v2
                    v29 := GenS1{F0: 0}
                    fmt.Println(v6, v4)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v6 = (v6 << ((683) & 63))
            } else {
                if (18446744073709551615 <= 42) {
                    v30 := GenS2{F0: v4, F1: 170, F2: 2}
                    v6 = v6
                    var v31 uint64 = 97
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        _ = len(v5)
                        _ = fi32
                    }
                }
            }
        default:
            if v4 {
                v4 = true
            } else {
                v2 = "foo bar"
            }
        }
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        fmt.Println(v6)
        if (65535 > 0) {
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 int64 = 9223372036854775807
                _ = fi34
                _ = v35
            }
            if v4 {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                _ = len(v19)
                for fi36 := 0; fi36 < 4; fi36++ {
                    v37 := GenS3{F0: v4, F1: 2, F2: 1}
                    var v38 uint = 594
                    if (0 != 42) {
                        switch v2 {
                        case "":
                            switch ((v2 + ("a" + (v2 + (v2 + (((("foo bar" + (v2 + "\n\t")) + ("\n\t" + ("\n\t" + v2))) + v2) + (("a" + "hello") + ("hello" + v2))))))) + (("foo bar" + v2) + "hello")) {
                            case "hello":
                                switch -530 {
                                case 0:
                                    var v39 int64 = 1
                                    _ = v39
                                case -11:
                                    var v40 string = "a"
                                    _ = v40
                                default:
                                    v4 = true
                                }
                            }
                        }
                        var v41 string = ("hello" + "\n\t")
                        _ = v41
                    } else {
                        var v42 int32 = 1
                        delete(v5, 1)
                        _ = v42
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
        } else {
            var v43 int64 = 595
            _ = v43
        }
        _ = fi33
    }
    v44 := GenS4{F0: 879}
    for fi45 := 0; fi45 < 16; fi45++ {
        v6 -= v6
        if len(v19) > 0 {
            v19[0] = 1
        }
        v4 = v4
        _ = fi45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
