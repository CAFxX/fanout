package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int64; F2 int8 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 uint8; F1 uint }

func main() {
    var v1 float64 = (((3.14159 / 1.5) - (0.1 / float64((1.5 - 0.1)))) / 0.1)
    v2 := []int{4294967295, 42, -9223372036854775808}
    var v3 uint64 = 0
    var v4 string = ("" + "a")
    var v5 int16 = 1
    if ((false || (975 <= 32767)) && false) {
        v5 = (v5 * v5)
        var v6 int16 = 530
        var v7 int16 = 0
        _ = v6
        _ = v7
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: -1}
        _ = fi8
        _ = v9
    }
    var v10 int32 = 42
    var v11 uint64 = v3
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            if len(v2) > 0 {
                v2[0] = (fi12 << ((65535) & 63))
            }
            if false {
                fmt.Println(fi13, v1, v5)
            } else {
                v5 = -1
                fmt.Println(fi12)
                var v14 uint8 = 70
                _ = v14
            }
            _ = fi13
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(fi15)
            fmt.Println(v1, v11, fi12)
            _ = fi15
        }
        _ = fi12
    }
    if true {
        var v16 uint = 4294967295
        v1 = v1
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS1{F0: -1}
            for fi19 := 0; fi19 < 1; fi19++ {
                _ = len(v2)
                if false {
                    v11 = v3
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        var v20 int32 = (v10 >> (((v16 / (((v16 | (v16 * v16))) | 1))) & 63))
                        if true {
                            v16 += v16
                            if false {
                                v3 += (^(v11))
                                v21 := GenS2{F0: 0}
                                if (0 <= int64(v10)) {
                                    v2 = append(v2, -999)
                                    v4 = (v4 + v4)
                                    switch (((((((((v4 + v4) + v4) + v4) + "foo bar") + ("\n\t" + (v4 + ("\n\t" + (("foo bar" + v4) + "foo bar"))))) + ("foo bar" + v4)) + "hello") + ("foo bar" + "hello")) + v4) {
                                    case "hello":
                                        v4 = v4
                                    case "foo bar":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "a":
                                        v2 = append(v2, (^(fi17)))
                                    }
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v22 uint8 = 1
                                var v23 int8 = 0
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                        } else {
                            var v24 string = ("a" + v4)
                            v25 := GenS1{F0: -2147483648}
                            if (false && (814 < 0)) {
                                switch v24 {
                                case "a":
                                    var v26 int = fi17
                                    _ = v26
                                default:
                                    if true {
                                        if true {
                                            v1 = ((v1 - (-2.25 / v1)) - v1)
                                            v27 := GenS3{F0: (v16 - (v16 % ((v16) | 1))), F1: 1}
                                            v20 -= 2147483647
                                            if len(v2) > 0 {
                                                v2[0] = (fi19 << ((944) & 63))
                                            }
                                            _ = v27
                                        } else {
                                            switch fi19 {
                                            case 42:
                                                if (false || ("" < v24)) {
                                                    var v28 uint32 = 65535
                                                    _ = v28
                                                }
                                            case 9223372036854775807:
                                                var v29 int32 = v20
                                                _ = v29
                                            case -2147483648:
                                                switch fi19 {
                                                case 937:
                                                    var v30 uint8 = 2
                                                    _ = v30
                                                case 128:
                                                    v24 = v24
                                                    var v31 int16 = (^(v5))
                                                    v32 := GenS4{F0: v10, F1: 42}
                                                    _ = v31
                                                    _ = v32
                                                default:
                                                    v33 := GenS5{F0: (false || ((true || false) && (v1 != v1)))}
                                                    _ = v33
                                                }
                                            }
                                        }
                                        switch ("\n\t" + (("" + ("foo bar" + v24)) + (((v4 + ("a" + "foo bar")) + v4) + "foo bar"))) {
                                        case "hello":
                                            var v34 uint16 = 42
                                            _ = v34
                                        }
                                        var v35 uint8 = 1
                                        _ = v35
                                    }
                                }
                            }
                            _ = v24
                            _ = v25
                        }
                        if (false && (v5 == v5)) {
                            v3 -= (v3 - (^(v11)))
                            v5 = (v5 - (v5 * v5))
                        }
                        _ = v20
                    } else {
                        switch v4 {
                        case "a":
                            var v36 float64 = v1
                            _ = v36
                        case "foo bar":
                            var v37 uint8 = 236
                            _ = v37
                        }
                    }
                    v10 -= (-(v10))
                    if (true || true) {
                        var v38 int16 = -357
                        v39 := GenS1{F0: 1}
                        switch v4 {
                        case "\n\t":
                            v40 := GenS4{F0: v10, F1: 858}
                            _ = v40
                        default:
                            if false {
                                if true {
                                    if (!(!false)) {
                                        var v41 float64 = v1
                                        _ = len(v2)
                                        _ = v41
                                    } else {
                                        v5 = 1000
                                    }
                                    v42 := GenS2{F0: -1}
                                    _ = v42
                                }
                            }
                        }
                        _ = v38
                        _ = v39
                    }
                    fmt.Println(v1)
                }
                _ = fi19
            }
            if (128 != uint8(v16)) {
                if true {
                    v43 := GenS6{F0: (("foo bar" + v4) + (((v4 + ((v4 + v4) + "\n\t")) + "a") + "hello")), F1: 42, F2: -128}
                    _ = v43
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if false {
                    v44 := GenS7{F0: (fi17 / ((890) | 1)), F1: 2147483648}
                    var v45 int64 = -9223372036854775808
                    _ = v44
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        v2 = append(v2, -2147483648)
                        _ = fi46
                    }
                    var v47 bool = true
                    _ = v47
                }
            } else {
                v3 = (v3 * v11)
                var v48 int8 = 9
                _ = v48
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v49 := GenS7{F0: 17, F1: 42}
    var v50 uint8 = 0
    if (false || true) {
        fmt.Println(v3, v4)
        for fi51 := 0; fi51 < 8; fi51++ {
            switch "foo bar" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = 408
                }
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi51
        }
        v2 = append(v2, -1)
    }
    switch -641 {
    case -409:
        var v52 uint16 = 586
        _ = v52
    }
    _ = len(v2)
    v53 := GenS8{F0: v50, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v53
}
