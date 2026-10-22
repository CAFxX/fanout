package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int32, p1 string, p2 uint) uint64 {
    v6 := GenS1{F0: 445, F1: "foo bar", F2: (2 > (^(p2)))}
    if false {
        if (42 > -32768) {
            var v7 uint16 = 318
            switch -2147483648 {
            case 540:
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int8 = -1
                    _ = fi8
                    _ = v9
                }
            case 42:
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := GenS2{F0: (127 != -110)}
                    _ = fi10
                    _ = v11
                }
            case -931:
                var v12 uint16 = 878
                _ = v12
            }
            _ = v7
        }
        var v13 float32 = -2.25
        if (161 >= 263) {
            var v14 int64 = -1
            fmt.Println(v13)
            _ = v14
        } else {
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v13 = 1.5
                    fmt.Println(p1)
                    v16 := []int{(fi15 << ((542) & 63)), (fi15 << ((p2) & 63))}
                    _ = fi15
                    _ = v16
                }
            } else {
                switch p1 {
                case "hello":
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v18 := []int{fi17}
                        if len(v18) > 0 {
                            v18[0] = (fi17 % ((-791) | 1))
                        }
                        p2 += p2
                        _ = fi17
                        _ = v18
                    }
                }
            }
            var v19 int32 = -1
            _ = v19
        }
        _ = v13
    } else {
        var v20 float64 = (float64(1.5) / 0.1)
        var v21 uint16 = 574
        _ = v20
        _ = v21
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        var v22 uint32 = 239
        _ = v22
    case "a":
        var v23 bool = true
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint32 = 463
    var v4 int = 4294967295
    var v5 int8 = 100
    var v24 uint32 = 1
    v1 = false
    var v25 int64 = 9223372036854775807
    v5 += (v5 << ((18446744073709551615) & 63))
    switch -2147483648 {
    case 509:
        switch -643 {
        case 127:
            var v26 int16 = -1
            _ = v26
        }
    default:
        v1 = v1
    }
    v27 := []int{}
    if v1 {
        switch v4 {
        case 32767:
            var v28 uint16 = 32767
            _ = v28
        default:
            for fi29 := 0; fi29 < 16; fi29++ {
                v1 = (true || (42 == 168))
                for fi30 := 0; fi30 < 12; fi30++ {
                    v27 = append(v27, (fi29 >> ((1) & 63)))
                    var v31 uint64 = 0
                    _ = fi30
                    _ = v31
                }
                v32 := GenS3{F0: 0, F1: uint16(v24), F2: int(v5)}
                _ = fi29
                _ = v32
            }
        }
        var v33 uint64 = 42
        var v34 int64 = (v25 + (^(v25)))
        _ = v33
        _ = v34
    }
    if (18446744073709551615 == 624) {
        for fi35 := 0; fi35 < 7; fi35++ {
            if v1 {
                var v36 string = "hello"
                _ = v36
            } else {
                v37 := GenS4{F0: v1}
                _ = v37
            }
            if false {
                v24 -= v3
                _ = len(v27)
                for fi38 := 0; fi38 < 2; fi38++ {
                    switch fi38 {
                    case -173:
                        var v39 uint16 = 180
                        _ = v39
                    case 612:
                        if v1 {
                            v40 := GenS1{F0: -303, F1: "foo bar", F2: (42 > 947)}
                            v2[""] = (^(v4))
                            _ = v40
                        }
                    case 9223372036854775807:
                        fmt.Println(v3, v25)
                    default:
                        v24 -= v24
                    }
                    _ = fi38
                }
            } else {
                v41 := GenS5{F0: v25}
                _ = v41
            }
            _ = fi35
        }
        if false {
            var v42 uint = 1
            switch "hello" {
            case "foo bar":
                v43 := GenS6{F0: (v4 >> ((v42) & 63))}
                _ = v43
            default:
                v4 = (v4 << ((v42) & 63))
            }
            _ = v42
        }
        var v44 int8 = 0
        _ = v44
    } else {
        _ = len(v27)
    }
    var v45 uint64 = 145
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
