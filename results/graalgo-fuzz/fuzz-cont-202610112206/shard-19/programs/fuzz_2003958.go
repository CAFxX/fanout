package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 int, p1 int) int8 {
    for fi4 := 0; fi4 < 9; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            switch ("" + ("hello" + ("foo bar" + ("" + ("foo bar" + ""))))) {
            case "hello":
                var v6 string = ("\n\t" + "a")
                _ = v6
            default:
                var v7 int64 = 42
                _ = v7
            }
            v8 := GenS1{F0: 835, F1: 382}
            v9 := []int{(p1 & p1), 128, (fi4 + (fi5 * (p1 << ((568) & 63)))), (p0 / ((p0) | 1))}
            _ = fi5
            _ = v8
            _ = v9
        }
        _ = fi4
    }
    var v10 float64 = 1e10
    _ = p0
    _ = p1
    _ = v10
    return 127
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 2
    var v3 uint64 = 18446744073709551615
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint16 = 32767
            var v14 uint = 255
            if false {
                var v15 string = "\n\t"
                v16 := []int{}
                if (!true) {
                    switch (fi12 & (fi12 >> ((v14) & 63))) {
                    case 128:
                        v17 := GenS1{F0: (v3 << (((^(v14))) & 63)), F1: uint32(v14)}
                        _ = v17
                    case 114:
                        switch fi12 {
                        case -527:
                            v18 := GenS1{F0: 42, F1: 4294967295}
                            _ = v18
                        case -644:
                            v19 := GenS2{F0: 32767}
                            _ = v19
                        }
                        switch (fi12 * (fi12 + fi11)) {
                        case -2147483648:
                            var v20 uint = 64
                            _ = v20
                        case 32767:
                            switch fi11 {
                            case -9223372036854775808:
                                var v21 uint = v14
                                _ = v21
                            case -1:
                                var v22 uint8 = 141
                                _ = v22
                            }
                            switch (fi11 * (fi11 | fi12)) {
                            case -9223372036854775808:
                                v3 = v3
                            default:
                                var v23 uint = (v14 >> (((^(v14))) & 63))
                                _ = v23
                            }
                        case 0:
                            v14 = (^(v14))
                        default:
                            var v24 float32 = (0.0 * 0.0)
                            _ = v24
                        }
                        if len(v16) > 0 {
                            v16[0] = (fi11 >> ((uint(v2)) & 63))
                        }
                    case 42:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                    if true {
                        if true {
                            v2 += v2
                            switch v15 {
                            case "foo bar":
                                fmt.Println(fi12)
                            case "":
                                fmt.Println(v1)
                            default:
                                v2 += 889
                            }
                            var v25 string = v15
                            _ = v25
                        }
                    }
                    v1 = (v1 >> ((v14) & 63))
                }
                _ = v15
                _ = v16
            } else {
                v26 := GenS3{F0: v3}
                var v27 int = (-(fi12))
                v28 := []int{}
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    v29 := GenS2{F0: 999}
    for fi30 := 0; fi30 < 16; fi30++ {
        v31 := []int{}
        v3 += v3
        var v32 int8 = -1
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        v34 := GenS4{F0: "", F1: 249}
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: 1, F1: 0}
    v2 = (v2 / (((v2 << ((370) & 63))) | 1))
    v1 = genFn0(976, 2147483647)
    fmt.Println(v2, v2, v1)
    v1 += 79
    if (65535 <= 0) {
        v1 = v1
    }
    var v36 uint64 = 4294967295
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 4; fi38++ {
            if ((v1 + v1) == 100) {
                v36 += v3
                v39 := []int{}
                _ = v39
            } else {
                var v40 uint8 = 1
                v41 := GenS2{F0: 42}
                if false {
                    var v42 int32 = -940
                    _ = v42
                } else {
                    v1 += genFn0((fi38 - (fi37 & fi38)), 9223372036854775807)
                    var v43 int8 = -53
                    var v44 int64 = 1
                    _ = v43
                    _ = v44
                }
                _ = v40
                _ = v41
            }
            v2 -= 65535
            _ = fi38
        }
        if ((!true) && (!false)) {
            var v45 int64 = -9223372036854775808
            var v46 float32 = (float32(0.0) * 0.1)
            _ = v45
            _ = v46
        } else {
            v2 = uint32(v36)
        }
        v47 := GenS5{F0: 1, F1: 997}
        _ = fi37
        _ = v47
    }
    var v48 int16 = 42
    v1 -= 100
    var v49 int32 = 18
    var v50 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v48
    _ = v49
    _ = v50
}
