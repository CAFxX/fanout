package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8; F1 string; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 string; F1 int8; F2 uint16 }
func genFn0(p0 uint16, p1 uint64) int32 {
    var v5 int8 = -93
    fmt.Println(p0, p0, p1)
    v5 -= 96
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() uint8 {
    if true {
        v6 := []int{42, -1}
        var v7 float64 = 1.5
        _ = v6
        _ = v7
    } else {
        if true {
            var v8 uint32 = 42
            var v9 int8 = -128
            v8 = v8
            _ = v8
            _ = v9
        } else {
            switch -936 {
            case 4294967295:
                var v10 uint = 411
                _ = v10
            case 582:
                for fi11 := 0; fi11 < 8; fi11++ {
                    v12 := []int{(fi11 | -2147483648)}
                    var v13 uint64 = 42
                    if (0 < 2) {
                        if true {
                            if (!(v13 <= (v13 & 1))) {
                                if (-2147483648 < (-(fi11))) {
                                    var v14 int8 = 127
                                    v15 := GenS1{F0: false, F1: fi11}
                                    fmt.Println(v14, fi11, fi11)
                                    _ = v14
                                    _ = v15
                                }
                                var v16 uint16 = 42
                                v12 = append(v12, fi11)
                                _ = v16
                            }
                        }
                    } else {
                        v17 := GenS2{F0: 91, F1: ((("foo bar" + "hello") + "\n\t") + ""), F2: 2}
                        if true {
                            switch "hello" {
                            case "":
                                var v18 int64 = 9223372036854775807
                                _ = v18
                            case "foo bar":
                                v13 += 647
                            default:
                                var v19 string = ("\n\t" + "hello")
                                _ = v19
                            }
                            var v20 string = "foo bar"
                            var v21 uint16 = 982
                            _ = v20
                            _ = v21
                        }
                        if false {
                            v13 = (v13 ^ v13)
                            v13 = (v13 | 42)
                        } else {
                            if true {
                                v13 += v13
                                switch (fi11 % ((-1) | 1)) {
                                case 2147483647:
                                    var v22 int16 = 419
                                    _ = v22
                                case 32767:
                                    v12 = append(v12, (-(fi11)))
                                case -9223372036854775808:
                                    var v23 uint32 = 65535
                                    _ = v23
                                }
                                var v24 uint16 = 705
                                _ = v24
                            } else {
                                if (-41 == -128) {
                                    v13 = 18446744073709551615
                                }
                            }
                            var v25 int8 = 127
                            _ = v25
                        }
                        _ = v17
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 2147483647:
                var v26 int16 = 1000
                _ = v26
            default:
                var v27 uint32 = 1
                _ = v27
            }
        }
    }
    var v28 int32 = 42
    for fi29 := 0; fi29 < 5; fi29++ {
        v28 += v28
        _ = fi29
    }
    v28 -= (v28 * (v28 | v28))
    if true {
        var v30 string = (("a" + "hello") + "a")
        _ = v30
    } else {
        v31 := GenS1{F0: false, F1: 2147483647}
        _ = v31
    }
    _ = v28
    return 1
}

func main() {
    v1 := []uint32{2147483648, 4294967295}
    v2 := make(map[string]bool)
    var v3 int16 = 42
    v4 := []int{-2147483648, 4294967295, -339, 4294967295}
    var v32 int = 147
    var v33 int = (v32 << ((42) & 63))
    if (!true) {
        switch 25 {
        case -9223372036854775808:
            var v34 int64 = -539
            _ = v34
        case 340:
            var v35 uint32 = 0
            _ = v35
        case 0:
            if true {
                v36 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: -9223372036854775808}
                _ = v36
            }
        }
        v37 := GenS3{F0: -2147483648, F1: 18446744073709551615, F2: -752}
        _ = v37
    } else {
        if (!(0 < 1)) {
            var v38 string = "a"
            var v39 int32 = 1
            switch 32 {
            case 9223372036854775807:
                var v40 int8 = 8
                _ = v40
            default:
                if (!(true || (genFn1() == 128))) {
                    var v41 int = v33
                    fmt.Println(v3, v32)
                    var v42 uint = 382
                    _ = v41
                    _ = v42
                }
            }
            _ = v38
            _ = v39
        }
    }
    fmt.Println(v32, v3, v33)
    for fi43 := 0; fi43 < 18; fi43++ {
        v44 := GenS3{F0: 42, F1: 2, F2: -147}
        _ = len(v2)
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 int = (v33 & v33)
        v46 -= v33
        _ = fi45
        _ = v46
    }
    v47 := GenS4{F0: "foo bar", F1: -57, F2: 981}
    if (!((0.1 - -2.25) >= (float64((0.1 - 1.5)) + (-2.25 / ((1e10 * 1.5) + 3.14159))))) {
        delete(v2, "hello")
    }
    if (!false) {
        for fi48 := 0; fi48 < 9; fi48++ {
            for fi49 := 0; fi49 < 14; fi49++ {
                var v50 int64 = 9223372036854775807
                var v51 uint = 255
                var v52 uint64 = 528
                _ = fi49
                _ = v50
                _ = v51
                _ = v52
            }
            _ = fi48
        }
    }
    var v53 string = "\n\t"
    v53 = "a"
    v53 = (((v53 + "\n\t") + (v53 + "\n\t")) + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v47
    _ = v53
}
