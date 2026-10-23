package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int8; F2 int16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint8) int32 {
    var v7 float64 = 0.1
    p0 = (p0 | p0)
    v8 := GenS1{F0: 42}
    fmt.Println(v7, p0)
    p0 -= 2
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 2
    var v3 int64 = 1
    var v4 uint8 = 0
    var v5 int32 = -2147483648
    v6 := []string{"a", ""}
    v9 := GenS1{F0: 2}
    switch -952 {
    case 128:
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 19; fi11++ {
                switch (fi10 & fi11) {
                case 9223372036854775807:
                    v12 := GenS1{F0: 2}
                    _ = v12
                case -9223372036854775808:
                    if (v3 == (v3 | v3)) {
                        _ = len(v6)
                        v13 := GenS2{F0: 0}
                        _ = v13
                    }
                case 2147483647:
                    if false {
                        if len(v6) > 0 {
                            v6[0] = (("hello" + "\n\t") + "")
                        }
                        v14 := v1[2147483648]
                        if true {
                            var v15 int32 = v5
                            _ = v15
                        } else {
                            var v16 float32 = (1e10 - 0.1)
                            if false {
                                var v17 uint64 = v2
                                _ = v17
                            } else {
                                if (false || (v4 <= v4)) {
                                    v18 := GenS3{F0: 2, F1: 0, F2: 537}
                                    v2 = (v2 + 2)
                                    _ = v18
                                }
                            }
                            var v19 bool = true
                            _ = v16
                            _ = v19
                        }
                        _ = v14
                    } else {
                        v2 -= (^(v2))
                        if (0 != v2) {
                            var v20 string = ("a" + "foo bar")
                            if (127 < 127) {
                                var v21 uint32 = 42
                                _ = len(v6)
                                _ = v21
                            }
                            _ = v20
                        } else {
                            v22 := GenS4{F0: v3, F1: 100}
                            _ = v22
                        }
                    }
                }
                _ = fi11
            }
            _ = fi10
        }
    case -2147483648:
        v2 = v2
    case -334:
        if len(v6) > 0 {
            v6[0] = "a"
        }
    }
    fmt.Println(v3, v4)
    var v23 uint64 = (v2 - (v2 * 679))
    var v24 float32 = float32(0.0)
    var v25 bool = (!(int(v3) > -254))
    delete(v1, 65535)
    for fi26 := 0; fi26 < 9; fi26++ {
        if v25 {
            for fi27 := 0; fi27 < 16; fi27++ {
                if (!true) {
                    switch "hello" {
                    case "":
                        if v25 {
                            var v28 uint32 = uint32(v23)
                            var v29 int32 = v5
                            _ = v28
                            _ = v29
                        } else {
                            if (!v25) {
                                v3 -= 69
                                v30 := v1[uint32(v2)]
                                v5 -= 1
                                if ((v5 ^ int32(v3)) >= (-(v5))) {
                                    switch v30 {
                                    case -1:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    v4 = (^(v4))
                                }
                                _ = v30
                            } else {
                                v4 = (v4 % (((v4 / ((0) | 1))) | 1))
                            }
                            fmt.Println(v3, v5)
                        }
                    case "\n\t":
                        v31 := v1[65535]
                        _ = v31
                    default:
                        switch "\n\t" {
                        case "a":
                            v4 += v4
                        case "":
                            var v32 int32 = (v5 - 2147483647)
                            _ = v32
                        case "\n\t":
                            var v33 uint8 = (v4 ^ v4)
                            _ = v33
                        default:
                            var v34 int8 = -128
                            _ = v34
                        }
                    }
                    v35 := GenS5{F0: 2}
                    fmt.Println(v5, v24, v5)
                    _ = v35
                }
                v2 = 42
                v36 := GenS6{F0: v4}
                _ = fi27
                _ = v36
            }
        } else {
            fmt.Println(v4, v25)
            if false {
                v5 = (^(v5))
                var v37 uint32 = 2
                _ = v37
            } else {
                switch fi26 {
                case 0:
                    v4 = v4
                case -9223372036854775808:
                    _ = len(v6)
                default:
                    if ((!(1 < 18446744073709551615)) || (!((fi26 < (fi26 % ((fi26) | 1))) && (v25 || (true && v25))))) {
                        v38 := GenS1{F0: (v23 >> ((853) & 63))}
                        _ = v38
                    } else {
                        var v39 uint64 = (^(v2))
                        var v40 int = fi26
                        _ = v39
                        _ = v40
                    }
                }
                fmt.Println(fi26, v3)
                fmt.Println(v24)
            }
        }
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
        for fi41 := 0; fi41 < 10; fi41++ {
            v5 += v5
            _ = fi41
        }
        _ = fi26
    }
    if false {
        v42 := GenS7{F0: 2147483647}
        v43 := GenS6{F0: (^(v4))}
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 uint = 18446744073709551615
        if len(v6) > 0 {
            v6[0] = ("a" + "foo bar")
        }
        var v46 string = ("\n\t" + (("" + (("foo bar" + "a") + (("hello" + "\n\t") + "hello"))) + ("hello" + "\n\t")))
        _ = fi44
        _ = v45
        _ = v46
    }
    v6 = append(v6, "a")
    v47 := GenS5{F0: 1}
    var v48 int64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v9
}
