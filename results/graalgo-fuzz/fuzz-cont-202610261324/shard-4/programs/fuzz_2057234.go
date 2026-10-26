package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "foo bar"
    v3 := make(map[int]uint32)
    var v4 float64 = 1.5
    var v5 float64 = 0.0
    v5 = 0.0
    if (false && false) {
        if false {
            var v6 bool = false
            switch (v2 + ((v2 + ("" + ("foo bar" + v2))) + v2)) {
            case "":
                fmt.Println(v2, v4)
            case "\n\t":
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 bool = (127 > -1)
                    fmt.Println(v2, v2, v8)
                    _ = fi7
                    _ = v8
                }
            case "hello":
                v9 := []int{}
                _ = v9
            }
            _ = v6
        }
    } else {
        var v10 float32 = (float32((float64(-2.25) / v4)) / (0.1 / (((1.5 * 0.1) / 0.1) / 1e10)))
        v11 := v3[-2]
        if false {
            if (0 >= 1) {
                var v12 bool = false
                v10 = 1e10
                v13 := GenS1{F0: -85}
                _ = v12
                _ = v13
            } else {
                _ = len(v1)
                v10 = v10
            }
        }
        _ = v10
        _ = v11
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint8 = 123
            _ = fi15
            _ = v16
        }
        var v17 int16 = 32767
        if true {
            var v18 bool = true
            var v19 int32 = -2147483648
            var v20 string = "hello"
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi14
        _ = v17
    }
    if (!(float64((0.0 + 0.0)) <= v5)) {
        if (!true) {
            if false {
                var v21 int64 = 0
                v3[32768] = 115
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 float32 = 1.5
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 float32 = 3.14159
                    var v26 int8 = 42
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
        }
        var v27 string = (v2 + "a")
        if true {
            v28 := GenS1{F0: 0}
            _ = v28
        } else {
            v29 := GenS1{F0: 42}
            v2 = "hello"
            switch -1 {
            case 32767:
                for fi30 := 0; fi30 < 3; fi30++ {
                    v2 = ("foo bar" + v2)
                    _ = fi30
                }
            case -1:
                var v31 bool = (42 <= 2147483647)
                _ = v31
            case 9223372036854775807:
                v5 = 1e10
            }
            _ = v29
        }
        _ = v27
    } else {
        v32 := GenS2{F0: -1, F1: 18446744073709551615}
        switch -1 {
        case -146:
            v33 := GenS1{F0: -128}
            _ = v33
        case -922:
            if (1 < 521) {
                v34 := GenS2{F0: 1, F1: 42}
                var v35 float32 = (1e10 * 0.1)
                _ = v34
                _ = v35
            } else {
                var v36 uint16 = 1
                if (59 < 42) {
                    switch 42 {
                    case -262:
                        switch "\n\t" {
                        case "a":
                            for fi37 := 0; fi37 < 15; fi37++ {
                                var v38 uint8 = 127
                                _ = fi37
                                _ = v38
                            }
                        case "":
                            if (-2.25 > 1e10) {
                                var v39 uint64 = 395
                                v2 = v2
                                v40 := GenS3{F0: true, F1: -181}
                                _ = v39
                                _ = v40
                            } else {
                                if true {
                                    v5 += v5
                                    v36 = 255
                                } else {
                                    var v41 uint64 = 18446744073709551615
                                    v1[32768] = -2147483648
                                    _ = v41
                                }
                                v4 += v4
                                for fi42 := 0; fi42 < 7; fi42++ {
                                    v2 = ((v2 + v2) + v2)
                                    _ = fi42
                                }
                            }
                        case "\n\t":
                            v5 += (-2.25 / v5)
                        default:
                            var v43 int64 = 538
                            _ = v43
                        }
                    case 128:
                        v5 += 1.5
                    }
                    var v44 int64 = 9223372036854775807
                    fmt.Println(v2)
                    _ = v44
                }
                v45 := GenS3{F0: (-965 <= -889), F1: -800}
                _ = v36
                _ = v45
            }
        default:
            var v46 string = (v2 + v2)
            _ = v46
        }
        _ = v32
    }
    v47 := GenS2{F0: -618, F1: 0}
    if true {
        switch (v2 + v2) {
        case "foo bar":
            v2 = v2
        case "":
            var v48 float32 = float32((v5 - v4))
            _ = v48
        }
    }
    fmt.Println(v5, v5)
    fmt.Println(v4, v5, v5)
    var v49 int16 = 1000
    var v50 int16 = (-(v49))
    _ = len(v1)
    for fi51 := 0; fi51 < 12; fi51++ {
        if true {
            fmt.Println(v49, v4)
        }
        _ = fi51
    }
    var v52 uint16 = 32767
    v5 -= (v5 + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v52
}
