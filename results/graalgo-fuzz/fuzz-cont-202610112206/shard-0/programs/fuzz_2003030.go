package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 bool }
func genFn0(p0 int) int8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        if (1 <= 288) {
            v7 := []int{32768}
            switch ("hello" + ("hello" + "")) {
            case "hello":
                for fi8 := 0; fi8 < 8; fi8++ {
                    if true {
                        fmt.Println(fi8, fi8, fi8)
                        p0 = fi8
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = fi8
                }
            default:
                var v9 int16 = 33
                _ = v9
            }
            var v10 int = -9223372036854775808
            _ = v10
            _ = v7
        }
        var v11 int16 = -944
        _ = fi6
        _ = v11
    }
    if (42 < 214) {
        for fi12 := 0; fi12 < 12; fi12++ {
            p0 += fi12
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(fi13, fi13)
                if true {
                    v14 := []int{-9223372036854775808, (fi13 + -1)}
                    _ = v14
                } else {
                    v15 := []int{}
                    _ = v15
                }
                v16 := GenS1{F0: 57, F1: 2, F2: 895}
                _ = fi13
                _ = v16
            }
            switch "a" {
            case "hello":
                var v17 string = "foo bar"
                _ = v17
            }
            _ = fi12
        }
        p0 = 1
        v18 := []int{(p0 - (p0 & p0))}
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        if ((65535 <= 2) && ((!false) && true)) {
            p0 -= 32767
        }
        var v20 int = fi19
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        if (!false) {
            v22 := GenS1{F0: -1, F1: 127, F2: 1}
            _ = v22
        } else {
            fmt.Println(p0, p0)
        }
        v23 := GenS1{F0: 2147483647, F1: 1, F2: 726}
        if false {
            var v24 int32 = -2147483648
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS2{F0: -32768, F1: true, F2: fi21}
                v27 := []int{p0, fi25, (^(fi21))}
                v28 := GenS2{F0: 0, F1: ("hello" <= ("a" + "hello")), F2: (fi21 / ((-1) | 1))}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            _ = v24
        }
        _ = fi21
        _ = v23
    }
    switch ((("\n\t" + "foo bar") + "a") + "hello") {
    case "":
        if (!false) {
            fmt.Println(p0)
            p0 += p0
            switch "hello" {
            case "\n\t":
                for fi29 := 0; fi29 < 9; fi29++ {
                    var v30 uint8 = 115
                    _ = fi29
                    _ = v30
                }
            case "hello":
                var v31 float32 = 1.5
                _ = v31
            case "foo bar":
                var v32 uint = 4294967295
                _ = v32
            }
        } else {
            v33 := []int{(p0 - (p0 >> ((1) & 63))), p0, p0, p0}
            _ = v33
        }
    case "hello":
        v34 := []int{32767, 42, (p0 << ((1) & 63))}
        _ = v34
    case "a":
        if true {
            if (("hello" > "foo bar") || (42 == 255)) {
                if (42 == p0) {
                    p0 = 0
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 string = "hello"
                        v37 := GenS2{F0: -1, F1: false, F2: p0}
                        fmt.Println(fi35, p0)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    var v38 int = 741
                    _ = v38
                }
            } else {
                fmt.Println(p0)
                fmt.Println(p0)
            }
            switch ("foo bar" + "a") {
            case "a":
                var v39 bool = true
                _ = v39
            case "foo bar":
                v40 := []int{(p0 | (^(p0))), -2147483648, p0, 66}
                _ = v40
            }
        } else {
            fmt.Println(p0)
            var v41 uint32 = 1
            _ = v41
        }
    default:
        v42 := GenS1{F0: 600, F1: 128, F2: -1}
        _ = v42
    }
    _ = p0
    return -1
}
func genFn1() int16 {
    switch ("\n\t" + "") {
    case "\n\t":
        v43 := []int{793, 1, 9223372036854775807, 312}
        _ = v43
    case "":
        v44 := GenS3{F0: -32768}
        _ = v44
    default:
        var v45 int8 = 100
        _ = v45
    }
    var v46 int16 = 32767
    v47 := []int{787}
    _ = v46
    _ = v47
    return (v46 % (((v46 * v46)) | 1))
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = 0.0
    var v3 bool = false
    var v4 string = ("a" + "\n\t")
    var v5 int64 = 0
    v4 = (v4 + v4)
    switch -2147483648 {
    case 2147483647:
        v48 := GenS3{F0: 1000}
        _ = v48
    case 127:
        var v49 int64 = 306
        _ = v49
    case -2147483648:
        if (!((^(v5)) > 491)) {
            v3 = (!v3)
        }
    default:
        v50 := GenS4{F0: "foo bar", F1: v3}
        _ = v50
    }
    if v3 {
        v3 = (v5 > (v5 % ((int64(v1)) | 1)))
    }
    v51 := []int{-14, 4294967295}
    v52 := GenS2{F0: 1, F1: (32767 < -32768), F2: 380}
    var v53 int = int(v1)
    fmt.Println(v5, v53)
    fmt.Println(v3)
    if len(v51) >= 2 {
        v51 = v51[1:len(v51)]
    }
    if v3 {
        var v54 int8 = 127
        for fi55 := 0; fi55 < 10; fi55++ {
            if false {
                for fi56 := 0; fi56 < 4; fi56++ {
                    var v57 int = (-(v53))
                    v54 += 0
                    _ = fi56
                    _ = v57
                }
            } else {
                fmt.Println(v54)
                _ = len(v51)
            }
            v58 := GenS3{F0: -728}
            v51 = append(v51, (fi55 ^ (^(fi55))))
            _ = fi55
            _ = v58
        }
        for fi59 := 0; fi59 < 20; fi59++ {
            switch (fi59 ^ fi59) {
            case 1:
                v60 := GenS3{F0: 1000}
                _ = v60
            case 127:
                v3 = (!v3)
            case -2147483648:
                for fi61 := 0; fi61 < 11; fi61++ {
                    fmt.Println(v5)
                    switch (fi61 >> ((42) & 63)) {
                    case 9223372036854775807:
                        if true {
                            var v62 uint8 = 0
                            v63 := GenS4{F0: (v4 + ("foo bar" + v4)), F1: false}
                            v4 = "hello"
                            _ = v62
                            _ = v63
                        } else {
                            var v64 int = -9223372036854775808
                            _ = v64
                        }
                    case 128:
                        var v65 string = "\n\t"
                        _ = v65
                    case -948:
                        v54 = v54
                    }
                    v66 := GenS1{F0: int32(v1), F1: 128, F2: -1}
                    _ = fi61
                    _ = v66
                }
            }
            var v67 float32 = 3.14159
            if v3 {
                var v68 uint = 0
                var v69 int8 = v54
                _ = v68
                _ = v69
            } else {
                if (!(1 <= 127)) {
                    var v70 string = v4
                    _ = v70
                }
            }
            _ = fi59
            _ = v67
        }
        _ = v54
    }
    for fi71 := 0; fi71 < 7; fi71++ {
        var v72 uint16 = 0
        _ = fi71
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
