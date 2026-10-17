package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
func genFn0(p0 string, p1 int16) uint64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{(fi7 % ((fi6) | 1)), fi6, fi6, (fi7 * fi6)}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{}
            p0 = (("foo bar" + p0) + p0)
            if len(v10) > 0 {
                v10[0] = fi6
            }
            _ = fi9
            _ = v10
        }
        switch p0 {
        case "a":
            var v11 uint16 = 0
            _ = v11
        default:
            v12 := []int{-770, fi6, fi6, fi6}
            _ = v12
        }
        _ = fi6
    }
    if (false || (0 != 42)) {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{-2147483648, fi13}
            if len(v14) > 0 {
                v14[0] = (fi13 - 9223372036854775807)
            }
            p1 = -853
            _ = fi13
            _ = v14
        }
        switch "a" {
        case "hello":
            if ((183 == 234) && true) {
                if (!false) {
                    var v15 int64 = -9223372036854775808
                    var v16 int16 = p1
                    _ = v15
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        p1 = 1000
                        if false {
                            if false {
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    if false {
                                        switch fi17 {
                                        case 127:
                                            var v19 int64 = -1
                                            _ = v19
                                        case 2147483647:
                                            var v20 int = 127
                                            _ = v20
                                        case 0:
                                            var v21 int64 = 0
                                            _ = v21
                                        }
                                    }
                                    var v22 int = (fi18 ^ 0)
                                    fmt.Println(p0)
                                    var v23 int64 = -153
                                    p0 = "foo bar"
                                    _ = fi18
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                v24 := GenS1{F0: 0, F1: "\n\t", F2: -1}
                                var v25 uint8 = 232
                                _ = v24
                                _ = v25
                            }
                            v26 := []int{fi17, 3, (fi17 & (-(fi17)))}
                            _ = v26
                        }
                        _ = fi17
                    }
                    var v27 string = "foo bar"
                    _ = v27
                }
            } else {
                if (-9223372036854775808 >= 0) {
                    p1 = 76
                }
                for fi28 := 0; fi28 < 4; fi28++ {
                    if true {
                        var v29 bool = (!(255 < 127))
                        _ = v29
                    } else {
                        if (1 > 238) {
                            var v30 string = "a"
                            var v31 float32 = 0.1
                            for fi32 := 0; fi32 < 7; fi32++ {
                                var v33 int32 = -1
                                _ = fi32
                                _ = v33
                            }
                            _ = v30
                            _ = v31
                        } else {
                            v34 := []int{fi28, fi28}
                            _ = v34
                        }
                        p0 = (((p0 + ((p0 + p0) + p0)) + (p0 + "\n\t")) + p0)
                    }
                    _ = fi28
                }
                switch -9223372036854775808 {
                case 32768:
                    p1 += p1
                }
            }
        }
    }
    p1 = 828
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int = -1
    v2 := []string{"a", ("foo bar" + "hello"), "foo bar"}
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    v5 := []uint32{}
    for fi35 := 0; fi35 < 6; fi35++ {
        v4["a"] = v1
        fmt.Println(v1)
        if true {
            v1 = -614
            for fi36 := 0; fi36 < 15; fi36++ {
                if (!(308 >= 4294967295)) {
                    v1 = -400
                    var v37 bool = (127 > 0)
                    _ = v37
                }
                _ = fi36
            }
        }
        _ = fi35
    }
    v1 = (v1 + -2147483648)
    v1 += (v1 * v1)
    v38 := GenS2{F0: -123, F1: -429}
    switch (v1 % ((v1) | 1)) {
    case 127:
        v1 -= v1
    case 9223372036854775807:
        switch "" {
        case "a":
            for fi39 := 0; fi39 < 1; fi39++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi39
            }
        }
    case -924:
        if true {
            v1 = v1
        }
    }
    var v40 float32 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v41 uint32 = 406
        _ = v41
    } else {
        var v42 int8 = -1
        _ = v42
    }
    var v43 uint8 = 128
    v1 += (v1 << ((425) & 63))
    switch "a" {
    case "\n\t":
        v40 -= v40
    case "hello":
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v43
    _ = v5
}
