package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    v1 := []int{}
    v2 := []int{-2147483648}
    var v3 float64 = ((1.5 + ((float64(1.5) * (1e10 + 0.1)) * -2.25)) - 3.14159)
    var v4 uint8 = 86
    var v5 uint8 = v4
    if true {
        v5 += 170
        switch "foo bar" {
        case "\n\t":
            v6 := GenS1{F0: 18446744073709551615}
            _ = v6
        default:
            v5 = v5
        }
    } else {
        fmt.Println(v3)
        switch 4294967295 {
        case -843:
            v5 = v5
        }
        v5 = 140
    }
    if false {
        switch (("a" + "\n\t") + "foo bar") {
        case "\n\t":
            if false {
                if len(v2) > 0 {
                    v2[0] = -844
                }
                v7 := GenS2{F0: true, F1: "a"}
                var v8 uint = 2
                _ = v7
                _ = v8
            }
        case "a":
            v5 = v5
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v4, v4, v3)
    } else {
        var v9 int8 = 0
        var v10 uint64 = 771
        _ = v10
        _ = v9
    }
    v1 = append(v1, 32768)
    v4 += 197
    v5 += v4
    if (!(32767 == 65535)) {
        for fi11 := 0; fi11 < 5; fi11++ {
            if len(v2) > 0 {
                v2[0] = (fi11 / ((fi11) | 1))
            }
            _ = fi11
        }
        var v12 int8 = 1
        _ = v12
    }
    v13 := GenS2{F0: false, F1: ("\n\t" + ("foo bar" + ""))}
    v14 := GenS2{F0: true, F1: ("hello" + "a")}
    v3 -= v3
    var v15 string = "\n\t"
    fmt.Println(v15, v15, v3)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint64 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v3, v5, fi16)
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint32 = 65535
        _ = fi18
        _ = v19
    }
    if len(v2) > 0 {
        v2[0] = -9223372036854775808
    }
    v4 += v4
    switch 2147483647 {
    case -1:
        switch 423 {
        case -801:
            if len(v2) > 0 {
                v2[0] = -1
            }
        case 42:
            switch v15 {
            case "hello":
                if false {
                    v20 := GenS3{F0: "hello"}
                    for fi21 := 0; fi21 < 9; fi21++ {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            v15 = ("a" + v15)
                            fmt.Println(fi21, fi21)
                            _ = fi22
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    v1 = append(v1, 117)
                    v2 = append(v2, 1)
                }
            }
        default:
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint8 = v4
                v5 = (v5 >> ((1) & 63))
                _ = fi23
                _ = v24
            }
        }
    case 128:
        _ = len(v2)
    case -849:
        if false {
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := GenS4{F0: 255, F1: fi25}
                v5 += 255
                var v27 uint = 933
                _ = fi25
                _ = v26
                _ = v27
            }
            if false {
                var v28 uint64 = 0
                _ = v28
            } else {
                switch v15 {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    _ = len(v1)
                case "":
                    if (!(v5 != (v4 % ((1) | 1)))) {
                        v4 += v4
                    } else {
                        v3 = v3
                        var v29 int16 = 0
                        v30 := GenS4{F0: 255, F1: -1}
                        _ = v29
                        _ = v30
                    }
                default:
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 uint8 = (v5 + 2)
                        _ = len(v2)
                        v4 = (v4 << ((726) & 63))
                        _ = fi31
                        _ = v32
                    }
                }
            }
        }
    default:
        if len(v1) > 0 {
            v1[0] = 32767
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
