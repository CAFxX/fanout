package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint {
    v4 := GenS1{F0: 128, F1: 399, F2: 0}
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    var v6 uint32 = 65535
    v7 := GenS1{F0: 1, F1: 32767, F2: 42}
    _ = v4
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 bool, p2 string) uint {
    p1 = p1
    var v8 int64 = 42
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{(-(fi9))}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    v3 := []uint32{165, 1, 42}
    var v11 int16 = -1
    var v12 int32 = -2147483648
    switch (("foo bar" + ("\n\t" + "\n\t")) + "") {
    case "":
        v2[752] = 42
    default:
        v1 = (v1 - 773)
    }
    v13 := GenS2{F0: 42}
    for fi14 := 0; fi14 < 2; fi14++ {
        v1 = 2
        if false {
            _ = len(v3)
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch 128 {
                case 2147483647:
                    v15 := GenS1{F0: 255, F1: int(v1), F2: 42}
                    _ = v15
                case 1:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 uint32 = 2
                        _ = len(v3)
                        v18 := GenS1{F0: 128, F1: fi14, F2: 128}
                        if false {
                            v19 := GenS3{F0: ((^(v1)) > (v1 ^ 18446744073709551615)), F1: false}
                            if false {
                                switch ("\n\t" + "") {
                                case "foo bar":
                                    var v20 uint32 = v17
                                    _ = v20
                                case "":
                                    if ((float32((0.1 * 0.1)) / (((0.1 - float32(((-2.25 - (0.0 * 1.5)) / ((-2.25 * (((-2.25 + (float64(3.14159) / 1.0)) / (1.5 * float64((3.14159 - float32(-2.25))))) * (3.14159 / 1.0))) * ((-2.25 / (1e10 + ((-2.25 + 3.14159) + 3.14159))) * -2.25))))) - (3.14159 * float32((-2.25 * ((0.0 / 0.1) + (0.0 - 1e10)))))) - float32((1.5 + 0.1)))) < 0.0) {
                                        fmt.Println(v1)
                                    }
                                default:
                                    fmt.Println(fi14, v17, fi16)
                                }
                                switch fi16 {
                                case 0:
                                    v12 = (^(v12))
                                default:
                                    fmt.Println(v1)
                                }
                            }
                            if false {
                                delete(v2, (^(fi16)))
                                var v21 uint8 = 127
                                var v22 uint32 = 405
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        }
                        var v23 int64 = 0
                        v12 = -252
                        if false {
                            fmt.Println(v11)
                            var v24 bool = (128 > 2)
                            var v25 float32 = float32(0.0)
                            _ = v24
                            _ = v25
                        }
                        var v26 int64 = v23
                        switch "foo bar" {
                        case "a":
                            v17 = uint32(v1)
                        case "hello":
                            v26 -= 1
                        default:
                            v12 = v12
                        }
                        switch fi16 {
                        case -9223372036854775808:
                            _ = len(v2)
                        case 37:
                            var v27 int32 = v12
                            _ = v27
                        case 1:
                            fmt.Println(v11, v1, v12)
                        default:
                            v1 = 4294967295
                        }
                        v12 = (v12 / ((v12) | 1))
                        fmt.Println(v23, v23)
                        delete(v2, (fi14 - 9223372036854775807))
                        v11 = (v11 << ((4294967295) & 63))
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v23
                        _ = v26
                    }
                case -683:
                    switch (fi14 << ((65535) & 63)) {
                    case 9223372036854775807:
                        if ((^(v1)) >= (v1 << ((0) & 63))) {
                            v12 -= (v12 & (v12 / (((-(v12))) | 1)))
                            switch "\n\t" {
                            case "foo bar":
                                var v28 uint8 = 128
                                _ = v28
                            }
                            v12 = v12
                        }
                    case 127:
                        fmt.Println(v12, v1, v12)
                    }
                default:
                    fmt.Println(v11)
                }
            }
            v12 = v12
        }
        _ = fi14
    }
    switch int(v12) {
    case 42:
        _ = len(v2)
    }
    var v29 bool = false
    var v30 int = 673
    var v31 int = v30
    if true {
        for fi32 := 0; fi32 < 2; fi32++ {
            fmt.Println(v11, fi32, v11)
            _ = fi32
        }
        var v33 bool = (!v29)
        _ = v33
    }
    switch (-(v30)) {
    case 55:
        var v34 uint64 = 42
        _ = v34
    case 2147483647:
        var v35 int = 42
        _ = v35
    case -523:
        v1 = 2
    default:
        delete(v2, v31)
    }
    var v36 uint8 = 2
    var v37 int = (v30 ^ (v31 ^ v31))
    fmt.Println(v31, v37)
    v38 := v2[4294967295]
    var v39 int32 = v12
    if (!(!false)) {
        switch ("hello" + "") {
        case "a":
            for fi40 := 0; fi40 < 16; fi40++ {
                var v41 uint16 = 65535
                _ = fi40
                _ = v41
            }
        case "":
            v37 += v31
        case "foo bar":
            fmt.Println(v1)
        }
        _ = len(v3)
    } else {
        var v42 bool = v29
        _ = v42
    }
    switch -122 {
    case 42:
        v43 := GenS2{F0: 959}
        _ = v43
    case 596:
        if (v31 <= (v37 / (((v38 << ((492) & 63))) | 1))) {
            var v44 int8 = 105
            _ = v44
        } else {
            v30 += (^(v30))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
