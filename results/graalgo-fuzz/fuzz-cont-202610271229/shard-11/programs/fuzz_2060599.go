package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 int16; F2 uint32 }
func genFn0(p0 uint8, p1 bool, p2 uint) uint32 {
    p1 = p1
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 = (v4 & v4)
        if false {
            if (!(32767 <= 6)) {
                if false {
                    var v6 uint = p2
                    _ = v6
                }
            }
            fmt.Println(v4, fi5, p2)
        }
        var v7 float32 = 1e10
        _ = fi5
        _ = v7
    }
    if (-128 != 0) {
        v8 := []int{32767}
        p1 = (-55 == 0)
        _ = v8
    }
    var v9 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return 65535
}
func genFn1(p0 uint16, p1 int16, p2 uint32) string {
    v10 := []int{-85, 641}
    switch "\n\t" {
    case "":
        v10 = append(v10, 922)
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        if (false || true) {
            _ = len(v10)
            p0 = (p0 & p0)
            switch "a" {
            case "a":
                _ = len(v10)
            case "hello":
                var v12 int16 = 32767
                _ = v12
            case "\n\t":
                if true {
                    if (0 != 2147483647) {
                        switch "\n\t" {
                        case "":
                            p0 -= p0
                        case "\n\t":
                            for fi13 := 0; fi13 < 9; fi13++ {
                                var v14 uint32 = (p2 + (p2 ^ p2))
                                var v15 uint16 = p0
                                p1 += p1
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                }
            }
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                switch 128 {
                case -9223372036854775808:
                    var v17 bool = true
                    _ = v17
                case 509:
                    var v18 float32 = (1e10 / 0.1)
                    _ = v18
                }
                v19 := GenS1{F0: 669, F1: 567, F2: (p2 % ((genFn0(128, false, 0)) | 1))}
                fmt.Println(p2)
                _ = fi16
                _ = v19
            }
        }
        v20 := GenS2{F0: 512, F1: ""}
        var v21 int8 = 42
        _ = fi11
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 1, F1: ""}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v22
    return (("hello" + "") + "\n\t")
}

func main() {
    var v1 int8 = -128
    var v2 float32 = 0.0
    v3 := []string{("foo bar" + "hello")}
    var v23 uint64 = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 uint = 4294967295
    var v25 uint = v24
    v26 := GenS3{F0: 0, F1: 108, F2: 255}
    var v27 uint = (v24 ^ 4294967295)
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 string = (("\n\t" + "") + (("a" + "foo bar") + ("\n\t" + "a")))
        if (true || true) {
            v25 = 1
            if len(v3) > 0 {
                v3[0] = (v29 + "hello")
            }
        } else {
            v25 = 0
            v3 = append(v3, v29)
        }
        _ = fi28
        _ = v29
    }
    switch "a" {
    case "":
        v30 := GenS4{F0: -1}
        _ = v30
    }
    var v31 int32 = -267
    fmt.Println(v25, v24, v31)
    var v32 uint16 = 255
    v2 -= v2
    v1 = 26
    if ((false || true) || false) {
        for fi33 := 0; fi33 < 15; fi33++ {
            if (!false) {
                var v34 uint = uint(v32)
                var v35 float32 = 1.5
                _ = v34
                _ = v35
            } else {
                for fi36 := 0; fi36 < 4; fi36++ {
                    var v37 int = fi36
                    fmt.Println(v25)
                    if (!false) {
                        if (!(true || (!(!true)))) {
                            if ((-9223372036854775808 == 359) && ((("" + "hello") + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) > "a")) {
                                var v38 int32 = (v31 - -183)
                                _ = v38
                            }
                            var v39 int16 = 1
                            v40 := GenS5{F0: (v31 * (v31 | (v31 % (((v31 ^ (v31 + v31))) | 1)))), F1: (v39 + v39), F2: uint32(v27)}
                            var v41 int64 = 0
                            if false {
                                if false {
                                    v42 := GenS2{F0: v41, F1: "hello"}
                                    _ = v42
                                }
                                var v43 uint8 = 42
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        v44 := GenS3{F0: v41, F1: (^(v32)), F2: v43}
                                        _ = v44
                                    case "a":
                                        var v45 uint64 = v23
                                        _ = v45
                                    case "foo bar":
                                        if (v1 >= (v1 & (v1 & v1))) {
                                            if len(v3) > 0 {
                                                v3[0] = ("a" + "hello")
                                            }
                                        } else {
                                            v41 = v41
                                        }
                                    }
                                }
                                _ = v43
                            } else {
                                fmt.Println(v23, v27, v2)
                                fmt.Println(v32, v41)
                                v32 += v32
                            }
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ("\n\t" + (("hello" + "hello") + ("" + ("hello" + "a"))))
                        }
                        fmt.Println(v2, fi36)
                        var v46 int16 = int16(v27)
                        _ = v46
                    }
                    _ = fi36
                    _ = v37
                }
                if true {
                    var v47 int = int(v31)
                    for fi48 := 0; fi48 < 18; fi48++ {
                        fmt.Println(v31)
                        _ = fi48
                    }
                    _ = v47
                } else {
                    fmt.Println(v1, v23)
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v50 := GenS4{F0: -2147483648}
                        _ = fi49
                        _ = v50
                    }
                }
            }
            fmt.Println(v23)
            if len(v3) > 0 {
                v3[0] = ("a" + (("hello" + "hello") + "hello"))
            }
            _ = fi33
        }
    } else {
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    var v51 uint64 = (v23 | 457)
    var v52 string = ""
    for fi53 := 0; fi53 < 3; fi53++ {
        switch (^(fi53)) {
        case -9223372036854775808:
            for fi54 := 0; fi54 < 17; fi54++ {
                var v55 float64 = 0.1
                _ = fi54
                _ = v55
            }
        case 2147483647:
            var v56 int64 = -314
            _ = v56
        }
        v51 -= v23
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v51
    _ = v52
}
