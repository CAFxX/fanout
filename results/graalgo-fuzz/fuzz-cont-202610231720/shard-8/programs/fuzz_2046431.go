package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 uint16, p1 bool, p2 uint) int16 {
    var v6 int32 = -565
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((-(fi7)) >= 0) {
                fmt.Println(p0)
                v9 := GenS1{F0: 1000, F1: "hello", F2: 845}
                _ = v9
            } else {
                var v10 int64 = 143
                fmt.Println(p0)
                fmt.Println(p2, p0, v10)
                _ = v10
            }
            var v11 int16 = 610
            _ = fi8
            _ = v11
        }
        _ = fi7
    }
    if p1 {
        p1 = (p1 || (true && ((!(false && (0.1 == 3.14159))) && true)))
        if (((3.14159 - 3.14159) + (float64(3.14159) + 0.1)) >= (((1.5 * 0.1) * (-2.25 - 0.1)) / -2.25)) {
            p0 = (p0 * p0)
        } else {
            fmt.Println(p0, p1, p1)
        }
        v12 := GenS1{F0: -1, F1: ((("" + "") + "hello") + "hello"), F2: 0}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            switch fi13 {
            case -397:
                v6 += (v6 / ((-1) | 1))
            case -251:
                fmt.Println(p2)
            }
            p1 = (1 <= 1)
            v14 := []int{fi13, 615}
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := []int{fi15, (fi15 + -857), 127}
            fmt.Println(v6, fi15)
            _ = fi15
            _ = v16
        }
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 bool = true
    v2 := []uint32{2147483648, 711, 2147483648}
    var v3 int64 = 9223372036854775807
    var v4 uint32 = 985
    var v5 uint = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v4, v1, v4)
    v2 = append(v2, uint32(v5))
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS2{F0: 1, F1: v1, F2: 2}
        if v1 {
            v1 = v1
            switch "a" {
            case "hello":
                switch 337 {
                case 42:
                    v19 := GenS3{F0: 35}
                    _ = v19
                case 2147483647:
                    fmt.Println(v4, v5, v4)
                case 909:
                    v20 := GenS3{F0: (fi17 | fi17)}
                    _ = v20
                default:
                    if (("hello" + ("hello" + "hello")) >= ((("\n\t" + "foo bar") + "\n\t") + ("a" + ("foo bar" + "a")))) {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 uint16 = uint16(v5)
                            v3 = 147
                            v3 = (v3 - v3)
                            v4 += 65535
                            _ = fi21
                            _ = v22
                        }
                    } else {
                        switch 9223372036854775807 {
                        case 32767:
                            fmt.Println(v3, fi17, v5)
                        case 883:
                            switch "" {
                            case "":
                                for fi23 := 0; fi23 < 8; fi23++ {
                                    _ = len(v2)
                                    var v24 uint64 = 18446744073709551615
                                    _ = fi23
                                    _ = v24
                                }
                            default:
                                v25 := GenS3{F0: -2147483648}
                                _ = v25
                            }
                        case 4294967295:
                            var v26 int32 = 42
                            _ = v26
                        }
                    }
                }
            case "a":
                for fi27 := 0; fi27 < 2; fi27++ {
                    if (1e10 <= 1e10) {
                        if (-67 <= -123) {
                            var v28 uint32 = (v4 / (((^(v4))) | 1))
                            _ = len(v2)
                            _ = v28
                        }
                        fmt.Println(v3, v4)
                        switch fi27 {
                        case -554:
                            v2 = append(v2, v4)
                        case -9223372036854775808:
                            var v29 int64 = (v3 << ((927) & 63))
                            _ = v29
                        default:
                            _ = len(v2)
                        }
                    } else {
                        v3 = v3
                    }
                    var v30 uint64 = 2
                    _ = fi27
                    _ = v30
                }
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    var v31 uint64 = 966
                    _ = v31
                case "":
                    v32 := GenS4{F0: fi17, F1: 0, F2: "\n\t"}
                    _ = v32
                }
            }
        } else {
            v3 -= v3
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            switch "\n\t" {
            case "\n\t":
                var v34 int32 = 496
                _ = v34
            }
            v35 := GenS4{F0: fi17, F1: 519, F2: "\n\t"}
            _ = fi33
            _ = v35
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v3)
    v36 := GenS2{F0: -967, F1: (!(v1 && (42 <= 2))), F2: 1}
    v2 = append(v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
