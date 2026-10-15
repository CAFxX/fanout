package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int64 = 686
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 uint = 2
    v3[""] = ((-(v1)) != (v1 / (((v1 + v1)) | 1)))
    var v6 float32 = 1.5
    v6 -= v6
    var v7 string = "\n\t"
    if true {
        v6 = v6
    } else {
        v8 := GenS1{F0: v5, F1: 2, F2: (1 == 239)}
        v9 := GenS2{F0: ((!false) || (-1 < 42)), F1: (v5 >> ((v5) & 63))}
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            v6 = ((-2.25 - (v6 + v6)) - 3.14159)
            v11 := []int{(fi10 % ((fi10) | 1)), fi10, fi10}
            if (!true) {
                if ((false || (!false)) || (!(743 <= -1))) {
                    if true {
                        v1 = -1
                        if (!((false && false) || (v6 < 0.1))) {
                            fmt.Println(v6)
                            var v12 uint8 = 1
                            _ = v12
                        }
                    }
                } else {
                    _ = len(v11)
                    switch fi10 {
                    case 756:
                        var v13 int8 = 42
                        _ = v13
                    default:
                        var v14 float32 = 0.1
                        _ = v14
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
            }
            _ = v11
        } else {
            _ = len(v3)
            for fi15 := 0; fi15 < 17; fi15++ {
                v16 := []int{(fi15 * 32768), fi15, 1, 883}
                if (116 < 1) {
                    v5 = (v5 % (((^(v5))) | 1))
                }
                _ = fi15
                _ = v16
            }
        }
        var v17 uint = v5
        _ = fi10
        _ = v17
    }
    if true {
        var v18 uint16 = 0
        v6 = v6
        if true {
            switch -2147483648 {
            case 32768:
                switch "\n\t" {
                case "\n\t":
                    switch ("foo bar" + v7) {
                    case "a":
                        var v19 int32 = 1
                        _ = v19
                    default:
                        var v20 int = -890
                        _ = v20
                    }
                case "foo bar":
                    v2[65535] = (18446744073709551615 != 4294967295)
                case "":
                    for fi21 := 0; fi21 < 18; fi21++ {
                        var v22 string = v7
                        _ = fi21
                        _ = v22
                    }
                }
            case 127:
                fmt.Println(v6, v5, v18)
            case -9223372036854775808:
                if (!true) {
                    v18 += (v18 / ((65535) | 1))
                    fmt.Println(v6)
                }
            default:
                fmt.Println(v1, v7)
            }
            v23 := []int{-9223372036854775808}
            fmt.Println(v18)
            _ = v23
        }
        _ = v18
    } else {
        var v24 uint = (v5 & v5)
        var v25 bool = ((!false) && false)
        _ = v24
        _ = v25
    }
    var v26 int16 = -808
    v1 = (v1 / ((1) | 1))
    switch 4294967295 {
    case -410:
        delete(v3, v7)
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v27 int8 = -1
        _ = v27
    case "hello":
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 7; fi29++ {
                v26 = (v26 << (((^(v5))) & 63))
                var v30 uint8 = 0
                var v31 uint64 = 42
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
    case "":
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint32 = 2147483647
            v7 = (v7 + (v7 + (v7 + (v7 + v7))))
            _ = fi32
            _ = v33
        }
    default:
        _ = len(v3)
    }
    var v34 int16 = v26
    for fi35 := 0; fi35 < 17; fi35++ {
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := v3[((v7 + v7) + ((v7 + "\n\t") + ("foo bar" + ("\n\t" + v7))))]
            var v38 bool = v37
            v34 = 42
            _ = fi36
            _ = v37
            _ = v38
        }
        if true {
            v39 := []int{(fi35 << ((65535) & 63)), 127, fi35, (-(fi35))}
            _ = v39
        } else {
            v5 = v5
        }
        v40 := GenS3{F0: (v5 >> ((v5) & 63)), F1: 255}
        _ = fi35
        _ = v40
    }
    v34 += (-(v26))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
