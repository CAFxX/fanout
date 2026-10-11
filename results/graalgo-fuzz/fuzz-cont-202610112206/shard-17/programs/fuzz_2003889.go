package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int8; F1 int32; F2 uint }
type GenS7 struct { F0 int8; F1 uint8 }
func genFn0(p0 bool) int {
    p0 = true
    switch ("\n\t" + ("hello" + "\n\t")) {
    case "a":
        v4 := []int{-9223372036854775808, 0, 32768, 1}
        _ = v4
    case "foo bar":
        v5 := []int{42}
        _ = v5
    }
    v6 := []int{}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            if len(v6) > 0 {
                v6[0] = 1
            }
            _ = fi8
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        _ = len(v6)
        if true {
            p0 = true
        }
        switch ("a" + ((("foo bar" + "hello") + ("a" + ("a" + "a"))) + "hello")) {
        case "hello":
            p0 = true
        case "a":
            p0 = ((-2.25 + (0.1 - -2.25)) < (3.14159 + 3.14159))
        case "\n\t":
            if ((((((0.1 / (1e10 - 0.0)) - (-2.25 * -2.25)) + 1e10) / (1.5 / -2.25)) / 1.0) >= 1e10) {
                switch fi9 {
                case -903:
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                case 127:
                    var v10 int64 = 1
                    _ = v10
                default:
                    if len(v6) > 0 {
                        v6[0] = fi9
                    }
                }
                v11 := GenS1{F0: 0, F1: 19}
                v6 = append(v6, fi9)
                _ = v11
            }
        }
        _ = fi9
    }
    _ = p0
    _ = v6
    return 32768
}
func genFn1(p0 uint8, p1 int64, p2 bool) uint8 {
    if p2 {
        if p2 {
            fmt.Println(p0)
        }
        v12 := []int{-9223372036854775808, 0}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v12
    } else {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int16 = -32768
            _ = fi13
            _ = v14
        }
        p2 = (1e10 < 0.1)
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 int64 = p1
            v17 := GenS1{F0: 32767, F1: -1}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v18 := []int{}
    for fi19 := 0; fi19 < 14; fi19++ {
        v18 = append(v18, fi19)
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return (p0 * 42)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 0
    var v3 int16 = -1
    switch ("hello" + "hello") {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            switch genFn0(false) {
            case 0:
                fmt.Println(v3, v3)
            case 128:
                var v20 string = "\n\t"
                _ = v20
            }
        case "a":
            var v21 string = "a"
            _ = v21
        }
    case "":
        v22 := GenS2{F0: 0, F1: -9223372036854775808}
        _ = v22
    case "hello":
        fmt.Println(v2, v2)
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 bool = true
            switch ("hello" + ("foo bar" + ("\n\t" + "\n\t"))) {
            case "hello":
                v2 -= (v2 % ((18446744073709551615) | 1))
            case "a":
                var v26 uint32 = 880
                fmt.Println(fi24, v25, v25)
                var v27 uint32 = 2147483648
                _ = v26
                _ = v27
            }
            v2 = v2
            _ = fi24
            _ = v25
        }
        v2 += (v2 / ((v2) | 1))
        var v28 uint32 = 27
        _ = fi23
        _ = v28
    }
    v2 = v2
    v29 := GenS2{F0: 0, F1: genFn0(true)}
    v30 := GenS3{F0: v2, F1: -2147483648}
    var v31 uint16 = 65535
    v32 := GenS4{F0: 2147483648}
    switch 427 {
    case -2147483648:
        fmt.Println(v2)
    case 2147483647:
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 bool = (!false)
            if v34 {
                var v35 uint = 666
                for fi36 := 0; fi36 < 5; fi36++ {
                    switch ("a" + "hello") {
                    case "foo bar":
                        var v37 uint = (v35 << ((v35) & 63))
                        _ = v37
                    case "hello":
                        var v38 float32 = (float32(0.1) + 0.1)
                        _ = v38
                    }
                    _ = fi36
                }
                _ = v35
            } else {
                fmt.Println(v31, v31)
                v2 = (v2 + v2)
            }
            var v39 uint = 42
            _ = fi33
            _ = v34
            _ = v39
        }
    default:
        v31 += 2
    }
    switch ((("foo bar" + (("a" + "\n\t") + "foo bar")) + "hello") + "hello") {
    case "a":
        var v40 float32 = 0.1
        _ = v40
    case "\n\t":
        fmt.Println(v3)
    }
    v41 := GenS5{F0: false, F1: 696}
    switch 4294967295 {
    case 1:
        for fi42 := 0; fi42 < 9; fi42++ {
            for fi43 := 0; fi43 < 13; fi43++ {
                switch ("" + "\n\t") {
                case "foo bar":
                    v44 := GenS6{F0: 0, F1: 0, F2: uint(v2)}
                    _ = v44
                case "":
                    switch 0 {
                    case 9223372036854775807:
                        var v45 int64 = int64(v3)
                        _ = v45
                    case 42:
                        var v46 uint8 = genFn1(2, -1, (-32768 < v3))
                        _ = v46
                    case 4294967295:
                        v47 := v1[2]
                        _ = v47
                    default:
                        v48 := GenS5{F0: (float32(1e10) != float32(0.1)), F1: 668}
                        _ = v48
                    }
                case "\n\t":
                    var v49 float32 = (-2.25 + 1.5)
                    _ = v49
                default:
                    var v50 uint16 = (v31 / ((v31) | 1))
                    _ = v50
                }
                v31 += 32767
                v2 = (v2 * 827)
                v51 := GenS2{F0: 42, F1: (^(fi43))}
                if true {
                    if (!(!true)) {
                        var v52 int = fi43
                        _ = v52
                    }
                    v53 := GenS5{F0: true, F1: 4294967295}
                    fmt.Println(v2, v31, v2)
                    _ = v53
                }
                _ = fi43
                _ = v51
            }
            var v54 uint = 255
            var v55 uint64 = v2
            _ = fi42
            _ = v54
            _ = v55
        }
    case -1:
        v3 = (^(v3))
    default:
        fmt.Println(v3)
    }
    if true {
        v2 = (v2 % (((v2 >> ((255) & 63))) | 1))
    }
    for fi56 := 0; fi56 < 19; fi56++ {
        fmt.Println(v2, fi56, v3)
        _ = fi56
    }
    fmt.Println(v2, v31)
    delete(v1, 0)
    switch "a" {
    case "\n\t":
        v57 := GenS3{F0: (v2 | (v2 - v2)), F1: 2147483647}
        _ = v57
    default:
        if false {
            v58 := GenS7{F0: 1, F1: 255}
            v3 -= (v3 & (v3 * (v3 - 42)))
            _ = v58
        } else {
            v59 := []int{}
            _ = v59
        }
    }
    v3 += (-(v3))
    var v60 int = genFn0(false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v41
    _ = v60
}
