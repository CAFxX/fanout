package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int16, p1 bool, p2 string) int32 {
    p1 = (!(!p1))
    p0 = p0
    p2 = p2
    var v7 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1(p0 uint8, p1 uint8, p2 uint) string {
    p0 = p0
    switch ("foo bar" + ("a" + "")) {
    case "\n\t":
        var v8 int64 = 42
        _ = v8
    case "":
        v9 := GenS1{F0: 1}
        _ = v9
    case "hello":
        p1 = (p0 % (((p0 ^ 255)) | 1))
    }
    if true {
        for fi10 := 0; fi10 < 5; fi10++ {
            if (0 > 4294967295) {
                fmt.Println(p0)
                fmt.Println(fi10, p2, fi10)
            } else {
                p2 = p2
                switch "hello" {
                case "\n\t":
                    switch ((("\n\t" + (("foo bar" + "\n\t") + "")) + "foo bar") + ("hello" + "hello")) {
                    case "":
                        v11 := []int{}
                        _ = v11
                    case "\n\t":
                        p0 -= (p0 << (((p2 & (p2 | (^(p2))))) & 63))
                    }
                case "":
                    if false {
                        p0 -= (p0 << (((p2 % ((485) | 1))) & 63))
                        v12 := []int{(fi10 * fi10), (fi10 + fi10), fi10}
                        switch (fi10 & 934) {
                        case 1:
                            if (true || true) {
                                p0 += p1
                                fmt.Println(p1, fi10, fi10)
                                p0 = p0
                            }
                        case -9223372036854775808:
                            p0 = 0
                        case 127:
                            fmt.Println(p1)
                        default:
                            if false {
                                var v13 bool = (p2 == (p2 | (p2 & (p2 ^ 2))))
                                v12 = append(v12, -1)
                                var v14 string = "foo bar"
                                _ = v13
                                _ = v14
                            } else {
                                fmt.Println(p1, fi10, fi10)
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        }
                        _ = v12
                    }
                default:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 bool = ((p1 << ((255) & 63)) > p0)
                        _ = fi15
                        _ = v16
                    }
                }
                v17 := []int{}
                _ = v17
            }
            switch "hello" {
            case "hello":
                fmt.Println(fi10)
            }
            switch (((("\n\t" + ("\n\t" + ("hello" + "a"))) + ("" + "a")) + "foo bar") + "\n\t") {
            case "hello":
                for fi18 := 0; fi18 < 8; fi18++ {
                    if (832 >= 42) {
                        var v19 uint16 = 65535
                        var v20 int64 = 275
                        _ = v19
                        _ = v20
                    } else {
                        switch ((("hello" + "hello") + "foo bar") + "a") {
                        case "":
                            fmt.Println(p0)
                        case "a":
                            v21 := GenS2{F0: -2147483648, F1: -9223372036854775808, F2: p2}
                            _ = v21
                        case "hello":
                            var v22 string = "\n\t"
                            _ = v22
                        }
                    }
                    v23 := []int{(fi10 + (fi18 ^ fi18)), fi10}
                    v23 = append(v23, 9223372036854775807)
                    _ = fi18
                    _ = v23
                }
            case "\n\t":
                v24 := []int{445, 127, 32767, fi10}
                _ = v24
            }
            _ = fi10
        }
    } else {
        if false {
            for fi25 := 0; fi25 < 6; fi25++ {
                p2 = 0
                for fi26 := 0; fi26 < 10; fi26++ {
                    v27 := []int{-879}
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            var v28 uint = 1
            var v29 float64 = (1.5 + -2.25)
            _ = v28
            _ = v29
        }
    }
    var v30 uint64 = 650
    if true {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS2{F0: 1, F1: int64(v30), F2: (p2 << ((36) & 63))}
            if false {
                p1 = p1
                v33 := GenS3{F0: (!("a" > "foo bar")), F1: 278}
                _ = v33
            }
            v34 := []int{(fi31 / ((2147483647) | 1)), (fi31 ^ (fi31 ^ int(p2))), fi31, 4294967295}
            _ = fi31
            _ = v32
            _ = v34
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return ""
}

func main() {
    var v1 int = -177
    var v2 uint16 = 287
    v3 := []int{(-(v1)), (v1 * v1)}
    var v4 uint64 = 739
    var v5 int16 = -1
    var v6 int32 = -354
    if (((v5 << ((4294967295) & 63)) < v5) && (int16(v2) <= v5)) {
        v4 = (v4 & (^(v4)))
        if (255 > 255) {
            var v35 uint32 = 2
            for fi36 := 0; fi36 < 7; fi36++ {
                v3 = append(v3, fi36)
                _ = fi36
            }
            _ = v35
        } else {
            var v37 uint32 = 834
            _ = v37
        }
        v6 = genFn0((v5 * -1), true, ("hello" + ("\n\t" + "foo bar")))
    } else {
        v38 := GenS4{F0: "", F1: 2}
        switch "a" {
        case "":
            v39 := GenS4{F0: "\n\t", F1: 2147483648}
            _ = v39
        }
        v40 := GenS4{F0: "\n\t", F1: 2147483648}
        _ = v38
        _ = v40
    }
    v2 -= 65535
    for fi41 := 0; fi41 < 5; fi41++ {
        fmt.Println(v4, v5)
        v42 := GenS5{F0: 134, F1: 1}
        _ = fi41
        _ = v42
    }
    if false {
        if len(v3) > 0 {
            v3[0] = (v1 - (v1 + v1))
        }
    } else {
        v43 := GenS4{F0: (("\n\t" + "hello") + ("foo bar" + "hello")), F1: 309}
        _ = v43
    }
    v44 := GenS6{F0: 578, F1: genFn0(-32768, (((1 >= (^(v4))) && ((v5 << ((930) & 63)) >= (v5 - int16(v1)))) || true), "a")}
    _ = len(v3)
    var v45 int = (v1 >> ((351) & 63))
    fmt.Println(v4)
    var v46 int16 = 287
    v47 := GenS7{F0: 0}
    if true {
        var v48 int64 = 1
        _ = v48
    }
    v2 -= (v2 ^ v2)
    switch v45 {
    case 4294967295:
        if false {
            fmt.Println(v5, v4)
            if (1 < 0) {
                switch v1 {
                case 4294967295:
                    var v49 int8 = -128
                    _ = v49
                case 0:
                    for fi50 := 0; fi50 < 14; fi50++ {
                        v45 = v1
                        v45 = 93
                        _ = fi50
                    }
                }
                v6 = genFn0(32767, true, ("a" + "foo bar"))
            }
            v3 = append(v3, 9223372036854775807)
        } else {
            if true {
                v4 = v4
                _ = len(v3)
            }
            fmt.Println(v45)
            var v51 uint32 = 65535
            _ = v51
        }
    case 32768:
        if ((v45 + (-(v1))) >= (v1 ^ v45)) {
            if true {
                for fi52 := 0; fi52 < 11; fi52++ {
                    var v53 string = (("" + "\n\t") + "foo bar")
                    _ = fi52
                    _ = v53
                }
            } else {
                if true {
                    for fi54 := 0; fi54 < 15; fi54++ {
                        v55 := GenS6{F0: 523, F1: -739}
                        _ = fi54
                        _ = v55
                    }
                    v1 += 0
                    v3 = append(v3, (v45 ^ (v45 | 4294967295)))
                }
                v1 = (v45 * v1)
            }
        }
    case 128:
        fmt.Println(v4, v2, v45)
    default:
        if false {
            if (true || false) {
                v5 = (v46 + v46)
                v6 -= v6
                fmt.Println(v2)
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = (^(v1))
                }
                switch 914 {
                case -393:
                    for fi56 := 0; fi56 < 12; fi56++ {
                        var v57 uint = 14
                        v5 = (v5 << ((0) & 63))
                        _ = fi56
                        _ = v57
                    }
                case 42:
                    for fi58 := 0; fi58 < 13; fi58++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v59 int = (fi58 & v45)
                        _ = fi58
                        _ = v59
                    }
                }
                v5 -= (v5 + 42)
            }
            if (2 > 2) {
                switch v45 {
                case -2147483648:
                    for fi60 := 0; fi60 < 4; fi60++ {
                        v61 := GenS6{F0: -9223372036854775808, F1: (v6 | -1)}
                        _ = fi60
                        _ = v61
                    }
                case 999:
                    fmt.Println(v2)
                case 42:
                    fmt.Println(v45, v45, v2)
                }
                fmt.Println(v6, v2, v5)
            }
        }
    }
    v1 = v45
    switch 4294967295 {
    case 1:
        switch (("foo bar" + "foo bar") + "") {
        case "a":
            var v62 bool = false
            _ = v62
        case "foo bar":
            var v63 int32 = v6
            _ = v63
        }
    case 867:
        v2 = (v2 << ((18446744073709551615) & 63))
    case 127:
        v64 := GenS6{F0: 0, F1: 505}
        _ = v64
    }
    v65 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v65
}
