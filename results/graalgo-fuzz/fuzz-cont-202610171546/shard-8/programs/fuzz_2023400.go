package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint8; F2 int8 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 string = "a"
    var v4 uint16 = 42
    v5 := make(map[uint32]uint32)
    v5[2147483648] = uint32(v4)
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 127}
        _ = fi6
        _ = v7
    }
    v3 = (v3 + ("a" + ""))
    switch v3 {
    case "\n\t":
        var v8 int64 = 822
        _ = v8
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            v1 = false
            v10 := GenS1{F0: 42}
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint8 = 255
                if false {
                    var v13 uint8 = 187
                    _ = v13
                } else {
                    var v14 int = (fi11 - fi9)
                    var v15 int32 = 1
                    var v16 int = (fi11 >> ((1) & 63))
                    _ = v14
                    _ = v15
                    _ = v16
                }
                v3 = v3
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v1, v1, v3)
    var v17 uint32 = 180
    v3 = "\n\t"
    var v18 uint = 214
    var v19 int64 = 224
    for fi20 := 0; fi20 < 12; fi20++ {
        switch 79 {
        case 2147483647:
            var v21 int64 = v19
            _ = v21
        case -2147483648:
            if false {
                delete(v5, (v17 % ((v17) | 1)))
            } else {
                switch v3 {
                case "":
                    if ((^(v18)) <= v18) {
                        v17 = (v17 & 4294967295)
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 float64 = 3.14159
                            var v24 uint32 = 4294967295
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS2{F0: 1000}
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 0}
                        var v27 string = v3
                        _ = v26
                        _ = v27
                    }
                }
                var v28 string = v3
                _ = v28
            }
        case 42:
            switch (("foo bar" + ("hello" + "foo bar")) + "") {
            case "hello":
                fmt.Println(v3)
            case "foo bar":
                var v29 string = "\n\t"
                _ = v29
            }
        default:
            v17 = (^(v17))
        }
        _ = fi20
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        fmt.Println(fi30, fi30)
        for fi31 := 0; fi31 < 1; fi31++ {
            if (uint64(v4) > 42) {
                v32 := GenS1{F0: 213}
                _ = v32
            } else {
                v33 := []int{}
                _ = v33
            }
            var v34 uint32 = v17
            var v35 int64 = (v19 / ((v19) | 1))
            _ = fi31
            _ = v34
            _ = v35
        }
        var v36 string = ("" + v3)
        _ = fi30
        _ = v36
    }
    fmt.Println(v1, v4)
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 2; fi38++ {
            switch (v3 + v3) {
            case "foo bar":
                switch v3 {
                case "hello":
                    if v1 {
                        if v1 {
                            if v1 {
                                var v39 int8 = 42
                                var v40 int8 = v39
                                _ = v39
                                _ = v40
                            }
                            if true {
                                fmt.Println(v1)
                                var v41 int = (fi37 ^ 128)
                                v1 = (uint64(v17) != 4294967295)
                                switch v3 {
                                case "":
                                    if false {
                                        fmt.Println(v4, v41)
                                        var v42 float64 = 0.1
                                        _ = v42
                                    }
                                }
                                fmt.Println(v1)
                                v4 -= uint16(v18)
                                _ = v41
                            }
                            var v43 uint32 = (v17 << (((v18 << (((^(v18))) & 63))) & 63))
                            _ = v43
                        }
                    }
                case "":
                    fmt.Println(v19)
                }
            case "":
                v4 = v4
            default:
                var v44 int64 = -9223372036854775808
                if (v17 >= 517) {
                    v45 := []int{fi38, (fi38 - 561)}
                    _ = v45
                } else {
                    v18 = v18
                    if v1 {
                        var v46 uint8 = 255
                        _ = v46
                    }
                }
                _ = v44
            }
            var v47 uint64 = 2
            v48 := GenS2{F0: -32768}
            _ = fi38
            _ = v47
            _ = v48
        }
        v17 += (^(v17))
        _ = fi37
    }
    fmt.Println(v4, v18, v3)
    v2[uint32(v18)] = v3
    v49 := []int{127, 1}
    if v1 {
        _ = len(v49)
    } else {
        v50 := GenS2{F0: 284}
        if (v19 >= (^(v19))) {
            fmt.Println(v1)
            v51 := GenS3{F0: "", F1: 127, F2: 42}
            v52 := GenS3{F0: "", F1: 0, F2: -1}
            _ = v51
            _ = v52
        }
        v18 = (v18 * 255)
        _ = v50
    }
    v53 := GenS4{F0: 484}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
}
