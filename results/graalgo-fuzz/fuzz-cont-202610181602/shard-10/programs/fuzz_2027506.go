package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32; F1 bool }
func genFn0(p0 int64) string {
    var v6 int32 = -298
    v6 = (v6 % ((v6) | 1))
    _ = p0
    _ = v6
    return "\n\t"
}
func genFn1(p0 bool, p1 uint, p2 uint16) bool {
    for fi7 := 0; fi7 < 9; fi7++ {
        if (!(p0 || true)) {
            fmt.Println(fi7, p2)
            v8 := GenS1{F0: "foo bar"}
            switch (^(fi7)) {
            case 9223372036854775807:
                if p0 {
                    var v9 bool = (p0 && p0)
                    if ((^(fi7)) != (fi7 << (((^(p1))) & 63))) {
                        var v10 uint16 = (p2 - (p2 << ((p1) & 63)))
                        var v11 string = "a"
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(fi7, p1, p2)
                        fmt.Println(p0)
                    }
                    var v12 uint64 = 262
                    _ = v12
                    _ = v9
                }
            case 32767:
                for fi13 := 0; fi13 < 16; fi13++ {
                    p0 = true
                    v14 := []int{(fi7 ^ (fi13 | 0)), fi13, 45}
                    if false {
                        v15 := GenS1{F0: "foo bar"}
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = v8
        } else {
            switch "\n\t" {
            case "hello":
                v16 := []int{fi7, fi7, -1}
                _ = v16
            default:
                var v17 uint = p1
                _ = v17
            }
            if (!true) {
                fmt.Println(p1, p1, p1)
            }
        }
        _ = fi7
    }
    p2 = p2
    for fi18 := 0; fi18 < 15; fi18++ {
        p1 += uint(p2)
        var v19 uint8 = 0
        switch (fi18 ^ fi18) {
        case 127:
            fmt.Println(v19, v19, v19)
        case 32767:
            p0 = (p0 && true)
        }
        _ = fi18
        _ = v19
    }
    switch 2147483647 {
    case 128:
        fmt.Println(p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 int8 = 42
    var v2 bool = false
    v3 := make(map[string]int)
    var v4 uint16 = 0
    v5 := []uint32{457, 4294967295, 853, 42}
    if true {
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(fi20)
            if true {
                for fi21 := 0; fi21 < 19; fi21++ {
                    switch ("a" + ("" + ("hello" + ("\n\t" + ("hello" + ("" + (("\n\t" + "hello") + "a"))))))) {
                    case "\n\t":
                        var v22 string = "a"
                        _ = v22
                    case "foo bar":
                        if v2 {
                            var v23 float64 = 1.5
                            switch (("a" + "") + ((("\n\t" + "a") + "a") + "\n\t")) {
                            case "\n\t":
                                v24 := GenS1{F0: ""}
                                _ = v24
                            }
                            v1 = v1
                            if v2 {
                                fmt.Println(v1)
                            } else {
                                _ = len(v3)
                                fmt.Println(v1)
                                v2 = (float64(0.1) <= float64(0.0))
                            }
                            v25 := GenS2{F0: 344}
                            _ = v23
                            _ = v25
                        } else {
                            switch "a" {
                            case "\n\t":
                                v26 := GenS1{F0: ("\n\t" + "\n\t")}
                                _ = v26
                            }
                            v1 = v1
                            v27 := GenS3{F0: -76, F1: v1, F2: 0}
                            _ = v27
                        }
                    case "":
                        v28 := GenS2{F0: (v4 & v4)}
                        _ = v28
                    }
                    var v29 uint16 = 65535
                    v30 := GenS2{F0: v4}
                    _ = fi21
                    _ = v29
                    _ = v30
                }
                delete(v3, ("a" + ("foo bar" + ("foo bar" + "a"))))
                fmt.Println(v2)
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
            _ = fi20
        }
        v4 = v4
        var v31 float64 = 0.1
        _ = v31
    }
    v1 = (v1 - (v1 * (-(v1))))
    v32 := GenS4{F0: 127, F1: 0}
    var v33 int = 206
    var v34 uint64 = 70
    v35 := GenS5{F0: -1, F1: (2 != 4294967295)}
    var v36 int64 = 9223372036854775807
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 bool = v2
        _ = fi37
        _ = v38
    }
    _ = len(v5)
    for fi39 := 0; fi39 < 15; fi39++ {
        for fi40 := 0; fi40 < 8; fi40++ {
            var v41 bool = v2
            _ = fi40
            _ = v41
        }
        fmt.Println(v2)
        _ = fi39
    }
    v1 -= (v1 & 42)
    if (!false) {
        switch 32768 {
        case 0:
            if ((((1.5 - 1.5) - ((-2.25 / (3.14159 + (1e10 + 0.1))) / float64(float32(0.1)))) - float64(1e10)) >= (float64(0.0) - -2.25)) {
                switch "foo bar" {
                case "a":
                    v2 = false
                case "\n\t":
                    v42 := GenS5{F0: 1, F1: false}
                    _ = v42
                default:
                    for fi43 := 0; fi43 < 18; fi43++ {
                        for fi44 := 0; fi44 < 8; fi44++ {
                            var v45 int = 9223372036854775807
                            var v46 uint64 = 385
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        _ = fi43
                    }
                }
            } else {
                _ = len(v3)
                fmt.Println(v36, v33, v33)
            }
        default:
            v36 = -1
        }
    }
    var v47 float64 = 1e10
    v1 -= (v1 >> ((845) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v47
    _ = v5
}
