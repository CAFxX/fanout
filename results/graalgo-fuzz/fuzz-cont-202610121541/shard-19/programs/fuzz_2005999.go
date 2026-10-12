package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]int)
    v3 := []string{}
    v4 := make(map[string]uint32)
    v5 := []uint32{67, 654}
    v6 := []uint32{}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 1
            _ = fi8
            _ = v9
        }
        if ((42 < 42) && (false && true)) {
            if ((!true) || false) {
                for fi10 := 0; fi10 < 10; fi10++ {
                    fmt.Println(fi7, fi7, fi10)
                    v1 = v1
                    _ = fi10
                }
                v1 = v1
            } else {
                var v11 uint32 = 4294967295
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint = 65535
                _ = fi12
                _ = v13
            }
            fmt.Println(fi7, v1)
        }
        v1 += (-(v1))
        _ = fi7
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint32 = 555
            switch (fi15 - fi14) {
            case 32767:
                _ = len(v4)
            }
            _ = fi15
            _ = v16
        }
        v1 += 0
        _ = fi14
    }
    if true {
        var v17 uint8 = 128
        switch -1 {
        case 128:
            v1 = (v1 / (((-(v1))) | 1))
        case 32767:
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 float32 = 1e10
                switch "" {
                case "a":
                    var v20 bool = (18446744073709551615 < 18446744073709551615)
                    _ = v20
                case "foo bar":
                    if true {
                        var v21 float64 = 0.0
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
        case -92:
            var v22 int32 = -1
            _ = v22
        }
        var v23 uint64 = 2
        _ = v17
        _ = v23
    } else {
        if (!false) {
            v1 -= (v1 & -1)
            v1 += v1
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS1{F0: 0, F1: 1}
                switch ("\n\t" + ("" + "hello")) {
                case "hello":
                    _ = len(v4)
                case "":
                    var v26 uint16 = 0
                    _ = v26
                }
                for fi27 := 0; fi27 < 12; fi27++ {
                    if true {
                        v28 := GenS2{F0: 128}
                        _ = v28
                    } else {
                        if false {
                            if len(v5) > 0 {
                                v5[0] = 65535
                            }
                            switch ("foo bar" + (("a" + "") + ("" + "\n\t"))) {
                            case "":
                                var v29 float32 = 3.14159
                                _ = v29
                            case "hello":
                                _ = len(v3)
                            case "foo bar":
                                v1 = v1
                            default:
                                var v30 float32 = 0.1
                                _ = v30
                            }
                        } else {
                            v1 = 1
                            v1 += v1
                            v31 := GenS1{F0: -471, F1: 0}
                            _ = v31
                        }
                    }
                    fmt.Println(v1)
                    _ = fi27
                }
                _ = fi24
                _ = v25
            }
        }
        var v32 bool = true
        _ = v32
    }
    v5 = append(v5, 376)
    v33 := v4["foo bar"]
    for fi34 := 0; fi34 < 9; fi34++ {
        fmt.Println(fi34, v33, v33)
        _ = fi34
    }
    delete(v4, "")
    var v35 string = ""
    v33 -= 368
    for fi36 := 0; fi36 < 1; fi36++ {
        if false {
            v1 = (v1 - 1)
        }
        _ = len(v4)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
