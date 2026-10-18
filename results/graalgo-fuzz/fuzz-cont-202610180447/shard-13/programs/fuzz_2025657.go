package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 int; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 int }
func genFn0(p0 int) int {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int = (p0 * -722)
        var v8 int16 = 1000
        var v9 bool = true
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 229, F1: "hello"}
    _ = p0
    _ = v10
    return 2147483647
}
func genFn1() uint16 {
    v11 := GenS1{F0: 128, F1: "hello"}
    var v12 string = ("\n\t" + "\n\t")
    if (v12 < "\n\t") {
        v13 := GenS1{F0: 2, F1: v12}
        v12 = "foo bar"
        _ = v13
    } else {
        v14 := []int{42, -1, -1, 0}
        _ = v14
    }
    switch v12 {
    case "":
        fmt.Println(v12)
    case "a":
        for fi15 := 0; fi15 < 20; fi15++ {
            switch (fi15 - fi15) {
            case 4294967295:
                var v16 int16 = -990
                _ = v16
            }
            v17 := []int{9223372036854775807, (fi15 & genFn0((^(fi15))))}
            _ = fi15
            _ = v17
        }
    case "hello":
        v18 := []int{}
        _ = v18
    }
    fmt.Println(v12)
    _ = v11
    _ = v12
    return 59
}

func main() {
    v1 := []string{"foo bar", "foo bar", "\n\t"}
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    var v5 uint8 = 1
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 243 {
        case -734:
            if (((("\n\t" + (("foo bar" + "") + "\n\t")) + ("foo bar" + "")) + "") <= "\n\t") {
                switch "a" {
                case "\n\t":
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v4 -= v4
                        var v20 int8 = 34
                        _ = fi19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: v5, F1: "\n\t"}
                var v22 uint16 = 389
                _ = v21
                _ = v22
            } else {
                v5 += (v5 ^ v5)
                var v23 bool = (true && (!true))
                _ = v23
            }
        case 32767:
            if true {
                if (true && true) {
                    if false {
                        var v24 int8 = 100
                        _ = v24
                    } else {
                        switch "" {
                        case "\n\t":
                            if (57 <= 66) {
                                v25 := GenS1{F0: (v5 / ((v5) | 1)), F1: ""}
                                v5 += 1
                                _ = v25
                            } else {
                                v26 := GenS1{F0: 42, F1: "foo bar"}
                                _ = v26
                            }
                        case "":
                            fmt.Println(v4, v2)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v4 -= v4
                    }
                    switch ("hello" + ("foo bar" + "")) {
                    case "foo bar":
                        v27 := GenS2{F0: 11, F1: 32767, F2: -1}
                        _ = v27
                    }
                    v28 := GenS3{F0: 255}
                    _ = v28
                } else {
                    _ = len(v3)
                    if false {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi29
                        }
                        for fi30 := 0; fi30 < 7; fi30++ {
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                            var v31 float64 = -2.25
                            var v32 string = "foo bar"
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                }
                _ = len(v3)
                var v33 string = (("foo bar" + ("foo bar" + (("hello" + "") + "foo bar"))) + "\n\t")
                _ = v33
            }
        case -1:
            var v34 int16 = -32768
            _ = v34
        }
    case 0:
        v4 -= v4
    case -1:
        v5 -= (^(v5))
    default:
        v35 := GenS2{F0: 32767, F1: -133, F2: 42}
        _ = v35
    }
    var v36 float64 = 3.14159
    var v37 float64 = v36
    if true {
        switch 32768 {
        case -2147483648:
            v36 += 3.14159
        case 9223372036854775807:
            v36 = -2.25
        }
        _ = len(v1)
        for fi38 := 0; fi38 < 8; fi38++ {
            v39 := GenS4{F0: 255, F1: (^(v4))}
            _ = fi38
            _ = v39
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v40 int16 = -957
    v3[975] = true
    if (true || (false && ((("" + "hello") != ("" + "hello")) || ((v2 % ((int64(v4)) | 1)) > v2)))) {
        v2 -= v2
        var v41 int64 = (v2 ^ 1)
        _ = v41
    } else {
        v42 := GenS5{F0: (v4 - v4), F1: ("foo bar" + ("\n\t" + "a")), F2: -9223372036854775808}
        _ = len(v3)
        var v43 string = "a"
        _ = v42
        _ = v43
    }
    if ((!false) && (42 <= 42)) {
        v40 -= v40
        v4 = v4
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        v40 = 42
        v36 = v36
        if false {
            fmt.Println(v36, v36)
        }
        _ = fi44
    }
    _ = len(v3)
    fmt.Println(v40, v2, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
