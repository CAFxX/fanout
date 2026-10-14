package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int16; F1 string; F2 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 uint32 = 42
    var v4 uint16 = 351
    v5 := make(map[string]uint32)
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            var v7 bool = true
            if v7 {
                v8 := GenS1{F0: 93, F1: 683, F2: -1}
                _ = v8
            }
            _ = v7
        }
        if (("\n\t" + ("foo bar" + "foo bar")) >= "\n\t") {
            var v9 uint8 = 0
            fmt.Println(fi6, v4)
            v10 := GenS2{F0: (!false)}
            _ = v10
            _ = v9
        } else {
            v4 = (v4 - v4)
        }
        _ = fi6
    }
    var v11 float32 = (((0.0 * 0.1) * 3.14159) / ((((0.1 + float32((-2.25 - 3.14159))) + -2.25) * 1e10) + 1.5))
    var v12 uint32 = v3
    if (127 >= 2147483647) {
        v13 := v2[42]
        fmt.Println(v3, v12, v13)
        _ = v13
    } else {
        fmt.Println(v11, v3, v3)
        v14 := []int{42, 32767, 0}
        fmt.Println(v3, v3)
        _ = v14
    }
    v12 -= (v12 >> ((42) & 63))
    if false {
        v15 := GenS3{F0: 0}
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 int32 = -1
                _ = len(v1)
                _ = fi17
                _ = v18
            }
            if true {
                if (fi16 < (fi16 % ((fi16) | 1))) {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v20 := GenS4{F0: 255, F1: ""}
                        _ = fi19
                        _ = v20
                    }
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v11, v12, v4)
                    }
                    var v21 int = fi16
                    _ = v21
                }
            } else {
                if false {
                    var v22 int64 = -824
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 int8 = 63
                        switch (fi16 & 0) {
                        case 288:
                            v4 -= v4
                        case -264:
                            v11 += 3.14159
                        case -9223372036854775808:
                            var v25 uint16 = (v4 + v4)
                            _ = v25
                        default:
                            var v26 uint8 = 2
                            _ = v26
                        }
                        v22 -= (v22 << ((409) & 63))
                        _ = fi23
                        _ = v24
                    }
                    _ = v22
                } else {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v11 = 0.1
                        v11 += (((((0.0 - v11) + (v11 / -2.25)) + v11) - (0.1 - 3.14159)) * 1e10)
                        v3 = (^(v3))
                        _ = fi27
                    }
                }
                if false {
                    v28 := GenS1{F0: 255, F1: -1, F2: 1}
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v12 += v12
                        v30 := GenS5{F0: (("\n\t" + "") + ("foo bar" + "\n\t")), F1: 0}
                        v12 -= (^(v3))
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                } else {
                    v31 := []int{fi16, fi16}
                    if false {
                        fmt.Println(v12)
                        _ = len(v5)
                    }
                    switch fi16 {
                    case -2147483648:
                        var v32 float64 = 1e10
                        _ = v32
                    }
                    _ = v31
                }
                fmt.Println(v12, fi16, v12)
            }
            _ = fi16
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            for fi34 := 0; fi34 < 4; fi34++ {
                if ((1.5 != -2.25) && (450 <= (fi33 + 4294967295))) {
                    var v35 uint32 = 65535
                    v36 := GenS3{F0: 1}
                    _ = v35
                    _ = v36
                } else {
                    var v37 bool = (false || (1e10 < -2.25))
                    var v38 int64 = 9223372036854775807
                    fmt.Println(v11, v4)
                    _ = v37
                    _ = v38
                }
                fmt.Println(v12)
                _ = fi34
            }
            _ = fi33
        }
        _ = v15
    } else {
        v39 := GenS6{F0: 1, F1: v3}
        _ = v39
    }
    v40 := GenS7{F0: -860, F1: ("\n\t" + "a"), F2: (v11 == 0.1)}
    v4 = v4
    v3 += v3
    v41 := []int{-332, 1, 127, -372}
    v42 := GenS7{F0: 920, F1: ("hello" + ""), F2: (0 < 18)}
    switch ("foo bar" + (("foo bar" + "") + ("\n\t" + "a"))) {
    case "foo bar":
        v3 = 1
    default:
        v43 := GenS2{F0: false}
        _ = v43
    }
    var v44 int16 = -29
    var v45 int8 = 0
    var v46 int32 = 0
    v41 = append(v41, 2147483647)
    v47 := GenS6{F0: -9223372036854775808, F1: v3}
    v3 += (v12 | (v3 / (((^(v12))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
