package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0() int {
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint32 = 166
        var v6 bool = (-1 != 42)
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 string = "foo bar"
    fmt.Println(v7)
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int32 = -1
        switch v7 {
        case "":
            if true {
                var v10 int64 = -1
                _ = v10
            }
        case "foo bar":
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := []int{fi8, fi8, (fi8 % ((fi8) | 1)), (^(fi8))}
                _ = fi11
                _ = v12
            }
        case "a":
            fmt.Println(v7, v7, fi8)
        }
        _ = fi8
        _ = v9
    }
    switch 4294967295 {
    case 1:
        if false {
            v7 = "\n\t"
            fmt.Println(v7)
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    fmt.Println(fi13, fi14, fi13)
                    v7 = (v7 + v7)
                    _ = fi14
                }
                v7 = ((v7 + "\n\t") + (("hello" + v7) + (v7 + v7)))
                v15 := GenS1{F0: 647}
                _ = fi13
                _ = v15
            }
            var v16 int8 = 42
            _ = v16
        }
    case -841:
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := []int{(fi17 * fi17), 32768}
            var v19 int = fi17
            _ = fi17
            _ = v18
            _ = v19
        }
    case -273:
        if (!false) {
            v20 := []int{-2147483648}
            switch (v7 + (v7 + v7)) {
            case "foo bar":
                var v21 int16 = 1
                _ = v21
            case "hello":
                var v22 int16 = 1
                _ = v22
            }
            var v23 int16 = 32767
            _ = v20
            _ = v23
        }
    }
    _ = v7
    return -215
}
func genFn1() int8 {
    switch "a" {
    case "hello":
        var v24 bool = false
        _ = v24
    case "foo bar":
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := []int{fi25}
            v27 := GenS1{F0: 42}
            if len(v26) > 0 {
                v26[0] = fi25
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    case "\n\t":
        v28 := []int{genFn0()}
        _ = v28
    default:
        var v29 uint64 = 42
        _ = v29
    }
    var v30 bool = false
    var v31 string = "hello"
    _ = v30
    _ = v31
    return -1
}

func main() {
    var v1 float64 = (1.5 / -2.25)
    var v2 int16 = -997
    v3 := []int{-588, 32767, 4294967295, 2147483647}
    v32 := GenS1{F0: 0}
    for fi33 := 0; fi33 < 6; fi33++ {
        v1 = v1
        for fi34 := 0; fi34 < 7; fi34++ {
            v1 = v1
            _ = fi34
        }
        var v35 int16 = v2
        _ = fi33
        _ = v35
    }
    fmt.Println(v2, v1)
    if true {
        var v36 int8 = -13
        _ = v36
    } else {
        var v37 bool = true
        var v38 string = ""
        for fi39 := 0; fi39 < 16; fi39++ {
            for fi40 := 0; fi40 < 12; fi40++ {
                v41 := GenS1{F0: 4294967295}
                _ = fi40
                _ = v41
            }
            var v42 bool = (!v37)
            var v43 uint8 = 127
            _ = fi39
            _ = v42
            _ = v43
        }
        _ = v37
        _ = v38
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        switch (^(fi44)) {
        case 1:
            if len(v3) > 0 {
                v3[0] = (fi44 % (((fi44 >> ((18446744073709551615) & 63))) | 1))
            }
        case -2147483648:
            for fi45 := 0; fi45 < 19; fi45++ {
                v46 := GenS1{F0: 18446744073709551615}
                _ = fi45
                _ = v46
            }
        case 20:
            v47 := GenS1{F0: 1}
            _ = v47
        default:
            v2 -= (v2 >> ((4294967295) & 63))
        }
        _ = len(v3)
        var v48 bool = true
        _ = fi44
        _ = v48
    }
    if true {
        _ = len(v3)
        v49 := GenS1{F0: 0}
        _ = len(v3)
        _ = v49
    }
    if (0 != 794) {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v50 float32 = (-2.25 / (3.14159 / -2.25))
            _ = v50
        }
    } else {
        for fi51 := 0; fi51 < 20; fi51++ {
            var v52 float64 = (1e10 + v1)
            var v53 float64 = v52
            _ = fi51
            _ = v52
            _ = v53
        }
        v2 = 1
        v1 = float64(3.14159)
    }
    var v54 int = 302
    var v55 int64 = 0
    v55 = (v55 / ((v55) | 1))
    if false {
        var v56 int8 = -1
        switch (((("\n\t" + "foo bar") + "a") + "a") + "") {
        case "\n\t":
            var v57 uint8 = 0
            _ = v57
        default:
            if len(v3) > 0 {
                v3[0] = v54
            }
        }
        switch ("hello" + "foo bar") {
        case "a":
            v3 = append(v3, v54)
        case "\n\t":
            switch 1 {
            case 32767:
                v58 := GenS2{F0: 65535, F1: 42, F2: false}
                _ = v58
            case 9223372036854775807:
                if (((!true) && false) || (genFn1() > int8(v2))) {
                    v1 = (1.5 / ((float64((3.14159 * (float32(float64(3.14159)) * float32(v1)))) + v1) * (((float64(1.5) - v1) - (v1 - v1)) - v1)))
                    switch (v54 >> ((18446744073709551615) & 63)) {
                    case -457:
                        _ = len(v3)
                    }
                    v3 = append(v3, (v54 * 32767))
                }
            case 2147483647:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                var v59 bool = (112 == 0)
                _ = v59
            }
        }
        _ = v56
    } else {
        v3 = append(v3, v54)
        fmt.Println(v54, v1)
    }
    var v60 int8 = -128
    if true {
        var v61 int = v54
        for fi62 := 0; fi62 < 1; fi62++ {
            var v63 int32 = -324
            v2 -= (-(v2))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi62
            _ = v63
        }
        v55 = (v55 | v55)
        _ = v61
    } else {
        var v64 uint16 = 817
        _ = v64
    }
    v1 = (v1 - v1)
    switch (v54 | (v54 ^ v54)) {
    case 9223372036854775807:
        var v65 int32 = 1
        _ = v65
    }
    v54 = -2147483648
    if false {
        var v66 float32 = 0.0
        v2 += 1000
        _ = v66
    }
    var v67 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v54
    _ = v55
    _ = v60
    _ = v67
}
