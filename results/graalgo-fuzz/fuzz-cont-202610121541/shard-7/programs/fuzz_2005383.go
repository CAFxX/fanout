package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = -2.25
    var v3 int = 127
    var v4 string = "\n\t"
    v5 := make(map[string]uint32)
    v6 := GenS1{F0: 399, F1: 0}
    v7 := GenS1{F0: -1, F1: 42}
    var v8 int32 = 2147483647
    fmt.Println(v8, v2, v3)
    v4 = "a"
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint32 = 0
        fmt.Println(v4, v3, v2)
        delete(v5, "foo bar")
        _ = fi9
        _ = v10
    }
    switch -9223372036854775808 {
    case 32767:
        v11 := GenS1{F0: 364, F1: 1}
        _ = v11
    default:
        var v12 uint16 = 2
        _ = v12
    }
    if (2 == 640) {
        switch (v4 + v4) {
        case "hello":
            fmt.Println(v2)
        case "\n\t":
            fmt.Println(v8, v3, v4)
        case "a":
            var v13 bool = ((!((565 == 762) && false)) || false)
            _ = v13
        default:
            fmt.Println(v8, v3)
        }
        if false {
            switch v4 {
            case "\n\t":
                v14 := GenS1{F0: (v8 | v8), F1: (-(v8))}
                _ = v14
            case "":
                v8 -= (v8 ^ (v8 - (v8 % (((-(v8))) | 1))))
            case "foo bar":
                switch v4 {
                case "a":
                    var v15 int = (v3 & v3)
                    _ = v15
                default:
                    v2 += v2
                }
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 float32 = v2
                for fi18 := 0; fi18 < 1; fi18++ {
                    if false {
                        var v19 uint64 = 4294967295
                        var v20 int16 = 32767
                        _ = v19
                        _ = v20
                    } else {
                        if true {
                            var v21 uint = 819
                            var v22 int16 = 0
                            _ = v21
                            _ = v22
                        } else {
                            v8 -= (v8 | v8)
                            if (!(!(fi16 != (fi16 / (((-(fi18))) | 1))))) {
                                fmt.Println(v8)
                                if (true && (65535 > 42)) {
                                    if false {
                                        v8 = v8
                                    }
                                    v23 := []int{v3, -1, 9223372036854775807}
                                    v8 = -874
                                    _ = v23
                                }
                                var v24 uint = 1
                                _ = v24
                            } else {
                                v5[("\n\t" + v4)] = 65535
                                v25 := GenS1{F0: v8, F1: 0}
                                _ = v25
                            }
                        }
                        var v26 uint64 = 18446744073709551615
                        _ = v26
                    }
                    v2 = ((v2 / 1.5) + (v17 + (1e10 + v17)))
                    fmt.Println(v8)
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        } else {
            for fi27 := 0; fi27 < 10; fi27++ {
                if (42 <= 32767) {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if (((v4 + v4) + v4) == "\n\t") {
                            if true {
                                var v29 float32 = 0.1
                                _ = v29
                            } else {
                                fmt.Println(fi28, v3, v4)
                                v2 += -2.25
                            }
                            v30 := []int{}
                            _ = v30
                        } else {
                            var v31 uint = 42
                            v32 := GenS2{F0: 42}
                            v33 := []int{(v3 | (v3 >> ((18446744073709551615) & 63))), (fi27 | fi28)}
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        if false {
                            _ = len(v5)
                        }
                        _ = fi28
                    }
                } else {
                    if (!(v8 > v8)) {
                        v34 := GenS2{F0: 576}
                        v3 = fi27
                        v3 = 42
                        _ = v34
                    }
                }
                _ = fi27
            }
            var v35 int16 = 1
            _ = v35
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            if true {
                var v37 string = (("a" + v4) + v4)
                var v38 int = fi36
                _ = v37
                _ = v38
            }
            v39 := GenS1{F0: (v8 << ((2) & 63)), F1: (v8 ^ 242)}
            _ = fi36
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        var v41 int8 = 1
        _ = fi40
        _ = v41
    }
    fmt.Println(v3, v3)
    v42 := GenS3{F0: 42, F1: 42}
    var v43 uint64 = 0
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 13; fi45++ {
            delete(v5, (v4 + v4))
            _ = fi45
        }
        _ = fi44
    }
    var v46 int64 = -558
    v47 := []int{v3, -1, 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
