package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int16; F1 int64; F2 bool }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 int8, p1 uint8, p2 bool) uint8 {
    p1 = p1
    fmt.Println(p1, p0)
    var v7 bool = p2
    var v8 int16 = -32768
    fmt.Println(v7, p2, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 int64 = -885
    v4 := make(map[uint32]bool)
    var v5 uint64 = 4294967295
    var v6 int8 = 109
    v4[2] = (2 <= (v5 + (v5 ^ (^(v5)))))
    var v9 uint16 = 544
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v9)
        _ = fi10
    }
    v3 = -484
    fmt.Println(v6, v6, v6)
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            if true {
                switch 826 {
                case 9223372036854775807:
                    var v13 int16 = 230
                    _ = v13
                case 1:
                    v14 := GenS1{F0: 1, F1: "foo bar"}
                    _ = v14
                default:
                    fmt.Println(fi11, v3)
                }
                _ = len(v2)
                if false {
                    fmt.Println(fi11, fi12, v5)
                    _ = len(v1)
                    fmt.Println(v3, v3)
                    fmt.Println(v9, v6, fi11)
                    v15 := GenS1{F0: 0, F1: "hello"}
                    _ = v15
                } else {
                    switch (("" + "a") + "a") {
                    case "":
                        var v16 uint16 = (v9 - v9)
                        _ = v16
                    }
                    if false {
                        v17 := GenS1{F0: int32(v6), F1: ""}
                        v1 = append(v1, "foo bar")
                        _ = v17
                    } else {
                        fmt.Println(v5, v6, v3)
                    }
                    v18 := GenS1{F0: -1, F1: ("hello" + "hello")}
                    _ = v18
                }
            }
            if true {
                v19 := GenS1{F0: -713, F1: "a"}
                _ = v19
            } else {
                switch "a" {
                case "foo bar":
                    if (!(0.0 >= 0.0)) {
                        v4[65535] = true
                        fmt.Println(v9, v3)
                    }
                }
                v5 -= v5
            }
            if (v6 < v6) {
                switch "hello" {
                case "\n\t":
                    v6 += 127
                case "hello":
                    if (1 >= fi12) {
                        var v20 int64 = -76
                        if true {
                            var v21 bool = false
                            var v22 uint = 65535
                            if v21 {
                                var v23 float32 = (3.14159 + 0.0)
                                var v24 string = "\n\t"
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(v3, fi12)
                            _ = v21
                            _ = v22
                        } else {
                            v20 = (-(v3))
                        }
                        var v25 int = 128
                        _ = v20
                        _ = v25
                    }
                case "a":
                    switch (fi12 / ((42) | 1)) {
                    case -1:
                        if true {
                            var v26 uint8 = genFn0(0, 0, false)
                            var v27 int8 = v6
                            v28 := GenS2{F0: 0, F1: 992, F2: genFn0(v6, genFn0(-1, genFn0(101, genFn0(-1, (v26 - (v26 >> ((2) & 63))), (!true)), (!true)), true), true)}
                            v29 := GenS1{F0: 512, F1: "\n\t"}
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            v3 = v3
                            v9 += (v9 % (((v9 + v9)) | 1))
                            delete(v2, "")
                        }
                    case 4294967295:
                        v9 = (^(v9))
                    case 312:
                        fmt.Println(v6, fi12)
                    }
                default:
                    switch "hello" {
                    case "a":
                        delete(v4, 2147483647)
                    case "":
                        var v30 string = "a"
                        _ = v30
                    }
                }
                v6 = (v6 % ((0) | 1))
                if (fi11 <= (fi11 ^ (fi12 | 1))) {
                    fmt.Println(v6, fi12, v3)
                    var v31 uint8 = 128
                    v32 := GenS3{F0: 1000, F1: (v3 - v3), F2: false}
                    _ = v31
                    _ = v32
                }
            } else {
                v6 += v6
            }
            _ = fi12
        }
        _ = fi11
    }
    v3 = v3
    _ = len(v4)
    switch ("" + "a") {
    case "\n\t":
        if false {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        } else {
            if (!(!(!("foo bar" < "\n\t")))) {
                v6 = (^(v6))
                switch 814 {
                case 2147483647:
                    if false {
                        if false {
                            v6 = (v6 - -1)
                            v9 -= 42
                            var v33 int32 = -2147483648
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v6 -= v6
                            var v35 uint16 = (v9 | v9)
                            fmt.Println(v3, v35)
                            _ = fi34
                            _ = v35
                        }
                        fmt.Println(v5)
                    }
                case 387:
                    v3 -= v3
                default:
                    delete(v4, 4294967295)
                }
            }
            fmt.Println(v9, v5)
            if true {
                v36 := GenS3{F0: 42, F1: -55, F2: false}
                _ = v36
            }
        }
    case "foo bar":
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 uint8 = 42
            fmt.Println(v9)
            var v39 uint8 = (v38 << ((688) & 63))
            _ = fi37
            _ = v38
            _ = v39
        }
    default:
        fmt.Println(v5)
    }
    v40 := GenS4{F0: 1, F1: -563}
    for fi41 := 0; fi41 < 1; fi41++ {
        fmt.Println(v5)
        _ = fi41
    }
    var v42 float64 = (1.5 / 1.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
    _ = v9
}
