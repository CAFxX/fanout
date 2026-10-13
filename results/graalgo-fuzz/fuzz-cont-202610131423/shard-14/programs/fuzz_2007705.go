package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0(p0 int16, p1 uint) int8 {
    var v5 uint32 = 1
    p0 += -32768
    var v6 int = 0
    v7 := GenS1{F0: p0, F1: true}
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 bool = ((661 <= 783) || (false && true))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint16, p1 bool) int8 {
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: -32768, F1: (0 > uint32(p0))}
        switch "foo bar" {
        case "a":
            p0 += (p0 * 42)
        case "hello":
            var v12 uint16 = p0
            _ = v12
        case "\n\t":
            p0 = (p0 - 626)
        default:
            switch (fi10 / (((fi10 ^ 9223372036854775807)) | 1)) {
            case 32767:
                p1 = false
            case 32768:
                v13 := []int{}
                _ = v13
            }
        }
        v14 := []int{fi10, (fi10 + (fi10 - (fi10 & fi10))), (fi10 * (fi10 & 2147483647))}
        _ = fi10
        _ = v11
        _ = v14
    }
    if (0 <= 2147483648) {
        var v15 float32 = float32(1.5)
        _ = v15
    }
    _ = p0
    _ = p1
    return -59
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 2147483647
    var v3 int32 = -38
    v4 := []uint32{}
    fmt.Println(v3, v3)
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                switch "\n\t" {
                case "foo bar":
                    v1 = (v1 ^ (-(v1)))
                }
                var v18 int8 = -37
                _ = fi17
                _ = v18
            }
            if (!(false || (fi16 <= (fi16 | (fi16 * (^(fi16))))))) {
                if true {
                    var v19 bool = (v2 <= v2)
                    var v20 uint64 = 2
                    switch fi16 {
                    case 32767:
                        if v19 {
                            switch (fi16 / ((fi16) | 1)) {
                            case 4294967295:
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 uint32 = (v2 + v2)
                                    _ = fi21
                                    _ = v22
                                }
                            case -572:
                                var v23 float32 = 1.5
                                _ = v23
                            default:
                                v20 += (v20 << ((2) & 63))
                            }
                        } else {
                            var v24 int64 = -9223372036854775808
                            _ = v24
                        }
                    case 0:
                        fmt.Println(fi16, v1)
                    }
                    _ = v19
                    _ = v20
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi16
        }
        _ = len(v4)
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 int16 = -1
        var v27 int32 = (v3 / ((v3) | 1))
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v2)
    fmt.Println(v1, v1)
    if ("a" > (("a" + "a") + ("a" + "a"))) {
        v3 = (v3 * -1)
        v1 += (v1 & genFn1(32767, true))
    }
    if false {
        var v28 float64 = float64(1e10)
        if (42 > 42) {
            var v29 string = ((("a" + ("" + "a")) + "\n\t") + ("\n\t" + "\n\t"))
            for fi30 := 0; fi30 < 8; fi30++ {
                _ = len(v4)
                _ = fi30
            }
            v1 += v1
            _ = v29
        } else {
            switch "a" {
            case "\n\t":
                fmt.Println(v1)
            case "":
                v2 -= 2147483647
            case "a":
                var v31 float32 = float32((v28 / (-2.25 + float64((1.5 - -2.25)))))
                _ = v31
            }
        }
        var v32 float64 = 0.1
        _ = v28
        _ = v32
    } else {
        var v33 int = 42
        for fi34 := 0; fi34 < 4; fi34++ {
            switch "foo bar" {
            case "\n\t":
                var v35 int16 = 0
                _ = v35
            case "hello":
                var v36 int8 = v1
                _ = v36
            case "foo bar":
                var v37 uint64 = 1
                _ = v37
            }
            var v38 float32 = 1e10
            var v39 uint16 = 42
            _ = fi34
            _ = v38
            _ = v39
        }
        _ = v33
    }
    v4 = append(v4, v2)
    var v40 uint8 = 112
    var v41 uint8 = v40
    v42 := GenS1{F0: 42, F1: true}
    v43 := GenS2{F0: 0, F1: -325}
    for fi44 := 0; fi44 < 9; fi44++ {
        v45 := GenS1{F0: -1, F1: (!false)}
        _ = fi44
        _ = v45
    }
    fmt.Println(v40, v3)
    fmt.Println(v2, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
