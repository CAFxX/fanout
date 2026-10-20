package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int; F2 int32 }
type GenS7 struct { F0 string; F1 int64; F2 int64 }
type GenS8 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 float64 = (-2.25 - 1e10)
    v5 := make(map[string]bool)
    v6 := []int{1, -9223372036854775808, 221, -173}
    fmt.Println(v4)
    v7 := GenS1{F0: 0, F1: true}
    if ((249 != 340) && ("" > "a")) {
        v2 += (v2 % ((382) | 1))
        var v8 bool = false
        _ = v8
    } else {
        v9 := GenS2{F0: 94}
        switch -1 {
        case -9223372036854775808:
            if (!(!(1 > 2147483648))) {
                var v10 uint8 = 242
                fmt.Println(v1, v10, v10)
                _ = v10
            } else {
                var v11 int64 = 133
                v3[0] = false
                _ = v11
            }
        case 0:
            v2 = 42
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v4 += v4
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v3[fi12] = true
        _ = fi12
    }
    v1 = (((-2.25 / v1) * v1) * float32(float64(0.0)))
    switch "a" {
    case "foo bar":
        if false {
            v13 := GenS2{F0: 0}
            v4 += 0.1
            for fi14 := 0; fi14 < 6; fi14++ {
                switch -1 {
                case 0:
                    for fi15 := 0; fi15 < 5; fi15++ {
                        switch fi14 {
                        case 4294967295:
                            var v16 uint = 777
                            _ = v16
                        default:
                            v4 += v4
                            v2 = 1
                            v2 += 1
                        }
                        _ = len(v6)
                        fmt.Println(v1, v2, v1)
                        _ = fi15
                    }
                default:
                    v17 := GenS3{F0: 100, F1: 2}
                    _ = v17
                }
                if false {
                    v2 = v2
                    v2 = 4294967295
                    var v18 uint16 = uint16(v2)
                    _ = v18
                } else {
                    if false {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            v2 = (v2 >> ((42) & 63))
                            _ = fi19
                        }
                    }
                    v6 = append(v6, fi14)
                }
                v2 -= 42
                _ = fi14
            }
            _ = v13
        } else {
            switch -105 {
            case -494:
                v2 += 2
            }
            if (822 >= 1) {
                if (true || ((v2 >> ((18) & 63)) >= 1)) {
                    v6 = append(v6, 32767)
                }
                var v20 int16 = 396
                if (1.5 < v4) {
                    if false {
                        v20 = 0
                        var v21 uint64 = 428
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                        _ = v21
                    } else {
                        v20 = (v20 ^ 1000)
                    }
                    var v22 int = 42
                    if (!false) {
                        v23 := GenS1{F0: 18446744073709551615, F1: (false && true)}
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v4 -= v4
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi24
                        }
                        for fi25 := 0; fi25 < 12; fi25++ {
                            switch int(v20) {
                            case 32767:
                                if (-9223372036854775808 > 0) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    for fi26 := 0; fi26 < 16; fi26++ {
                                        if false {
                                            var v27 uint8 = 135
                                            var v28 uint8 = (^(v27))
                                            v22 = (v22 / (((^(fi25))) | 1))
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v20 += (v20 % ((-32768) | 1))
                                            if ((true && (2147483647 == 0)) && (v20 == v20)) {
                                                var v29 int64 = 0
                                                v30 := GenS1{F0: 824, F1: false}
                                                switch (fi26 * (^(fi26))) {
                                                case 42:
                                                    var v31 uint32 = 1
                                                    _ = v31
                                                case -91:
                                                    if false {
                                                        v32 := GenS4{F0: 0}
                                                        _ = v32
                                                    }
                                                }
                                                v29 = -1
                                                _ = v29
                                                _ = v30
                                            }
                                            v33 := GenS2{F0: 0}
                                            _ = v33
                                        }
                                        _ = fi26
                                    }
                                    v34 := GenS5{F0: ("\n\t" + "a")}
                                    _ = v34
                                } else {
                                    switch fi25 {
                                    case -1:
                                        var v35 int16 = (v20 >> ((607) & 63))
                                        _ = v35
                                    default:
                                        var v36 uint8 = 48
                                        _ = v36
                                    }
                                    v20 += (v20 * 1)
                                }
                            case 32768:
                                v37 := GenS6{F0: -1, F1: v22, F2: 2147483647}
                                _ = v37
                            case -387:
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    var v39 uint = 2
                                    v2 = v2
                                    var v40 uint = 2
                                    fmt.Println(v20, fi38, fi25)
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi25
                        }
                    }
                    _ = v22
                }
                _ = v20
            }
            v41 := GenS3{F0: 33, F1: 127}
            _ = v41
        }
    case "":
        var v42 uint8 = 128
        _ = v42
    default:
        v6 = append(v6, 2147483647)
    }
    if ((true || false) || false) {
        var v43 uint32 = 317
        _ = v43
    } else {
        var v44 int = -1
        for fi45 := 0; fi45 < 13; fi45++ {
            if len(v6) > 0 {
                v6[0] = 127
            }
            var v46 int32 = -1
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := GenS7{F0: "hello", F1: 538, F2: -839}
                v49 := v5[""]
                _ = fi47
                _ = v48
                _ = v49
            }
            _ = fi45
            _ = v46
        }
        fmt.Println(v4)
        _ = v44
    }
    v1 -= (v1 + (v1 + float32(0.1)))
    v50 := GenS8{F0: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
