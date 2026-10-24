package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint8 = 1
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v2 = 42
    v6 := GenS1{F0: 0}
    v2 = 161
    switch -161 {
    case 0:
        if false {
            for fi7 := 0; fi7 < 11; fi7++ {
                v1 = v1
                _ = fi7
            }
        }
    case -9223372036854775808:
        v8 := GenS2{F0: 0, F1: 1, F2: 0}
        _ = v8
    case -817:
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                var v10 int64 = v2
                switch "" {
                case "foo bar":
                    delete(v3, 65535)
                case "a":
                    var v11 uint32 = 2147483648
                    _ = v11
                case "hello":
                    var v12 float32 = 0.1
                    _ = v12
                }
                _ = v10
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 float64 = 1e10
                _ = fi13
                _ = v14
            }
            fmt.Println(v1)
            _ = fi9
        }
    default:
        v2 = v2
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint8 = v1
        switch "foo bar" {
        case "foo bar":
            v2 -= (v2 >> ((2) & 63))
        case "hello":
            v2 += v2
        case "":
            switch (-(fi15)) {
            case 32767:
                var v17 int8 = 100
                _ = v17
            case -2147483648:
                v18 := GenS2{F0: 42, F1: fi15, F2: -1}
                _ = v18
            default:
                if false {
                    fmt.Println(v16)
                    switch fi15 {
                    case 4294967295:
                        fmt.Println(v16)
                    case -590:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            v20 := GenS2{F0: 4294967295, F1: (fi15 << ((255) & 63)), F2: 42}
                            v21 := GenS2{F0: 2, F1: (fi15 / ((fi15) | 1)), F2: -47}
                            v22 := GenS3{F0: 1, F1: v2}
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                    case -9223372036854775808:
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 int = fi15
                            var v25 string = "foo bar"
                            v26 := []int{}
                            v24 -= 1
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                } else {
                    switch (("a" + "hello") + "a") {
                    case "\n\t":
                        if false {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                var v28 uint8 = (^(v16))
                                _ = fi27
                                _ = v28
                            }
                            var v29 uint32 = 65535
                            _ = v29
                        }
                    }
                    var v30 int8 = -16
                    v16 += (^(v16))
                    _ = v30
                }
            }
        default:
            for fi31 := 0; fi31 < 3; fi31++ {
                v16 -= 76
                _ = fi31
            }
        }
        var v32 uint32 = 2
        _ = fi15
        _ = v16
        _ = v32
    }
    if (42 == 32767) {
        if false {
            for fi33 := 0; fi33 < 2; fi33++ {
                if (0.0 != 1e10) {
                    var v34 uint8 = 68
                    v1 += v1
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v1 = 42
                        v36 := GenS2{F0: 1, F1: (fi35 << ((606) & 63)), F2: -1}
                        v37 := []int{-137, (^(fi33))}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                } else {
                    switch -665 {
                    case 128:
                        var v38 int32 = -1
                        _ = v38
                    case 42:
                        v39 := []int{(fi33 >> ((2) & 63)), 32768, -9223372036854775808, 0}
                        _ = v39
                    }
                }
                var v40 uint = 42
                v41 := v5["hello"]
                _ = fi33
                _ = v40
                _ = v41
            }
            _ = len(v4)
        }
        v1 -= (v1 / (((v1 >> ((42) & 63))) | 1))
        for fi42 := 0; fi42 < 5; fi42++ {
            _ = len(v3)
            var v43 uint8 = v1
            _ = fi42
            _ = v43
        }
    } else {
        v44 := GenS4{F0: 285}
        if true {
            switch ("hello" + "foo bar") {
            case "":
                v1 = v1
            case "hello":
                var v45 bool = true
                _ = v45
            default:
                for fi46 := 0; fi46 < 13; fi46++ {
                    var v47 int = 344
                    var v48 int32 = -2147483648
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            }
        }
        _ = v44
    }
    v1 -= (v1 / ((v1) | 1))
    v2 += -1
    var v49 int64 = 406
    for fi50 := 0; fi50 < 16; fi50++ {
        v1 -= (v1 ^ v1)
        _ = fi50
    }
    v1 += v1
    for fi51 := 0; fi51 < 11; fi51++ {
        if (true && (2 < 42)) {
            fmt.Println(v49, v2)
        }
        if (true || true) {
            v2 = v49
            v49 -= v2
            var v52 int32 = -1
            _ = v52
        } else {
            var v53 int32 = 647
            v49 = v2
            _ = v53
        }
        _ = fi51
    }
    if false {
        var v54 int32 = 42
        _ = v54
    } else {
        for fi55 := 0; fi55 < 4; fi55++ {
            for fi56 := 0; fi56 < 20; fi56++ {
                v57 := GenS2{F0: 4294967295, F1: fi56, F2: -1}
                fmt.Println(v2, v49, v2)
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
        v49 -= -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
