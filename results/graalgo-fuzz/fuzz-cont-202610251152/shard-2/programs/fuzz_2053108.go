package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 int; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 string }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: "a", F1: 0}
        if true {
            if true {
                if false {
                    for fi7 := 0; fi7 < 15; fi7++ {
                        var v8 int = fi7
                        v9 := []int{}
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    var v10 int64 = -1
                    var v11 uint64 = 842
                    _ = v10
                    _ = v11
                }
            }
        } else {
            fmt.Println(fi5)
        }
        _ = fi5
        _ = v6
    }
    var v12 uint64 = 2
    v12 = (v12 | v12)
    _ = v12
    return 32767
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (1000 != int16(v1))
    v3 := make(map[int]int)
    v4 := []int{127, 4294967295, -1}
    fmt.Println(v1)
    for fi13 := 0; fi13 < 1; fi13++ {
        switch "foo bar" {
        case "\n\t":
            v1 = v1
        case "":
            v1 += (v1 << ((2) & 63))
        default:
            v2 = v2
        }
        if (544 == 0) {
            var v14 uint64 = (v1 | (v1 | v1))
            _ = v14
        } else {
            v2 = (((42 >= 11) || (-1 == 442)) && (v2 && (v2 || (!v2))))
            v1 = (^(v1))
            if false {
                v2 = true
            }
        }
        v15 := GenS2{F0: fi13, F1: 521, F2: fi13}
        _ = fi13
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(fi16, v2)
        var v17 float64 = float64(0.0)
        _ = fi16
        _ = v17
    }
    _ = len(v4)
    for fi18 := 0; fi18 < 18; fi18++ {
        v2 = false
        _ = fi18
    }
    switch "\n\t" {
    case "a":
        for fi19 := 0; fi19 < 1; fi19++ {
            if (v1 == v1) {
                v2 = true
                var v20 uint8 = 54
                _ = v20
            } else {
                if v2 {
                    fmt.Println(fi19, fi19)
                    v1 = (v1 - v1)
                    switch ("\n\t" + ((("a" + "a") + ("\n\t" + ("" + "\n\t"))) + (("hello" + (("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + ("hello" + ("a" + "\n\t"))))) + "\n\t")) + (("hello" + "hello") + "foo bar")))) {
                    case "a":
                        fmt.Println(v2)
                    default:
                        v1 -= 2
                    }
                }
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 string = "a"
                if ((fi19 / ((-543) | 1)) < 128) {
                    v23 := GenS1{F0: (v22 + (v22 + "")), F1: 672}
                    _ = v23
                } else {
                    v3[(fi21 << ((0) & 63))] = (-(fi21))
                    var v24 float64 = float64(1.5)
                    _ = v24
                }
                v22 = "foo bar"
                if ("hello" > v22) {
                    v1 = v1
                    if v2 {
                        var v25 uint16 = 34
                        _ = v25
                    } else {
                        v2 = v2
                        v22 = "hello"
                    }
                    switch v22 {
                    case "":
                        if (false || true) {
                            v22 = v22
                            var v26 uint16 = 20
                            _ = v26
                        }
                    case "\n\t":
                        v2 = v2
                    case "a":
                        var v27 uint32 = 0
                        _ = v27
                    default:
                        if (true || v2) {
                            if v2 {
                                v22 = ("foo bar" + "foo bar")
                                var v28 uint8 = 255
                                _ = v28
                            } else {
                                _ = len(v4)
                            }
                        } else {
                            v29 := GenS3{F0: v1, F1: 9223372036854775807, F2: 1}
                            var v30 uint8 = 80
                            _ = v29
                            _ = v30
                        }
                    }
                    var v31 int16 = -192
                    if v2 {
                        v32 := GenS3{F0: 4294967295, F1: (fi21 | fi19), F2: 2147483648}
                        delete(v3, (fi21 << ((255) & 63)))
                        _ = v32
                    } else {
                        var v33 int8 = 127
                        _ = v33
                    }
                    _ = v31
                } else {
                    v34 := GenS4{F0: 42}
                    v22 = "foo bar"
                    v3[128] = (fi19 | fi19)
                    if v2 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v35 float64 = 0.1
                        if ((((v35 * (1.5 + v35)) / v35) != (float64(0.0) + v35)) && false) {
                            var v36 uint32 = 42
                            var v37 uint = 4294967295
                            _ = v36
                            _ = v37
                        }
                        _ = v35
                    }
                    _ = v34
                }
                fmt.Println(v22, v2, fi19)
                v22 = v22
                switch v22 {
                case "foo bar":
                    if v2 {
                        switch (^(fi21)) {
                        case 0:
                            v38 := GenS5{F0: 997, F1: 0}
                            _ = v38
                        case 185:
                            v39 := GenS5{F0: v1, F1: 4294967295}
                            _ = v39
                        case 811:
                            switch fi19 {
                            case 127:
                                v40 := GenS5{F0: (v1 * (v1 ^ 2)), F1: 1}
                                _ = v40
                            case 1:
                                if true {
                                    switch fi19 {
                                    case 1:
                                        switch 1 {
                                        case 304:
                                            if ((!v2) || false) {
                                                v41 := GenS6{F0: 0}
                                                fmt.Println(v2)
                                                var v42 string = ""
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                    case 127:
                                        switch ("" + v22) {
                                        case "hello":
                                            var v43 string = v22
                                            _ = v43
                                        }
                                    default:
                                        if false {
                                            v44 := v3[(fi21 ^ 32768)]
                                            v45 := GenS6{F0: 9223372036854775807}
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                }
                            case 531:
                                v46 := v3[(fi21 * (^(fi19)))]
                                _ = v46
                            }
                        }
                        fmt.Println(v2, v1, v2)
                        var v47 float64 = float64(1.5)
                        _ = v47
                    }
                    if (((-9223372036854775808 < fi21) && (0 > 1)) && (false || (fi19 < 2147483647))) {
                        var v48 int32 = -2147483648
                        delete(v3, 2147483647)
                        _ = v48
                    } else {
                        v49 := GenS7{F0: -1, F1: ("hello" + (v22 + (v22 + "a")))}
                        var v50 uint = 65535
                        v50 = (v50 ^ (v50 << ((4294967295) & 63)))
                        _ = v49
                        _ = v50
                    }
                    var v51 float64 = -2.25
                    _ = v51
                default:
                    v1 = v1
                }
                _ = fi21
                _ = v22
            }
            _ = fi19
        }
    case "hello":
        v1 += v1
    case "\n\t":
        for fi52 := 0; fi52 < 10; fi52++ {
            var v53 float32 = (0.0 / ((3.14159 / (1e10 - -2.25)) + 0.1))
            if v2 {
                var v54 uint64 = (v1 & (v1 & (^(v1))))
                v55 := GenS4{F0: -483}
                _ = v54
                _ = v55
            }
            _ = fi52
            _ = v53
        }
    default:
        var v56 uint8 = 127
        _ = v56
    }
    var v57 float64 = 1e10
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!false) {
        var v58 int32 = 1
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
