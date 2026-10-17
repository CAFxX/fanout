package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint32) int32 {
    if (false || false) {
        p1 = 1
    } else {
        if false {
            for fi7 := 0; fi7 < 6; fi7++ {
                switch (fi7 / (((fi7 & fi7)) | 1)) {
                case 367:
                    if true {
                        for fi8 := 0; fi8 < 8; fi8++ {
                            v9 := GenS1{F0: 0, F1: "a"}
                            p1 -= p1
                            p1 = (^(p1))
                            var v10 int16 = -158
                            v11 := GenS1{F0: 127, F1: "a"}
                            _ = fi8
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                    }
                }
                _ = fi7
            }
            var v12 int32 = (p0 | (p0 >> ((6) & 63)))
            p0 += (p0 / (((p0 % (((p0 - p0)) | 1))) | 1))
            _ = v12
        }
    }
    v13 := GenS2{F0: (42 != 0), F1: 4294967295, F2: true}
    p1 -= (^(p1))
    var v14 int8 = -128
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return p0
}
func genFn1(p0 int64, p1 int) string {
    var v15 uint32 = 2
    v16 := []int{}
    switch "\n\t" {
    case "a":
        v15 = (v15 * 1)
    case "\n\t":
        for fi17 := 0; fi17 < 11; fi17++ {
            switch ("\n\t" + "a") {
            case "foo bar":
                if true {
                    if len(v16) > 0 {
                        v16[0] = p1
                    }
                    if ((int64(fi17) >= (p0 * (p0 - int64(v15)))) || (true || (!true))) {
                        var v18 uint8 = 2
                        _ = v18
                    } else {
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        var v19 float64 = float64((1e10 + 0.1))
                        v15 -= (^(v15))
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v15 -= (v15 % (((^(v15))) | 1))
                        v15 = 0
                        _ = fi20
                    }
                }
            }
            if true {
                v21 := GenS1{F0: 1, F1: (("foo bar" + "") + ((("foo bar" + "a") + "\n\t") + ((("\n\t" + "\n\t") + "a") + ("foo bar" + "a"))))}
                _ = v21
            } else {
                var v22 uint16 = 255
                _ = v22
            }
            _ = fi17
        }
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return ""
}

func main() {
    v1 := []uint32{0, 1, 139}
    var v2 int = 32767
    v3 := make(map[uint32]string)
    var v4 int16 = 32767
    v5 := make(map[uint32]int)
    var v6 bool = (true || true)
    v23 := GenS3{F0: 0}
    var v24 uint32 = 65535
    v25 := GenS4{F0: 0}
    v6 = v6
    if (0 > -161) {
        if (100 > 127) {
            var v26 uint = 65535
            v26 += v26
            var v27 int64 = -9223372036854775808
            _ = v26
            _ = v27
        } else {
            var v28 float32 = 1e10
            _ = v28
        }
        v29 := GenS2{F0: v6, F1: 17, F2: (127 != -91)}
        fmt.Println(v6, v4)
        _ = v29
    } else {
        v2 -= -750
        if v6 {
            switch "hello" {
            case "\n\t":
                var v30 uint32 = 1
                _ = v30
            default:
                var v31 int16 = (v4 * 595)
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 15; fi32++ {
                v2 = v2
                fmt.Println(v6)
                _ = fi32
            }
            _ = len(v1)
        }
        v4 -= 473
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        for fi34 := 0; fi34 < 9; fi34++ {
            if (v6 && ((1.5 + (1.5 / 1.0)) > (1.5 / 1e10))) {
                v35 := GenS2{F0: v6, F1: 2, F2: v6}
                _ = v35
            } else {
                if ("a" >= "") {
                    switch fi34 {
                    case -1:
                        fmt.Println(v4)
                    }
                }
                v6 = v6
            }
            v2 += 9223372036854775807
            _ = fi34
        }
        _ = fi33
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v37 := GenS5{F0: 132, F1: 216}
        v38 := GenS6{F0: (v24 << ((492) & 63))}
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v24, v4)
    _ = len(v5)
    switch "hello" {
    case "hello":
        v39 := GenS4{F0: 315}
        _ = v39
    }
    v6 = v6
    v4 = 42
    fmt.Println(v2, v2, v24)
    var v40 int16 = v4
    v6 = true
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
