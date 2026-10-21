package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
func genFn0(p0 int32) int32 {
    for fi5 := 0; fi5 < 9; fi5++ {
        if (!false) {
            var v6 int32 = p0
            for fi7 := 0; fi7 < 6; fi7++ {
                v8 := GenS1{F0: 100, F1: fi5}
                _ = fi7
                _ = v8
            }
            var v9 string = "hello"
            _ = v6
            _ = v9
        }
        v10 := []int{}
        var v11 uint8 = 0
        _ = fi5
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 9223372036854775807}
    v13 := []int{2147483647, -9223372036854775808, -732}
    _ = p0
    _ = v12
    _ = v13
    return 265
}
func genFn1(p0 string, p1 int16, p2 int16) int {
    for fi14 := 0; fi14 < 10; fi14++ {
        p0 = p0
        _ = fi14
    }
    switch "foo bar" {
    case "hello":
        if (0.1 != (((0.1 + float64(3.14159)) * (1e10 * -2.25)) - (1.5 - 1.5))) {
            p2 = p2
        } else {
            var v15 uint = 1
            var v16 float32 = 1.5
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 int = fi17
                switch p0 {
                case "":
                    if false {
                        v18 = 32767
                    }
                case "hello":
                    if true {
                        p0 = p0
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 int16 = 1
                            fmt.Println(v16, v18, v16)
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        var v21 uint64 = 4294967295
                        var v22 int8 = 1
                        _ = v21
                        _ = v22
                    }
                case "\n\t":
                    v23 := GenS2{F0: ((true || true) || (2 <= 218))}
                    _ = v23
                }
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
    }
    var v24 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 4294967295
}

func main() {
    var v1 uint32 = 567
    var v2 int64 = -9223372036854775808
    var v3 uint = 255
    v4 := []int{127}
    switch "foo bar" {
    case "":
        for fi25 := 0; fi25 < 18; fi25++ {
            switch "\n\t" {
            case "":
                fmt.Println(v3, v3, v2)
            }
            v2 += 1
            _ = fi25
        }
    default:
        v26 := GenS2{F0: (-1 != -128)}
        _ = v26
    }
    var v27 uint64 = 2
    var v28 bool = true
    for fi29 := 0; fi29 < 8; fi29++ {
        switch genFn1(("hello" + "foo bar"), -1, 32767) {
        case 32768:
            if (true || ((true || v28) || v28)) {
                var v30 float32 = 1e10
                _ = len(v4)
                for fi31 := 0; fi31 < 12; fi31++ {
                    switch (fi29 << (((v3 * (v3 - v3))) & 63)) {
                    case 0:
                        var v32 uint8 = 2
                        _ = v32
                    }
                    v3 -= (v3 & 857)
                    _ = fi31
                }
                _ = v30
            }
        }
        v27 = (^(v27))
        var v33 uint16 = 32767
        _ = fi29
        _ = v33
    }
    v2 += v2
    v3 += v3
    switch 32768 {
    case -2147483648:
        var v34 float32 = (-2.25 / 1.5)
        _ = v34
    case 2147483647:
        var v35 int64 = 42
        _ = v35
    }
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v36 int8 = 74
    var v37 int64 = (^(v2))
    var v38 uint = (v3 / ((18446744073709551615) | 1))
    var v39 uint16 = 42
    var v40 uint8 = 2
    if v28 {
        if false {
            var v41 uint8 = (^(v40))
            v36 = (v36 * (v36 - v36))
            v1 = v1
            _ = v41
        }
        if true {
            if v28 {
                for fi42 := 0; fi42 < 5; fi42++ {
                    if v28 {
                        switch ("\n\t" + "hello") {
                        case "hello":
                            v36 = -1
                        default:
                            v43 := GenS1{F0: (v36 + 127), F1: fi42}
                            _ = v43
                        }
                    }
                    v44 := GenS1{F0: v36, F1: (-(fi42))}
                    switch fi42 {
                    case 395:
                        v28 = false
                    case -264:
                        for fi45 := 0; fi45 < 19; fi45++ {
                            if v28 {
                                v46 := GenS3{F0: 1}
                                _ = v46
                            }
                            v36 += v36
                            _ = fi45
                        }
                    case 9223372036854775807:
                        v47 := GenS4{F0: ("foo bar" < "\n\t"), F1: 0, F2: ((!false) || ((("hello" + "hello") + "") < "hello"))}
                        _ = v47
                    default:
                        var v48 string = ("hello" + "\n\t")
                        _ = v48
                    }
                    _ = fi42
                    _ = v44
                }
                switch 1 {
                case 128:
                    var v49 string = "\n\t"
                    _ = v49
                }
            }
        }
        if v28 {
            v1 += uint32(v40)
            for fi50 := 0; fi50 < 4; fi50++ {
                fmt.Println(v38, v39)
                _ = fi50
            }
        }
    }
    v51 := GenS2{F0: true}
    if len(v4) > 0 {
        v4[0] = 127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v51
}
