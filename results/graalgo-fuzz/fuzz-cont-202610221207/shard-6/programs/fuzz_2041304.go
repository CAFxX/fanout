package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int16; F1 int; F2 int32 }
type GenS3 struct { F0 string; F1 int; F2 int32 }

func main() {
    v1 := []string{}
    var v2 int16 = 42
    var v3 int32 = 0
    v4 := []uint32{4294967295, 65535, 42}
    v5 := GenS1{F0: 1, F1: (v3 * (v3 % (((v3 - v3)) | 1)))}
    var v6 string = "hello"
    fmt.Println(v2, v2)
    if (1 <= 42) {
        if false {
            if true {
                v6 = "a"
                if len(v4) > 0 {
                    v4[0] = 24
                }
                v2 = v2
            } else {
                if (!true) {
                    for fi7 := 0; fi7 < 11; fi7++ {
                        v2 += 32767
                        _ = fi7
                    }
                    v4 = append(v4, 692)
                    v3 -= v3
                } else {
                    if false {
                        var v8 int = 127
                        v1 = append(v1, v6)
                        _ = v8
                    }
                    fmt.Println(v6)
                }
            }
            if false {
                v6 = ("" + "a")
                v3 += -267
            } else {
                for fi9 := 0; fi9 < 13; fi9++ {
                    if false {
                        v10 := GenS2{F0: (v2 >> ((42) & 63)), F1: fi9, F2: v3}
                        _ = v10
                    }
                    v3 -= -1
                    v3 = 1
                    _ = fi9
                }
                fmt.Println(v2, v6)
            }
        }
        fmt.Println(v2)
    } else {
        v6 = (("" + (v6 + v6)) + v6)
        if true {
            var v11 float32 = 0.0
            var v12 int64 = -463
            v6 = v6
            _ = v11
            _ = v12
        } else {
            if true {
                if len(v4) > 0 {
                    v4[0] = 815
                }
                if ((0.1 / (float32(1e10) / 0.1)) != float32(((0.0 - 0.0) - 1e10))) {
                    v6 = v6
                    v13 := GenS3{F0: v6, F1: -1, F2: v3}
                    var v14 uint16 = 1
                    _ = v13
                    _ = v14
                } else {
                    var v15 float64 = -2.25
                    fmt.Println(v6)
                    _ = v15
                }
            } else {
                v16 := GenS2{F0: v2, F1: -71, F2: 2147483647}
                _ = v16
            }
        }
    }
    var v17 int16 = (v2 | v2)
    var v18 uint16 = 751
    v3 = v3
    if true {
        if true {
            v1 = append(v1, "a")
        } else {
            v6 = ("hello" + v6)
            var v19 uint = 979
            if true {
                for fi20 := 0; fi20 < 9; fi20++ {
                    fmt.Println(v17, v19)
                    v18 = v18
                    _ = fi20
                }
                var v21 uint16 = v18
                _ = v21
            } else {
                fmt.Println(v19, v18)
                var v22 uint64 = 18446744073709551615
                v22 -= v22
                _ = v22
            }
            _ = v19
        }
        if (true || ("" == (("\n\t" + v6) + v6))) {
            var v23 string = ((v6 + "") + (((v6 + "a") + (v6 + (v6 + "\n\t"))) + "hello"))
            _ = v23
        } else {
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 bool = true
                    _ = fi25
                    _ = v26
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    v2 = v17
                    switch v6 {
                    case "foo bar":
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                switch -2147483648 {
                                case 4294967295:
                                    var v28 string = "hello"
                                    _ = v28
                                case 128:
                                    v1 = append(v1, "hello")
                                }
                                var v29 int16 = v2
                                var v30 uint8 = 42
                                _ = v29
                                _ = v30
                            }
                            var v31 int8 = 42
                            fmt.Println(v3, v3, v31)
                            v17 = (v17 - v2)
                            var v32 uint64 = 4294967295
                            v18 -= (v18 % ((v18) | 1))
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v33 float64 = 0.1
                                if false {
                                    v34 := GenS3{F0: v6, F1: (fi24 >> ((1) & 63)), F2: (v3 - 42)}
                                    var v35 uint32 = 864
                                    _ = v34
                                    _ = v35
                                }
                                _ = v33
                            }
                            _ = v31
                            _ = v32
                        }
                    }
                    var v36 float64 = (1e10 / ((((0.0 * -2.25) - -2.25) / (1e10 * (1.5 + 0.0))) / float64(1.5)))
                    _ = fi27
                    _ = v36
                }
                _ = fi24
            }
            v18 -= (v18 | (v18 - v18))
            fmt.Println(v17, v6, v2)
        }
    } else {
        v37 := GenS1{F0: 0, F1: (^(v3))}
        var v38 uint64 = 784
        _ = v37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
