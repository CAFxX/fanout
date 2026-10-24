package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 string; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int32; F1 string; F2 uint64 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t"}
    var v2 int32 = 0
    v3 := []uint32{4294967295, 2147483648, 42, 987}
    v4 := make(map[int]int)
    var v5 bool = true
    var v6 uint64 = 42
    var v7 int = 689
    delete(v4, -63)
    if false {
        v2 -= -472
        switch v7 {
        case -429:
            v5 = (!(-1 > -939))
        case 2147483647:
            if v5 {
                var v8 bool = true
                v9 := GenS1{F0: -266, F1: 32767, F2: 2147483647}
                for fi10 := 0; fi10 < 18; fi10++ {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v4[-1] = fi11
                        var v12 uint64 = v6
                        _ = fi11
                        _ = v12
                    }
                    if true {
                        v1 = append(v1, "")
                    } else {
                        if v5 {
                            for fi13 := 0; fi13 < 6; fi13++ {
                                switch "" {
                                case "":
                                    switch (v7 * -504) {
                                    case 32767:
                                        fmt.Println(fi10)
                                        fmt.Println(v6, fi13)
                                        _ = len(v1)
                                    }
                                case "foo bar":
                                    v14 := v4[(fi10 / ((4294967295) | 1))]
                                    _ = v14
                                }
                                switch "\n\t" {
                                case "a":
                                    if (v5 || v8) {
                                        if (false && (-1 < 89)) {
                                            v15 := GenS2{F0: 0, F1: 0, F2: 9223372036854775807}
                                            v8 = v8
                                            _ = len(v4)
                                            _ = v15
                                        }
                                        v5 = false
                                    }
                                }
                                _ = fi13
                            }
                            fmt.Println(v7, fi10, v5)
                            if true {
                                fmt.Println(fi10, fi10, v7)
                                v16 := GenS2{F0: 32767, F1: (v2 % ((v2) | 1)), F2: -1}
                                if (!v5) {
                                    var v17 string = ""
                                    _ = v17
                                } else {
                                    for fi18 := 0; fi18 < 4; fi18++ {
                                        var v19 float64 = -2.25
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                _ = v16
                            } else {
                                v20 := GenS3{F0: ("a" + "\n\t"), F1: uint16(v6), F2: 95}
                                v2 = (v2 / ((v2) | 1))
                                v7 = v7
                                _ = v20
                            }
                        }
                        fmt.Println(v7)
                        v8 = v5
                    }
                    _ = fi10
                }
                _ = v8
                _ = v9
            }
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 bool = (v6 > (^(v6)))
            var v23 uint16 = 47
            for fi24 := 0; fi24 < 4; fi24++ {
                switch 4294967295 {
                case -229:
                    var v25 uint64 = 4294967295
                    _ = v25
                case 2147483647:
                    v4[v7] = 0
                case -1:
                    v6 = 0
                }
                fmt.Println(fi21, v7, fi21)
                if len(v1) > 0 {
                    v1[0] = ("a" + "")
                }
                _ = fi24
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        var v26 int16 = -220
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        v5 = v5
        var v28 bool = (false || false)
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(fi29, fi29, v7)
            _ = fi30
        }
        v31 := GenS4{F0: (v2 % ((2147483647) | 1)), F1: "hello", F2: v6}
        _ = fi29
        _ = v31
    }
    if v5 {
        if v5 {
            v32 := GenS5{F0: -278}
            _ = v32
        }
        switch 9223372036854775807 {
        case -2147483648:
            switch 42 {
            case -543:
                v33 := GenS2{F0: 161, F1: v2, F2: -9223372036854775808}
                _ = v33
            }
        default:
            var v34 int64 = 0
            _ = v34
        }
    } else {
        v2 = v2
        fmt.Println(v2)
    }
    var v35 bool = v5
    var v36 uint16 = 255
    var v37 uint64 = v6
    v38 := GenS5{F0: -9223372036854775808}
    if v5 {
        _ = len(v1)
        var v39 string = ""
        _ = v39
    } else {
        switch v7 {
        case 42:
            switch (((("\n\t" + "foo bar") + ("" + (("foo bar" + "a") + "foo bar"))) + "a") + ("a" + "")) {
            case "":
                var v40 uint64 = v6
                _ = v40
            default:
                v41 := GenS1{F0: -9223372036854775808, F1: (^(v7)), F2: v7}
                _ = v41
            }
        }
        var v42 uint8 = 128
        _ = v42
    }
    var v43 string = "hello"
    fmt.Println(v36)
    fmt.Println(v6, v5, v35)
    for fi44 := 0; fi44 < 20; fi44++ {
        switch v7 {
        case 9223372036854775807:
            var v45 float32 = 1.5
            _ = v45
        default:
            v46 := GenS4{F0: (v2 << ((175) & 63)), F1: v43, F2: (v37 + (v37 + 18446744073709551615))}
            _ = v46
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi44
    }
    v47 := v4[964]
    fmt.Println(v2, v2)
    v1 = append(v1, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
