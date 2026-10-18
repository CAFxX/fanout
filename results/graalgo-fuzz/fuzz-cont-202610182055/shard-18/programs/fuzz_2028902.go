package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = -97
    var v2 bool = (false && false)
    var v3 uint8 = 2
    v4 := make(map[int]int)
    v5 := []uint32{2147483647, 0}
    var v6 bool = v2
    fmt.Println(v3, v1, v3)
    switch -790 {
    case 127:
        switch 9223372036854775807 {
        case 2147483647:
            delete(v4, 848)
        case 193:
            v7 := GenS1{F0: 0}
            _ = v7
        case -9223372036854775808:
            v5 = append(v5, 65535)
        default:
            delete(v4, -564)
        }
    case 640:
        v8 := GenS1{F0: -9223372036854775808}
        _ = v8
    }
    v3 = v3
    fmt.Println(v2, v2)
    v6 = (690 <= 2)
    switch -881 {
    case 42:
        v9 := v4[-1]
        _ = v9
    case -1:
        switch 779 {
        case 4294967295:
            switch 234 {
            case 0:
                if false {
                    switch ("" + "hello") {
                    case "foo bar":
                        v1 = v1
                    }
                    if (!(true || v2)) {
                        var v10 uint64 = 504
                        _ = len(v4)
                        _ = v10
                    }
                }
            case 42:
                var v11 string = ""
                _ = v11
            default:
                var v12 bool = v6
                _ = v12
            }
        case -1:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int16 = 32767
                var v15 int64 = 9223372036854775807
                v2 = v6
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            fmt.Println(v2)
        }
    case 0:
        for fi16 := 0; fi16 < 3; fi16++ {
            switch ("a" + "a") {
            case "a":
                fmt.Println(fi16)
            }
            _ = fi16
        }
    }
    var v17 uint64 = 0
    v18 := GenS1{F0: 983}
    fmt.Println(v1, v2)
    v19 := GenS2{F0: "", F1: (v17 * (^(v17))), F2: -36}
    if v2 {
        if true {
            v2 = true
        } else {
            v17 = 0
            if (!(!(1.5 < float64(1.5)))) {
                v2 = (v6 || v2)
            } else {
                var v20 uint64 = v17
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch "" {
                    case "foo bar":
                        if len(v5) > 0 {
                            v5[0] = uint32(v3)
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            var v22 int8 = v1
                            _ = v22
                        }
                    case "":
                        if v6 {
                            v3 = v3
                            var v23 uint64 = (v20 | 2)
                            _ = v23
                        }
                    default:
                        if ("" >= ("\n\t" + "a")) {
                            fmt.Println(v17, fi21)
                            switch fi21 {
                            case 0:
                                fmt.Println(v20, v6, fi21)
                            }
                            v20 = 547
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                            switch fi21 {
                            case 381:
                                v6 = (!false)
                            }
                        }
                    }
                    _ = fi21
                }
                switch 42 {
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 18; fi24++ {
                        if false {
                            switch -561 {
                            case 32768:
                                for fi25 := 0; fi25 < 6; fi25++ {
                                    switch "" {
                                    case "":
                                        if ((v1 - v1) == v1) {
                                            switch "" {
                                            case "hello":
                                                v6 = (true && (382 < 295))
                                            }
                                        } else {
                                            var v26 int8 = (v1 + v1)
                                            var v27 string = ""
                                            var v28 int32 = 266
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    case "hello":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case "a":
                                        v2 = v2
                                    default:
                                        var v29 uint8 = (v3 + v3)
                                        _ = v29
                                    }
                                    if false {
                                        v30 := GenS1{F0: 42}
                                        var v31 bool = v6
                                        var v32 int8 = v1
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = fi25
                                }
                            }
                            v33 := GenS3{F0: v3}
                            v2 = false
                            _ = v33
                        }
                        _ = fi24
                    }
                default:
                    var v34 uint8 = 42
                    _ = v34
                }
                _ = v20
            }
            if (-2.25 <= ((3.14159 + -2.25) / 0.1)) {
                v3 = 127
                if len(v5) > 0 {
                    v5[0] = 1
                }
                v4[502] = 32767
            }
        }
        v35 := v4[352]
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 int64 = 75
            v3 = 235
            _ = fi37
            _ = v38
        }
        if (v17 > (v17 >> ((2) & 63))) {
            v2 = v2
        } else {
            if v2 {
                var v39 uint8 = 42
                _ = v39
            }
            var v40 bool = (18446744073709551615 >= 0)
            _ = v40
        }
        _ = fi36
    }
    fmt.Println(v17)
    var v41 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
