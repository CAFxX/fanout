package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint64; F1 string }

func main() {
    v1 := []string{"\n\t", ("\n\t" + ("hello" + "")), "", "hello"}
    var v2 uint32 = 2147483647
    v3 := []int{}
    var v4 uint16 = 255
    v5 := GenS1{F0: 9223372036854775807}
    if true {
        switch (("a" + "hello") + "hello") {
        case "":
            v6 := GenS1{F0: 0}
            _ = v6
        default:
            var v7 uint64 = 2
            _ = v7
        }
        v8 := GenS1{F0: -1}
        _ = v8
    }
    v2 = (v2 & 4294967295)
    var v9 int8 = -15
    var v10 uint16 = 1
    fmt.Println(v10, v4, v4)
    v9 -= v9
    if true {
        var v11 int16 = 1
        _ = v11
    } else {
        var v12 string = "a"
        if (-753 >= 42) {
            switch v12 {
            case "hello":
                var v13 bool = (1 > 861)
                _ = v13
            case "\n\t":
                v3 = append(v3, 32768)
            default:
                if ((v2 << ((18446744073709551615) & 63)) < (v2 * (v2 + 4294967295))) {
                    var v14 string = "hello"
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int64 = 319
                        switch (fi15 - (fi15 << ((343) & 63))) {
                        case 1:
                            v12 = v12
                        case 32768:
                            var v17 uint64 = 673
                            _ = v17
                        case 127:
                            switch (fi15 / ((32768) | 1)) {
                            case -26:
                                var v18 float64 = 1e10
                                _ = v18
                            case 301:
                                if (0.1 >= 1.5) {
                                    for fi19 := 0; fi19 < 12; fi19++ {
                                        v12 = (v12 + ("a" + v12))
                                        _ = fi19
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (v12 + "foo bar")
                                    }
                                } else {
                                    var v20 uint16 = (v4 & v10)
                                    v21 := GenS1{F0: v16}
                                    _ = v20
                                    _ = v21
                                }
                            default:
                                var v22 int32 = 42
                                _ = v22
                            }
                        default:
                            v23 := GenS2{F0: 1, F1: "", F2: v9}
                            _ = v23
                        }
                        var v24 float64 = 0.1
                        _ = fi15
                        _ = v16
                        _ = v24
                    }
                    fmt.Println(v10)
                }
            }
            var v25 uint64 = 2
            _ = v25
        } else {
            var v26 uint64 = 346
            _ = v26
        }
        v4 = v10
        _ = v12
    }
    if (1e10 == 0.0) {
        fmt.Println(v2, v2)
    } else {
        var v27 uint64 = 511
        if true {
            if len(v1) > 0 {
                v1[0] = ""
            }
        } else {
            v28 := GenS2{F0: 717, F1: "hello", F2: (v9 + v9)}
            if ((526 >= 9223372036854775807) && (!true)) {
                fmt.Println(v2, v9)
            } else {
                if len(v1) > 0 {
                    v1[0] = (("" + "") + "a")
                }
                fmt.Println(v9)
            }
            switch -2147483648 {
            case 1:
                if true {
                    v29 := GenS1{F0: 1}
                    _ = v29
                }
            case -223:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 9223372036854775807:
                v27 = v27
            }
            _ = v28
        }
        _ = v27
    }
    v30 := GenS2{F0: 42, F1: ("\n\t" + "a"), F2: v9}
    for fi31 := 0; fi31 < 19; fi31++ {
        v32 := GenS3{F0: fi31, F1: -253}
        for fi33 := 0; fi33 < 8; fi33++ {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            v3 = append(v3, fi33)
            _ = fi33
        }
        v34 := GenS4{F0: 18446744073709551615, F1: ("a" + "")}
        _ = fi31
        _ = v32
        _ = v34
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 float32 = (0.0 - (float32((1.5 + 0.0)) + (1e10 + 1.5)))
    var v36 uint32 = (v2 % (((v2 << ((1) & 63))) | 1))
    v9 = v9
    var v37 uint8 = 127
    var v38 bool = true
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS2{F0: v2, F1: ("foo bar" + "foo bar"), F2: 1}
        v41 := GenS4{F0: 1, F1: "hello"}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
