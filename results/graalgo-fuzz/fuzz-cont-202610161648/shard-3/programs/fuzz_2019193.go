package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int64; F2 int32 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 127
    var v3 uint64 = 42
    var v4 int16 = -32768
    var v5 int64 = 9223372036854775807
    _ = len(v1)
    var v6 string = "foo bar"
    v3 -= v3
    switch ((v6 + "foo bar") + "\n\t") {
    case "a":
        v3 = 0
    default:
        v7 := []int{}
        _ = v7
    }
    delete(v1, 909)
    switch 127 {
    case 127:
        if true {
            switch 32768 {
            case 9223372036854775807:
                v6 = ("hello" + (v6 + v6))
            default:
                if false {
                    if true {
                        var v8 bool = false
                        _ = v8
                    } else {
                        v4 = (v4 * -276)
                        v1[32768] = v6
                    }
                    switch 42 {
                    case 4294967295:
                        if (false && (2 > 65535)) {
                            if (119 == 1) {
                                v9 := []int{-542, 641, -1, 703}
                                v1[0] = v6
                                v10 := v1[580]
                                _ = v10
                                _ = v9
                            }
                            if false {
                                delete(v1, 813)
                            }
                            if false {
                                if (775 != -878) {
                                    var v11 uint8 = 1
                                    var v12 float64 = (1e10 + 3.14159)
                                    _ = v11
                                    _ = v12
                                } else {
                                    var v13 string = "foo bar"
                                    v6 = "a"
                                    _ = v13
                                }
                                for fi14 := 0; fi14 < 2; fi14++ {
                                    for fi15 := 0; fi15 < 7; fi15++ {
                                        var v16 uint64 = (^(v3))
                                        _ = fi15
                                        _ = v16
                                    }
                                    v17 := GenS1{F0: fi14, F1: 492, F2: 127}
                                    _ = fi14
                                    _ = v17
                                }
                                v3 = (v3 >> ((729) & 63))
                            } else {
                                fmt.Println(v2)
                            }
                        }
                    case 643:
                        var v18 bool = (!(!(128 <= -2147483648)))
                        _ = v18
                    default:
                        v19 := GenS1{F0: -2147483648, F1: 882, F2: v2}
                        _ = v19
                    }
                    switch -1 {
                    case 0:
                        v20 := GenS2{F0: (v3 / ((v3) | 1))}
                        _ = v20
                    case 1:
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v6 = v6
                            _ = fi21
                        }
                    default:
                        v22 := GenS3{F0: ((-12 > 611) && false)}
                        _ = v22
                    }
                } else {
                    v3 -= v3
                }
            }
        } else {
            var v23 string = ("hello" + (v6 + "a"))
            _ = v23
        }
    }
    v4 = v4
    v24 := []int{42, -347, 4294967295}
    var v25 uint16 = 193
    var v26 uint32 = 2
    if true {
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 uint64 = 4294967295
            v3 = v3
            _ = fi27
            _ = v28
        }
        v4 -= (v4 - v4)
    }
    if (v26 >= v26) {
        if true {
            switch v6 {
            case "":
                if len(v24) > 0 {
                    v24[0] = 32767
                }
            case "hello":
                for fi29 := 0; fi29 < 15; fi29++ {
                    if false {
                        _ = len(v1)
                    } else {
                        for fi30 := 0; fi30 < 16; fi30++ {
                            if ((v5 | (^(v5))) == v5) {
                                v31 := GenS1{F0: fi30, F1: v3, F2: v2}
                                if len(v24) > 0 {
                                    v24[0] = (fi29 + (fi30 ^ (fi29 + (-(fi29)))))
                                }
                                _ = v31
                            }
                            switch v6 {
                            case "foo bar":
                                if false {
                                    if true {
                                        if true {
                                            var v32 int64 = v5
                                            _ = v32
                                        } else {
                                            var v33 int16 = (v4 >> ((255) & 63))
                                            _ = v33
                                        }
                                    } else {
                                        if false {
                                            if true {
                                                v2 -= v2
                                                v1[(^(fi29))] = v6
                                            } else {
                                                v34 := GenS4{F0: -475, F1: v5, F2: 408}
                                                _ = v34
                                            }
                                            var v35 uint64 = (v3 << ((1) & 63))
                                            var v36 int8 = (-(v2))
                                            v4 -= (^(v4))
                                            _ = len(v24)
                                            _ = v35
                                            _ = v36
                                        } else {
                                            v5 = 0
                                            var v37 uint8 = uint8(v3)
                                            v37 -= v37
                                            _ = v37
                                        }
                                    }
                                    v38 := GenS2{F0: (^(v3))}
                                    switch v6 {
                                    case "hello":
                                        if len(v24) > 0 {
                                            v24[0] = (-(fi29))
                                        }
                                    default:
                                        _ = len(v24)
                                    }
                                    _ = v38
                                }
                            case "a":
                                v39 := v1[fi30]
                                _ = v39
                            default:
                                if ((1 <= 0) || false) {
                                    var v40 int8 = v2
                                    _ = v40
                                }
                            }
                            _ = fi30
                        }
                        v41 := GenS5{F0: (-(v4))}
                        v42 := GenS3{F0: true}
                        _ = v41
                        _ = v42
                    }
                    if (!true) {
                        if false {
                            if len(v24) > 0 {
                                v24[0] = (-(fi29))
                            }
                            var v43 uint8 = 137
                            _ = v43
                        } else {
                            var v44 int64 = 1
                            fmt.Println(v4)
                            var v45 int8 = v2
                            _ = v44
                            _ = v45
                        }
                        var v46 uint32 = v26
                        var v47 int = fi29
                        _ = v46
                        _ = v47
                    }
                    _ = fi29
                }
            default:
                var v48 uint32 = v26
                _ = v48
            }
        }
        for fi49 := 0; fi49 < 4; fi49++ {
            var v50 string = "a"
            _ = fi49
            _ = v50
        }
    }
    fmt.Println(v6, v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
