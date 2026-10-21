package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 bool; F2 uint64 }
func genFn0(p0 uint32) uint32 {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            p0 = p0
            var v7 uint32 = (p0 << ((1) & 63))
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            v10 := GenS1{F0: 18446744073709551615, F1: 0}
            v11 := []int{(^(fi6)), -9223372036854775808}
            _ = v10
            _ = v11
        }
        p0 += p0
        _ = fi6
    }
    _ = p0
    return (p0 << ((18446744073709551615) & 63))
}
func genFn1(p0 uint8, p1 uint32) uint64 {
    if (false || (!false)) {
        if true {
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS1{F0: 0, F1: 42}
                _ = fi12
                _ = v13
            }
            v14 := []int{260, 128, 42, -2147483648}
            _ = v14
        } else {
            var v15 float64 = 1.5
            p1 = p1
            if true {
                v16 := GenS1{F0: 42, F1: 625}
                var v17 string = ("hello" + (("\n\t" + "hello") + "foo bar"))
                _ = v16
                _ = v17
            }
            _ = v15
        }
        var v18 bool = ((false && (64 >= 127)) || (0 > 2))
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS2{F0: -2147483648, F1: 0}
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        if true {
            v21 := GenS3{F0: 2, F1: true, F2: 0}
            _ = v21
        } else {
            fmt.Println(p1, p0)
            p0 = p0
        }
        var v22 uint32 = 2147483648
        _ = v22
    }
    var v23 int32 = 2147483647
    v24 := []int{}
    switch int(p1) {
    case 2147483647:
        for fi25 := 0; fi25 < 17; fi25++ {
            if (-128 < 34) {
                if ((1e10 == 3.14159) || false) {
                    for fi26 := 0; fi26 < 7; fi26++ {
                        p1 -= 4294967295
                        var v27 bool = false
                        v28 := GenS3{F0: 18446744073709551615, F1: v27, F2: 42}
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                } else {
                    fmt.Println(p1, v23, fi25)
                    var v29 int16 = -32768
                    _ = v29
                }
            }
            _ = len(v24)
            var v30 string = ("\n\t" + "a")
            _ = fi25
            _ = v30
        }
    case 803:
        p1 += 1
    case 0:
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
    default:
        var v31 int8 = -1
        _ = v31
    }
    var v32 int64 = 797
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v32
    return 0
}

func main() {
    var v1 uint = 757
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 uint64 = 42
    var v5 uint16 = 553
    v33 := GenS1{F0: uint64(v5), F1: (v3 - (-(v3)))}
    v34 := []int{42}
    v35 := v2[("hello" + "a")]
    fmt.Println(v35)
    var v36 int64 = 0
    fmt.Println(v1, v5, v35)
    if v35 {
        var v37 int = 32768
        var v38 int32 = 2147483647
        _ = v37
        _ = v38
    } else {
        v3 = -538
    }
    var v39 uint64 = (v4 << (((^(v1))) & 63))
    for fi40 := 0; fi40 < 2; fi40++ {
        v36 = int64(fi40)
        switch ("hello" + "hello") {
        case "":
            fmt.Println(v5, v35, v36)
        case "\n\t":
            var v41 int32 = v3
            _ = v41
        case "foo bar":
            v42 := GenS2{F0: v3, F1: -1}
            _ = v42
        default:
            var v43 float32 = float32(0.1)
            _ = v43
        }
        _ = fi40
    }
    v44 := GenS1{F0: (v4 | (v39 * v4)), F1: v3}
    v45 := GenS3{F0: v1, F1: true, F2: (v4 / ((v39) | 1))}
    v1 = (v1 ^ (v1 & 0))
    var v46 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
