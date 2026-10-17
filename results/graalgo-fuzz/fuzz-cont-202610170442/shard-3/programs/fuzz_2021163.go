package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }
func genFn0(p0 bool, p1 int8) int16 {
    var v4 int8 = p1
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int = 0
        v7 := GenS1{F0: 128}
        v8 := GenS1{F0: 128}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint16 = 546
        switch fi9 {
        case 1:
            var v11 uint32 = 629
            _ = v11
        default:
            var v12 bool = ((^(p1)) > 42)
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    if p0 {
        var v13 int = 2147483647
        v14 := GenS1{F0: 42}
        v15 := []int{-573}
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int = 2147483647
        var v17 int32 = -1
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v4
    return -1
}
func genFn1(p0 int64, p1 int64, p2 uint32) int64 {
    if (42 > 29) {
        v18 := GenS1{F0: 127}
        v19 := []int{618, 0}
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: 1}
        var v21 float32 = 1e10
        v22 := GenS1{F0: 128}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: 1}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 0
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v24 := []int{32768, -9223372036854775808, 32768, 971}
    fmt.Println(v1)
    switch -303 {
    case -58:
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 bool = false
            v26 = (fi25 == fi25)
            if v26 {
                var v27 uint = 807
                _ = v27
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    fmt.Println(v1)
                    if v26 {
                        var v29 int = (-(fi28))
                        v30 := GenS2{F0: 14, F1: 1, F2: ("hello" + (("a" + "a") + ("" + "\n\t")))}
                        _ = v29
                        _ = v30
                    }
                    var v31 int16 = 607
                    _ = fi28
                    _ = v31
                }
            }
            _ = fi25
            _ = v26
        }
    }
    var v32 string = "a"
    v1 = (v1 ^ 120)
    switch ("hello" + v32) {
    case "":
        var v33 uint = 2
        _ = v33
    case "a":
        var v34 string = (v32 + v32)
        _ = v34
    case "hello":
        switch 261 {
        case 1:
            switch v32 {
            case "\n\t":
                for fi35 := 0; fi35 < 3; fi35++ {
                    for fi36 := 0; fi36 < 11; fi36++ {
                        if true {
                            var v37 int8 = v1
                            _ = v37
                        }
                        v1 += v1
                        _ = fi36
                    }
                    for fi38 := 0; fi38 < 5; fi38++ {
                        if false {
                            v32 = (v32 + "a")
                            v32 = ""
                        } else {
                            var v39 int64 = 42
                            _ = len(v24)
                            var v40 float32 = 1.5
                            _ = v39
                            _ = v40
                        }
                        _ = fi38
                    }
                    _ = fi35
                }
            case "foo bar":
                delete(v3, 42)
            default:
                v41 := GenS1{F0: 1}
                _ = v41
            }
        }
    default:
        switch v32 {
        case "":
            var v42 int8 = 6
            _ = v42
        default:
            for fi43 := 0; fi43 < 18; fi43++ {
                fmt.Println(v1, v1, v1)
                switch v32 {
                case "hello":
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v32 = v32
                        var v45 int8 = v1
                        _ = fi44
                        _ = v45
                    }
                case "":
                    v32 = v32
                case "a":
                    if (42 > genFn0(true, -13)) {
                        var v46 int = fi43
                        var v47 int16 = -1
                        _ = v46
                        _ = v47
                    }
                default:
                    var v48 int8 = v1
                    _ = v48
                }
                _ = fi43
            }
        }
    }
    v49 := GenS1{F0: 99}
    fmt.Println(v32)
    switch "" {
    case "\n\t":
        if ((v32 <= v32) && ((true || (!true)) || false)) {
            var v50 float64 = (0.1 / 1.0)
            _ = v50
        } else {
            v51 := GenS3{F0: 1}
            var v52 float64 = 1.5
            v53 := GenS1{F0: 54}
            _ = v51
            _ = v52
            _ = v53
        }
    case "":
        var v54 int = 940
        _ = v54
    }
    v1 = -128
    if false {
        if (255 < 65535) {
            for fi55 := 0; fi55 < 3; fi55++ {
                var v56 uint = 65535
                var v57 int64 = 510
                v56 += (^(v56))
                _ = fi55
                _ = v56
                _ = v57
            }
            if (18446744073709551615 == 1) {
                var v58 int32 = 1
                _ = v58
            }
        } else {
            v1 = v1
            var v59 uint32 = 93
            _ = v59
        }
        v60 := GenS3{F0: 264}
        v1 -= 127
        _ = v60
    }
    v32 = v32
    v61 := GenS4{F0: 1, F1: 18446744073709551615, F2: 32767}
    for fi62 := 0; fi62 < 3; fi62++ {
        fmt.Println(v32, fi62, v32)
        var v63 float32 = -2.25
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v49
    _ = v61
}
