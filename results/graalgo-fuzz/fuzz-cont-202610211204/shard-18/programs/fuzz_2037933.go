package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint64; F1 int16; F2 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }

func main() {
    v1 := []uint32{2, 107, 377, 551}
    var v2 uint = 65535
    var v3 uint8 = 41
    v4 := []int{-435, 32767, -385}
    var v5 string = ("" + "foo bar")
    var v6 uint32 = 65535
    fmt.Println(v2, v2, v2)
    for fi7 := 0; fi7 < 3; fi7++ {
        switch fi7 {
        case -9223372036854775808:
            v8 := GenS1{F0: 9223372036854775807, F1: -2147483648}
            _ = v8
        case 2147483647:
            v9 := GenS1{F0: -9223372036854775808, F1: int(v3)}
            _ = v9
        case -692:
            _ = len(v1)
        default:
            v5 = v5
        }
        v5 = "\n\t"
        if len(v4) > 0 {
            v4[0] = (^(fi7))
        }
        _ = fi7
    }
    var v10 int32 = 42
    var v11 uint8 = v3
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: 262, F1: fi12}
        for fi14 := 0; fi14 < 10; fi14++ {
            if (!((false || false) || ((false && false) && ((!false) || false)))) {
                _ = len(v4)
                switch "\n\t" {
                case "foo bar":
                    switch fi14 {
                    case -947:
                        v5 = (v5 + v5)
                    case 32768:
                        v3 = 2
                    case -9223372036854775808:
                        switch (("a" + v5) + "foo bar") {
                        case "hello":
                            var v15 float64 = (float64(1e10) * 3.14159)
                            _ = v15
                        case "a":
                            _ = len(v1)
                            var v16 uint32 = v6
                            _ = v16
                        default:
                            fmt.Println(v10)
                        }
                    default:
                        var v17 int16 = -1
                        v6 -= uint32(v3)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v17
                    }
                }
            }
            _ = fi14
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 uint = 0
            _ = fi18
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    v20 := GenS2{F0: 2, F1: -1, F2: ((v5 + (v5 + (v5 + v5))) + v5)}
    var v21 uint8 = (v3 % (((^(v11))) | 1))
    for fi22 := 0; fi22 < 8; fi22++ {
        v23 := GenS3{F0: 797, F1: v11}
        v2 += 1
        _ = fi22
        _ = v23
    }
    if false {
        v3 += v21
    } else {
        v5 = "\n\t"
    }
    fmt.Println(v2, v10)
    var v24 int = 378
    switch ((("" + ("foo bar" + v5)) + "hello") + v5) {
    case "":
        switch v5 {
        case "hello":
            v25 := GenS4{F0: 65535, F1: true}
            _ = v25
        case "":
            v5 = v5
        }
    case "\n\t":
        v24 += -742
    case "hello":
        for fi26 := 0; fi26 < 7; fi26++ {
            if ((!true) || (!true)) {
                v5 = v5
                if false {
                    v21 = v3
                    _ = len(v1)
                } else {
                    if len(v4) > 0 {
                        v4[0] = (fi26 ^ (v24 + 32768))
                    }
                    fmt.Println(v24)
                    v6 = (v6 | (v6 - 572))
                }
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS5{F0: 500}
                    _ = fi27
                    _ = v28
                }
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := GenS6{F0: 42}
                if (v10 > (v10 >> ((v2) & 63))) {
                    switch v5 {
                    case "a":
                        fmt.Println(fi26)
                    case "foo bar":
                        fmt.Println(v2)
                    }
                } else {
                    if (620 > 1) {
                        v11 += (v21 / (((v3 ^ 1)) | 1))
                    } else {
                        if false {
                            v31 := GenS6{F0: -1}
                            _ = v31
                        } else {
                            v10 += (-(v10))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch (fi29 >> ((v2) & 63)) {
                            case 1:
                                v32 := GenS2{F0: 4294967295, F1: -489, F2: "a"}
                                _ = v32
                            case 9223372036854775807:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case -9223372036854775808:
                                switch (v5 + "a") {
                                case "":
                                    v33 := GenS6{F0: 0}
                                    _ = v33
                                }
                            }
                            v21 = (v3 << ((868) & 63))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v34 uint16 = 255
                            _ = len(v4)
                            _ = v34
                        }
                        switch -621 {
                        case 127:
                            var v35 uint64 = 42
                            _ = v35
                        }
                    }
                }
                _ = fi29
                _ = v30
            }
            _ = fi26
        }
    }
    v10 -= -2147483648
    v3 = v21
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
