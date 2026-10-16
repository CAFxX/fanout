package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 bool; F2 int8 }
type GenS5 struct { F0 uint8; F1 int; F2 int64 }
type GenS6 struct { F0 bool }

func main() {
    v1 := []uint32{65535}
    var v2 int64 = 42
    var v3 uint32 = 0
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    var v6 uint8 = 1
    v7 := GenS1{F0: v2, F1: 532, F2: 4294967295}
    v6 = v6
    v2 = v2
    var v8 float32 = ((float32(1e10) / -2.25) / 1.0)
    v9 := GenS2{F0: (^(v3))}
    v3 -= (v3 + v3)
    v10 := GenS3{F0: true}
    for fi11 := 0; fi11 < 7; fi11++ {
        if (3.14159 != 1.5) {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 float64 = 0.0
                var v14 uint32 = (v3 * v3)
                v15 := GenS4{F0: 2147483647, F1: false, F2: -128}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 float32 = -2.25
            _ = v16
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        switch ((("hello" + ("" + ((("" + ("" + "a")) + "") + "a"))) + "\n\t") + ("foo bar" + ("foo bar" + ((("hello" + "") + "a") + "hello")))) {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "a":
            if (!(42 > -1)) {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v1 = append(v1, v3)
                    v19 := GenS4{F0: 638, F1: true, F2: 119}
                    v3 -= v3
                    _ = fi18
                    _ = v19
                }
            } else {
                fmt.Println(v3, v3, v6)
                if true {
                    v1 = append(v1, v3)
                } else {
                    fmt.Println(v3, v2, v8)
                    v2 = -399
                }
                if (!true) {
                    v8 = (0.1 / -2.25)
                    var v20 string = "\n\t"
                    for fi21 := 0; fi21 < 13; fi21++ {
                        if true {
                            v4[(v20 + v20)] = v3
                            v2 = v2
                            fmt.Println(v6, v8)
                        }
                        if true {
                            delete(v5, (v20 + v20))
                            if (255 < 18446744073709551615) {
                                v22 := GenS5{F0: (v6 ^ (^(v6))), F1: fi21, F2: 9223372036854775807}
                                v2 = -65
                                _ = v22
                            } else {
                                v2 += v2
                                fmt.Println(fi17, v20, fi21)
                                v2 = 192
                            }
                            switch (fi21 * (^(fi21))) {
                            case -150:
                                var v23 int32 = 2147483647
                                var v24 uint = 2
                                switch fi21 {
                                case 106:
                                    var v25 int32 = (-(v23))
                                    _ = v25
                                default:
                                    var v26 uint = v24
                                    _ = v26
                                }
                                _ = v23
                                _ = v24
                            case 857:
                                delete(v4, "foo bar")
                            default:
                                switch (fi21 % (((fi17 & 4294967295)) | 1)) {
                                case 42:
                                    var v27 float64 = -2.25
                                    _ = v27
                                case 513:
                                    if true {
                                        v20 = v20
                                        var v28 float64 = 0.1
                                        var v29 uint = 255
                                        _ = v28
                                        _ = v29
                                    }
                                case 2147483647:
                                    var v30 int64 = (^(v2))
                                    _ = v30
                                default:
                                    v31 := GenS6{F0: true}
                                    _ = v31
                                }
                            }
                        } else {
                            if true {
                                v6 = v6
                            } else {
                                var v32 uint64 = 1
                                v3 -= v3
                                v6 = uint8(v32)
                                v1 = append(v1, 65535)
                                _ = v32
                            }
                            _ = len(v5)
                            if true {
                                v6 -= (v6 << ((395) & 63))
                                var v33 uint32 = (^(v3))
                                _ = v33
                            }
                            v5["hello"] = v20
                            fmt.Println(fi21)
                            var v34 uint = 255
                            v35 := v4[v20]
                            _ = v34
                            _ = v35
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    v36 := v4[("" + "a")]
                    var v37 bool = false
                    _ = v36
                    _ = v37
                }
            }
        case "":
            fmt.Println(fi17, v6)
        }
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
