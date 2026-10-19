package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int32 }

func main() {
    var v1 int16 = 42
    var v2 uint32 = 30
    var v3 uint16 = 1
    fmt.Println(v1)
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            v1 = -32768
            for fi5 := 0; fi5 < 18; fi5++ {
                v6 := []int{(fi4 + (fi4 % ((fi5) | 1))), (-(fi5)), -9223372036854775808, -2147483648}
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            switch "hello" {
            case "a":
                fmt.Println(v1)
            }
            fmt.Println(v1)
            var v8 int8 = int8(v1)
            _ = fi7
            _ = v8
        }
    } else {
        fmt.Println(v3, v2, v2)
        var v9 uint = 539
        _ = v9
    }
    v2 += v2
    v1 += (v1 + (v1 | v1))
    v2 -= v2
    v1 = 93
    if false {
        var v10 int16 = v1
        if (32767 == (-(v1))) {
            for fi11 := 0; fi11 < 13; fi11++ {
                if true {
                    var v12 uint16 = v3
                    var v13 int16 = (v10 ^ v1)
                    _ = v12
                    _ = v13
                }
                if (true || (("a" + "hello") == "\n\t")) {
                    v2 = (v2 * v2)
                    for fi14 := 0; fi14 < 5; fi14++ {
                        fmt.Println(v3)
                        v15 := GenS1{F0: ("foo bar" + "hello"), F1: 857, F2: int16(v3)}
                        _ = fi14
                        _ = v15
                    }
                    fmt.Println(fi11, fi11, v3)
                } else {
                    if true {
                        v16 := GenS2{F0: (v2 % (((v2 % ((2) | 1))) | 1)), F1: 100}
                        if true {
                            if true {
                                fmt.Println(fi11, v3)
                                switch (("" + "foo bar") + "a") {
                                case "\n\t":
                                    v17 := []int{(fi11 & 0), (-(fi11)), 42, fi11}
                                    _ = v17
                                default:
                                    if ((float32(float64((0.1 - float32((0.0 * (3.14159 - 0.0)))))) == 0.1) && ((-(fi11)) <= (fi11 / ((fi11) | 1)))) {
                                        switch "\n\t" {
                                        case "":
                                            switch fi11 {
                                            case 32767:
                                                if true {
                                                    v18 := []int{(fi11 / ((fi11) | 1)), 42, fi11}
                                                    v2 = (v2 ^ 42)
                                                    v3 = v3
                                                    _ = v18
                                                }
                                            }
                                        }
                                        var v19 string = "foo bar"
                                        v20 := []int{}
                                        _ = v19
                                        _ = v20
                                    }
                                }
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    v2 -= (v2 * v2)
                                    v22 := GenS1{F0: "", F1: -536, F2: 32767}
                                    var v23 int64 = 1
                                    v2 = (^(v2))
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            }
                        } else {
                            var v24 int8 = 127
                            switch "" {
                            case "":
                                switch ("a" + "hello") {
                                case "foo bar":
                                    var v25 int64 = -9223372036854775808
                                    _ = v25
                                }
                            case "a":
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    v1 += (v1 << ((255) & 63))
                                    v2 += 4294967295
                                    _ = fi26
                                }
                            case "foo bar":
                                if false {
                                    v1 += (v10 + (v10 % ((v1) | 1)))
                                    v24 = 1
                                    if true {
                                        v27 := GenS3{F0: 1}
                                        var v28 uint16 = v3
                                        _ = v27
                                        _ = v28
                                    } else {
                                        var v29 float32 = (3.14159 / 3.14159)
                                        fmt.Println(fi11, v2, v1)
                                        _ = v29
                                    }
                                }
                            }
                            v30 := GenS4{F0: (fi11 & (^(fi11))), F1: -2147483648}
                            _ = v24
                            _ = v30
                        }
                        _ = v16
                    }
                    var v31 uint16 = v3
                    var v32 int64 = 674
                    _ = v31
                    _ = v32
                }
                fmt.Println(fi11, v3)
                _ = fi11
            }
            v33 := []int{2147483647, 74, -399, 4294967295}
            _ = v33
        } else {
            v34 := []int{32767}
            v34 = append(v34, -315)
            _ = v34
        }
        _ = v10
    } else {
        var v35 string = "a"
        _ = v35
    }
    switch 32768 {
    case 2147483647:
        for fi36 := 0; fi36 < 11; fi36++ {
            v1 += -1
            _ = fi36
        }
    case 32768:
        switch (("\n\t" + ("" + "foo bar")) + "hello") {
        case "foo bar":
            v37 := GenS2{F0: 0, F1: 127}
            _ = v37
        }
    default:
        var v38 uint = 42
        _ = v38
    }
    v39 := []int{4294967295, 4294967295}
    v39 = append(v39, 2147483647)
    switch 868 {
    case 2147483647:
        var v40 uint16 = v3
        _ = v40
    }
    if true {
        fmt.Println(v1, v2)
        v2 = v2
        v1 = v1
    }
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    v39 = append(v39, int(v3))
    var v41 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v41
}
