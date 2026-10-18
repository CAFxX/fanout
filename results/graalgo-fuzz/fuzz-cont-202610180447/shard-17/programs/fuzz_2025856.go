package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    if true {
        v2 = (((float32(3.14159) * v2) + v2) * 1.5)
        for fi4 := 0; fi4 < 19; fi4++ {
            v3[(fi4 ^ (fi4 - 0))] = "a"
            var v5 uint32 = 4294967295
            _ = fi4
            _ = v5
        }
    } else {
        for fi6 := 0; fi6 < 8; fi6++ {
            v3[fi6] = ""
            v1 += 424
            for fi7 := 0; fi7 < 3; fi7++ {
                v2 = v2
                _ = fi7
            }
            _ = fi6
        }
        v1 -= (v1 * (v1 * (v1 & (v1 ^ v1))))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v1, v1)
        if false {
            fmt.Println(v1)
        } else {
            v1 = 9
            v2 -= v2
        }
        if ("" <= ("a" + ("hello" + "hello"))) {
            v1 += v1
            v9 := GenS1{F0: 109, F1: 42}
            _ = v9
        } else {
            var v10 int64 = -9223372036854775808
            v11 := []int{}
            v1 += (v1 - (-(v1)))
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v12 := []int{-2147483648, -603, 215, 32768}
    v2 += -2.25
    v1 = v1
    v13 := GenS2{F0: 2147483648, F1: 1}
    if false {
        v14 := GenS3{F0: 42, F1: 65535}
        switch 32768 {
        case 0:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        case -310:
            if true {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v2 -= 1.5
                    for fi16 := 0; fi16 < 12; fi16++ {
                        switch "" {
                        case "\n\t":
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        }
                        if ((("foo bar" + "\n\t") + "") >= ("" + "")) {
                            delete(v3, (fi16 ^ 32767))
                        } else {
                            v17 := GenS1{F0: 389, F1: -9223372036854775808}
                            v1 -= v1
                            _ = v17
                        }
                        v2 += v2
                        _ = fi16
                    }
                    v1 += (v1 ^ v1)
                    _ = fi15
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 int = (^(fi18))
                    _ = fi18
                    _ = v19
                }
            }
        case -174:
            v20 := GenS2{F0: 2147483647, F1: 117}
            _ = v20
        }
        var v21 string = "\n\t"
        _ = v14
        _ = v21
    }
    if true {
        var v22 uint64 = 2
        var v23 uint64 = 2
        var v24 uint64 = 0
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println(v2, v1)
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int64 = -1
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 float32 = ((float32(1.5) + ((v2 * (v2 * 0.1)) + float32(0.1))) / 3.14159)
            v29 := GenS4{F0: 194, F1: 42}
            var v30 bool = true
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 int = fi25
        _ = fi25
        _ = v26
        _ = v31
    }
    if (0 >= 42) {
        switch "a" {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                v2 += (0.0 - 3.14159)
            }
        case "":
            v1 -= 1
        case "hello":
            var v32 uint = 4294967295
            _ = v32
        default:
            v33 := GenS3{F0: 65535, F1: 817}
            _ = v33
        }
        var v34 uint32 = 351
        _ = v34
    } else {
        switch -2147483648 {
        case -1:
            if len(v12) > 0 {
                v12[0] = 1
            }
        case 32768:
            var v35 uint64 = 604
            _ = v35
        case -484:
            switch ("a" + ("hello" + "hello")) {
            case "hello":
                var v36 int16 = -74
                _ = v36
            case "":
                fmt.Println(v2, v2)
            case "foo bar":
                if len(v12) > 0 {
                    v12[0] = 4294967295
                }
            }
        default:
            _ = len(v3)
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            switch "foo bar" {
            case "":
                if ((fi37 ^ (fi37 + (-(fi37)))) > fi37) {
                    var v38 int64 = -1
                    _ = v38
                }
            case "\n\t":
                for fi39 := 0; fi39 < 14; fi39++ {
                    fmt.Println(v1, fi37, v1)
                    var v40 string = ""
                    _ = fi39
                    _ = v40
                }
            }
            var v41 uint8 = 128
            _ = fi37
            _ = v41
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
