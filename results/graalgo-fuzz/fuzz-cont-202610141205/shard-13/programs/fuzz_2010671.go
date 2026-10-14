package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0() int16 {
    if false {
        v7 := GenS1{F0: false}
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint16 = 2
            var v10 int64 = 914
            _ = fi8
            _ = v10
            _ = v9
        }
        switch "a" {
        case "":
            v11 := GenS1{F0: (255 <= 2)}
            _ = v11
        case "hello":
            var v12 string = "hello"
            _ = v12
        case "foo bar":
            var v13 int32 = -2147483648
            _ = v13
        default:
            v14 := GenS1{F0: (-230 == 350)}
            _ = v14
        }
        _ = v7
    }
    switch (("a" + "foo bar") + "") {
    case "a":
        var v15 int16 = 42
        _ = v15
    case "":
        v16 := GenS2{F0: 42}
        _ = v16
    }
    switch "" {
    case "\n\t":
        switch 32768 {
        case 32767:
            var v17 uint32 = 0
            _ = v17
        default:
            switch "" {
            case "hello":
                var v18 string = "a"
                _ = v18
            case "a":
                if (2 < 4294967295) {
                    if (-2147483648 >= 1) {
                        var v19 float64 = -2.25
                        var v20 float32 = 1.5
                        _ = v19
                        _ = v20
                    } else {
                        v21 := []int{239, 1}
                        var v22 bool = true
                        if (true || (-1 != -322)) {
                            if len(v21) > 0 {
                                v21[0] = -1
                            }
                            var v23 bool = false
                            v23 = ((1 != -1) && v22)
                            _ = v23
                        }
                        _ = v21
                        _ = v22
                    }
                }
            default:
                if (!false) {
                    v24 := []int{4294967295, 127, -1, -9223372036854775808}
                    v25 := GenS1{F0: ((3.14159 * (float64((1e10 / float32(-2.25))) + (0.0 / ((0.0 * (-2.25 * 0.1)) + 0.1)))) > ((0.0 / 1.5) / 1.0))}
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint = 782
                    _ = v26
                }
            }
        }
    default:
        v27 := GenS2{F0: 2147483648}
        _ = v27
    }
    var v28 int = 127
    _ = v28
    return 32767
}

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 680
    v3 := []int{}
    v4 := []int{1, -1}
    v5 := make(map[string]bool)
    v6 := []uint32{2147483648, 2147483648, 915, 2}
    var v29 uint = 1
    v4 = append(v4, -710)
    if false {
        switch "a" {
        case "foo bar":
            v2 = 65535
        case "":
            v30 := GenS3{F0: (795 >= int32(v29)), F1: false}
            _ = v30
        default:
            v29 = v29
        }
    } else {
        var v31 uint32 = 42
        _ = v31
    }
    _ = len(v6)
    if false {
        switch 4294967295 {
        case 128:
            v29 = v29
        }
    } else {
        if true {
            v29 -= v29
            switch "hello" {
            case "a":
                v1 += v1
            case "hello":
                var v32 float32 = -2.25
                _ = v32
            default:
                v1 += v1
            }
        }
    }
    if (143 >= 1) {
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 int64 = 1
            v35 := GenS2{F0: 520}
            _ = fi33
            _ = v34
            _ = v35
        }
        v2 -= (v2 % ((v2) | 1))
        switch -83 {
        case 588:
            v36 := GenS2{F0: 65535}
            _ = v36
        case 32767:
            var v37 int = 42
            _ = v37
        default:
            fmt.Println(v1, v2, v1)
        }
    } else {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 uint64 = 297
            if ((v2 + 32767) >= v2) {
                var v40 uint32 = 65535
                var v41 int64 = 9223372036854775807
                _ = v40
                _ = v41
            }
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v39, v39)
            }
            _ = fi38
            _ = v39
        }
        var v42 int32 = -1
        if len(v6) > 0 {
            v6[0] = 65535
        }
        _ = v42
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 int16 = -1
        _ = fi43
        _ = v44
    }
    if (v29 != v29) {
        if false {
            fmt.Println(v2, v29)
            for fi45 := 0; fi45 < 11; fi45++ {
                v2 = 2
                v46 := GenS3{F0: ((!((0.0 * 0.1) > 1.5)) || false), F1: (1000 != 1)}
                _ = fi45
                _ = v46
            }
            for fi47 := 0; fi47 < 11; fi47++ {
                v48 := GenS2{F0: 0}
                v29 = v29
                delete(v5, "foo bar")
                _ = fi47
                _ = v48
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                var v49 int16 = -1
                _ = v49
            default:
                v2 -= (^(v2))
            }
            for fi50 := 0; fi50 < 20; fi50++ {
                v51 := GenS4{F0: -330, F1: 42}
                _ = fi50
                _ = v51
            }
        }
        var v52 bool = (188 != 42)
        var v53 uint8 = 127
        _ = v52
        _ = v53
    } else {
        var v54 int64 = 1
        var v55 float32 = (3.14159 - (1e10 - ((float32(v1) - (((float32(v1) - (((((0.0 * 0.0) * 0.0) / -2.25) * ((3.14159 / (1e10 - 0.0)) / 1.0)) + -2.25)) * float32(1e10)) - float32((0.1 + (v1 + 1e10))))) + (0.1 + 1.5))))
        _ = v54
        _ = v55
    }
    v2 = (v2 & (v2 ^ 774))
    if len(v4) > 0 {
        v4[0] = 1
    }
    fmt.Println(v2, v2)
    var v56 int8 = 1
    if true {
        v29 += v29
        var v57 bool = false
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
