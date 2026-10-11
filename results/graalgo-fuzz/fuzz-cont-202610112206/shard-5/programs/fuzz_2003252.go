package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int64, p1 int, p2 uint8) bool {
    if (1 >= (p0 + (-(p0)))) {
        var v4 int = 0
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := GenS1{F0: -1}
            _ = fi5
            _ = v6
        }
        _ = v4
    } else {
        if (true && true) {
            var v7 string = ""
            fmt.Println(v7)
            _ = v7
        }
    }
    v8 := []int{(^(p1))}
    if true {
        if true {
            switch p1 {
            case -1:
                v8 = append(v8, 127)
            case -2147483648:
                p1 -= (p1 ^ (p1 >> ((1) & 63)))
            default:
                var v9 int = (^(p1))
                _ = v9
            }
        }
        var v10 int = (p1 << ((65535) & 63))
        var v11 uint = 4294967295
        _ = v10
        _ = v11
    }
    var v12 int64 = p0
    v13 := GenS1{F0: 8}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v8
    return false
}

func main() {
    var v1 int16 = 1000
    var v2 bool = (-70 > 88)
    v3 := []string{"a"}
    v1 += v1
    fmt.Println(v2, v2)
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                v3 = append(v3, "")
                v3 = append(v3, "hello")
                _ = fi15
            }
            _ = fi14
        }
        v1 = v1
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 uint = 942
            _ = fi16
            _ = v17
        }
    } else {
        v1 += v1
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS1{F0: -128}
                v21 := GenS2{F0: 0}
                fmt.Println(v1, v2)
                var v22 string = (("" + "foo bar") + "a")
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = fi18
        }
        var v23 int16 = (^(v1))
        _ = v23
    }
    var v24 int = 128
    var v25 float32 = -2.25
    v1 = v1
    var v26 string = ""
    switch (v24 % ((v24) | 1)) {
    case 127:
        fmt.Println(v1)
    case -572:
        for fi27 := 0; fi27 < 18; fi27++ {
            v28 := GenS1{F0: -1}
            if (!v2) {
                v29 := GenS2{F0: 255}
                _ = v29
            } else {
                if true {
                    v30 := GenS3{F0: 127, F1: 2147483647}
                    _ = v30
                } else {
                    switch v26 {
                    case "":
                        var v31 int8 = 0
                        _ = v31
                    case "a":
                        var v32 uint16 = 0
                        _ = v32
                    }
                    if true {
                        v26 = "a"
                        var v33 bool = (-25 < -555)
                        _ = v33
                    }
                }
                if (1 != 18446744073709551615) {
                    v2 = false
                }
                v24 -= (v24 << ((220) & 63))
            }
            switch "a" {
            case "":
                var v34 uint8 = 11
                _ = v34
            }
            _ = fi27
            _ = v28
        }
    }
    var v35 bool = false
    switch v26 {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "\n\t":
        v3 = append(v3, v26)
    case "":
        v36 := GenS3{F0: 42, F1: 65535}
        _ = v36
    default:
        fmt.Println(v1)
    }
    switch v24 {
    case 312:
        if (255 <= 32767) {
            fmt.Println(v35)
        } else {
            if (0 != -9223372036854775808) {
                v1 -= (v1 % ((-1) | 1))
            } else {
                v37 := GenS1{F0: 42}
                _ = v37
            }
            var v38 int16 = (v1 | v1)
            for fi39 := 0; fi39 < 12; fi39++ {
                if v35 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v40 int64 = 111
                    v25 += (v25 * (v25 + -2.25))
                    _ = v40
                }
                _ = fi39
            }
            _ = v38
        }
    }
    var v41 float64 = (0.1 - ((float64(1.5) - 0.1) / 1.0))
    for fi42 := 0; fi42 < 8; fi42++ {
        _ = len(v3)
        v25 = (v25 * v25)
        v43 := GenS3{F0: -1, F1: 4294967295}
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        v25 += float32(3.14159)
        _ = fi44
    }
    v1 = v1
    var v45 uint32 = 0
    switch v24 {
    case 127:
        v46 := GenS4{F0: 9223372036854775807, F1: ((!v35) || v2)}
        _ = v46
    case 0:
        if (!((((v2 || v35) || v35) || v2) || (((v25 - -2.25) > float32(1e10)) && v2))) {
            v1 = v1
        } else {
            var v47 uint32 = v45
            _ = v47
        }
    case -772:
        v48 := GenS4{F0: 939, F1: v35}
        _ = v48
    default:
        fmt.Println(v45, v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v41
    _ = v45
}
