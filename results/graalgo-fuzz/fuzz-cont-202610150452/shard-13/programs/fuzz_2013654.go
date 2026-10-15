package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint8 }
func genFn0() uint8 {
    var v6 bool = (255 < 255)
    v7 := GenS1{F0: 0}
    var v8 uint8 = 1
    var v9 uint32 = 601
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v8 - (v8 | (v8 - 127)))
}
func genFn1(p0 int64, p1 int8) int32 {
    fmt.Println(p1)
    for fi10 := 0; fi10 < 14; fi10++ {
        if true {
            if ((true && true) && true) {
                var v11 uint64 = 2
                _ = v11
            }
            p0 += p0
            switch fi10 {
            case 2147483647:
                fmt.Println(p0)
            default:
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 bool = false
                    if (-93 <= 42) {
                        switch "\n\t" {
                        case "a":
                            v14 := GenS1{F0: 2}
                            _ = v14
                        }
                    } else {
                        v15 := GenS2{F0: 42, F1: 2, F2: 856}
                        _ = v15
                    }
                    switch "foo bar" {
                    case "":
                        v16 := []int{0, (^(fi12)), 686, -912}
                        _ = v16
                    case "\n\t":
                        p1 = p1
                    case "hello":
                        var v17 int16 = 57
                        _ = v17
                    }
                    _ = fi12
                    _ = v13
                }
            }
        } else {
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS2{F0: 1, F1: 255, F2: 0}
                _ = fi18
                _ = v19
            }
        }
        switch -2147483648 {
        case 0:
            fmt.Println(p1)
        }
        _ = fi10
    }
    if false {
        if false {
            p0 -= p0
        } else {
            v20 := GenS1{F0: 255}
            p0 += -885
            fmt.Println(p0, p0)
            _ = v20
        }
        switch ("" + "a") {
        case "hello":
            for fi21 := 0; fi21 < 16; fi21++ {
                p0 = (p0 >> ((2) & 63))
                _ = fi21
            }
        case "\n\t":
            v22 := []int{-9223372036854775808, -9223372036854775808}
            _ = v22
        default:
            p0 -= -9223372036854775808
        }
    } else {
        v23 := []int{0, 123}
        v23 = append(v23, 4294967295)
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                var v26 uint8 = 128
                fmt.Println(v26, fi25, fi24)
                var v27 uint8 = (^(v26))
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
        _ = v23
    }
    v28 := GenS3{F0: 127, F1: 42}
    p1 = (p1 + 0)
    _ = p0
    _ = p1
    _ = v28
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 42
    v3 := []int{241}
    v4 := make(map[string]int)
    var v5 uint64 = 2
    v2 -= (^(v2))
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS2{F0: 42, F1: 0, F2: 475}
        if false {
            fmt.Println(fi29, v5, v2)
            if (!(1 != -686)) {
                v31 := v1[(fi29 + (fi29 >> ((0) & 63)))]
                var v32 uint32 = v31
                if true {
                    v33 := v1[fi29]
                    v31 = (v31 << ((338) & 63))
                    var v34 int64 = -1
                    _ = v33
                    _ = v34
                }
                _ = v31
                _ = v32
            }
            for fi35 := 0; fi35 < 18; fi35++ {
                v5 = (v5 | v5)
                v36 := GenS3{F0: 128, F1: uint16(v2)}
                v5 -= v5
                _ = fi35
                _ = v36
            }
        }
        v2 -= v2
        _ = fi29
        _ = v30
    }
    fmt.Println(v5, v2)
    var v37 bool = true
    var v38 uint16 = 0
    fmt.Println(v37)
    if v37 {
        v5 = (v5 % ((18446744073709551615) | 1))
    } else {
        switch "\n\t" {
        case "foo bar":
            var v39 uint32 = 0
            _ = v39
        case "":
            var v40 uint16 = v38
            _ = v40
        default:
            var v41 bool = ((^(v2)) < 246)
            _ = v41
        }
        v2 = v2
        if v37 {
            if (true && v37) {
                var v42 uint64 = (v5 % ((4294967295) | 1))
                var v43 uint64 = 4294967295
                if (genFn1(-516, 127) <= 2147483647) {
                    var v44 int = 42
                    var v45 string = "foo bar"
                    var v46 int16 = 1
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v47 := GenS4{F0: genFn1(-485, int8(v43)), F1: 0}
                    _ = v47
                }
                _ = v42
                _ = v43
            }
        }
    }
    fmt.Println(v2)
    fmt.Println(v2, v38, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
