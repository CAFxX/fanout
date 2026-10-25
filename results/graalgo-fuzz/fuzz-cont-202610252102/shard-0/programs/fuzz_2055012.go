package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 string; F2 string }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0() int16 {
    var v5 bool = (243 <= 849)
    v5 = true
    v5 = true
    var v6 int = 128
    _ = v5
    _ = v6
    return -164
}
func genFn1(p0 uint) uint {
    var v7 uint32 = 340
    fmt.Println(p0, p0)
    p0 = 1
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: -1}
        if true {
            var v10 uint64 = 0
            switch 128 {
            case 4294967295:
                if (32767 == 32767) {
                    if true {
                        var v11 uint8 = 127
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 6; fi12++ {
                            if false {
                                p0 = 2
                                p0 = p0
                            }
                            if true {
                                fmt.Println(p0)
                                v7 += v7
                            }
                            _ = fi12
                        }
                        if false {
                            if (true && false) {
                                p0 = 9
                                v13 := GenS2{F0: (v10 - (^(v10)))}
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    v15 := GenS3{F0: 0, F1: "", F2: "foo bar"}
                                    if (!false) {
                                        v16 := GenS1{F0: -1}
                                        var v17 uint16 = 0
                                        _ = v16
                                        _ = v17
                                    }
                                    switch fi14 {
                                    case 42:
                                        var v18 string = "hello"
                                        _ = v18
                                    case -1:
                                        p0 += (p0 | (p0 % ((p0) | 1)))
                                    }
                                    p0 += (p0 + (p0 % ((p0) | 1)))
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v13
                            }
                        } else {
                            p0 += p0
                            v19 := GenS4{F0: 641, F1: 0, F2: v7}
                            fmt.Println(p0)
                            _ = v19
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                v7 += uint32(p0)
                                v7 += v7
                            } else {
                                fmt.Println(p0, v10)
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    v21 := GenS5{F0: (v7 * 1), F1: (v7 & v7), F2: 1}
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        case "hello":
                            v22 := GenS4{F0: 42, F1: 42, F2: (^(v7))}
                            _ = v22
                        }
                    }
                }
            case 542:
                var v23 float32 = 3.14159
                _ = v23
            case -1:
                v24 := GenS6{F0: 773, F1: 1}
                _ = v24
            }
            _ = v10
        }
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS7{F0: ((-(fi8)) > fi8), F1: 373}
            v27 := GenS8{F0: (^(v7)), F1: 11, F2: -901}
            if (!false) {
                var v28 int16 = 1000
                var v29 uint64 = 1
                v30 := []int{fi25, fi8}
                _ = v28
                _ = v29
                _ = v30
            } else {
                v31 := []int{(fi25 * fi8), (fi25 * (fi8 << ((p0) & 63))), 806, 42}
                _ = v31
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi8
        _ = v9
    }
    v32 := GenS3{F0: -546, F1: "", F2: "hello"}
    _ = p0
    _ = v32
    _ = v7
    return p0
}

func main() {
    var v1 int32 = -1
    v2 := []int{32767, 42, 128, 42}
    v3 := []int{}
    var v4 int8 = 111
    v4 -= v4
    v4 = (v4 - v4)
    var v33 int16 = -32768
    fmt.Println(v4, v33)
    if true {
        var v34 int16 = (^(v33))
        v35 := GenS5{F0: 2, F1: 4294967295, F2: 666}
        var v36 uint = genFn1(0)
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v4, v4)
    var v37 string = ((("" + "foo bar") + (("hello" + "foo bar") + "")) + "")
    v33 = v33
    switch "foo bar" {
    case "a":
        var v38 uint8 = 227
        _ = v38
    case "":
        if true {
            var v39 string = v37
            _ = v39
        } else {
            var v40 uint16 = 0
            switch ("foo bar" + "\n\t") {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            }
            _ = v40
        }
    case "\n\t":
        if true {
            v1 = 1
            var v41 bool = true
            _ = v41
        }
    }
    v4 -= -26
    v1 = (v1 << ((18446744073709551615) & 63))
    var v42 string = v37
    fmt.Println(v42, v42)
    if true {
        for fi43 := 0; fi43 < 12; fi43++ {
            v4 += (v4 / ((-1) | 1))
            _ = fi43
        }
    } else {
        _ = len(v2)
        var v44 int8 = (v4 / ((127) | 1))
        fmt.Println(v44, v42, v44)
        _ = v44
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        for fi46 := 0; fi46 < 4; fi46++ {
            var v47 uint16 = 83
            _ = fi46
            _ = v47
        }
        _ = fi45
    }
    fmt.Println(v42, v37, v42)
    v33 -= (v33 >> ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v42
}
