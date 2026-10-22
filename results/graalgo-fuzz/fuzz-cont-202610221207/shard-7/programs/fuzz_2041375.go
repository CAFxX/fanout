package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int16; F1 int16 }
func genFn0(p0 uint8) int16 {
    v6 := GenS1{F0: 1, F1: 732}
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 += 1
        switch -9223372036854775808 {
        case -899:
            v8 := GenS1{F0: 42, F1: 2147483647}
            _ = v8
        }
        var v9 uint64 = 18446744073709551615
        _ = fi7
        _ = v9
    }
    p0 = (p0 / ((p0) | 1))
    switch -1 {
    case -421:
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS2{F0: 0}
            var v12 int64 = -307
            _ = fi10
            _ = v11
            _ = v12
        }
    case 435:
        if (!false) {
            switch -589 {
            case -1:
                fmt.Println(p0)
            }
            p0 += 128
        }
    }
    v13 := GenS3{F0: 1}
    _ = p0
    _ = v13
    _ = v6
    return 32767
}
func genFn1() uint {
    var v14 uint = 18446744073709551615
    fmt.Println(v14)
    v15 := []int{55}
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 uint16 = 292
    var v2 string = "\n\t"
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 1
    var v5 bool = false
    v16 := GenS1{F0: genFn0(127), F1: (v4 << ((4294967295) & 63))}
    v4 = (v4 - v4)
    if v5 {
        v2 = ("foo bar" + v2)
        switch -2147483648 {
        case 32767:
            v4 += 0
        }
        delete(v3, v4)
    }
    switch -592 {
    case -635:
        var v17 int8 = -1
        _ = v17
    case 127:
        if true {
            _ = len(v3)
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS1{F0: -32768, F1: (^(v4))}
                var v20 int64 = 42
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            fmt.Println(v5, v4, v5)
            v21 := GenS2{F0: -1}
            delete(v3, (v4 & (v4 / ((4294967295) | 1))))
            _ = v21
        }
    default:
        switch v2 {
        case "foo bar":
            v5 = (!(-179 >= 32768))
        case "\n\t":
            v5 = v5
        case "hello":
            v2 = (v2 + v2)
        }
    }
    var v22 uint = 0
    switch v2 {
    case "\n\t":
        var v23 float64 = 1e10
        _ = v23
    default:
        switch -9223372036854775808 {
        case 32767:
            if (-233 == 9223372036854775807) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 int64 = 42
                    v25 = v25
                    _ = fi24
                    _ = v25
                }
                if ((182 >= 0) || (!(-54 > 42))) {
                    v5 = (216 != 126)
                    v26 := GenS2{F0: -1}
                    _ = v26
                }
            } else {
                v27 := GenS2{F0: 1000}
                _ = v27
            }
        case 9223372036854775807:
            for fi28 := 0; fi28 < 9; fi28++ {
                if (!v5) {
                    v22 -= genFn1()
                    switch fi28 {
                    case 32768:
                        switch v2 {
                        case "foo bar":
                            v2 = v2
                        default:
                            var v29 bool = (0 <= (v1 | (v1 + 65535)))
                            _ = v29
                        }
                    default:
                        v2 = v2
                    }
                } else {
                    fmt.Println(v22, v4)
                    v22 = 2
                }
                var v30 int16 = -32768
                _ = fi28
                _ = v30
            }
        default:
            if (!v5) {
                for fi31 := 0; fi31 < 20; fi31++ {
                    v22 = v22
                    _ = fi31
                }
                v32 := []int{-748, -9223372036854775808, 96, -1}
                _ = v32
            } else {
                if false {
                    v33 := GenS4{F0: (^(v4)), F1: 65535, F2: -1}
                    switch (v2 + (v2 + "foo bar")) {
                    case "\n\t":
                        v2 = (v2 + "hello")
                    case "foo bar":
                        for fi34 := 0; fi34 < 4; fi34++ {
                            v35 := []int{}
                            _ = fi34
                            _ = v35
                        }
                    }
                    _ = v33
                } else {
                    var v36 int = 4294967295
                    v37 := []int{}
                    _ = v36
                    _ = v37
                }
                v2 = (v2 + "\n\t")
                if v5 {
                    v1 -= 0
                    v4 = (^(v4))
                }
            }
        }
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := []int{fi38}
        delete(v3, v4)
        _ = fi38
        _ = v39
    }
    v2 = ("a" + v2)
    if (42 <= -128) {
        v40 := GenS2{F0: 1000}
        v41 := []int{32767, 9223372036854775807, 2147483647, 32768}
        _ = len(v41)
        _ = v40
        _ = v41
    } else {
        v42 := GenS5{F0: genFn0(42), F1: (v5 || ((0.0 / 1.5) >= (0.1 * -2.25)))}
        v43 := GenS6{F0: 415, F1: 0}
        fmt.Println(v1, v2)
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
