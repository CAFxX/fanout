package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 42
    v5 := GenS1{F0: 941, F1: 0}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: 581, F1: 868}
        fmt.Println(v4)
        var v8 uint16 = 800
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 65535, F1: 42}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint = 0
        _ = fi10
        _ = v11
    }
    v4 -= 0
    v12 := []int{42, 32767, 9223372036854775807}
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                v4 = v4
                switch "" {
                case "hello":
                    v15 := v3[0]
                    _ = v15
                case "":
                    fmt.Println(fi13, fi14)
                case "a":
                    v4 += (-(v4))
                default:
                    fmt.Println(fi14, fi14)
                }
                _ = fi14
            }
            var v16 bool = (true && false)
            switch "\n\t" {
            case "hello":
                v17 := GenS1{F0: 65535, F1: 0}
                _ = v17
            case "a":
                switch (^(fi13)) {
                case 530:
                    v4 -= 127
                case 127:
                    _ = len(v3)
                case 4294967295:
                    var v18 float32 = -2.25
                    _ = v18
                default:
                    switch fi13 {
                    case 127:
                        var v19 bool = ("foo bar" != (((("" + "foo bar") + "a") + "") + ("hello" + "a")))
                        _ = v19
                    }
                }
            case "\n\t":
                v16 = true
            }
            _ = fi13
            _ = v16
        }
        for fi20 := 0; fi20 < 1; fi20++ {
            v1[2] = ("" + ("\n\t" + "a"))
            if (!true) {
                delete(v1, 1)
                var v21 uint = 2
                if false {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 string = "foo bar"
                        switch (fi20 & (^(fi22))) {
                        case 0:
                            v23 = v23
                        }
                        _ = fi22
                        _ = v23
                    }
                    var v24 uint = (v21 >> ((v21) & 63))
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v24
                } else {
                    switch fi20 {
                    case 42:
                        if false {
                            if true {
                                v25 := GenS2{F0: 65535, F1: (^(fi20))}
                                _ = v25
                            } else {
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    if false {
                                        fmt.Println(v4, v4)
                                        if (((!true) && true) && (!true)) {
                                            switch -505 {
                                            case 569:
                                                v27 := GenS2{F0: 843, F1: (fi26 & (-(fi20)))}
                                                _ = v27
                                            case 4294967295:
                                                v28 := v3[1]
                                                _ = v28
                                            default:
                                                var v29 bool = true
                                                fmt.Println(v4, v21)
                                                _ = v29
                                            }
                                            v21 += v21
                                            fmt.Println(v21, fi20, fi20)
                                        }
                                    }
                                    v21 = (v21 & v21)
                                    _ = fi26
                                }
                                for fi30 := 0; fi30 < 4; fi30++ {
                                    v21 += (v21 | 0)
                                    v31 := GenS3{F0: 255}
                                    if true {
                                        switch ((("" + (("a" + "hello") + "hello")) + ("foo bar" + "hello")) + "hello") {
                                        case "\n\t":
                                            switch (^(fi20)) {
                                            case 622:
                                                var v32 int64 = -832
                                                _ = v32
                                            case 32767:
                                                v21 = v21
                                            }
                                        case "hello":
                                            v33 := GenS2{F0: 32767, F1: (-(fi20))}
                                            _ = v33
                                        case "":
                                            if false {
                                                switch ("\n\t" + "\n\t") {
                                                case "hello":
                                                    var v34 uint16 = 1
                                                    _ = v34
                                                case "a":
                                                    var v35 string = (("hello" + ("\n\t" + "a")) + "a")
                                                    _ = v35
                                                default:
                                                    var v36 uint64 = 513
                                                    _ = v36
                                                }
                                                if true {
                                                    v37 := GenS4{F0: (^(v4))}
                                                    v21 += v21
                                                    var v38 uint64 = 18446744073709551615
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v39 := v3[535]
                                                _ = v39
                                            }
                                        }
                                    } else {
                                        switch 0 {
                                        case 302:
                                            v21 += 2
                                        case 4294967295:
                                            v21 = (v21 % (((v21 % (((v21 & (v21 << ((386) & 63)))) | 1))) | 1))
                                        default:
                                            fmt.Println(fi20, v21)
                                        }
                                        var v40 uint64 = 434
                                        v4 = 1
                                        _ = v40
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                v41 := GenS5{F0: 1, F1: 309}
                                _ = v41
                            }
                            fmt.Println(fi20, fi20)
                        } else {
                            var v42 string = ("\n\t" + "a")
                            _ = v42
                        }
                    case 0:
                        v3[341] = 441
                    case -521:
                        delete(v2, (fi20 / ((-9223372036854775808) | 1)))
                    }
                }
                _ = v21
            }
            if true {
                switch "a" {
                case "\n\t":
                    if false {
                        if false {
                            var v43 float32 = (1.5 - 3.14159)
                            _ = v43
                        } else {
                            _ = len(v12)
                            if true {
                                for fi44 := 0; fi44 < 19; fi44++ {
                                    v45 := GenS2{F0: 65535, F1: fi44}
                                    _ = fi44
                                    _ = v45
                                }
                                if (!(false && (true && (true || (-44 > 1))))) {
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        if ((0.0 + (3.14159 * -2.25)) == (0.1 / -2.25)) {
                                            var v47 int16 = 42
                                            _ = v47
                                        }
                                        var v48 int64 = -9223372036854775808
                                        v48 -= (v48 % (((-(v48))) | 1))
                                        v48 = v48
                                        _ = fi46
                                        _ = v48
                                    }
                                }
                                v4 = -1
                            } else {
                                fmt.Println(v4, fi20)
                            }
                            _ = len(v12)
                        }
                        v4 -= 21
                        if (!true) {
                            v1[2] = "hello"
                            if (2 <= 308) {
                                if false {
                                    var v49 int8 = v4
                                    for fi50 := 0; fi50 < 8; fi50++ {
                                        v51 := GenS6{F0: ""}
                                        _ = len(v12)
                                        _ = fi50
                                        _ = v51
                                    }
                                    _ = v49
                                } else {
                                    if true {
                                        if false {
                                            var v52 string = "hello"
                                            _ = v52
                                        }
                                    }
                                }
                            } else {
                                var v53 int64 = -792
                                _ = v53
                            }
                            v54 := GenS5{F0: 2, F1: 1}
                            _ = v54
                        } else {
                            if (!false) {
                                fmt.Println(fi20, v4)
                                v3[1] = 42
                            }
                            v4 += -90
                            v55 := GenS2{F0: 1, F1: -517}
                            _ = v55
                        }
                    } else {
                        var v56 int16 = 42
                        var v57 uint8 = 2
                        var v58 uint8 = 2
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                case "":
                    var v59 string = "\n\t"
                    _ = v59
                }
            } else {
                var v60 int16 = 32767
                var v61 int8 = (v4 * v4)
                v1[660] = ("foo bar" + "foo bar")
                _ = v60
                _ = v61
            }
            _ = fi20
        }
        v4 += (v4 - v4)
    } else {
        v4 = v4
        v62 := GenS6{F0: "foo bar"}
        _ = v62
    }
    if false {
        v63 := v2[-608]
        _ = v63
    } else {
        var v64 uint32 = 4294967295
        _ = v64
    }
    fmt.Println(v4)
    var v65 float64 = 1e10
    var v66 float64 = (((float64(3.14159) / v65) - 1.5) / (((((((((3.14159 / v65) - float64(((float32(v65) / (((((0.1 - 0.0) + -2.25) * 0.1) + 1e10) * float32(1.5))) - 0.0))) - v65) * 1e10) / v65) / (0.0 + (float64((-2.25 - -2.25)) + v65))) / -2.25) - 3.14159) * v65))
    var v67 uint64 = 2
    v65 -= v66
    fmt.Println(v4, v4)
    v68 := GenS5{F0: (^(v67)), F1: 2147483647}
    v65 += (((v66 - v65) / v66) - 1.5)
    switch 2147483647 {
    case -529:
        var v69 uint16 = 65535
        _ = v69
    case 2147483647:
        v2[32768] = ((("" + "hello") + "\n\t") + "foo bar")
    default:
        fmt.Println(v67, v67)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v9
}
