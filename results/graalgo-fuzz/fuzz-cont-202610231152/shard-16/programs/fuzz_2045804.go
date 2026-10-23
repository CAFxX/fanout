package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    var v3 string = v2
    var v4 int32 = 0
    v5 := []string{"\n\t"}
    v6 := []string{v2, ("a" + ((("foo bar" + "foo bar") + (v2 + v3)) + (("hello" + "") + "\n\t"))), (v3 + v2)}
    _ = len(v6)
    v3 = "a"
    if true {
        _ = len(v1)
        if false {
            var v7 int64 = 42
            for fi8 := 0; fi8 < 18; fi8++ {
                v2 = "\n\t"
                var v9 int64 = v7
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            v2 = v2
            switch "\n\t" {
            case "\n\t":
                v10 := GenS1{F0: 2, F1: 42, F2: 530}
                _ = v10
            }
        }
        if (100 >= 58) {
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v2 = ""
                    if (65535 > 0) {
                        switch v2 {
                        case "":
                            v4 = -904
                        case "a":
                            if (0 <= 0) {
                                switch fi12 {
                                case -677:
                                    v13 := v1[fi12]
                                    _ = v13
                                case 9223372036854775807:
                                    v1[(fi12 * fi11)] = (fi11 & (fi12 << ((1) & 63)))
                                case 2147483647:
                                    switch (fi12 + fi11) {
                                    case 128:
                                        var v14 uint = 1
                                        _ = v14
                                    case 819:
                                        v15 := GenS1{F0: 65535, F1: (fi12 ^ (fi11 | (fi12 & fi11))), F2: 2}
                                        _ = v15
                                    case 1:
                                        switch ("foo bar" + v2) {
                                        case "\n\t":
                                            if false {
                                                v16 := GenS1{F0: 2147483647, F1: (^(fi12)), F2: 2147483647}
                                                var v17 uint8 = 0
                                                _ = v16
                                                _ = v17
                                            } else {
                                                if (4294967295 <= 18446744073709551615) {
                                                    v2 = ""
                                                } else {
                                                    v3 = (v2 + "\n\t")
                                                    v4 += v4
                                                }
                                                v18 := GenS2{F0: 42, F1: 1}
                                                _ = v18
                                            }
                                        case "":
                                            var v19 uint16 = 255
                                            _ = v19
                                        case "hello":
                                            v20 := GenS3{F0: 361, F1: 1000}
                                            _ = v20
                                        default:
                                            v2 = ("hello" + "foo bar")
                                        }
                                    default:
                                        switch ((("hello" + v3) + "hello") + v2) {
                                        case "foo bar":
                                            var v21 int = (fi12 + (fi11 - fi11))
                                            _ = v21
                                        case "":
                                            var v22 float64 = 1.5
                                            _ = v22
                                        case "\n\t":
                                            fmt.Println(fi11)
                                        }
                                    }
                                }
                            } else {
                                var v23 uint16 = 440
                                v24 := GenS3{F0: 0, F1: 1}
                                var v25 int64 = -89
                                if (148 >= v25) {
                                    fmt.Println(v25, v3)
                                }
                                v26 := GenS2{F0: 127, F1: 1}
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                    }
                    _ = fi12
                }
                v1[fi11] = (^(fi11))
                _ = fi11
            }
        }
    } else {
        v3 = "\n\t"
        switch 298 {
        case 0:
            if true {
                for fi27 := 0; fi27 < 5; fi27++ {
                    var v28 uint32 = 2
                    _ = fi27
                    _ = v28
                }
                v29 := GenS1{F0: 4294967295, F1: -255, F2: 234}
                _ = v29
            } else {
                v4 += -729
                v2 = v2
            }
        case -93:
            if (42 > 0) {
                var v30 uint16 = 0
                var v31 int8 = 0
                v32 := GenS2{F0: 127, F1: 32767}
                _ = v30
                _ = v31
                _ = v32
            } else {
                v33 := GenS2{F0: 1, F1: 65535}
                if true {
                    var v34 int16 = -32768
                    _ = v34
                } else {
                    var v35 bool = (true && (!(!(899 >= 562))))
                    _ = v35
                }
                fmt.Println(v4, v2)
                _ = v33
            }
        case 128:
            var v36 uint8 = 244
            _ = v36
        }
    }
    v2 = ("\n\t" + v2)
    if (!((!true) || ((!false) && true))) {
        v5 = append(v5, "hello")
        v2 = ""
    }
    fmt.Println(v2)
    fmt.Println(v4, v3, v2)
    for fi37 := 0; fi37 < 16; fi37++ {
        var v38 float32 = float32(0.1)
        var v39 float32 = v38
        _ = fi37
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        v2 = "hello"
        v41 := GenS4{F0: 1, F1: -1}
        _ = fi40
        _ = v41
    }
    if false {
        v1[32767] = -479
        var v42 int64 = 0
        _ = v42
    }
    switch 2147483647 {
    case -2147483648:
        v43 := GenS2{F0: 42, F1: 944}
        _ = v43
    default:
        switch "a" {
        case "":
            v2 = (v3 + "")
        case "a":
            var v44 uint = 57
            _ = v44
        case "\n\t":
            var v45 int16 = -1
            _ = v45
        default:
            if true {
                var v46 string = v3
                if (true || (-128 > 100)) {
                    for fi47 := 0; fi47 < 7; fi47++ {
                        v48 := GenS2{F0: 0, F1: 255}
                        if (false || (v46 > (v46 + "a"))) {
                            var v49 uint64 = 4294967295
                            fmt.Println(v3)
                            _ = v49
                        }
                        _ = fi47
                        _ = v48
                    }
                } else {
                    v50 := GenS2{F0: 2, F1: 2}
                    _ = v50
                }
                _ = v46
            } else {
                var v51 uint8 = 14
                v1[739] = 4294967295
                var v52 float64 = 0.0
                _ = v51
                _ = v52
            }
        }
    }
    v53 := GenS5{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
}
