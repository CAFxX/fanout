package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 int8 }
func genFn0() bool {
    v6 := GenS1{F0: false, F1: -1}
    switch "" {
    case "":
        v7 := GenS1{F0: true, F1: 127}
        _ = v7
    default:
        switch (("\n\t" + "") + "a") {
        case "":
            var v8 float32 = ((0.0 - float32(((0.0 / 1.0) * float64(0.0)))) - 0.0)
            _ = v8
        case "foo bar":
            switch 32767 {
            case 65:
                var v9 string = ((("foo bar" + "") + "") + "")
                _ = v9
            }
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v10 string = "a"
        _ = v10
    case "hello":
        v11 := []int{161, -670, -9223372036854775808}
        _ = v11
    case "":
        for fi12 := 0; fi12 < 17; fi12++ {
            if ((!(1 > 0)) || false) {
                v13 := []int{fi12, 370, 416, (fi12 % ((fi12) | 1))}
                _ = len(v13)
                _ = v13
            } else {
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 int64 = 0
                    _ = fi14
                    _ = v15
                }
                if (!true) {
                    switch "hello" {
                    case "hello":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v17 := []int{(fi16 ^ (^(fi16))), fi12}
                            if false {
                                if len(v17) > 0 {
                                    v17[0] = fi12
                                }
                            }
                            if (4294967295 <= 28) {
                                var v18 int = 42
                                var v19 bool = (v18 > fi16)
                                fmt.Println(fi12, fi12, fi16)
                                var v20 float64 = float64(3.14159)
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            var v21 int = (^(fi12))
                            fmt.Println(v21, v21)
                            _ = fi16
                            _ = v17
                            _ = v21
                        }
                    }
                }
            }
            var v22 int64 = -362
            for fi23 := 0; fi23 < 3; fi23++ {
                switch "" {
                case "hello":
                    if (false || true) {
                        v22 += -388
                        v24 := []int{(fi23 >> ((445) & 63)), fi12, 343, fi23}
                        var v25 float32 = 0.1
                        if ("hello" != "hello") {
                            var v26 int32 = -1
                            switch (("hello" + "\n\t") + "a") {
                            case "foo bar":
                                fmt.Println(v22)
                            case "\n\t":
                                fmt.Println(fi23)
                            }
                            if len(v24) >= 2 {
                                v24 = v24[1:len(v24)]
                            }
                            _ = v26
                        }
                        v25 += 3.14159
                        _ = v24
                        _ = v25
                    } else {
                        var v27 int32 = 0
                        v28 := GenS1{F0: ((!(((v22 >> ((324) & 63)) <= v22) && true)) || (-1 < -128)), F1: 40}
                        var v29 float64 = 1e10
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                default:
                    fmt.Println(fi12)
                }
                _ = fi23
            }
            _ = fi12
            _ = v22
        }
    }
    _ = v6
    return false
}
func genFn1() int32 {
    v30 := GenS1{F0: (true || false), F1: 0}
    if true {
        if (!(!true)) {
            v31 := GenS2{F0: 557}
            _ = v31
        }
    } else {
        var v32 int64 = -984
        switch (("" + ((("" + "a") + (("\n\t" + ("foo bar" + "foo bar")) + ("\n\t" + "foo bar"))) + (("hello" + "\n\t") + ("\n\t" + "hello")))) + "") {
        case "":
            var v33 bool = (860 >= 65535)
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 19; fi34++ {
                v32 = 0
                switch "foo bar" {
                case "foo bar":
                    v32 = -1
                case "a":
                    for fi35 := 0; fi35 < 1; fi35++ {
                        fmt.Println(v32, fi35)
                        v32 = (v32 | -857)
                        _ = fi35
                    }
                default:
                    v36 := []int{32768}
                    _ = v36
                }
                _ = fi34
            }
        case "\n\t":
            v37 := GenS2{F0: 42}
            _ = v37
        }
        _ = v32
    }
    _ = v30
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint = 42
    var v3 int8 = 100
    v4 := []string{"", "a"}
    v5 := []int{793, 590, -252}
    v4 = append(v4, "foo bar")
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 float32 = float32(-2.25)
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 uint16 = 32767
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi40
            _ = v41
        }
        v3 -= (v3 + (v3 >> (((v2 | 1)) & 63)))
        _ = fi38
        _ = v39
    }
    var v42 uint64 = 848
    v1 = append(v1, 131)
    v43 := GenS1{F0: ("hello" > "a"), F1: (v3 * (-(v3)))}
    fmt.Println(v42)
    v44 := GenS3{F0: (v42 >> (((v2 / ((31) | 1))) & 63)), F1: 42, F2: 1}
    if (float32(((1.5 + 0.1) * 0.0)) >= (0.1 / (float32((1.5 * (1e10 / 1e10))) * 0.1))) {
        var v45 bool = (127 >= 127)
        for fi46 := 0; fi46 < 6; fi46++ {
            switch fi46 {
            case 32768:
                for fi47 := 0; fi47 < 14; fi47++ {
                    var v48 int64 = -9223372036854775808
                    v45 = (!(!(1 > 568)))
                    if ((((float32(0.1) - float32(3.14159)) - (1.5 - float32(-2.25))) == (-2.25 + -2.25)) && v45) {
                        v48 -= (v48 - (v48 & (v48 & v48)))
                    } else {
                        var v49 int32 = -2147483648
                        var v50 float32 = (-2.25 / float32((3.14159 / (0.1 - (float64(0.0) * 0.1)))))
                        v50 = v50
                        _ = v49
                        _ = v50
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                switch "foo bar" {
                case "":
                    v2 += (^(v2))
                default:
                    var v51 float64 = 1.5
                    _ = v51
                }
            }
            for fi52 := 0; fi52 < 14; fi52++ {
                fmt.Println(fi46)
                v53 := GenS1{F0: v45, F1: (v3 << (((v2 / (((v2 & (v2 + 255))) | 1))) & 63))}
                _ = fi52
                _ = v53
            }
            for fi54 := 0; fi54 < 4; fi54++ {
                switch ("\n\t" + ("hello" + "a")) {
                case "":
                    var v55 uint64 = v42
                    _ = v55
                case "a":
                    switch "hello" {
                    case "hello":
                        v1 = append(v1, fi46)
                    case "a":
                        v42 = v42
                    }
                case "\n\t":
                    v2 += v2
                }
                v45 = (628 != 71)
                fmt.Println(fi54, v45)
                _ = fi54
            }
            _ = fi46
        }
        _ = v45
    } else {
        v56 := GenS3{F0: v42, F1: uint16(v42), F2: v3}
        _ = v56
    }
    for fi57 := 0; fi57 < 6; fi57++ {
        v58 := GenS4{F0: false}
        v59 := GenS5{F0: 1}
        _ = fi57
        _ = v58
        _ = v59
    }
    var v60 uint = 757
    v61 := GenS6{F0: 2147483647}
    if ((v60 * v60) == 914) {
        fmt.Println(v3)
    } else {
        var v62 int64 = 0
        switch "hello" {
        case "hello":
            switch 161 {
            case 0:
                var v63 int16 = 32767
                _ = v63
            case 1:
                for fi64 := 0; fi64 < 13; fi64++ {
                    v60 += 255
                    if false {
                        v65 := GenS7{F0: -32768, F1: (v3 & v3)}
                        for fi66 := 0; fi66 < 5; fi66++ {
                            var v67 bool = (0.0 < (0.1 * 1.5))
                            _ = fi66
                            _ = v67
                        }
                        var v68 uint16 = 32767
                        _ = v65
                        _ = v68
                    }
                    _ = fi64
                }
            default:
                if true {
                    fmt.Println(v42)
                    var v69 int64 = -1
                    _ = v69
                }
            }
        }
        _ = v62
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v60
    _ = v61
}
