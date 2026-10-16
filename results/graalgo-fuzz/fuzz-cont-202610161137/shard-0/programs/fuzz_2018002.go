package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }

func main() {
    var v1 uint32 = 0
    var v2 float64 = 3.14159
    v3 := []string{"a"}
    v4 := []string{}
    v5 := GenS1{F0: 0, F1: -143, F2: 556}
    var v6 string = ("" + "foo bar")
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint = 65535
        _ = fi7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint32 = 2147483647
        _ = fi9
        _ = v10
    }
    switch "" {
    case "\n\t":
        v6 = "hello"
    case "a":
        v11 := GenS1{F0: -34, F1: 42, F2: 1000}
        _ = v11
    case "foo bar":
        var v12 uint16 = 664
        _ = v12
    default:
        if false {
            if len(v3) > 0 {
                v3[0] = (v6 + v6)
            }
            v1 = v1
            v13 := GenS2{F0: 249, F1: 0, F2: -285}
            _ = v13
        }
    }
    if len(v4) > 0 {
        v4[0] = ("" + v6)
    }
    v1 += v1
    if ((2 <= 2) || true) {
        v14 := GenS2{F0: 42, F1: 161, F2: 1}
        _ = v14
    }
    if ((true || false) && true) {
        v1 = v1
        switch 950 {
        case 855:
            if ((!(4294967295 > 4294967295)) && true) {
                if (v2 <= 0.0) {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        var v16 float64 = (v2 + v2)
                        var v17 uint32 = 885
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                var v18 int64 = 1
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch v6 {
                    case "foo bar":
                        if false {
                            for fi20 := 0; fi20 < 8; fi20++ {
                                fmt.Println(v1)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v6, fi20)
                                _ = fi20
                            }
                            switch (v6 + "") {
                            case "\n\t":
                                for fi21 := 0; fi21 < 19; fi21++ {
                                    fmt.Println(fi19, fi21)
                                    fmt.Println(v18)
                                    if (!(!true)) {
                                        if (v2 != v2) {
                                            if true {
                                                switch (^(fi21)) {
                                                case 42:
                                                    var v22 uint8 = 234
                                                    _ = v22
                                                case 436:
                                                    v6 = v6
                                                }
                                            } else {
                                                v18 -= 9223372036854775807
                                                v23 := GenS2{F0: 42, F1: v18, F2: (^(fi19))}
                                                v1 = 525
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v23
                                            }
                                        }
                                        v1 = v1
                                    }
                                    _ = fi21
                                }
                            case "a":
                                var v24 float32 = 0.0
                                _ = v24
                            }
                        } else {
                            if true {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    v4 = append(v4, "\n\t")
                                    _ = fi25
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                for fi26 := 0; fi26 < 6; fi26++ {
                                    v27 := GenS2{F0: 106, F1: 210, F2: 32768}
                                    _ = fi26
                                    _ = v27
                                }
                            }
                        }
                    case "":
                        v2 += ((v2 + 0.0) - 0.1)
                    case "hello":
                        var v28 int = fi19
                        _ = v28
                    default:
                        switch (^(fi19)) {
                        case 1:
                            var v29 int16 = 0
                            _ = v29
                        }
                    }
                    v2 -= v2
                    if true {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            if ((true || (((v1 ^ 1) != v1) && (156 == 2))) || (18446744073709551615 > 526)) {
                                if true {
                                    v31 := GenS3{F0: 1, F1: -57, F2: 9223372036854775807}
                                    _ = v31
                                }
                                var v32 int = fi19
                                v33 := GenS3{F0: 42, F1: 0, F2: 42}
                                _ = v32
                                _ = v33
                            }
                            v1 -= v1
                            _ = fi30
                        }
                        v6 = "foo bar"
                        var v34 int64 = -17
                        _ = v34
                    }
                    _ = fi19
                }
                _ = v18
            }
        case 981:
            v6 = "hello"
        case -158:
            if true {
                v35 := GenS1{F0: 127, F1: 127, F2: -706}
                _ = v35
            } else {
                var v36 bool = false
                fmt.Println(v1, v2, v6)
                _ = v36
            }
        }
    }
    v1 = 1
    fmt.Println(v1, v2)
    if false {
        if false {
            var v37 int16 = 0
            var v38 string = v6
            var v39 bool = ((v1 ^ v1) < v1)
            _ = v37
            _ = v38
            _ = v39
        }
        v1 = v1
    }
    switch 9223372036854775807 {
    case 1:
        v40 := GenS3{F0: 160, F1: -380, F2: 0}
        _ = v40
    case 32767:
        var v41 float32 = (0.1 / 1.0)
        _ = v41
    }
    var v42 int64 = 1
    switch 2147483647 {
    case 635:
        v2 += v2
    case 4294967295:
        fmt.Println(v42)
    case 42:
        var v43 uint32 = v1
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
