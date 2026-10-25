package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = (v1 & v1)
    var v3 uint = 4294967295
    var v4 int8 = 0
    v5 := []uint32{}
    for fi6 := 0; fi6 < 3; fi6++ {
        switch fi6 {
        case 32768:
            v7 := GenS1{F0: (v4 >> ((42) & 63)), F1: (v4 | 42), F2: (v3 / ((v3) | 1))}
            _ = v7
        }
        v8 := GenS2{F0: -291}
        _ = fi6
        _ = v8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch ("foo bar" + ("" + "a")) {
    case "foo bar":
        if (false && (!(!(true && (v3 < v3))))) {
            if (((((3.14159 * float32((1.5 + -2.25))) * 1.5) / 1.0) + 1.5) <= 3.14159) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int32 = 42
                    _ = fi9
                    _ = v10
                }
                v3 = (v3 * v3)
            }
        } else {
            if (v2 > v1) {
                switch ("" + "foo bar") {
                case "hello":
                    var v11 uint64 = 548
                    _ = v11
                case "":
                    v12 := GenS2{F0: 0}
                    _ = v12
                case "\n\t":
                    if true {
                        switch ("\n\t" + ("\n\t" + "a")) {
                        case "":
                            var v13 string = "foo bar"
                            _ = v13
                        case "hello":
                            if true {
                                v14 := GenS2{F0: 0}
                                _ = v14
                            } else {
                                var v15 string = ("hello" + "foo bar")
                                _ = v15
                            }
                        }
                        for fi16 := 0; fi16 < 3; fi16++ {
                            _ = len(v5)
                            if true {
                                if (321 > -551) {
                                    if true {
                                        v17 := GenS2{F0: 726}
                                        _ = v17
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483648
                                        }
                                    }
                                }
                                var v18 float32 = 1.5
                                _ = v18
                            }
                            fmt.Println(fi16)
                            _ = fi16
                        }
                    } else {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v3 -= (v3 << ((2) & 63))
                            var v20 int = (fi19 - fi19)
                            _ = fi19
                            _ = v20
                        }
                        v5 = append(v5, 42)
                        var v21 int = -142
                        _ = v21
                    }
                }
                var v22 string = "foo bar"
                _ = v22
            } else {
                switch -56 {
                case 4294967295:
                    v3 -= 2
                case -2147483648:
                    var v23 float32 = (-2.25 + (float32(((3.14159 * -2.25) + 1e10)) + 0.0))
                    _ = v23
                default:
                    if true {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS2{F0: -1}
                            _ = fi24
                            _ = v25
                        }
                        var v26 uint16 = (v2 & v1)
                        v2 = 2
                        _ = v26
                    } else {
                        _ = len(v5)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v27 := GenS2{F0: -9223372036854775808}
                        _ = v27
                    }
                }
            }
            if false {
                v28 := GenS1{F0: (v4 % (((v4 >> ((v3) & 63))) | 1)), F1: (v4 - (v4 - v4)), F2: (v3 | v3)}
                _ = len(v5)
                _ = v28
            }
        }
    case "a":
        fmt.Println(v1, v2)
    case "hello":
        fmt.Println(v3, v1)
    }
    var v29 uint = (v3 << ((v3) & 63))
    switch -634 {
    case 1:
        if (false || (!(int8(v2) > (v4 & 100)))) {
            var v30 int = 153
            v31 := GenS3{F0: v4, F1: 128, F2: 42}
            var v32 int8 = v4
            _ = v30
            _ = v31
            _ = v32
        } else {
            v29 = v29
            v4 -= v4
            fmt.Println(v29, v2, v4)
        }
    case 42:
        var v33 int8 = 0
        _ = v33
    case -2147483648:
        fmt.Println(v2, v4, v2)
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        for fi35 := 0; fi35 < 4; fi35++ {
            v36 := GenS1{F0: v4, F1: (-(v4)), F2: (^(v3))}
            _ = fi35
            _ = v36
        }
        var v37 int = (fi34 + (fi34 << ((v29) & 63)))
        _ = fi34
        _ = v37
    }
    switch ((("" + "a") + "") + "hello") {
    case "\n\t":
        var v38 uint16 = v1
        _ = v38
    }
    v39 := GenS2{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
