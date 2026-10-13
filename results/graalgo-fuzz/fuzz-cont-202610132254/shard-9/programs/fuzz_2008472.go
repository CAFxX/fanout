package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 string; F1 int16; F2 bool }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
func genFn0() int64 {
    v4 := GenS1{F0: -723, F1: 575}
    var v5 float32 = (0.0 - float32(1.5))
    if false {
        if true {
            v5 = v5
        }
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int = -766
            _ = fi6
            _ = v7
        }
        if (42 <= 2147483647) {
            var v8 bool = (219 < 255)
            var v9 uint64 = 512
            v5 = 1.5
            _ = v8
            _ = v9
        } else {
            fmt.Println(v5)
            v5 = float32(1e10)
        }
    } else {
        var v10 int16 = 0
        _ = v10
    }
    v11 := GenS1{F0: -1, F1: 65535}
    _ = v11
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 uint) int32 {
    v12 := GenS1{F0: 1, F1: 1}
    if (!true) {
        p0 = (p0 + p0)
        switch -1 {
        case 9223372036854775807:
            if false {
                switch "foo bar" {
                case "a":
                    p1 = 99
                case "":
                    fmt.Println(p0, p0)
                case "hello":
                    p2 = p2
                default:
                    v13 := []int{-1, 32767, 32767, 731}
                    _ = v13
                }
                v14 := GenS1{F0: 0, F1: 2147483647}
                fmt.Println(p2, p0)
                _ = v14
            } else {
                v15 := GenS1{F0: 1, F1: 42}
                v16 := GenS2{F0: p1, F1: ("foo bar" + ("foo bar" + "\n\t")), F2: 1}
                fmt.Println(p0, p2, p1)
                _ = v15
                _ = v16
            }
        case 2147483647:
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 float32 = (3.14159 * 0.1)
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 string = "a"
                    p0 -= p0
                    _ = fi19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
        case -9223372036854775808:
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS2{F0: (p1 / (((p1 >> ((18446744073709551615) & 63))) | 1)), F1: "a", F2: (p0 / (((p0 >> (((p2 / ((p2) | 1))) & 63))) | 1))}
                p0 += (p0 * p0)
                var v23 float32 = 1.5
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            p1 = 0
            if true {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 int16 = 32767
                    if (!false) {
                        v27 := []int{}
                        _ = v27
                    } else {
                        v28 := []int{fi24}
                        if ("a" <= "\n\t") {
                            v29 := GenS1{F0: (^(v26)), F1: 0}
                            _ = v29
                        }
                        _ = v28
                    }
                    var v30 int16 = (v26 / (((v26 / ((v26) | 1))) | 1))
                    _ = fi25
                    _ = v26
                    _ = v30
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 string = ("hello" + "hello")
                    v32 = ""
                    v33 := []int{9223372036854775807, 4294967295, -9223372036854775808}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                var v34 string = "foo bar"
                _ = v34
            }
            _ = fi24
        }
    } else {
        p2 -= 0
        var v35 float64 = 3.14159
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 uint8 = (p1 - 127)
        p2 = (p2 ^ 423)
        if (771 > uint32(p2)) {
            if (fi36 == 42) {
                var v38 int16 = 1000
                v39 := []int{(fi36 * (fi36 | (fi36 * -9223372036854775808)))}
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
                _ = v38
                _ = v39
            }
        }
        _ = fi36
        _ = v37
    }
    v40 := GenS3{F0: 1, F1: p0, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v40
    return -346
}

func main() {
    v1 := []string{"\n\t", "foo bar"}
    v2 := []uint32{42}
    var v3 uint64 = 277
    v2 = append(v2, 348)
    v1 = append(v1, "foo bar")
    v41 := GenS4{F0: 0, F1: 0}
    v42 := GenS5{F0: ("hello" + ((("" + "a") + (("a" + "foo bar") + ("foo bar" + ("hello" + "a")))) + "hello")), F1: -32768, F2: false}
    v3 = v3
    v43 := GenS6{F0: 42, F1: 127, F2: ("\n\t" + "\n\t")}
    v3 = v3
    v44 := GenS6{F0: 161, F1: -1, F2: "a"}
    switch 32768 {
    case 9223372036854775807:
        v3 = v3
    case 135:
        var v45 uint16 = 0
        _ = v45
    case 128:
        for fi46 := 0; fi46 < 5; fi46++ {
            switch fi46 {
            case 883:
                var v47 uint64 = v3
                _ = v47
            default:
                var v48 string = "foo bar"
                _ = v48
            }
            v3 = v3
            var v49 int16 = 751
            _ = fi46
            _ = v49
        }
    }
    v3 = v3
    if len(v1) > 0 {
        v1[0] = ("foo bar" + ((("foo bar" + "") + "\n\t") + ("foo bar" + "\n\t")))
    }
    switch "\n\t" {
    case "a":
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := GenS4{F0: -289, F1: 0}
            if (29 <= -128) {
                for fi52 := 0; fi52 < 9; fi52++ {
                    var v53 int64 = 9223372036854775807
                    var v54 int32 = 1
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "\n\t" {
                case "\n\t":
                    v3 = (v3 ^ v3)
                }
            }
            _ = fi50
            _ = v51
        }
    case "foo bar":
        switch -694 {
        case 42:
            v3 -= (v3 - (^(v3)))
        case 32767:
            fmt.Println(v3)
        case 592:
            v3 = (v3 ^ (v3 >> ((255) & 63)))
        default:
            if ((float64(-2.25) * ((3.14159 / -2.25) * 1e10)) == 1e10) {
                var v55 uint32 = 517
                _ = v55
            }
        }
    case "hello":
        v3 += 2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
}
