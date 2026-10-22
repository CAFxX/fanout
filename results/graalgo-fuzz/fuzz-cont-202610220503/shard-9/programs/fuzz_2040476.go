package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
func genFn0(p0 bool) int32 {
    v6 := GenS1{F0: 42, F1: 722}
    v7 := []int{-2147483648, 127, 127, -2147483648}
    var v8 int8 = 100
    var v9 bool = false
    for fi10 := 0; fi10 < 2; fi10++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi10
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1() uint16 {
    v11 := GenS1{F0: 2, F1: 4294967295}
    var v12 float64 = ((0.0 - 3.14159) + float64(((1e10 + (1.5 - 3.14159)) - 0.1)))
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 string = "hello"
            _ = fi13
            _ = v14
        }
    case "":
        v15 := GenS1{F0: 42, F1: 42}
        _ = v15
    default:
        fmt.Println(v12)
    }
    switch ("a" + (("\n\t" + ("foo bar" + "foo bar")) + "hello")) {
    case "":
        var v16 uint32 = 2147483648
        _ = v16
    case "a":
        v17 := []int{778, -2147483648, 42}
        _ = v17
    default:
        v18 := GenS1{F0: 29, F1: 4294967295}
        _ = v18
    }
    _ = v11
    _ = v12
    return 448
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 2
    var v3 uint16 = v2
    v4 := []int{32768, -422, -554}
    v5 := make(map[uint32]string)
    if (-1 >= -364) {
        fmt.Println(v2, v3)
        _ = len(v1)
    }
    var v19 float32 = 3.14159
    delete(v1, "\n\t")
    v19 -= -2.25
    if false {
        var v20 float64 = 1.5
        switch -115 {
        case 4294967295:
            v19 = v19
        case -286:
            var v21 uint = 611
            _ = v21
        default:
            v22 := GenS2{F0: "a", F1: 4294967295, F2: genFn1()}
            _ = v22
        }
        _ = v20
    } else {
        if (787 < v3) {
            var v23 uint = 4294967295
            var v24 float32 = v19
            fmt.Println(v24, v3, v3)
            _ = v23
            _ = v24
        }
    }
    if false {
        var v25 uint64 = 2
        var v26 string = "\n\t"
        _ = v25
        _ = v26
    }
    v2 = 1
    var v27 string = ("\n\t" + (("hello" + ("\n\t" + "a")) + (("hello" + "\n\t") + "\n\t")))
    if (42 > 2) {
        var v28 float32 = (v19 + v19)
        var v29 int64 = 458
        _ = v28
        _ = v29
    }
    var v30 uint = 710
    var v31 float32 = (v19 + ((v19 + (-2.25 + 0.0)) / (v19 * v19)))
    switch (((v27 + v27) + v27) + (v27 + "foo bar")) {
    case "foo bar":
        v32 := GenS2{F0: ("hello" + v27), F1: (v30 ^ uint(v3)), F2: 107}
        _ = v32
    case "a":
        switch 4294967295 {
        case -9223372036854775808:
            v27 = v27
        case 0:
            if (1000 == 42) {
                v33 := GenS1{F0: 128, F1: 268}
                switch -2147483648 {
                case -1:
                    fmt.Println(v30, v2)
                case 127:
                    v34 := GenS2{F0: v27, F1: v30, F2: genFn1()}
                    _ = v34
                }
                _ = v33
            } else {
                if (0 <= 398) {
                    fmt.Println(v30, v27, v27)
                    v27 = (v27 + v27)
                }
                var v35 uint32 = 4294967295
                switch -770 {
                case 4294967295:
                    switch -623 {
                    case 128:
                        fmt.Println(v3, v2)
                    case -1:
                        if true {
                            v19 = v19
                            v35 = v35
                            v5[(v35 >> ((0) & 63))] = "\n\t"
                        } else {
                            v27 = (("" + v27) + "a")
                            if false {
                                var v36 int64 = 1
                                v30 += v30
                                v2 += genFn1()
                                _ = v36
                            }
                            v35 -= 2
                        }
                    default:
                        if false {
                            fmt.Println(v27, v30)
                            var v37 uint = (v30 % (((^(v30))) | 1))
                            v4 = append(v4, -918)
                            _ = v37
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 32768
                            }
                            var v38 uint64 = 42
                            v35 -= v35
                            _ = v38
                        }
                    }
                case 127:
                    var v39 int16 = -640
                    _ = v39
                case 32767:
                    if (uint16(v35) > v3) {
                        var v40 uint32 = (v35 ^ (v35 / ((v35) | 1)))
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            for fi42 := 0; fi42 < 6; fi42++ {
                                v43 := GenS1{F0: 127, F1: 0}
                                var v44 uint = v30
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                            _ = fi41
                        }
                        v27 = "\n\t"
                    }
                }
                _ = v35
            }
        }
    }
    v45 := GenS1{F0: 119, F1: 2}
    var v46 float32 = v19
    fmt.Println(v19)
    v47 := v1[(v27 + v27)]
    var v48 int = -897
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
