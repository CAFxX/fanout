package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint }
type GenS6 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0() uint16 {
    v4 := GenS1{F0: 738}
    var v5 int8 = 127
    var v6 float32 = 1.5
    _ = v4
    _ = v5
    _ = v6
    return 720
}
func genFn1() uint64 {
    v7 := GenS2{F0: 128, F1: 63}
    var v8 uint32 = 1
    _ = v7
    _ = v8
    return 104
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 4294967295
    v3 := []uint32{}
    v2 += 0
    v2 = 995
    var v9 int32 = -383
    switch 32767 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: (fi10 + 921)}
            if true {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v9 = -71
                    fmt.Println(fi12)
                    var v13 uint64 = 2
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 bool = (!false)
                v15 := GenS2{F0: (fi10 + (fi10 * 128)), F1: v2}
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint32 = 1
                    switch (fi10 + (fi10 ^ (fi16 % (((fi16 | -566)) | 1)))) {
                    case 752:
                        fmt.Println(v17, v14, v17)
                    default:
                        switch fi10 {
                        case 128:
                            fmt.Println(v17, fi16, fi10)
                        case 32767:
                            if v14 {
                                var v18 int8 = -1
                                v9 += (v9 | -1)
                                _ = v18
                            }
                        default:
                            v1[("foo bar" + "a")] = true
                        }
                    }
                    var v19 uint32 = v17
                    if (v2 >= 42) {
                        if (v2 > 296) {
                            var v20 bool = (v14 || (v14 || (42 == -1)))
                            fmt.Println(v19, v19, v19)
                            switch "\n\t" {
                            case "\n\t":
                                var v21 uint8 = 1
                                _ = v21
                            case "":
                                v9 -= 961
                            }
                            _ = v20
                        } else {
                            var v22 uint32 = v19
                            var v23 int8 = 0
                            delete(v1, ("a" + "\n\t"))
                            if ((v9 & v9) > 941) {
                                var v24 int64 = 9223372036854775807
                                _ = v24
                            }
                            fmt.Println(v19)
                            _ = v22
                            _ = v23
                        }
                        v25 := GenS1{F0: -2147483648}
                        _ = v25
                    }
                    _ = fi16
                    _ = v17
                    _ = v19
                }
                _ = v14
                _ = v15
            }
            var v26 uint16 = genFn0()
            _ = fi10
            _ = v11
            _ = v26
        }
    }
    fmt.Println(v9, v9)
    if (42 >= 18446744073709551615) {
        if (!false) {
            var v27 int16 = -1
            if (true || true) {
                for fi28 := 0; fi28 < 4; fi28++ {
                    v29 := GenS1{F0: 32767}
                    fmt.Println(fi28)
                    var v30 int8 = 95
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            var v31 int32 = (v9 + 188)
            _ = v27
            _ = v31
        } else {
            v9 = v9
        }
        v32 := GenS3{F0: false, F1: v9}
        _ = v32
    } else {
        delete(v1, "\n\t")
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := GenS4{F0: (v2 >> (((v2 | (^(v2)))) & 63))}
        fmt.Println(v9)
        for fi35 := 0; fi35 < 4; fi35++ {
            v36 := GenS5{F0: 65535, F1: 389, F2: 0}
            v37 := GenS6{F0: 42, F1: ("a" + (("a" + "foo bar") + "")), F2: 18446744073709551615}
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    v38 := GenS4{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v9
}
