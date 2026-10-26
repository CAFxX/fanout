package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint = 255
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 uint64 = 224
    var v6 int = 519
    v7 := v3[v2]
    delete(v3, v2)
    if false {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := v3[v2]
            _ = fi8
            _ = v9
        }
    } else {
        if (-2147483648 != -1) {
            v5 = v5
        } else {
            fmt.Println(v7, v2)
        }
        v10 := []int{}
        _ = v10
    }
    if (275 > 1) {
        if (1 >= int32(v5)) {
            var v11 int16 = -787
            _ = v11
        } else {
            v12 := v3[uint32(v5)]
            if true {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS1{F0: 286, F1: v1, F2: v5}
                    _ = fi13
                    _ = v14
                }
                var v15 bool = true
                if v15 {
                    var v16 int64 = int64(v1)
                    if v15 {
                        if v15 {
                            fmt.Println(v2, v6)
                            if true {
                                var v17 bool = (!false)
                                _ = v17
                            }
                            v15 = ((1e10 / ((1.5 - 1e10) - 0.1)) < 1.5)
                        }
                        v2 -= v2
                        for fi18 := 0; fi18 < 14; fi18++ {
                            for fi19 := 0; fi19 < 7; fi19++ {
                                v20 := GenS1{F0: 1, F1: v1, F2: (v5 << (((v1 | v1)) & 63))}
                                switch 643 {
                                case 0:
                                    v21 := GenS2{F0: (v1 - v1), F1: ""}
                                    _ = v21
                                }
                                _ = fi19
                                _ = v20
                            }
                            for fi22 := 0; fi22 < 2; fi22++ {
                                v3[uint32(v1)] = (v7 >> ((524) & 63))
                                var v23 float64 = 1.5
                                v7 = v12
                                v12 = 1
                                v5 = (v5 | 42)
                                _ = fi22
                                _ = v23
                            }
                            _ = fi18
                        }
                    }
                    v5 = v5
                    _ = v16
                }
                _ = v15
            } else {
                _ = len(v3)
                v24 := v3[(v2 / ((v2) | 1))]
                _ = v24
            }
            var v25 uint8 = 75
            _ = v12
            _ = v25
        }
        v6 -= 4294967295
        for fi26 := 0; fi26 < 17; fi26++ {
            switch ("a" + "\n\t") {
            case "a":
                for fi27 := 0; fi27 < 17; fi27++ {
                    fmt.Println(v5)
                    _ = fi27
                }
            default:
                v5 += v5
            }
            _ = fi26
        }
    } else {
        v2 = uint32(v5)
    }
    v6 = 4294967295
    switch v6 {
    case 0:
        v28 := v3[(^(v7))]
        _ = v28
    case 127:
        var v29 uint8 = 29
        _ = v29
    default:
        if true {
            var v30 uint16 = 255
            fmt.Println(v5)
            for fi31 := 0; fi31 < 11; fi31++ {
                v32 := []int{(fi31 % ((42) | 1)), fi31}
                _ = fi31
                _ = v32
            }
            _ = v30
        }
    }
    v7 = (^(v2))
    if false {
        var v33 uint = v1
        fmt.Println(v7, v7, v5)
        _ = v33
    } else {
        v2 -= uint32(v5)
    }
    v7 = v2
    for fi34 := 0; fi34 < 20; fi34++ {
        if true {
            for fi35 := 0; fi35 < 6; fi35++ {
                switch "" {
                case "foo bar":
                    v36 := []int{(v6 / ((int(v5)) | 1))}
                    _ = v36
                default:
                    v1 = v1
                }
                _ = fi35
            }
        } else {
            var v37 bool = (-2.25 == ((3.14159 * (3.14159 * (3.14159 * 0.1))) + (((1e10 / 1.0) * -2.25) - 1.5)))
            var v38 bool = true
            var v39 int32 = 0
            _ = v37
            _ = v38
            _ = v39
        }
        var v40 float64 = -2.25
        _ = fi34
        _ = v40
    }
    v5 = (^(v5))
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 18; fi42++ {
            if (true || (v1 <= v1)) {
                v2 -= v2
                v3[v2] = (v2 | v2)
            }
            _ = len(v3)
            v43 := GenS3{F0: 9223372036854775807}
            _ = fi42
            _ = v43
        }
        v44 := []int{}
        _ = fi41
        _ = v44
    }
    if false {
        var v45 uint64 = (v5 / ((v5) | 1))
        v46 := GenS4{F0: 255}
        v47 := GenS4{F0: 2}
        _ = v45
        _ = v46
        _ = v47
    } else {
        var v48 float32 = 1.5
        _ = v48
    }
    var v49 uint16 = 65535
    for fi50 := 0; fi50 < 6; fi50++ {
        if true {
            var v51 float64 = 3.14159
            _ = v51
        } else {
            v7 += (v7 * (v7 ^ 821))
            var v52 uint64 = (v5 * (v5 & (v5 | 2)))
            _ = v52
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
