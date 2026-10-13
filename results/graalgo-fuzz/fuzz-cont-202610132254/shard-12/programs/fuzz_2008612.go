package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 bool }
func genFn0(p0 uint32, p1 uint16, p2 uint64) uint8 {
    switch 1 {
    case -1:
        if (!(2 == 873)) {
            v5 := GenS1{F0: -998, F1: 1}
            var v6 float32 = 0.1
            fmt.Println(p2)
            _ = v5
            _ = v6
        } else {
            p0 = (^(p0))
        }
    case 0:
        if (9223372036854775807 != 42) {
            var v7 string = (("a" + ("a" + "a")) + ("foo bar" + "hello"))
            _ = v7
        }
    case 1:
        var v8 bool = true
        _ = v8
    default:
        switch "foo bar" {
        case "hello":
            p1 = 407
        case "\n\t":
            v9 := []int{-9223372036854775808, -9223372036854775808, 9223372036854775807}
            _ = v9
        }
    }
    var v10 uint64 = p2
    p1 = uint16(p0)
    v11 := GenS1{F0: 1, F1: 112}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 113
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 795
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t", ""}
    switch "\n\t" {
    case "\n\t":
        v12 := GenS2{F0: 0, F1: v2}
        _ = v12
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "":
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int64 = 1
                _ = fi14
                _ = v15
            }
            if (337 >= 79) {
                var v16 uint = 65535
                v16 = (^(v16))
                if true {
                    v17 := GenS3{F0: fi13}
                    _ = v17
                } else {
                    switch (^(fi13)) {
                    case 281:
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v19 := GenS2{F0: 564, F1: (v2 * (-(v2)))}
                            if (42 >= 492) {
                                switch fi18 {
                                case -9223372036854775808:
                                    v20 := GenS4{F0: (v16 % (((v16 | (v16 - v16))) | 1))}
                                    _ = v20
                                default:
                                    var v21 bool = false
                                    _ = v21
                                }
                            } else {
                                delete(v3, (fi13 / ((fi13) | 1)))
                                if true {
                                    if (true && true) {
                                        v2 = 308
                                    } else {
                                        switch (fi13 + (fi18 | (fi13 | fi18))) {
                                        case -2147483648:
                                            fmt.Println(fi18, v2, v2)
                                        case 9223372036854775807:
                                            v22 := GenS4{F0: v16}
                                            _ = v22
                                        case -9223372036854775808:
                                            var v23 float32 = 0.1
                                            var v24 int64 = 0
                                            v25 := GenS1{F0: -30, F1: 0}
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 int = 4294967295
                                        v16 += v16
                                        _ = v26
                                    }
                                } else {
                                    switch ("" + "hello") {
                                    case "a":
                                        if (false && true) {
                                            v2 = -32768
                                        }
                                    }
                                    v16 -= (v16 >> ((255) & 63))
                                }
                                v16 += v16
                            }
                            _ = fi18
                            _ = v19
                        }
                    case 42:
                        v3[-1] = 65535
                    case -1:
                        v16 -= 1
                    default:
                        var v27 string = ((("hello" + ("\n\t" + ((("hello" + ("" + "foo bar")) + ((("" + "\n\t") + "hello") + ("hello" + ("\n\t" + "foo bar")))) + "a"))) + "foo bar") + "\n\t")
                        _ = v27
                    }
                }
                _ = v16
            } else {
                v2 += (v2 & (v2 >> ((255) & 63)))
            }
            _ = fi13
        }
    }
    v28 := GenS4{F0: 2}
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 uint32 = 65535
        _ = fi29
        _ = v30
    }
    v4 = append(v4, "foo bar")
    v31 := GenS4{F0: 255}
    switch ("foo bar" + "hello") {
    case "":
        var v32 uint = 191
        _ = v32
    case "a":
        fmt.Println(v2)
    case "foo bar":
        var v33 string = "\n\t"
        _ = v33
    default:
        var v34 string = "\n\t"
        _ = v34
    }
    var v35 uint16 = 885
    v3[4294967295] = 4294967295
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 bool = ((1 != 43) || false)
        _ = fi36
        _ = v37
    }
    v3[1] = 65535
    switch 9223372036854775807 {
    case -2147483648:
        v2 = (v2 - (v2 / (((v2 % (((v2 - (v2 % ((v2) | 1)))) | 1))) | 1)))
    case 128:
        switch 2147483647 {
        case 2147483647:
            var v38 int = 332
            _ = v38
        }
    }
    var v39 int = 4294967295
    v40 := GenS5{F0: (^(v39)), F1: false}
    for fi41 := 0; fi41 < 18; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            v35 = v35
            _ = fi42
        }
        fmt.Println(v39)
        _ = fi41
    }
    var v43 uint64 = 2
    fmt.Println(v2, v39, v43)
    if ((v39 / ((2147483647) | 1)) >= v39) {
        fmt.Println(v39, v2)
        var v44 uint64 = 412
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v43
}
