package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool }

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{115, 1, 253, 4294967295}
    v3 := []int{9223372036854775807}
    var v4 int16 = 1000
    var v5 string = (("" + ("\n\t" + "\n\t")) + "a")
    if true {
        if (2147483647 != 226) {
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := GenS1{F0: v5, F1: 641, F2: 232}
                fmt.Println(fi6)
                _ = fi6
                _ = v7
            }
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                v4 -= 42
                _ = fi8
            }
            v4 = v4
            var v9 bool = true
            _ = v9
        }
    }
    var v10 int8 = 127
    v11 := GenS1{F0: "\n\t", F1: 4294967295, F2: 1}
    for fi12 := 0; fi12 < 15; fi12++ {
        v10 = (v10 / ((v10) | 1))
        var v13 int8 = v10
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: true}
    if (0.0 != 0.0) {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint16 = 134
            fmt.Println(v16, fi15, v4)
            v17 := GenS1{F0: v5, F1: 0, F2: 18446744073709551615}
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        if true {
            v1 = v1
        } else {
            if (2 < 710) {
                if false {
                    var v18 uint = 255
                    if true {
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 bool = false
                            _ = fi19
                            _ = v20
                        }
                        var v21 int16 = -32768
                        _ = v21
                    }
                    _ = v18
                } else {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v5 = "foo bar"
                        _ = fi22
                    }
                }
                var v23 float32 = -2.25
                _ = v23
            } else {
                v24 := GenS1{F0: v5, F1: 2147483647, F2: 18446744073709551615}
                for fi25 := 0; fi25 < 4; fi25++ {
                    fmt.Println(v1)
                    var v26 int32 = -204
                    _ = fi25
                    _ = v26
                }
                _ = v24
            }
        }
        fmt.Println(v1, v5)
    }
    fmt.Println(v1, v10, v1)
    if true {
        switch "\n\t" {
        case "a":
            if false {
                v27 := GenS3{F0: 1, F1: false}
                for fi28 := 0; fi28 < 11; fi28++ {
                    fmt.Println(v5, v10, fi28)
                    if (false && (v5 < (("\n\t" + v5) + v5))) {
                        for fi29 := 0; fi29 < 7; fi29++ {
                            switch fi29 {
                            case -852:
                                var v30 int8 = (-(v10))
                                _ = v30
                            case 32767:
                                switch fi29 {
                                case 971:
                                    fmt.Println(v1, v4, fi29)
                                }
                            default:
                                v5 = ((((v5 + v5) + ("foo bar" + v5)) + "foo bar") + ((v5 + v5) + ("" + (v5 + v5))))
                            }
                            _ = fi29
                        }
                    }
                    _ = fi28
                }
                _ = v27
            } else {
                if (4294967295 >= 869) {
                    var v31 uint = 42
                    v10 = 1
                    v31 = (v31 >> ((247) & 63))
                    _ = v31
                }
                _ = len(v2)
            }
        case "\n\t":
            v5 = v5
        }
    } else {
        var v32 uint = 531
        if (((^(v32)) > 1) && (true || true)) {
            for fi33 := 0; fi33 < 7; fi33++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi33
            }
            v4 = 1
            v3 = append(v3, 374)
        } else {
            v1 = 0.1
        }
        var v34 string = v5
        _ = v32
        _ = v34
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS1{F0: v5, F1: 630, F2: 2}
        v4 += v4
        for fi37 := 0; fi37 < 15; fi37++ {
            v1 = 0.0
            fmt.Println(fi37, v5, fi37)
            v38 := GenS3{F0: 42, F1: true}
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
