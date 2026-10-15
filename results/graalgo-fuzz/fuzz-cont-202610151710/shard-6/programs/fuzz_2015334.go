package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0() uint16 {
    if (530 <= 0) {
        var v7 uint64 = 0
        v7 += (v7 - v7)
        _ = v7
    } else {
        if true {
            var v8 uint64 = 2
            var v9 int8 = -128
            switch 68 {
            case 956:
                if false {
                    var v10 float32 = (0.0 + (3.14159 - ((3.14159 * float32(-2.25)) - 0.0)))
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v10 = v10
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v13 := []int{(fi12 * int(v8)), (fi12 + (fi12 + 507)), fi12, fi12}
                        switch ("a" + "hello") {
                        case "hello":
                            v14 := GenS1{F0: 1}
                            _ = v14
                        }
                        v13 = append(v13, fi12)
                        _ = fi12
                        _ = v13
                    }
                    _ = v10
                }
            case -9223372036854775808:
                if (!(4294967295 <= 2147483647)) {
                    switch 32767 {
                    case 32768:
                        for fi15 := 0; fi15 < 14; fi15++ {
                            var v16 uint64 = v8
                            var v17 int32 = 42
                            switch ("hello" + "foo bar") {
                            case "":
                                fmt.Println(v16, v16)
                            default:
                                for fi18 := 0; fi18 < 2; fi18++ {
                                    if true {
                                        v9 += (v9 - v9)
                                        fmt.Println(v16)
                                    } else {
                                        v19 := GenS1{F0: 2}
                                        _ = v19
                                    }
                                    var v20 int64 = -662
                                    _ = fi18
                                    _ = v20
                                }
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case -2147483648:
                        v21 := []int{}
                        _ = v21
                    case 1:
                        var v22 int32 = 2147483647
                        _ = v22
                    }
                } else {
                    var v23 uint8 = 0
                    fmt.Println(v9, v8, v9)
                    v23 = 171
                    _ = v23
                }
            case 42:
                var v24 uint = 255
                _ = v24
            default:
                var v25 int16 = 1000
                _ = v25
            }
            _ = v8
            _ = v9
        } else {
            var v26 uint = 42
            _ = v26
        }
        var v27 float64 = -2.25
        _ = v27
    }
    var v28 uint16 = 255
    var v29 int16 = 1
    _ = v28
    _ = v29
    return (v28 - (v28 / (((v28 - 255)) | 1)))
}

func main() {
    v1 := []int{32767, 0, -2147483648, 4294967295}
    v2 := []int{214, 4294967295, 9223372036854775807}
    var v3 int16 = 1
    var v4 uint64 = 363
    var v5 uint8 = 2
    var v6 string = ("\n\t" + "\n\t")
    var v30 uint32 = 1
    v31 := GenS1{F0: v5}
    for fi32 := 0; fi32 < 16; fi32++ {
        v33 := GenS2{F0: 0}
        for fi34 := 0; fi34 < 7; fi34++ {
            if (!false) {
                var v35 uint16 = 0
                _ = v35
            } else {
                v4 += v4
            }
            if len(v2) > 0 {
                v2[0] = 564
            }
            var v36 bool = true
            _ = fi34
            _ = v36
        }
        v6 = "\n\t"
        _ = fi32
        _ = v33
    }
    v37 := GenS3{F0: 2}
    v38 := GenS2{F0: 0}
    v39 := GenS1{F0: 128}
    v4 -= (v4 ^ (v4 / ((v4) | 1)))
    var v40 int32 = 2147483647
    v5 = 255
    var v41 uint = 65535
    if ((false || (!(42 < genFn0()))) && true) {
        var v42 uint32 = (v30 - 4294967295)
        if true {
            fmt.Println(v3)
            v43 := GenS2{F0: int8(v4)}
            _ = v43
        }
        _ = v42
    } else {
        var v44 uint = 646
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
