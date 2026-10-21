package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS8 struct { F0 string; F1 uint32; F2 int }
type GenS9 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS10 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 int64 = 1
    v4 := make(map[uint32]string)
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            if false {
                var v7 int = fi6
                v3 -= -9223372036854775808
                v8 := GenS1{F0: 0, F1: 42}
                _ = v7
                _ = v8
            } else {
                v3 -= (v3 - -148)
                v3 = (v3 | (v3 | (v3 + (v3 * (v3 - -9223372036854775808)))))
                v3 = -363
            }
            var v9 uint32 = 4294967295
            v10 := GenS1{F0: 18446744073709551615, F1: fi5}
            _ = fi6
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: ((true || false) && (-1 >= -32768)), F1: 2}
        _ = fi5
        _ = v11
    }
    v12 := []int{4294967295, 774, -796, -576}
    if false {
        if (-115 < -128) {
            v12 = append(v12, 4294967295)
        }
    } else {
        if true {
            switch "" {
            case "foo bar":
                for fi13 := 0; fi13 < 5; fi13++ {
                    fmt.Println(v3, v3)
                    var v14 int16 = 1
                    fmt.Println(v3, v3)
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                if ((((float32((1.5 - 3.14159)) * 0.1) - 0.0) - (0.1 - 1e10)) < (1e10 * (0.1 - 1.5))) {
                    _ = len(v4)
                    if (3.14159 != float32((0.0 + (-2.25 / float64(1.5))))) {
                        var v15 uint16 = 338
                        _ = len(v12)
                        _ = v15
                    }
                    _ = len(v4)
                } else {
                    if (("a" + "hello") != "foo bar") {
                        var v16 float64 = -2.25
                        if len(v12) > 0 {
                            v12[0] = 0
                        }
                        _ = v16
                    }
                }
            default:
                var v17 string = "foo bar"
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 11; fi18++ {
                v3 = v3
                _ = fi18
            }
            v19 := GenS2{F0: false, F1: 42}
            _ = v19
        }
    }
    v3 = v3
    v3 += (v3 % ((v3) | 1))
    var v20 uint = 278
    v21 := GenS3{F0: "hello", F1: int8(v3)}
    var v22 float32 = (-2.25 * 1e10)
    switch "a" {
    case "hello":
        delete(v2, ("hello" + ("\n\t" + "foo bar")))
    case "":
        v3 += v3
    case "a":
        fmt.Println(v22, v3)
    default:
        if (-1 < -542) {
            for fi23 := 0; fi23 < 5; fi23++ {
                if (251 <= 42) {
                    v3 = 0
                    var v24 float32 = v22
                    _ = v24
                } else {
                    switch "" {
                    case "":
                        v25 := GenS4{F0: 1000}
                        _ = v25
                    default:
                        var v26 int32 = 960
                        _ = v26
                    }
                    _ = len(v12)
                    if (!(42 != 21)) {
                        v27 := v1[65535]
                        v28 := GenS5{F0: 42}
                        v22 += (v22 + (v22 + (v22 / v22)))
                        _ = v27
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 15; fi29++ {
                            v22 = v22
                            v3 += v3
                            _ = fi29
                        }
                        delete(v2, "hello")
                        _ = len(v2)
                    }
                }
                _ = fi23
            }
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 int64 = 787
                switch int(v31) {
                case -1:
                    v32 := GenS6{F0: 66, F1: 339, F2: 2}
                    _ = v32
                default:
                    v20 = 1
                }
                v12 = append(v12, fi30)
                _ = fi30
                _ = v31
            }
            v33 := GenS7{F0: uint8(v20), F1: -48, F2: 42}
            _ = len(v4)
            _ = v33
        }
    }
    v34 := v4[65535]
    var v35 int8 = 127
    var v36 uint16 = 255
    var v37 float32 = float32(0.0)
    var v38 string = "\n\t"
    switch -997 {
    case 42:
        for fi39 := 0; fi39 < 20; fi39++ {
            for fi40 := 0; fi40 < 15; fi40++ {
                if ("" < ((v38 + ("hello" + v34)) + "foo bar")) {
                    v1[2147483647] = fi39
                    v41 := GenS8{F0: v38, F1: 65535, F2: (fi39 / ((fi40) | 1))}
                    _ = v41
                } else {
                    switch (fi40 ^ (fi40 % ((fi39) | 1))) {
                    case -721:
                        v34 = v34
                    }
                    v42 := GenS4{F0: 1000}
                    _ = v42
                }
                v43 := GenS1{F0: v20, F1: fi39}
                _ = fi40
                _ = v43
            }
            _ = fi39
        }
    case -213:
        var v44 uint16 = v36
        _ = v44
    case 4294967295:
        if (42 == -480) {
            v45 := GenS5{F0: -2147483648}
            v46 := GenS9{F0: 2147483648, F1: 2, F2: v35}
            _ = v45
            _ = v46
        } else {
            v47 := GenS7{F0: 40, F1: v35, F2: 2}
            _ = v47
        }
    }
    v48 := GenS10{F0: 127, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v48
}
