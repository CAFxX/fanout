package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint64; F2 bool }
type GenS8 struct { F0 string; F1 uint32; F2 int }

func main() {
    var v1 uint16 = 1
    v2 := []string{"a", "a", "\n\t", ("hello" + "foo bar")}
    v3 := []string{}
    var v4 float32 = (3.14159 + 3.14159)
    if (-1 != 1) {
        if true {
            v5 := GenS1{F0: 548, F1: 0}
            if true {
                var v6 uint16 = 1
                var v7 float32 = (((1e10 / v4) / (v4 - (v4 / 3.14159))) - v4)
                _ = v6
                _ = v7
            }
            _ = v5
        } else {
            var v8 float32 = 0.0
            v9 := GenS1{F0: 42, F1: uint(v1)}
            _ = v8
            _ = v9
        }
        v4 -= v4
    } else {
        v10 := GenS1{F0: 1, F1: 609}
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int8 = 127
            if ("foo bar" != "") {
                for fi13 := 0; fi13 < 12; fi13++ {
                    v12 -= 42
                    fmt.Println(fi11, v4)
                    var v14 uint = 4294967295
                    _ = fi13
                    _ = v14
                }
                v2 = append(v2, "foo bar")
                if ((!(781 >= (fi11 / ((fi11) | 1)))) || false) {
                    v4 -= float32((0.1 + 1.5))
                    v1 += 2
                } else {
                    fmt.Println(v12)
                    var v15 int32 = 42
                    v1 = 76
                    _ = v15
                }
            }
            _ = fi11
            _ = v12
        }
        fmt.Println(v1)
        _ = v10
    }
    v4 -= v4
    v2 = append(v2, "\n\t")
    v16 := GenS2{F0: 0}
    switch "\n\t" {
    case "a":
        v4 = -2.25
    default:
        v4 -= v4
    }
    fmt.Println(v4, v4)
    for fi17 := 0; fi17 < 17; fi17++ {
        switch ("hello" + ("\n\t" + "hello")) {
        case "a":
            if false {
                switch fi17 {
                case 783:
                    v4 = ((0.1 + -2.25) + v4)
                case -2147483648:
                    v4 = float32(1e10)
                case -811:
                    v18 := GenS1{F0: -861, F1: 962}
                    _ = v18
                }
            } else {
                var v19 int8 = 37
                var v20 int64 = int64(v19)
                _ = v19
                _ = v20
            }
        case "\n\t":
            for fi21 := 0; fi21 < 10; fi21++ {
                switch fi17 {
                case 128:
                    if ((int(v1) > fi21) && (127 >= -1)) {
                        if false {
                            var v22 int8 = 0
                            if (549 >= -724) {
                                v1 += 2
                                var v23 uint = 4294967295
                                _ = v23
                            } else {
                                if false {
                                    v1 -= v1
                                    fmt.Println(fi21, fi21, v4)
                                } else {
                                    var v24 uint = 2
                                    v25 := GenS2{F0: -151}
                                    _ = v24
                                    _ = v25
                                }
                            }
                            _ = v22
                        }
                        _ = len(v3)
                        v4 -= ((0.1 * float32(1e10)) / v4)
                    } else {
                        v4 = v4
                        var v26 int64 = -1
                        v27 := GenS2{F0: -193}
                        if ((128 >= 34) || ((18446744073709551615 < 4294967295) || (2147483647 == 65535))) {
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            switch (fi21 + fi21) {
                            case -9223372036854775808:
                                switch 42 {
                                case -1:
                                    if true {
                                        fmt.Println(v26, v1)
                                        fmt.Println(fi17, v26)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        var v28 int16 = 42
                                        fmt.Println(v1, fi17, v26)
                                        switch "" {
                                        case "a":
                                            if (true || (!false)) {
                                                if len(v2) > 0 {
                                                    v2[0] = ("a" + (("" + "\n\t") + "\n\t"))
                                                }
                                                var v29 string = ("a" + "foo bar")
                                                if ((!false) && (0 > 159)) {
                                                    var v30 string = (v29 + ((("hello" + v29) + "\n\t") + ""))
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        fmt.Println(fi21, fi17)
                                                    case "foo bar":
                                                        if (!((42 <= v28) && true)) {
                                                            v31 := GenS3{F0: -1, F1: (fi17 << ((42) & 63))}
                                                            _ = v31
                                                        } else {
                                                            v29 = "a"
                                                            v29 = (v30 + v30)
                                                            var v32 float64 = 0.0
                                                            _ = v32
                                                        }
                                                    }
                                                    v29 = v30
                                                    _ = v30
                                                } else {
                                                    fmt.Println(v29, v29)
                                                    if false {
                                                        v33 := GenS4{F0: -957, F1: fi21}
                                                        _ = v33
                                                    } else {
                                                        v4 = (v4 / v4)
                                                        var v34 uint = 0
                                                        _ = v34
                                                    }
                                                    if false {
                                                        v35 := GenS5{F0: v29}
                                                        v36 := GenS6{F0: (-(fi17))}
                                                        v37 := GenS4{F0: -703, F1: fi17}
                                                        _ = v35
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                }
                                                _ = v29
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case "foo bar":
                                            var v38 bool = true
                                            _ = v38
                                        }
                                        _ = v28
                                    }
                                case 127:
                                    fmt.Println(fi21, fi17, v26)
                                    switch -2147483648 {
                                    case 32767:
                                        var v39 uint = 2
                                        _ = v39
                                    case 2147483647:
                                        v26 = 902
                                        v2 = append(v2, ("hello" + ""))
                                        fmt.Println(fi17, fi21, fi17)
                                    default:
                                        var v40 int = (fi21 | -9223372036854775808)
                                        _ = v40
                                    }
                                case -187:
                                    v1 -= v1
                                }
                            }
                            var v41 uint = 65535
                            _ = v41
                        } else {
                            if true {
                                v4 += 1.5
                            }
                            var v42 int16 = -602
                            _ = v42
                        }
                        var v43 uint8 = 0
                        _ = v26
                        _ = v27
                        _ = v43
                    }
                case 32767:
                    v4 += v4
                default:
                    v44 := GenS7{F0: 368, F1: 196, F2: (2147483647 < -1)}
                    _ = v44
                }
                v45 := GenS7{F0: 0, F1: 18446744073709551615, F2: false}
                _ = fi21
                _ = v45
            }
        case "hello":
            v4 = float32(((1.5 * -2.25) + 1e10))
        }
        _ = fi17
    }
    var v46 int32 = 0
    for fi47 := 0; fi47 < 1; fi47++ {
        var v48 int = fi47
        _ = fi47
        _ = v48
    }
    var v49 int = -1
    if ((-67 >= 122) && ((100 >= -1) || false)) {
        if len(v3) > 0 {
            v3[0] = ((("a" + "foo bar") + "foo bar") + ((("a" + (("" + "\n\t") + "foo bar")) + "a") + "foo bar"))
        }
    }
    if len(v2) > 0 {
        v2[0] = (("foo bar" + ("hello" + "hello")) + "\n\t")
    }
    var v50 float64 = (0.1 / (0.1 - 0.0))
    switch "a" {
    case "\n\t":
        for fi51 := 0; fi51 < 17; fi51++ {
            v52 := GenS1{F0: 0, F1: 1}
            if false {
                switch 128 {
                case 127:
                    v53 := GenS8{F0: "", F1: 2147483648, F2: (fi51 + v49)}
                    _ = v53
                case 449:
                    switch "a" {
                    case "a":
                        v54 := GenS6{F0: fi51}
                        _ = v54
                    }
                case 9223372036854775807:
                    var v55 bool = true
                    _ = v55
                }
                var v56 string = "\n\t"
                _ = v56
            } else {
                var v57 uint32 = 2147483648
                _ = v57
            }
            v3 = append(v3, ("hello" + ("a" + ("" + "foo bar"))))
            _ = fi51
            _ = v52
        }
    case "":
        _ = len(v3)
    case "foo bar":
        for fi58 := 0; fi58 < 1; fi58++ {
            if false {
                fmt.Println(fi58, v49, v1)
            } else {
                _ = len(v3)
            }
            v3 = append(v3, ("hello" + "\n\t"))
            v46 = v46
            _ = fi58
        }
    }
    if true {
        var v59 int8 = -50
        if (!(!true)) {
            for fi60 := 0; fi60 < 19; fi60++ {
                v61 := GenS5{F0: ""}
                if false {
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "hello")
                    }
                }
                _ = fi60
                _ = v61
            }
        } else {
            var v62 int16 = 881
            _ = v62
        }
        v49 += v49
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v49
    _ = v50
}
