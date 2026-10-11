package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint32; F2 string }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int32 }

func main() {
    v1 := []int{32767, -768}
    v2 := make(map[string]int)
    var v3 uint16 = 2
    var v4 int8 = -96
    var v5 float64 = 0.1
    if (!(("\n\t" + "foo bar") != ("a" + "hello"))) {
        if (-1 < 0) {
            switch "hello" {
            case "hello":
                v3 += 42
            }
            _ = len(v2)
        } else {
            var v6 int = -541
            var v7 uint = 4294967295
            switch "\n\t" {
            case "hello":
                v3 = v3
            case "\n\t":
                fmt.Println(v4, v3)
            }
            _ = v6
            _ = v7
        }
        var v8 float64 = float64(1e10)
        v9 := GenS1{F0: (v3 << ((4294967295) & 63))}
        _ = v8
        _ = v9
    }
    v4 -= (v4 * 0)
    fmt.Println(v4, v5)
    v10 := GenS1{F0: v3}
    var v11 uint16 = 255
    if true {
        if (!false) {
            var v12 string = ("\n\t" + ("hello" + "foo bar"))
            _ = v12
        }
    }
    v13 := GenS2{F0: 977, F1: 280}
    var v14 uint8 = 1
    if (!true) {
        if (((-(v4)) >= (^(v4))) && true) {
            var v15 float64 = v5
            v16 := GenS1{F0: 42}
            var v17 int64 = 0
            _ = v15
            _ = v16
            _ = v17
        }
    }
    v4 = v4
    if ((!true) && true) {
        switch -619 {
        case -916:
            if false {
                for fi18 := 0; fi18 < 5; fi18++ {
                    switch (("\n\t" + ("hello" + (("a" + "\n\t") + "foo bar"))) + "foo bar") {
                    case "":
                        var v19 uint = 2
                        _ = v19
                    case "\n\t":
                        for fi20 := 0; fi20 < 8; fi20++ {
                            fmt.Println(v5, fi18)
                            if (!(true && true)) {
                                v21 := GenS2{F0: 298, F1: 2}
                                var v22 float32 = 1e10
                                if true {
                                    fmt.Println(v5, v4, fi20)
                                }
                                _ = v21
                                _ = v22
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (false && true) {
                                    var v23 float32 = (-2.25 - (3.14159 / (float32(v5) * 3.14159)))
                                    _ = v23
                                } else {
                                    var v24 uint16 = 1
                                    var v25 uint32 = 2147483648
                                    var v26 uint32 = (v25 | uint32(v24))
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(fi20, v3, v3)
                                switch (("a" + (("hello" + "a") + (("a" + "foo bar") + (("foo bar" + "foo bar") + "")))) + (("hello" + "a") + "a")) {
                                case "\n\t":
                                    switch 9223372036854775807 {
                                    case 760:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case 32768:
                                        fmt.Println(v14, fi18, fi20)
                                    case 156:
                                        fmt.Println(v11, v14, v3)
                                    }
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi18))
                                    }
                                }
                                v3 -= 42
                                if true {
                                    v27 := GenS3{F0: 168}
                                    var v28 int64 = -9223372036854775808
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 uint64 = 18446744073709551615
                                    v30 := GenS1{F0: v11}
                                    switch (("" + ("hello" + ("" + ("hello" + ("foo bar" + "hello"))))) + "a") {
                                    case "\n\t":
                                        v29 -= v29
                                    }
                                    _ = v29
                                    _ = v30
                                }
                                var v31 uint16 = (v3 >> ((255) & 63))
                                _ = v31
                            }
                            _ = fi20
                        }
                    case "foo bar":
                        delete(v2, ("a" + ""))
                    default:
                        v32 := GenS3{F0: 1}
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v34 := GenS1{F0: v11}
                        v35 := GenS2{F0: 184, F1: 421}
                        var v36 uint16 = (^(v11))
                        v4 -= -1
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    if ((v11 / ((v3) | 1)) >= 32767) {
                        for fi37 := 0; fi37 < 11; fi37++ {
                            fmt.Println(v14)
                            var v38 int64 = 1
                            var v39 float64 = v5
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    } else {
                        v4 = (v4 % ((v4) | 1))
                    }
                    _ = fi18
                }
            }
        case 127:
            for fi40 := 0; fi40 < 2; fi40++ {
                for fi41 := 0; fi41 < 6; fi41++ {
                    if (!true) {
                        v42 := GenS4{F0: ("\n\t" + ""), F1: 65535, F2: ("" + ("foo bar" + ("" + "a")))}
                        var v43 uint8 = v14
                        if false {
                            var v44 uint = 4294967295
                            v5 -= ((v5 / float64((3.14159 + 1.5))) - (v5 + (-2.25 / (v5 - (1.5 + 1e10)))))
                            fmt.Println(v11, v14)
                            _ = v44
                        }
                        _ = v42
                        _ = v43
                    }
                    var v45 float32 = -2.25
                    _ = fi41
                    _ = v45
                }
                for fi46 := 0; fi46 < 6; fi46++ {
                    v3 += 65535
                    if ((42 <= 0) && (!(42 >= -186))) {
                        if (527 == -312) {
                            if true {
                                v14 += 127
                                switch (fi46 - fi40) {
                                case 667:
                                    v47 := GenS4{F0: "", F1: 141, F2: "\n\t"}
                                    _ = v47
                                }
                            }
                            switch ("a" + ("\n\t" + (("\n\t" + ("foo bar" + ("" + "a"))) + ""))) {
                            case "":
                                v48 := GenS1{F0: 1}
                                _ = v48
                            default:
                                v49 := GenS5{F0: 2147483647, F1: fi40}
                                _ = v49
                            }
                        } else {
                            if (!false) {
                                _ = len(v2)
                            }
                        }
                        v50 := GenS6{F0: 2147483647, F1: uint64(v3)}
                        v51 := GenS7{F0: 0}
                        _ = v50
                        _ = v51
                    } else {
                        v11 += 42
                        v52 := GenS1{F0: (v3 >> ((18446744073709551615) & 63))}
                        if len(v1) > 0 {
                            v1[0] = fi40
                        }
                        if (42 != 2147483647) {
                            v14 = (^(v14))
                            var v53 uint8 = (v14 - (^(v14)))
                            _ = v53
                        }
                        _ = v52
                    }
                    fmt.Println(v5, v3, fi46)
                    _ = fi46
                }
                if ((fi40 | 1) != (-(fi40))) {
                    if ((-756 == -2147483648) || true) {
                        for fi54 := 0; fi54 < 10; fi54++ {
                            delete(v2, "\n\t")
                            var v55 string = ("" + "hello")
                            if false {
                                _ = len(v2)
                                switch 610 {
                                case 1:
                                    if (v11 >= (^(v3))) {
                                        fmt.Println(v55, fi40)
                                        v56 := GenS6{F0: 384, F1: 698}
                                        _ = v56
                                    }
                                case 32768:
                                    _ = len(v1)
                                case 4294967295:
                                    var v57 uint = 2
                                    _ = v57
                                default:
                                    if (!false) {
                                        v4 -= (v4 % ((-1) | 1))
                                    }
                                }
                                v3 += (v3 << ((18446744073709551615) & 63))
                            } else {
                                switch fi40 {
                                case 128:
                                    var v58 string = v55
                                    _ = v58
                                case -805:
                                    var v59 uint32 = 1
                                    _ = v59
                                case -816:
                                    _ = len(v2)
                                default:
                                    if true {
                                        var v60 int16 = 1
                                        v61 := GenS4{F0: v55, F1: 65535, F2: ""}
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v5 = v5
                                        _ = len(v1)
                                        var v62 string = (v55 + "\n\t")
                                        _ = v62
                                    }
                                }
                                delete(v2, v55)
                                _ = len(v2)
                            }
                            _ = fi54
                            _ = v55
                        }
                        switch "" {
                        case "foo bar":
                            var v63 uint16 = v11
                            _ = v63
                        default:
                            v3 = (^(v3))
                        }
                    }
                    delete(v2, ((("a" + (("foo bar" + "hello") + (("foo bar" + "") + ("hello" + "")))) + "hello") + ""))
                    fmt.Println(v4, fi40, v4)
                }
                _ = fi40
            }
        case -2147483648:
            v11 -= 65535
        default:
            v3 = (v3 - (^(v11)))
        }
        v3 -= (v3 % ((v11) | 1))
    }
    fmt.Println(v11, v5, v14)
    fmt.Println(v11, v4)
    var v64 int = 0
    if (((v4 >> ((255) & 63)) > -45) && false) {
        if false {
            fmt.Println(v14, v5)
            _ = len(v2)
        } else {
            var v65 int64 = 1
            var v66 uint8 = 20
            v3 = 2
            _ = v65
            _ = v66
        }
        fmt.Println(v64, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
}
