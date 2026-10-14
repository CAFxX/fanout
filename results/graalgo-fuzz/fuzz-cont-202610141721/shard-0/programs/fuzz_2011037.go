package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
func genFn0(p0 uint) uint16 {
    if true {
        p0 = (p0 >> ((p0) & 63))
    } else {
        var v7 int16 = 976
        _ = v7
    }
    if true {
        var v8 bool = (0 < int16(p0))
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint = p0
            _ = fi9
            _ = v10
        }
        var v11 int = 127
        _ = v11
        _ = v8
    }
    switch "foo bar" {
    case "foo bar":
        var v12 uint64 = 2
        _ = v12
    default:
        if (1.5 != 1.5) {
            var v13 uint16 = 2
            _ = v13
        } else {
            v14 := GenS1{F0: 128, F1: 99, F2: 171}
            v15 := []int{}
            if (float64(0.1) == (0.0 + 1.5)) {
                for fi16 := 0; fi16 < 2; fi16++ {
                    switch (-(fi16)) {
                    case -9223372036854775808:
                        var v17 int8 = 127
                        _ = v17
                    case 128:
                        p0 += p0
                    default:
                        for fi18 := 0; fi18 < 8; fi18++ {
                            var v19 int16 = 104
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = fi16
                }
            } else {
                if ("a" != "foo bar") {
                    if len(v15) > 0 {
                        v15[0] = 9223372036854775807
                    }
                    var v20 uint16 = 299
                    _ = v20
                }
                if true {
                    switch 128 {
                    case 2147483647:
                        if (1 < 463) {
                            var v21 int16 = 1
                            _ = v21
                        } else {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                v23 := GenS1{F0: fi22, F1: -51, F2: -9223372036854775808}
                                _ = fi22
                                _ = v23
                            }
                            switch ((("foo bar" + "foo bar") + "foo bar") + "foo bar") {
                            case "hello":
                                p0 -= p0
                            case "":
                                v24 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: 0}
                                _ = v24
                            case "\n\t":
                                p0 -= p0
                            }
                        }
                    case -67:
                        var v25 int8 = -1
                        _ = v25
                    default:
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                    var v26 int64 = -651
                    for fi27 := 0; fi27 < 6; fi27++ {
                        for fi28 := 0; fi28 < 15; fi28++ {
                            v29 := GenS2{F0: 2147483648, F1: 980, F2: 0}
                            fmt.Println(fi28, v26, fi27)
                            if (true && false) {
                                p0 = p0
                                v26 = v26
                            }
                            _ = fi28
                            _ = v29
                        }
                        p0 += p0
                        fmt.Println(v26)
                        _ = fi27
                    }
                    _ = v26
                }
            }
            _ = v14
            _ = v15
        }
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{0}
    v2 := make(map[int]string)
    var v3 uint32 = 317
    v4 := make(map[uint32]string)
    v5 := []int{}
    v6 := make(map[int]string)
    fmt.Println(v3)
    v3 += 2147483647
    if (2 < 1) {
        switch int(v3) {
        case 42:
            if false {
                var v30 int32 = 0
                switch "hello" {
                case "":
                    var v31 float64 = 1.5
                    _ = v31
                case "a":
                    fmt.Println(v3, v3)
                case "\n\t":
                    v32 := GenS1{F0: 0, F1: 42, F2: 1}
                    _ = v32
                default:
                    v30 = (v30 >> ((191) & 63))
                }
                if len(v5) > 0 {
                    v5[0] = 9223372036854775807
                }
                _ = v30
            } else {
                v33 := GenS1{F0: 4294967295, F1: -2147483648, F2: 108}
                _ = v33
            }
        case 9223372036854775807:
            v3 = (^(v3))
        default:
            v3 += v3
        }
        v34 := GenS1{F0: 32768, F1: 647, F2: -374}
        var v35 uint16 = 1
        _ = v34
        _ = v35
    } else {
        v36 := GenS1{F0: 9223372036854775807, F1: 42, F2: 627}
        _ = v36
    }
    var v37 int64 = 9223372036854775807
    fmt.Println(v37)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v38 uint = 65535
    var v39 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
