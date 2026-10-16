package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int16; F1 uint64; F2 uint16 }
func genFn0() int32 {
    switch 127 {
    case -131:
        var v6 uint = 42
        _ = v6
    default:
        var v7 uint32 = 1
        _ = v7
    }
    var v8 uint = 205
    fmt.Println(v8)
    _ = v8
    return -1
}
func genFn1(p0 int, p1 int, p2 bool) string {
    v9 := GenS1{F0: p2, F1: 655}
    switch -1 {
    case 128:
        v10 := []int{p1}
        _ = v10
    case -860:
        v11 := GenS2{F0: p2}
        _ = v11
    default:
        var v12 int16 = 32767
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        fmt.Println(fi13)
        p2 = (0 != 0)
        _ = fi13
    }
    v14 := GenS2{F0: (42 != 18446744073709551615)}
    var v15 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v9
    return ("hello" + "\n\t")
}

func main() {
    var v1 int16 = -729
    v2 := []string{("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))}
    v3 := []int{127}
    var v4 bool = false
    var v5 int = -676
    var v16 int16 = 937
    switch "foo bar" {
    case "foo bar":
        var v17 bool = v4
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 2; fi18++ {
            _ = len(v2)
            _ = fi18
        }
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                if (!("hello" > "a")) {
                    if false {
                        v16 = v16
                        if ((255 >= 128) && v4) {
                            v4 = v4
                        }
                    } else {
                        if false {
                            fmt.Println(v1)
                            if (genFn0() >= -641) {
                                v1 -= 0
                                if true {
                                    if (!v4) {
                                        _ = len(v2)
                                        var v21 uint8 = 20
                                        _ = v21
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi19, v4, v1)
                                } else {
                                    v22 := GenS3{F0: 2, F1: 2147483647, F2: ""}
                                    fmt.Println(v4)
                                    _ = v22
                                }
                                switch "\n\t" {
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                case "a":
                                    if v4 {
                                        var v23 int8 = 45
                                        _ = v23
                                    }
                                default:
                                    var v24 int = -1
                                    _ = v24
                                }
                            } else {
                                v25 := GenS4{F0: 0}
                                v26 := GenS2{F0: v4}
                                _ = len(v2)
                                _ = v25
                                _ = v26
                            }
                            v2 = append(v2, "\n\t")
                            fmt.Println(fi20, v4, v5)
                        } else {
                            fmt.Println(fi19, v5, fi20)
                        }
                        v2 = append(v2, "")
                        var v27 int64 = 333
                        _ = v27
                    }
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v28 := GenS5{F0: 966}
                    v5 += -804
                    _ = v28
                }
                v2 = append(v2, ("\n\t" + "a"))
                _ = fi20
            }
            var v29 int8 = 41
            v16 = (v16 | 0)
            _ = fi19
            _ = v29
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v16, v1)
    } else {
        v16 = -32768
        v1 = v16
    }
    v16 -= (-(v1))
    if (((32767 == 1) || v4) && (int16(v5) == (^(v1)))) {
        v16 = (v1 >> ((6) & 63))
    } else {
        if false {
            switch v5 {
            case 1:
                v30 := GenS6{F0: (^(v1)), F1: 105}
                _ = v30
            case 138:
                v1 = (^(v1))
            case 0:
                var v31 int64 = 1
                _ = v31
            }
            v32 := GenS7{F0: (v16 - -1), F1: 0, F2: 32767}
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 float32 = -2.25
                _ = fi33
                _ = v34
            }
            _ = v32
        }
    }
    v35 := GenS5{F0: 2}
    v4 = true
    if v4 {
        var v36 int8 = 42
        v36 = 1
        _ = v36
    } else {
        fmt.Println(v16, v1, v4)
    }
    switch "hello" {
    case "foo bar":
        v4 = v4
    case "a":
        var v37 string = "foo bar"
        _ = v37
    case "":
        _ = len(v3)
    }
    var v38 int64 = -231
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v39 int8 = 111
    if len(v3) > 0 {
        v3[0] = v5
    }
    if true {
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 uint32 = 2147483647
            _ = fi40
            _ = v41
        }
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    } else {
        v38 -= v38
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v42 uint32 = 42
    v43 := GenS7{F0: -32768, F1: 519, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
