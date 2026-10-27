package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0() int8 {
    if ("hello" != "") {
        var v5 int64 = 0
        var v6 int = int(v5)
        var v7 uint64 = 2
        _ = v5
        _ = v6
        _ = v7
    }
    if false {
        var v8 float64 = 1.5
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v9 uint8 = 128
            _ = v9
        case 4294967295:
            var v10 float32 = 3.14159
            _ = v10
        default:
            var v11 float64 = ((v8 * (float64(0.1) / 1.5)) * float64(3.14159))
            _ = v11
        }
        _ = v8
    } else {
        var v12 int32 = -26
        v13 := GenS1{F0: 2, F1: 65535}
        _ = v12
        _ = v13
    }
    if ((true && true) || (1 != 1)) {
        if (-1 <= -290) {
            if (0.0 < 0.1) {
                v14 := []int{127}
                v15 := GenS2{F0: 128}
                var v16 int16 = 0
                _ = v14
                _ = v15
                _ = v16
            }
            if (42 > -970) {
                var v17 uint8 = 38
                v18 := GenS1{F0: 267, F1: 2147483648}
                if (-2147483648 == 2147483647) {
                    v19 := GenS1{F0: 556, F1: 0}
                    var v20 int32 = 1
                    _ = v19
                    _ = v20
                }
                _ = v17
                _ = v18
            }
        }
        var v21 uint64 = 0
        _ = v21
    } else {
        v22 := []int{127, -2147483648, 4294967295, -1}
        _ = v22
    }
    return 100
}
func genFn1(p0 uint8) int64 {
    p0 += p0
    p0 = (p0 * p0)
    _ = p0
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = 42
    var v3 int64 = 1
    v4 := make(map[uint32]uint32)
    if false {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := []int{fi23, fi23}
            switch fi23 {
            case 0:
                for fi25 := 0; fi25 < 14; fi25++ {
                    switch 908 {
                    case 4294967295:
                        var v26 uint16 = 255
                        _ = v26
                    default:
                        v27 := GenS3{F0: (false && false)}
                        _ = v27
                    }
                    _ = fi25
                }
            case -9223372036854775808:
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 int16 = 42
                    fmt.Println(fi23, fi28)
                    if (100 == 100) {
                        v3 += v3
                        v3 -= 9223372036854775807
                        var v30 uint = 717
                        _ = v30
                    }
                    if true {
                        var v31 float32 = (((0.1 / 0.1) * 1.5) - 1e10)
                        var v32 uint = 255
                        _ = v31
                        _ = v32
                    } else {
                        v33 := GenS2{F0: 128}
                        switch fi28 {
                        case 610:
                            var v34 int8 = 127
                            _ = v34
                        case -606:
                            v3 += -9223372036854775808
                        }
                        fmt.Println(v2, v3, v29)
                        _ = v33
                    }
                    _ = fi28
                    _ = v29
                }
            default:
                v35 := GenS1{F0: 42, F1: 65535}
                _ = v35
            }
            v4[0] = 4294967295
            _ = fi23
            _ = v24
        }
        var v36 int32 = 42
        _ = v36
    } else {
        for fi37 := 0; fi37 < 1; fi37++ {
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS1{F0: 501, F1: 0}
                v40 := GenS4{F0: 0, F1: 65535}
                _ = fi38
                _ = v39
                _ = v40
            }
            if true {
                var v41 int8 = 127
                v1 = (^(v1))
                _ = v41
            }
            v2 = -806
            _ = fi37
        }
    }
    _ = len(v4)
    var v42 uint = 42
    var v43 int32 = 42
    for fi44 := 0; fi44 < 18; fi44++ {
        switch (fi44 & -1) {
        case 127:
            v45 := GenS2{F0: 127}
            _ = v45
        case 572:
            v42 = v42
        }
        fmt.Println(v2, v43, v42)
        _ = fi44
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := GenS1{F0: 2, F1: 441}
        fmt.Println(v3, v43)
        _ = fi46
        _ = v47
    }
    var v48 int16 = -32768
    v49 := []int{0}
    v50 := GenS2{F0: 128}
    fmt.Println(v3, v48, v2)
    v51 := GenS5{F0: 255, F1: 1}
    if false {
        for fi52 := 0; fi52 < 8; fi52++ {
            delete(v4, 0)
            var v53 uint = v42
            var v54 int = (fi52 >> ((2) & 63))
            _ = fi52
            _ = v53
            _ = v54
        }
    } else {
        v55 := GenS5{F0: 136, F1: -2147483648}
        if ("hello" != "a") {
            switch "" {
            case "a":
                var v56 int64 = v1
                _ = v56
            case "\n\t":
                var v57 uint = v42
                _ = v57
            default:
                v3 = v2
            }
        }
        v3 = -1
        _ = v55
    }
    v58 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v58
}
