package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8 }
func genFn0() int8 {
    var v6 int16 = 42
    var v7 uint8 = 2
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 int32) uint64 {
    if false {
        if false {
            v8 := GenS1{F0: 2, F1: -116}
            v9 := GenS1{F0: 0, F1: -56}
            var v10 uint32 = 65535
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := []int{}
            _ = v11
        }
    }
    p0 = 1
    p0 = p0
    _ = p0
    return 667
}

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 255
    v3 := []string{}
    var v4 float32 = 3.14159
    v5 := make(map[uint32]string)
    var v12 int64 = 9223372036854775807
    v12 -= 1
    v1 += (v1 | 65535)
    v13 := GenS1{F0: v2, F1: 42}
    for fi14 := 0; fi14 < 11; fi14++ {
        switch (fi14 ^ fi14) {
        case -2147483648:
            v12 = (v12 / (((v12 % ((9223372036854775807) | 1))) | 1))
        case 0:
            if (true || true) {
                v4 -= (v4 * v4)
            }
        default:
            var v15 bool = false
            _ = v15
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v4 = v4
        switch (fi16 % ((32768) | 1)) {
        case 128:
            switch ((("\n\t" + "hello") + "a") + ("a" + "a")) {
            case "hello":
                var v17 int16 = 1000
                _ = v17
            case "\n\t":
                fmt.Println(v2)
            case "":
                fmt.Println(v12, v2)
            default:
                switch (fi16 + (fi16 & int(v12))) {
                case 1:
                    var v18 uint16 = 255
                    _ = v18
                default:
                    if false {
                        if true {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v4 = v4
                                _ = fi19
                            }
                            fmt.Println(v2)
                        } else {
                            var v20 uint16 = 1
                            if (!(genFn1(-933) >= 90)) {
                                var v21 uint8 = (^(v2))
                                for fi22 := 0; fi22 < 19; fi22++ {
                                    if false {
                                        var v23 uint8 = 93
                                        v12 -= (v12 * (-(v12)))
                                        var v24 int32 = -559
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 bool = (!false)
                                        _ = len(v3)
                                        if (!(42 < 0)) {
                                            fmt.Println(v12, fi22)
                                            var v26 int64 = v12
                                            _ = v26
                                        }
                                        if v25 {
                                            v27 := GenS1{F0: (v2 / ((v2) | 1)), F1: -1}
                                            v4 += (-2.25 + ((3.14159 - 0.1) * v4))
                                            _ = v27
                                        } else {
                                            v12 -= (-(v12))
                                            switch "hello" {
                                            case "foo bar":
                                                v28 := v5[uint32(v21)]
                                                _ = v28
                                            case "hello":
                                                v21 = 135
                                            }
                                            v12 = (^(v12))
                                        }
                                        fmt.Println(v2)
                                        _ = v25
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    _ = fi22
                                }
                                _ = v21
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            for fi29 := 0; fi29 < 16; fi29++ {
                                v5[234] = ""
                                fmt.Println(v2, v1, v2)
                                if true {
                                    v30 := GenS2{F0: ("a" + "foo bar")}
                                    _ = v30
                                } else {
                                    fmt.Println(v20, fi29)
                                    var v31 int = (fi29 | (fi16 & fi29))
                                    _ = v31
                                }
                                fmt.Println(fi29, v2)
                                v32 := GenS2{F0: ""}
                                _ = fi29
                                _ = v32
                            }
                            _ = v20
                        }
                        fmt.Println(v2, v4)
                    }
                }
            }
        }
        v12 -= 126
        _ = fi16
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := v5[v1]
        _ = fi33
        _ = v34
    }
    var v35 int = 128
    if (int8(v35) < 0) {
        v4 += v4
        v1 += (v1 >> ((42) & 63))
        var v36 int64 = 9223372036854775807
        _ = v36
    }
    fmt.Println(v12, v2, v12)
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 string = ""
        fmt.Println(v2, v35)
        v38 = v38
        _ = fi37
        _ = v38
    }
    if true {
        v1 += v1
        v12 -= (v12 + v12)
    } else {
        var v39 bool = (42 == -1)
        var v40 int8 = -1
        fmt.Println(v35, v40)
        _ = v39
        _ = v40
    }
    v41 := GenS1{F0: v2, F1: 100}
    if (!((v1 % ((v1) | 1)) >= (v1 % ((42) | 1)))) {
        v5[uint32(v2)] = "foo bar"
    } else {
        delete(v5, 2)
        var v42 uint8 = 127
        fmt.Println(v2)
        _ = v42
    }
    var v43 uint32 = (^(v1))
    switch v35 {
    case 9223372036854775807:
        v1 = (v43 ^ 65535)
    case 4294967295:
        if (false || true) {
            var v44 float32 = 1.5
            _ = v44
        }
    case -9223372036854775808:
        var v45 string = "a"
        _ = v45
    default:
        switch (^(v35)) {
        case 1:
            var v46 int8 = -128
            _ = v46
        case -62:
            fmt.Println(v2, v35)
        case 42:
            v47 := GenS3{F0: true, F1: v35}
            _ = v47
        default:
            v48 := GenS4{F0: v2}
            _ = v48
        }
    }
    var v49 uint = 724
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v41
    _ = v43
    _ = v49
    _ = v5
}
