package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0() int32 {
    v4 := []int{848, -703}
    var v5 string = ""
    _ = v4
    _ = v5
    return -795
}
func genFn1(p0 uint8) int64 {
    if true {
        v6 := []int{127, 128, 4294967295}
        v6 = append(v6, 4294967295)
        _ = v6
    }
    var v7 bool = true
    var v8 int32 = 505
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 1
    var v3 string = ""
    for fi9 := 0; fi9 < 19; fi9++ {
        v3 = ((((v3 + (("hello" + "foo bar") + (v3 + "foo bar"))) + (v3 + (("" + "hello") + "hello"))) + v3) + v3)
        var v10 uint = 42
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int8 = v1
        _ = fi11
        _ = v12
    }
    switch v3 {
    case "":
        v1 = (^(v1))
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                if (true || (true && true)) {
                    var v15 bool = (true || (v3 > v3))
                    if (0 >= 2147483647) {
                        v3 = v3
                        fmt.Println(fi13)
                    } else {
                        v16 := GenS1{F0: 178, F1: 2147483648, F2: v15}
                        v1 = v1
                        switch fi14 {
                        case 9223372036854775807:
                            var v17 int16 = 457
                            _ = v17
                        }
                        v18 := []int{(fi13 % ((42) | 1)), fi14}
                        v15 = true
                        _ = v16
                        _ = v18
                    }
                    _ = v15
                }
                v19 := []int{fi13}
                _ = fi14
                _ = v19
            }
            _ = fi13
        }
    }
    v3 = (("a" + ("hello" + "\n\t")) + (v3 + "hello"))
    var v20 int16 = 32767
    var v21 bool = true
    v22 := []int{128, -1}
    fmt.Println(v1, v3)
    switch "" {
    case "a":
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
    case "":
        var v23 string = ((v3 + (("" + ((v3 + ("\n\t" + v3)) + v3)) + v3)) + "foo bar")
        _ = v23
    }
    if (0.1 < 0.0) {
        for fi24 := 0; fi24 < 18; fi24++ {
            switch "" {
            case "foo bar":
                for fi25 := 0; fi25 < 7; fi25++ {
                    if len(v22) > 0 {
                        v22[0] = fi24
                    }
                    _ = fi25
                }
            }
            v2 -= (v2 + v2)
            _ = fi24
        }
    } else {
        var v26 int8 = v1
        _ = v26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        switch "" {
        case "\n\t":
            if false {
                fmt.Println(v1, v21)
                var v28 uint8 = 42
                fmt.Println(v28, v1)
                _ = v28
            } else {
                v1 += (v1 - v1)
                switch v3 {
                case "foo bar":
                    var v29 string = v3
                    _ = v29
                case "a":
                    fmt.Println(fi27, v1)
                default:
                    var v30 uint = 929
                    _ = v30
                }
            }
        case "hello":
            for fi31 := 0; fi31 < 9; fi31++ {
                var v32 uint32 = 65535
                _ = fi31
                _ = v32
            }
        case "foo bar":
            if true {
                v2 -= (v2 | v2)
                if len(v22) > 0 {
                    v22[0] = 32768
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    v20 -= 493
                    fmt.Println(v2, v20, v21)
                    var v34 int32 = -2147483648
                    _ = fi33
                    _ = v34
                }
            } else {
                _ = len(v22)
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
        default:
            var v35 int32 = -474
            _ = v35
        }
        v36 := GenS1{F0: 2, F1: 92, F2: false}
        switch fi27 {
        case 988:
            fmt.Println(fi27, fi27)
        case 2147483647:
            if (!true) {
                var v37 uint = 255
                for fi38 := 0; fi38 < 17; fi38++ {
                    v22 = append(v22, fi38)
                    _ = len(v22)
                    _ = fi38
                }
                switch v3 {
                case "foo bar":
                    for fi39 := 0; fi39 < 6; fi39++ {
                        switch (fi39 * 2147483647) {
                        case 646:
                            _ = len(v22)
                        case 42:
                            v40 := GenS1{F0: 0, F1: 2, F2: v21}
                            _ = v40
                        }
                        _ = fi39
                    }
                default:
                    v41 := GenS1{F0: 128, F1: 1, F2: v21}
                    _ = v41
                }
                _ = v37
            }
        }
        _ = fi27
        _ = v36
    }
    fmt.Println(v1, v1, v2)
    v42 := GenS1{F0: 164, F1: 2, F2: v21}
    v22 = append(v22, 0)
    for fi43 := 0; fi43 < 15; fi43++ {
        switch (v3 + ((v3 + v3) + v3)) {
        case "\n\t":
            v44 := GenS1{F0: 2, F1: 1, F2: v21}
            _ = v44
        case "hello":
            v20 = v20
        default:
            v45 := GenS1{F0: 1, F1: 65535, F2: (1.5 >= 0.1)}
            _ = v45
        }
        if (!(!v21)) {
            var v46 int16 = -32768
            if (false && (!true)) {
                if ((v3 >= v3) && (383 == 2147483647)) {
                    v22 = append(v22, (fi43 - (fi43 ^ (^(fi43)))))
                    v22 = append(v22, -556)
                    v47 := GenS1{F0: 1, F1: 763, F2: (65535 > 0)}
                    _ = v47
                }
                v3 = v3
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            } else {
                fmt.Println(v3)
                var v48 int64 = 0
                _ = v48
            }
            var v49 uint64 = 1
            _ = v46
            _ = v49
        } else {
            v50 := GenS2{F0: v3, F1: -381, F2: (v3 + v3)}
            v51 := GenS3{F0: 1, F1: "hello"}
            for fi52 := 0; fi52 < 15; fi52++ {
                v53 := GenS1{F0: 66, F1: 699, F2: (75 <= -1)}
                switch ("foo bar" + (("a" + ((v3 + (v3 + v3)) + ((v3 + "hello") + (v3 + "hello")))) + v3)) {
                case "\n\t":
                    v3 = v3
                case "":
                    var v54 uint64 = 4294967295
                    _ = v54
                case "foo bar":
                    v55 := GenS4{F0: v21, F1: 72, F2: (v1 & (v1 / (((v1 - 100)) | 1)))}
                    _ = v55
                default:
                    switch v3 {
                    case "foo bar":
                        v56 := GenS4{F0: v21, F1: 2, F2: (v1 ^ -1)}
                        _ = v56
                    }
                }
                var v57 int32 = 42
                _ = fi52
                _ = v53
                _ = v57
            }
            _ = v50
            _ = v51
        }
        var v58 float64 = (0.1 / 3.14159)
        _ = fi43
        _ = v58
    }
    var v59 uint16 = 255
    fmt.Println(v20, v21, v59)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v42
    _ = v59
}
