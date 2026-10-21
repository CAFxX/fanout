package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 bool }
func genFn0(p0 string) uint {
    switch 1 {
    case 603:
        for fi4 := 0; fi4 < 19; fi4++ {
            v5 := GenS1{F0: "foo bar", F1: 32767}
            switch (p0 + ("" + p0)) {
            case "hello":
                p0 = "hello"
            case "a":
                switch p0 {
                case "\n\t":
                    var v6 int64 = -9223372036854775808
                    _ = v6
                default:
                    if false {
                        if true {
                            if false {
                                if (!false) {
                                    p0 = p0
                                    v7 := []int{fi4, (fi4 * fi4), 69, fi4}
                                    var v8 uint = 0
                                    _ = v7
                                    _ = v8
                                }
                                var v9 uint = 11
                                v10 := []int{283, -429, fi4, fi4}
                                _ = v10
                                _ = v9
                            }
                        }
                        switch p0 {
                        case "\n\t":
                            var v11 uint8 = 1
                            _ = v11
                        case "":
                            p0 = ("" + p0)
                        default:
                            v12 := GenS2{F0: (fi4 / ((fi4) | 1))}
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 15; fi13++ {
                            if (4294967295 > 33) {
                                if (false || (!(606 >= 0))) {
                                    p0 = p0
                                }
                            }
                            v14 := []int{(-(fi13))}
                            p0 = p0
                            _ = fi13
                            _ = v14
                        }
                    }
                }
            case "\n\t":
                v15 := GenS3{F0: 636, F1: -1, F2: (p0 + "hello")}
                _ = v15
            default:
                v16 := []int{}
                _ = v16
            }
            _ = fi4
            _ = v5
        }
    case -9223372036854775808:
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 string = (p0 + "\n\t")
                p0 = v19
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    case 4294967295:
        var v20 int32 = -737
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 float32 = 0.0
        _ = fi21
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 8; fi23++ {
            fmt.Println(fi23, fi23)
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := []int{fi24, fi23, (fi24 + 9223372036854775807)}
                if ((false && (!(!false))) || true) {
                    var v26 float32 = 3.14159
                    _ = v26
                } else {
                    var v27 int = 9223372036854775807
                    var v28 uint64 = 18446744073709551615
                    var v29 uint8 = 255
                    v30 := GenS4{F0: 9223372036854775807, F1: 106}
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi24
                _ = v25
            }
            if true {
                if (481 == 0) {
                    p0 = (p0 + "")
                }
            } else {
                var v31 uint32 = 0
                v32 := []int{32767, fi23}
                var v33 int8 = 42
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi23
        }
        var v34 bool = (1.5 <= 1e10)
        _ = v34
    } else {
        v35 := []int{}
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        _ = v35
    }
    _ = p0
    return 2
}

func main() {
    var v1 int16 = 952
    v2 := []int{}
    v3 := []uint32{4294967295, 2147483647, 0}
    if (false && (3.14159 < 0.1)) {
        v1 = v1
    }
    var v36 string = "a"
    if (!(-128 < 0)) {
        v1 = (^(v1))
        v1 = v1
    } else {
        v3 = append(v3, 0)
    }
    var v37 string = ((v36 + "") + "a")
    if false {
        if len(v3) > 0 {
            v3[0] = 65535
        }
    }
    fmt.Println(v36)
    v36 = v36
    v1 = -32768
    var v38 uint16 = 255
    var v39 uint8 = 0
    if false {
        if true {
            v36 = v36
        } else {
            var v40 int16 = (-(v1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v3)
            _ = v40
        }
    }
    v37 = "hello"
    v38 += v38
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := GenS5{F0: true}
        var v43 int16 = v1
        for fi44 := 0; fi44 < 6; fi44++ {
            v36 = "a"
            _ = fi44
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    v3 = append(v3, 1)
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 uint32 = 1
        _ = fi45
        _ = v46
    }
    v47 := GenS5{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v47
}
