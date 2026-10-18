package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64, p1 int8, p2 int32) int {
    var v7 uint16 = 489
    fmt.Println(p1, p1)
    v8 := GenS1{F0: p0, F1: 755, F2: p2}
    if (-72 >= (p0 ^ p0)) {
        var v9 float64 = float64(1e10)
        switch "a" {
        case "":
            v7 = 1
        case "hello":
            var v10 int = -2147483648
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(p2, p0)
            v7 = 255
            _ = fi11
        }
        _ = v9
    } else {
        switch 133 {
        case 32767:
            var v12 uint16 = v7
            _ = v12
        default:
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS2{F0: v7, F1: "foo bar"}
                _ = fi13
                _ = v14
            }
        }
        v7 += 0
    }
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -833
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -1
    v3 := []uint32{42, 101, 2147483648, 42}
    v4 := []string{}
    v5 := make(map[string]int)
    v6 := []string{}
    var v15 int8 = -1
    v15 += 100
    var v16 float64 = 0.1
    switch "\n\t" {
    case "\n\t":
        if false {
            var v17 uint8 = 22
            _ = len(v1)
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint = 1
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    }
    if (false || false) {
        v15 += (v15 + -1)
        var v20 int = -1
        _ = v20
    }
    v16 += float64(1.5)
    var v21 int16 = int16(v15)
    var v22 float32 = 1e10
    fmt.Println(v16, v22, v16)
    var v23 string = ("hello" + "a")
    switch int(v2) {
    case 1:
        if (false && false) {
            v16 -= (v16 * v16)
            switch v23 {
            case "a":
                v24 := GenS2{F0: 42, F1: "\n\t"}
                _ = v24
            case "":
                fmt.Println(v2, v22, v15)
            case "\n\t":
                for fi25 := 0; fi25 < 9; fi25++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi25
                }
            }
            switch genFn0((v2 ^ v2), v15, 353) {
            case 845:
                if false {
                    v16 = (3.14159 / ((v16 - (((-2.25 - ((v16 + 0.0) + (v16 * (v16 / 1e10)))) / 1.5) * 0.1)) + 3.14159))
                }
            case 32768:
                var v26 string = (v23 + "hello")
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 int16 = (v21 << ((1) & 63))
                fmt.Println(v28, v2)
                _ = fi27
                _ = v28
            }
            if true {
                for fi29 := 0; fi29 < 17; fi29++ {
                    switch fi29 {
                    case -804:
                        v30 := GenS3{F0: v15}
                        _ = v30
                    case 4294967295:
                        v15 = 1
                    }
                    _ = fi29
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    v15 += v15
                    switch (fi31 >> ((42) & 63)) {
                    case 32767:
                        if len(v4) > 0 {
                            v4[0] = v23
                        }
                    case 2147483647:
                        if true {
                            for fi32 := 0; fi32 < 17; fi32++ {
                                var v33 bool = (!false)
                                if v33 {
                                    var v34 string = (v23 + "")
                                    _ = v34
                                }
                                _ = fi32
                                _ = v33
                            }
                        }
                    case 4294967295:
                        v22 = v22
                    default:
                        fmt.Println(v22)
                    }
                    if (true && true) {
                        v35 := GenS4{F0: 32767}
                        if false {
                            v16 = float64((v22 - ((v22 + (3.14159 - -2.25)) - v22)))
                        }
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                var v36 int16 = (v21 - (v21 / (((v21 & (v21 - (v21 * v21)))) | 1)))
                                _ = v36
                            case "hello":
                                v37 := GenS1{F0: v2, F1: 2, F2: -1}
                                _ = v37
                            case "a":
                                v23 = ""
                            }
                            v38 := GenS2{F0: 0, F1: "foo bar"}
                            _ = v38
                        }
                        _ = v35
                    }
                    _ = fi31
                }
            } else {
                if (v22 > v22) {
                    v23 = (("foo bar" + v23) + "")
                    v4 = append(v4, v23)
                    var v39 int = 42
                    _ = v39
                }
            }
            v40 := GenS1{F0: v2, F1: 2, F2: 42}
            _ = v40
        }
    }
    var v41 int16 = v21
    v42 := GenS5{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
