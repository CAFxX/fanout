package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int16 = -332
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 uint16 = 377
    delete(v2, "\n\t")
    if ((167 <= 4294967295) || true) {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 float64 = 1e10
            v7 := []int{fi5, (fi5 % (((fi5 - -41)) | 1)), (fi5 * 128), (fi5 & fi5)}
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 string = ("a" + (("hello" + "a") + "foo bar"))
        _ = v8
    }
    fmt.Println(v4)
    switch 32767 {
    case -145:
        for fi9 := 0; fi9 < 18; fi9++ {
            if true {
                v10 := []int{(fi9 * (fi9 / (((fi9 | (fi9 & fi9))) | 1))), (fi9 << ((4294967295) & 63)), (-(fi9))}
                _ = v10
            } else {
                for fi11 := 0; fi11 < 9; fi11++ {
                    if (!false) {
                        if false {
                            switch fi11 {
                            case 216:
                                switch (fi9 - 128) {
                                case 4294967295:
                                    var v12 int32 = int32(fi9)
                                    _ = v12
                                case -86:
                                    if false {
                                        _ = len(v2)
                                    }
                                case -820:
                                    if ((v4 / (((^(v4))) | 1)) >= (v4 / ((v4) | 1))) {
                                        fmt.Println(fi11, fi9)
                                        fmt.Println(fi9, v4)
                                    } else {
                                        fmt.Println(v1, fi11)
                                        var v13 uint = 42
                                        var v14 int32 = -795
                                        v15 := []int{-9223372036854775808, (fi11 ^ (fi9 / ((542) | 1)))}
                                        v16 := GenS1{F0: 4294967295}
                                        _ = v13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                }
                            case -2147483648:
                                var v17 int16 = -32768
                                _ = v17
                            default:
                                v18 := GenS1{F0: 197}
                                _ = v18
                            }
                            var v19 int64 = -9223372036854775808
                            delete(v3, "")
                            var v20 int16 = 68
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v1, fi11)
                        }
                        v4 -= v4
                    } else {
                        _ = len(v2)
                        if false {
                            var v21 uint64 = 0
                            switch fi9 {
                            case -1:
                                v4 = (v4 % (((v4 % (((^(v4))) | 1))) | 1))
                            case 4294967295:
                                var v22 uint = 65535
                                _ = v22
                            case -2147483648:
                                v23 := GenS2{F0: -1, F1: 42, F2: 285}
                                _ = v23
                            }
                            _ = v21
                        } else {
                            delete(v3, (("" + "hello") + ("" + "foo bar")))
                        }
                        fmt.Println(fi9)
                        v1 -= v1
                    }
                    _ = fi11
                }
                var v24 int = (fi9 - fi9)
                for fi25 := 0; fi25 < 1; fi25++ {
                    delete(v3, "hello")
                    var v26 string = ("" + "\n\t")
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
            fmt.Println(fi9, fi9, v1)
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS2{F0: 232, F1: 127, F2: 65535}
                v29 := GenS1{F0: 209}
                _ = len(v2)
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi9
        }
    case 32767:
        if (!(!(!false))) {
            if (0 == 42) {
                v30 := GenS1{F0: 330}
                _ = v30
            } else {
                v31 := GenS3{F0: 18446744073709551615, F1: (v1 >> ((4294967295) & 63))}
                _ = v31
            }
            var v32 int8 = -1
            var v33 int64 = 9223372036854775807
            _ = v32
            _ = v33
        } else {
            var v34 int64 = -1
            v35 := []int{}
            if len(v35) >= 2 {
                v35 = v35[1:len(v35)]
            }
            _ = v34
            _ = v35
        }
    case -1:
        var v36 int64 = -9223372036854775808
        _ = v36
    default:
        v4 -= (v4 << ((4294967295) & 63))
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v4 += v4
        var v38 uint64 = 1
        _ = fi37
        _ = v38
    }
    var v39 uint16 = v4
    switch 2147483647 {
    case 32768:
        v40 := GenS1{F0: 713}
        _ = v40
    }
    fmt.Println(v4, v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
