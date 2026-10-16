package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 uint64 = 820
    var v2 uint32 = 0
    var v3 string = ""
    v4 := make(map[string]string)
    var v5 string = "hello"
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            v1 -= 42
            fmt.Println(fi7, v3)
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: v2}
                fmt.Println(fi8, fi8)
                var v10 uint = 0
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{fi7, (fi7 + fi7), (fi7 + 2147483647), (fi7 << ((0) & 63))}
            _ = v11
        }
        if true {
            var v12 uint32 = (v2 & (v2 % (((v2 * v2)) | 1)))
            fmt.Println(v5, v1, v3)
            _ = v12
        }
        _ = fi7
    }
    if false {
        v3 = ("foo bar" + (v3 + "hello"))
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int16 = -109
        v15 := GenS1{F0: (v2 | 543)}
        var v16 float64 = 0.1
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float32 = -2.25
        v19 := []int{}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := []int{268, -9223372036854775808, 32767, 32768}
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 18; fi22++ {
            v3 = v3
            _ = fi22
        }
        v23 := GenS1{F0: v2}
        _ = fi21
        _ = v23
    }
    var v24 uint64 = v1
    v5 = v5
    if false {
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                v1 -= v24
                v1 -= v24
                _ = fi26
            }
            for fi27 := 0; fi27 < 6; fi27++ {
                if (!false) {
                    var v28 uint32 = (v2 * (v2 >> ((65535) & 63)))
                    _ = v28
                }
                _ = fi27
            }
            if false {
                _ = len(v6)
            } else {
                v29 := GenS1{F0: (v2 >> ((615) & 63))}
                v24 += v24
                _ = v29
            }
            _ = fi25
        }
    } else {
        fmt.Println(v24, v1)
        for fi30 := 0; fi30 < 10; fi30++ {
            delete(v6, (^(v2)))
            if ((((^(v2)) != v2) && (v1 > v1)) || false) {
                var v31 int8 = -1
                if (!false) {
                    v24 = v1
                    if false {
                        if true {
                            fmt.Println(v5, fi30)
                            if (!true) {
                                v32 := GenS1{F0: (v2 * (v2 | 442))}
                                v1 = v24
                                _ = v32
                            } else {
                                _ = len(v20)
                                fmt.Println(v5)
                            }
                            if true {
                                var v33 int64 = -445
                                v3 = ""
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    v2 = (v2 - (v2 >> ((42) & 63)))
                                    v1 = v1
                                    switch ("foo bar" + v5) {
                                    case "a":
                                        v35 := v6[v2]
                                        _ = v35
                                    case "":
                                        if (!true) {
                                            v33 = v33
                                            v36 := GenS2{F0: -462}
                                            _ = v36
                                        }
                                    }
                                    _ = fi34
                                }
                                _ = v33
                            }
                        }
                        v24 += v1
                    }
                } else {
                    v37 := GenS2{F0: 726}
                    _ = v37
                }
                _ = v31
            } else {
                v38 := GenS3{F0: v5, F1: "foo bar"}
                if (!true) {
                    _ = len(v6)
                }
                v3 = "foo bar"
                _ = v38
            }
            var v39 int8 = 42
            _ = fi30
            _ = v39
        }
    }
    var v40 int8 = 127
    for fi41 := 0; fi41 < 19; fi41++ {
        var v42 bool = ((v2 == v2) && (v40 == v40))
        _ = fi41
        _ = v42
    }
    v43 := GenS1{F0: (v2 & v2)}
    v3 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
