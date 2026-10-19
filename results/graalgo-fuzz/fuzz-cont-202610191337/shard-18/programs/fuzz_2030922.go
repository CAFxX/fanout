package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int16 }
func genFn0(p0 int16, p1 int16, p2 uint32) int64 {
    var v4 int8 = -38
    var v5 int = 127
    v6 := GenS1{F0: "a"}
    if false {
        if (18446744073709551615 <= 106) {
            v7 := GenS1{F0: ("foo bar" + "hello")}
            v8 := GenS2{F0: (^(v5))}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int32 = -2147483648
                var v11 int = (fi9 * v5)
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        v12 := GenS1{F0: ("\n\t" + (("a" + "\n\t") + "a"))}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -649
}
func genFn1(p0 uint64) uint64 {
    if false {
        v13 := GenS2{F0: 1}
        v14 := GenS1{F0: "hello"}
        v15 := GenS3{F0: -1, F1: -1}
        _ = v13
        _ = v14
        _ = v15
    } else {
        fmt.Println(p0)
        for fi16 := 0; fi16 < 4; fi16++ {
            p0 = p0
            p0 -= (p0 << ((4294967295) & 63))
            switch (fi16 ^ 32767) {
            case 758:
                v17 := GenS2{F0: -973}
                _ = v17
            case 128:
                var v18 int8 = 0
                _ = v18
            case 191:
                var v19 float32 = 0.0
                _ = v19
            default:
                if false {
                    p0 = 1
                    p0 = p0
                    v20 := GenS4{F0: -1, F1: -1}
                    _ = v20
                }
            }
            _ = fi16
        }
        var v21 uint = 0
        _ = v21
    }
    v22 := GenS1{F0: ""}
    fmt.Println(p0)
    v23 := GenS2{F0: 32768}
    _ = p0
    _ = v22
    _ = v23
    return p0
}

func main() {
    var v1 bool = true
    var v2 int64 = 342
    v3 := []uint32{2147483648, 42, 4294967295}
    v24 := GenS5{F0: 214, F1: 61}
    v25 := GenS6{F0: 0, F1: "hello"}
    var v26 float64 = 3.14159
    if (!(!("hello" != ""))) {
        var v27 uint32 = 465
        _ = v27
    } else {
        var v28 int64 = v2
        fmt.Println(v2, v28)
        v29 := GenS7{F0: 32767}
        _ = v28
        _ = v29
    }
    var v30 int64 = (v2 + -414)
    var v31 uint8 = 24
    v32 := GenS2{F0: 127}
    fmt.Println(v30)
    if ("a" >= "") {
        v1 = false
        if v1 {
            var v33 uint32 = 690
            var v34 uint = 148
            if ((!(1 <= (v2 << ((v34) & 63)))) || v1) {
                v1 = (0 <= 18446744073709551615)
                v1 = v1
                v33 -= 528
            } else {
                var v35 uint64 = 4294967295
                v3 = append(v3, 0)
                var v36 int8 = -113
                _ = v35
                _ = v36
            }
            _ = v33
            _ = v34
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int8 = -128
                switch (fi37 - fi37) {
                case -862:
                    if len(v3) > 0 {
                        v3[0] = 89
                    }
                case 1:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -837:
                    switch "hello" {
                    case "foo bar":
                        _ = len(v3)
                    case "":
                        var v39 int64 = v30
                        _ = v39
                    default:
                        if v1 {
                            for fi40 := 0; fi40 < 19; fi40++ {
                                var v41 uint16 = 1
                                v31 = (v31 << ((4294967295) & 63))
                                _ = fi40
                                _ = v41
                            }
                            v1 = v1
                            fmt.Println(v38, v30)
                        }
                    }
                }
                fmt.Println(v2, v38, v1)
                _ = fi37
                _ = v38
            }
            if v1 {
                var v42 bool = v1
                _ = v42
            }
            var v43 string = ("hello" + "\n\t")
            _ = v43
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v44 uint64 = 42
    if (((-2147483648 != 609) && (v1 || v1)) && false) {
        var v45 bool = (441 < 908)
        var v46 uint32 = 2
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 bool = v1
            var v49 bool = false
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v44
}
