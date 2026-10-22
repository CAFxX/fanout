package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0() uint32 {
    if (false && false) {
        switch "foo bar" {
        case "a":
            for fi5 := 0; fi5 < 8; fi5++ {
                if (!false) {
                    var v6 int32 = 42
                    for fi7 := 0; fi7 < 12; fi7++ {
                        var v8 string = (("\n\t" + "") + "\n\t")
                        fmt.Println(fi7)
                        _ = fi7
                        _ = v8
                    }
                    _ = v6
                } else {
                    if ((-2147483648 <= -2147483648) || (1.5 == (1.5 * (1e10 - ((1.5 / 1.5) * 3.14159))))) {
                        v9 := []int{fi5, -9223372036854775808, (fi5 | -2147483648)}
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 12; fi10++ {
                            var v11 uint32 = 2147483647
                            _ = fi10
                            _ = v11
                        }
                    }
                    var v12 int8 = 0
                    v12 -= 0
                    _ = v12
                }
                v13 := GenS1{F0: 128}
                _ = fi5
                _ = v13
            }
        case "hello":
            var v14 bool = (!(2 < 2))
            _ = v14
        case "\n\t":
            var v15 int = 2147483647
            _ = v15
        }
    } else {
        var v16 float64 = 1e10
        switch "hello" {
        case "foo bar":
            v17 := []int{6, 880, 32767, 32767}
            _ = v17
        default:
            v16 = v16
        }
        v16 = (((v16 * (v16 * v16)) + float64(((0.1 * 0.0) / 1.0))) - v16)
        _ = v16
    }
    var v18 float64 = (0.0 * 1.5)
    if (9223372036854775807 != -598) {
        v19 := GenS2{F0: "a", F1: 1, F2: 565}
        _ = v19
    } else {
        if false {
            switch "" {
            case "hello":
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 int = (fi20 << ((758) & 63))
                    _ = fi20
                    _ = v21
                }
            case "a":
                v18 -= v18
            }
            switch -1 {
            case -9223372036854775808:
                v18 = v18
            }
        } else {
            var v22 int8 = 127
            v18 -= (0.0 - 3.14159)
            _ = v22
        }
        v18 = 0.0
    }
    _ = v18
    return 664
}
func genFn1(p0 int8) uint16 {
    p0 -= (^(p0))
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(p0)
        _ = fi23
    }
    _ = p0
    return 65535
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int64 = -1
    _ = len(v3)
    v24 := []int{127, 128, 32768}
    var v25 uint32 = (v1 >> ((0) & 63))
    v1 = 0
    v1 += 757
    switch int(v1) {
    case 42:
        v26 := GenS3{F0: "a", F1: 127}
        _ = v26
    case 71:
        var v27 uint16 = 2
        _ = v27
    case 32767:
        v24 = append(v24, 2147483647)
    }
    if (true && (32767 >= -9223372036854775808)) {
        _ = len(v24)
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v1 = (v1 >> ((4294967295) & 63))
        for fi29 := 0; fi29 < 8; fi29++ {
            if false {
                var v30 int8 = 100
                if ((255 <= 128) || (!(v30 <= v30))) {
                    v1 = v25
                }
                _ = v30
            }
            v4 = (v4 << ((1) & 63))
            v31 := GenS3{F0: "hello", F1: 128}
            _ = fi29
            _ = v31
        }
        fmt.Println(v1, v1, v1)
        _ = fi28
    }
    if false {
        v24 = append(v24, 2147483647)
        var v32 int8 = 0
        _ = v32
    } else {
        var v33 float32 = (0.0 - 1e10)
        _ = v33
    }
    v34 := GenS3{F0: "hello", F1: 127}
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint16 = 1
        for fi37 := 0; fi37 < 2; fi37++ {
            v38 := GenS2{F0: "hello", F1: v25, F2: 0}
            var v39 uint8 = 255
            fmt.Println(v4)
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    var v40 int = 120
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v40
}
