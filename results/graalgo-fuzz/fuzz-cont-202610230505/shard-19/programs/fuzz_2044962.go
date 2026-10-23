package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 bool) uint8 {
    if (42 != -1) {
        v6 := []int{-1, -310}
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int8 = 0
            var v9 float32 = (1e10 - 1e10)
            var v10 int64 = 9223372036854775807
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
    } else {
        var v11 uint32 = 720
        v12 := GenS1{F0: 100}
        v13 := GenS1{F0: 127}
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    var v3 uint64 = 867
    var v4 uint16 = 32767
    var v5 int32 = 1
    fmt.Println(v5, v5, v3)
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            v17 := GenS2{F0: fi16, F1: fi15}
            v2 = "foo bar"
            _ = fi16
            _ = v17
        }
        switch fi15 {
        case -9223372036854775808:
            if false {
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 uint = 65535
                    v20 := []int{fi15, fi15, fi18}
                    if false {
                        if (v2 == v2) {
                            var v21 uint16 = (v4 + (v4 + 2))
                            v22 := v1[1]
                            v2 = v2
                            _ = v21
                            _ = v22
                        } else {
                            v3 += v3
                            var v23 uint64 = (v3 / ((v3) | 1))
                            _ = v23
                        }
                        v24 := GenS3{F0: (^(v4)), F1: -2147483648}
                        _ = v24
                    }
                    var v25 int16 = -81
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v25
                }
                switch v2 {
                case "hello":
                    var v26 uint16 = (v4 >> ((65535) & 63))
                    _ = v26
                case "":
                    v27 := GenS3{F0: (v4 | 605), F1: v5}
                    _ = v27
                }
                v28 := GenS4{F0: 723, F1: 2, F2: v5}
                _ = v28
            } else {
                v29 := GenS2{F0: fi15, F1: (fi15 & (fi15 & int(v3)))}
                switch v2 {
                case "":
                    v30 := GenS4{F0: -9223372036854775808, F1: 382, F2: -1}
                    _ = v30
                case "\n\t":
                    v4 -= (v4 - 2)
                }
                v5 = (v5 << ((18446744073709551615) & 63))
                _ = v29
            }
        case 609:
            fmt.Println(v3, v4)
        case 1:
            fmt.Println(v5)
        default:
            switch v2 {
            case "hello":
                v2 = v2
            }
        }
        if false {
            var v31 string = (v2 + ("hello" + (("foo bar" + v2) + "hello")))
            _ = v31
        } else {
            fmt.Println(v2, fi15, v2)
            v2 = v2
        }
        _ = fi15
    }
    v3 += (^(v3))
    var v32 uint = 0
    v1[210] = v2
    v1[65535] = v2
    var v33 float32 = (1e10 + 3.14159)
    for fi34 := 0; fi34 < 10; fi34++ {
        v33 -= v33
        _ = fi34
    }
    var v35 int8 = -58
    var v36 int8 = v35
    var v37 float32 = 3.14159
    switch -1 {
    case 32768:
        if true {
            var v38 uint16 = v4
            _ = v38
        }
    case -747:
        v33 = v33
    case 2147483647:
        if false {
            if (!true) {
                v35 -= v36
                var v39 int64 = 42
                _ = v39
            }
            v40 := GenS5{F0: 2, F1: -1}
            _ = v40
        } else {
            var v41 int16 = 490
            var v42 uint8 = 0
            _ = v41
            _ = v42
        }
    default:
        fmt.Println(v5, v33, v35)
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
