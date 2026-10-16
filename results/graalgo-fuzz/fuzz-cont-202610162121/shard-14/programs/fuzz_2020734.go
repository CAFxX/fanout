package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 int16 }
func genFn0(p0 uint32, p1 string, p2 uint16) int8 {
    v5 := []int{127, 127, -2147483648, 0}
    var v6 uint64 = 487
    var v7 uint64 = v6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 bool, p1 uint8, p2 uint) uint16 {
    fmt.Println(p1)
    v8 := GenS1{F0: -9223372036854775808, F1: 233}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{580, 106, 270, 689}
    var v3 int = 0
    var v4 uint8 = 169
    v4 += v4
    v9 := GenS2{F0: 0}
    v10 := GenS3{F0: (0 != 0), F1: 308}
    v3 = v3
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint8 = v4
        var v13 string = "\n\t"
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS4{F0: 1}
    fmt.Println(v4, v4)
    v4 = (^(v4))
    if (!(!true)) {
        switch v3 {
        case 127:
            v3 = 448
        case 128:
            if true {
                v4 -= v4
            } else {
                var v15 uint32 = 202
                v16 := GenS1{F0: 413, F1: 589}
                if true {
                    fmt.Println(v3, v3, v4)
                } else {
                    switch "foo bar" {
                    case "":
                        if (!true) {
                            if true {
                                v17 := v1[("" + (("a" + ("foo bar" + "\n\t")) + "a"))]
                                _ = v17
                            }
                            if true {
                                _ = len(v2)
                            }
                            if false {
                                var v18 uint32 = v15
                                v19 := GenS3{F0: false, F1: -41}
                                v15 += (v18 + v15)
                                _ = v18
                                _ = v19
                            } else {
                                v3 -= v3
                                v20 := GenS1{F0: 9223372036854775807, F1: 1}
                                _ = v20
                            }
                        }
                    default:
                        var v21 float32 = 0.1
                        _ = v21
                    }
                    v1["a"] = ""
                    v22 := GenS3{F0: ("hello" <= "\n\t"), F1: 1}
                    _ = v22
                }
                _ = v15
                _ = v16
            }
        default:
            var v23 int16 = 1000
            _ = v23
        }
        if false {
            v4 -= v4
        }
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v3 = -827
            _ = fi24
        }
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        var v26 int64 = 1
        fmt.Println(v26)
        _ = fi25
        _ = v26
    }
    v4 = 255
    for fi27 := 0; fi27 < 19; fi27++ {
        if true {
            var v28 string = ("" + "foo bar")
            var v29 int32 = -2147483648
            var v30 uint16 = 0
            _ = v28
            _ = v29
            _ = v30
        } else {
            if true {
                var v31 string = (("\n\t" + ("a" + ("hello" + "\n\t"))) + "a")
                if true {
                    v3 = (v3 << ((65535) & 63))
                }
                if true {
                    v32 := GenS5{F0: 18446744073709551615, F1: 638, F2: 32767}
                    _ = v32
                } else {
                    fmt.Println(fi27, fi27)
                    v2 = append(v2, 2147483648)
                    if true {
                        v33 := GenS6{F0: 2147483647, F1: -1}
                        v3 = (v3 - fi27)
                        _ = v33
                    }
                }
                _ = v31
            }
            v3 = v3
            var v34 int16 = 42
            _ = v34
        }
        _ = fi27
    }
    var v35 uint16 = 255
    fmt.Println(v35, v3, v4)
    v3 -= v3
    if (!true) {
        v36 := v1["a"]
        v36 = v36
        if true {
            v1[((v36 + "") + (v36 + v36))] = v36
        } else {
            var v37 int = 9223372036854775807
            v38 := GenS1{F0: 9223372036854775807, F1: 0}
            _ = v37
            _ = v38
        }
        _ = v36
    }
    fmt.Println(v3, v35)
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v9
}
