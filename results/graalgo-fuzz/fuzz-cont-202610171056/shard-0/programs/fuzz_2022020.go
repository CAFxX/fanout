package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint64 }
type GenS8 struct { F0 bool }
func genFn0(p0 int, p1 int32, p2 uint16) int {
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 719, F1: p2}
        var v6 int = p0
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 string = ("\n\t" + "foo bar")
        for fi9 := 0; fi9 < 17; fi9++ {
            switch (fi9 << ((4294967295) & 63)) {
            case -9223372036854775808:
                v10 := GenS1{F0: 1000, F1: 97}
                var v11 uint = 255
                _ = v10
                _ = v11
            case 127:
                if false {
                    var v12 uint32 = 4294967295
                    _ = v12
                }
                v8 = "\n\t"
                if true {
                    v13 := []int{fi9, (fi7 + (^(p0))), fi9}
                    fmt.Println(p0)
                    _ = v13
                }
                p1 += -902
                switch fi9 {
                case 9223372036854775807:
                    v14 := []int{p0}
                    p1 -= (^(p1))
                    if len(v14) > 0 {
                        v14[0] = (^(fi7))
                    }
                    _ = v14
                }
            case 42:
                p2 = p2
            }
            v15 := []int{(-(p0)), (fi9 % ((fi9) | 1))}
            _ = fi9
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    var v16 float32 = ((1e10 / 1.0) - float32(1.5))
    v17 := GenS2{F0: false, F1: p2}
    var v18 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return 32767
}

