package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint, p1 uint) int32 {
    var v5 int32 = -2147483648
    var v6 bool = true
    p1 = p0
    v7 := []int{-76, -9223372036854775808, 32767}
    p0 += 380
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -29
}
func genFn1(p0 int32, p1 int8) uint8 {
    var v8 int32 = (p0 ^ p0)
    if (1.5 < -2.25) {
        var v9 uint = 659
        if (-9223372036854775808 <= 128) {
            v10 := GenS1{F0: p1}
            p1 = (p1 * p1)
            _ = v10
        }
        var v11 int16 = 1
        _ = v11
        _ = v9
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = v8
    return 127
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -397
    var v3 string = "foo bar"
    v4 := []uint32{}
    v12 := GenS2{F0: 771}
    var v13 uint32 = 0
    delete(v1, 1)
    switch 9223372036854775807 {
    case 2147483647:
        if true {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        } else {
            switch 1 {
            case -1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 786:
                v13 = v13
            default:
                var v14 uint64 = 1
                _ = v14
            }
            if true {
                switch 2147483647 {
                case 9223372036854775807:
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if false {
                            var v16 uint64 = 863
                            _ = v16
                        }
                        var v17 int8 = -128
                        v17 -= (-(v17))
                        _ = fi15
                        _ = v17
                    }
                case -1:
                    switch -516 {
                    case -1:
                        v3 = "\n\t"
                    case 4294967295:
                        if len(v4) > 0 {
                            v4[0] = (^(v13))
                        }
                    default:
                        if (!false) {
                            fmt.Println(v2, v2, v13)
                            _ = len(v4)
                        }
                    }
                case 32768:
                    for fi18 := 0; fi18 < 5; fi18++ {
                        fmt.Println(v13, fi18, v2)
                        switch ("\n\t" + v3) {
                        case "a":
                            var v19 int = (-(fi18))
                            _ = v19
                        case "":
                            var v20 string = ("foo bar" + v3)
                            _ = v20
                        default:
                            v2 = 0
                        }
                        _ = fi18
                    }
                }
                fmt.Println(v13)
                _ = len(v1)
            }
            var v21 string = v3
            _ = v21
        }
    case -1:
        if (true || true) {
            var v22 uint8 = 128
            _ = v22
        } else {
            var v23 int = 32767
            fmt.Println(v3)
            v24 := GenS2{F0: 240}
            _ = v23
            _ = v24
        }
    }
    var v25 uint16 = 255
    v26 := GenS1{F0: -1}
    if ((0.1 / -2.25) >= float64(3.14159)) {
        for fi27 := 0; fi27 < 11; fi27++ {
            fmt.Println(v2, v3)
            _ = fi27
        }
        if (!true) {
            fmt.Println(v25, v25, v25)
            var v28 uint8 = 54
            var v29 float32 = 0.1
            _ = v28
            _ = v29
        } else {
            var v30 int32 = 1
            _ = v30
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v3)
            switch (-(fi31)) {
            case 129:
                v32 := v1[fi31]
                _ = v32
            default:
                fmt.Println(v13)
            }
            _ = fi31
        }
    } else {
        if len(v4) > 0 {
            v4[0] = v13
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            if (!false) {
                if true {
                    var v34 int = fi33
                    for fi35 := 0; fi35 < 13; fi35++ {
                        fmt.Println(v25, v34)
                        var v36 uint16 = (v25 - 1)
                        v37 := GenS1{F0: -14}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                }
                v1[fi33] = "foo bar"
                var v38 int32 = 563
                _ = v38
            }
            _ = fi33
        }
        v25 -= (v25 & 0)
    }
    v25 += 42
    v3 = "\n\t"
    _ = len(v4)
    v2 = v2
    switch "hello" {
    case "a":
        fmt.Println(v2)
    default:
        var v39 int16 = v2
        _ = v39
    }
    fmt.Println(v2, v13, v2)
    var v40 int64 = -119
    var v41 uint16 = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
