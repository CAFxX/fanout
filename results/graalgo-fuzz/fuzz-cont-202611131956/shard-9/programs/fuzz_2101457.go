package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 int) int64 {
    p0 = p0
    var v4 string = "hello"
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int32 = -2147483648
        _ = fi5
        _ = v6
    }
    var v7 int = p0
    _ = p0
    _ = v4
    _ = v7
    return -56
}

func main() {
    var v1 int32 = 973
    var v2 uint64 = 379
    var v3 uint8 = 43
    v2 = (v2 % ((2) | 1))
    for fi8 := 0; fi8 < 17; fi8++ {
        switch 2147483647 {
        case -557:
            var v9 uint = 4294967295
            _ = v9
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int = int(v1)
        var v12 uint64 = uint64(v3)
        _ = fi10
        _ = v11
        _ = v12
    }
    v1 += (^(v1))
    v2 -= uint64(v3)
    var v13 bool = (false && (false && false))
    v14 := GenS1{F0: -382, F1: 1, F2: 42}
    if v13 {
        if (!false) {
            var v15 float32 = 1e10
            v2 += 303
            _ = v15
        }
        var v16 float64 = (1e10 / 0.1)
        _ = v16
    } else {
        v2 -= 154
    }
    var v17 uint32 = 589
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 2; fi19++ {
            v1 -= v1
            switch 32767 {
            case 1:
                var v20 string = "foo bar"
                _ = v20
            case 42:
                v21 := GenS2{F0: -1}
                _ = v21
            case -2147483648:
                if (v17 != v17) {
                    v2 = v2
                    fmt.Println(v1, v13, v13)
                    v22 := GenS1{F0: (v1 ^ 2147483647), F1: 1, F2: 32767}
                    _ = v22
                } else {
                    var v23 bool = v13
                    v2 -= 4294967295
                    var v24 int16 = 244
                    _ = v23
                    _ = v24
                }
            default:
                var v25 uint8 = 238
                _ = v25
            }
            var v26 bool = true
            var v27 uint = 173
            _ = fi19
            _ = v26
            _ = v27
        }
        _ = fi18
    }
    if v13 {
        v28 := GenS2{F0: (v1 % ((0) | 1))}
        var v29 int8 = 42
        if v13 {
            v30 := []int{}
            _ = v30
        } else {
            v29 += 1
            v31 := GenS1{F0: (v1 - (-(v1))), F1: genFn0(-2147483648), F2: 1}
            _ = v31
        }
        _ = v28
        _ = v29
    } else {
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 uint16 = 0
            _ = fi32
            _ = v33
        }
        v34 := GenS2{F0: (v1 | int32(v17))}
        for fi35 := 0; fi35 < 7; fi35++ {
            switch "foo bar" {
            case "hello":
                v36 := GenS1{F0: int32(v17), F1: 1, F2: 255}
                _ = v36
            }
            v17 += v17
            v37 := GenS1{F0: (v1 & v1), F1: int64(v3), F2: 0}
            _ = fi35
            _ = v37
        }
        _ = v34
    }
    if false {
        v2 += v2
        if (!(!false)) {
            var v38 float64 = 3.14159
            v39 := GenS2{F0: (v1 / (((v1 | (v1 - 522))) | 1))}
            _ = v38
            _ = v39
        }
    } else {
        var v40 uint8 = (^(v3))
        _ = v40
    }
    var v41 float32 = 1e10
    for fi42 := 0; fi42 < 17; fi42++ {
        v43 := []int{}
        switch (^(fi42)) {
        case 881:
            v17 = (v17 >> ((0) & 63))
        case -401:
            for fi44 := 0; fi44 < 18; fi44++ {
                v41 += 0.0
                switch (fi44 + (-(fi42))) {
                case -994:
                    if false {
                        _ = len(v43)
                    } else {
                        v41 = v41
                    }
                case 32768:
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                }
                _ = fi44
            }
        case 128:
            v45 := GenS2{F0: int32(v3)}
            _ = v45
        }
        var v46 int32 = v1
        _ = fi42
        _ = v43
        _ = v46
    }
    switch 794 {
    case 2147483647:
        v41 += float32(1.5)
    }
    v47 := GenS3{F0: 127, F1: ((v3 > v3) || (659 <= 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v41
    _ = v47
}
