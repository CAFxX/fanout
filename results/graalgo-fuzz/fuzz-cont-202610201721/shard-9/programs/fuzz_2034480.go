package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "\n\t"
    var v3 uint16 = 1
    var v4 uint32 = 2147483647
    fmt.Println(v2, v2, v3)
    var v5 string = ""
    var v6 uint32 = 2147483647
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            if false {
                delete(v1, "")
                fmt.Println(v6)
            }
            _ = fi7
        }
    } else {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{(^(fi8))}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 string = "a"
            delete(v1, "")
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 100, F1: 4294967295, F2: 104}
        _ = v12
    }
    var v13 uint16 = (v3 | v3)
    v14 := GenS1{F0: 100, F1: (v4 - v6), F2: 2}
    var v15 int16 = -32768
    v2 = "a"
    if (!true) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v4 += 65535
            var v17 int8 = 30
            switch "a" {
            case "":
                v18 := []int{(-(fi16)), 42, fi16, fi16}
                _ = v18
            case "a":
                if true {
                    var v19 int32 = 42
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 string = "hello"
                        if (float64(-2.25) != float64(-2.25)) {
                            v21 = "foo bar"
                            if false {
                                v2 = "\n\t"
                                v5 = v5
                            }
                            if (!(!(false || true))) {
                                v22 := GenS2{F0: (v19 & 2147483647), F1: (^(fi16)), F2: 1}
                                if (185 < 0) {
                                    var v23 int16 = -1
                                    _ = v23
                                }
                                _ = v22
                            } else {
                                var v24 uint64 = 42
                                fmt.Println(v21, v2)
                                v17 -= 105
                                _ = v24
                            }
                        } else {
                            v25 := GenS1{F0: (v17 % ((v17) | 1)), F1: 2147483647, F2: 607}
                            var v26 int64 = -9223372036854775808
                            _ = v25
                            _ = v26
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v13, v4, v2)
        var v27 string = v2
        _ = v27
    }
    if (false && (0.1 > 1e10)) {
        fmt.Println(v2, v4)
        if true {
            var v28 float64 = 1.5
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 string = v5
                _ = fi29
                _ = v30
            }
            var v31 int8 = 0
            _ = v28
            _ = v31
        }
    } else {
        switch -558 {
        case -2147483648:
            if true {
                for fi32 := 0; fi32 < 16; fi32++ {
                    fmt.Println(v5)
                    var v33 int = -427
                    _ = fi32
                    _ = v33
                }
            } else {
                v1[v5] = v4
                fmt.Println(v6, v4, v3)
                v34 := GenS1{F0: -128, F1: (^(v4)), F2: 1}
                _ = v34
            }
        case 128:
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 bool = (33 > 2)
                delete(v1, v2)
                v3 = (v13 >> ((0) & 63))
                _ = fi35
                _ = v36
            }
        case 25:
            v37 := v1[v2]
            _ = v37
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            var v39 string = "hello"
            _ = fi38
            _ = v39
        }
    }
    if (!true) {
        var v40 uint64 = 42
        v41 := []int{2147483647, 0}
        _ = v40
        _ = v41
    } else {
        v42 := GenS2{F0: 2147483647, F1: 64, F2: 9223372036854775807}
        v43 := GenS1{F0: -9, F1: v6, F2: 779}
        _ = v42
        _ = v43
    }
    var v44 int = 46
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
