package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 uint64) uint8 {
    p0 += 822
    p0 = p0
    _ = p0
    return 86
}
func genFn1(p0 int, p1 uint8, p2 int64) uint {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 1000, F1: -241}
        if true {
            p1 = 255
        }
        v8 := GenS1{F0: 1000, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 = p1
    var v9 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 159
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"foo bar", ("\n\t" + "\n\t"), "\n\t", "\n\t"}
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 bool = (1e10 == float32(float64((3.14159 / float32(3.14159)))))
    for fi10 := 0; fi10 < 18; fi10++ {
        v5 = (fi10 > 42)
        _ = fi10
    }
    v5 = (!(!true))
    _ = len(v1)
    var v11 int16 = 32767
    if v5 {
        fmt.Println(v5)
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                if ((v11 + v11) > v11) {
                    switch (fi13 ^ fi13) {
                    case -124:
                        v11 = (v11 % (((v11 ^ (v11 & v11))) | 1))
                    case -597:
                        v5 = v5
                    }
                } else {
                    fmt.Println(v5, v5)
                    switch 0 {
                    case 128:
                        v14 := GenS2{F0: -1, F1: -2147483648, F2: "foo bar"}
                        _ = v14
                    case 127:
                        switch "hello" {
                        case "":
                            v2 = append(v2, (("hello" + (("a" + "hello") + "foo bar")) + "foo bar"))
                        case "foo bar":
                            var v15 uint16 = 813
                            _ = v15
                        default:
                            v5 = (!false)
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                if (true && (4294967295 != 0)) {
                                    var v16 float64 = 0.0
                                    fmt.Println(v11, v5)
                                    _ = v16
                                }
                            case "\n\t":
                                var v17 float32 = (0.1 - -2.25)
                                _ = v17
                            }
                        }
                    }
                }
                fmt.Println(fi13, v11)
                _ = fi13
            }
            _ = fi12
        }
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            v11 -= (v11 - (v11 | v11))
            for fi19 := 0; fi19 < 18; fi19++ {
                v1[32768] = 934
                var v20 bool = v5
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        if v5 {
            for fi21 := 0; fi21 < 2; fi21++ {
                if (!(1 != 255)) {
                    fmt.Println(fi21, v11)
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint16 = 2
                        switch "foo bar" {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        case "":
                            v24 := GenS1{F0: (v11 * v11), F1: 42}
                            _ = v24
                        case "foo bar":
                            if (("a" + "foo bar") != "hello") {
                                if v5 {
                                    if len(v2) > 0 {
                                        v2[0] = "hello"
                                    }
                                }
                            }
                        default:
                            if true {
                                v4["\n\t"] = fi22
                            }
                        }
                        delete(v3, (fi22 << ((0) & 63)))
                        var v25 bool = true
                        v25 = (!false)
                        _ = fi22
                        _ = v23
                        _ = v25
                    }
                } else {
                    v26 := GenS2{F0: -71, F1: 1, F2: "hello"}
                    _ = v26
                }
                _ = fi21
            }
        }
        if (!v5) {
            var v27 int64 = -9223372036854775808
            _ = v27
        }
    }
    v2 = append(v2, "")
    v28 := GenS1{F0: -1, F1: (v11 & v11)}
    if true {
        var v29 string = ("" + "a")
        v30 := GenS3{F0: 375, F1: 2147483647, F2: 238}
        v2 = append(v2, "foo bar")
        _ = v29
        _ = v30
    } else {
        fmt.Println(v11, v5)
        for fi31 := 0; fi31 < 4; fi31++ {
            v11 = (v11 % ((v11) | 1))
            v32 := GenS4{F0: 1, F1: 2147483647, F2: genFn1(fi31, genFn0(990), -689)}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
            _ = v32
        }
    }
    v5 = false
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
