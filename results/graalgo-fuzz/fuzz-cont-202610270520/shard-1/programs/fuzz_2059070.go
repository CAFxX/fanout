package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int; F1 uint }
type GenS8 struct { F0 int; F1 uint }
func genFn0() int8 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 string = "\n\t"
        if false {
            var v6 string = "foo bar"
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    switch ("hello" + "\n\t") {
    case "hello":
        if (true && ((46 < 42) && (470 >= 0))) {
            v7 := GenS1{F0: 0, F1: 427}
            _ = v7
        } else {
            var v8 uint64 = 2
            _ = v8
        }
    default:
        var v9 int8 = 39
        _ = v9
    }
    if (0 >= 190) {
        v10 := GenS2{F0: true, F1: 536}
        var v11 int16 = 1
        switch -645 {
        case 32768:
            v12 := []int{9223372036854775807, -2147483648, -121}
            _ = v12
        case 42:
            v13 := GenS3{F0: 844, F1: 1}
            _ = v13
        case 127:
            var v14 uint8 = 128
            _ = v14
        default:
            var v15 int16 = 0
            _ = v15
        }
        _ = v10
        _ = v11
    }
    if true {
        var v16 bool = (true || false)
        _ = v16
    } else {
        var v17 int8 = -128
        var v18 int8 = 100
        _ = v17
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        switch 128 {
        case 127:
            var v19 uint = 255
            _ = v19
        case 0:
            var v20 int8 = 0
            _ = v20
        default:
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 int64 = -9223372036854775808
                v23 := GenS4{F0: 2, F1: fi21}
                var v24 uint16 = 289
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
        }
    case "foo bar":
        var v25 uint64 = 2
        _ = v25
    case "a":
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 int16 = -712
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 int = -2147483648
    v3 := []uint32{349}
    if (true || (122 == (v1 * v1))) {
        v2 = (v2 - v2)
        fmt.Println(v1, v1)
    }
    v29 := GenS1{F0: 18446744073709551615, F1: 708}
    for fi30 := 0; fi30 < 9; fi30++ {
        v2 = (v2 - -9223372036854775808)
        fmt.Println(v2)
        _ = len(v3)
        _ = fi30
    }
    v2 = v2
    fmt.Println(v2, v1)
    var v31 int32 = -2147483648
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS4{F0: 42, F1: (v2 + (v2 | fi32))}
        v31 = 374
        switch "\n\t" {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                var v34 uint = 42
                _ = v34
            case "a":
                v2 -= 32768
            }
        case "\n\t":
            switch fi32 {
            case 551:
                var v35 int8 = -128
                _ = v35
            }
        case "":
            var v36 string = "foo bar"
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    var v37 int = v2
    var v38 int32 = -1
    if true {
        var v39 uint64 = 884
        v40 := GenS3{F0: (^(v39)), F1: 2}
        _ = v39
        _ = v40
    }
    fmt.Println(v1)
    var v41 float32 = 1e10
    switch (v37 | v2) {
    case 127:
        v41 = (((v41 * v41) * v41) + 0.0)
    }
    if (!true) {
        var v42 string = "foo bar"
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v42
    } else {
        if true {
            for fi43 := 0; fi43 < 6; fi43++ {
                for fi44 := 0; fi44 < 7; fi44++ {
                    v45 := GenS5{F0: 474, F1: (fi43 << ((0) & 63))}
                    var v46 uint = 0
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                fmt.Println(v1, fi43, v37)
                _ = fi43
            }
            v1 -= (^(v1))
        } else {
            for fi47 := 0; fi47 < 2; fi47++ {
                if false {
                    v1 -= (v1 * 164)
                    var v48 uint16 = v1
                    _ = v48
                } else {
                    v31 = 1
                    v38 -= (v38 | v31)
                    v37 -= fi47
                }
                if true {
                    switch -492 {
                    case 9223372036854775807:
                        var v49 int32 = v38
                        _ = v49
                    case 1:
                        if false {
                            v3 = append(v3, 2)
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = 126
                    }
                }
                _ = fi47
            }
        }
        v37 += (v37 - (v2 | 1))
        v50 := GenS6{F0: (v2 / ((v37) | 1)), F1: "\n\t"}
        _ = v50
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        for fi52 := 0; fi52 < 7; fi52++ {
            if ((-(v31)) < 42) {
                v31 = -1
                v31 -= -2147483648
            }
            v53 := GenS7{F0: fi51, F1: 65535}
            v54 := GenS8{F0: (fi51 + fi52), F1: 773}
            v31 += v31
            var v55 string = "\n\t"
            _ = fi52
            _ = v53
            _ = v54
            _ = v55
        }
        if true {
            fmt.Println(v1, v2)
            fmt.Println(v41, v41, v37)
        } else {
            if ((v2 & v2) >= v2) {
                switch v2 {
                case -2147483648:
                    v56 := GenS6{F0: (v2 | (v37 | fi51)), F1: ""}
                    _ = v56
                }
                switch ((("\n\t" + "foo bar") + "") + "") {
                case "hello":
                    v57 := GenS5{F0: 0, F1: fi51}
                    _ = v57
                case "a":
                    v1 += v1
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = fi51
    }
    var v58 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v37
    _ = v38
    _ = v41
    _ = v58
}
