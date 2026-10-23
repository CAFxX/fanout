package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int16; F2 int16 }
type GenS7 struct { F0 uint32; F1 int; F2 uint8 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint64, p2 int64) uint8 {
    v6 := GenS1{F0: 951, F1: (p0 % ((-863) | 1)), F2: -678}
    p0 -= (^(p2))
    v7 := GenS2{F0: 128, F1: 1}
    var v8 int = 0
    var v9 int64 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint16, p1 uint8) uint64 {
    if (false && (!(!false))) {
        if ((true || true) && true) {
            for fi10 := 0; fi10 < 2; fi10++ {
                p1 += genFn0(770, 1, int64(fi10))
                switch (fi10 & fi10) {
                case 547:
                    for fi11 := 0; fi11 < 14; fi11++ {
                        p1 += p1
                        _ = fi11
                    }
                case -2147483648:
                    v12 := GenS3{F0: "\n\t", F1: (42 == 65535), F2: -32768}
                    _ = v12
                case 812:
                    var v13 string = "\n\t"
                    _ = v13
                }
                _ = fi10
            }
        } else {
            p0 = 2
            p1 = (p1 & p1)
            var v14 int32 = 451
            _ = v14
        }
        var v15 bool = true
        _ = v15
    }
    var v16 bool = true
    p1 -= 128
    for fi17 := 0; fi17 < 14; fi17++ {
        v16 = v16
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(fi17)
            switch fi17 {
            case 42:
                switch "a" {
                case "a":
                    p1 += p1
                    v19 := GenS4{F0: 400}
                    _ = v19
                }
            }
            switch "a" {
            case "":
                p0 = 1
            case "hello":
                var v20 uint32 = 2147483647
                _ = v20
            }
            switch (fi18 - (fi17 | fi18)) {
            case 6:
                if ((718 >= fi18) || (v16 || v16)) {
                    v21 := []int{fi18}
                    _ = len(v21)
                    v16 = (84 == -3)
                    _ = v21
                } else {
                    switch 32768 {
                    case 42:
                        v22 := GenS5{F0: -607, F1: -9223372036854775808, F2: (p1 ^ (^(p1)))}
                        _ = v22
                    }
                }
            case -9223372036854775808:
                v23 := []int{fi18, fi18, (fi18 ^ int(p0))}
                _ = v23
            case 32767:
                var v24 uint16 = p0
                _ = v24
            }
            v25 := GenS6{F0: (p0 << ((261) & 63)), F1: 656, F2: -32768}
            _ = fi18
            _ = v25
        }
        _ = fi17
    }
    switch "\n\t" {
    case "a":
        var v26 bool = v16
        _ = v26
    case "hello":
        switch 1 {
        case 42:
            var v27 float64 = ((3.14159 * float64(3.14159)) + 1.5)
            _ = v27
        case 37:
            if false {
                v28 := GenS7{F0: 987, F1: 127, F2: 128}
                if v16 {
                    var v29 uint8 = 0
                    v16 = v16
                    var v30 int32 = 42
                    _ = v29
                    _ = v30
                }
                _ = v28
            }
        case 128:
            var v31 float32 = 3.14159
            _ = v31
        default:
            v32 := []int{-391, 9223372036854775807, -413, 68}
            _ = v32
        }
    case "foo bar":
        fmt.Println(v16, v16)
    default:
        v33 := []int{128, 770, 418, 32767}
        _ = v33
    }
    _ = p0
    _ = p1
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 bool = ((2147483648 > 415) && false)
    v2 := []int{-922}
    var v3 string = "foo bar"
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v34 int = 42
    v4[4294967295] = v34
    if false {
        v35 := GenS8{F0: 127}
        var v36 string = ("\n\t" + v3)
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 uint = 255
        _ = fi37
        _ = v38
    }
    fmt.Println(v34, v34, v34)
    if (0 >= 0) {
        for fi39 := 0; fi39 < 2; fi39++ {
            v3 = "foo bar"
            v40 := GenS5{F0: -648, F1: 0, F2: 132}
            _ = fi39
            _ = v40
        }
        fmt.Println(v1, v34, v1)
        var v41 uint64 = 0
        _ = v41
    }
    v42 := GenS4{F0: 139}
    var v43 int64 = 656
    for fi44 := 0; fi44 < 7; fi44++ {
        v1 = false
        v34 = fi44
        _ = fi44
    }
    switch v34 {
    case 438:
        v3 = (v3 + "")
    case 32767:
        v3 = "\n\t"
    case 32768:
        for fi45 := 0; fi45 < 15; fi45++ {
            for fi46 := 0; fi46 < 14; fi46++ {
                if ((v43 * v43) > v43) {
                    var v47 uint8 = 123
                    v34 = (^(v34))
                    v48 := GenS1{F0: 733, F1: (v43 % ((v43) | 1)), F2: -1}
                    _ = v47
                    _ = v48
                }
                if len(v2) > 0 {
                    v2[0] = fi45
                }
                fmt.Println(v1, v3)
                _ = fi46
            }
            var v49 uint64 = 873
            v50 := v5[(v34 << ((255) & 63))]
            _ = fi45
            _ = v49
            _ = v50
        }
    }
    v4[42] = (v34 - 0)
    var v51 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v51
}
