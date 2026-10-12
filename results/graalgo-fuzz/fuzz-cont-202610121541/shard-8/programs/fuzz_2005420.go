package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint64 }

func main() {
    var v1 int = 128
    var v2 uint64 = 634
    var v3 string = (((("a" + (("" + "hello") + "hello")) + "a") + "\n\t") + "foo bar")
    v4 := []uint32{2147483647, 452}
    var v5 int32 = 1
    var v6 string = v3
    switch int(v5) {
    case 9223372036854775807:
        if (3.14159 == 3.14159) {
            switch "a" {
            case "\n\t":
                v7 := GenS1{F0: v1, F1: 2147483648, F2: 1}
                _ = v7
            case "a":
                var v8 bool = true
                _ = v8
            }
            v6 = ("a" + (((v6 + "") + v3) + "a"))
            for fi9 := 0; fi9 < 4; fi9++ {
                if true {
                    var v10 int8 = -33
                    _ = v10
                }
                _ = fi9
            }
        }
    case 275:
        var v11 int8 = -128
        _ = v11
    case 128:
        var v12 bool = false
        _ = v12
    default:
        if len(v4) > 0 {
            v4[0] = 2
        }
    }
    if true {
        for fi13 := 0; fi13 < 17; fi13++ {
            v4 = append(v4, 42)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS2{F0: false, F1: 42, F2: v3}
        switch 533 {
        case 2147483647:
            var v16 uint32 = 887
            _ = v16
        default:
            fmt.Println(v1, v6, v1)
        }
        _ = fi14
        _ = v15
    }
    v17 := GenS2{F0: false, F1: 255, F2: (("foo bar" + (v6 + "\n\t")) + "foo bar")}
    for fi18 := 0; fi18 < 18; fi18++ {
        fmt.Println(fi18, v2, v2)
        switch (fi18 & 42) {
        case -630:
            _ = len(v4)
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v5, v3, v1)
            _ = fi19
        }
        _ = fi18
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        switch v1 {
        case 32767:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -2147483648:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case -1:
            if true {
                var v20 float32 = (0.0 + 1e10)
                v3 = ("" + v6)
                if (!(!(!(100 != 275)))) {
                    v1 = (^(v1))
                } else {
                    for fi21 := 0; fi21 < 14; fi21++ {
                        switch v6 {
                        case "":
                            for fi22 := 0; fi22 < 15; fi22++ {
                                if (((65535 > 2) || false) && (!(!false))) {
                                    fmt.Println(v6, v6)
                                    var v23 uint = 65535
                                    fmt.Println(v6, v1, v5)
                                    _ = v23
                                } else {
                                    switch (v6 + ("a" + (v6 + "\n\t"))) {
                                    case "":
                                        v3 = "hello"
                                    case "foo bar":
                                        switch (v6 + v6) {
                                        case "foo bar":
                                            v24 := GenS2{F0: (0.0 > ((float32(1.5) * v20) + 1e10)), F1: 127, F2: "a"}
                                            _ = v24
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 0
                                        }
                                    case "a":
                                        var v25 string = v6
                                        _ = v25
                                    }
                                    switch v6 {
                                    case "\n\t":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "":
                                        v26 := GenS3{F0: "hello", F1: 214}
                                        _ = v26
                                    }
                                    var v27 int64 = -9223372036854775808
                                    _ = v27
                                }
                                _ = fi22
                            }
                        case "foo bar":
                            fmt.Println(v3, v1, v2)
                        }
                        v1 = (v1 - fi21)
                        var v28 int = v1
                        _ = fi21
                        _ = v28
                    }
                }
                _ = v20
            } else {
                fmt.Println(v5)
                v29 := GenS4{F0: 431}
                var v30 uint = 2
                _ = v29
                _ = v30
            }
        default:
            v6 = v6
        }
        var v31 uint32 = 42
        _ = v31
    }
    v2 = 2
    v32 := GenS5{F0: "", F1: (v2 - (v2 + v2))}
    var v33 uint8 = 127
    v34 := GenS3{F0: "foo bar", F1: -9223372036854775808}
    var v35 int8 = -120
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint16 = 58
        _ = fi36
        _ = v37
    }
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
