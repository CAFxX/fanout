package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32, p1 bool, p2 int) uint32 {
    p2 = p2
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := []int{}
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 bool = p1
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        switch "" {
        case "foo bar":
            var v9 uint64 = 18446744073709551615
            _ = v9
        case "":
            p2 = int(p0)
        case "hello":
            v10 := GenS1{F0: 4294967295}
            _ = v10
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            p2 += fi8
            if false {
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    switch (-(fi8)) {
                    case 42:
                        switch "" {
                        case "foo bar":
                            var v12 uint8 = 255
                            _ = v12
                        case "a":
                            p2 -= p2
                        }
                    default:
                        p1 = ((!p1) && (true || (311 == 4294967295)))
                    }
                case "hello":
                    var v13 int64 = 0
                    _ = v13
                }
            }
            _ = fi11
        }
        _ = fi8
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 100
    var v3 int64 = 42
    var v14 int = 42
    var v15 int16 = 92
    var v16 string = "a"
    var v17 uint = 341
    var v18 int64 = (v3 & (v3 + (v3 ^ v3)))
    v15 -= v15
    fmt.Println(v2)
    switch (v16 + (v16 + "foo bar")) {
    case "\n\t":
        if (127 <= 255) {
            v1 += (v1 << ((v17) & 63))
            var v19 uint16 = (^(v1))
            _ = v19
        } else {
            switch "foo bar" {
            case "":
                fmt.Println(v17, v16)
            case "foo bar":
                for fi20 := 0; fi20 < 3; fi20++ {
                    v14 += (v14 + fi20)
                    var v21 uint64 = 867
                    if (254 <= 1) {
                        var v22 uint16 = uint16(v17)
                        v1 += v22
                        _ = v22
                    } else {
                        var v23 int = (fi20 % (((v14 / ((1) | 1))) | 1))
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            case "\n\t":
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 int8 = (-(v2))
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 uint8 = 128
                        v28 := GenS1{F0: v17}
                        var v29 int32 = 1
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi24
                    _ = v25
                }
            }
            if false {
                fmt.Println(v2, v3, v14)
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int16 = (-(v15))
                    for fi32 := 0; fi32 < 9; fi32++ {
                        v33 := []int{}
                        _ = fi32
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                v34 := GenS2{F0: false, F1: 255}
                for fi35 := 0; fi35 < 19; fi35++ {
                    fmt.Println(v3, fi35)
                    v17 = v17
                    _ = fi35
                }
                var v36 uint8 = 0
                _ = v34
                _ = v36
            }
        }
    case "hello":
        v37 := GenS3{F0: genFn0(1, (1.5 > 1.5), v14)}
        _ = v37
    }
    var v38 float32 = float32(0.0)
    v17 = 0
    var v39 uint32 = 4294967295
    for fi40 := 0; fi40 < 9; fi40++ {
        for fi41 := 0; fi41 < 12; fi41++ {
            v3 += 42
            var v42 int32 = 0
            _ = fi41
            _ = v42
        }
        fmt.Println(v1, v3, v18)
        v3 += int64(v15)
        _ = fi40
    }
    v17 += 690
    var v43 int32 = -2147483648
    for fi44 := 0; fi44 < 10; fi44++ {
        switch (fi44 >> ((501) & 63)) {
        case 9223372036854775807:
            v45 := GenS3{F0: 0}
            _ = v45
        default:
            v46 := GenS3{F0: uint32(v17)}
            _ = v46
        }
        fmt.Println(v43, v43)
        switch -87 {
        case 32767:
            v47 := []int{}
            _ = v47
        default:
            v18 -= v3
        }
        _ = fi44
    }
    var v48 int8 = 0
    v49 := []int{127, 42, v14, v14}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v43
    _ = v48
    _ = v49
}
