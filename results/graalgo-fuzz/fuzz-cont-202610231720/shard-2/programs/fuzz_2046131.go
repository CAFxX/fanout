package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 string }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int16 }
func genFn0() uint8 {
    v4 := []int{2147483647}
    var v5 string = "\n\t"
    var v6 int16 = 32767
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 255
    var v3 uint8 = 174
    if (404 != -754) {
        var v7 uint16 = 32767
        _ = v7
    }
    v8 := []int{}
    if false {
        if (199 == 2147483647) {
            if true {
                if (!false) {
                    v3 = v3
                }
                var v9 uint16 = 2
                switch 42 {
                case -773:
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 bool = ((fi10 & 4294967295) <= 4294967295)
                        _ = fi10
                        _ = v11
                    }
                case 32767:
                    if true {
                        v2 = 691
                        var v12 int = 32768
                        v13 := GenS1{F0: (^(v3)), F1: 224, F2: 32767}
                        _ = v12
                        _ = v13
                    }
                default:
                    v14 := GenS2{F0: 1}
                    _ = v14
                }
                _ = v9
            }
            if (false && (!((true || ((v3 + v3) < 1)) && true))) {
                v1[0] = false
            } else {
                if false {
                    if (!((("" + ("a" + "\n\t")) + ("hello" + "\n\t")) <= "\n\t")) {
                        var v15 string = ("\n\t" + "foo bar")
                        var v16 float32 = 1e10
                        _ = v15
                        _ = v16
                    } else {
                        v17 := GenS3{F0: 792}
                        var v18 string = "a"
                        _ = v17
                        _ = v18
                    }
                    switch ("hello" + ("hello" + ("" + ((("a" + ("foo bar" + ("a" + "\n\t"))) + "hello") + "a")))) {
                    case "foo bar":
                        for fi19 := 0; fi19 < 16; fi19++ {
                            for fi20 := 0; fi20 < 4; fi20++ {
                                switch ("\n\t" + (("" + "foo bar") + (("foo bar" + ("\n\t" + ((("hello" + "\n\t") + ("hello" + "hello")) + "\n\t"))) + ("a" + ("\n\t" + "foo bar"))))) {
                                case "a":
                                    _ = len(v8)
                                case "\n\t":
                                    v21 := GenS3{F0: 959}
                                    _ = v21
                                default:
                                    v22 := GenS2{F0: 18446744073709551615}
                                    _ = v22
                                }
                                _ = fi20
                            }
                            v23 := GenS4{F0: v2, F1: 32767, F2: (("foo bar" + "") + "")}
                            _ = fi19
                            _ = v23
                        }
                    case "hello":
                        if true {
                            if true {
                                var v24 float64 = -2.25
                                v2 = 18446744073709551615
                                _ = v24
                            }
                            v3 = 127
                        }
                    case "":
                        v25 := GenS5{F0: 2, F1: -992}
                        _ = v25
                    default:
                        switch ("\n\t" + "") {
                        case "a":
                            var v26 uint64 = 287
                            _ = v26
                        default:
                            if false {
                                _ = len(v1)
                            }
                        }
                    }
                    _ = len(v8)
                } else {
                    var v27 int64 = -1
                    var v28 int16 = -785
                    _ = v27
                    _ = v28
                }
                var v29 int16 = 757
                _ = v29
            }
            fmt.Println(v2, v2)
        } else {
            if ((v2 == 255) || true) {
                var v30 uint64 = 904
                _ = v30
            }
            v3 = (v3 >> ((4294967295) & 63))
        }
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(fi31)
            _ = fi31
        }
    } else {
        v3 = 2
        for fi32 := 0; fi32 < 18; fi32++ {
            v33 := GenS4{F0: v2, F1: 170, F2: "hello"}
            for fi34 := 0; fi34 < 2; fi34++ {
                v2 = v2
                var v35 int16 = 32767
                v36 := GenS1{F0: v3, F1: 572, F2: -32768}
                var v37 string = "\n\t"
                v35 -= (v35 * v35)
                switch (v37 + "foo bar") {
                case "hello":
                    var v38 int16 = 32767
                    _ = v38
                default:
                    if len(v8) > 0 {
                        v8[0] = (fi34 + fi34)
                    }
                }
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
            _ = fi32
            _ = v33
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            var v40 bool = (-2.25 == (float64((((0.0 + 3.14159) * 1.5) + 1.5)) - (1.5 * 1.5)))
            switch ("a" + "hello") {
            case "hello":
                if v40 {
                    v2 = v2
                }
            case "foo bar":
                var v41 uint16 = 42
                _ = v41
            default:
                if len(v8) > 0 {
                    v8[0] = fi39
                }
            }
            _ = fi39
            _ = v40
        }
    }
    _ = len(v8)
    _ = len(v1)
    var v42 int = 42
    if true {
        v3 -= v3
    } else {
        for fi43 := 0; fi43 < 7; fi43++ {
            for fi44 := 0; fi44 < 4; fi44++ {
                v3 = (v3 ^ genFn0())
                v3 += (^(v3))
                _ = fi44
            }
            if (false || true) {
                v8 = append(v8, v42)
                _ = len(v8)
            } else {
                v42 += (fi43 / ((-713) | 1))
            }
            _ = fi43
        }
        var v45 int32 = int32(v42)
        for fi46 := 0; fi46 < 10; fi46++ {
            fmt.Println(fi46, v3, v3)
            v47 := GenS6{F0: 1}
            switch (v42 % (((v42 / (((v42 | (fi46 ^ int(v2)))) | 1))) | 1)) {
            case 93:
                _ = len(v1)
            case 42:
                if false {
                    v48 := v1[-2147483648]
                    _ = v48
                } else {
                    var v49 int32 = (v45 * v45)
                    for fi50 := 0; fi50 < 12; fi50++ {
                        v49 += v45
                        _ = fi50
                    }
                    _ = v49
                }
            default:
                fmt.Println(fi46, v45, v3)
            }
            _ = fi46
            _ = v47
        }
        _ = v45
    }
    if len(v8) > 0 {
        v8[0] = v42
    }
    _ = len(v1)
    var v51 float32 = 0.1
    for fi52 := 0; fi52 < 4; fi52++ {
        var v53 int32 = 121
        _ = fi52
        _ = v53
    }
    var v54 uint = v2
    if len(v8) > 0 {
        v8[0] = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v51
    _ = v54
    _ = v8
}
