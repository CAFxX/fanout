package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int32; F1 int16; F2 int16 }
type GenS4 struct { F0 int32; F1 bool; F2 uint8 }
type GenS5 struct { F0 int64; F1 int32; F2 int64 }
type GenS6 struct { F0 string; F1 int; F2 bool }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    var v3 uint32 = 4294967295
    v4 := GenS1{F0: 2, F1: 100, F2: true}
    if (true && ((((42 <= 42) && false) && ((!(1 > 1)) && true)) && (1 > 18446744073709551615))) {
        v5 := GenS2{F0: -2147483648, F1: 128, F2: 0}
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float64 = 3.14159
            v8 := GenS2{F0: -1, F1: 1, F2: 882}
            v9 := []int{fi6}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 string = "a"
        _ = v10
        _ = v5
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 uint32 = v3
                v13 = v13
                var v14 int = 32768
                _ = fi12
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS3{F0: -851, F1: 32767, F2: int16(fi15)}
                var v17 int8 = 1
                switch "a" {
                case "a":
                    if false {
                        if (false && false) {
                            fmt.Println(v17, fi11, v1)
                            v2[65535] = (fi11 + 128)
                        } else {
                            var v18 int = (^(fi15))
                            v19 := GenS2{F0: 2147483647, F1: 2, F2: int16(fi11)}
                            v3 = 1
                            _ = v18
                            _ = v19
                        }
                    }
                default:
                    var v20 float64 = -2.25
                    _ = v20
                }
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                v3 += (v3 + (v3 + (v3 % ((v3) | 1))))
                _ = fi21
            }
            var v22 int8 = 42
            _ = v22
        }
        var v23 int64 = 9223372036854775807
        v23 -= v23
        _ = fi11
        _ = v23
    }
    v3 -= (v3 / ((2147483647) | 1))
    _ = len(v2)
    for fi24 := 0; fi24 < 18; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            v26 := GenS3{F0: 1, F1: 0, F2: 566}
            if ((fi25 ^ -323) <= fi25) {
                var v27 int64 = 1
                var v28 uint8 = 0
                switch ("foo bar" + ((("" + (("hello" + "\n\t") + ("hello" + "a"))) + "\n\t") + ("a" + ("a" + "")))) {
                case "foo bar":
                    v27 = v27
                    v29 := GenS4{F0: 42, F1: false, F2: (^(v28))}
                    var v30 int8 = 42
                    _ = v29
                    _ = v30
                case "hello":
                    v28 = v28
                case "\n\t":
                    var v31 uint32 = (v3 - (v3 - (v3 | (v3 << ((1) & 63)))))
                    _ = v31
                default:
                    if (true || false) {
                        switch 764 {
                        case 42:
                            v32 := v2[v3]
                            _ = v32
                        case -9223372036854775808:
                            v3 += (v3 % ((v3) | 1))
                            var v33 float32 = (v1 + v1)
                            var v34 int64 = v27
                            _ = v33
                            _ = v34
                        case -964:
                            v1 += v1
                        }
                        if true {
                            if true {
                                var v35 int8 = 1
                                _ = v35
                            }
                        }
                        switch 128 {
                        case 42:
                            fmt.Println(fi24, fi24, v3)
                        case 0:
                            v36 := []int{(fi25 % ((fi24) | 1))}
                            _ = v36
                        }
                    } else {
                        var v37 uint8 = (^(v28))
                        v38 := GenS5{F0: (^(v27)), F1: -1, F2: v27}
                        _ = v37
                        _ = v38
                    }
                    v27 += v27
                }
                v1 += float32((1e10 / ((1.5 * 1e10) + 1e10)))
                fmt.Println(v3, v28)
                _ = v27
                _ = v28
            } else {
                v39 := GenS6{F0: "foo bar", F1: 42, F2: (!(!false))}
                v40 := []int{fi25, (fi24 / ((4294967295) | 1)), 42}
                fmt.Println(v3, fi24)
                _ = v39
                _ = v40
            }
            _ = fi25
            _ = v26
        }
        v41 := GenS5{F0: -1, F1: -1, F2: -285}
        _ = fi24
        _ = v41
    }
    var v42 bool = (v3 != v3)
    delete(v2, v3)
    v3 = 1
    v3 -= (v3 >> ((18446744073709551615) & 63))
    if v42 {
        switch 42 {
        case 1:
            if (true && v42) {
                var v43 int64 = -1
                _ = v43
            }
        case 9223372036854775807:
            fmt.Println(v42, v42, v42)
        case -9223372036854775808:
            if v42 {
                v44 := []int{}
                _ = v44
            }
        default:
            fmt.Println(v42)
        }
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
