package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int16; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 int8) bool {
    v6 := GenS1{F0: 4294967295, F1: -914, F2: 42}
    p0 = 127
    v7 := GenS1{F0: -454, F1: int16(p0), F2: 42}
    switch "\n\t" {
    case "hello":
        v8 := GenS2{F0: 665, F1: -9223372036854775808}
        _ = v8
    case "\n\t":
        switch 32767 {
        case 127:
            v9 := []int{4294967295, 128, int(p0)}
            _ = v9
        case 32767:
            var v10 int = 1
            _ = v10
        }
    default:
        p0 += p0
    }
    _ = p0
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 string) uint32 {
    var v11 float64 = 3.14159
    fmt.Println(p0, p0)
    switch 0 {
    case 127:
        if false {
            v12 := GenS2{F0: 2, F1: 1}
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{873, fi13}
            v11 = (v11 / v11)
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v11, v11)
    var v15 float32 = (((1.5 - float32(v11)) - 1e10) - 1.5)
    _ = p0
    _ = v11
    _ = v15
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{(("" + "a") + ""), "a", ("" + "hello"), ""}
    v3 := make(map[string]uint32)
    v4 := []int{4294967295, 128}
    var v5 uint = 1
    if false {
        switch (("" + "a") + "foo bar") {
        case "hello":
            v5 = v5
        case "\n\t":
            v5 = v5
        case "foo bar":
            v16 := GenS3{F0: 766, F1: -589}
            _ = v16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(fi17)
            v18 := GenS4{F0: -32768, F1: 1, F2: 1}
            v5 -= 18446744073709551615
            _ = fi17
            _ = v18
        }
    } else {
        var v19 float64 = -2.25
        switch "" {
        case "a":
            var v20 uint64 = 0
            _ = v20
        case "hello":
            for fi21 := 0; fi21 < 1; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 float32 = ((((float32(0.1) + -2.25) * 1e10) - (0.0 + 0.1)) - -2.25)
                    _ = fi22
                    _ = v23
                }
                var v24 uint32 = 2147483648
                for fi25 := 0; fi25 < 4; fi25++ {
                    fmt.Println(fi25)
                    v26 := v3["a"]
                    _ = fi25
                    _ = v26
                }
                _ = fi21
                _ = v24
            }
        case "":
            if (0 > 2) {
                for fi27 := 0; fi27 < 16; fi27++ {
                    if ("a" > "hello") {
                        v19 = 0.0
                    }
                    v19 = v19
                    var v28 uint = v5
                    _ = fi27
                    _ = v28
                }
                if true {
                    _ = len(v2)
                    v19 -= (((1.5 - (float64(1.5) + v19)) + v19) - (v19 - (v19 * v19)))
                } else {
                    for fi29 := 0; fi29 < 15; fi29++ {
                        _ = len(v1)
                        _ = fi29
                    }
                    delete(v1, (("foo bar" + ("" + "\n\t")) + "\n\t"))
                }
                _ = len(v2)
            }
        }
        _ = v19
    }
    fmt.Println(v5)
    v30 := GenS5{F0: 42, F1: 42}
    fmt.Println(v5)
    fmt.Println(v5)
    v5 = v5
    for fi31 := 0; fi31 < 11; fi31++ {
        v5 += 802
        var v32 uint16 = 177
        switch "" {
        case "hello":
            var v33 uint8 = 188
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 8; fi34++ {
                fmt.Println(v32, v5, fi31)
                if false {
                    v1[""] = (fi31 & 2147483647)
                    var v35 bool = (!true)
                    switch ("a" + "a") {
                    case "\n\t":
                        if (uint8(v5) < uint8(v5)) {
                            var v36 int64 = 0
                            v5 += (v5 + (v5 % ((4294967295) | 1)))
                            _ = v36
                        } else {
                            fmt.Println(fi34)
                            var v37 uint = 255
                            var v38 int = fi34
                            _ = v37
                            _ = v38
                        }
                    case "foo bar":
                        if v35 {
                            v2 = append(v2, "\n\t")
                            v39 := GenS5{F0: -935, F1: 4294967295}
                            _ = v39
                        } else {
                            if false {
                                var v40 bool = (v35 || v35)
                                fmt.Println(fi31)
                                _ = v40
                            } else {
                                fmt.Println(v5, v5)
                                fmt.Println(fi34)
                            }
                            v35 = (uint64(v32) >= 0)
                            v32 += (v32 ^ (v32 << (((v5 / ((v5) | 1))) & 63)))
                        }
                    }
                    _ = v35
                }
                switch (fi31 / (((fi31 ^ fi31)) | 1)) {
                case 665:
                    var v41 uint8 = 42
                    _ = v41
                case -2147483648:
                    v42 := GenS5{F0: -32768, F1: 18446744073709551615}
                    _ = v42
                }
                _ = fi34
            }
        case "a":
            v43 := GenS2{F0: genFn1(("" + "a")), F1: 971}
            _ = v43
        default:
            var v44 int64 = 9223372036854775807
            _ = v44
        }
        _ = fi31
        _ = v32
    }
    v45 := v1["\n\t"]
    v45 = v45
    var v46 uint8 = 136
    v5 = v5
    var v47 string = ""
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 int = (fi48 % (((fi48 << (((v5 & 255)) & 63))) | 1))
        var v50 string = ("hello" + "")
        _ = fi48
        _ = v49
        _ = v50
    }
    v51 := GenS1{F0: (v45 ^ v45), F1: 0, F2: 2147483647}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v52 bool = (127 >= 50)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v51
    _ = v52
}
