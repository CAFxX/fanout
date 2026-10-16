package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint8) uint8 {
    switch "" {
    case "foo bar":
        v7 := GenS1{F0: -128, F1: 1, F2: 811}
        _ = v7
    }
    var v8 uint8 = (p0 + p0)
    _ = p0
    _ = v8
    return 2
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]string)
    var v3 int8 = -11
    var v4 string = v1
    v5 := []string{v1, v4, "", v4}
    var v6 float32 = (0.1 - 0.1)
    for fi9 := 0; fi9 < 8; fi9++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        delete(v2, 2)
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 = (v3 - (-(v3)))
        var v11 int64 = 1
        var v12 uint64 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    v6 += (((v6 + (v6 + (1.5 - 3.14159))) + float32(((3.14159 - ((0.1 / 0.1) * -2.25)) - 1e10))) - -2.25)
    var v13 uint16 = 32767
    var v14 int64 = -9223372036854775808
    if true {
        _ = len(v2)
        switch "foo bar" {
        case "\n\t":
            var v15 int32 = 0
            _ = v15
        case "":
            v16 := GenS2{F0: -287}
            _ = v16
        }
        v17 := GenS3{F0: 32767}
        _ = v17
    } else {
        v4 = v4
        var v18 bool = (!(true || (v14 > v14)))
        fmt.Println(v13, v3)
        _ = v18
    }
    switch -1000 {
    case -307:
        if (0 <= 42) {
            fmt.Println(v6)
            v6 -= v6
        } else {
            var v19 int32 = 1
            for fi20 := 0; fi20 < 11; fi20++ {
                v14 += v14
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 bool = false
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            var v23 uint8 = 2
            _ = v19
            _ = v23
        }
    case 32768:
        var v24 uint8 = 128
        _ = v24
    }
    v25 := GenS1{F0: v3, F1: v3, F2: 0}
    if (0 >= v3) {
        var v26 uint8 = 0
        v4 = v4
        _ = v26
    } else {
        var v27 bool = (!true)
        _ = v27
    }
    v28 := GenS3{F0: 914}
    v1 = ("" + "hello")
    var v29 int8 = -1
    var v30 int64 = (v14 ^ v14)
    switch -876 {
    case 1:
        var v31 uint16 = (v13 % ((v13) | 1))
        _ = v31
    }
    if (v14 >= (v14 * (v30 ^ v14))) {
        if false {
            for fi32 := 0; fi32 < 2; fi32++ {
                var v33 float32 = v6
                fmt.Println(v1, v3, v6)
                _ = fi32
                _ = v33
            }
            v34 := GenS2{F0: 42}
            _ = v34
        }
        if true {
            v5 = append(v5, "foo bar")
        }
    } else {
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS4{F0: (v30 / ((v14) | 1))}
            v6 = 0.0
            switch (v1 + v1) {
            case "\n\t":
                var v37 bool = false
                _ = v37
            case "hello":
                v1 = v1
            case "a":
                for fi38 := 0; fi38 < 20; fi38++ {
                    if ((fi35 >> ((561) & 63)) < fi38) {
                        if (!true) {
                            v3 = v3
                            if (true || true) {
                                if false {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if false {
                                        var v39 uint = 0
                                        if len(v5) > 0 {
                                            v5[0] = v1
                                        }
                                        fmt.Println(v13, v6)
                                        var v40 uint16 = (v13 / ((v13) | 1))
                                        _ = v39
                                        _ = v40
                                    }
                                    v41 := GenS2{F0: (-(fi38))}
                                    _ = v41
                                }
                                fmt.Println(fi38, v1, v4)
                            }
                            var v42 uint64 = 2
                            _ = v42
                        }
                        var v43 uint32 = 237
                        _ = v43
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi38
                }
            }
            _ = fi35
            _ = v36
        }
        var v44 uint8 = 59
        _ = v44
    }
    v4 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
