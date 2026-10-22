package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }
type GenS7 struct { F0 int64; F1 uint }
func genFn0(p0 uint) uint16 {
    if (((!(2147483648 <= 65535)) || false) && false) {
        p0 += (p0 >> (((p0 | p0)) & 63))
        p0 = (p0 / ((531) | 1))
        switch 32768 {
        case 42:
            p0 -= (^(p0))
        }
    } else {
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 -= p0
        p0 = p0
        p0 = (^(p0))
        _ = fi8
    }
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1(p0 uint64, p1 int32, p2 uint) int16 {
    v9 := GenS1{F0: 2147483648, F1: (p1 % ((p1) | 1))}
    var v10 uint16 = genFn0((p2 & 65535))
    if (42 != -578) {
        var v11 int = 0
        _ = v11
    } else {
        p2 = 572
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int64 = 42
            switch ("\n\t" + "") {
            case "foo bar":
                if true {
                    var v14 int8 = -128
                    p1 += p1
                    var v15 uint16 = (v10 - (v10 - (v10 >> ((513) & 63))))
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS2{F0: v10, F1: 0, F2: p1}
                    _ = v16
                }
            default:
                p1 += (p1 * p1)
            }
            v10 = v10
            _ = fi12
            _ = v13
        }
        switch 4294967295 {
        case -615:
            v17 := GenS3{F0: -19, F1: 9223372036854775807}
            _ = v17
        }
    }
    var v18 uint16 = v10
    var v19 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v18
    _ = v19
    _ = v9
    return 1
}

func main() {
    v1 := []string{"a", "a"}
    v2 := []uint32{4294967295, 65535, 42}
    var v3 uint32 = 817
    v4 := []uint32{597}
    var v5 uint16 = 36
    var v6 int64 = -194
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 642 {
    case -1:
        for fi20 := 0; fi20 < 16; fi20++ {
            if false {
                var v21 int = (-(fi20))
                v22 := GenS4{F0: (v3 ^ (v3 - 281)), F1: 1}
                v3 = 1
                _ = v21
                _ = v22
            } else {
                _ = len(v2)
            }
            _ = fi20
        }
    case 1:
        var v23 uint16 = (v5 >> ((675) & 63))
        _ = v23
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        if false {
            if false {
                var v24 string = ("hello" + "foo bar")
                v25 := GenS3{F0: 1, F1: v6}
                _ = v24
                _ = v25
            }
            if ((3.14159 - 1e10) >= 1.5) {
                v5 = genFn0(2)
            }
            v5 = (^(v5))
        } else {
            var v26 uint16 = (v5 % (((^(v5))) | 1))
            var v27 uint = 0
            for fi28 := 0; fi28 < 8; fi28++ {
                v4 = append(v4, v3)
                fmt.Println(v27, v6)
                if (!((v27 + v27) < (^(v27)))) {
                    for fi29 := 0; fi29 < 8; fi29++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v26, fi29)
                        var v30 int = (fi28 << (((v27 & 207)) & 63))
                        var v31 uint64 = 4294967295
                        v31 = (v31 << (((^(v27))) & 63))
                        switch "hello" {
                        case "a":
                            v2 = append(v2, 4294967295)
                        case "\n\t":
                            v3 -= v3
                        default:
                            switch "hello" {
                            case "":
                                _ = len(v2)
                            case "a":
                                switch (-(fi29)) {
                                case 2147483647:
                                    v27 += uint(v5)
                                }
                                var v32 uint8 = 0
                                var v33 float32 = 1.5
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    if true {
                        var v34 int = fi28
                        v35 := GenS4{F0: (v3 - 4294967295), F1: 4294967295}
                        _ = len(v4)
                        _ = v34
                        _ = v35
                    }
                    _ = len(v2)
                } else {
                    var v36 int = fi28
                    var v37 uint8 = uint8(v26)
                    _ = v36
                    _ = v37
                }
                _ = fi28
            }
            _ = v26
            _ = v27
        }
    default:
        for fi38 := 0; fi38 < 5; fi38++ {
            fmt.Println(fi38)
            var v39 int64 = v6
            v40 := GenS4{F0: (v3 % ((495) | 1)), F1: 2}
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    v5 -= v5
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 int = fi41
        for fi43 := 0; fi43 < 12; fi43++ {
            var v44 string = "hello"
            _ = fi43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    v45 := GenS2{F0: 32767, F1: 0, F2: -1}
    var v46 uint32 = v3
    v5 += v5
    v47 := GenS5{F0: false, F1: "a"}
    v48 := GenS6{F0: 37, F1: v6, F2: 18446744073709551615}
    fmt.Println(v5, v6)
    v49 := GenS3{F0: 42, F1: v6}
    var v50 uint8 = 161
    v50 = 0
    if (1 == -9223372036854775808) {
        v51 := GenS7{F0: v6, F1: 2}
        _ = v51
    }
    _ = len(v2)
    fmt.Println(v46, v50, v46)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
