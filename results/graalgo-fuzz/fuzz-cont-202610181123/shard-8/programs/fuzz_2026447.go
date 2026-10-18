package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int8, p1 string, p2 uint) uint64 {
    v5 := GenS1{F0: 999}
    var v6 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1() uint64 {
    switch (("foo bar" + ("foo bar" + "foo bar")) + "a") {
    case "\n\t":
        switch "" {
        case "\n\t":
            var v7 uint8 = 211
            _ = v7
        case "a":
            v8 := GenS2{F0: 521, F1: -1, F2: 2}
            _ = v8
        default:
            var v9 int64 = 42
            _ = v9
        }
    case "a":
        var v10 int8 = 127
        _ = v10
    }
    if true {
        var v11 int8 = -14
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint64 = genFn0(127, ("a" + "foo bar"), 42)
            for fi14 := 0; fi14 < 3; fi14++ {
                v11 = 1
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        var v15 uint16 = 65535
        v16 := GenS1{F0: -422}
        v17 := GenS2{F0: v15, F1: 1, F2: v15}
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: (2147483648 <= 365)}
    for fi19 := 0; fi19 < 3; fi19++ {
        if (false || true) {
            v20 := []int{-701, (fi19 / ((1) | 1)), fi19}
            var v21 bool = true
            _ = v20
            _ = v21
        } else {
            switch "" {
            case "hello":
                fmt.Println(fi19)
            default:
                fmt.Println(fi19)
            }
        }
        if (316 != 32767) {
            switch "\n\t" {
            case "foo bar":
                if (2147483647 != -992) {
                    var v22 uint16 = 65535
                    fmt.Println(fi19, v22)
                    _ = v22
                }
            }
        }
        v23 := GenS2{F0: 2, F1: -329, F2: 32767}
        _ = fi19
        _ = v23
    }
    if true {
        var v24 float64 = ((0.0 + 1e10) / ((3.14159 * (((0.1 - (0.1 + 1.5)) * ((1.5 * float64(-2.25)) + (3.14159 * 1e10))) * 3.14159)) * (3.14159 + 3.14159)))
        var v25 int64 = -1
        v25 -= (v25 | 1)
        _ = v24
        _ = v25
    } else {
        var v26 uint32 = 571
        switch "a" {
        case "hello":
            v27 := GenS1{F0: 32768}
            _ = v27
        case "\n\t":
            v26 -= v26
        case "":
            var v28 uint = 0
            _ = v28
        default:
            var v29 uint = 2
            _ = v29
        }
        _ = v26
    }
    _ = v18
    return 2
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]bool)
    var v3 float32 = float32(((v1 / 1.5) / v1))
    v4 := []int{128}
    v30 := GenS3{F0: (((false || ((v3 - v3) > (v3 / 1.5))) && true) && ((-2.25 - (1.5 * (v1 / v1))) <= ((v1 / ((v1 + 1.5) / (v1 + 1.5))) * v1)))}
    _ = len(v2)
    var v31 bool = (0 != 1)
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int64 = -9223372036854775808
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: -993}
    v1 -= v1
    v31 = v31
    delete(v2, 32767)
    if len(v4) > 0 {
        v4[0] = -1
    }
    v31 = (!true)
    if v31 {
        if ((-9223372036854775808 <= -402) || (v31 && false)) {
            var v35 bool = true
            var v36 int8 = 100
            if (436 > -2147483648) {
                v4 = append(v4, -532)
                v37 := GenS4{F0: 255, F1: 2}
                _ = v37
            }
            _ = v35
            _ = v36
        }
        var v38 int = 42
        _ = v38
    } else {
        switch -583 {
        case 4294967295:
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 string = ("hello" + "a")
                var v41 string = v40
                if (genFn0(87, "a", 729) != 2) {
                    var v42 int64 = 0
                    fmt.Println(v42, v1)
                    for fi43 := 0; fi43 < 2; fi43++ {
                        var v44 int8 = 127
                        var v45 int64 = v42
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    _ = v42
                }
                _ = fi39
                _ = v40
                _ = v41
            }
        case -248:
            v46 := GenS3{F0: (0 >= -646)}
            _ = v46
        }
        var v47 int32 = -2147483648
        if (0 >= 0) {
            var v48 bool = v31
            v3 = v3
            if (!(-2.25 < 0.1)) {
                var v49 int = 0
                switch "hello" {
                case "a":
                    for fi50 := 0; fi50 < 18; fi50++ {
                        for fi51 := 0; fi51 < 17; fi51++ {
                            fmt.Println(v3, v48)
                            _ = fi51
                        }
                        _ = fi50
                    }
                default:
                    var v52 uint64 = 0
                    _ = v52
                }
                _ = v49
            } else {
                v53 := GenS3{F0: true}
                for fi54 := 0; fi54 < 2; fi54++ {
                    var v55 uint32 = 606
                    var v56 int32 = 2147483647
                    delete(v2, 32767)
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                if v31 {
                    v57 := GenS1{F0: -2147483648}
                    v58 := GenS2{F0: 65535, F1: v47, F2: 0}
                    v31 = v48
                    _ = v57
                    _ = v58
                }
                _ = v53
            }
            _ = v48
        }
        _ = v47
    }
    v1 = v1
    var v59 uint32 = 0
    var v60 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v59
    _ = v60
}
