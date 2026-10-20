package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS6 struct { F0 uint16; F1 bool; F2 int16 }
type GenS7 struct { F0 int8; F1 bool }

func main() {
    v1 := []uint32{811, 282}
    var v2 int8 = 116
    var v3 int = 0
    var v4 uint = 510
    var v5 int8 = v2
    for fi6 := 0; fi6 < 19; fi6++ {
        v1 = append(v1, 42)
        v7 := GenS1{F0: -32768}
        _ = fi6
        _ = v7
    }
    if (!((0.0 - ((0.0 + 3.14159) * 1.5)) == float64(3.14159))) {
        var v8 uint64 = 42
        _ = v8
    }
    v5 += v2
    v3 -= -9223372036854775808
    for fi9 := 0; fi9 < 4; fi9++ {
        switch ("foo bar" + ((("\n\t" + (("\n\t" + ("hello" + "\n\t")) + (("hello" + "foo bar") + "foo bar"))) + "foo bar") + "hello")) {
        case "hello":
            if (!((fi9 / (((fi9 / ((v3) | 1))) | 1)) != (fi9 & fi9))) {
                for fi10 := 0; fi10 < 5; fi10++ {
                    if false {
                        _ = len(v1)
                    } else {
                        v4 = (v4 ^ v4)
                        var v11 string = "hello"
                        var v12 uint64 = 2
                        v4 += v4
                        v13 := GenS2{F0: -637, F1: 4294967295}
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 101}
                    if false {
                        var v15 uint16 = 65535
                        if ((^(v15)) == (v15 - (v15 ^ v15))) {
                            if (18446744073709551615 <= 1) {
                                v2 = -1
                                var v16 int16 = -86
                                _ = v16
                            } else {
                                v17 := GenS2{F0: 1, F1: 686}
                                var v18 int = 1
                                v5 -= (v2 + v2)
                                _ = v17
                                _ = v18
                            }
                            fmt.Println(v15)
                            var v19 uint32 = 4294967295
                            _ = v19
                        } else {
                            v20 := GenS1{F0: 1}
                            _ = v20
                        }
                        var v21 bool = ((v5 | v2) == v5)
                        _ = v15
                        _ = v21
                    } else {
                        v3 -= -9223372036854775808
                    }
                    _ = fi10
                    _ = v14
                }
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS1{F0: 0}
                    v24 := GenS3{F0: 2}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if true {
                    for fi25 := 0; fi25 < 14; fi25++ {
                        if ((0 < -373) && (!(false && true))) {
                            var v26 int64 = 9223372036854775807
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            v5 = (v5 + (v5 + v5))
                            _ = v26
                        } else {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v2 -= (v2 ^ 39)
                                if false {
                                    v27 := GenS4{F0: 100}
                                    _ = v27
                                }
                            }
                            if true {
                                v5 = (v5 - -128)
                                v3 = fi9
                            } else {
                                var v28 uint64 = 42
                                _ = v28
                            }
                            var v29 uint64 = 18446744073709551615
                            _ = v29
                        }
                        v5 = 42
                        var v30 int64 = 9223372036854775807
                        _ = fi25
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v32 := GenS5{F0: 127, F1: 206, F2: 42}
                        v33 := GenS4{F0: (v2 * (v2 ^ (v5 - v2)))}
                        var v34 bool = false
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                }
            }
        case "":
            v35 := GenS5{F0: 29, F1: 42, F2: 2}
            _ = v35
        }
        if (3.14159 < (0.1 * ((0.0 - 0.1) * ((0.1 * 0.1) / 1e10)))) {
            v1 = append(v1, 1)
            for fi36 := 0; fi36 < 17; fi36++ {
                if (!false) {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    var v37 int64 = 0
                    v5 += (^(v2))
                    v5 = (v2 + -128)
                    _ = v37
                }
                _ = len(v1)
                v4 = 4294967295
                var v38 int = 32768
                fmt.Println(v4, v38)
                _ = fi36
                _ = v38
            }
            _ = len(v1)
        }
        var v39 bool = true
        _ = fi9
        _ = v39
    }
    switch v3 {
    case -105:
        v40 := GenS1{F0: 32767}
        _ = v40
    case -579:
        for fi41 := 0; fi41 < 4; fi41++ {
            v2 += v2
            switch fi41 {
            case 804:
                var v42 int = -951
                _ = v42
            case 128:
                var v43 string = "foo bar"
                _ = v43
            }
            _ = fi41
        }
    case 127:
        var v44 int8 = (-(v2))
        _ = v44
    }
    switch (^(v3)) {
    case 498:
        v45 := GenS6{F0: 187, F1: true, F2: 1}
        _ = v45
    case -1:
        for fi46 := 0; fi46 < 8; fi46++ {
            v47 := GenS1{F0: 42}
            var v48 string = "hello"
            _ = fi46
            _ = v47
            _ = v48
        }
    }
    var v49 uint = v4
    var v50 bool = ((42 != 937) && false)
    if true {
        v51 := GenS7{F0: 42, F1: (v3 < v3)}
        v49 -= v49
        _ = v51
    }
    v2 -= v2
    v50 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
