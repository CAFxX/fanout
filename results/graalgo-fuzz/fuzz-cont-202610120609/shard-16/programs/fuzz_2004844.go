package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 string }
type GenS5 struct { F0 int16; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -2147483648
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int64 = 42
        fmt.Println(v6, v3, v3)
        _ = fi5
        _ = v6
    }
    var v7 bool = false
    switch -482 {
    case -658:
        v8 := GenS1{F0: 127}
        _ = v8
    case -2147483648:
        if (42 > 42) {
            var v9 uint = 2
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(v7)
                if (false && (v3 != -289)) {
                    v9 -= v9
                    switch 510 {
                    case 128:
                        if (v7 && ((^(v9)) <= (v9 + (^(v9))))) {
                            v11 := GenS2{F0: v3, F1: "hello"}
                            _ = v11
                        } else {
                            if (fi10 > 249) {
                                v12 := GenS3{F0: 127}
                                _ = v12
                            } else {
                                for fi13 := 0; fi13 < 11; fi13++ {
                                    v7 = (!v7)
                                    _ = fi13
                                }
                                v14 := []int{}
                                _ = v14
                            }
                            v15 := v2[2]
                            switch ("a" + "a") {
                            case "hello":
                                v9 -= (v9 ^ (v9 + v9))
                            case "\n\t":
                                for fi16 := 0; fi16 < 17; fi16++ {
                                    v9 = v9
                                    var v17 string = "\n\t"
                                    v17 = v17
                                    var v18 int64 = -9223372036854775808
                                    var v19 float64 = (0.1 + 1.5)
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                            case "a":
                                v20 := GenS3{F0: 187}
                                _ = v20
                            }
                            _ = v15
                        }
                    case -14:
                        v4 -= (v4 + v4)
                    case 42:
                        if (!(!v7)) {
                            v9 = 65535
                        } else {
                            v21 := GenS4{F0: (v4 / (((v4 * v4)) | 1)), F1: ("a" + ""), F2: "hello"}
                            _ = v21
                        }
                    default:
                        v3 = (v3 ^ int32(v4))
                    }
                    for fi22 := 0; fi22 < 6; fi22++ {
                        var v23 uint16 = v4
                        _ = fi22
                        _ = v23
                    }
                }
                _ = fi10
            }
            var v24 uint16 = (v4 & v4)
            _ = v24
            _ = v9
        }
    }
    v3 -= 1
    switch ("hello" + "a") {
    case "":
        switch ("" + ("hello" + "a")) {
        case "a":
            fmt.Println(v7)
        case "foo bar":
            var v25 int64 = -558
            _ = v25
        case "\n\t":
            var v26 bool = (!v7)
            _ = v26
        }
    }
    v27 := GenS4{F0: v4, F1: "hello", F2: "\n\t"}
    var v28 string = "\n\t"
    v29 := GenS2{F0: (v3 - 205), F1: v28}
    if ((!v7) && v7) {
        v2[748] = 4294967295
        for fi30 := 0; fi30 < 9; fi30++ {
            v28 = ""
            v4 = 335
            for fi31 := 0; fi31 < 10; fi31++ {
                v3 = (v3 ^ -161)
                var v32 float32 = (float32(float64(0.1)) - (0.1 - (-2.25 - 0.0)))
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        if (593 > 42) {
            v3 += v3
            var v33 float64 = 0.1
            for fi34 := 0; fi34 < 10; fi34++ {
                v35 := GenS3{F0: 85}
                var v36 string = ((v28 + "") + "a")
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v33
        }
    }
    v37 := GenS4{F0: (v4 / ((v4) | 1)), F1: (v28 + v28), F2: v28}
    if (false && false) {
        fmt.Println(v7)
        switch v28 {
        case "\n\t":
            v7 = (-1 > -1)
        }
    } else {
        v38 := GenS5{F0: 0, F1: 42, F2: (v28 + ("hello" + (v28 + ((v28 + v28) + "hello"))))}
        v7 = (756 <= 0)
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
    _ = v7
}
