package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 string }
func genFn0(p0 int64, p1 uint32, p2 int16) int {
    fmt.Println(p0)
    if true {
        fmt.Println(p1)
        v5 := []int{2147483647, -9223372036854775808, 32768, 32768}
        for fi6 := 0; fi6 < 4; fi6++ {
            switch (fi6 % ((fi6) | 1)) {
            case 0:
                var v7 int32 = -2147483648
                _ = v7
            case 127:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v8 := GenS1{F0: 255}
            _ = len(v5)
            _ = fi6
            _ = v8
        }
        _ = v5
    } else {
        v9 := GenS1{F0: 255}
        p1 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -686
}
func genFn1(p0 bool, p1 int32, p2 uint8) uint32 {
    fmt.Println(p2, p2, p1)
    var v10 bool = ((p2 ^ (p2 << ((255) & 63))) < (p2 * 255))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 393
}

func main() {
    var v1 uint = 255
    v2 := []string{}
    v3 := make(map[int]int)
    var v4 int64 = 42
    if (false && (true || false)) {
        if ("a" < "") {
            v11 := GenS1{F0: (^(v1))}
            v4 = 179
            var v12 uint32 = 2
            _ = v11
            _ = v12
        }
        if (693 <= 1) {
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v4 -= -1
                    _ = len(v2)
                    fmt.Println(fi14, fi14)
                    _ = fi14
                }
                var v15 uint = v1
                if (!true) {
                    if false {
                        v3[32768] = (-(fi13))
                        if false {
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                var v16 int32 = -63
                                _ = v16
                            case "hello":
                                v15 += v1
                            default:
                                delete(v3, fi13)
                            }
                            switch fi13 {
                            case 2147483647:
                                switch fi13 {
                                case 956:
                                    switch (fi13 >> (((v15 % ((v1) | 1))) & 63)) {
                                    case 42:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 32768:
                                        v4 = (v4 & v4)
                                    case -1:
                                        var v17 int8 = 100
                                        _ = v17
                                    default:
                                        v15 = v15
                                    }
                                default:
                                    var v18 uint = v15
                                    _ = v18
                                }
                            case 128:
                                for fi19 := 0; fi19 < 7; fi19++ {
                                    if true {
                                        var v20 int32 = 888
                                        v21 := GenS2{F0: 2, F1: (fi19 | (-(fi19)))}
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v22 uint8 = 128
                                            _ = v22
                                        case "":
                                            v23 := GenS1{F0: v1}
                                            _ = v23
                                        }
                                        _ = v20
                                        _ = v21
                                    }
                                    var v24 bool = ("a" < ((((("hello" + "a") + ("" + "\n\t")) + "") + "\n\t") + ("a" + "")))
                                    _ = fi19
                                    _ = v24
                                }
                            }
                        }
                    }
                    for fi25 := 0; fi25 < 16; fi25++ {
                        v4 = v4
                        var v26 int = 0
                        var v27 uint8 = 128
                        v2 = append(v2, (((("\n\t" + "\n\t") + ("foo bar" + "")) + ("\n\t" + "\n\t")) + "\n\t"))
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = ("\n\t" + ("\n\t" + "\n\t"))
                    }
                    var v28 int = (fi13 ^ -9223372036854775808)
                    _ = v28
                }
                _ = fi13
                _ = v15
            }
            var v29 int8 = 1
            if true {
                v30 := GenS3{F0: 1, F1: 255}
                _ = v30
            }
            _ = v29
        } else {
            fmt.Println(v1, v1)
            var v31 int8 = 0
            v1 = 65535
            _ = v31
        }
    }
    v4 += (v4 >> ((v1) & 63))
    v32 := GenS1{F0: 65535}
    if len(v2) > 0 {
        v2[0] = ""
    }
    v33 := GenS4{F0: 1, F1: (v4 & (^(v4))), F2: 423}
    var v34 int16 = -1
    fmt.Println(v4)
    var v35 int32 = 42
    fmt.Println(v1)
    v36 := GenS3{F0: 2, F1: 127}
    var v37 uint = (^(v1))
    v1 = (v37 << ((146) & 63))
    switch ("foo bar" + ("\n\t" + ("a" + (("a" + "hello") + "hello")))) {
    case "":
        var v38 string = (("" + ("\n\t" + "hello")) + "\n\t")
        _ = v38
    }
    for fi39 := 0; fi39 < 2; fi39++ {
        v40 := GenS5{F0: 403, F1: "\n\t"}
        _ = fi39
        _ = v40
    }
    v41 := v3[331]
    switch (v41 % ((genFn0(v4, 447, (v34 + (^(v34))))) | 1)) {
    case 127:
        for fi42 := 0; fi42 < 13; fi42++ {
            v1 = (v1 - 2)
            _ = fi42
        }
    case -938:
        fmt.Println(v37)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
