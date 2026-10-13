package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint64) string {
    p1 += 2
    v7 := []int{-530, -825, 128}
    p1 = 207
    _ = p0
    _ = p1
    _ = v7
    return ("hello" + "\n\t")
}

func main() {
    var v1 float64 = (3.14159 / 1.0)
    var v2 float64 = -2.25
    var v3 int64 = 0
    v4 := make(map[string]string)
    v5 := []string{("a" + "a"), "\n\t"}
    v6 := []int{-294, 2147483647, -56}
    delete(v4, "")
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            if (32767 > 32767) {
                if false {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 string = ("\n\t" + ("hello" + "foo bar"))
                        if true {
                            var v11 uint16 = 2
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v10 = (v10 + "")
                            _ = v11
                        } else {
                            fmt.Println(fi8)
                        }
                        _ = fi9
                        _ = v10
                    }
                } else {
                    v12 := GenS1{F0: v3, F1: 32767}
                    _ = v12
                }
                var v13 int64 = (^(v3))
                if ("hello" >= ("" + "hello")) {
                    if false {
                        delete(v4, "foo bar")
                        var v14 int64 = v13
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 uint16 = 2
                            switch fi8 {
                            case 4294967295:
                                v13 = (v13 | (v3 + (^(v3))))
                            }
                            fmt.Println(v16, v13)
                            _ = fi15
                            _ = v16
                        }
                        v17 := GenS1{F0: -877, F1: -1}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := GenS1{F0: 42, F1: -32768}
                        fmt.Println(fi18)
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v13
            } else {
                v3 += v3
            }
            v20 := GenS2{F0: -1, F1: (0 != 0)}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 float32 = ((0.1 * 0.1) - 1.5)
                v2 += (v2 * 1.5)
                _ = len(v6)
                v23 := GenS2{F0: 17, F1: (((((0.0 - ((v2 + (v2 - 0.0)) / -2.25)) / (0.0 - 3.14159)) * 1e10) >= v2) && true)}
                fmt.Println(v2, v22, v3)
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: -601}
            _ = v24
        }
        fmt.Println(v3, v3, fi8)
        v25 := GenS2{F0: 100, F1: (true || ((693 > 65535) || ((-9 == -1) || (!false))))}
        _ = fi8
        _ = v25
    }
    v4[("\n\t" + "")] = "a"
    v3 -= -1
    var v26 bool = (511 >= -1)
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 uint32 = 65535
        var v29 int8 = -1
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 uint = 859
    for fi31 := 0; fi31 < 16; fi31++ {
        v4[("hello" + "hello")] = (("" + (("a" + ("foo bar" + "hello")) + "foo bar")) + ((("foo bar" + ("hello" + ("hello" + "hello"))) + ("" + "hello")) + ("" + "")))
        _ = fi31
    }
    var v32 int32 = 1
    if false {
        v1 = 0.0
    }
    var v33 int8 = -48
    if (!(42 < 471)) {
        var v34 int = -814
        v34 = -1
        v32 += (-(v32))
        _ = v34
    } else {
        for fi35 := 0; fi35 < 17; fi35++ {
            fmt.Println(v1, v3)
            _ = fi35
        }
        var v36 int32 = v32
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
