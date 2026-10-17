package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint; F2 int32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 bool) int16 {
    p1 = (42 == 1000)
    if p1 {
        if ((^(p0)) < 9223372036854775807) {
            p1 = (!false)
            p0 -= p0
            var v4 float64 = 0.1
            _ = v4
        } else {
            var v5 int = p0
            var v6 int64 = 101
            v5 += p0
            _ = v5
            _ = v6
        }
        v7 := []int{p0, p0, (p0 << ((0) & 63)), p0}
        _ = v7
    } else {
        p0 = 2147483647
        var v8 string = "hello"
        _ = v8
    }
    var v9 uint16 = 2
    _ = p0
    _ = p1
    _ = v9
    return 32767
}
func genFn1(p0 uint) int8 {
    var v10 string = ""
    v11 := GenS1{F0: 1}
    switch 4294967295 {
    case 401:
        fmt.Println(p0, p0)
    case 42:
        var v12 uint = (p0 + p0)
        _ = v12
    case -9223372036854775808:
        v13 := GenS2{F0: 486}
        _ = v13
    default:
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS3{F0: fi14, F1: (^(p0)), F2: 0}
            _ = fi14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := []int{fi16}
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v10
    _ = v11
    return 14
}

func main() {
    v1 := []int{655, 42}
    v2 := make(map[uint32]string)
    var v3 int32 = -2147483648
    v18 := GenS3{F0: 1, F1: 65535, F2: v3}
    var v19 uint32 = 1
    var v20 uint8 = 4
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 int = -1
        fmt.Println(v22)
        for fi23 := 0; fi23 < 13; fi23++ {
            _ = len(v1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    v24 := GenS2{F0: 0}
    switch -2147483648 {
    case -293:
        if (-2147483648 >= (^(v3))) {
            var v25 int64 = 9223372036854775807
            v26 := GenS3{F0: 32767, F1: 42, F2: v3}
            v1 = append(v1, 1)
            _ = v25
            _ = v26
        }
    case -395:
        fmt.Println(v3, v19)
    default:
        switch 2147483647 {
        case -2147483648:
            v3 = v3
        case -232:
            switch "hello" {
            case "foo bar":
                switch int(v19) {
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    switch 2147483647 {
                    case -2147483648:
                        var v27 int8 = -1
                        _ = v27
                    case 0:
                        var v28 int64 = -562
                        _ = v28
                    case 127:
                        if false {
                            var v29 int32 = v3
                            for fi30 := 0; fi30 < 9; fi30++ {
                                v31 := GenS2{F0: 708}
                                _ = fi30
                                _ = v31
                            }
                            _ = v29
                        } else {
                            var v32 int8 = 127
                            _ = v32
                        }
                    }
                }
            case "\n\t":
                if true {
                    v19 = v19
                    var v33 bool = (((("hello" + "\n\t") + "\n\t") < "hello") && true)
                    v34 := GenS4{F0: 662, F1: 81}
                    _ = v33
                    _ = v34
                }
            case "a":
                if (!(1 <= 42)) {
                    fmt.Println(v3, v3)
                } else {
                    fmt.Println(v3, v3)
                    v20 += 1
                    v20 -= v20
                }
            }
        default:
            v1 = append(v1, 2147483647)
        }
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        _ = len(v2)
        fmt.Println(v3, v3, fi35)
        for fi36 := 0; fi36 < 10; fi36++ {
            var v37 string = "a"
            v38 := GenS4{F0: 2, F1: 1}
            var v39 uint64 = 56
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = fi35
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        if (true || true) {
            v41 := GenS3{F0: -195, F1: 18446744073709551615, F2: v3}
            v3 = 42
            var v42 int64 = 0
            _ = v41
            _ = v42
        }
        switch (fi40 >> ((4294967295) & 63)) {
        case 2147483647:
            var v43 int32 = 42
            _ = v43
        case -1:
            fmt.Println(v19)
        default:
            var v44 uint16 = 65535
            _ = v44
        }
        _ = fi40
    }
    var v45 int16 = -32768
    v3 = -49
    if (!false) {
        if false {
            v20 += v20
            v45 -= (v45 | (v45 >> ((4294967295) & 63)))
        } else {
            v45 += v45
            v46 := v2[v19]
            v47 := GenS5{F0: -12}
            _ = v46
            _ = v47
        }
        _ = len(v1)
    }
    fmt.Println(v19)
    fmt.Println(v20)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v45
}