func main() {
    v1 := []string{("" + ("foo bar" + ("hello" + ("hello" + ("" + (("hello" + "\n\t") + (((("a" + (("hello" + ("" + "a")) + "a")) + "hello") + ("" + "foo bar")) + "foo bar"))))))), ""}
    var v2 int8 = 100
    v3 := make(map[string]bool)
    switch ((("foo bar" + "") + "a") + ("" + ("\n\t" + "\n\t"))) {
    case "a":
        fmt.Println(v2)
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint8 = 0
        if false {
            var v21 float64 = ((1e10 / (float64(float32(0.1)) - -2.25)) + 1e10)
            var v22 int64 = 229
            _ = v21
            _ = v22
        } else {
            if true {
                fmt.Println(fi19, v20)
                for fi23 := 0; fi23 < 15; fi23++ {
                    if true {
                        if false {
                            _ = len(v3)
                        }
                        switch (("hello" + ("" + "\n\t")) + (("" + ((((("foo bar" + "a") + ("" + ("\n\t" + "hello"))) + "a") + ("hello" + "\n\t")) + ((("foo bar" + "\n\t") + "hello") + ""))) + "foo bar")) {
                        case "\n\t":
                            if (v2 > 0) {
                                if (-2147483648 == fi23) {
                                    switch (fi19 * (fi19 ^ fi19)) {
                                    case 128:
                                        v2 = v2
                                    default:
                                        v24 := GenS2{F0: (9223372036854775807 == 965), F1: 255}
                                        _ = v24
                                    }
                                    switch ("a" + "hello") {
                                    case "a":
                                        v20 = (v20 << ((2) & 63))
                                    case "foo bar":
                                        if (4294967295 <= 530) {
                                            v2 = v2
                                            v25 := GenS3{F0: (0.0 != 1.5), F1: 0}
                                            _ = v25
                                        } else {
                                            fmt.Println(fi23, fi23)
                                        }
                                    }
                                    v26 := GenS2{F0: false, F1: 0}
                                    _ = v26
                                }
                            } else {
                                v27 := GenS4{F0: 1}
                                var v28 uint64 = 149
                                _ = v27
                                _ = v28
                            }
                        case "a":
                            v29 := GenS5{F0: false}
                            _ = v29
                        }
                        var v30 uint8 = (^(v20))
                        if ((v2 << ((4294967295) & 63)) < (v2 | 100)) {
                            v31 := GenS2{F0: (1.5 != (3.14159 / float64((1e10 / 1.0)))), F1: 139}
                            _ = v31
                        }
                        var v32 uint64 = 18446744073709551615
                        _ = v30
                        _ = v32
                    } else {
                        delete(v3, ("foo bar" + "a"))
                        v33 := GenS4{F0: 0}
                        var v34 int8 = 1
                        _ = v33
                        _ = v34
                    }
                    _ = fi23
                }
                var v35 int32 = 1
                _ = v35
            } else {
                v20 = v20
            }
        }
        _ = fi19
        _ = v20
    }
    if false {
        var v36 uint16 = 968
        _ = v36
    } else {
        if true {
            v2 += (v2 % ((v2) | 1))
        } else {
            if true {
                v37 := GenS1{F0: -578, F1: 255}
                _ = v37
            } else {
                var v38 float32 = 1.5
                _ = v38
            }
        }
        var v39 int32 = 0
        _ = v39
    }
    var v40 int8 = v2
    switch "a" {
    case "":
        var v41 bool = true
        _ = v41
    case "\n\t":
        var v42 uint32 = 4294967295
        _ = v42
    case "foo bar":
        var v43 int = -870
        _ = v43
    }
    if (!(1 != -760)) {
        v3["hello"] = ((1.5 <= 1.5) && false)
        var v44 int16 = -1
        if false {
            switch 32768 {
            case 488:
                var v45 uint = 4294967295
                _ = v45
            case -119:
                v40 = v2
            case -530:
                v46 := GenS6{F0: (-(v2))}
                _ = v46
            }
        } else {
            for fi47 := 0; fi47 < 3; fi47++ {
                if false {
                    switch 128 {
                    case 147:
                        var v48 string = (("foo bar" + "a") + ("foo bar" + "\n\t"))
                        _ = v48
                    case -2147483648:
                        v2 = v2
                    case 42:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v40 = 68
                    }
                    v2 = 1
                    var v49 float64 = float64((1e10 - 1e10))
                    _ = v49
                }
                var v50 uint8 = 1
                switch int(v50) {
                case 32768:
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                case 128:
                    switch "foo bar" {
                    case "a":
                        v2 = (v2 | -128)
                    case "hello":
                        for fi51 := 0; fi51 < 8; fi51++ {
                            fmt.Println(fi51, fi51, v44)
                            _ = fi51
                        }
                    }
                }
                _ = fi47
                _ = v50
            }
            var v52 float32 = ((-2.25 * -2.25) * 0.0)
            var v53 uint64 = 382
            _ = v52
            _ = v53
        }
        _ = v44
    } else {
        v54 := GenS1{F0: 1000, F1: 32767}
        var v55 int8 = (v2 >> ((2) & 63))
        for fi56 := 0; fi56 < 4; fi56++ {
            var v57 uint32 = 273
            v57 -= (v57 | 405)
            if true {
                v40 = v40
                for fi58 := 0; fi58 < 20; fi58++ {
                    v55 = v40
                    v2 -= 1
                    _ = fi58
                }
                v59 := GenS3{F0: ((!true) || false), F1: 48}
                _ = v59
            }
            _ = fi56
            _ = v57
        }
        _ = v54
        _ = v55
    }
    if true {
        v60 := GenS5{F0: false}
        var v61 int = -890
        _ = v60
        _ = v61
    } else {
        var v62 string = "hello"
        _ = v62
    }
    switch 42 {
    case 127:
        v63 := GenS7{F0: 561, F1: 468}
        _ = v63
    case 170:
        v64 := GenS8{F0: false}
        _ = v64
    case -9223372036854775808:
        v65 := GenS4{F0: 0}
        _ = v65
    }
    if (774 < 616) {
        var v66 string = "a"
        var v67 uint = 4294967295
        v68 := v3["\n\t"]
        _ = v66
        _ = v67
        _ = v68
    } else {
        _ = len(v3)
        var v69 int64 = 1
        fmt.Println(v69, v2, v2)
        _ = v69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
}
