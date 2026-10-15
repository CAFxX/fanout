package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int; F1 int }
func genFn0() int8 {
    if false {
        var v7 float32 = float32(float64(-2.25))
        _ = v7
    } else {
        switch "a" {
        case "foo bar":
            var v8 uint64 = 2
            _ = v8
        case "":
            var v9 uint = 0
            _ = v9
        default:
            var v10 uint8 = 255
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: ""}
        _ = v11
    }
    if true {
        var v12 uint64 = 1
        var v13 uint32 = 216
        fmt.Println(v12, v12)
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 string = "a"
            for fi16 := 0; fi16 < 10; fi16++ {
                switch (((v15 + "") + "foo bar") + "") {
                case "a":
                    var v17 int64 = 9223372036854775807
                    _ = v17
                case "hello":
                    switch fi14 {
                    case 32767:
                        fmt.Println(v15)
                        v18 := GenS1{F0: 51, F1: "hello"}
                        v15 = ("hello" + ("a" + (v15 + v15)))
                        v19 := GenS1{F0: 59, F1: (("a" + v15) + (v15 + v15))}
                        _ = v18
                        _ = v19
                    case -1:
                        v20 := GenS1{F0: 0, F1: "foo bar"}
                        var v21 bool = ((32767 >= 32767) && true)
                        v21 = false
                        switch (v15 + "a") {
                        case "hello":
                            v15 = v15
                        }
                        var v22 float32 = (1e10 - 0.1)
                        v21 = v21
                        _ = v20
                        _ = v21
                        _ = v22
                    default:
                        switch (fi14 & (-(fi16))) {
                        case -468:
                            v15 = ("" + v15)
                        case -1:
                            v15 = ""
                        case -2147483648:
                            if false {
                                if false {
                                    fmt.Println(fi14)
                                    v15 = (v15 + (v15 + ((("foo bar" + (v15 + "foo bar")) + (("foo bar" + v15) + (v15 + "\n\t"))) + v15)))
                                    v15 = "foo bar"
                                    v23 := []int{fi14}
                                    _ = v23
                                } else {
                                    fmt.Println(v15)
                                    v15 = ((((v15 + v15) + "") + (v15 + v15)) + v15)
                                }
                                var v24 int32 = 698
                                _ = v24
                            }
                        }
                    }
                case "foo bar":
                    var v25 uint = 255
                    _ = v25
                }
                var v26 string = v15
                _ = fi16
                _ = v26
            }
            _ = fi14
            _ = v15
        }
    }
    v27 := GenS2{F0: 128, F1: 32767, F2: -1}
    _ = v27
    return 1
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]string)
    v3 := []string{}
    var v4 uint8 = 149
    var v5 int8 = -14
    v6 := make(map[uint32]string)
    v28 := GenS1{F0: 100, F1: (("" + ("foo bar" + "a")) + "\n\t")}
    v5 = (v5 + (v5 * (v5 * 1)))
    v5 = (^(v5))
    fmt.Println(v5, v4)
    v29 := GenS2{F0: 32768, F1: 32767, F2: 0}
    if false {
        for fi30 := 0; fi30 < 15; fi30++ {
            v5 = (v5 >> ((65535) & 63))
            _ = fi30
        }
    }
    v4 += v4
    v5 -= genFn0()
    var v31 int64 = 0
    if (-123 >= -1) {
        v32 := GenS1{F0: (^(v5)), F1: "foo bar"}
        var v33 uint32 = 1
        switch "" {
        case "\n\t":
            v31 -= v31
        case "foo bar":
            v31 -= v31
        case "hello":
            v34 := GenS3{F0: (^(v4)), F1: 2}
            _ = v34
        default:
            v31 -= (v31 >> ((uint(v4)) & 63))
        }
        _ = v32
        _ = v33
    }
    fmt.Println(v4)
    var v35 int32 = 877
    var v36 float64 = float64(3.14159)
    for fi37 := 0; fi37 < 5; fi37++ {
        v5 -= (-(v5))
        for fi38 := 0; fi38 < 1; fi38++ {
            v6[15] = "\n\t"
            if (0 < 419) {
                v35 = v35
            }
            _ = fi38
        }
        switch fi37 {
        case -1:
            for fi39 := 0; fi39 < 12; fi39++ {
                _ = len(v3)
                var v40 float32 = 1e10
                _ = fi39
                _ = v40
            }
        case 32767:
            for fi41 := 0; fi41 < 5; fi41++ {
                var v42 float64 = 0.1
                var v43 float64 = (0.0 + 3.14159)
                fmt.Println(v43, v43, v31)
                _ = fi41
                _ = v42
                _ = v43
            }
        case 32768:
            var v44 uint16 = 605
            _ = v44
        default:
            if true {
                for fi45 := 0; fi45 < 12; fi45++ {
                    v3 = append(v3, "hello")
                    var v46 string = "hello"
                    _ = fi45
                    _ = v46
                }
                v47 := GenS4{F0: fi37, F1: fi37}
                _ = v47
            } else {
                var v48 int64 = (v31 & v31)
                _ = v48
            }
        }
        _ = fi37
    }
    v31 -= (v31 % (((-(v31))) | 1))
    for fi49 := 0; fi49 < 2; fi49++ {
        v5 = (v5 << ((484) & 63))
        _ = fi49
    }
    var v50 uint16 = 270
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
