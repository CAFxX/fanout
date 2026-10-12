package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int64, p1 int16) int32 {
    p0 = 0
    var v4 uint64 = 659
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v1 = (v1 + (v1 | (v1 % (((^(v1))) | 1))))
    if false {
        for fi5 := 0; fi5 < 16; fi5++ {
            if false {
                v6 := GenS1{F0: -774, F1: -1}
                if (true && (((("a" + ("\n\t" + "\n\t")) + ("\n\t" + "\n\t")) + (((("" + "") + "foo bar") + (("hello" + "hello") + "foo bar")) + "a")) >= "")) {
                    var v7 uint8 = 255
                    v8 := GenS2{F0: v7, F1: 0, F2: 1}
                    v7 -= v7
                    _ = v7
                    _ = v8
                } else {
                    v1 = v1
                    v9 := []int{(fi5 ^ 1), 2147483647, fi5}
                    switch "foo bar" {
                    case "":
                        fmt.Println(v1)
                    default:
                        if (true && ("foo bar" >= ("foo bar" + ""))) {
                            var v10 uint8 = 255
                            v1 -= (v1 ^ v1)
                            for fi11 := 0; fi11 < 6; fi11++ {
                                if ((true || ("\n\t" == "")) && (genFn0(952, -599) == 42)) {
                                    switch (^(fi11)) {
                                    case -9223372036854775808:
                                        switch 990 {
                                        case -384:
                                            var v12 bool = true
                                            _ = v12
                                        case -9223372036854775808:
                                            v13 := GenS3{F0: 2, F1: -2147483648}
                                            _ = v13
                                        case 738:
                                            if false {
                                                var v14 uint64 = 42
                                                var v15 uint8 = (v10 & 22)
                                                v1 -= (v14 + (v1 % ((v14) | 1)))
                                                _ = v14
                                                _ = v15
                                            } else {
                                                var v16 int32 = 1
                                                var v17 uint8 = (v10 | (v10 / (((v10 % ((v10) | 1))) | 1)))
                                                _ = v16
                                                _ = v17
                                            }
                                        }
                                    case -469:
                                        switch 91 {
                                        case -9223372036854775808:
                                            var v18 bool = false
                                            _ = v18
                                        }
                                    }
                                    if false {
                                        v1 += (v1 - (^(v1)))
                                    } else {
                                        v19 := GenS2{F0: v10, F1: 42, F2: 0}
                                        _ = v19
                                    }
                                    _ = len(v2)
                                }
                                var v20 uint = 348
                                v1 = (^(v1))
                                _ = fi11
                                _ = v20
                            }
                            _ = v10
                        } else {
                            var v21 uint32 = 2
                            _ = v21
                        }
                    }
                    _ = v9
                }
                _ = v6
            } else {
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(fi5)
                    v23 := GenS1{F0: 73, F1: -89}
                    _ = fi22
                    _ = v23
                }
            }
            v24 := GenS2{F0: 1, F1: 0, F2: -70}
            if (!(518 == 65535)) {
                var v25 bool = false
                v25 = v25
                _ = v25
            } else {
                var v26 bool = (fi5 > fi5)
                _ = v26
            }
            _ = fi5
            _ = v24
        }
    }
    v27 := []int{2147483647, -7}
    for fi28 := 0; fi28 < 16; fi28++ {
        for fi29 := 0; fi29 < 17; fi29++ {
            if (32767 > 1000) {
                v1 -= 4294967295
            }
            fmt.Println(v1, fi28)
            _ = fi29
        }
        _ = fi28
    }
    fmt.Println(v1)
    v1 -= v1
    var v30 int64 = -627
    switch "foo bar" {
    case "":
        v30 -= (v30 << ((184) & 63))
    case "foo bar":
        _ = len(v2)
    case "a":
        switch 127 {
        case 32768:
            fmt.Println(v30)
        case 4294967295:
            if true {
                fmt.Println(v30)
                switch (("hello" + "a") + ((((("foo bar" + "\n\t") + ("" + "")) + ("a" + (("hello" + "\n\t") + "foo bar"))) + "a") + "")) {
                case "hello":
                    var v31 int16 = 774
                    _ = v31
                default:
                    var v32 uint16 = 32767
                    _ = v32
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    v1 -= (v1 / ((0) | 1))
                    for fi34 := 0; fi34 < 13; fi34++ {
                        if true {
                            if (false && true) {
                                v1 += v1
                            }
                            delete(v2, "foo bar")
                        }
                        _ = fi34
                    }
                    var v35 float64 = -2.25
                    _ = fi33
                    _ = v35
                }
            }
        case -153:
            var v36 int8 = 127
            _ = v36
        default:
            v1 = (^(v1))
        }
    }
    v37 := GenS3{F0: 316, F1: -2147483648}
    if len(v27) > 0 {
        v27[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v37
}
