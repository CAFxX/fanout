package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint16; F1 int }
func genFn0(p0 uint16) uint32 {
    switch ("a" + "a") {
    case "hello":
        v4 := []int{-1}
        _ = v4
    case "":
        var v5 int32 = 497
        _ = v5
    default:
        fmt.Println(p0)
    }
    switch "" {
    case "\n\t":
        p0 = p0
    case "a":
        switch -436 {
        case 4294967295:
            v6 := []int{-1}
            _ = v6
        case 127:
            var v7 int = 0
            _ = v7
        }
    }
    _ = p0
    return 2
}

func main() {
    var v1 float32 = (1.5 - 0.0)
    var v2 float64 = 0.1
    var v3 uint16 = 2
    switch "\n\t" {
    case "foo bar":
        var v8 int8 = 0
        _ = v8
    case "\n\t":
        fmt.Println(v2, v1)
    case "a":
        fmt.Println(v1)
    default:
        v9 := GenS1{F0: 0}
        _ = v9
    }
    var v10 uint = 166
    var v11 string = ""
    var v12 uint16 = v3
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                if true {
                    var v15 int32 = 797
                    fmt.Println(v2, v12, fi13)
                    var v16 uint = v10
                    _ = v15
                    _ = v16
                } else {
                    if false {
                        if false {
                            v17 := []int{(fi14 + fi13)}
                            _ = v17
                        }
                    } else {
                        fmt.Println(fi13, v1)
                        if (true || false) {
                            v12 -= v12
                        } else {
                            switch (fi14 << ((v10) & 63)) {
                            case 9223372036854775807:
                                var v18 int8 = 0
                                _ = v18
                            default:
                                var v19 int64 = -501
                                _ = v19
                            }
                            v1 = v1
                        }
                        fmt.Println(v12, fi14)
                        v11 = "\n\t"
                    }
                    var v20 int64 = 1
                    v21 := []int{fi13, fi13, (fi13 - fi14), (fi14 << ((v10) & 63))}
                    _ = v20
                    _ = v21
                }
            } else {
                v12 = (v12 & v12)
            }
            fmt.Println(fi14, v3, v3)
            _ = fi14
        }
        var v22 float64 = -2.25
        _ = fi13
        _ = v22
    }
    v23 := GenS2{F0: (("hello" + "\n\t") + ""), F1: (0 == 0), F2: 128}
    v12 -= (v3 + v3)
    switch (("hello" + (v11 + (v11 + (v11 + ("\n\t" + v11))))) + "foo bar") {
    case "hello":
        v24 := []int{126, 32768, -725, 0}
        _ = v24
    case "a":
        if true {
            var v25 bool = false
            v26 := GenS1{F0: -1}
            _ = v25
            _ = v26
        } else {
            fmt.Println(v3, v12, v12)
        }
    default:
        if false {
            switch 128 {
            case -1:
                for fi27 := 0; fi27 < 11; fi27++ {
                    v3 -= v12
                    v2 = float64(0.1)
                    v1 = v1
                    _ = fi27
                }
            case 771:
                v28 := GenS2{F0: "foo bar", F1: false, F2: 2}
                _ = v28
            case 32767:
                for fi29 := 0; fi29 < 17; fi29++ {
                    if (2 >= 891) {
                        fmt.Println(v11, v3)
                        if true {
                            v2 = v2
                            switch (fi29 * -250) {
                            case 0:
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    var v31 float64 = v2
                                    _ = fi30
                                    _ = v31
                                }
                            case -9223372036854775808:
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    v2 -= v2
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v33 := GenS3{F0: 808}
                                            _ = v33
                                        }
                                    } else {
                                        fmt.Println(v2)
                                        var v34 uint = (^(v10))
                                        v3 += (v12 * 1)
                                        _ = v34
                                    }
                                    _ = fi32
                                }
                            case 127:
                                v10 += v10
                            }
                        }
                        v1 += 0.0
                    }
                    v1 -= 0.1
                    _ = fi29
                }
            }
            switch v11 {
            case "\n\t":
                fmt.Println(v1, v2)
            case "foo bar":
                var v35 string = (v11 + ("hello" + (v11 + "")))
                _ = v35
            case "a":
                switch (v11 + "\n\t") {
                case "hello":
                    if false {
                        v36 := GenS3{F0: v12}
                        v37 := GenS2{F0: "foo bar", F1: (false || (32767 <= -32768)), F2: 2}
                        _ = v36
                        _ = v37
                    } else {
                        v2 -= 0.1
                        if (42 != 252) {
                            v38 := GenS4{F0: 42}
                            _ = v38
                        } else {
                            v39 := GenS5{F0: 208, F1: 18446744073709551615}
                            _ = v39
                        }
                    }
                default:
                    v40 := GenS6{F0: (^(v3)), F1: -9223372036854775808}
                    _ = v40
                }
            default:
                if (0 != -493) {
                    var v41 float32 = v1
                    v3 = uint16(v10)
                    fmt.Println(v12, v2)
                    _ = v41
                } else {
                    v42 := GenS6{F0: (v12 - v3), F1: 4294967295}
                    _ = v42
                }
            }
        }
    }
    fmt.Println(v11, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
}
