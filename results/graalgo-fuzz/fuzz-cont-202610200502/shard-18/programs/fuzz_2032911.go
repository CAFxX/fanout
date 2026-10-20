package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 uint, p2 int64) uint8 {
    switch (p0 ^ 979) {
    case 128:
        v7 := []int{}
        _ = v7
    case -9223372036854775808:
        var v8 int64 = (^(p2))
        _ = v8
    default:
        p2 -= (p2 * p2)
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 1000, F1: -2147483648}
        switch "foo bar" {
        case "":
            v11 := GenS1{F0: 1000, F1: -102}
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    if (true || false) {
        var v12 uint32 = 575
        _ = v12
    }
    fmt.Println(p2)
    var v13 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2
}
func genFn1(p0 bool, p1 int8) uint32 {
    fmt.Println(p0, p1)
    switch 347 {
    case 32768:
        var v14 float64 = -2.25
        _ = v14
    case -2147483648:
        p1 -= (p1 / ((100) | 1))
    default:
        var v15 bool = p0
        _ = v15
    }
    _ = p0
    _ = p1
    return 331
}

func main() {
    var v1 bool = (!true)
    var v2 int32 = -1
    v3 := []int{-9223372036854775808, 4294967295, 4294967295}
    var v4 bool = (!v1)
    var v5 int32 = v2
    var v6 uint64 = 0
    var v16 bool = (false && (!false))
    var v17 uint64 = v6
    var v18 uint = 2
    var v19 int8 = 0
    var v20 bool = v16
    v17 += (v17 >> (((v18 + (v18 - v18))) & 63))
    fmt.Println(v6, v17)
    fmt.Println(v1)
    switch 789 {
    case -2147483648:
        if (!(true && (v20 && (90 != 32767)))) {
            var v21 string = "a"
            _ = v21
        }
    case 2147483647:
        if len(v3) > 0 {
            v3[0] = 32768
        }
    case 4294967295:
        switch "hello" {
        case "foo bar":
            if v1 {
                var v22 int32 = -844
                var v23 uint64 = (^(v17))
                for fi24 := 0; fi24 < 15; fi24++ {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v23 += (v6 - v23)
                        _ = fi25
                    }
                    fmt.Println(v20)
                    v26 := GenS1{F0: 1000, F1: -1}
                    _ = fi24
                    _ = v26
                }
                _ = v22
                _ = v23
            }
        case "\n\t":
            if true {
                v18 = 42
                if v16 {
                    var v27 int32 = (v5 % (((v2 + v2)) | 1))
                    var v28 string = "a"
                    _ = v27
                    _ = v28
                }
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    fmt.Println(v5, v16)
                    _ = fi29
                }
                switch "\n\t" {
                case "\n\t":
                    switch 1 {
                    case 127:
                        if false {
                            fmt.Println(v5, v4, v4)
                            switch "foo bar" {
                            case "\n\t":
                                switch -263 {
                                case -965:
                                    var v30 int32 = v2
                                    _ = v30
                                case 127:
                                    v31 := GenS1{F0: -1, F1: v5}
                                    _ = v31
                                case 128:
                                    if v16 {
                                        var v32 int32 = int32(v6)
                                        _ = v32
                                    }
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = 32767
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v2, v19)
                            case "":
                                var v33 int64 = 42
                                _ = v33
                            default:
                                if (0 != int(v5)) {
                                    v5 = (v5 - v5)
                                    var v34 uint16 = 65535
                                    _ = v34
                                } else {
                                    switch 32767 {
                                    case 32768:
                                        if false {
                                            if v4 {
                                                v18 += (v18 ^ v18)
                                                v35 := GenS1{F0: 195, F1: v2}
                                                _ = v35
                                            } else {
                                                fmt.Println(v16)
                                                if len(v3) > 0 {
                                                    v3[0] = -143
                                                }
                                            }
                                            if v16 {
                                                v6 += v17
                                                v6 = 0
                                                if v16 {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            v1 = (uint64(v18) < (v17 - (v6 % ((42) | 1))))
                                        } else {
                                            _ = len(v3)
                                            for fi36 := 0; fi36 < 9; fi36++ {
                                                v4 = (((("\n\t" + (("" + "a") + (("" + "foo bar") + "a"))) + "\n\t") + "a") == ("hello" + ""))
                                                v5 += v5
                                                fmt.Println(v17, v17, v20)
                                                _ = fi36
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            var v37 int32 = (v5 << (((^(v18))) & 63))
                            _ = v37
                        }
                    case -980:
                        v2 += (v5 << ((v18) & 63))
                    default:
                        _ = len(v3)
                    }
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = 9223372036854775807
                    }
                default:
                    v38 := GenS2{F0: 434}
                    _ = v38
                }
                v19 -= v19
            }
        }
    }
    v16 = v16
    switch 2147483647 {
    case 127:
        v3 = append(v3, 9223372036854775807)
    case -9223372036854775808:
        var v39 uint = v18
        _ = v39
    default:
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 int = 2147483647
            v16 = (v17 >= v6)
            switch fi40 {
            case -644:
                v42 := GenS3{F0: v18}
                _ = v42
            case -9223372036854775808:
                for fi43 := 0; fi43 < 5; fi43++ {
                    v44 := GenS2{F0: -225}
                    fmt.Println(v20, v41)
                    var v45 uint8 = 42
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
            _ = fi40
            _ = v41
        }
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
