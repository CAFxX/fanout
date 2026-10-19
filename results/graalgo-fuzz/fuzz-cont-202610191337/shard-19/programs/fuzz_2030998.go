package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int32; F2 int16 }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -9223372036854775808
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = (1.5 - ((1e10 - (-2.25 * 0.1)) + float64((((-2.25 * float32(1e10)) + 1.5) + (1e10 * 3.14159)))))
        v8 := GenS1{F0: 18446744073709551615, F1: (42 < 0)}
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint8 = 127
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 = (v5 % ((507) | 1))
    var v11 float64 = 1e10
    fmt.Println(v3, v5, v5)
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS2{F0: ("hello" + ""), F1: (("hello" + "a") + ("" + "\n\t"))}
        if true {
            var v14 int64 = 0
            for fi15 := 0; fi15 < 12; fi15++ {
                switch "a" {
                case "a":
                    v16 := GenS3{F0: 392}
                    _ = v16
                case "foo bar":
                    if false {
                        v2 -= 32768
                        switch "\n\t" {
                        case "a":
                            v14 -= v14
                        case "":
                            v14 += v14
                        default:
                            var v17 float32 = -2.25
                            _ = v17
                        }
                    }
                }
                v5 += 42
                switch ("a" + (((("\n\t" + ("" + "\n\t")) + "foo bar") + "a") + "foo bar")) {
                case "a":
                    v18 := GenS2{F0: "\n\t", F1: "a"}
                    _ = v18
                default:
                    var v19 uint8 = 63
                    _ = v19
                }
                _ = fi15
            }
            if (v3 && true) {
                v20 := GenS3{F0: 2}
                for fi21 := 0; fi21 < 6; fi21++ {
                    v22 := []int{}
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    v5 -= (v5 - v5)
                    _ = fi23
                }
                _ = v20
            }
            _ = v14
        } else {
            fmt.Println(fi12)
            v11 += float64(0.1)
            if (1 == 372) {
                v24 := []int{(v2 | fi12), 0, (v2 << ((2) & 63))}
                _ = v24
            } else {
                switch -820 {
                case 128:
                    v5 -= v5
                case 2147483647:
                    var v25 uint32 = 2
                    _ = v25
                default:
                    var v26 uint16 = 32767
                    _ = v26
                }
                v27 := GenS2{F0: "hello", F1: ("" + "hello")}
                _ = v27
            }
        }
        v3 = (!true)
        _ = fi12
        _ = v13
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if v3 {
            for fi29 := 0; fi29 < 9; fi29++ {
                v1[fi28] = v3
                v30 := GenS4{F0: (v5 - v5)}
                if v3 {
                    v31 := []int{}
                    _ = v31
                } else {
                    var v32 uint32 = 2147483648
                    _ = v32
                }
                if (v3 || (!((v11 + v11) > float64(1.5)))) {
                    var v33 int64 = -1
                    switch (("a" + "\n\t") + "foo bar") {
                    case "hello":
                        _ = len(v4)
                        var v34 int16 = -72
                        _ = v34
                    }
                    var v35 int = fi29
                    _ = v33
                    _ = v35
                } else {
                    v36 := GenS5{F0: v5, F1: 1, F2: -802}
                    var v37 uint64 = v5
                    v38 := GenS3{F0: 255}
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi29
                _ = v30
            }
            switch "" {
            case "foo bar":
                fmt.Println(v5)
            }
            v39 := GenS3{F0: 376}
            _ = v39
        }
        _ = fi28
    }
    var v40 int8 = 1
    if true {
        v41 := GenS2{F0: "hello", F1: "foo bar"}
        _ = len(v4)
        _ = v41
    }
    var v42 uint8 = 2
    for fi43 := 0; fi43 < 9; fi43++ {
        v2 = 246
        fmt.Println(v3, v3)
        _ = fi43
    }
    v42 += 186
    v5 -= (v5 + (v5 + v5))
    v44 := GenS2{F0: "", F1: "hello"}
    v45 := GenS2{F0: "a", F1: ""}
    var v46 int16 = -213
    if (!(v5 <= (^(v5)))) {
        var v47 int8 = v40
        if v3 {
            switch 233 {
            case 4294967295:
                v48 := GenS6{F0: (v5 + 625)}
                _ = v48
            case 1:
                var v49 int32 = 42
                _ = v49
            }
        }
        if v3 {
            v50 := GenS2{F0: ("a" + ""), F1: "hello"}
            _ = v50
        } else {
            switch "\n\t" {
            case "foo bar":
                v51 := GenS4{F0: 1}
                _ = v51
            case "a":
                v46 = v46
            }
            var v52 int8 = (v47 | (-(v47)))
            v42 = 127
            _ = v52
        }
        _ = v47
    } else {
        var v53 uint32 = 674
        _ = v53
    }
    v40 -= v40
    v40 += v40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
