package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = (("\n\t" + "a") + "a")
    v3 := make(map[uint32]string)
    var v4 uint16 = 1
    v5 := make(map[int]bool)
    var v6 int8 = 1
    var v7 string = (v2 + (v2 + (v2 + v2)))
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: v6, F1: 0}
        switch fi8 {
        case 128:
            if true {
                for fi10 := 0; fi10 < 16; fi10++ {
                    v11 := GenS2{F0: 281, F1: 980}
                    v12 := GenS3{F0: v6, F1: 0}
                    v13 := []int{fi10, fi8}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            }
        }
        v7 = v7
        _ = fi8
        _ = v9
    }
    v14 := []int{42}
    v6 = v6
    v2 = (v2 + (((("hello" + "\n\t") + (("\n\t" + "foo bar") + "\n\t")) + ("\n\t" + v7)) + v2))
    fmt.Println(v7)
    if false {
        var v15 bool = false
        switch 1 {
        case 128:
            if (v2 != "a") {
                fmt.Println(v2)
            } else {
                var v16 int16 = 42
                _ = v16
            }
        case 2147483647:
            var v17 uint16 = (^(v4))
            _ = v17
        case 726:
            var v18 int64 = -30
            _ = v18
        default:
            v4 = (v4 & (v4 + 0))
        }
        _ = v15
    }
    v6 -= -50
    v2 = v2
    fmt.Println(v4)
    fmt.Println(v6)
    v2 = v2
    var v19 uint64 = 1
    if true {
        v20 := GenS2{F0: 196, F1: 854}
        var v21 uint8 = 2
        var v22 uint64 = 378
        _ = v20
        _ = v21
        _ = v22
    } else {
        var v23 int64 = -9223372036854775808
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v19 = (v19 * (v19 >> ((831) & 63)))
            var v26 int = fi24
            fmt.Println(v26)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 13; fi28++ {
            if true {
                fmt.Println(fi27, v4, fi28)
                _ = len(v14)
                var v29 int8 = int8(v4)
                _ = v29
            } else {
                v30 := GenS1{F0: 100, F1: 255}
                v4 = (^(v4))
                _ = v30
            }
            switch "a" {
            case "\n\t":
                v4 -= v4
            case "hello":
                if (((v2 != (v2 + v7)) && (127 >= 1)) && true) {
                    if false {
                        v14 = append(v14, (-(fi27)))
                        v4 = v4
                        var v31 int8 = v6
                        _ = v31
                    }
                } else {
                    if (18446744073709551615 != 0) {
                        v4 -= v4
                        fmt.Println(fi27)
                        if true {
                            var v32 uint = 0
                            if (!false) {
                                if (!(0 != 127)) {
                                    v7 = "hello"
                                } else {
                                    var v33 int = (fi28 + fi27)
                                    _ = v33
                                }
                                v7 = (v2 + v7)
                                if true {
                                    v2 = ("foo bar" + "foo bar")
                                    fmt.Println(v2)
                                    switch v7 {
                                    case "a":
                                        v34 := GenS3{F0: v6, F1: 0}
                                        _ = v34
                                    default:
                                        v3[42] = ("a" + v7)
                                    }
                                    delete(v3, uint32(v19))
                                    v35 := GenS3{F0: (v6 >> ((v32) & 63)), F1: 2}
                                    _ = v35
                                } else {
                                    var v36 float64 = (0.1 + -2.25)
                                    var v37 float32 = 0.0
                                    _ = v36
                                    _ = v37
                                }
                            } else {
                                v14 = append(v14, (fi28 ^ 4294967295))
                            }
                            _ = v32
                        } else {
                            v2 = "foo bar"
                            var v38 uint = 1
                            var v39 uint32 = 1
                            _ = v38
                            _ = v39
                        }
                    } else {
                        v7 = v7
                        v40 := GenS3{F0: v6, F1: 2}
                        _ = v40
                    }
                }
            case "":
                v6 = (-(v6))
                fmt.Println(fi28)
            }
            _ = fi28
        }
        _ = fi27
    }
    if (-252 > 1) {
        if true {
            switch (((((v7 + v7) + v2) + "") + v2) + v7) {
            case "hello":
                if false {
                    if (true && false) {
                        var v41 string = v7
                        _ = v41
                    } else {
                        v42 := GenS1{F0: (v6 << ((42) & 63)), F1: 1}
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 5; fi43++ {
                        v14 = append(v14, fi43)
                        v6 += int8(fi43)
                        fmt.Println(v2, v4)
                        _ = fi43
                    }
                    v19 -= v19
                }
            case "foo bar":
                for fi44 := 0; fi44 < 20; fi44++ {
                    v7 = (v7 + v2)
                    v45 := GenS4{F0: (v7 + v2), F1: 1000}
                    _ = fi44
                    _ = v45
                }
            }
            fmt.Println(v19)
        }
    } else {
        if false {
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 int16 = 42
                _ = fi46
                _ = v47
            }
        }
        for fi48 := 0; fi48 < 16; fi48++ {
            switch v7 {
            case "foo bar":
                if len(v14) > 0 {
                    v14[0] = fi48
                }
            case "hello":
                if len(v14) > 0 {
                    v14[0] = fi48
                }
            case "":
                v49 := GenS5{F0: -32768, F1: 2147483647}
                _ = v49
            }
            for fi50 := 0; fi50 < 2; fi50++ {
                v51 := v3[338]
                _ = fi50
                _ = v51
            }
            _ = fi48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
