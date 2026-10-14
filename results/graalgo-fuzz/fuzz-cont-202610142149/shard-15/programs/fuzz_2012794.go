package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint, p2 uint16) int64 {
    v7 := GenS1{F0: 1000, F1: int16(p2), F2: 0}
    p0 = p0
    v8 := []int{2147483647, -2147483648, 9223372036854775807}
    var v9 float64 = float64(0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []string{}
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 float32 = (3.14159 / 1.0)
    v6 := make(map[uint32]string)
    var v10 string = "foo bar"
    if ((-2.25 >= 1.5) || (0 <= 0)) {
        var v11 int64 = 0
        switch -1 {
        case -372:
            if (true || true) {
                if true {
                    _ = len(v6)
                    v5 -= ((v5 * v5) * v5)
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        for fi13 := 0; fi13 < 13; fi13++ {
                            v14 := GenS1{F0: 32767, F1: -1, F2: 34}
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(v11)
                        _ = fi12
                    }
                }
                switch (v10 + (v10 + v10)) {
                case "hello":
                    var v15 int16 = int16(v11)
                    _ = v15
                case "\n\t":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 float32 = v5
                        v18 := GenS1{F0: 1, F1: 42, F2: 198}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                case "foo bar":
                    var v19 uint64 = 887
                    _ = v19
                }
                fmt.Println(v10)
            }
        case 9223372036854775807:
            fmt.Println(v5, v10)
        case 32768:
            v20 := GenS1{F0: 0, F1: 42, F2: 255}
            _ = v20
        }
        _ = v11
    }
    v5 = ((float32(-2.25) - v5) - (v5 * v5))
    var v21 int32 = -592
    if true {
        var v22 int32 = v21
        switch int(v21) {
        case 9223372036854775807:
            if false {
                var v23 bool = true
                if v23 {
                    v1 = append(v1, (v10 + v10))
                } else {
                    var v24 int8 = 42
                    v5 += (float32(1e10) / v5)
                    _ = v24
                }
                _ = v23
            } else {
                if (42 > 32768) {
                    switch 477 {
                    case 4294967295:
                        var v25 int16 = 42
                        _ = v25
                    case 42:
                        for fi26 := 0; fi26 < 5; fi26++ {
                            var v27 int = 127
                            _ = fi26
                            _ = v27
                        }
                    default:
                        for fi28 := 0; fi28 < 13; fi28++ {
                            fmt.Println(v10, v22, v21)
                            _ = fi28
                        }
                    }
                    var v29 uint8 = 127
                    v10 = v10
                    _ = v29
                } else {
                    fmt.Println(v22)
                }
                switch 42 {
                case 581:
                    fmt.Println(v10)
                case -367:
                    var v30 int8 = 42
                    _ = v30
                default:
                    v31 := v2[((v10 + v10) + v10)]
                    _ = v31
                }
                v2[("foo bar" + ((((v10 + ("\n\t" + "")) + v10) + "a") + v10))] = 4294967295
            }
        case -9223372036854775808:
            if true {
                switch v10 {
                case "\n\t":
                    v32 := GenS2{F0: 2}
                    _ = v32
                case "foo bar":
                    v10 = v10
                case "":
                    v5 = v5
                }
            }
        default:
            var v33 int = -9223372036854775808
            _ = v33
        }
        var v34 uint16 = 42
        _ = v22
        _ = v34
    } else {
        if false {
            fmt.Println(v10, v5)
            var v35 uint16 = 250
            v21 = (-(v21))
            _ = v35
        } else {
            if (387 == 2147483647) {
                switch -423 {
                case -545:
                    var v36 int32 = (v21 - -178)
                    _ = v36
                }
                if false {
                    for fi37 := 0; fi37 < 13; fi37++ {
                        v38 := GenS3{F0: ((372 > 65535) && true)}
                        fmt.Println(v10)
                        _ = fi37
                        _ = v38
                    }
                    var v39 int = -387
                    v40 := GenS2{F0: 194}
                    _ = v39
                    _ = v40
                } else {
                    v6[291] = ("hello" + ("" + ("\n\t" + "\n\t")))
                }
                v5 += -2.25
            } else {
                var v41 float32 = (v5 - v5)
                var v42 uint32 = 964
                _ = v41
                _ = v42
            }
        }
        v10 = (v10 + "\n\t")
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = v10
    }
    switch -1 {
    case -173:
        delete(v3, -2147483648)
    case -1:
        v43 := v3[128]
        _ = v43
    case 32768:
        var v44 float64 = ((0.1 + (-2.25 / float64(v5))) + 1e10)
        _ = v44
    default:
        var v45 uint32 = 65535
        _ = v45
    }
    if (212 >= -1) {
        switch "" {
        case "":
            if false {
                if (!false) {
                    for fi46 := 0; fi46 < 11; fi46++ {
                        var v47 uint = 1
                        _ = fi46
                        _ = v47
                    }
                } else {
                    var v48 float64 = (0.0 / 1.0)
                    _ = v48
                }
            }
        case "hello":
            var v49 int16 = -1
            _ = v49
        case "\n\t":
            if false {
                v50 := GenS1{F0: 1, F1: -599, F2: 42}
                v51 := GenS1{F0: 1000, F1: 786, F2: 255}
                _ = v50
                _ = v51
            }
        default:
            switch 2147483647 {
            case -287:
                var v52 uint16 = 111
                _ = v52
            case -218:
                if true {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        var v54 int = fi53
                        fmt.Println(fi53)
                        _ = len(v3)
                        _ = fi53
                        _ = v54
                    }
                }
            case -206:
                var v55 int16 = 1
                _ = v55
            default:
                switch v10 {
                case "":
                    if false {
                        var v56 uint = 438
                        _ = v56
                    }
                }
            }
        }
        v57 := GenS2{F0: 0}
        v58 := v6[2]
        _ = v57
        _ = v58
    } else {
        v5 = v5
    }
    var v59 float32 = float32((-2.25 * (0.1 / (1.5 / float64(float32((((1e10 / (0.0 + 1e10)) / -2.25) + 3.14159)))))))
    v59 += 0.1
    v10 = "hello"
    v21 = -563
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
