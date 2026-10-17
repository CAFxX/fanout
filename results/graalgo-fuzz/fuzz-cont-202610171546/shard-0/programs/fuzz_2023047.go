package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int; F1 uint16; F2 uint32 }
func genFn0(p0 int16) uint32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int32 = -2147483648
        switch fi5 {
        case 127:
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := []int{2147483647, fi7}
                fmt.Println(fi7, v6)
                _ = fi7
                _ = v8
            }
        case 732:
            for fi9 := 0; fi9 < 16; fi9++ {
                fmt.Println(v6, p0, v6)
                _ = fi9
            }
        case -930:
            v10 := []int{(fi5 ^ (fi5 / (((fi5 | (fi5 & (fi5 % (((fi5 << ((255) & 63))) | 1))))) | 1))), (fi5 / ((0) | 1)), (fi5 | fi5)}
            _ = v10
        }
        _ = fi5
        _ = v6
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            if (12 <= 128) {
                v13 := GenS1{F0: 0, F1: -1, F2: 32767}
                _ = v13
            }
            v14 := []int{fi12, -693, -645}
            var v15 uint16 = 424
            _ = fi12
            _ = v14
            _ = v15
        }
        switch fi11 {
        case 9223372036854775807:
            var v16 uint32 = 42
            _ = v16
        case 42:
            for fi17 := 0; fi17 < 6; fi17++ {
                if true {
                    var v18 int = (fi11 / (((fi11 % (((fi11 ^ (-(fi11)))) | 1))) | 1))
                    v19 := GenS1{F0: 255, F1: 924, F2: 949}
                    _ = v18
                    _ = v19
                }
                p0 += p0
                var v20 int32 = -1
                _ = fi17
                _ = v20
            }
        case 32768:
            for fi21 := 0; fi21 < 15; fi21++ {
                switch fi21 {
                case 42:
                    v22 := GenS2{F0: 524, F1: 42, F2: 1}
                    _ = v22
                }
                var v23 int16 = (-(p0))
                _ = fi21
                _ = v23
            }
        }
        p0 += p0
        _ = fi11
    }
    p0 -= -32768
    var v24 float64 = 3.14159
    _ = p0
    _ = v24
    return 0
}

func main() {
    var v1 int64 = 610
    v2 := make(map[uint32]uint32)
    var v3 int8 = int8(v1)
    v4 := make(map[string]bool)
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 uint16 = 0
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            v3 = (v3 - v3)
            _ = fi28
        }
        _ = fi27
    }
    if true {
        var v29 int = 9223372036854775807
        v29 += v29
        _ = v29
    }
    v3 = v3
    for fi30 := 0; fi30 < 11; fi30++ {
        switch fi30 {
        case -2147483648:
            v31 := GenS2{F0: 842, F1: 0, F2: v3}
            _ = v31
        default:
            v3 = v3
        }
        switch ("a" + ("" + "")) {
        case "foo bar":
            if true {
                v32 := []int{-19, 1}
                v33 := GenS3{F0: (genFn0(-1) <= 1), F1: 18446744073709551615}
                _ = v32
                _ = v33
            } else {
                var v34 float64 = 1e10
                v35 := GenS3{F0: true, F1: 18446744073709551615}
                _ = v34
                _ = v35
            }
        case "\n\t":
            fmt.Println(v1, fi30, v1)
        case "":
            var v36 float32 = 0.0
            _ = v36
        default:
            var v37 int8 = -1
            _ = v37
        }
        for fi38 := 0; fi38 < 19; fi38++ {
            v39 := GenS2{F0: 65535, F1: 2, F2: (v3 << ((0) & 63))}
            v3 -= (v3 + (-(v3)))
            fmt.Println(fi38, fi30, fi30)
            _ = fi38
            _ = v39
        }
        _ = fi30
    }
    var v40 bool = (!(!false))
    for fi41 := 0; fi41 < 1; fi41++ {
        if (false && v40) {
            v3 = (v3 & (v3 * (v3 - v3)))
            switch "foo bar" {
            case "foo bar":
                v42 := GenS2{F0: 0, F1: 65535, F2: -128}
                _ = v42
            case "":
                switch (fi41 >> ((255) & 63)) {
                case 42:
                    if ((1.5 / 1.0) <= ((-2.25 / (3.14159 - 1e10)) + -2.25)) {
                        var v43 int8 = 0
                        _ = v43
                    } else {
                        fmt.Println(fi41)
                    }
                }
            default:
                if (v40 || true) {
                    var v44 uint8 = 1
                    for fi45 := 0; fi45 < 5; fi45++ {
                        v46 := []int{(fi41 ^ 42), fi41, fi45}
                        _ = fi45
                        _ = v46
                    }
                    var v47 int = 182
                    _ = v44
                    _ = v47
                }
            }
        }
        v1 += (v1 | v1)
        _ = fi41
    }
    fmt.Println(v3, v1, v1)
    v48 := GenS3{F0: v40, F1: 2}
    v40 = v40
    v3 += 53
    var v49 uint64 = 969
    switch "a" {
    case "foo bar":
        v50 := GenS4{F0: 128, F1: 2147483648}
        _ = v50
    case "hello":
        _ = len(v4)
    default:
        switch int(v49) {
        case -445:
            if v40 {
                v3 -= v3
                v40 = ((v40 || ((-2147483648 >= 0) || (1.5 >= (0.0 - (1.5 - 0.0))))) && v40)
                v51 := GenS5{F0: (true || v40), F1: (v49 >> ((125) & 63))}
                _ = v51
            } else {
                fmt.Println(v3, v1)
                if ((-1 < -128) && (!v40)) {
                    v2[2147483647] = 42
                    switch 642 {
                    case -715:
                        v3 += (^(v3))
                    case 297:
                        v52 := GenS6{F0: int(v49), F1: 2, F2: genFn0(-32768)}
                        _ = v52
                    }
                    delete(v4, "")
                }
            }
        case -972:
            switch 2147483647 {
            case 0:
                var v53 uint8 = 2
                _ = v53
            case -2147483648:
                var v54 float64 = 3.14159
                _ = v54
            case 32767:
                var v55 int64 = (v1 << ((2) & 63))
                _ = v55
            }
        case 0:
            v1 = v1
        default:
            for fi56 := 0; fi56 < 19; fi56++ {
                if ((v40 && (!v40)) || v40) {
                    v40 = ((v49 % ((2) | 1)) > (v49 / ((1000) | 1)))
                }
                switch "a" {
                case "foo bar":
                    var v57 int64 = v1
                    _ = v57
                case "a":
                    for fi58 := 0; fi58 < 8; fi58++ {
                        v59 := []int{1}
                        _ = fi58
                        _ = v59
                    }
                case "\n\t":
                    fmt.Println(v49)
                }
                _ = fi56
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v48
    _ = v49
}
