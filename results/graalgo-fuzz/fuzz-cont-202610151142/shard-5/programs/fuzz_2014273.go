package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32, p2 string) uint8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 2}
        v9 := GenS2{F0: true, F1: fi7, F2: 2}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        fmt.Println(p0, p0)
        for fi10 := 0; fi10 < 9; fi10++ {
            if (-390 > 42) {
                if (p0 <= p0) {
                    v11 := GenS1{F0: 2}
                    var v12 uint8 = 128
                    var v13 int = (fi10 >> ((18446744073709551615) & 63))
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(fi10, p0, p2)
                    var v14 int16 = 32767
                    var v15 int32 = 1
                    _ = v14
                    _ = v15
                }
                switch 32768 {
                case 32767:
                    if (!false) {
                        var v16 string = "foo bar"
                        for fi17 := 0; fi17 < 7; fi17++ {
                            switch "hello" {
                            case "hello":
                                if (2 <= 22) {
                                    var v18 int8 = 8
                                    if (((int16(fi10) > -1) && (!false)) && (v16 < p2)) {
                                        var v19 int16 = -1
                                        v19 -= v19
                                        _ = v19
                                    } else {
                                        var v20 int = fi17
                                        p2 = p2
                                        _ = v20
                                    }
                                    if (42 <= -1) {
                                        var v21 uint16 = 32767
                                        _ = v21
                                    }
                                    _ = v18
                                } else {
                                    p1 = (p1 / (((p1 * p1)) | 1))
                                    var v22 int16 = 1000
                                    _ = v22
                                }
                            }
                            v23 := GenS3{F0: true, F1: 99, F2: false}
                            if (501 >= 728) {
                                v24 := GenS1{F0: 2}
                                p2 = ("foo bar" + p2)
                                v16 = "a"
                                _ = v24
                            } else {
                                var v25 int64 = 212
                                fmt.Println(p2)
                                var v26 uint = (^(p0))
                                _ = v25
                                _ = v26
                            }
                            var v27 uint16 = 32767
                            var v28 int32 = 655
                            _ = fi17
                            _ = v23
                            _ = v27
                            _ = v28
                        }
                        p0 = p0
                        _ = v16
                    } else {
                        p1 += p1
                    }
                default:
                    fmt.Println(p2, p2)
                }
            } else {
                fmt.Println(p0)
                switch ((p2 + "\n\t") + p2) {
                case "\n\t":
                    var v29 int8 = -48
                    _ = v29
                default:
                    if (993 <= 212) {
                        p1 += 2
                        var v30 uint8 = 42
                        switch fi10 {
                        case 128:
                            v31 := GenS4{F0: 100}
                            _ = v31
                        default:
                            v32 := []int{(fi10 >> ((65535) & 63))}
                            _ = v32
                        }
                        _ = v30
                    }
                }
                p0 += p0
            }
            p0 = 65535
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 uint = 617
    v2 := make(map[string]int)
    var v3 uint64 = uint64(v1)
    v4 := []uint32{}
    v5 := make(map[int]string)
    v6 := []int{}
    v1 = (v1 + 65535)
    var v33 int32 = 42
    _ = len(v4)
    v34 := v2[("\n\t" + ("" + "\n\t"))]
    if len(v4) > 0 {
        v4[0] = 0
    }
    v3 = (v3 ^ (v3 + 412))
    if ("hello" != "hello") {
        if ((false && true) || false) {
            if ("a" != ("hello" + ("" + "foo bar"))) {
                var v35 uint32 = 2
                v35 = (v35 / (((v35 | 2)) | 1))
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            if true {
                delete(v2, (("foo bar" + "hello") + "hello"))
            } else {
                var v37 int8 = -128
                _ = v37
            }
            var v38 int8 = -7
            _ = fi36
            _ = v38
        }
        switch -717 {
        case 42:
            var v39 int8 = -1
            _ = v39
        default:
            var v40 uint = v1
            _ = v40
        }
    } else {
        v34 -= -1
    }
    v1 = v1
    v1 -= v1
    var v41 int32 = v33
    var v42 uint8 = 127
    var v43 uint = v1
    if false {
        v42 = v42
        var v44 int64 = 1
        var v45 int16 = -32768
        _ = v44
        _ = v45
    } else {
        delete(v2, "hello")
        fmt.Println(v3, v33, v33)
    }
    fmt.Println(v33, v34)
    if len(v4) > 0 {
        v4[0] = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
