package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0(p0 int, p1 uint16, p2 uint8) int {
    var v7 bool = true
    if (!true) {
        fmt.Println(p1)
        p1 = p1
    }
    v7 = v7
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 * (p0 / ((p0) | 1)))
}
func genFn1(p0 uint8) uint32 {
    var v8 uint32 = 176
    v9 := []int{1, 894, 4294967295}
    fmt.Println(p0)
    _ = p0
    _ = v8
    _ = v9
    return (v8 & v8)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 955
    v3 := make(map[int]string)
    var v4 int32 = 0
    v5 := make(map[string]int)
    v6 := []uint32{411, 0}
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                v2 = 1000
                var v12 uint = 396
                _ = v12
            } else {
                var v13 string = "\n\t"
                _ = v13
            }
            _ = fi11
        }
        _ = fi10
    }
    var v14 int16 = (v2 & (-(v2)))
    v4 = -420
    for fi15 := 0; fi15 < 2; fi15++ {
        switch "foo bar" {
        case "hello":
            v14 += -32768
        case "foo bar":
            if (!false) {
                switch 2147483647 {
                case -2147483648:
                    var v16 int16 = v14
                    _ = v16
                }
            }
        default:
            fmt.Println(v14, v14, v4)
        }
        if (992 <= 42) {
            for fi17 := 0; fi17 < 3; fi17++ {
                fmt.Println(v2, fi17)
                v4 = int32(fi17)
                if (true && false) {
                    if true {
                        var v18 int64 = 9223372036854775807
                        v19 := v1[65535]
                        var v20 string = "hello"
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    v4 += v4
                } else {
                    var v21 int32 = v4
                    v14 = (-(v14))
                    var v22 int = (fi15 + genFn0(fi17, 65535, 39))
                    var v23 uint64 = 42
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi17
            }
            if ((v2 * v14) > v14) {
                v24 := GenS1{F0: 134}
                fmt.Println(fi15, v14)
                switch "hello" {
                case "\n\t":
                    if (-395 <= 42) {
                        fmt.Println(v2)
                        if (false && ((v2 << ((190) & 63)) <= v14)) {
                            v25 := GenS1{F0: 128}
                            v4 += v4
                            _ = v25
                        } else {
                            v26 := GenS1{F0: 255}
                            if ((v4 ^ (-(v4))) >= v4) {
                                v27 := GenS1{F0: 118}
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    v29 := GenS1{F0: 171}
                                    _ = fi28
                                    _ = v29
                                }
                                _ = v27
                            } else {
                                switch (fi15 << ((0) & 63)) {
                                case -2147483648:
                                    v14 = v2
                                }
                                v2 = (v14 * v14)
                            }
                            for fi30 := 0; fi30 < 7; fi30++ {
                                v2 -= (v2 ^ -1)
                                _ = fi30
                            }
                            _ = v26
                        }
                        for fi31 := 0; fi31 < 3; fi31++ {
                            v2 = (v2 - 0)
                            _ = fi31
                        }
                    }
                }
                _ = v24
            } else {
                for fi32 := 0; fi32 < 16; fi32++ {
                    fmt.Println(fi32, v4)
                    v33 := GenS2{F0: ""}
                    _ = fi32
                    _ = v33
                }
            }
            var v34 string = "hello"
            _ = v34
        }
        switch genFn0(fi15, 42, 131) {
        case 32767:
            switch ("hello" + ("hello" + "\n\t")) {
            case "hello":
                _ = len(v5)
            default:
                var v35 uint64 = 4294967295
                _ = v35
            }
        default:
            fmt.Println(v14, v2, v4)
        }
        _ = fi15
    }
    var v36 int = 676
    if false {
        var v37 float64 = float64(float32(float64(((((3.14159 * 1e10) * 0.1) * 3.14159) / (1.5 + 1.5)))))
        for fi38 := 0; fi38 < 20; fi38++ {
            v36 = v36
            fmt.Println(v14, v2, v2)
            _ = fi38
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            for fi40 := 0; fi40 < 15; fi40++ {
                v3[(v36 - (^(fi40)))] = "hello"
                v41 := GenS1{F0: 218}
                if (!(!(!false))) {
                    v42 := GenS2{F0: ("\n\t" + "hello")}
                    _ = len(v6)
                    var v43 string = "foo bar"
                    var v44 uint8 = 2
                    v6 = append(v6, 2147483648)
                    _ = v42
                    _ = v43
                    _ = v44
                } else {
                    var v45 uint8 = 2
                    _ = v45
                }
                _ = fi40
                _ = v41
            }
            var v46 bool = (1 <= 255)
            _ = fi39
            _ = v46
        }
        _ = v37
    }
    v47 := GenS3{F0: 151, F1: 4294967295, F2: -355}
    fmt.Println(v14, v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
