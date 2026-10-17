package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 string; F2 int8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool, p1 int16) uint32 {
    var v7 string = "\n\t"
    var v8 bool = (42 >= -1)
    v9 := []int{-648}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 bool, p1 string) int64 {
    var v10 int32 = -71
    var v11 uint = 65535
    var v12 int32 = (v10 * -1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 string = "a"
    var v2 int8 = -1
    v3 := []string{v1, "hello", v1}
    var v4 int = 128
    v5 := make(map[string]int)
    v6 := make(map[int]uint32)
    var v13 uint64 = 2
    if (!false) {
        var v14 uint32 = genFn0(false, -589)
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS1{F0: v13, F1: 9223372036854775807, F2: 85}
        fmt.Println(v13, v1)
        v17 := GenS1{F0: (v13 ^ (v13 >> ((0) & 63))), F1: int64(fi15), F2: 825}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float32 = 1e10
    switch v4 {
    case 434:
        fmt.Println(v2, v13, v18)
    }
    v19 := GenS1{F0: 4294967295, F1: 1, F2: 2}
    if (v13 >= v13) {
        v1 = v1
        switch -596 {
        case -1:
            var v20 int = v4
            _ = v20
        }
    } else {
        var v21 uint16 = 42
        _ = len(v3)
        switch (v1 + ("" + "\n\t")) {
        case "a":
            v22 := GenS2{F0: genFn1((!true), "")}
            _ = v22
        }
        _ = v21
    }
    if ((v2 ^ -128) >= v2) {
        if true {
            v4 = v4
        } else {
            if true {
                fmt.Println(v18, v4)
                v18 -= v18
            } else {
                if false {
                    var v23 uint = 4294967295
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        if (!true) {
                            v1 = ("foo bar" + v1)
                            v25 := GenS3{F0: 32768, F1: (v1 + v1), F2: v2}
                            for fi26 := 0; fi26 < 19; fi26++ {
                                v27 := GenS2{F0: 1}
                                v1 = "hello"
                                _ = fi26
                                _ = v27
                            }
                            _ = v25
                        }
                        v28 := v5["\n\t"]
                        _ = fi24
                        _ = v28
                    }
                }
                if false {
                    if len(v3) > 0 {
                        v3[0] = v1
                    }
                } else {
                    v13 = v13
                    fmt.Println(v18, v2, v1)
                }
                v13 = v13
            }
            var v29 int8 = v2
            _ = v29
        }
        var v30 float32 = v18
        v31 := GenS2{F0: 1}
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            var v33 int64 = 1
            _ = fi32
            _ = v33
        }
        _ = len(v5)
    }
    v1 = (v1 + ("\n\t" + ((v1 + "") + "hello")))
    v1 = ""
    if (false && false) {
        for fi34 := 0; fi34 < 11; fi34++ {
            if (32767 > 65535) {
                v13 = 42
            } else {
                if false {
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v18 = 0.1
                        v36 := GenS1{F0: 148, F1: 459, F2: 2}
                        _ = fi35
                        _ = v36
                    }
                } else {
                    if (false || false) {
                        for fi37 := 0; fi37 < 8; fi37++ {
                            if (128 >= 2) {
                                switch v1 {
                                case "":
                                    v18 -= v18
                                case "\n\t":
                                    var v38 float64 = 3.14159
                                    _ = v38
                                case "hello":
                                    v39 := GenS1{F0: v13, F1: -1, F2: 1}
                                    _ = v39
                                default:
                                    v3 = append(v3, (("foo bar" + v1) + v1))
                                    if true {
                                        switch "hello" {
                                        case "foo bar":
                                            v18 = (float32(-2.25) + v18)
                                        case "a":
                                            var v40 int16 = -1
                                            _ = v40
                                        case "":
                                            v3 = append(v3, "\n\t")
                                        }
                                        fmt.Println(v18, v13, v18)
                                        v4 -= v4
                                    } else {
                                        v13 -= (v13 + v13)
                                        fmt.Println(v13)
                                        var v41 int = fi37
                                        v42 := GenS2{F0: 42}
                                        var v43 string = (("a" + v1) + ("" + v1))
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    }
                                    v44 := GenS1{F0: v13, F1: 42, F2: 1}
                                    if true {
                                        switch ((v1 + v1) + v1) {
                                        case "":
                                            v13 = (v13 << ((926) & 63))
                                        }
                                        var v45 string = v1
                                        _ = v45
                                    } else {
                                        if (!false) {
                                            delete(v5, v1)
                                        }
                                        v46 := GenS1{F0: v13, F1: -9223372036854775808, F2: 883}
                                        v2 = -1
                                        _ = v46
                                    }
                                    if false {
                                        if false {
                                            v47 := GenS2{F0: -9223372036854775808}
                                            _ = v47
                                        } else {
                                            var v48 uint16 = 0
                                            if len(v3) > 0 {
                                                v3[0] = v1
                                            }
                                            _ = v48
                                        }
                                    }
                                    _ = v44
                                }
                                if (42 > 32) {
                                    var v49 uint64 = v13
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: 0}
                                    _ = v50
                                }
                                v4 += 1
                            } else {
                                var v51 bool = (v1 != v1)
                                _ = v51
                            }
                            _ = fi37
                        }
                        v1 = "\n\t"
                        _ = len(v3)
                    }
                }
            }
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            for fi52 := 0; fi52 < 18; fi52++ {
                v4 += fi52
                fmt.Println(v13)
                var v53 float64 = (1.5 / -2.25)
                _ = fi52
                _ = v53
            }
            _ = fi34
        }
    }
    v13 += (v13 % ((4294967295) | 1))
    v13 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
