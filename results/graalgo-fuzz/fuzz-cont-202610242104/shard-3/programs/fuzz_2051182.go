package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int = 891
    v2 := []string{}
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.0
    if false {
        v5 := GenS1{F0: 32767}
        if false {
            v1 = (v1 * 2147483647)
            switch (-(v1)) {
            case -9223372036854775808:
                switch (-(v1)) {
                case 832:
                    if (0 > 42) {
                        var v6 uint16 = 646
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v6
                    } else {
                        var v7 string = ("a" + "hello")
                        _ = v7
                    }
                case -228:
                    var v8 int32 = -870
                    _ = v8
                }
            default:
                var v9 float64 = 3.14159
                _ = v9
            }
            v4 = 1.5
        } else {
            var v10 bool = (0 > -406)
            v11 := GenS2{F0: -23, F1: -128, F2: 9223372036854775807}
            v1 = v1
            _ = v10
            _ = v11
        }
        v4 += v4
        _ = v5
    } else {
        var v12 uint64 = 129
        if (562 >= 65535) {
            var v13 int16 = 42
            var v14 bool = true
            v13 += v13
            _ = v13
            _ = v14
        } else {
            var v15 uint64 = (^(v12))
            v16 := GenS2{F0: -128, F1: 0, F2: 1}
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(v4, v4, fi17)
            var v18 int = (v1 & (v1 / ((-301) | 1)))
            _ = fi17
            _ = v18
        }
        _ = v12
    }
    var v19 int16 = 0
    for fi20 := 0; fi20 < 20; fi20++ {
        switch "a" {
        case "foo bar":
            v21 := GenS3{F0: -128, F1: 239}
            _ = v21
        case "hello":
            if true {
                v3[2] = ("a" > ("\n\t" + "hello"))
            } else {
                if true {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "")
                            }
                            switch 1 {
                            case -1:
                                var v23 string = ("hello" + "foo bar")
                                _ = v23
                            case 127:
                                fmt.Println(fi22, fi22)
                            }
                        }
                        v19 = v19
                        var v24 uint32 = 42
                        delete(v3, v24)
                        v2 = append(v2, "\n\t")
                        v19 += 32767
                        _ = fi22
                        _ = v24
                    }
                } else {
                    v1 -= (-(fi20))
                }
                for fi25 := 0; fi25 < 9; fi25++ {
                    v19 -= 271
                    _ = fi25
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint64 = 89
                    delete(v3, 4294967295)
                    var v28 uint = 255
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        default:
            v19 += (^(v19))
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            switch "a" {
            case "hello":
                v30 := GenS3{F0: 2, F1: -9223372036854775808}
                _ = v30
            }
            _ = fi29
        }
        v4 -= (v4 * v4)
        _ = fi20
    }
    var v31 int32 = 0
    var v32 float64 = (v4 - v4)
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(v19)
        v34 := GenS1{F0: (v19 - (v19 - (^(v19))))}
        _ = fi33
        _ = v34
    }
    if ((true && true) && (243 <= 1)) {
        switch 32767 {
        case -1:
            fmt.Println(v31)
        }
    }
    v35 := GenS4{F0: 2, F1: 144}
    var v36 int16 = (v19 ^ (v19 - -1))
    for fi37 := 0; fi37 < 4; fi37++ {
        v1 += (^(v1))
        _ = fi37
    }
    switch "" {
    case "":
        v19 = v19
    }
    v38 := GenS4{F0: 65535, F1: 127}
    var v39 bool = ((v1 << ((112) & 63)) > (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
}
