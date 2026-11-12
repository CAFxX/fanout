package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 bool; F1 uint32 }
type GenS7 struct { F0 string; F1 uint16; F2 int }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 bool = (350 != 610)
    var v2 int16 = -32768
    var v3 bool = false
    var v4 int8 = 100
    v5 := GenS1{F0: (v2 | (v2 % (((v2 | (v2 & (-(v2))))) | 1)))}
    switch "a" {
    case "\n\t":
        var v6 uint32 = 0
        _ = v6
    case "hello":
        fmt.Println(v4, v4, v4)
    }
    v7 := GenS1{F0: v2}
    v1 = v3
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "hello" {
        case "a":
            v2 += (v2 % ((v2) | 1))
        case "foo bar":
            fmt.Println(v1)
        case "hello":
            v9 := []int{(fi8 ^ 583)}
            _ = v9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v4 = (v4 - (v4 << ((255) & 63)))
        if v1 {
            if (255 == 255) {
                switch (fi10 << ((0) & 63)) {
                case 209:
                    switch -145 {
                    case 42:
                        v11 := GenS2{F0: 148, F1: "a", F2: 65535}
                        _ = v11
                    }
                case 42:
                    v3 = v1
                case 32767:
                    if v1 {
                        var v12 int32 = 2147483647
                        _ = v12
                    }
                }
                for fi13 := 0; fi13 < 14; fi13++ {
                    if (42 <= 0) {
                        var v14 int16 = v2
                        _ = v14
                    }
                    v15 := []int{fi13}
                    _ = fi13
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := []int{(fi16 & (fi10 - int(v2))), (fi10 & -9223372036854775808), fi16}
                    if (true || v3) {
                        v18 := GenS3{F0: 469, F1: "a"}
                        v2 += v2
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
        _ = fi10
    }
    v19 := []int{-950, -1, -1, -94}
    _ = len(v19)
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 int32 = 682
        var v22 uint8 = 54
        v1 = (857 == 273)
        _ = fi20
        _ = v21
        _ = v22
    }
    if v3 {
        switch -9223372036854775808 {
        case -2147483648:
            v23 := GenS4{F0: (v2 >> ((742) & 63)), F1: -128}
            _ = v23
        case 854:
            v24 := GenS5{F0: "a", F1: 0}
            _ = v24
        case 0:
            if v1 {
                for fi25 := 0; fi25 < 5; fi25++ {
                    if (v4 != v4) {
                        var v26 int = (fi25 | -475)
                        v4 += v4
                        for fi27 := 0; fi27 < 19; fi27++ {
                            var v28 int32 = -859
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    }
                    _ = fi25
                }
                if (42 < 0) {
                    v29 := GenS1{F0: -834}
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v4 -= 122
                        _ = fi30
                    }
                    v4 += 1
                }
            } else {
                v2 = 0
                v4 -= (v4 & (v4 | -1))
                if len(v19) > 0 {
                    v19[0] = 42
                }
            }
        }
        v2 = -1
    } else {
        v31 := GenS6{F0: v3, F1: 2147483648}
        if (!v1) {
            v1 = ((v1 && v3) && true)
            v32 := GenS1{F0: (v2 - (^(v2)))}
            if (1 < -853) {
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 uint8 = 0
                    v35 := GenS7{F0: "a", F1: 1, F2: (fi33 * -1)}
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    var v37 string = "\n\t"
                    _ = fi36
                    _ = v37
                }
            }
            _ = v32
        }
        if v1 {
            var v38 int = -9223372036854775808
            v19 = append(v19, -505)
            v3 = false
            _ = v38
        } else {
            for fi39 := 0; fi39 < 19; fi39++ {
                v1 = true
                for fi40 := 0; fi40 < 6; fi40++ {
                    if v3 {
                        var v41 float32 = float32(-2.25)
                        var v42 bool = (0 > 1)
                        v3 = v1
                        _ = v41
                        _ = v42
                    }
                    if (!(true && false)) {
                        switch ("foo bar" + "hello") {
                        case "hello":
                            if v1 {
                                var v43 string = (("foo bar" + "a") + "\n\t")
                                v44 := GenS4{F0: (v2 * v2), F1: v4}
                                _ = v43
                                _ = v44
                            }
                        case "foo bar":
                            v45 := GenS3{F0: 2, F1: "\n\t"}
                            _ = v45
                        default:
                            fmt.Println(fi40, fi39, v1)
                        }
                    }
                    v3 = true
                    v46 := GenS8{F0: 0}
                    _ = fi40
                    _ = v46
                }
                if v3 {
                    if (4294967295 == 4294967295) {
                        if v1 {
                            fmt.Println(v1, v1)
                            var v47 uint = 187
                            _ = v47
                        }
                    }
                }
                _ = fi39
            }
            var v48 uint64 = 0
            _ = v48
        }
        _ = v31
    }
    if len(v19) > 0 {
        v19[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
