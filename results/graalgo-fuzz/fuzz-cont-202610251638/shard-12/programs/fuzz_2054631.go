package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint16; F1 string; F2 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 int) uint64 {
    p0 = p0
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return 0
}
func genFn1(p0 uint16) string {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 / ((9223372036854775807) | 1))}
        if false {
            if (-2147483648 > -172) {
                v8 := GenS1{F0: p0, F1: (p0 - (p0 * 91))}
                var v9 uint64 = 4294967295
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 int = fi10
                    v11 = v11
                    p0 = p0
                    _ = fi10
                    _ = v11
                }
                switch (fi6 >> ((42) & 63)) {
                case 9223372036854775807:
                    v12 := GenS1{F0: p0, F1: 255}
                    _ = v12
                default:
                    var v13 int16 = 1000
                    _ = v13
                }
                p0 += (p0 << ((4294967295) & 63))
            }
        }
        _ = len(v7)
        _ = fi6
        _ = v7
    }
    var v14 float32 = (0.0 * float32(3.14159))
    _ = p0
    _ = v14
    return (("hello" + ("\n\t" + "foo bar")) + "\n\t")
}

func main() {
    var v1 int16 = -32768
    v2 := []string{("" + "hello"), "\n\t"}
    var v3 float64 = (0.1 + 1e10)
    v4 := make(map[int]uint32)
    var v5 float32 = float32((v3 - 1.5))
    for fi15 := 0; fi15 < 17; fi15++ {
        switch (fi15 * 9223372036854775807) {
        case 4294967295:
            if (42 == -650) {
                if false {
                    fmt.Println(v3, v1, v3)
                }
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    fmt.Println(v1, v3)
                    v5 = (v5 * v5)
                    _ = fi16
                }
                v2 = append(v2, "\n\t")
                var v17 int = fi15
                _ = v17
            }
        }
        switch (fi15 << ((255) & 63)) {
        case 2147483647:
            v1 = -32768
        case 42:
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 int32 = 2147483647
                _ = fi18
                _ = v19
            }
        case -2147483648:
            var v20 int32 = 42
            _ = v20
        default:
            fmt.Println(v5, v3, v5)
        }
        delete(v4, fi15)
        _ = fi15
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        fmt.Println(v1)
        v22 := v4[fi21]
        var v23 int64 = -191
        _ = fi21
        _ = v22
        _ = v23
    }
    v5 = float32(v3)
    v24 := GenS2{F0: ""}
    if (-2147483648 < 0) {
        if true {
            _ = len(v4)
            for fi25 := 0; fi25 < 10; fi25++ {
                for fi26 := 0; fi26 < 10; fi26++ {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    if (32767 >= 526) {
                        _ = len(v4)
                        switch "" {
                        case "hello":
                            var v27 int = fi26
                            _ = v27
                        case "foo bar":
                            if false {
                                v28 := GenS3{F0: 32767, F1: genFn0(fi26), F2: 78}
                                _ = v28
                            }
                            var v29 uint8 = 2
                            _ = v29
                        }
                        switch fi25 {
                        case 127:
                            v30 := v4[(fi26 & fi26)]
                            _ = v30
                        case 42:
                            if false {
                                fmt.Println(v3, v5, v1)
                            } else {
                                var v31 uint16 = 0
                                v5 = 0.0
                                var v32 float32 = v5
                                fmt.Println(v31)
                                _ = v31
                                _ = v32
                            }
                        case 1:
                            fmt.Println(fi25, v5, fi26)
                            var v33 int32 = 1
                            _ = v33
                        default:
                            v5 += v5
                        }
                    }
                    v5 = 0.1
                    var v34 string = ((((("hello" + ("" + "\n\t")) + "\n\t") + "\n\t") + ("hello" + "a")) + "\n\t")
                    _ = fi26
                    _ = v34
                }
                var v35 int = (fi25 - fi25)
                v35 = fi25
                _ = fi25
                _ = v35
            }
            v5 = v5
        } else {
            var v36 uint32 = 2147483648
            for fi37 := 0; fi37 < 9; fi37++ {
                if true {
                    for fi38 := 0; fi38 < 2; fi38++ {
                        fmt.Println(v1, v3, v1)
                        v39 := v4[(fi37 / (((^(fi38))) | 1))]
                        var v40 float32 = v5
                        v41 := GenS3{F0: 42, F1: 4294967295, F2: 42}
                        _ = fi38
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    v42 := v4[4294967295]
                    _ = v42
                }
                _ = fi37
            }
            if true {
                var v43 int = 4294967295
                switch (v43 & (v43 * -9223372036854775808)) {
                case 127:
                    if (42 >= -2147483648) {
                        v44 := GenS4{F0: 1, F1: v43}
                        v4[42] = 271
                        _ = v44
                    }
                default:
                    _ = len(v4)
                }
                _ = v43
            }
            _ = v36
        }
        v5 = v5
        if true {
            if (-2147483648 != 0) {
                v5 = v5
                v45 := v4[128]
                _ = v45
            } else {
                fmt.Println(v5, v5, v3)
            }
            var v46 float64 = (v3 - 1e10)
            _ = len(v4)
            _ = v46
        }
    }
    v47 := GenS3{F0: 1, F1: 0, F2: 0}
    if ("a" == "foo bar") {
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 float32 = 1.5
            v5 += v5
            _ = len(v4)
            _ = fi48
            _ = v49
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v50 := GenS5{F0: 65535, F1: "\n\t", F2: (v1 >> ((2) & 63))}
        _ = v50
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        v4[fi51] = 1
        _ = fi51
    }
    if false {
        switch "foo bar" {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        case "":
            v1 = (v1 & 1000)
        case "\n\t":
            var v52 float64 = (v3 / float64((float32(v3) * 0.0)))
            _ = v52
        }
        for fi53 := 0; fi53 < 13; fi53++ {
            if (false && true) {
                var v54 string = "a"
                _ = v54
            } else {
                fmt.Println(v5)
            }
            _ = fi53
        }
        if (-1 <= 71) {
            var v55 float64 = 1.5
            for fi56 := 0; fi56 < 8; fi56++ {
                var v57 int32 = 42
                _ = fi56
                _ = v57
            }
            if (!false) {
                v58 := GenS1{F0: 255, F1: 255}
                var v59 float64 = float64(float32(v3))
                fmt.Println(v5, v55)
                _ = v58
                _ = v59
            }
            _ = v55
        } else {
            switch 42 {
            case 128:
                fmt.Println(v1, v5, v1)
            case 32767:
                v60 := GenS6{F0: true}
                _ = v60
            case -2147483648:
                v5 += v5
            default:
                v3 = 1.5
            }
            delete(v4, 0)
        }
    }
    v61 := GenS3{F0: 32767, F1: 4294967295, F2: 239}
    var v62 int64 = 1
    delete(v4, 77)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v61
    _ = v62
}
