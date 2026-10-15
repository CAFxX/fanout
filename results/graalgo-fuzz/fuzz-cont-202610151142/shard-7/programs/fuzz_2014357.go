package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
func genFn0() int32 {
    var v7 float64 = 0.0
    var v8 int = -2147483648
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -1
    var v3 uint32 = 2
    var v4 string = "a"
    var v5 int64 = -9223372036854775808
    var v6 uint32 = 42
    fmt.Println(v2, v4, v2)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if false {
                fmt.Println(v3)
                var v11 int64 = -675
                v12 := GenS1{F0: 206, F1: (fi10 / (((fi9 | (fi10 ^ (fi9 << ((18446744073709551615) & 63))))) | 1))}
                _ = v11
                _ = v12
            }
            var v13 uint64 = 42
            _ = fi10
            _ = v13
        }
        var v14 uint8 = 42
        for fi15 := 0; fi15 < 1; fi15++ {
            if ((128 <= v14) || false) {
                v16 := GenS1{F0: v14, F1: (fi9 & fi9)}
                _ = v16
            }
            if (v2 <= (v2 & (v2 << ((18446744073709551615) & 63)))) {
                v17 := GenS1{F0: v14, F1: -987}
                _ = v17
            } else {
                if false {
                    var v18 int64 = 0
                    var v19 bool = (1 >= 2)
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS2{F0: (!(0 <= 1000))}
                    var v21 uint64 = 4294967295
                    _ = v20
                    _ = v21
                }
                v6 -= (^(v3))
            }
            var v22 int8 = (v2 + 42)
            _ = fi15
            _ = v22
        }
        _ = fi9
        _ = v14
    }
    v5 = v5
    var v23 int32 = 0
    if (!(42 <= 0)) {
        switch "" {
        case "foo bar":
            for fi24 := 0; fi24 < 18; fi24++ {
                if (((0.0 * -2.25) / 0.1) != 0.1) {
                    v6 = (^(v3))
                }
                var v25 int = 86
                _ = fi24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(v2, v4, fi26)
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 string = "\n\t"
                v5 = 9223372036854775807
                v28 = v28
                v29 := GenS1{F0: 146, F1: (fi26 >> ((2) & 63))}
                v30 := GenS1{F0: 42, F1: -9223372036854775808}
                v23 = 0
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            var v31 string = ("hello" + "foo bar")
            _ = fi26
            _ = v31
        }
        delete(v1, -915)
    }
    var v32 uint8 = 255
    switch -9223372036854775808 {
    case -759:
        for fi33 := 0; fi33 < 19; fi33++ {
            _ = len(v1)
            v3 = 65535
            fmt.Println(v2)
            _ = fi33
        }
    default:
        v34 := GenS2{F0: true}
        _ = v34
    }
    fmt.Println(v3, v4, v6)
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := []int{(fi35 + fi35), fi35}
        _ = fi35
        _ = v36
    }
    v5 -= v5
    v37 := GenS3{F0: (v2 >> ((4294967295) & 63))}
    v38 := []int{1}
    v32 += v32
    switch -204 {
    case 680:
        if true {
            var v39 uint = 0
            fmt.Println(v5, v23)
            _ = v39
        }
    case -611:
        for fi40 := 0; fi40 < 18; fi40++ {
            v41 := GenS1{F0: v32, F1: fi40}
            var v42 int = (^(fi40))
            _ = fi40
            _ = v41
            _ = v42
        }
    default:
        v43 := GenS1{F0: v32, F1: 32768}
        _ = v43
    }
    v44 := GenS3{F0: (v2 + (v2 + 1))}
    switch 32768 {
    case -752:
        v45 := GenS2{F0: (v32 == v32)}
        _ = v45
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        for fi47 := 0; fi47 < 16; fi47++ {
            v32 -= 125
            var v48 float64 = -2.25
            _ = fi47
            _ = v48
        }
        var v49 uint16 = 544
        fmt.Println(v49, v6, v3)
        _ = fi46
        _ = v49
    }
    v6 = (v3 / (((v6 + (v3 << ((2) & 63)))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
