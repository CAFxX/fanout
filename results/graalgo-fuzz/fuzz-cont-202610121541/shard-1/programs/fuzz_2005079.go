package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0() string {
    var v6 uint16 = 65535
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v6)
    case "\n\t":
        v6 = (v6 % ((77) | 1))
    }
    if (false || (42 > 2147483648)) {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (false && true) {
                var v8 string = "hello"
                v8 = (v8 + ((v8 + v8) + ("" + v8)))
                var v9 int = (fi7 | fi7)
                _ = v8
                _ = v9
            }
            var v10 int = 62
            fmt.Println(v10, fi7, v10)
            _ = fi7
            _ = v10
        }
    }
    if true {
        for fi11 := 0; fi11 < 15; fi11++ {
            if (3.14159 > 0.0) {
                var v12 string = ("\n\t" + "a")
                v12 = ""
                _ = v12
            } else {
                v13 := []int{fi11}
                v14 := GenS1{F0: 986, F1: (true && true)}
                switch 829 {
                case -112:
                    var v15 uint8 = 1
                    _ = v15
                case 9223372036854775807:
                    v6 += 2
                }
                _ = v13
                _ = v14
            }
            switch "foo bar" {
            case "foo bar":
                var v16 string = ("" + "foo bar")
                _ = v16
            default:
                if (-866 >= -2147483648) {
                    if true {
                        var v17 int = (fi11 >> ((1) & 63))
                        var v18 uint64 = uint64(v6)
                        _ = v17
                        _ = v18
                    } else {
                        v6 -= (v6 & v6)
                    }
                } else {
                    if (906 <= 18446744073709551615) {
                        for fi19 := 0; fi19 < 10; fi19++ {
                            v20 := []int{-9223372036854775808}
                            if true {
                                var v21 int16 = 0
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                    }
                    v22 := []int{}
                    v23 := GenS1{F0: 6, F1: (int32(v6) == 1)}
                    _ = v22
                    _ = v23
                }
            }
            _ = fi11
        }
        v24 := []int{1}
        _ = len(v24)
        _ = v24
    }
    if true {
        var v25 float32 = 0.1
        v26 := []int{-1, 128}
        var v27 string = "a"
        _ = v25
        _ = v26
        _ = v27
    } else {
        if false {
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := GenS2{F0: -172, F1: 68, F2: 65535}
                v6 += v6
                if true {
                    v30 := GenS3{F0: 74}
                    _ = v30
                }
                _ = fi28
                _ = v29
            }
        }
    }
    _ = v6
    return "a"
}

func main() {
    var v1 uint = 255
    var v2 int32 = -1
    var v3 uint = v1
    v4 := make(map[string]bool)
    v5 := make(map[string]bool)
    v31 := GenS4{F0: 1000}
    var v32 bool = true
    fmt.Println(v3, v32, v32)
    var v33 float64 = float64(-2.25)
    v33 -= v33
    var v34 float64 = v33
    v35 := []int{}
    var v36 int8 = -9
    var v37 int32 = (v2 + 711)
    if v32 {
        fmt.Println(v32, v32)
    }
    v38 := v5["a"]
    var v39 float32 = 0.0
    switch "\n\t" {
    case "\n\t":
        v34 -= v33
    case "":
        v40 := v5["a"]
        _ = v40
    case "a":
        if false {
            for fi41 := 0; fi41 < 6; fi41++ {
                v42 := GenS2{F0: v2, F1: 174, F2: 0}
                switch fi41 {
                case 4294967295:
                    var v43 float32 = 1.5
                    _ = v43
                case -1:
                    switch ("" + "hello") {
                    case "":
                        var v44 uint64 = 219
                        _ = v44
                    default:
                        for fi45 := 0; fi45 < 1; fi45++ {
                            if false {
                                v3 -= (v3 >> ((v1) & 63))
                                switch fi45 {
                                case 0:
                                    var v46 float64 = ((-2.25 * (v33 / float64(v39))) * v34)
                                    _ = v46
                                case -9223372036854775808:
                                    delete(v5, "")
                                case 32767:
                                    v47 := GenS5{F0: 127, F1: 42}
                                    _ = v47
                                default:
                                    v48 := v4["\n\t"]
                                    _ = v48
                                }
                                var v49 int64 = -592
                                var v50 bool = v32
                                fmt.Println(v2)
                                var v51 string = ("hello" + ("\n\t" + "a"))
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                            _ = fi45
                        }
                    }
                case -813:
                    var v52 int = (fi41 / (((fi41 << ((1) & 63))) | 1))
                    _ = v52
                default:
                    for fi53 := 0; fi53 < 2; fi53++ {
                        var v54 float64 = v34
                        _ = fi53
                        _ = v54
                    }
                }
                _ = fi41
                _ = v42
            }
            var v55 int32 = v2
            _ = v55
        }
    default:
        var v56 bool = v38
        _ = v56
    }
    if ((0.1 - float64(v39)) <= v33) {
        var v57 int64 = -890
        _ = v57
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v34)
    default:
        switch 127 {
        case 2147483647:
            for fi58 := 0; fi58 < 11; fi58++ {
                var v59 uint = 255
                var v60 int8 = v36
                _ = fi58
                _ = v59
                _ = v60
            }
        case -2147483648:
            var v61 uint32 = 2147483647
            _ = v61
        case 366:
            var v62 uint64 = 2
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
