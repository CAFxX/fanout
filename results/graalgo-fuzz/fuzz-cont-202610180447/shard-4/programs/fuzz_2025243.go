package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 string, p1 uint) uint8 {
    p1 = 0
    p1 = 4294967295
    _ = p0
    _ = p1
    return 128
}
func genFn1(p0 int64, p1 int, p2 int32) int16 {
    fmt.Println(p2, p1, p2)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return -738
}

func main() {
    v1 := []uint32{144, 42}
    var v2 float32 = (0.0 * 1.5)
    var v3 int = 0
    var v4 int64 = -9223372036854775808
    v5 := []string{"\n\t", "foo bar", ("foo bar" + ("foo bar" + (("" + "a") + "hello")))}
    v6 := make(map[string]int)
    var v7 int8 = 100
    var v8 string = "a"
    if true {
        var v9 int8 = 39
        v3 = (v3 / (((v3 + v3)) | 1))
        _ = v9
    } else {
        v4 = (-(v4))
    }
    var v10 uint8 = 127
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 65535
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v13 int32 = 2147483647
            var v14 uint64 = 4294967295
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    switch v3 {
    case 42:
        var v15 uint8 = (^(v10))
        _ = v15
    case 128:
        var v16 int16 = 1000
        _ = v16
    case -1:
        _ = len(v1)
    }
    v1 = append(v1, 4294967295)
    if (1 <= 2) {
        v10 = (v10 & v10)
        if (v2 < 0.1) {
            for fi17 := 0; fi17 < 13; fi17++ {
                v10 = (v10 << ((65535) & 63))
                _ = fi17
            }
        }
        var v18 string = v8
        _ = v18
    } else {
        var v19 int = (v3 & v3)
        _ = v19
    }
    v20 := v6[("" + v8)]
    switch v3 {
    case 32768:
        v5 = append(v5, v8)
    case 632:
        for fi21 := 0; fi21 < 1; fi21++ {
            switch v3 {
            case 0:
                var v22 float32 = 0.1
                _ = v22
            case 832:
                switch (fi21 * (fi21 / ((9223372036854775807) | 1))) {
                case -9223372036854775808:
                    fmt.Println(v7, v10, v7)
                }
            }
            _ = fi21
        }
    case -233:
        var v23 uint8 = (v10 % (((v10 & v10)) | 1))
        _ = v23
    default:
        switch (v3 ^ v3) {
        case 301:
            v10 -= v10
        }
    }
    fmt.Println(v8, v4)
    var v24 float64 = (float64(v2) + (float64(0.1) + -2.25))
    switch 415 {
    case 0:
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float64 = (float64((1e10 + ((v2 - ((v2 * (-2.25 + (v2 - (v2 * v2)))) + -2.25)) + v2))) / float64((v2 * v2)))
            v2 = v2
            v20 = (v20 ^ v20)
            _ = fi25
            _ = v26
        }
    case 128:
        if false {
            for fi27 := 0; fi27 < 7; fi27++ {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v29 := GenS1{F0: v8, F1: 65535}
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            v5 = append(v5, v8)
            if true {
                v4 = v4
            } else {
                if false {
                    var v30 float64 = float64(v2)
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v20 = 2147483647
                        var v32 uint = 42
                        fmt.Println(v8)
                        _ = fi31
                        _ = v32
                    }
                    delete(v6, (v8 + ""))
                    _ = v30
                }
                if (!false) {
                    v33 := GenS2{F0: v7}
                    var v34 uint16 = 2
                    v35 := GenS3{F0: 59}
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    var v36 bool = (0 <= 0)
                    for fi37 := 0; fi37 < 19; fi37++ {
                        switch (v20 & (v3 * v3)) {
                        case -2147483648:
                            fmt.Println(v10)
                        case 42:
                            var v38 uint16 = 255
                            _ = v38
                        case -1:
                            fmt.Println(v10, v4)
                        default:
                            v39 := GenS3{F0: 448}
                            _ = v39
                        }
                        _ = fi37
                    }
                    _ = v36
                }
            }
        }
    case -530:
        var v40 uint8 = 188
        _ = v40
    default:
        var v41 uint32 = 681
        _ = v41
    }
    var v42 int8 = (^(v7))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v43 string = v8
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 bool = false
        if v45 {
            v46 := GenS4{F0: 811, F1: 0, F2: 499}
            if (0.0 >= v2) {
                v20 = fi44
            }
            _ = v46
        } else {
            var v47 uint = 0
            _ = v47
        }
        for fi48 := 0; fi48 < 16; fi48++ {
            v7 += -62
            if v45 {
                delete(v6, "foo bar")
                if (417 >= 45) {
                    var v49 uint16 = 1
                    switch int(v4) {
                    case 127:
                        switch (v20 * int(v49)) {
                        case -9223372036854775808:
                            var v50 uint64 = 0
                            _ = v50
                        case 1:
                            var v51 string = ((v43 + "") + (v8 + (((v8 + v43) + v8) + (v8 + "hello"))))
                            _ = v51
                        case -147:
                            v52 := GenS5{F0: 100, F1: 42, F2: ""}
                            _ = v52
                        }
                    case -2147483648:
                        fmt.Println(fi48)
                    case -9223372036854775808:
                        v49 = 643
                    default:
                        v53 := v6[v8]
                        _ = v53
                    }
                    _ = v49
                }
            }
            v2 = 3.14159
            _ = fi48
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
