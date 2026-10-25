package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 string; F1 int; F2 bool }

func main() {
    var v1 uint16 = 255
    var v2 float64 = -2.25
    var v3 int16 = -32768
    v4 := make(map[string]bool)
    var v5 uint16 = (v1 + 651)
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                switch fi7 {
                case -654:
                    v9 := []int{fi8, fi8, fi7, fi7}
                    _ = v9
                default:
                    var v10 float32 = (((3.14159 + 1e10) * 0.1) * 1.5)
                    _ = v10
                }
            } else {
                if (true && (1 != -1)) {
                    var v11 int8 = 87
                    v12 := v4[("foo bar" + "\n\t")]
                    var v13 uint8 = 2
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := v4[""]
                var v15 int64 = 1
                var v16 uint64 = 18446744073709551615
                _ = v14
                _ = v15
                _ = v16
            }
            switch (fi7 >> (((v6 % ((v6) | 1))) & 63)) {
            case -1:
                if false {
                    fmt.Println(v1, fi7, v1)
                }
            case 1:
                fmt.Println(v2, v1, v3)
            default:
                switch "foo bar" {
                case "\n\t":
                    if (18446744073709551615 <= 2) {
                        v4["hello"] = ((65535 < 255) && false)
                        v1 = v5
                    } else {
                        if (42 == 45) {
                            delete(v4, "foo bar")
                            fmt.Println(fi7, v3)
                            v3 = (v3 / ((-32768) | 1))
                            v17 := GenS1{F0: v5}
                            _ = v17
                        }
                        fmt.Println(v5, v5, v2)
                        if true {
                            v5 = v1
                            v3 = -32768
                        } else {
                            var v18 int8 = 1
                            _ = v18
                        }
                    }
                default:
                    if ((!false) && (false || true)) {
                        v19 := GenS1{F0: v5}
                        _ = v19
                    }
                }
            }
            _ = fi8
        }
        if true {
            for fi20 := 0; fi20 < 12; fi20++ {
                fmt.Println(fi20, fi20, fi7)
                _ = fi20
            }
            var v21 int8 = -120
            var v22 bool = true
            _ = v21
            _ = v22
        }
        fmt.Println(v2, v5, v6)
        _ = fi7
    }
    fmt.Println(v5)
    var v23 int64 = -1
    v24 := GenS2{F0: 32767}
    v25 := GenS2{F0: 1}
    v26 := GenS2{F0: 1}
    v27 := []int{-661, -505, 32768, -9223372036854775808}
    delete(v4, "")
    _ = len(v4)
    v28 := GenS3{F0: "a", F1: 9}
    fmt.Println(v3)
    if false {
        if true {
            v29 := GenS2{F0: 127}
            v1 = (v1 ^ (v5 % (((^(v5))) | 1)))
            _ = v29
        } else {
            if true {
                v2 = (3.14159 / 1e10)
                if false {
                    _ = len(v27)
                    var v30 int8 = 88
                    _ = v30
                } else {
                    if (v5 > uint16(v6)) {
                        if false {
                            v31 := GenS1{F0: (v1 + 65535)}
                            v32 := GenS1{F0: (v1 | v5)}
                            if false {
                                if false {
                                    if false {
                                        v33 := v4["hello"]
                                        fmt.Println(v33)
                                        if v33 {
                                            v34 := v4["a"]
                                            v34 = v34
                                            delete(v4, "hello")
                                            _ = v34
                                        }
                                        _ = v33
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v35 int8 = 53
                                            _ = v35
                                        case "a":
                                            if false {
                                                v36 := GenS4{F0: v3, F1: "hello"}
                                                fmt.Println(v3)
                                                _ = v36
                                            }
                                        default:
                                            v5 = v5
                                        }
                                        if ((v23 * v23) < v23) {
                                            v37 := GenS5{F0: ("a" + "a"), F1: 0, F2: false}
                                            _ = v37
                                        }
                                        var v38 int32 = 2147483647
                                        _ = v38
                                    }
                                }
                                if len(v27) >= 2 {
                                    v27 = v27[1:len(v27)]
                                }
                            } else {
                                v5 -= 0
                            }
                            _ = v31
                            _ = v32
                        }
                        if false {
                            v39 := v4["foo bar"]
                            _ = v39
                        }
                        v2 = v2
                    }
                    v23 += (v23 - v23)
                    delete(v4, "")
                }
                if (-128 >= 42) {
                    v3 = 0
                    if false {
                        var v40 uint64 = 748
                        for fi41 := 0; fi41 < 1; fi41++ {
                            var v42 int64 = v23
                            var v43 bool = false
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        _ = v40
                    }
                    v44 := v4[""]
                    _ = v44
                } else {
                    if ((^(v6)) < 320) {
                        v5 -= v5
                        fmt.Println(v2, v6, v23)
                        v45 := GenS1{F0: (^(v1))}
                        _ = v45
                    } else {
                        v2 -= (v2 + 1.5)
                        fmt.Println(v3)
                    }
                    if true {
                        if (true && ((v1 - 979) < v1)) {
                            for fi46 := 0; fi46 < 7; fi46++ {
                                v4[("foo bar" + ("foo bar" + "hello"))] = false
                                _ = fi46
                            }
                        } else {
                            var v47 int32 = -215
                            v4[(("a" + "foo bar") + "foo bar")] = (!(1e10 > ((v2 * ((0.0 / v2) * (1.5 * ((v2 + (v2 + 1.5)) + v2)))) - float64((0.0 / 0.1)))))
                            _ = v47
                        }
                    } else {
                        var v48 uint64 = 19
                        var v49 uint32 = 65535
                        var v50 uint16 = v1
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                    _ = len(v4)
                }
            } else {
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            }
        }
        fmt.Println(v23)
    }
    v51 := GenS4{F0: v3, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
