package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 int32) uint8 {
    v6 := GenS1{F0: p1, F1: "foo bar", F2: 187}
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: (-(p1)), F1: "a", F2: 2147483647}
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int64 = 9223372036854775807
            p1 -= p1
            _ = fi9
            _ = v10
        }
        v11 := []int{}
        _ = fi7
        _ = v11
        _ = v8
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 string = (("hello" + "") + ("hello" + ""))
        switch ((v13 + v13) + (v13 + v13)) {
        case "a":
            v14 := GenS1{F0: p1, F1: v13, F2: 4294967295}
            _ = v14
        case "foo bar":
            var v15 string = (v13 + "\n\t")
            _ = v15
        default:
            switch ((v13 + v13) + ("foo bar" + "\n\t")) {
            case "":
                if (4294967295 < 2) {
                    v16 := []int{(fi12 & (fi12 / ((2147483647) | 1))), int(p0)}
                    var v17 bool = (4294967295 >= 689)
                    v18 := GenS2{F0: p1, F1: v13}
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case "\n\t":
                if ((^(fi12)) > (fi12 & fi12)) {
                    if true {
                        var v19 float32 = 0.1
                        _ = v19
                    }
                } else {
                    var v20 int64 = -9223372036854775808
                    v21 := []int{}
                    _ = v20
                    _ = v21
                }
            case "foo bar":
                switch -2147483648 {
                case -1:
                    switch v13 {
                    case "foo bar":
                        p1 -= -1
                    case "hello":
                        v13 = ("" + "")
                    case "a":
                        if false {
                            var v22 int64 = 42
                            _ = v22
                        }
                    }
                case 42:
                    v13 = ("\n\t" + "")
                case 128:
                    v23 := GenS2{F0: p1, F1: v13}
                    _ = v23
                default:
                    p0 += (p0 >> ((2) & 63))
                }
            default:
                v24 := []int{(-(fi12)), (fi12 - fi12), fi12, (fi12 % (((fi12 + 474)) | 1))}
                _ = v24
            }
        }
        switch -2147483648 {
        case -2147483648:
            var v25 uint16 = p0
            _ = v25
        default:
            if false {
                for fi26 := 0; fi26 < 7; fi26++ {
                    switch v13 {
                    case "\n\t":
                        p1 = (p1 + (p1 | p1))
                        if false {
                            v13 = "foo bar"
                            v27 := []int{fi12, (fi12 ^ 32767)}
                            _ = v27
                        }
                        fmt.Println(fi26, fi26)
                        p0 = p0
                    case "":
                        var v28 uint16 = (p0 * p0)
                        v29 := []int{(fi26 << ((357) & 63)), (fi12 % ((fi12) | 1)), fi26, 32768}
                        _ = v28
                        _ = v29
                    }
                    switch v13 {
                    case "hello":
                        fmt.Println(p1, p1, v13)
                    case "foo bar":
                        p1 += (p1 + -909)
                    default:
                        v30 := GenS2{F0: -2147483648, F1: "foo bar"}
                        _ = v30
                    }
                    v31 := []int{524}
                    _ = fi26
                    _ = v31
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    p0 = (^(p0))
                    p1 = 0
                    var v33 uint16 = (p0 / ((0) | 1))
                    _ = fi32
                    _ = v33
                }
                v34 := []int{(fi12 - fi12), (fi12 >> ((566) & 63)), (-(fi12))}
                _ = v34
            }
        }
        _ = fi12
        _ = v13
    }
    p1 = (-(p1))
    var v35 float32 = ((3.14159 - (-2.25 + (-2.25 - 0.0))) + (1.5 - 1.5))
    _ = p0
    _ = p1
    _ = v35
    _ = v6
    return 100
}
func genFn1(p0 uint64) uint32 {
    switch 2147483647 {
    case 42:
        v36 := []int{2147483647}
        _ = v36
    case 341:
        fmt.Println(p0)
    default:
        for fi37 := 0; fi37 < 20; fi37++ {
            for fi38 := 0; fi38 < 5; fi38++ {
                fmt.Println(p0, p0)
                p0 += p0
                var v39 int16 = -443
                _ = fi38
                _ = v39
            }
            var v40 int = -553
            _ = fi37
            _ = v40
        }
    }
    var v41 string = ""
    v42 := GenS3{F0: 4294967295}
    var v43 uint = 537
    for fi44 := 0; fi44 < 2; fi44++ {
        v43 = 255
        _ = fi44
    }
    _ = p0
    _ = v41
    _ = v42
    _ = v43
    return 65535
}

func main() {
    var v1 float64 = ((float64(1e10) * (1e10 * ((0.0 * (1.5 / 1e10)) - 1e10))) - float64(0.1))
    var v2 uint32 = 712
    var v3 uint64 = 1
    var v4 uint64 = 1
    var v5 uint64 = 18446744073709551615
    v45 := []int{1}
    switch "" {
    case "a":
        v3 += v3
    case "\n\t":
        v46 := GenS3{F0: uint(v4)}
        _ = v46
    }
    fmt.Println(v1, v2, v4)
    v47 := GenS1{F0: -2147483648, F1: "foo bar", F2: v2}
    if (!true) {
        v48 := GenS2{F0: -991, F1: (((("a" + ("foo bar" + "\n\t")) + "") + "hello") + "")}
        _ = v48
    }
    switch -2147483648 {
    case 32767:
        for fi49 := 0; fi49 < 19; fi49++ {
            v50 := GenS1{F0: -829, F1: "a", F2: v2}
            _ = fi49
            _ = v50
        }
    case -954:
        var v51 uint = 337
        _ = v51
    case -939:
        var v52 uint8 = 0
        _ = v52
    default:
        var v53 uint64 = (v3 - v5)
        _ = v53
    }
    var v54 uint = 532
    v55 := GenS3{F0: v54}
    for fi56 := 0; fi56 < 2; fi56++ {
        fmt.Println(fi56, fi56, v5)
        if ((!true) && true) {
            if len(v45) >= 2 {
                v45 = v45[1:len(v45)]
            }
        }
        _ = fi56
    }
    var v57 bool = (-1 >= 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v5
    _ = v54
    _ = v55
    _ = v57
}
