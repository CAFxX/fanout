package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 bool; F1 int; F2 string }
type GenS5 struct { F0 bool; F1 uint64; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{-863, 4294967295, 315}
    var v3 int8 = -59
    var v4 int16 = 0
    var v5 uint8 = 255
    if true {
        switch "a" {
        case "a":
            var v6 uint8 = v5
            _ = v6
        case "":
            switch -400 {
            case 425:
                if (v3 < -128) {
                    for fi7 := 0; fi7 < 10; fi7++ {
                        if false {
                            var v8 int16 = (v4 << ((18446744073709551615) & 63))
                            _ = v8
                        }
                        _ = fi7
                    }
                } else {
                    fmt.Println(v4, v3)
                    if len(v2) > 0 {
                        v2[0] = 32767
                    }
                    switch 4294967295 {
                    case 2147483647:
                        var v9 uint32 = 2
                        _ = v9
                    }
                }
            case -2147483648:
                v5 -= (v5 + v5)
            case -12:
                for fi10 := 0; fi10 < 9; fi10++ {
                    switch (fi10 >> ((18446744073709551615) & 63)) {
                    case 168:
                        v11 := GenS1{F0: 42}
                        _ = v11
                    case -350:
                        v5 = 1
                    case 32768:
                        if (42 != 2147483647) {
                            v5 = v5
                        } else {
                            if true {
                                for fi12 := 0; fi12 < 14; fi12++ {
                                    _ = len(v1)
                                    _ = fi12
                                }
                                var v13 string = (((("a" + ("a" + ("\n\t" + "foo bar"))) + "a") + "hello") + ("foo bar" + ("hello" + "a")))
                                _ = v13
                            }
                            var v14 string = ("" + "a")
                            _ = v14
                        }
                    }
                    var v15 uint16 = 0
                    if false {
                        v1[326] = 2147483648
                        switch 100 {
                        case 4294967295:
                            v16 := GenS1{F0: 65535}
                            _ = v16
                        case 0:
                            v17 := GenS2{F0: (fi10 % ((fi10) | 1)), F1: -32768}
                            _ = v17
                        case 888:
                            v15 += 1
                        default:
                            for fi18 := 0; fi18 < 13; fi18++ {
                                v5 += v5
                                fmt.Println(v15, v4, fi10)
                                _ = fi18
                            }
                        }
                    }
                    _ = fi10
                    _ = v15
                }
            }
        default:
            switch "foo bar" {
            case "\n\t":
                delete(v1, 4294967295)
            case "hello":
                v4 = (v4 >> ((1) & 63))
            case "foo bar":
                _ = len(v1)
            }
        }
    }
    var v19 float64 = (-2.25 / 0.1)
    var v20 float32 = (float32(v19) + 1.5)
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 bool = false
            if false {
                var v24 int = fi21
                _ = v24
            } else {
                if (2147483648 > 734) {
                    var v25 float32 = 1.5
                    _ = v25
                } else {
                    v4 -= v4
                }
                switch fi21 {
                case -2147483648:
                    var v26 uint32 = 2147483648
                    _ = v26
                case 738:
                    switch fi21 {
                    case 1:
                        v27 := v1[2]
                        _ = v27
                    default:
                        v28 := GenS3{F0: 42, F1: 18446744073709551615}
                        _ = v28
                    }
                default:
                    if (0 == 309) {
                        var v29 int = fi21
                        _ = v29
                    }
                }
                v19 = v19
                v30 := GenS3{F0: 2, F1: 2}
                _ = v30
            }
            _ = fi22
            _ = v23
        }
        if false {
            switch ("foo bar" + "") {
            case "\n\t":
                fmt.Println(fi21)
            case "hello":
                var v31 uint64 = 75
                _ = v31
            case "foo bar":
                var v32 int8 = (v3 / ((v3) | 1))
                _ = v32
            default:
                var v33 string = "foo bar"
                _ = v33
            }
        }
        _ = fi21
    }
    fmt.Println(v3, v19)
    var v34 uint64 = 42
    for fi35 := 0; fi35 < 12; fi35++ {
        fmt.Println(v19)
        v34 = (v34 | 577)
        for fi36 := 0; fi36 < 3; fi36++ {
            v37 := GenS4{F0: (18446744073709551615 != 318), F1: (fi35 >> ((18446744073709551615) & 63)), F2: "foo bar"}
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    var v38 uint32 = 2
    v39 := GenS3{F0: 65535, F1: 255}
    switch ((("a" + "\n\t") + "\n\t") + "a") {
    case "a":
        v2 = append(v2, -194)
    }
    var v40 uint = 687
    v40 += v40
    var v41 int16 = v4
    if true {
        var v42 float64 = v19
        _ = v42
    } else {
        v43 := GenS5{F0: (512 == (^(v38))), F1: (v34 ^ v34), F2: 440}
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
