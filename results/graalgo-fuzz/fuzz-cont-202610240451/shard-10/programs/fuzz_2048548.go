package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS5 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64, p1 int8) string {
    var v7 string = ((("\n\t" + "a") + "") + "a")
    var v8 int = 4294967295
    v7 = ((("a" + ("a" + ((v7 + ((((v7 + v7) + (v7 + v7)) + v7) + v7)) + v7))) + "") + "a")
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ("\n\t" + "hello")
}
func genFn1(p0 string, p1 int16) string {
    p1 = p1
    if true {
        if (false && ((false && (!false)) || (0 < 255))) {
            var v9 int = 4294967295
            if true {
                var v10 float32 = 1.5
                v11 := []int{(v9 + v9), v9, v9, (v9 + -1)}
                _ = v10
                _ = v11
            }
            _ = v9
        }
        v12 := GenS1{F0: 0}
        _ = v12
    } else {
        var v13 int32 = 2147483647
        switch -690 {
        case 32767:
            fmt.Println(p0)
        default:
            var v14 uint = 18446744073709551615
            _ = v14
        }
        _ = v13
    }
    fmt.Println(p1, p0)
    var v15 uint = 481
    _ = p0
    _ = p1
    _ = v15
    return (p0 + "foo bar")
}

func main() {
    var v1 string = "hello"
    var v2 int = 32767
    var v3 string = ((v1 + ("\n\t" + "hello")) + v1)
    var v4 bool = ((v2 | v2) < -915)
    v5 := make(map[string]bool)
    var v6 int8 = -1
    for fi16 := 0; fi16 < 11; fi16++ {
        if true {
            v17 := []int{(-(fi16)), (fi16 & v2)}
            if (1 >= 2147483647) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    if (42 > -730) {
                        v3 = "foo bar"
                        v4 = v4
                    }
                    _ = fi18
                }
                fmt.Println(v3, v6, v6)
                var v19 float32 = 0.0
                _ = v19
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                switch (fi20 % ((fi16) | 1)) {
                case 127:
                    v21 := GenS2{F0: 32767}
                    _ = v21
                case 4294967295:
                    var v22 uint = 2
                    _ = v22
                }
                var v23 bool = (0 < 0)
                _ = fi20
                _ = v23
            }
            _ = v17
        } else {
            v6 = (v6 - v6)
            for fi24 := 0; fi24 < 4; fi24++ {
                v25 := []int{9223372036854775807, fi24}
                v1 = v1
                switch (v2 & fi24) {
                case -441:
                    if v4 {
                        v1 = "\n\t"
                        var v26 int64 = 42
                        _ = v26
                    } else {
                        if v4 {
                            delete(v5, ((v3 + v1) + v1))
                        } else {
                            v4 = v4
                            var v27 int = fi16
                            if (1 <= 887) {
                                var v28 int = (v2 | 32768)
                                switch "\n\t" {
                                case "foo bar":
                                    _ = len(v5)
                                    var v29 uint = 42
                                    v29 -= (v29 * 0)
                                    var v30 string = ""
                                    var v31 uint64 = 2
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                case "a":
                                    var v32 uint8 = 170
                                    _ = v32
                                }
                                if false {
                                    v33 := GenS2{F0: 65535}
                                    v6 += (v6 - -1)
                                    _ = v33
                                } else {
                                    var v34 float32 = 0.0
                                    _ = v34
                                }
                                if (v4 || v4) {
                                    v35 := GenS2{F0: 959}
                                    if v4 {
                                        var v36 uint8 = 128
                                        _ = v36
                                    }
                                    _ = v35
                                } else {
                                    v37 := GenS3{F0: 65535}
                                    _ = v37
                                }
                                _ = v28
                            } else {
                                v27 -= (-(fi24))
                                v38 := GenS3{F0: 2}
                                var v39 int16 = 1000
                                if v4 {
                                    v40 := GenS1{F0: 128}
                                    v41 := GenS2{F0: 2}
                                    _ = v40
                                    _ = v41
                                } else {
                                    switch (-(fi24)) {
                                    case 0:
                                        var v42 int8 = int8(v39)
                                        _ = v42
                                    case 2147483647:
                                        v43 := GenS4{F0: 65535, F1: 2, F2: (v39 | int16(v27))}
                                        _ = v43
                                    default:
                                        v27 -= v2
                                    }
                                }
                                if v4 {
                                    v39 = (-(v39))
                                    v44 := GenS3{F0: 2}
                                    _ = v44
                                } else {
                                    if v4 {
                                        fmt.Println(v27, v39, fi24)
                                    } else {
                                        if len(v25) > 0 {
                                            v25[0] = fi16
                                        }
                                    }
                                }
                                v45 := GenS5{F0: v6, F1: 479}
                                _ = v38
                                _ = v39
                                _ = v45
                            }
                            fmt.Println(v27)
                            _ = v27
                        }
                        v4 = v4
                        v4 = true
                        switch v2 {
                        case 0:
                            v2 -= v2
                        case -121:
                            var v46 int32 = 42
                            _ = v46
                        default:
                            v2 += 357
                        }
                    }
                case 2147483647:
                    var v47 uint16 = 652
                    _ = v47
                default:
                    var v48 int64 = -9223372036854775808
                    _ = v48
                }
                _ = fi24
                _ = v25
            }
            v49 := []int{1}
            _ = v49
        }
        _ = fi16
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        v3 = ((v1 + ("a" + v1)) + (("\n\t" + (v1 + "a")) + ((v1 + "") + (v3 + "\n\t"))))
        delete(v5, "a")
        var v51 string = (v1 + v3)
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        var v53 uint16 = 0
        _ = fi52
        _ = v53
    }
    v3 = v1
    var v54 float32 = 1.5
    switch (v2 + 32767) {
    case -913:
        delete(v5, "a")
    case 127:
        var v55 string = ("a" + v3)
        _ = v55
    case 466:
        var v56 uint = 255
        _ = v56
    default:
        var v57 int16 = -1
        _ = v57
    }
    v58 := GenS1{F0: 128}
    fmt.Println(v3, v4, v1)
    if (606 >= 0) {
        fmt.Println(v6)
        var v59 uint = 718
        _ = v59
    }
    v6 = (^(v6))
    v3 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v58
    _ = v6
}
