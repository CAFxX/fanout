package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 int8; F1 uint8; F2 int16 }
func genFn0(p0 uint32, p1 int64, p2 string) int8 {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = 195
        if (("hello" >= p2) || false) {
            v8 -= v8
            var v9 bool = (9223372036854775807 >= (fi7 / (((fi7 & fi7)) | 1)))
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            switch fi10 {
            case -162:
                var v11 int8 = -24
                _ = v11
            case 32767:
                p0 = p0
            case 109:
                fmt.Println(p1)
            }
            fmt.Println(p1, p0, p0)
            _ = fi10
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1() uint32 {
    var v12 uint16 = 964
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v12 += v12
            v15 := []int{fi14, 32767, fi13, fi13}
            _ = fi14
            _ = v15
        }
        v16 := []int{}
        for fi17 := 0; fi17 < 6; fi17++ {
            v12 = (v12 - 258)
            v12 = v12
            _ = fi17
        }
        _ = fi13
        _ = v16
    }
    v12 = (^(v12))
    v12 = 255
    _ = v12
    return 926
}

func main() {
    var v1 uint32 = 42
    var v2 int = 32768
    v3 := make(map[int]bool)
    var v4 uint64 = 42
    v5 := make(map[uint32]uint32)
    v6 := []string{"hello", "foo bar", ("\n\t" + ""), "a"}
    for fi18 := 0; fi18 < 5; fi18++ {
        v2 = v2
        fmt.Println(v2)
        if (true || false) {
            for fi19 := 0; fi19 < 7; fi19++ {
                v1 = v1
                _ = fi19
            }
            fmt.Println(v4)
            var v20 uint32 = uint32(v4)
            _ = v20
        } else {
            v2 = fi18
        }
        _ = fi18
    }
    switch v2 {
    case -2147483648:
        v21 := GenS1{F0: 255, F1: 1}
        _ = v21
    case -781:
        switch "a" {
        case "hello":
            var v22 uint64 = v4
            _ = v22
        }
    default:
        v1 += 530
    }
    if (1 > 781) {
        v3[v2] = (!false)
        fmt.Println(v2, v2, v2)
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int16 = 94
            _ = fi23
            _ = v24
        }
    }
    _ = len(v6)
    v1 = v1
    _ = len(v6)
    if (true || true) {
        if false {
            var v25 uint8 = 128
            _ = v25
        } else {
            for fi26 := 0; fi26 < 20; fi26++ {
                v2 = -208
                var v27 int16 = -487
                switch "foo bar" {
                case "":
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 uint8 = 1
                        _ = fi28
                        _ = v29
                    }
                case "foo bar":
                    v30 := GenS1{F0: 42, F1: 389}
                    _ = v30
                case "a":
                    if (3.14159 > 3.14159) {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v32 := GenS2{F0: 32767, F1: (^(v4)), F2: 52}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if true {
                                var v33 uint32 = (v1 | (v1 + v1))
                                v34 := GenS3{F0: 1}
                                _ = v33
                                _ = v34
                            } else {
                                if true {
                                    _ = len(v5)
                                    v27 = 1000
                                } else {
                                    v27 += v27
                                    v4 -= (v4 | (^(v4)))
                                    v35 := GenS4{F0: ""}
                                    _ = v35
                                }
                                v2 = v2
                                if len(v6) > 0 {
                                    v6[0] = ((("hello" + "\n\t") + ("\n\t" + "foo bar")) + ("" + ((("\n\t" + "hello") + (("\n\t" + "\n\t") + ("\n\t" + ("\n\t" + "")))) + "foo bar")))
                                }
                                v36 := GenS5{F0: (^(v4))}
                                _ = v36
                            }
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        var v37 uint32 = (v1 & genFn1())
                        if true {
                            switch "foo bar" {
                            case "a":
                                v37 -= (v37 % ((v1) | 1))
                            case "foo bar":
                                v38 := GenS5{F0: (v4 ^ (^(v4)))}
                                _ = v38
                            default:
                                var v39 int8 = -128
                                _ = v39
                            }
                            fmt.Println(v1, v37, v1)
                            v4 = 2
                        }
                        fmt.Println(fi26, v1, fi26)
                        _ = v37
                    }
                default:
                    _ = len(v5)
                }
                _ = fi26
                _ = v27
            }
            if true {
                v40 := GenS6{F0: 747}
                switch ("hello" + "\n\t") {
                case "":
                    var v41 uint16 = 57
                    _ = v41
                case "foo bar":
                    for fi42 := 0; fi42 < 1; fi42++ {
                        fmt.Println(v2, v4, fi42)
                        _ = fi42
                    }
                default:
                    fmt.Println(v2)
                }
                for fi43 := 0; fi43 < 2; fi43++ {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        if (1 != 1) {
                            var v45 uint = 65535
                            v46 := GenS7{F0: 1}
                            _ = v45
                            _ = v46
                        } else {
                            v1 = 1
                        }
                        var v47 string = ""
                        var v48 int = 29
                        _ = fi44
                        _ = v47
                        _ = v48
                    }
                    if (406 < 1) {
                        var v49 bool = (1.5 <= (0.1 * (float32((3.14159 / float64(0.1))) + 1e10)))
                        if (255 > 2) {
                            for fi50 := 0; fi50 < 17; fi50++ {
                                var v51 uint = 0
                                v52 := GenS8{F0: ((("\n\t" + (((("foo bar" + "") + "") + "a") + "")) + ((("" + "\n\t") + (("hello" + ("" + "\n\t")) + "a")) + ((("hello" + "hello") + ("foo bar" + "\n\t")) + ("foo bar" + ("foo bar" + "hello"))))) + ("" + ("hello" + "hello")))}
                                switch fi50 {
                                case 1:
                                    if v49 {
                                        var v53 string = "\n\t"
                                        var v54 bool = ((v49 && (true && false)) || (-128 != 42))
                                        _ = v53
                                        _ = v54
                                    } else {
                                        var v55 int = -2147483648
                                        switch (v55 - 0) {
                                        case 32767:
                                            var v56 float32 = float32((1e10 * -2.25))
                                            _ = v56
                                        case 0:
                                            if (0 == v4) {
                                                v2 = (fi43 + 933)
                                            } else {
                                                var v57 int16 = -32768
                                                fmt.Println(v4, v4)
                                                _ = v57
                                            }
                                        default:
                                            if (!(32767 >= 1000)) {
                                                fmt.Println(v2)
                                                if v49 {
                                                    var v58 int64 = 1
                                                    var v59 uint8 = 107
                                                    _ = v58
                                                    _ = v59
                                                }
                                            }
                                        }
                                        v49 = (!((v4 % (((v4 + (v4 | v4))) | 1)) <= 0))
                                        _ = v55
                                    }
                                    v1 = (v1 & v1)
                                case 9223372036854775807:
                                    v60 := GenS9{F0: "a", F1: 127}
                                    _ = v60
                                case 773:
                                    var v61 float32 = 3.14159
                                    switch "" {
                                    case "\n\t":
                                        v61 = ((0.1 + v61) + ((v61 / float32(1.5)) * (v61 + (-2.25 - v61))))
                                    case "hello":
                                        var v62 int = (fi43 << (((v51 + v51)) & 63))
                                        _ = v62
                                    }
                                    _ = v61
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v1 = 2147483647
                            v49 = true
                        }
                        v63 := GenS10{F0: 0, F1: 1}
                        _ = v49
                        _ = v63
                    }
                    fmt.Println(v1, v4, fi43)
                    _ = fi43
                }
                _ = v40
            } else {
                switch "" {
                case "hello":
                    v2 = (v2 * 127)
                }
                var v64 bool = (!false)
                v3[v2] = (128 > 128)
                _ = v64
            }
        }
        if len(v6) > 0 {
            v6[0] = "hello"
        }
    } else {
        switch (("foo bar" + "a") + "") {
        case "foo bar":
            v65 := GenS3{F0: 2}
            _ = v65
        case "a":
            v4 -= v4
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = "hello"
            }
        }
    }
    v2 = 32768
    v2 = 2147483647
    v1 -= genFn1()
    switch ((((((("hello" + "a") + "a") + ("hello" + "\n\t")) + "foo bar") + "") + "\n\t") + ("\n\t" + "a")) {
    case "hello":
        var v66 float32 = -2.25
        _ = v66
    case "a":
        var v67 uint16 = 642
        _ = v67
    }
    if (18446744073709551615 > 2) {
        v2 = v2
        var v68 float32 = 0.1
        _ = v68
    }
    for fi69 := 0; fi69 < 7; fi69++ {
        v2 += (v2 * (v2 + 127))
        if (("a" + "hello") <= "\n\t") {
            v1 += 2147483647
        }
        v70 := GenS11{F0: 127, F1: 42, F2: int16(v4)}
        _ = fi69
        _ = v70
    }
    var v71 int = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
