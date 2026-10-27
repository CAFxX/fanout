package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint32) string {
    p0 = 2
    p0 = p0
    p0 = 2
    var v7 uint32 = p1
    p1 = 244
    _ = p0
    _ = p1
    _ = v7
    return "foo bar"
}
func genFn1(p0 int32, p1 bool) int64 {
    v8 := []int{}
    var v9 uint64 = 42
    v8 = append(v8, 2147483647)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{32767}
    var v3 float64 = float64(-2.25)
    v4 := make(map[uint32]string)
    var v5 int = 2147483647
    var v6 uint = 18446744073709551615
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(fi10, fi10)
        if len(v2) > 0 {
            v2[0] = (v5 << (((v6 ^ v6)) & 63))
        }
        _ = len(v1)
        _ = fi10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: -2147483648, F1: 750}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: (v5 << (((v6 % ((v6) | 1))) & 63)), F1: -836}
    _ = len(v4)
    v3 += 0.0
    for fi14 := 0; fi14 < 7; fi14++ {
        fmt.Println(v6, v3, fi14)
        _ = fi14
    }
    _ = len(v4)
    switch "hello" {
    case "foo bar":
        if true {
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v5)
                _ = fi15
            }
            v6 = (v6 & (v6 | 65535))
        } else {
            if (true && ((v3 + v3) < (v3 / float64(((3.14159 + 1.5) * 0.1))))) {
                fmt.Println(v3, v3, v6)
            } else {
                switch "hello" {
                case "":
                    if true {
                        v16 := GenS1{F0: -9223372036854775808, F1: 1}
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v6 = (v6 / (((v6 - 65535)) | 1))
                            if false {
                                v18 := GenS1{F0: -9223372036854775808, F1: -10}
                                _ = v18
                            } else {
                                v3 += ((v3 * v3) + v3)
                            }
                            v19 := GenS1{F0: (fi17 ^ 595), F1: 881}
                            _ = fi17
                            _ = v19
                        }
                        _ = v16
                    } else {
                        v5 += v5
                        v5 += v5
                    }
                default:
                    var v20 string = "foo bar"
                    _ = v20
                }
                switch (v5 | 852) {
                case 32768:
                    switch (v5 % ((v5) | 1)) {
                    case 32768:
                        for fi21 := 0; fi21 < 7; fi21++ {
                            switch fi21 {
                            case 128:
                                v5 = 1
                            case 42:
                                fmt.Println(v6, v3, v5)
                            }
                            for fi22 := 0; fi22 < 11; fi22++ {
                                v23 := GenS1{F0: fi22, F1: 369}
                                _ = len(v4)
                                v6 = (v6 % ((v6) | 1))
                                fmt.Println(fi21)
                                _ = fi22
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 2; fi24++ {
                                v25 := GenS1{F0: fi24, F1: 1000}
                                if true {
                                    var v26 uint = v6
                                    _ = v26
                                }
                                if (168 != 255) {
                                    var v27 int16 = 1000
                                    _ = v27
                                }
                                _ = fi24
                                _ = v25
                            }
                            _ = fi21
                        }
                    case 9223372036854775807:
                        v6 -= (v6 >> (((v6 & v6)) & 63))
                    case 2147483647:
                        if (v6 == v6) {
                            var v28 bool = (false || true)
                            var v29 uint32 = 129
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS1{F0: -9223372036854775808, F1: 32767}
                            _ = v30
                        }
                    default:
                        v5 += (v5 / ((v5) | 1))
                    }
                case -1:
                    v3 = v3
                case 1:
                    var v31 uint32 = 0
                    _ = v31
                default:
                    var v32 int8 = 0
                    _ = v32
                }
                if false {
                    v33 := GenS2{F0: 253, F1: (("foo bar" == "hello") || (-32768 >= 1000))}
                    v3 = 0.0
                    v34 := GenS3{F0: 2, F1: (true || false)}
                    _ = v33
                    _ = v34
                } else {
                    if true {
                        var v35 float32 = 1.5
                        _ = v35
                    }
                    var v36 uint64 = 42
                    v36 = v36
                    _ = v36
                }
            }
        }
    case "":
        if (false && false) {
            for fi37 := 0; fi37 < 18; fi37++ {
                if true {
                    var v38 uint64 = 18446744073709551615
                    fmt.Println(v3, v5, v3)
                    _ = v38
                }
                _ = fi37
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch (("" + "foo bar") + "\n\t") {
            case "\n\t":
                v5 += v5
            case "hello":
                fmt.Println(v5, v5)
            }
        } else {
            v5 = -645
        }
    case "\n\t":
        for fi39 := 0; fi39 < 11; fi39++ {
            v40 := GenS4{F0: true}
            delete(v1, v5)
            _ = fi39
            _ = v40
        }
    default:
        var v41 int8 = 42
        _ = v41
    }
    v42 := v4[65535]
    v5 += -1
    v43 := GenS1{F0: (v5 | v5), F1: 42}
    if true {
        v44 := GenS5{F0: -732}
        v45 := GenS5{F0: -550}
        if false {
            if false {
                v46 := GenS3{F0: 2, F1: true}
                _ = v46
            } else {
                var v47 int64 = -9223372036854775808
                _ = len(v2)
                v48 := v4[2]
                _ = v47
                _ = v48
            }
        }
        _ = v44
        _ = v45
    } else {
        var v49 int32 = 0
        fmt.Println(v3, v42)
        _ = v49
    }
    var v50 int = v5
    if (!(2147483647 <= 46)) {
        for fi51 := 0; fi51 < 14; fi51++ {
            v5 -= 127
            var v52 uint64 = 42
            _ = fi51
            _ = v52
        }
    } else {
        v50 = (v50 - -2147483648)
        v53 := GenS3{F0: 888, F1: ((v6 & (v6 - (^(v6)))) <= (v6 * (v6 >> (((v6 << ((1) & 63))) & 63))))}
        _ = v53
    }
    fmt.Println(v6, v42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v50
    _ = v6
}
