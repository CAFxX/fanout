package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 uint64, p1 int16) uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        p1 -= p1
        p0 = (p0 >> ((18446744073709551615) & 63))
        p1 -= (p1 | 42)
        _ = fi4
    }
    if true {
        var v5 string = ("" + (("a" + "a") + ("a" + (("a" + "foo bar") + "foo bar"))))
        p0 = 4294967295
        switch 0 {
        case 476:
            v6 := GenS1{F0: 18446744073709551615, F1: (p1 != -209), F2: p1}
            _ = v6
        case 42:
            p0 -= (p0 + 722)
        default:
            fmt.Println(p1, p0, p1)
        }
        _ = v5
    } else {
        v7 := GenS2{F0: 42, F1: (0.1 != (0.0 - (float32(3.14159) + ((1e10 - (0.1 + float32(0.1))) + (1e10 / 1.0))))), F2: 42}
        for fi8 := 0; fi8 < 10; fi8++ {
            p1 = (p1 * p1)
            _ = fi8
        }
        var v9 uint32 = 204
        _ = v7
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() bool {
    var v10 float32 = 1.5
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v10)
    case 42:
        for fi11 := 0; fi11 < 11; fi11++ {
            switch "foo bar" {
            case "a":
                v10 += 1e10
            case "":
                var v12 uint16 = 0
                _ = v12
            case "\n\t":
                v13 := GenS3{F0: 512, F1: 2147483647, F2: 128}
                _ = v13
            }
            _ = fi11
        }
    default:
        v14 := GenS2{F0: 0, F1: false, F2: 1}
        _ = v14
    }
    var v15 float32 = v10
    v15 = v10
    _ = v10
    _ = v15
    return false
}

func main() {
    var v1 int8 = 100
    var v2 float32 = 1.5
    var v3 bool = true
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if true {
                var v18 bool = v3
                var v19 int16 = -136
                if (2147483647 >= 520) {
                    v3 = v3
                    v20 := []int{(fi16 % ((1) | 1)), fi16, fi16, -505}
                    _ = v20
                }
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 uint = 497
            if v3 {
                fmt.Println(v22, v2, v22)
            }
            v2 -= float32(-2.25)
            _ = fi21
            _ = v22
        }
        var v23 int32 = 0
        _ = fi16
        _ = v23
    }
    switch "\n\t" {
    case "":
        switch "a" {
        case "foo bar":
            if v3 {
                switch "foo bar" {
                case "hello":
                    var v24 int = 42
                    _ = v24
                case "foo bar":
                    var v25 uint32 = 0
                    _ = v25
                }
                v2 += 0.1
                if v3 {
                    var v26 uint = 4294967295
                    for fi27 := 0; fi27 < 15; fi27++ {
                        for fi28 := 0; fi28 < 7; fi28++ {
                            if (!v3) {
                                v1 = (v1 - 30)
                                var v29 string = ""
                                _ = v29
                            } else {
                                fmt.Println(fi28, v3)
                                var v30 uint32 = 521
                                fmt.Println(v2, v2)
                                v26 += v26
                                fmt.Println(v2, v2, v2)
                                _ = v30
                            }
                            v31 := GenS4{F0: "\n\t", F1: 42}
                            _ = fi28
                            _ = v31
                        }
                        v32 := []int{(fi27 % ((fi27) | 1))}
                        _ = fi27
                        _ = v32
                    }
                    v2 = v2
                    _ = v26
                }
            } else {
                var v33 int64 = -1
                _ = v33
            }
        }
    default:
        var v34 string = "hello"
        _ = v34
    }
    v3 = v3
    v2 += v2
    var v35 bool = false
    switch 1 {
    case 32768:
        v36 := []int{}
        _ = v36
    case 1:
        var v37 int = 894
        _ = v37
    case 0:
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 uint8 = 0
            v40 := GenS1{F0: 0, F1: true, F2: 0}
            v2 -= (v2 - (1e10 - v2))
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        v1 += (v1 << ((genFn0(367, 0)) & 63))
        for fi42 := 0; fi42 < 8; fi42++ {
            if v3 {
                var v43 bool = v35
                _ = v43
            } else {
                v44 := GenS4{F0: ("\n\t" + (("foo bar" + "a") + ("foo bar" + ("" + ("" + "foo bar"))))), F1: 255}
                _ = v44
            }
            _ = fi42
        }
        _ = fi41
    }
    v35 = (v35 && false)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
}
