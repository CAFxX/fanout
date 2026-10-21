package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint8 }
func genFn0(p0 int32, p1 int, p2 string) uint {
    p2 = p2
    var v7 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 118
}

func main() {
    var v1 uint16 = 644
    v2 := make(map[string]uint32)
    var v3 int8 = -128
    var v4 uint32 = 1
    v5 := make(map[string]string)
    v6 := make(map[string]string)
    var v8 string = ("a" + "a")
    if false {
        var v9 uint64 = 4294967295
        _ = v9
    }
    _ = len(v2)
    v10 := GenS1{F0: -308}
    delete(v6, (v8 + ((v8 + v8) + "hello")))
    var v11 uint8 = 48
    v12 := v5[v8]
    v8 = (v12 + (v12 + "\n\t"))
    v13 := []int{9223372036854775807, -2147483648, -2147483648}
    for fi14 := 0; fi14 < 5; fi14++ {
        v8 = (v12 + (v12 + ("a" + v12)))
        if ((1.5 + 0.1) <= 1.5) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v15 := GenS2{F0: false, F1: (^(v11))}
            v12 = v12
            _ = v15
        } else {
            switch fi14 {
            case 128:
                v13 = append(v13, (fi14 << ((2) & 63)))
            case -74:
                for fi16 := 0; fi16 < 8; fi16++ {
                    v8 = ("\n\t" + ("\n\t" + v8))
                    var v17 uint64 = 18446744073709551615
                    delete(v2, v8)
                    _ = fi16
                    _ = v17
                }
            case -1:
                var v18 bool = (!(v1 != (v1 & 225)))
                _ = v18
            default:
                v13 = append(v13, (fi14 / ((fi14) | 1)))
            }
        }
        _ = fi14
    }
    var v19 uint8 = (v11 << ((0) & 63))
    if (663 == -9223372036854775808) {
        if false {
            v4 = 2147483648
            if true {
                v20 := v2["\n\t"]
                v4 -= v4
                _ = v20
            } else {
                switch 9223372036854775807 {
                case 32767:
                    if (true && false) {
                        var v21 int16 = 1
                        if false {
                            v11 += v11
                            v5[("" + (v8 + v8))] = v12
                        } else {
                            v22 := GenS3{F0: 4294967295, F1: -1}
                            switch (("hello" + v12) + v12) {
                            case "hello":
                                v2[("foo bar" + "\n\t")] = (v4 << ((2) & 63))
                            case "foo bar":
                                _ = len(v6)
                            case "\n\t":
                                _ = len(v13)
                            }
                            _ = v22
                        }
                        v12 = (((v12 + v12) + ("hello" + v12)) + (v8 + ((v12 + v8) + ("a" + ("foo bar" + "\n\t")))))
                        _ = v21
                    }
                case 1:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v3 = (v3 ^ int8(v1))
                        if (!true) {
                            switch ("a" + v8) {
                            case "hello":
                                var v24 string = v12
                                _ = v24
                            }
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            v25 := GenS4{F0: false, F1: 0, F2: (v11 ^ (v11 >> ((1) & 63)))}
                            _ = v25
                        } else {
                            v26 := GenS1{F0: -782}
                            _ = v26
                        }
                        var v27 float64 = 0.1
                        _ = fi23
                        _ = v27
                    }
                }
                v12 = v12
                if false {
                    _ = len(v6)
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v29 := GenS3{F0: 649, F1: 9223372036854775807}
                        var v30 int64 = int64(v3)
                        for fi31 := 0; fi31 < 8; fi31++ {
                            var v32 int64 = v30
                            switch v8 {
                            case "a":
                                if false {
                                    var v33 uint32 = 499
                                    _ = v33
                                }
                            case "hello":
                                var v34 uint8 = (v11 % (((v19 + 224)) | 1))
                                _ = v34
                            }
                            if (v1 <= (v1 - v1)) {
                                var v35 int16 = 32767
                                v4 += v4
                                _ = v35
                            }
                            if ((255 == v11) && true) {
                                switch v8 {
                                case "\n\t":
                                    v36 := GenS4{F0: false, F1: 42, F2: (^(v11))}
                                    _ = v36
                                case "a":
                                    var v37 string = v8
                                    _ = v37
                                default:
                                    v38 := GenS5{F0: -469, F1: 42}
                                    _ = v38
                                }
                                if len(v13) > 0 {
                                    v13[0] = (fi28 >> ((450) & 63))
                                }
                                var v39 int16 = -1
                                v40 := GenS6{F0: (v1 == v1), F1: 0, F2: v11}
                                _ = v39
                                _ = v40
                            }
                            v41 := v6[v12]
                            _ = fi31
                            _ = v32
                            _ = v41
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    var v42 int64 = -1
                    _ = v42
                }
            }
        }
        fmt.Println(v8)
        var v43 int32 = 42
        _ = v43
    }
    v44 := GenS1{F0: -924}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v8
}
