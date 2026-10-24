package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int }
func genFn0(p0 string, p1 int8) uint64 {
    var v4 int64 = 1
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int16 = int16(fi5)
        p1 += -81
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (!(!true))
    switch -909 {
    case 2147483647:
        v9 := GenS1{F0: (42 == 4294967295), F1: p1}
        _ = v9
    case 4294967295:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 bool = ((fi10 - (-(fi10))) <= (fi10 << ((2) & 63)))
            _ = fi10
            _ = v11
        }
    case 128:
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{}
                switch fi13 {
                case 9223372036854775807:
                    v15 := GenS1{F0: (v8 || false), F1: (^(p1))}
                    _ = v15
                case 127:
                    var v16 int32 = -1
                    _ = v16
                default:
                    if (!((fi13 >> ((4294967295) & 63)) == (fi13 - fi12))) {
                        if (0 > 1) {
                            p1 -= 42
                            var v17 uint8 = 1
                            if v8 {
                                var v18 int64 = 116
                                if ("\n\t" > p0) {
                                    fmt.Println(v18, v8, p1)
                                    if true {
                                        fmt.Println(p0, v8, v8)
                                    } else {
                                        v4 = (v4 % ((v4) | 1))
                                    }
                                    var v19 float32 = 1e10
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v20 int8 = p1
                                        _ = v20
                                    case "a":
                                        v18 = 9223372036854775807
                                    default:
                                        if len(v14) > 0 {
                                            v14[0] = fi13
                                        }
                                    }
                                    _ = v19
                                } else {
                                    var v21 float32 = ((1.5 + float32((1.5 + (3.14159 * ((-2.25 / 1.0) - 3.14159))))) * 0.1)
                                    v22 := GenS2{F0: 0}
                                    v18 = v18
                                    _ = v21
                                    _ = v22
                                }
                                if ((42 != 255) || (v8 || (v8 && (false && ((p1 * (p1 / ((p1) | 1))) <= 100))))) {
                                    var v23 uint = 4294967295
                                    v24 := GenS2{F0: -738}
                                    fmt.Println(v17)
                                    var v25 uint16 = 0
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 int = 4294967295
                                    if (2 > v17) {
                                        var v27 uint = 4294967295
                                        _ = v27
                                    }
                                    _ = v26
                                }
                                _ = v18
                            }
                            p0 = (p0 + "foo bar")
                            _ = v17
                        } else {
                            var v28 uint = 2
                            v29 := GenS1{F0: (((v28 << (((^(v28))) & 63)) > v28) && (v8 && ((p1 & (p1 << ((18446744073709551615) & 63))) == -6))), F1: p1}
                            _ = v28
                            _ = v29
                        }
                        switch 886 {
                        case 0:
                            v14 = append(v14, fi13)
                        case -9223372036854775808:
                            fmt.Println(v4, fi13)
                        case -918:
                            v30 := GenS3{F0: (v4 * v4), F1: 65535, F2: 797}
                            switch (^(fi13)) {
                            case -886:
                                v14 = append(v14, -901)
                            case 4294967295:
                                switch (fi13 >> ((255) & 63)) {
                                case -339:
                                    v14 = append(v14, 1)
                                }
                            default:
                                fmt.Println(fi12, v8, v8)
                            }
                            _ = v30
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
            var v31 int = 0
            var v32 string = p0
            _ = fi12
            _ = v31
            _ = v32
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 2
}

func main() {
    var v1 int16 = -32768
    var v2 int16 = -720
    var v3 int64 = -9223372036854775808
    v33 := []int{127, -309}
    var v34 string = (("foo bar" + ("" + "")) + "foo bar")
    v3 -= (v3 >> ((4294967295) & 63))
    fmt.Println(v3, v34, v34)
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := GenS4{F0: 42, F1: v34}
        _ = fi35
        _ = v36
    }
    var v37 int16 = -32768
    for fi38 := 0; fi38 < 20; fi38++ {
        var v39 int64 = (^(v3))
        var v40 uint = 42
        _ = fi38
        _ = v39
        _ = v40
    }
    v41 := GenS5{F0: 0}
    if true {
        if false {
            fmt.Println(v2, v1, v2)
        }
    }
    switch 32768 {
    case -302:
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    case -1:
        if true {
            for fi42 := 0; fi42 < 3; fi42++ {
                if len(v33) > 0 {
                    v33[0] = fi42
                }
                for fi43 := 0; fi43 < 12; fi43++ {
                    var v44 uint8 = 127
                    v3 = v3
                    if false {
                        var v45 bool = (!false)
                        var v46 uint8 = (v44 >> ((0) & 63))
                        v1 -= (v1 * v37)
                        _ = v45
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                v34 = v34
                _ = fi42
            }
            switch "hello" {
            case "hello":
                v1 -= 740
            }
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
        }
    }
    v33 = append(v33, -574)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v41
}
