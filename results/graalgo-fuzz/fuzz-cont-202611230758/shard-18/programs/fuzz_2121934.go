package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int8 }

func main() {
    var v1 int32 = -1
    var v2 string = ""
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    v5 := []int{0, 4294967295}
    v6 := GenS1{F0: 1, F1: 32767, F2: 327}
    if ((2 <= 2147483648) || false) {
        var v7 uint64 = 214
        _ = v7
    } else {
        v8 := v4[-9223372036854775808]
        var v9 uint8 = 245
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 42}
    v2 = "a"
    v1 = -109
    v11 := GenS3{F0: 249, F1: 42}
    v12 := GenS2{F0: -1}
    if (!true) {
        v13 := v4[9223372036854775807]
        switch int(v1) {
        case 394:
            switch "a" {
            case "a":
                fmt.Println(v13, v2, v1)
            case "":
                fmt.Println(v13, v13, v2)
            }
        case 128:
            v14 := v3[v2]
            _ = v14
        default:
            if (0 >= 2147483647) {
                switch ("" + (v2 + (v2 + v2))) {
                case "foo bar":
                    fmt.Println(v2)
                default:
                    fmt.Println(v13, v13)
                }
                v2 = "\n\t"
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 string = (v2 + v2)
                    var v17 int = (-(fi15))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 10; fi18++ {
                    if v13 {
                        v19 := GenS4{F0: 42}
                        if (905 == 1) {
                            var v20 float32 = (3.14159 * 0.0)
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v13 = ((^(fi18)) >= 4294967295)
                            fmt.Println(v1)
                            v1 = (v1 | 42)
                            _ = len(v5)
                            _ = fi21
                        }
                        _ = v19
                    } else {
                        v1 = 0
                        if (true || ((v2 + "") < v2)) {
                            switch fi18 {
                            case -2147483648:
                                v2 = (("\n\t" + ((("a" + "hello") + v2) + v2)) + (v2 + ((v2 + v2) + (v2 + (v2 + "\n\t")))))
                            case 42:
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    fmt.Println(v1)
                                    v13 = (!v13)
                                    _ = len(v3)
                                    _ = fi22
                                }
                            }
                        }
                    }
                    _ = fi18
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 bool = (17 >= 18446744073709551615)
                    switch v2 {
                    case "a":
                        v25 := GenS3{F0: 0, F1: -386}
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 9; fi26++ {
                            var v27 int32 = (v1 / ((v1) | 1))
                            fmt.Println(fi26, v24, v27)
                            _ = fi26
                            _ = v27
                        }
                    }
                    fmt.Println(v1, v1)
                    _ = fi23
                    _ = v24
                }
            }
        }
        v13 = (float64(-2.25) > (-2.25 * 0.0))
        _ = v13
    } else {
        v2 = ("a" + (v2 + "\n\t"))
        if (0 >= -1) {
            _ = len(v5)
            switch ((v2 + "") + ("a" + v2)) {
            case "\n\t":
                v1 += v1
            case "hello":
                if true {
                    v28 := GenS5{F0: v2, F1: 42}
                    delete(v3, "\n\t")
                    if false {
                        var v29 int32 = (v1 + -1)
                        var v30 bool = false
                        var v31 int8 = -1
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        var v32 string = ("hello" + "")
                        _ = v32
                    }
                    _ = v28
                } else {
                    v1 = 42
                }
            default:
                var v33 bool = (1 >= 255)
                _ = v33
            }
        } else {
            if len(v5) > 0 {
                v5[0] = 32767
            }
            var v34 uint32 = 0
            v5 = append(v5, 4294967295)
            _ = v34
        }
        _ = len(v5)
    }
    switch (v2 + ("\n\t" + v2)) {
    case "\n\t":
        v2 = "foo bar"
    case "foo bar":
        delete(v3, ("" + ((v2 + (((v2 + v2) + (("a" + v2) + v2)) + (v2 + "hello"))) + v2)))
    case "a":
        _ = len(v5)
    default:
        fmt.Println(v2, v2)
    }
    v2 = (v2 + v2)
    v35 := GenS4{F0: 2}
    for fi36 := 0; fi36 < 2; fi36++ {
        if (42 < 0) {
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint64 = 42
                v39 := GenS2{F0: fi36}
                _ = fi37
                _ = v38
                _ = v39
            }
        }
        var v40 int = (^(fi36))
        if (!(false || false)) {
            var v41 uint16 = 255
            _ = v41
        } else {
            var v42 bool = (1 == 83)
            switch v40 {
            case 127:
                fmt.Println(v1, v40)
            default:
                _ = len(v5)
            }
            _ = v42
        }
        _ = fi36
        _ = v40
    }
    var v43 uint = 42
    fmt.Println(v2)
    v1 = v1
    var v44 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
