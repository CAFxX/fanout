package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 uint32) int {
    var v6 int8 = -1
    v6 += v6
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{"foo bar"}
    var v3 uint64 = 0
    var v4 int16 = 32767
    var v5 uint = 543
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 += v1
        v5 -= 755
        _ = fi7
    }
    if len(v2) > 0 {
        v2[0] = (("a" + "foo bar") + "foo bar")
    }
    fmt.Println(v3)
    v8 := GenS1{F0: (!true), F1: (v5 ^ v5)}
    var v9 int8 = 37
    v10 := GenS1{F0: false, F1: v5}
    v2 = append(v2, "foo bar")
    for fi11 := 0; fi11 < 5; fi11++ {
        v1 -= 1e10
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: true, F1: 414}
            var v14 int16 = (v4 * (v4 % (((v4 & v4)) | 1)))
            switch (fi11 % (((-(fi11))) | 1)) {
            case 32767:
                v2 = append(v2, "")
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        if true {
            var v15 float32 = 1e10
            if (!("a" != ("" + "foo bar"))) {
                v3 -= v3
                for fi16 := 0; fi16 < 13; fi16++ {
                    v4 = int16(fi16)
                    var v17 uint32 = 997
                    if (int(v4) <= fi11) {
                        if true {
                            var v18 int8 = int8(fi11)
                            var v19 uint = v5
                            _ = len(v2)
                            _ = v18
                            _ = v19
                        } else {
                            v17 = (v17 / ((v17) | 1))
                        }
                    } else {
                        fmt.Println(fi11, v15, fi16)
                    }
                    _ = fi16
                    _ = v17
                }
                v5 = v5
            }
            if (!true) {
                v15 = v15
                var v20 int16 = (v4 * v4)
                _ = v20
            }
            _ = v15
        } else {
            var v21 float32 = v1
            for fi22 := 0; fi22 < 18; fi22++ {
                v21 -= v1
                switch fi11 {
                case 32768:
                    if true {
                        if false {
                            _ = len(v2)
                            v3 = (v3 + v3)
                            switch genFn0(2) {
                            case 0:
                                v23 := GenS1{F0: (2147483647 != 65535), F1: v5}
                                _ = v23
                            case 145:
                                v24 := GenS1{F0: false, F1: (v5 | v5)}
                                _ = v24
                            }
                        }
                        switch (fi22 & -339) {
                        case 0:
                            if false {
                                fmt.Println(fi11, v9, v1)
                                fmt.Println(fi22)
                                if false {
                                    v4 += 0
                                    fmt.Println(v21, v3)
                                    fmt.Println(v3, v4, v5)
                                }
                            }
                        case 9223372036854775807:
                            var v25 uint64 = 501
                            _ = v25
                        }
                    }
                case 42:
                    fmt.Println(fi11)
                }
                _ = fi22
            }
            fmt.Println(v1, v9)
            _ = v21
        }
        _ = fi11
    }
    v5 -= 2
    switch "a" {
    case "hello":
        v4 = int16(v3)
    default:
        switch ("hello" + "a") {
        case "foo bar":
            if true {
                var v26 string = "a"
                _ = v26
            }
        case "":
            v4 = v4
        default:
            var v27 int16 = (v4 | (-(v4)))
            _ = v27
        }
    }
    var v28 int64 = 393
    for fi29 := 0; fi29 < 7; fi29++ {
        if ((2 >= 255) && (("" < "") || ((v4 | v4) > v4))) {
            for fi30 := 0; fi30 < 15; fi30++ {
                var v31 uint64 = v3
                var v32 uint = 65535
                v33 := GenS1{F0: false, F1: 18446744073709551615}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            var v34 int32 = 2147483647
            v35 := GenS2{F0: 1}
            _ = v34
            _ = v35
        }
        var v36 float32 = 1e10
        var v37 int = -9223372036854775808
        _ = fi29
        _ = v36
        _ = v37
    }
    if ((^(v3)) < v3) {
        v38 := GenS3{F0: "", F1: "foo bar"}
        _ = v38
    } else {
        if len(v2) > 0 {
            v2[0] = ("\n\t" + ("" + ("\n\t" + "foo bar")))
        }
        var v39 uint64 = (v3 & (v3 ^ v3))
        _ = v39
    }
    fmt.Println(v1, v28, v5)
    v2 = append(v2, ("\n\t" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
