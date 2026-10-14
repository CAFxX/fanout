package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 int16, p1 int16) int64 {
    switch 9223372036854775807 {
    case 70:
        if false {
            v5 := GenS1{F0: (p1 / ((-32768) | 1)), F1: (p1 ^ p0)}
            _ = v5
        } else {
            if false {
                var v6 int16 = 1000
                p1 = (v6 * (p1 + p1))
                _ = v6
            }
            v7 := GenS2{F0: (("a" + "\n\t") + ((("\n\t" + "a") + "\n\t") + "a")), F1: 127}
            switch 705 {
            case -1:
                p0 = p0
            case -2147483648:
                switch 61 {
                case 128:
                    var v8 int8 = 127
                    _ = v8
                case 868:
                    p1 = (p0 - (^(p1)))
                case 802:
                    if true {
                        p1 += p0
                    }
                default:
                    if false {
                        for fi9 := 0; fi9 < 6; fi9++ {
                            switch ("hello" + ("a" + ("foo bar" + "foo bar"))) {
                            case "foo bar":
                                var v10 int64 = -558
                                _ = v10
                            case "":
                                if (!true) {
                                    var v11 int64 = 891
                                    v12 := GenS2{F0: "foo bar", F1: 70}
                                    p1 += 1
                                    _ = v11
                                    _ = v12
                                } else {
                                    if false {
                                        fmt.Println(p1)
                                        for fi13 := 0; fi13 < 6; fi13++ {
                                            v14 := GenS2{F0: "foo bar", F1: 42}
                                            v15 := GenS1{F0: p1, F1: (p0 | (-(p1)))}
                                            v16 := []int{(fi13 - fi9), 128}
                                            _ = fi13
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                    }
                                }
                            }
                            v17 := GenS1{F0: 1, F1: p1}
                            for fi18 := 0; fi18 < 3; fi18++ {
                                v19 := []int{(^(fi18))}
                                fmt.Println(fi9)
                                var v20 bool = (!((p1 << ((717) & 63)) > (p0 - (^(p0)))))
                                v20 = (-53 != 373)
                                var v21 int64 = 42
                                _ = fi18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = fi9
                            _ = v17
                        }
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(p1)
                        case "foo bar":
                            var v22 uint = 0
                            _ = v22
                        case "hello":
                            p0 -= (^(p1))
                        default:
                            var v23 string = (("a" + "hello") + "hello")
                            _ = v23
                        }
                    }
                }
            }
            _ = v7
        }
    case -2147483648:
        v24 := []int{}
        _ = v24
    default:
        p0 = p1
    }
    var v25 int8 = int8(p0)
    _ = p0
    _ = p1
    _ = v25
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    v3 := []string{"\n\t", "", v2}
    v4 := []string{"hello", v2, ""}
    var v26 int32 = -1
    for fi27 := 0; fi27 < 19; fi27++ {
        v26 += (v26 - (v26 | -572))
        _ = fi27
    }
    if false {
        v26 += -2147483648
    }
    var v28 int16 = 1
    if true {
        v26 += -888
        if true {
            v28 = v28
        } else {
            fmt.Println(v2, v28)
            _ = len(v3)
        }
    }
    v1[v2] = true
    fmt.Println(v28)
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v28, fi29)
        _ = fi29
    }
    if (3.14159 != 1.5) {
        if (1e10 != 0.1) {
            fmt.Println(v28, v26, v2)
            v30 := GenS3{F0: (!(!true))}
            switch -61 {
            case -1:
                fmt.Println(v2, v28)
            }
            _ = v30
        }
    }
    v31 := GenS2{F0: v2, F1: 255}
    var v32 int32 = (v26 >> ((351) & 63))
    var v33 uint64 = 895
    _ = len(v3)
    switch 9223372036854775807 {
    case 277:
        for fi34 := 0; fi34 < 6; fi34++ {
            if (42 >= 37) {
                v32 += v32
            }
            _ = fi34
        }
    case 1:
        var v35 int32 = v26
        _ = v35
    case 9223372036854775807:
        v36 := GenS4{F0: (v33 << ((900) & 63)), F1: ("" + (v2 + ("foo bar" + "\n\t"))), F2: (v26 != (v32 << ((895) & 63)))}
        _ = v36
    default:
        v37 := GenS4{F0: v33, F1: v2, F2: true}
        _ = v37
    }
    v2 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
