package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
func genFn0(p0 int32) bool {
    p0 -= (p0 >> ((0) & 63))
    fmt.Println(p0)
    _ = p0
    return true
}
func genFn1(p0 uint16, p1 uint8) int {
    v7 := GenS1{F0: 2147483647}
    for fi8 := 0; fi8 < 1; fi8++ {
        if (!false) {
            if (!(!(!true))) {
                var v9 uint32 = 2147483648
                _ = v9
            }
        }
        var v10 int16 = -1
        _ = fi8
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 52
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    v3 := make(map[int]uint32)
    var v4 int8 = 127
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    v4 -= v4
    switch 127 {
    case 128:
        for fi11 := 0; fi11 < 2; fi11++ {
            v4 += v4
            fmt.Println(v4)
            _ = fi11
        }
    case -559:
        v4 -= (v4 / ((v4) | 1))
    case -498:
        var v12 int8 = -1
        _ = v12
    default:
        var v13 string = "hello"
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint = 0
        var v16 string = "a"
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        var v17 int = 2147483647
        var v18 int8 = 42
        var v19 float32 = 0.0
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int8 = v4
            v23 := v1[("" + "hello")]
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            switch "" {
            case "a":
                fmt.Println(v4, fi20)
            default:
                fmt.Println(fi20, v2, v4)
            }
            var v25 int = (fi24 + (fi20 >> ((2) & 63)))
            _ = fi24
            _ = v25
        }
        v2 = false
        _ = fi20
    }
    if (18446744073709551615 == 4294967295) {
        fmt.Println(v2)
        v2 = false
    }
    switch 429 {
    case 42:
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 uint = 392
            _ = fi26
            _ = v27
        }
    case 127:
        for fi28 := 0; fi28 < 6; fi28++ {
            if ("hello" < "\n\t") {
                var v29 bool = (42 == 899)
                v29 = (1e10 <= 0.0)
                _ = v29
            } else {
                var v30 bool = true
                v31 := GenS2{F0: (false || false), F1: 65535}
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS1{F0: 1}
                    _ = fi32
                    _ = v33
                }
                _ = v30
                _ = v31
            }
            var v34 int32 = 0
            for fi35 := 0; fi35 < 16; fi35++ {
                v36 := []int{(fi35 >> ((255) & 63)), 4294967295}
                v37 := GenS1{F0: v34}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi28
            _ = v34
        }
    case -9223372036854775808:
        var v38 uint8 = 42
        _ = v38
    default:
        var v39 string = ("foo bar" + "a")
        _ = v39
    }
    v2 = v2
    var v40 int32 = 42
    if v2 {
        v4 -= (-(v4))
        delete(v6, 42)
    } else {
        for fi41 := 0; fi41 < 5; fi41++ {
            for fi42 := 0; fi42 < 12; fi42++ {
                v43 := GenS2{F0: (3.14159 < -2.25), F1: 42}
                _ = len(v3)
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
    }
    v44 := GenS3{F0: -77}
    var v45 uint16 = 32767
    if (v4 != v4) {
        if true {
            v2 = false
            for fi46 := 0; fi46 < 8; fi46++ {
                v4 = (v4 % ((v4) | 1))
                v47 := GenS4{F0: 65535, F1: 32767, F2: int8(v40)}
                _ = fi46
                _ = v47
            }
            var v48 int8 = v4
            _ = v48
        } else {
            for fi49 := 0; fi49 < 14; fi49++ {
                var v50 int16 = 1000
                _ = fi49
                _ = v50
            }
            v4 += 42
            delete(v1, (("a" + "a") + ("foo bar" + ("\n\t" + "a"))))
        }
    } else {
        v40 -= (^(v40))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
