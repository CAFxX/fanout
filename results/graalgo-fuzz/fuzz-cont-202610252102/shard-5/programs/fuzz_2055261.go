package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 float64 = (3.14159 - 1e10)
    var v3 bool = ((("" + "\n\t") + "foo bar") <= "\n\t")
    v4 := []uint32{127}
    var v5 uint = 42
    v6 := make(map[int]string)
    fmt.Println(v3, v3)
    var v7 string = ("foo bar" + "a")
    v8 := GenS1{F0: 42, F1: 255, F2: v5}
    _ = len(v6)
    v7 = ""
    if v3 {
        switch "" {
        case "":
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 10; fi10++ {
                    v11 := GenS2{F0: 0, F1: -9223372036854775808}
                    var v12 int64 = 621
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                fmt.Println(fi9, fi9, v1)
                _ = fi9
            }
        case "foo bar":
            switch 32767 {
            case 957:
                v1 += (v1 * 2)
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        case "hello":
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int8 = 127
                _ = fi13
                _ = v14
            }
        default:
            switch ("hello" + ("a" + "\n\t")) {
            case "foo bar":
                var v15 float64 = 1.5
                _ = v15
            case "a":
                var v16 int64 = 1
                _ = v16
            case "hello":
                v5 = (v5 + 1)
            }
        }
        v17 := v6[-1]
        v4 = append(v4, v1)
        _ = v17
    } else {
        for fi18 := 0; fi18 < 3; fi18++ {
            _ = len(v4)
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 uint = (v5 + v5)
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        var v21 int32 = -1
        _ = v21
    }
    v3 = true
    switch -568 {
    case -2147483648:
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v4) > 0 {
                v4[0] = v1
            }
            _ = fi22
        }
    case 42:
        switch 127 {
        case 127:
            if v3 {
                var v23 uint16 = 298
                v3 = (581 <= -669)
                switch 63 {
                case -866:
                    for fi24 := 0; fi24 < 15; fi24++ {
                        v7 = v7
                        v7 = v7
                        _ = fi24
                    }
                case 887:
                    v2 += v2
                case 128:
                    v25 := GenS1{F0: 42, F1: 249, F2: 42}
                    _ = v25
                default:
                    v3 = (0 >= 42)
                }
                _ = v23
            }
        case 9223372036854775807:
            if (!v3) {
                var v26 int32 = 2147483647
                for fi27 := 0; fi27 < 20; fi27++ {
                    fmt.Println(v1, v1, v2)
                    v3 = v3
                    _ = fi27
                }
                if v3 {
                    v7 = "a"
                    v28 := GenS3{F0: 207, F1: 962, F2: (v26 ^ v26)}
                    _ = v28
                } else {
                    if true {
                        if (128 < 128) {
                            if true {
                                v2 = 3.14159
                                v29 := GenS1{F0: (^(v26)), F1: 127, F2: (v5 * 1)}
                                _ = v29
                            } else {
                                var v30 int32 = -1
                                v31 := GenS3{F0: 55, F1: 0, F2: -102}
                                v32 := GenS4{F0: 357, F1: (-(v30)), F2: false}
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            v1 -= v1
                            if v3 {
                                if v3 {
                                    v4 = append(v4, v1)
                                    if false {
                                        var v33 int64 = 1
                                        fmt.Println(v5, v33)
                                        _ = v33
                                    }
                                } else {
                                    var v34 string = "a"
                                    _ = v34
                                }
                                _ = len(v4)
                                if false {
                                    v3 = (!v3)
                                    v35 := GenS5{F0: -1}
                                    _ = v35
                                }
                            }
                        }
                        v36 := GenS3{F0: 128, F1: 42, F2: -712}
                        _ = v36
                    }
                }
                _ = v26
            } else {
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 int32 = 1
                    _ = fi37
                    _ = v38
                }
                if (-128 > 42) {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        if v3 {
                            fmt.Println(v1)
                            fmt.Println(v1, fi39)
                        } else {
                            var v40 int32 = -2147483648
                            _ = v40
                        }
                        _ = fi39
                    }
                    _ = len(v6)
                } else {
                    delete(v6, -2147483648)
                    fmt.Println(v3, v3, v2)
                }
            }
        case 560:
            if false {
                switch 72 {
                case -400:
                    v41 := GenS6{F0: v5, F1: 42}
                    _ = v41
                case 128:
                    fmt.Println(v1, v7)
                case 278:
                    v4 = append(v4, (v1 * (v1 & 1)))
                }
            } else {
                _ = len(v6)
            }
        default:
            v42 := GenS7{F0: 42}
            _ = v42
        }
    case 154:
        if ((int32(v1) > 429) || (4294967295 > 42)) {
            for fi43 := 0; fi43 < 1; fi43++ {
                v5 = 2
                var v44 int16 = -32768
                for fi45 := 0; fi45 < 13; fi45++ {
                    fmt.Println(v2, v7)
                    if ((fi45 & fi43) <= (fi43 << ((v5) & 63))) {
                        fmt.Println(v44, fi43, fi45)
                        var v46 string = (("a" + v7) + v7)
                        v44 += v44
                        _ = v46
                    } else {
                        if false {
                            switch "" {
                            case "\n\t":
                                v4 = append(v4, uint32(v5))
                            case "a":
                                var v47 float32 = float32((v2 * v2))
                                _ = v47
                            case "hello":
                                var v48 bool = (!((50 >= -2) || false))
                                _ = v48
                            }
                        } else {
                            if v3 {
                                v44 = (v44 ^ (v44 << ((18446744073709551615) & 63)))
                            } else {
                                if (42 < -1) {
                                    v7 = v7
                                } else {
                                    var v49 uint16 = 1
                                    fmt.Println(v3, v5)
                                    v7 = "a"
                                    _ = v49
                                }
                            }
                            v2 = v2
                            v1 -= 178
                            v44 += -886
                            if (v3 && ((v44 >= (^(v44))) && true)) {
                                var v50 uint32 = (^(v1))
                                _ = v50
                            } else {
                                var v51 int16 = v44
                                v52 := GenS7{F0: 18446744073709551615}
                                if false {
                                    if v3 {
                                        var v53 int8 = 127
                                        fmt.Println(v51, v44)
                                        _ = len(v4)
                                        var v54 int32 = 1
                                        _ = v53
                                        _ = v54
                                    }
                                    v44 += (-(v51))
                                    fmt.Println(v7, v5, fi43)
                                }
                                _ = v51
                                _ = v52
                            }
                            v55 := GenS4{F0: 255, F1: 1, F2: (v44 > (v44 + v44))}
                            _ = v55
                        }
                    }
                    _ = fi45
                }
                _ = fi43
                _ = v44
            }
        }
    }
    if v3 {
        var v56 uint = (v5 & v5)
        _ = v56
    }
    fmt.Println(v2)
    var v57 uint = (v5 + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
    _ = v8
}
