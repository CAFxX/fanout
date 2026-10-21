package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 int32; F1 int32; F2 string }
type GenS7 struct { F0 string; F1 int64; F2 int16 }
type GenS8 struct { F0 uint; F1 string }
func genFn0(p0 int32, p1 uint32) int16 {
    p1 -= (^(p1))
    v7 := []int{0, 128, 127}
    var v8 float32 = float32((1.5 / 0.1))
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]string)
    var v3 uint32 = 0
    var v4 float32 = 1.5
    var v5 uint64 = 484
    v6 := []uint32{(^(v3)), (v3 - (v3 >> ((v1) & 63)))}
    fmt.Println(v1)
    v9 := GenS1{F0: -32768, F1: 2147483647}
    v4 = -2.25
    if true {
        if false {
            for fi10 := 0; fi10 < 16; fi10++ {
                if true {
                    fmt.Println(v4)
                }
                fmt.Println(v4, v3, v4)
                var v11 float32 = 1.5
                _ = fi10
                _ = v11
            }
            v4 = v4
        }
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                var v13 float32 = v4
                v2[""] = "hello"
                if true {
                    v6 = append(v6, (v3 << ((v1) & 63)))
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 uint8 = 78
                        v4 += (v4 + (v4 + ((v13 * -2.25) - v4)))
                        _ = fi14
                        _ = v15
                    }
                } else {
                    v2[("" + "")] = ("\n\t" + "")
                }
                _ = v13
            }
            var v16 int = (-(fi12))
            _ = fi12
            _ = v16
        }
        var v17 float64 = ((0.0 / -2.25) / -2.25)
        _ = v17
    }
    v18 := GenS2{F0: -157, F1: (v3 * (v3 >> ((0) & 63))), F2: v3}
    var v19 int32 = -187
    v2["hello"] = "hello"
    v6 = append(v6, v3)
    var v20 int32 = v19
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 float64 = 0.1
            var v24 int = (fi21 << ((v1) & 63))
            _ = fi22
            _ = v23
            _ = v24
        }
        v1 = 604
        _ = fi21
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        if (true || (int32(v1) > v19)) {
            switch -9223372036854775808 {
            case 896:
                switch (fi25 * (fi25 << (((v1 >> ((0) & 63))) & 63))) {
                case -2147483648:
                    var v26 uint32 = v3
                    _ = v26
                default:
                    if true {
                        v27 := GenS2{F0: 115, F1: (^(v3)), F2: v3}
                        switch "\n\t" {
                        case "":
                            if true {
                                var v28 uint8 = 255
                                v5 -= 57
                                fmt.Println(v3)
                                _ = v28
                            }
                        case "\n\t":
                            v29 := GenS3{F0: 191}
                            _ = v29
                        default:
                            v5 = v5
                        }
                        _ = v27
                    }
                }
            case -710:
                switch "\n\t" {
                case "\n\t":
                    for fi30 := 0; fi30 < 10; fi30++ {
                        if (-18 != -1) {
                            var v31 int32 = -549
                            var v32 string = "a"
                            _ = v31
                            _ = v32
                        } else {
                            v1 = v1
                            switch ("\n\t" + ("foo bar" + "\n\t")) {
                            case "a":
                                if true {
                                    var v33 uint32 = v3
                                    var v34 bool = true
                                    fmt.Println(v34, v5)
                                    switch ("hello" + "foo bar") {
                                    case "\n\t":
                                        v19 = (v19 & (v19 % ((v20) | 1)))
                                    case "foo bar":
                                        v35 := GenS4{F0: (v20 << ((v1) & 63))}
                                        _ = v35
                                    default:
                                        if false {
                                            v4 -= 1.5
                                            v36 := v2[("" + "a")]
                                            v34 = v34
                                            _ = v36
                                        }
                                    }
                                    v37 := GenS5{F0: (v1 / ((4294967295) | 1)), F1: v1, F2: 259}
                                    var v38 int16 = 32767
                                    _ = v33
                                    _ = v34
                                    _ = v37
                                    _ = v38
                                }
                            case "":
                                v39 := GenS6{F0: v20, F1: (v20 / ((v20) | 1)), F2: ("a" + ("foo bar" + "foo bar"))}
                                _ = v39
                            case "foo bar":
                                delete(v2, "")
                            }
                            v1 += v1
                        }
                        _ = fi30
                    }
                default:
                    var v40 uint64 = (v5 + v5)
                    _ = v40
                }
            }
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        }
        v20 += v19
        _ = fi25
    }
    if (!true) {
        v41 := GenS1{F0: 1000, F1: (v20 << (((v1 | uint(v5))) & 63))}
        _ = v41
    }
    switch 0 {
    case -367:
        var v42 uint = v1
        _ = v42
    default:
        fmt.Println(v4, v5, v3)
    }
    switch 192 {
    case 32767:
        for fi43 := 0; fi43 < 1; fi43++ {
            switch "a" {
            case "foo bar":
                v44 := GenS7{F0: "foo bar", F1: 9223372036854775807, F2: -32768}
                _ = v44
            case "hello":
                switch ("" + (("" + "\n\t") + "\n\t")) {
                case "\n\t":
                    var v45 string = "foo bar"
                    _ = v45
                }
            default:
                for fi46 := 0; fi46 < 19; fi46++ {
                    if true {
                        v3 = 4294967295
                    }
                    v6 = append(v6, uint32(v1))
                    _ = fi46
                }
            }
            _ = fi43
        }
    case -199:
        v1 -= 36
    }
    var v47 int = 9223372036854775807
    for fi48 := 0; fi48 < 11; fi48++ {
        v49 := GenS8{F0: 1, F1: "\n\t"}
        _ = fi48
        _ = v49
    }
    v1 += (v1 % (((v1 ^ v1)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v9
}
