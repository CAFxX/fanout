package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint16; F1 int32; F2 int64 }
type GenS8 struct { F0 uint32; F1 uint64; F2 uint }

func main() {
    var v1 uint8 = 108
    v2 := []uint32{4294967295, 2147483648, 2, 42}
    v3 := []int{42, 32768, -863, 4294967295}
    var v4 int64 = 42
    var v5 int8 = 127
    var v6 int = 42
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint8 = v1
        v4 = (-(v4))
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    if (false || (-32768 >= 1)) {
        if (749 != uint64(v1)) {
            v9 := GenS1{F0: 48, F1: v1, F2: (true || ((!false) && ((v5 / ((127) | 1)) <= -1)))}
            var v10 float32 = -2.25
            _ = len(v3)
            _ = v10
            _ = v9
        }
        fmt.Println(v1, v4)
        if true {
            v6 = v6
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int8 = (v5 << ((18446744073709551615) & 63))
                var v13 int8 = v5
                var v14 int32 = -1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            v5 = v5
        }
    } else {
        v15 := GenS2{F0: -32768, F1: 1000, F2: "hello"}
        _ = v15
    }
    v16 := GenS3{F0: (v6 >> ((330) & 63)), F1: 1}
    var v17 uint32 = 602
    fmt.Println(v4, v5)
    for fi18 := 0; fi18 < 17; fi18++ {
        if (v1 > (^(v1))) {
            fmt.Println(v5, v1, fi18)
            for fi19 := 0; fi19 < 5; fi19++ {
                if (true && true) {
                    if true {
                        v5 = v5
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v17 = uint32(v1)
                    }
                    v20 := GenS4{F0: 255, F1: (v1 & v1), F2: 1}
                    _ = v20
                }
                switch fi18 {
                case -2147483648:
                    v5 = v5
                }
                _ = fi19
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v4 += v4
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int64 = 1
                v24 := GenS5{F0: 0, F1: -32768}
                _ = fi22
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS6{F0: v17, F1: (fi25 >> ((811) & 63))}
                v27 := GenS7{F0: 969, F1: -2147483648, F2: v4}
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := GenS8{F0: v17, F1: 852, F2: 0}
            _ = v28
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            if len(v3) > 0 {
                v3[0] = 0
            }
            v1 -= 62
            v30 := GenS8{F0: v17, F1: 18446744073709551615, F2: 2}
            if true {
                v2 = append(v2, 197)
                switch fi29 {
                case 42:
                    var v31 uint8 = (^(v1))
                    _ = v31
                default:
                    if (v1 != 2) {
                        if ("foo bar" <= "foo bar") {
                            fmt.Println(fi29, v6, v4)
                            v32 := GenS5{F0: (fi29 << ((18446744073709551615) & 63)), F1: 0}
                            v6 = 32768
                            var v33 uint32 = (v17 % ((v17) | 1))
                            _ = v32
                            _ = v33
                        }
                        if false {
                            var v34 float64 = 3.14159
                            v4 -= v4
                            var v35 int32 = 663
                            _ = v34
                            _ = v35
                        } else {
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    v5 = 127
                                    var v36 int16 = 42
                                    _ = v36
                                }
                            }
                            var v37 float32 = 1.5
                            var v38 uint32 = 2
                            _ = v37
                            _ = v38
                        }
                        var v39 int32 = 831
                        _ = v39
                    }
                }
            }
            _ = fi29
            _ = v30
        }
        _ = fi18
    }
    _ = len(v2)
    if true {
        var v40 uint8 = v1
        _ = len(v3)
        _ = v40
    } else {
        v6 += v6
        v4 = 9223372036854775807
    }
    if len(v2) > 0 {
        v2[0] = 173
    }
    if (567 < 2147483648) {
        v1 = 2
    }
    var v41 uint64 = 42
    if len(v3) > 0 {
        v3[0] = int(v41)
    }
    var v42 int8 = v5
    if (2 < 0) {
        switch "\n\t" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = (v6 % ((v6) | 1))
            }
        }
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        var v44 int64 = v4
        switch 901 {
        case 32768:
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 int16 = 32767
                fmt.Println(v46, v41)
                _ = fi45
                _ = v46
            }
        default:
            v1 = (^(v1))
        }
        switch "foo bar" {
        case "a":
            for fi47 := 0; fi47 < 9; fi47++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                v6 = (^(v6))
                _ = fi47
            }
        }
        _ = fi43
        _ = v44
    }
    v41 = (v41 | v41)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
