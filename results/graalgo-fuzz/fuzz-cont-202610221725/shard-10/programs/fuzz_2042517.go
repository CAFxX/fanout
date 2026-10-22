package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int16; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 int8; F2 bool }
type GenS8 struct { F0 uint16; F1 int16; F2 uint }
type GenS9 struct { F0 uint16 }
func genFn0() uint16 {
    var v7 uint = 4294967295
    var v8 int32 = 1
    v8 += v8
    _ = v7
    _ = v8
    return 394
}

func main() {
    var v1 bool = false
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 bool = (("\n\t" + (v2 + v2)) != "foo bar")
    var v6 uint16 = 1
    switch "" {
    case "":
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: fi9}
            switch v2 {
            case "hello":
                var v11 float32 = 1.5
                _ = v11
            case "a":
                fmt.Println(fi9, v1)
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int32 = -2147483648
                var v14 int8 = 0
                if (!false) {
                    fmt.Println(v2, fi9, fi9)
                } else {
                    fmt.Println(v5)
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
    case "\n\t":
        v2 = "\n\t"
    case "hello":
        var v15 uint = 65535
        _ = v15
    }
    v5 = (!true)
    switch (v2 + "") {
    case "\n\t":
        var v16 uint8 = 127
        _ = v16
    case "foo bar":
        fmt.Println(v1, v1, v5)
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        if v1 {
            if (!(true && true)) {
                switch fi17 {
                case 4294967295:
                    if true {
                        switch (fi17 * fi17) {
                        case 42:
                            fmt.Println(fi17)
                        case -1:
                            var v18 int32 = -2147483648
                            _ = v18
                        }
                    }
                case 2147483647:
                    if (1 >= 255) {
                        var v19 int32 = 2147483647
                        _ = v19
                    }
                case 32767:
                    if (11 < 0) {
                        fmt.Println(fi17, v5, v2)
                    }
                default:
                    fmt.Println(v5, v5)
                }
            } else {
                v2 = (v2 + (("hello" + ("a" + v2)) + v2))
            }
            v20 := GenS1{F0: (fi17 ^ (fi17 / ((9223372036854775807) | 1)))}
            switch v2 {
            case "":
                v21 := GenS1{F0: (fi17 + (fi17 >> ((56) & 63)))}
                _ = v21
            }
            _ = v20
        } else {
            if v5 {
                for fi22 := 0; fi22 < 6; fi22++ {
                    var v23 float32 = 1.5
                    v4[(v2 + "a")] = fi22
                    _ = fi22
                    _ = v23
                }
                v24 := []int{fi17}
                if v1 {
                    var v25 uint16 = (v6 ^ v6)
                    _ = v25
                }
                _ = v24
            } else {
                switch fi17 {
                case -1:
                    var v26 uint32 = 4294967295
                    _ = v26
                case 74:
                    fmt.Println(v6, v6)
                case 690:
                    if false {
                        var v27 int = fi17
                        v6 += 32767
                        v6 += (v6 & v6)
                        _ = v27
                    } else {
                        v28 := GenS2{F0: 9223372036854775807, F1: 784}
                        v29 := GenS3{F0: 32767, F1: 2147483647, F2: 1}
                        v5 = ((((v2 + v2) + "\n\t") != (v2 + "hello")) && (0 <= uint32(v6)))
                        _ = v28
                        _ = v29
                    }
                }
                var v30 uint16 = (v6 / (((v6 / ((32767) | 1))) | 1))
                var v31 int8 = 127
                _ = v30
                _ = v31
            }
            v32 := GenS3{F0: v6, F1: 0, F2: -2147483648}
            v33 := []int{fi17, 127, (fi17 + (fi17 | fi17)), fi17}
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 int = (fi17 & (fi17 + (fi17 & fi17)))
            v5 = true
            _ = fi34
            _ = v35
        }
        _ = fi17
    }
    switch 127 {
    case -9223372036854775808:
        var v36 uint64 = 1
        _ = v36
    case 128:
        var v37 bool = false
        _ = v37
    case -1:
        var v38 float64 = (0.0 * (((0.0 - (0.1 / -2.25)) - (-2.25 - 3.14159)) + -2.25))
        _ = v38
    default:
        var v39 int32 = 2147483647
        _ = v39
    }
    switch v2 {
    case "\n\t":
        switch 127 {
        case 0:
            var v40 float32 = ((1e10 - 0.0) / 1.5)
            _ = v40
        case -916:
            v41 := GenS4{F0: 0}
            _ = v41
        }
    case "":
        v4["hello"] = 128
    }
    switch v2 {
    case "hello":
        if true {
            if (!(42 > 42)) {
                switch -277 {
                case 1:
                    var v42 float64 = -2.25
                    _ = v42
                }
                v43 := GenS1{F0: 572}
                _ = v43
            } else {
                switch 128 {
                case -337:
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v44 := GenS5{F0: 65535, F1: 1000, F2: v1}
                        _ = v44
                    case "a":
                        for fi45 := 0; fi45 < 18; fi45++ {
                            for fi46 := 0; fi46 < 15; fi46++ {
                                v47 := []int{(fi46 / ((fi45) | 1)), fi45, fi46}
                                v48 := GenS6{F0: 59}
                                v49 := GenS1{F0: -2147483648}
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            switch (^(fi45)) {
                            case 4294967295:
                                switch (^(fi45)) {
                                case 309:
                                    if (-969 <= 42) {
                                        var v50 float64 = 3.14159
                                        _ = v50
                                    } else {
                                        if false {
                                            v2 = "foo bar"
                                        } else {
                                            var v51 int16 = -1
                                            _ = v51
                                        }
                                        switch (fi45 >> ((18446744073709551615) & 63)) {
                                        case 128:
                                            switch v2 {
                                            case "a":
                                                if false {
                                                    for fi52 := 0; fi52 < 10; fi52++ {
                                                        v5 = v5
                                                        _ = fi52
                                                    }
                                                }
                                            case "":
                                                fmt.Println(v6, fi45, fi45)
                                            case "\n\t":
                                                v53 := GenS7{F0: 1, F1: 100, F2: (42 > 42)}
                                                _ = v53
                                            default:
                                                var v54 uint8 = 0
                                                _ = v54
                                            }
                                        }
                                        var v55 int8 = 127
                                        _ = v55
                                    }
                                case 42:
                                    delete(v3, 4294967295)
                                case 127:
                                    v2 = v2
                                default:
                                    v56 := GenS7{F0: 18, F1: 1, F2: v1}
                                    _ = v56
                                }
                            case -9223372036854775808:
                                v57 := GenS8{F0: (v6 + (v6 >> ((255) & 63))), F1: 32767, F2: 269}
                                _ = v57
                            case 128:
                                v1 = (!false)
                            default:
                                var v58 int64 = -156
                                _ = v58
                            }
                            _ = fi45
                        }
                    case "\n\t":
                        fmt.Println(v5, v2)
                    }
                }
            }
            delete(v4, v2)
        } else {
            switch ("\n\t" + ("a" + v2)) {
            case "a":
                _ = len(v4)
            case "foo bar":
                v4[(("foo bar" + "\n\t") + ("foo bar" + (v2 + "a")))] = -9223372036854775808
            }
        }
    case "":
        v59 := GenS9{F0: (v6 - v6)}
        _ = v59
    default:
        if (-954 > 2147483647) {
            var v60 uint64 = 2
            _ = v60
        }
    }
    v61 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
