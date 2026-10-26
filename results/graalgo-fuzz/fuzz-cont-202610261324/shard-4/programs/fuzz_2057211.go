package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -2147483648
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 16; fi4++ {
        if (true || (2 >= 4294967295)) {
            if true {
                v5 := GenS1{F0: -9223372036854775808, F1: 807}
                if false {
                    var v6 uint16 = 566
                    _ = v6
                }
                _ = v5
            } else {
                v7 := GenS1{F0: 42, F1: -32768}
                var v8 int8 = 100
                if (!true) {
                    v2 -= v2
                } else {
                    v9 := GenS1{F0: -1, F1: 32767}
                    _ = v9
                }
                _ = v7
                _ = v8
            }
            var v10 float64 = (v3 * (0.0 - 0.0))
            v11 := GenS1{F0: 1, F1: 42}
            _ = v10
            _ = v11
        }
        _ = fi4
    }
    var v12 int64 = 1
    v13 := GenS1{F0: (^(v12)), F1: 32767}
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: (v12 >> ((4294967295) & 63)), F1: 1}
        v2 += v2
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            fmt.Println(fi17)
            if ((0 == 98) || true) {
                fmt.Println(v2, v12, v3)
                var v18 int16 = 42
                _ = v18
            } else {
                fmt.Println(v12, v3, v3)
                if (42 >= 65535) {
                    switch ("\n\t" + "") {
                    case "foo bar":
                        if (159 == 139) {
                            delete(v1, ("hello" + (((("" + ("hello" + "hello")) + (("a" + ("" + "\n\t")) + "hello")) + "hello") + ("a" + (("a" + "\n\t") + "")))))
                        } else {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    if (255 != 127) {
                                        v19 := GenS2{F0: -726, F1: -37, F2: v2}
                                        fmt.Println(v12, v3, v12)
                                        var v20 uint = 0
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := v1[("\n\t" + "")]
                                        if ((!v21) || v21) {
                                            switch 32768 {
                                            case 2147483647:
                                                v22 := v1["\n\t"]
                                                _ = v22
                                            case 4294967295:
                                                fmt.Println(v3)
                                            }
                                            switch fi17 {
                                            case 1:
                                                if true {
                                                    v23 := GenS1{F0: v12, F1: 776}
                                                    _ = v23
                                                } else {
                                                    v24 := GenS3{F0: (("a" + "a") + "\n\t"), F1: fi16}
                                                    v12 -= (v12 ^ (v12 * (-(v12))))
                                                    _ = v24
                                                }
                                            default:
                                                var v25 int32 = (v2 & (v2 / ((v2) | 1)))
                                                _ = v25
                                            }
                                            var v26 int16 = 0
                                            _ = v26
                                        }
                                        _ = v21
                                    }
                                }
                            default:
                                v3 -= v3
                            }
                            v2 = v2
                        }
                    case "":
                        delete(v1, "hello")
                    }
                    var v27 float32 = 0.1
                    _ = v27
                }
                v12 = v12
                var v28 string = "hello"
                v2 += (v2 % ((v2) | 1))
                _ = v28
            }
            v12 = -1
            _ = fi17
        }
        _ = fi16
    }
    v29 := v1["\n\t"]
    var v30 bool = true
    if (-842 == 42) {
        v30 = v30
        v29 = ((42 != 42) && (65535 >= 42))
    } else {
        if v30 {
            if (!(!true)) {
                var v31 int = -623
                v32 := v1["\n\t"]
                _ = v31
                _ = v32
            } else {
                v33 := []int{760}
                if ((1 > 42) || false) {
                    v34 := GenS4{F0: 42, F1: 2, F2: 0}
                    _ = v34
                }
                _ = v33
            }
            v29 = (1e10 >= ((v3 * 1.5) + 1.5))
        } else {
            v12 = (v12 ^ (v12 * 1))
        }
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            delete(v1, ("\n\t" + "a"))
        } else {
            var v36 string = ("a" + "foo bar")
            _ = v36
        }
        var v37 uint8 = 127
        fmt.Println(fi35, v12)
        _ = fi35
        _ = v37
    }
    var v38 bool = (v30 && v30)
    v29 = (((0.1 - float32((0.1 + 0.1))) - (0.1 * 1.5)) >= ((0.0 - 1.5) + 0.1))
    var v39 bool = v38
    var v40 int64 = 954
    var v41 int16 = -526
    v30 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
