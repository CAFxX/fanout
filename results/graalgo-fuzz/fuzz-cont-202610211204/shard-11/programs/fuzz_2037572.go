package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
func genFn0(p0 uint16, p1 uint32, p2 uint8) int32 {
    fmt.Println(p1)
    if false {
        var v6 string = ""
        _ = v6
    } else {
        p0 = p0
        p1 += p1
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(p1)
            p1 = 65535
            var v8 uint32 = 2147483647
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        p2 = 62
        p0 += (^(p0))
        _ = fi9
    }
    var v10 int8 = 1
    if false {
        p2 = (p2 & (p2 << ((2) & 63)))
        p2 -= 0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 717
}
func genFn1() int8 {
    var v11 int64 = 613
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 string = ("hello" + "")
        _ = fi12
        _ = v13
    }
    _ = v11
    return -72
}

func main() {
    var v1 uint32 = 0
    var v2 uint64 = 479
    var v3 int32 = 0
    v4 := []string{"\n\t", "foo bar"}
    v5 := make(map[string]string)
    switch -1 {
    case 32767:
        v3 -= 1
    default:
        var v14 bool = false
        _ = v14
    }
    v3 = v3
    v15 := GenS1{F0: -1}
    fmt.Println(v3, v3, v1)
    if (true && false) {
        if false {
            switch "hello" {
            case "a":
                switch ("" + "\n\t") {
                case "foo bar":
                    _ = len(v5)
                case "":
                    var v16 float32 = 3.14159
                    _ = v16
                default:
                    v5["\n\t"] = ("a" + "")
                }
            case "\n\t":
                if (0 != 1) {
                    v3 = v3
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v17 uint32 = v1
                    if ((1 != 42) || (false && true)) {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            for fi19 := 0; fi19 < 10; fi19++ {
                                fmt.Println(v17, v3)
                                if false {
                                    v3 = v3
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                _ = fi19
                            }
                            v20 := GenS1{F0: 1}
                            _ = fi18
                            _ = v20
                        }
                    } else {
                        var v21 float32 = 3.14159
                        var v22 int16 = -32768
                        v22 = (v22 + v22)
                        _ = v21
                        _ = v22
                    }
                    _ = v17
                }
            case "hello":
                v23 := GenS1{F0: -9223372036854775808}
                _ = v23
            }
            switch "" {
            case "":
                v24 := v5["foo bar"]
                _ = v24
            }
        }
        v3 += v3
    }
    v1 += (v1 + 2)
    var v25 int16 = 1000
    var v26 int = 128
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS1{F0: -468}
        for fi29 := 0; fi29 < 6; fi29++ {
            if true {
                fmt.Println(fi27, v1, v25)
                var v30 int32 = (v3 << ((65535) & 63))
                v31 := GenS1{F0: 4294967295}
                _ = v30
                _ = v31
            }
            if true {
                v1 = 42
                switch (-(fi27)) {
                case -207:
                    v1 = (^(v1))
                case 127:
                    var v32 uint32 = (v1 / ((v1) | 1))
                    _ = v32
                default:
                    v3 += v3
                }
            } else {
                if true {
                    v33 := GenS2{F0: v2, F1: 1, F2: 18446744073709551615}
                    if true {
                        fmt.Println(v1)
                    }
                    if len(v4) > 0 {
                        v4[0] = ("a" + "hello")
                    }
                    var v34 uint16 = 65
                    _ = v33
                    _ = v34
                } else {
                    delete(v5, "foo bar")
                }
                var v35 float64 = 1.5
                v5["hello"] = "\n\t"
                _ = v35
            }
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    var v36 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
