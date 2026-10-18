package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int32; F2 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 uint; F2 uint16 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 int; F1 string; F2 uint16 }
func genFn0() string {
    var v5 float64 = 3.14159
    fmt.Println(v5)
    for fi6 := 0; fi6 < 11; fi6++ {
        v5 -= 0.1
        _ = fi6
    }
    var v7 uint16 = 0
    v5 -= v5
    _ = v5
    _ = v7
    return (("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + "")
}
func genFn1(p0 bool, p1 int16, p2 uint32) uint32 {
    for fi8 := 0; fi8 < 5; fi8++ {
        if ((true || p0) || p0) {
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
        var v10 int64 = -1
        p2 = 0
        _ = fi8
        _ = v10
    }
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := []uint32{}
    var v2 int = -415
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    for fi11 := 0; fi11 < 3; fi11++ {
        switch "hello" {
        case "foo bar":
            var v12 bool = false
            _ = v12
        case "":
            fmt.Println(v2, v2, v2)
        default:
            var v13 int64 = int64(v2)
            _ = v13
        }
        fmt.Println(fi11, fi11, v4)
        _ = fi11
    }
    v14 := GenS1{F0: -1, F1: 244}
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v4, v2, fi15)
        if (!(173 >= 42)) {
            v4 = (v4 % ((v4) | 1))
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 int16 = 32767
                _ = len(v1)
                _ = fi16
                _ = v17
            }
            if (("a" < "\n\t") || false) {
                for fi18 := 0; fi18 < 20; fi18++ {
                    if (true && false) {
                        var v19 int8 = -54
                        v2 -= fi18
                        var v20 float32 = 3.14159
                        switch (("a" + "a") + "") {
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            if (801 == 65535) {
                                v2 -= v2
                                fmt.Println(v20, fi15)
                            } else {
                                fmt.Println(fi15, fi18)
                                v2 += -848
                                var v21 uint8 = 1
                                _ = v21
                            }
                            switch fi15 {
                            case -2147483648:
                                switch 555 {
                                case 32768:
                                    if (true || (65535 >= 32767)) {
                                        v2 += fi15
                                        switch "" {
                                        case "foo bar":
                                            v2 = (fi18 - 128)
                                        case "a":
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                        default:
                                            switch (-(fi18)) {
                                            case -2147483648:
                                                if (false || false) {
                                                    if true {
                                                        if false {
                                                            v22 := GenS2{F0: 905}
                                                            v20 += 3.14159
                                                            var v23 int = fi15
                                                            var v24 int8 = v19
                                                            _ = v22
                                                            _ = v23
                                                            _ = v24
                                                        } else {
                                                            v4 += v4
                                                            fmt.Println(v2, v19, fi18)
                                                        }
                                                        v25 := v3[("foo bar" + "")]
                                                        if (true || true) {
                                                            fmt.Println(v19, v4)
                                                            var v26 int64 = v4
                                                            _ = v26
                                                        }
                                                        _ = v25
                                                    } else {
                                                        v2 = 0
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v1 = append(v1, 437)
                                                    }
                                                    var v27 uint8 = 110
                                                    v20 += v20
                                                    delete(v3, "foo bar")
                                                    _ = v27
                                                }
                                            case 32767:
                                                if true {
                                                    v28 := GenS3{F0: (^(v4))}
                                                    var v29 int = fi15
                                                    if (("\n\t" >= (("a" + "a") + "\n\t")) && false) {
                                                        v30 := GenS4{F0: v19}
                                                        v2 = v2
                                                        switch (v29 | 9223372036854775807) {
                                                        case 4294967295:
                                                            v31 := GenS5{F0: 42, F1: 42, F2: 800}
                                                            _ = v31
                                                        }
                                                        _ = v30
                                                    }
                                                    var v32 string = (("\n\t" + ("a" + "\n\t")) + "foo bar")
                                                    _ = len(v3)
                                                    _ = v28
                                                    _ = v29
                                                    _ = v32
                                                }
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        v19 += (v19 << ((1) & 63))
                                        v33 := GenS6{F0: 255}
                                        _ = v33
                                    }
                                }
                            case 128:
                                if (!true) {
                                    var v34 int64 = (v4 & 42)
                                    if (!(0 != 291)) {
                                        v20 = (v20 - v20)
                                        v4 -= 774
                                    }
                                    v20 = ((v20 / -2.25) * -2.25)
                                    _ = v34
                                }
                            default:
                                switch "foo bar" {
                                case "":
                                    var v35 float32 = v20
                                    _ = v35
                                case "hello":
                                    v36 := GenS6{F0: 0}
                                    _ = v36
                                case "foo bar":
                                    if true {
                                        v37 := GenS7{F0: 65535}
                                        switch "foo bar" {
                                        case "":
                                            v4 = (v4 - v4)
                                        case "a":
                                            _ = len(v3)
                                        case "hello":
                                            if true {
                                                switch "hello" {
                                                case "":
                                                    var v38 int16 = 1000
                                                    _ = v38
                                                case "foo bar":
                                                    v4 = (v4 ^ (v4 >> ((42) & 63)))
                                                }
                                            }
                                        }
                                        var v39 uint64 = 0
                                        _ = v37
                                        _ = v39
                                    }
                                }
                            }
                            fmt.Println(v20)
                        case "":
                            fmt.Println(v20)
                        default:
                            fmt.Println(v20, fi15)
                        }
                        _ = v19
                        _ = v20
                    }
                    v4 = (v4 >> ((73) & 63))
                    _ = fi18
                }
                v4 -= v4
                v40 := GenS8{F0: (-115 <= 34), F1: 2, F2: 743}
                _ = v40
            }
        } else {
            switch fi15 {
            case 606:
                v4 = (-(v4))
            case 32767:
                var v41 uint8 = 42
                _ = v41
            }
            v42 := GenS5{F0: 0, F1: -2147483648, F2: -1}
            _ = v42
        }
        v43 := GenS9{F0: 96, F1: false}
        _ = fi15
        _ = v43
    }
    v4 -= 42
    _ = len(v3)
    v44 := GenS10{F0: (^(v2)), F1: "foo bar", F2: 2}
    var v45 float32 = ((3.14159 + (0.1 * (-2.25 * 1e10))) / 1.0)
    if false {
        for fi46 := 0; fi46 < 18; fi46++ {
            v2 = (fi46 + v2)
            _ = fi46
        }
    }
    fmt.Println(v2, v45)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
