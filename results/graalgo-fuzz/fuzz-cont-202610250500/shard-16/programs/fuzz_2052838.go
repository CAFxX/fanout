package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int8; F1 string }

func main() {
    var v1 uint8 = 197
    v2 := make(map[int]uint32)
    var v3 uint = 65535
    var v4 bool = false
    var v5 string = "a"
    v6 := make(map[uint32]int)
    v2[128] = 4294967295
    if v4 {
        v5 = ""
        switch 127 {
        case 128:
            if (0 <= 0) {
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 float32 = -2.25
                    _ = fi7
                    _ = v8
                }
                switch ("\n\t" + (v5 + v5)) {
                case "hello":
                    switch -600 {
                    case 32768:
                        v9 := GenS1{F0: 698}
                        _ = v9
                    case -725:
                        for fi10 := 0; fi10 < 12; fi10++ {
                            var v11 int = fi10
                            if v4 {
                                for fi12 := 0; fi12 < 15; fi12++ {
                                    var v13 uint = (v3 * 255)
                                    _ = fi12
                                    _ = v13
                                }
                                for fi14 := 0; fi14 < 1; fi14++ {
                                    var v15 int64 = 0
                                    v16 := GenS1{F0: 0}
                                    v17 := GenS1{F0: fi14}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 2; fi18++ {
                                    v19 := GenS1{F0: v11}
                                    _ = fi18
                                    _ = v19
                                }
                            }
                            _ = fi10
                            _ = v11
                        }
                    default:
                        if ((!((v4 || v4) || v4)) && v4) {
                            v20 := GenS2{F0: -1}
                            _ = v20
                        }
                    }
                case "a":
                    v21 := GenS3{F0: -32768}
                    _ = v21
                case "foo bar":
                    switch v5 {
                    case "\n\t":
                        v6[2] = 42
                    case "hello":
                        switch v5 {
                        case "foo bar":
                            v4 = false
                        case "hello":
                            var v22 int64 = -9223372036854775808
                            _ = v22
                        case "":
                            if false {
                                if v4 {
                                    fmt.Println(v3, v3)
                                }
                                var v23 uint16 = 65535
                                _ = v23
                            }
                        default:
                            _ = len(v6)
                        }
                    }
                default:
                    var v24 int8 = 71
                    _ = v24
                }
            } else {
                if ((1 <= 100) && v4) {
                    fmt.Println(v3, v1, v3)
                    var v25 uint8 = 2
                    v26 := v2[2147483647]
                    _ = v25
                    _ = v26
                }
                var v27 int32 = 42
                for fi28 := 0; fi28 < 12; fi28++ {
                    if (v4 || false) {
                        v4 = (((v27 * -77) <= (v27 % ((v27) | 1))) || (!v4))
                    } else {
                        if ((false || v4) || v4) {
                            v29 := GenS2{F0: 0}
                            _ = v29
                        }
                        v3 = v3
                    }
                    for fi30 := 0; fi30 < 4; fi30++ {
                        if (v4 || v4) {
                            var v31 string = ("a" + (("\n\t" + ((v5 + ((v5 + v5) + v5)) + (((((v5 + ((v5 + "foo bar") + v5)) + v5) + ("" + v5)) + v5) + v5))) + "a"))
                            _ = v31
                        }
                        if (-2147483648 != (fi30 * fi30)) {
                            var v32 int16 = 1000
                            _ = v32
                        }
                        var v33 float32 = -2.25
                        _ = fi30
                        _ = v33
                    }
                    _ = fi28
                }
                _ = v27
            }
        case 0:
            v4 = v4
        default:
            v34 := GenS1{F0: -1}
            _ = v34
        }
        var v35 int64 = 0
        _ = v35
    } else {
        if v4 {
            var v36 int32 = 42
            _ = v36
        } else {
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 int = fi37
                var v39 float32 = 3.14159
                if v4 {
                    switch "\n\t" {
                    case "hello":
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v5 = (("foo bar" + "hello") + (v5 + (v5 + v5)))
                            var v41 int64 = 47
                            _ = fi40
                            _ = v41
                        }
                    case "foo bar":
                        v42 := GenS4{F0: 1, F1: 264}
                        _ = v42
                    }
                }
                _ = fi37
                _ = v38
                _ = v39
            }
            if v4 {
                for fi43 := 0; fi43 < 17; fi43++ {
                    var v44 int8 = 100
                    _ = fi43
                    _ = v44
                }
            }
            v45 := GenS1{F0: -418}
            _ = v45
        }
    }
    if (73 >= (v1 >> ((42) & 63))) {
        v46 := GenS5{F0: v3, F1: (^(v3))}
        for fi47 := 0; fi47 < 18; fi47++ {
            v1 -= (v1 + (v1 | 127))
            fmt.Println(fi47, fi47)
            v4 = v4
            _ = fi47
        }
        _ = v46
    } else {
        v48 := GenS1{F0: -9223372036854775808}
        v1 -= (v1 * v1)
        _ = v48
    }
    v49 := []int{}
    var v50 int = 2147483647
    _ = len(v6)
    for fi51 := 0; fi51 < 18; fi51++ {
        v52 := GenS6{F0: 8, F1: v5}
        var v53 uint8 = 128
        var v54 uint = uint(v53)
        _ = fi51
        _ = v52
        _ = v53
        _ = v54
    }
    v1 -= (v1 % (((^(v1))) | 1))
    v50 = v50
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
