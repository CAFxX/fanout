package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 bool) bool {
    var v7 float32 = (3.14159 - 1e10)
    v8 := []int{}
    if (2 < 1) {
        var v9 int32 = -2147483648
        for fi10 := 0; fi10 < 8; fi10++ {
            if p0 {
                switch (-(fi10)) {
                case 32768:
                    v8 = append(v8, (fi10 + (fi10 >> ((772) & 63))))
                default:
                    var v11 uint = 536
                    _ = v11
                }
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 uint16 = 819
                    _ = fi12
                    _ = v13
                }
                if (p0 && p0) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v14 uint64 = 42
                    _ = v14
                } else {
                    v15 := GenS1{F0: 4294967295, F1: v9}
                    _ = v15
                }
            } else {
                if false {
                    var v16 int16 = -1
                    _ = v16
                }
                switch "foo bar" {
                case "foo bar":
                    switch fi10 {
                    case -1:
                        var v17 uint = 1
                        _ = v17
                    case 127:
                        var v18 string = "foo bar"
                        _ = v18
                    case 32767:
                        var v19 int16 = 1
                        _ = v19
                    }
                case "":
                    var v20 string = "a"
                    _ = v20
                }
                switch (fi10 >> ((0) & 63)) {
                case -2147483648:
                    v9 = v9
                case 42:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        if (0 == 48) {
                            p0 = ((76 != 42) && (!(int64(v9) == -218)))
                        } else {
                            switch (fi21 ^ (fi10 + -822)) {
                            case 2147483647:
                                v7 -= float32((1.5 - -2.25))
                            case 32767:
                                v9 -= -1
                                if false {
                                    var v22 int = fi10
                                    _ = v22
                                }
                            case 185:
                                var v23 uint = 65535
                                var v24 int16 = 42
                                if true {
                                    switch "a" {
                                    case "hello":
                                        if len(v8) > 0 {
                                            v8[0] = fi10
                                        }
                                    case "a":
                                        v9 += (v9 + 0)
                                    case "":
                                        var v25 uint16 = 42
                                        _ = v25
                                    }
                                    if false {
                                        if p0 {
                                            _ = len(v8)
                                            v26 := GenS2{F0: (^(v9)), F1: 65535}
                                            _ = v26
                                        }
                                        v7 = 3.14159
                                        v27 := GenS2{F0: (v9 ^ v9), F1: 2147483647}
                                        _ = v27
                                    }
                                    if true {
                                        v28 := GenS2{F0: 0, F1: 1}
                                        _ = v28
                                    } else {
                                        var v29 bool = (!p0)
                                        _ = v29
                                    }
                                }
                                var v30 int64 = 627
                                _ = v23
                                _ = v24
                                _ = v30
                            }
                        }
                        fmt.Println(v9, fi21)
                        if (true && false) {
                            var v31 int = 9223372036854775807
                            fmt.Println(v9, v7, fi10)
                            _ = v31
                        } else {
                            if p0 {
                                var v32 string = "hello"
                                p0 = p0
                                _ = v32
                            } else {
                                switch 1 {
                                case 128:
                                    v33 := GenS1{F0: 608, F1: (v9 >> ((42) & 63))}
                                    _ = v33
                                default:
                                    p0 = (!false)
                                }
                            }
                            v34 := GenS3{F0: ((((("hello" + "") + ("a" + "hello")) + "a") + ("foo bar" + ("foo bar" + ("\n\t" + "foo bar")))) + "hello"), F1: 792}
                            _ = v34
                        }
                        _ = fi21
                    }
                }
            }
            if true {
                var v35 uint16 = 2
                _ = v35
            }
            _ = fi10
        }
        _ = v9
    }
    v7 = (((v7 * (v7 * ((0.0 / v7) / ((3.14159 + (0.0 / v7)) / (-2.25 + 1e10))))) / (1e10 - v7)) - v7)
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint = 4294967295
    v2 := []uint32{2147483647, 2147483647, 0, 2147483647}
    var v3 int16 = 403
    var v4 uint = v1
    v5 := []string{("\n\t" + "foo bar"), ("hello" + "hello")}
    var v6 int8 = 0
    v6 += v6
    for fi36 := 0; fi36 < 2; fi36++ {
        if false {
            fmt.Println(v6, v4, v4)
            fmt.Println(fi36)
            v37 := GenS1{F0: (v4 >> ((v1) & 63)), F1: int32(v6)}
            _ = v37
        }
        var v38 int32 = -1
        _ = fi36
        _ = v38
    }
    v6 = (v6 ^ v6)
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch 32767 {
    case 2147483647:
        v4 += (v1 & (v4 ^ v1))
    case -2147483648:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -292:
        fmt.Println(v1)
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v5 = append(v5, "\n\t")
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = fi40
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi39
    }
    for fi41 := 0; fi41 < 11; fi41++ {
        fmt.Println(v3, v4, v3)
        _ = fi41
    }
    fmt.Println(v6, v3)
    if ((1 <= 42) && false) {
        if false {
            if (!false) {
                if (42 >= 1) {
                    if true {
                        v42 := GenS1{F0: v4, F1: -1}
                        v43 := GenS4{F0: v3}
                        _ = v42
                        _ = v43
                    }
                    if true {
                        var v44 int32 = 0
                        _ = v44
                    } else {
                        var v45 int64 = int64(v6)
                        _ = v45
                    }
                    v1 -= v4
                } else {
                    var v46 uint = (v1 << ((2) & 63))
                    _ = v46
                }
                var v47 float32 = 0.0
                _ = v47
            }
        } else {
            fmt.Println(v4, v6)
            var v48 float32 = 0.0
            switch "a" {
            case "":
                if (("a" + "") >= "foo bar") {
                    var v49 uint16 = 2
                    v50 := GenS2{F0: -160, F1: 2}
                    _ = v49
                    _ = v50
                } else {
                    var v51 int = -464
                    var v52 int32 = 0
                    _ = v51
                    _ = v52
                }
            case "\n\t":
                var v53 float64 = 1.5
                _ = v53
            default:
                var v54 int8 = -118
                _ = v54
            }
            _ = v48
        }
        v6 = -102
    } else {
        v1 = (^(v4))
        var v55 uint32 = 2147483647
        _ = v55
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
