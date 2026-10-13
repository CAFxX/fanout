package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int16; F1 uint16; F2 int }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 227
    var v3 int = v1
    var v4 float64 = -2.25
    v5 := make(map[int]bool)
    var v6 int32 = 2147483647
    v2 += v2
    switch v3 {
    case 1:
        if ((v2 | v2) < (^(v2))) {
            var v7 uint16 = 2
            _ = v7
        } else {
            var v8 uint = 18446744073709551615
            for fi9 := 0; fi9 < 13; fi9++ {
                v6 += v6
                _ = fi9
            }
            _ = v8
        }
    case 130:
        if false {
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v4, fi10, v4)
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: -1, F1: (false && (!((1.5 * 1e10) > 0.1)))}
                    var v13 int8 = 100
                    v3 = (v1 << ((1) & 63))
                    if true {
                        if ((0.1 / 1.5) == float32(1e10)) {
                            fmt.Println(fi10, v2)
                            v14 := []int{fi10, fi10}
                            if (!false) {
                                v2 += (v2 << ((4294967295) & 63))
                            }
                            _ = v14
                        }
                        v13 += v13
                    }
                    var v15 uint64 = 4294967295
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v15
                }
                _ = fi10
            }
            v3 -= v1
            if (0 != 1000) {
                if false {
                    var v16 uint64 = 759
                    _ = v16
                } else {
                    v17 := GenS1{F0: -411, F1: false}
                    if ((!(false && (-89 >= 58))) && true) {
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v2 += (v2 * v2)
                            v3 = v3
                            v19 := GenS2{F0: 32767, F1: fi18, F2: 1}
                            _ = fi18
                            _ = v19
                        }
                        var v20 float64 = v4
                        fmt.Println(v20, v20, v6)
                        _ = v20
                    } else {
                        v3 = -551
                        fmt.Println(v2)
                    }
                    _ = v17
                }
                for fi21 := 0; fi21 < 7; fi21++ {
                    switch 539 {
                    case 42:
                        v2 = (^(v2))
                    case 1:
                        fmt.Println(fi21, v3, v3)
                    case 4294967295:
                        v1 = -1
                    }
                    _ = fi21
                }
            } else {
                v3 = v3
            }
        } else {
            v2 = v2
            for fi22 := 0; fi22 < 11; fi22++ {
                if ((-2.25 == float32(3.14159)) || true) {
                    var v23 uint = 255
                    v1 = 0
                    if (!true) {
                        v23 += (v23 ^ uint(v2))
                        for fi24 := 0; fi24 < 18; fi24++ {
                            switch fi22 {
                            case 0:
                                var v25 uint64 = 42
                                _ = v25
                            case 504:
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            case 32767:
                                switch "hello" {
                                case "foo bar":
                                    v4 = v4
                                case "\n\t":
                                    v4 = v4
                                }
                            }
                            v27 := GenS2{F0: 42, F1: v3, F2: 42}
                            _ = fi24
                            _ = v27
                        }
                        v28 := GenS2{F0: 65535, F1: (fi22 | -1), F2: 65535}
                        _ = v28
                    } else {
                        if ("\n\t" < "hello") {
                            var v29 int = 0
                            var v30 int = 4294967295
                            _ = v29
                            _ = v30
                        }
                        v31 := v5[(-(v1))]
                        _ = v31
                    }
                    _ = v23
                }
                _ = fi22
            }
            v32 := GenS3{F0: 1, F1: -1}
            _ = v32
        }
    }
    switch "" {
    case "hello":
        var v33 int64 = 42
        _ = v33
    default:
        if true {
            fmt.Println(v4)
        }
    }
    var v34 int16 = 42
    if false {
        var v35 string = "\n\t"
        _ = v35
    }
    switch (v1 % (((v3 - (v1 << ((265) & 63)))) | 1)) {
    case 2147483647:
        switch -2147483648 {
        case -662:
            var v36 uint32 = 2
            _ = v36
        default:
            v34 += (^(v34))
        }
    case 128:
        if (v34 != v34) {
            v34 = (v34 % ((-32768) | 1))
            var v37 int8 = 100
            var v38 uint64 = 2
            _ = v37
            _ = v38
        } else {
            fmt.Println(v2)
            fmt.Println(v1, v4, v6)
        }
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            v40 := []int{(fi39 + fi39), fi39, fi39}
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS4{F0: (v34 - (v34 << ((2) & 63))), F1: 2, F2: (-(v1))}
    v42 := GenS2{F0: 32767, F1: (v1 << ((369) & 63)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
