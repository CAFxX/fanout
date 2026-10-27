package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        if (2 >= 1) {
            v6 := []int{}
            _ = v6
        } else {
            if (1 <= 2) {
                v7 := []int{-258}
                if (true || true) {
                    switch (fi5 ^ fi5) {
                    case 0:
                        v8 := GenS1{F0: 127, F1: -1}
                        _ = v8
                    case 4294967295:
                        if (true || false) {
                            switch "a" {
                            case "foo bar":
                                if len(v7) > 0 {
                                    v7[0] = fi5
                                }
                            case "hello":
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                var v9 float64 = 0.1
                                _ = v9
                            }
                            var v10 uint64 = 0
                            for fi11 := 0; fi11 < 18; fi11++ {
                                var v12 float64 = 0.0
                                _ = fi11
                                _ = v12
                            }
                            _ = v10
                        } else {
                            if len(v7) > 0 {
                                v7[0] = -9223372036854775808
                            }
                        }
                    }
                    fmt.Println(fi5)
                    v13 := GenS1{F0: 0, F1: (fi5 | 2147483647)}
                    _ = v13
                } else {
                    if ((!(!true)) || (!(!(!false)))) {
                        var v14 uint8 = 1
                        v15 := GenS1{F0: 180, F1: (fi5 * fi5)}
                        v7 = append(v7, fi5)
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS1{F0: 128, F1: fi5}
                    _ = v16
                }
                if (((fi5 << ((162) & 63)) < fi5) && (4294967295 <= 42)) {
                    var v17 int8 = -1
                    switch 32767 {
                    case 32767:
                        if true {
                            fmt.Println(fi5)
                            fmt.Println(fi5)
                            var v18 uint64 = 351
                            _ = v18
                        }
                    case -9223372036854775808:
                        fmt.Println(v17)
                    default:
                        v17 = v17
                    }
                    _ = v17
                }
                _ = v7
            }
            var v19 int = (fi5 >> ((63) & 63))
            _ = v19
        }
        var v20 uint16 = 65535
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := []int{(fi21 << ((uint(v20)) & 63)), (fi5 - (fi21 << ((42) & 63)))}
            if len(v22) > 0 {
                v22[0] = fi5
            }
            var v23 int = (-(fi21))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi5
        _ = v20
    }
    var v24 float32 = ((3.14159 - (1.5 + (0.1 * 0.1))) - (-2.25 + 0.0))
    switch 0 {
    case 2147483647:
        var v25 uint16 = 152
        _ = v25
    case 866:
        fmt.Println(v24)
    default:
        v26 := []int{127, 32767, -2147483648}
        _ = v26
    }
    _ = v24
    return 127
}
func genFn1() uint8 {
    v27 := []int{0}
    _ = len(v27)
    for fi28 := 0; fi28 < 3; fi28++ {
        v29 := GenS1{F0: 23, F1: fi28}
        v30 := GenS1{F0: 255, F1: 127}
        var v31 uint8 = 2
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    _ = v27
    return 42
}

func main() {
    var v1 float32 = float32(float64(0.0))
    var v2 uint16 = 2
    var v3 int32 = 42
    v4 := make(map[string]uint32)
    v32 := GenS2{F0: 75, F1: -690, F2: 589}
    for fi33 := 0; fi33 < 18; fi33++ {
        if false {
            fmt.Println(fi33, fi33)
            var v34 int64 = -1
            _ = v34
        } else {
            var v35 uint8 = 255
            _ = v35
        }
        if (true && false) {
            v36 := v4["foo bar"]
            fmt.Println(v36, v1, v3)
            _ = v36
        }
        fmt.Println(v3, v1)
        _ = fi33
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 int16 = 32767
        v38 = (^(v38))
        v38 = (^(v38))
        _ = fi37
        _ = v38
    }
    fmt.Println(v3)
    fmt.Println(v3, v3, v1)
    switch -624 {
    case -686:
        var v39 uint32 = 186
        _ = v39
    case -2147483648:
        switch -9223372036854775808 {
        case 0:
            v1 = v1
        default:
            var v40 float32 = v1
            _ = v40
        }
    default:
        if ((962 > -32768) && (0 != 18446744073709551615)) {
            v2 = (v2 | 68)
            v3 = (v3 ^ v3)
            v41 := v4[""]
            _ = v41
        }
    }
    v42 := GenS3{F0: -128, F1: 737}
    switch ("a" + (("foo bar" + "hello") + ("a" + "a"))) {
    case "foo bar":
        switch 1 {
        case -669:
            fmt.Println(v3, v1, v2)
        }
    default:
        delete(v4, "\n\t")
    }
    var v43 int8 = 42
    var v44 uint64 = 4294967295
    v45 := GenS2{F0: 137, F1: 0, F2: 42}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
