package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 int; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0() int64 {
    var v5 uint32 = 0
    v5 = (v5 % ((v5) | 1))
    _ = v5
    return -1
}
func genFn1(p0 uint32, p1 uint) bool {
    fmt.Println(p0, p0)
    if false {
        v6 := GenS1{F0: -1}
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 float64 = -2.25
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := []int{(fi9 << ((0) & 63)), fi7, (fi9 * (fi9 ^ (-(fi9))))}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    p0 = (^(p0))
    if true {
        v11 := GenS1{F0: 32767}
        v12 := []int{128, 22, 261}
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint = (^(p1))
            if len(v12) > 0 {
                v12[0] = fi13
            }
            _ = fi13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    switch ("hello" + (((("\n\t" + "a") + "foo bar") + "") + ("\n\t" + "a"))) {
    case "":
        p0 -= (p0 ^ 4294967295)
    default:
        if false {
            var v15 uint = (p1 - (p1 - p1))
            switch "hello" {
            case "":
                fmt.Println(p0, v15)
            case "hello":
                var v16 int8 = 127
                _ = v16
            }
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return (-9223372036854775808 >= 1)
}

func main() {
    v1 := []string{"\n\t", "\n\t", ("hello" + "hello")}
    v2 := make(map[string]string)
    var v3 int64 = 0
    var v4 uint = 588
    var v17 string = ""
    if (((v4 & (v4 % (((v4 / ((v4) | 1))) | 1))) != (v4 - v4)) || (55 >= -810)) {
        v4 = (v4 << ((1) & 63))
        v3 = (v3 >> ((42) & 63))
        switch "a" {
        case "\n\t":
            v3 += 711
        }
    } else {
        switch v17 {
        case "\n\t":
            if true {
                v18 := v2[("hello" + "a")]
                switch (v17 + "") {
                case "":
                    v17 = "hello"
                }
                if len(v1) > 0 {
                    v1[0] = v18
                }
                _ = v18
            }
        case "":
            fmt.Println(v17, v17, v3)
        }
        v19 := v2[(v17 + v17)]
        if false {
            if (false || (true || (v4 < (v4 ^ v4)))) {
                if true {
                    var v20 uint32 = 42
                    var v21 int8 = 42
                    _ = v20
                    _ = v21
                }
                fmt.Println(v17)
                if (!false) {
                    var v22 float64 = (1.5 * (0.0 + 0.1))
                    if false {
                        fmt.Println(v17, v17)
                        v23 := v2["\n\t"]
                        _ = v23
                    }
                    _ = v22
                } else {
                    v19 = "\n\t"
                    var v24 bool = (2 < 18446744073709551615)
                    v25 := GenS1{F0: 32767}
                    _ = v24
                    _ = v25
                }
            } else {
                switch "" {
                case "":
                    _ = len(v2)
                case "a":
                    var v26 float64 = 1.5
                    _ = v26
                case "hello":
                    fmt.Println(v19)
                }
                v27 := GenS2{F0: 665}
                for fi28 := 0; fi28 < 7; fi28++ {
                    if (float32(3.14159) >= 1.5) {
                        if (false || (32767 == -963)) {
                            for fi29 := 0; fi29 < 10; fi29++ {
                                v30 := GenS3{F0: 997, F1: 2, F2: (!(!false))}
                                v3 -= v3
                                switch v17 {
                                case "hello":
                                    fmt.Println(v19, v3)
                                }
                                var v31 uint64 = 500
                                var v32 int16 = 1000
                                v19 = "hello"
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 15; fi33++ {
                                var v34 uint8 = 42
                                var v35 uint = 831
                                v4 += (v35 >> ((v35) & 63))
                                fmt.Println(v3, v3)
                                var v36 int64 = (v3 % (((v3 * (v3 * -9223372036854775808))) | 1))
                                _ = fi33
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            switch (("foo bar" + v19) + v17) {
                            case "hello":
                                for fi37 := 0; fi37 < 19; fi37++ {
                                    switch ((v19 + "a") + "a") {
                                    case "hello":
                                        v3 += genFn0()
                                    case "foo bar":
                                        v38 := GenS4{F0: fi28, F1: -550, F2: 2}
                                        _ = v38
                                    case "\n\t":
                                        _ = len(v2)
                                    default:
                                        var v39 float64 = (float64(((float32((((((-2.25 + 0.0) + -2.25) * 0.1) - -2.25) + 1e10)) - float32((1e10 + 1e10))) - (-2.25 / 1.5))) + 3.14159)
                                        if false {
                                            v1 = append(v1, "a")
                                            v19 = v19
                                            var v40 string = ("a" + "foo bar")
                                            _ = v40
                                        }
                                        _ = v39
                                    }
                                    v19 = ((v19 + "") + v19)
                                    v17 = ("foo bar" + (v19 + ((v19 + "foo bar") + (v19 + "\n\t"))))
                                    v41 := GenS3{F0: 4294967295, F1: 4294967295, F2: (2147483647 >= -1)}
                                    _ = fi37
                                    _ = v41
                                }
                            case "foo bar":
                                fmt.Println(v3)
                            case "":
                                var v42 int64 = int64(v4)
                                _ = v42
                            }
                        } else {
                            var v43 int16 = 0
                            _ = v43
                        }
                    }
                    if (!true) {
                        v17 = v17
                    } else {
                        var v44 int64 = (v3 - v3)
                        for fi45 := 0; fi45 < 5; fi45++ {
                            switch "a" {
                            case "foo bar":
                                v3 = -9223372036854775808
                            default:
                                var v46 uint8 = 128
                                _ = v46
                            }
                            v3 = (v3 + 0)
                            _ = fi45
                        }
                        for fi47 := 0; fi47 < 19; fi47++ {
                            var v48 string = ("a" + v19)
                            _ = fi47
                            _ = v48
                        }
                        _ = v44
                    }
                    v49 := GenS4{F0: (fi28 << ((4294967295) & 63)), F1: -1, F2: 1}
                    _ = fi28
                    _ = v49
                }
                _ = v27
            }
            v50 := v2["foo bar"]
            var v51 float32 = 0.1
            _ = v50
            _ = v51
        }
        _ = v19
    }
    if (65535 >= 254) {
        switch -1 {
        case -2147483648:
            v52 := GenS2{F0: 42}
            _ = v52
        }
    }
    switch 42 {
    case 42:
        var v53 bool = true
        _ = v53
    case -200:
        for fi54 := 0; fi54 < 16; fi54++ {
            for fi55 := 0; fi55 < 1; fi55++ {
                var v56 int32 = -898
                _ = fi55
                _ = v56
            }
            switch fi54 {
            case 776:
                v17 = v17
            case 32767:
                var v57 uint = v4
                _ = v57
            default:
                if false {
                    var v58 int8 = -128
                    _ = v58
                }
            }
            _ = fi54
        }
    default:
        v59 := GenS2{F0: 2}
        _ = v59
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        if false {
            v61 := GenS3{F0: 65535, F1: 2, F2: true}
            _ = v61
        }
        var v62 int32 = 0
        switch fi60 {
        case 127:
            fmt.Println(fi60, v17)
        case 0:
            if false {
                fmt.Println(v3, v17)
                var v63 string = ""
                _ = v63
            }
        }
        _ = fi60
        _ = v62
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        v3 = genFn0()
        _ = fi64
    }
    var v65 int = 777
    v1 = append(v1, (v17 + "hello"))
    v17 = v17
    v66 := GenS5{F0: (-2147483648 <= -2147483648), F1: 0, F2: (v3 | v3)}
    v3 = (v3 ^ genFn0())
    var v67 uint8 = 11
    fmt.Println(v4, v67)
    v4 = v4
    fmt.Println(v3, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
    _ = v67
}
