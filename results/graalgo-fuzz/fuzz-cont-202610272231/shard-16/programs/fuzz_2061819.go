package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint) uint16 {
    var v4 int32 = 0
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 float32 = 1e10
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    var v8 int = 1
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    switch "foo bar" {
    case "hello":
        if false {
            v9 := GenS1{F0: 127, F1: 128, F2: (^(v1))}
            v3 = (((v3 + v3) + ("\n\t" + "foo bar")) + v3)
            _ = v9
        }
    case "\n\t":
        v10 := []int{4294967295}
        _ = v10
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        fmt.Println(v3, v1)
    }
    if (!true) {
        var v11 string = ((v3 + "foo bar") + v3)
        v3 = v3
        _ = v11
    }
    v1 = v1
    v12 := GenS1{F0: -128, F1: -515, F2: v1}
    v13 := v2[0]
    v3 = v3
    v13 = (^(v13))
    if (2147483647 < 2147483647) {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 float32 = 1.5
            v16 := GenS1{F0: 1, F1: (fi14 * fi14), F2: (v1 - v1)}
            _ = fi14
            _ = v15
            _ = v16
        }
        if true {
            var v17 bool = true
            _ = v17
        }
    }
    v3 = v3
    fmt.Println(v13)
    v3 = "foo bar"
    var v18 bool = (false || true)
    fmt.Println(v3)
    for fi19 := 0; fi19 < 2; fi19++ {
        switch fi19 {
        case -1:
            fmt.Println(v1)
        case 128:
            for fi20 := 0; fi20 < 19; fi20++ {
                v3 = v3
                var v21 int64 = 42
                _ = fi20
                _ = v21
            }
        case 517:
            if false {
                var v22 bool = ((734 != 794) && (1000 > -638))
                if v22 {
                    _ = len(v2)
                }
                v23 := []int{128}
                _ = v22
                _ = v23
            } else {
                if v18 {
                    var v24 int = -1
                    if v18 {
                        fmt.Println(v24)
                    }
                    var v25 uint = 575
                    _ = v24
                    _ = v25
                } else {
                    delete(v2, (^(v13)))
                    v13 = (v13 | v13)
                }
                if true {
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 uint8 = 241
                        var v28 uint8 = (^(v27))
                        v29 := GenS1{F0: -117, F1: fi26, F2: v1}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    if (false && true) {
                        var v30 uint = 65535
                        v18 = v18
                        _ = v30
                    }
                    var v31 int32 = 64
                    _ = v31
                } else {
                    v1 = v1
                    var v32 string = (((("\n\t" + v3) + v3) + ((v3 + "foo bar") + v3)) + v3)
                    _ = v32
                }
            }
        default:
            v3 = ""
        }
        var v33 uint = 42
        if (!(42 > 1)) {
            switch "hello" {
            case "foo bar":
                var v34 uint16 = 65535
                _ = v34
            }
            var v35 uint8 = 128
            switch v3 {
            case "foo bar":
                for fi36 := 0; fi36 < 2; fi36++ {
                    switch v3 {
                    case "foo bar":
                        fmt.Println(v3, fi19, v35)
                    case "":
                        switch -423 {
                        case 1:
                            v37 := GenS2{F0: v13}
                            _ = v37
                        case 562:
                            v38 := GenS1{F0: 127, F1: fi36, F2: (v1 * 18446744073709551615)}
                            _ = v38
                        case -296:
                            fmt.Println(fi19)
                        }
                    default:
                        fmt.Println(v13, v18)
                    }
                    if (!true) {
                        fmt.Println(v13, v1)
                        v33 -= 65535
                    } else {
                        v39 := GenS2{F0: v13}
                        fmt.Println(v3, v35, v1)
                        _ = v39
                    }
                    switch -2147483648 {
                    case -9223372036854775808:
                        var v40 float32 = 1.5
                        _ = v40
                    case 42:
                        v41 := GenS3{F0: (v33 % ((v33) | 1))}
                        _ = v41
                    case -526:
                        v18 = (32767 < 65535)
                    default:
                        switch v3 {
                        case "foo bar":
                            v3 = (v3 + "a")
                        default:
                            fmt.Println(v35, v35)
                        }
                    }
                    _ = fi36
                }
            }
            _ = v35
        }
        _ = fi19
        _ = v33
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        switch "\n\t" {
        case "\n\t":
            v18 = (810 > 42)
        }
        _ = fi42
    }
    var v43 float64 = float64(0.0)
    if true {
        v44 := GenS3{F0: 177}
        for fi45 := 0; fi45 < 16; fi45++ {
            v43 += ((v43 / v43) - 0.0)
            for fi46 := 0; fi46 < 20; fi46++ {
                v3 = "hello"
                _ = fi46
            }
            _ = fi45
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v43
}
