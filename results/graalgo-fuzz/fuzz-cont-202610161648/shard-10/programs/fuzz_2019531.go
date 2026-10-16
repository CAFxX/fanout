package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 string }
func genFn0(p0 uint64, p1 int64) int {
    v4 := GenS1{F0: p0, F1: 9223372036854775807}
    for fi5 := 0; fi5 < 1; fi5++ {
        if (2 <= 0) {
            v6 := []int{128}
            _ = v6
        }
        p0 = p0
        _ = fi5
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = -2
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return 264
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    var v3 int8 = 100
    for fi9 := 0; fi9 < 13; fi9++ {
        _ = len(v1)
        _ = fi9
    }
    if (!false) {
        v2 += v2
    }
    fmt.Println(v3, v3)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch "hello" {
        case "foo bar":
            fmt.Println(v3)
        default:
            v2 -= v2
        }
        _ = fi10
    }
    var v11 uint = 65535
    var v12 uint16 = 495
    if (-656 != 2147483647) {
        var v13 float64 = -2.25
        fmt.Println(v2)
        var v14 uint = v11
        _ = v13
        _ = v14
    }
    switch "a" {
    case "hello":
        v11 = (v11 ^ v11)
    case "\n\t":
        if false {
            v15 := GenS1{F0: 1, F1: 605}
            if true {
                var v16 int32 = 0
                _ = v16
            }
            _ = v15
        } else {
            fmt.Println(v12, v12)
            if (v12 < v12) {
                switch -9223372036854775808 {
                case 0:
                    var v17 float64 = 3.14159
                    _ = v17
                case 537:
                    var v18 string = ("\n\t" + (("" + ((("\n\t" + "a") + ("" + "foo bar")) + ("a" + "a"))) + (("" + ((((("\n\t" + "hello") + "") + "a") + "\n\t") + "\n\t")) + "\n\t")))
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if false {
                            v1[65535] = (1 >= 547)
                            switch 665 {
                            case -923:
                                var v20 float32 = 0.1
                                _ = v20
                            default:
                                if false {
                                    for fi21 := 0; fi21 < 18; fi21++ {
                                        v22 := []int{128}
                                        fmt.Println(v11, fi21)
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                            }
                            var v23 int = (fi19 - 9223372036854775807)
                            _ = v23
                        }
                        v11 += v11
                        _ = fi19
                    }
                }
                v12 += (v12 + v12)
            }
            fmt.Println(v12, v3)
        }
    case "":
        switch "foo bar" {
        case "":
            v2 = 0.1
        case "a":
            v24 := GenS2{F0: v11}
            _ = v24
        case "\n\t":
            for fi25 := 0; fi25 < 20; fi25++ {
                v11 += v11
                var v26 int64 = 42
                _ = fi25
                _ = v26
            }
        }
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint32 = 1
            _ = fi27
            _ = v28
        }
    }
    var v29 uint16 = (v12 - v12)
    switch "hello" {
    case "":
        if (42 < 255) {
            v30 := v1[1]
            var v31 float64 = (-2.25 - 0.0)
            _ = v30
            _ = v31
        }
    case "\n\t":
        for fi32 := 0; fi32 < 20; fi32++ {
            if false {
                if (!(v12 <= v12)) {
                    if ((false && ("" != "")) && (!(int64(fi32) == 145))) {
                        if false {
                            v33 := GenS1{F0: 405, F1: (fi32 + genFn0(1, 42))}
                            v34 := GenS3{F0: 0, F1: 140, F2: genFn0(114, -808)}
                            _ = v33
                            _ = v34
                        } else {
                            for fi35 := 0; fi35 < 4; fi35++ {
                                var v36 float64 = 1e10
                                _ = fi35
                                _ = v36
                            }
                        }
                        var v37 int64 = 42
                        var v38 int = 32767
                        _ = v37
                        _ = v38
                    } else {
                        v2 = (float32((0.1 / (1.5 / 1.5))) * 3.14159)
                    }
                    fmt.Println(v11, v12, v2)
                    var v39 uint32 = 65535
                    _ = v39
                } else {
                    switch "foo bar" {
                    case "hello":
                        v3 = v3
                    case "\n\t":
                        var v40 string = ("a" + (("hello" + ("hello" + "\n\t")) + (("\n\t" + (("" + ("hello" + "hello")) + ("" + "hello"))) + "foo bar")))
                        _ = v40
                    default:
                        var v41 uint32 = 444
                        _ = v41
                    }
                }
                var v42 int = -70
                _ = v42
            } else {
                var v43 int8 = (v3 % (((^(v3))) | 1))
                v44 := GenS4{F0: "\n\t"}
                _ = v43
                _ = v44
            }
            v29 = v29
            v3 = (v3 + (v3 << ((v11) & 63)))
            _ = fi32
        }
    case "hello":
        fmt.Println(v3, v11)
    default:
        if false {
            v2 -= v2
        }
    }
    v45 := v1[2147483648]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v45
}
