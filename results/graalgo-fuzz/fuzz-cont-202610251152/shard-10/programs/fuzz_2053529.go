package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{32767}
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 int8 = 1
    if (521 >= 410) {
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 int16 = 42
            v4 = (v4 * v4)
            var v7 uint8 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
        v4 = (^(v4))
        for fi8 := 0; fi8 < 4; fi8++ {
            _ = len(v3)
            fmt.Println(v4)
            if true {
                v9 := GenS1{F0: -118, F1: 42, F2: -9223372036854775808}
                v10 := GenS1{F0: 0, F1: 127, F2: -224}
                _ = v10
                _ = v9
            } else {
                v4 = (v4 % ((-128) | 1))
                if (((((1e10 / 0.1) / 1e10) * 3.14159) - ((0.0 * 0.1) / (1.5 / 3.14159))) < 0.0) {
                    if true {
                        var v11 bool = false
                        _ = v11
                    } else {
                        v4 = (v4 >> ((2) & 63))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi12 := 0; fi12 < 18; fi12++ {
                            if true {
                                v4 += v4
                                var v13 int = (^(fi8))
                                _ = v13
                            } else {
                                var v14 string = "foo bar"
                                var v15 int16 = 1
                                _ = v14
                                _ = v15
                            }
                            v4 = -1
                            _ = fi12
                        }
                    }
                    switch 32768 {
                    case 2147483647:
                        var v16 string = "a"
                        _ = v16
                    case 128:
                        fmt.Println(fi8, v4)
                    default:
                        v4 = v4
                    }
                }
                var v17 bool = false
                _ = v17
            }
            _ = fi8
        }
    } else {
        v4 += (v4 * v4)
    }
    switch "" {
    case "\n\t":
        v4 += v4
    default:
        fmt.Println(v4)
    }
    if (false || (true || (v4 > 34))) {
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 uint64 = 42
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        var v21 string = "foo bar"
        _ = v21
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            v4 += v4
            var v23 int = fi22
            _ = fi22
            _ = v23
        }
        delete(v3, (("a" + "hello") + "hello"))
    }
    if false {
        v4 += 1
        v4 = (v4 & v4)
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v4, v4)
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: 875, F1: 41, F2: -1}
            _ = fi25
            _ = v26
        }
        switch (fi24 * fi24) {
        case -2147483648:
            var v27 int32 = 0
            _ = v27
        case 0:
            _ = len(v1)
        case -9223372036854775808:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi24
    }
    if true {
        if (255 < 127) {
            v28 := GenS2{F0: -909}
            _ = len(v1)
            _ = v28
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 int64 = 1
            _ = fi29
            _ = v30
        }
        v4 += (v4 ^ v4)
    } else {
        delete(v2, ("\n\t" + "foo bar"))
        fmt.Println(v4)
    }
    if (false && (-1 > 9223372036854775807)) {
        v4 = v4
        for fi31 := 0; fi31 < 13; fi31++ {
            fmt.Println(v4, fi31)
            v32 := GenS2{F0: 992}
            if (1 <= 623) {
                var v33 float64 = 0.1
                var v34 int = (fi31 / (((fi31 >> ((42) & 63))) | 1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
    } else {
        switch "a" {
        case "\n\t":
            v35 := v2["foo bar"]
            _ = v35
        }
        var v36 string = "hello"
        v4 -= (v4 << ((2) & 63))
        _ = v36
    }
    if (-155 <= 42) {
        for fi37 := 0; fi37 < 13; fi37++ {
            if (4294967295 > 781) {
                v4 -= v4
            } else {
                var v38 bool = ("\n\t" != "")
                var v39 uint16 = 1
                v38 = (!v38)
                _ = v38
                _ = v39
            }
            _ = fi37
        }
        if len(v1) > 0 {
            v1[0] = 346
        }
        var v40 string = "\n\t"
        _ = v40
    }
    var v41 int64 = 42
    var v42 int = 642
    v43 := v3["\n\t"]
    v43 = "a"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
