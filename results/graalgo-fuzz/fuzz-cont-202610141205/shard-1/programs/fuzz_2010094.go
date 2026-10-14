package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 bool, p1 bool) uint {
    if true {
        if false {
            p1 = true
            if false {
                for fi5 := 0; fi5 < 1; fi5++ {
                    fmt.Println(p0)
                    var v6 int = fi5
                    p1 = (p0 && ((false || p1) && (!(!true))))
                    _ = fi5
                    _ = v6
                }
                fmt.Println(p0, p1)
            }
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 2; fi8++ {
                    v9 := GenS1{F0: -1, F1: "foo bar", F2: 42}
                    if (!p0) {
                        v10 := []int{}
                        v11 := GenS1{F0: 1000, F1: "a", F2: 127}
                        switch "" {
                        case "foo bar":
                            switch (("\n\t" + "hello") + "hello") {
                            case "hello":
                                var v12 float32 = 0.1
                                _ = v12
                            case "\n\t":
                                var v13 uint32 = 42
                                _ = v13
                            case "foo bar":
                                if p1 {
                                    p1 = p0
                                }
                            default:
                                if len(v10) > 0 {
                                    v10[0] = fi7
                                }
                            }
                        case "":
                            var v14 bool = (-521 < 0)
                            _ = v14
                        case "a":
                            p1 = (0.0 == (0.0 - float32((-2.25 - 3.14159))))
                        default:
                            if (fi8 <= fi7) {
                                if (127 >= 128) {
                                    var v15 string = "foo bar"
                                    _ = v15
                                }
                            }
                        }
                        _ = v10
                        _ = v11
                    } else {
                        if ((460 >= 255) && p1) {
                            var v16 uint64 = 660
                            var v17 uint64 = (v16 ^ (v16 ^ v16))
                            p1 = p1
                            _ = v16
                            _ = v17
                        } else {
                            p1 = false
                            v18 := GenS2{F0: "a"}
                            _ = v18
                        }
                        fmt.Println(fi8, p1)
                        var v19 int = (fi8 ^ fi8)
                        v19 -= 498
                        var v20 bool = (!(778 == 666))
                        _ = v19
                        _ = v20
                    }
                    _ = fi8
                    _ = v9
                }
                v21 := []int{42, fi7}
                p0 = p0
                _ = fi7
                _ = v21
            }
        } else {
            var v22 uint = 728
            if p1 {
                v23 := GenS2{F0: "\n\t"}
                p1 = p1
                _ = v23
            } else {
                var v24 float64 = (3.14159 + (0.1 / (1e10 - 3.14159)))
                _ = v24
            }
            _ = v22
        }
    } else {
        v25 := []int{32767, 4294967295, 3}
        fmt.Println(p1, p1)
        v25 = append(v25, 562)
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint8 = 2
            v29 := GenS3{F0: 1}
            if p1 {
                if false {
                    var v30 uint16 = 836
                    v31 := GenS1{F0: -1, F1: ("foo bar" + "foo bar"), F2: (v28 ^ v28)}
                    _ = v30
                    _ = v31
                } else {
                    v32 := []int{(fi27 ^ fi27), (fi27 | fi27), fi26}
                    _ = v32
                }
            } else {
                p1 = true
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        var v33 uint16 = 365
        _ = fi26
        _ = v33
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 bool = false
    var v2 int16 = 1
    v3 := make(map[string]int)
    v4 := []int{32768, -371, 952, -869}
    if true {
        var v34 int16 = (v2 & (v2 - 0))
        v35 := GenS4{F0: -1, F1: v1, F2: -563}
        _ = v34
        _ = v35
    } else {
        v1 = (4294967295 != 2147483647)
        v2 += v2
    }
    v1 = true
    v1 = (((!((false && v1) || (v2 <= 1))) || true) || v1)
    v36 := GenS5{F0: 906, F1: 1}
    var v37 uint64 = 18446744073709551615
    v38 := v3["foo bar"]
    _ = len(v4)
    switch v38 {
    case 42:
        _ = len(v3)
    default:
        fmt.Println(v1, v2, v1)
    }
    v39 := GenS3{F0: (v37 | (^(v37)))}
    var v40 bool = (v1 && false)
    v37 -= (v37 & v37)
    v41 := GenS2{F0: ((((("hello" + ("\n\t" + "")) + "foo bar") + ("hello" + ("hello" + "hello"))) + "hello") + ("hello" + "foo bar"))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
