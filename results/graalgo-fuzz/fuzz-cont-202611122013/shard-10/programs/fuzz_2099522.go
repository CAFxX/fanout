package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 string, p1 int8, p2 int16) string {
    if true {
        for fi4 := 0; fi4 < 12; fi4++ {
            for fi5 := 0; fi5 < 15; fi5++ {
                fmt.Println(p1)
                _ = fi5
            }
            _ = fi4
        }
        switch p0 {
        case "":
            if ((-450 >= 42) || (604 < 4294967295)) {
                v6 := GenS1{F0: 1, F1: -9223372036854775808}
                for fi7 := 0; fi7 < 11; fi7++ {
                    fmt.Println(fi7, fi7, p2)
                    var v8 uint64 = 824
                    if false {
                        v8 = v8
                        switch fi7 {
                        case 4294967295:
                            var v9 int8 = (p1 >> ((131) & 63))
                            _ = v9
                        }
                        v10 := GenS2{F0: 2147483647, F1: 4294967295}
                        _ = v10
                    }
                    _ = fi7
                    _ = v8
                }
                var v11 int16 = -32768
                _ = v11
                _ = v6
            }
        }
        var v12 float64 = 3.14159
        _ = v12
    }
    v13 := GenS2{F0: 1, F1: 65535}
    switch p0 {
    case "foo bar":
        var v14 string = "hello"
        _ = v14
    case "":
        var v15 string = "foo bar"
        _ = v15
    }
    fmt.Println(p1, p2)
    var v16 int64 = -149
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v16
    return p0
}
func genFn1() int8 {
    v17 := []int{-1, 42, 643}
    var v18 string = "a"
    var v19 uint16 = 576
    _ = v17
    _ = v18
    _ = v19
    return 107
}

func main() {
    var v1 int = 42
    v2 := []int{}
    v3 := make(map[uint32]bool)
    v1 = (v1 << ((4294967295) & 63))
    if false {
        _ = len(v3)
        if (128 != 97) {
            for fi20 := 0; fi20 < 7; fi20++ {
                v21 := GenS1{F0: 42, F1: fi20}
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS1{F0: 0, F1: fi22}
                    _ = fi22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch 0 {
            case 2147483647:
                var v24 uint8 = 42
                _ = v24
            }
        }
    } else {
        var v25 uint32 = 42
        v26 := GenS2{F0: 0, F1: 554}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        v1 = (v1 | v1)
        _ = fi27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 float64 = -2.25
        v30 := GenS1{F0: -594, F1: 128}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        switch fi31 {
        case -2147483648:
            fmt.Println(v1, v1)
        default:
            var v32 uint32 = 2147483648
            _ = v32
        }
        _ = fi31
    }
    var v33 int64 = -1
    v33 -= v33
    switch v1 {
    case 511:
        var v34 uint32 = 2147483647
        _ = v34
    }
    v33 = (v33 * (v33 * (v33 & -586)))
    switch v1 {
    case 835:
        v35 := GenS3{F0: (("hello" + "foo bar") + "\n\t"), F1: 430}
        _ = v35
    case 1:
        for fi36 := 0; fi36 < 14; fi36++ {
            if ((32767 >= -1) && (3.14159 == (0.0 + 0.0))) {
                var v37 uint8 = 255
                fmt.Println(v1, fi36)
                _ = v37
            }
            _ = fi36
        }
    default:
        for fi38 := 0; fi38 < 9; fi38++ {
            switch (v1 + fi38) {
            case 584:
                var v39 bool = false
                _ = v39
            case -1:
                v40 := v3[653]
                _ = v40
            case 128:
                fmt.Println(fi38, fi38)
            }
            _ = fi38
        }
    }
    var v41 uint32 = 2147483647
    v3[v41] = true
    var v42 uint16 = 42
    var v43 uint16 = v42
    switch "foo bar" {
    case "\n\t":
        v43 -= v42
    case "":
        switch v1 {
        case 1:
            var v44 float64 = ((1e10 - (0.1 + (0.1 / (1.5 - 0.0)))) + float64(-2.25))
            _ = v44
        }
    default:
        var v45 int32 = 42
        _ = v45
    }
    for fi46 := 0; fi46 < 20; fi46++ {
        var v47 int8 = 100
        var v48 int = 399
        _ = fi46
        _ = v47
        _ = v48
    }
    var v49 uint16 = v43
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v41
    _ = v42
    _ = v43
    _ = v49
}
