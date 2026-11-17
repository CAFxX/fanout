package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint8 }
func genFn0() uint16 {
    var v4 int8 = -128
    var v5 int64 = -9223372036854775808
    var v6 uint32 = 871
    v6 += (^(v6))
    _ = v4
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int64 = -1
    var v2 bool = false
    var v3 int32 = 0
    if (29 < (v1 + 9223372036854775807)) {
        for fi7 := 0; fi7 < 11; fi7++ {
            v2 = (!(42 <= 850))
            if (2147483647 == 2147483648) {
                v1 = -1
            } else {
                fmt.Println(v1, v2, v2)
                v8 := GenS1{F0: ("a" + "")}
                _ = v8
            }
            _ = fi7
        }
        var v9 int64 = v1
        v10 := []int{2147483647, 9223372036854775807, 4294967295, 865}
        _ = v10
        _ = v9
    } else {
        var v11 string = ""
        _ = v11
    }
    switch -2147483648 {
    case 1:
        var v12 float32 = 1e10
        _ = v12
    default:
        v2 = v2
    }
    v13 := GenS1{F0: "a"}
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(fi14)
        if false {
            fmt.Println(v3, v2)
        }
        _ = fi14
    }
    switch 699 {
    case -9223372036854775808:
        v3 = 845
    case 149:
        v1 = int64(v3)
    case -2147483648:
        if v2 {
            switch 42 {
            case -2147483648:
                v3 += (v3 >> ((2) & 63))
            case 4294967295:
                if v2 {
                    var v15 int = -1
                    _ = v15
                } else {
                    switch 32768 {
                    case 127:
                        v16 := []int{}
                        _ = v16
                    }
                    v2 = v2
                    fmt.Println(v1, v1)
                }
            default:
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int = fi17
                    _ = fi17
                    _ = v18
                }
            }
        }
    default:
        fmt.Println(v3, v2, v1)
    }
    var v19 int = 42
    v20 := []int{v19}
    var v21 uint16 = genFn0()
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS1{F0: "foo bar"}
        v24 := GenS2{F0: 1000, F1: (^(v21)), F2: 127}
        v20 = append(v20, -2147483648)
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            _ = len(v20)
            v1 = v1
            _ = fi26
        }
        _ = fi25
    }
    fmt.Println(v3, v2, v2)
    switch (-(v19)) {
    case -9223372036854775808:
        var v27 float64 = (float64(0.1) / float64((((0.1 + (1.5 - 3.14159)) * 0.0) - -2.25)))
        _ = v27
    case 4294967295:
        for fi28 := 0; fi28 < 8; fi28++ {
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi28
        }
    case 0:
        var v29 uint64 = 42
        _ = v29
    default:
        if ((v3 + v3) > v3) {
            var v30 uint32 = 2
            v1 = v1
            for fi31 := 0; fi31 < 12; fi31++ {
                var v32 uint16 = 65535
                for fi33 := 0; fi33 < 4; fi33++ {
                    switch (^(fi31)) {
                    case 32768:
                        if v2 {
                            var v34 uint = 255
                            v35 := GenS1{F0: ("foo bar" + ("a" + ""))}
                            _ = len(v20)
                            v36 := GenS3{F0: v1, F1: 2}
                            switch (v19 & fi31) {
                            case 9223372036854775807:
                                switch (fi33 % (((fi33 & int(v34))) | 1)) {
                                case -112:
                                    fmt.Println(v1)
                                }
                            }
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    case 0:
                        v2 = (!(v2 && v2))
                    default:
                        fmt.Println(v1, fi31)
                    }
                    _ = fi33
                }
                if len(v20) > 0 {
                    v20[0] = (-(v19))
                }
                _ = fi31
                _ = v32
            }
            _ = v30
        } else {
            v2 = v2
            switch "a" {
            case "a":
                v37 := GenS4{F0: -1}
                _ = v37
            case "\n\t":
                if v2 {
                    var v38 int32 = 537
                    v39 := GenS5{F0: "foo bar", F1: 177}
                    var v40 uint16 = (v21 + (v21 * (v21 << ((679) & 63))))
                    _ = v38
                    _ = v39
                    _ = v40
                } else {
                    _ = len(v20)
                    for fi41 := 0; fi41 < 1; fi41++ {
                        v42 := GenS6{F0: (v3 << ((1) & 63)), F1: (^(v21)), F2: 28}
                        var v43 uint64 = 18446744073709551615
                        v20 = append(v20, 628)
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                }
            default:
                v1 = (v1 / (((-(v1))) | 1))
            }
            v1 = v1
        }
    }
    fmt.Println(v1)
    if v2 {
        switch v19 {
        case -372:
            if (0 < 0) {
                var v44 string = "\n\t"
                fmt.Println(v1, v3)
                _ = v44
            }
        default:
            if v2 {
                if (true || false) {
                    if len(v20) > 0 {
                        v20[0] = 0
                    }
                }
                v3 -= v3
                var v45 float64 = 1e10
                _ = v45
            }
        }
        var v46 int = -429
        if (!(65535 < 190)) {
            v47 := GenS4{F0: v1}
            if false {
                var v48 uint32 = 42
                v3 = v3
                _ = v48
            }
            _ = v47
        } else {
            v49 := GenS2{F0: 42, F1: v21, F2: 106}
            if (-1 != v3) {
                var v50 int = v46
                _ = v50
            }
            _ = v49
        }
        _ = v46
    }
    _ = len(v20)
    fmt.Println(v2, v21, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
