package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 int8 = 51
    var v4 string = (("" + "hello") + (("a" + ("hello" + "\n\t")) + "foo bar"))
    _ = len(v1)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int = fi5
        fmt.Println(v6, v4, v4)
        _ = fi5
        _ = v6
    }
    if ((false || (float64((float32(1.5) - (float32((0.1 - (((1.5 / (1e10 + 0.1)) - ((1e10 * 1.5) * (0.1 * -2.25))) * ((((3.14159 - 1.5) - float64(1e10)) * 1.5) * 3.14159)))) + float32((-2.25 + -2.25))))) == -2.25)) || (v3 > v3)) {
        var v7 float32 = 3.14159
        _ = v7
    } else {
        v3 = (-(v3))
        var v8 uint32 = 4294967295
        switch 9223372036854775807 {
        case 4294967295:
            _ = len(v2)
        }
        _ = v8
    }
    v4 = ((v4 + ("\n\t" + ((v4 + "") + "hello"))) + ((v4 + v4) + v4))
    var v9 uint8 = 2
    switch 140 {
    case 128:
        _ = len(v2)
    default:
        delete(v2, v4)
    }
    _ = len(v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := []int{(fi11 << ((769) & 63)), (fi10 + fi11)}
            var v13 int64 = 9223372036854775807
            if (0 > 4294967295) {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            } else {
                var v14 string = ("hello" + v4)
                var v15 int8 = 42
                if (true || true) {
                    var v16 uint64 = 1
                    _ = v16
                } else {
                    switch fi10 {
                    case 769:
                        if (3.14159 != float64(0.1)) {
                            var v17 int64 = (-(v13))
                            _ = v17
                        } else {
                            var v18 float32 = -2.25
                            _ = v18
                        }
                        delete(v2, (v4 + ("\n\t" + "\n\t")))
                    default:
                        var v19 uint16 = 32767
                        _ = v19
                    }
                }
                switch v14 {
                case "hello":
                    v15 -= -128
                case "foo bar":
                    v20 := GenS1{F0: "hello", F1: 32767}
                    _ = v20
                case "":
                    var v21 int = fi10
                    _ = v21
                }
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v22 uint8 = 0
        v9 -= (v9 - v22)
        _ = fi10
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            v4 = v4
            fmt.Println(v4, fi24, v9)
            _ = fi24
        }
        v4 = v4
        _ = fi23
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v4 = v4
        _ = fi25
    }
    v4 = ""
    var v26 int16 = -1
    for fi27 := 0; fi27 < 17; fi27++ {
        if true {
            if (v26 > 32767) {
                _ = len(v1)
                switch (fi27 - fi27) {
                case 128:
                    switch fi27 {
                    case -108:
                        fmt.Println(v26, v9, v4)
                    case -1:
                        switch ("" + "\n\t") {
                        case "foo bar":
                            var v28 uint = 42
                            _ = v28
                        case "hello":
                            var v29 float32 = -2.25
                            _ = v29
                        }
                    case -658:
                        v30 := []int{-842, (fi27 | fi27), -761}
                        _ = v30
                    }
                case -9223372036854775808:
                    switch (("foo bar" + v4) + (v4 + (((v4 + "hello") + "foo bar") + ((v4 + v4) + v4)))) {
                    case "hello":
                        v9 = (v9 & v9)
                    default:
                        v3 = v3
                    }
                default:
                    v31 := GenS2{F0: 408, F1: 2, F2: 0}
                    _ = v31
                }
                v32 := GenS1{F0: "hello", F1: v26}
                _ = v32
            }
            v4 = "a"
            if false {
                if (42 <= 924) {
                    v4 = "hello"
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v26 = v26
                        v3 -= 1
                        var v34 float64 = 0.0
                        _ = fi33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 11; fi35++ {
                        if true {
                            var v36 int8 = (v3 * v3)
                            _ = v36
                        } else {
                            if false {
                                var v37 uint16 = 255
                                _ = v37
                            }
                            v3 -= v3
                        }
                        v38 := []int{}
                        v3 -= (^(v3))
                        _ = fi35
                        _ = v38
                    }
                } else {
                    v39 := []int{(fi27 << ((726) & 63)), (fi27 % ((fi27) | 1)), -619, (fi27 ^ fi27)}
                    v4 = v4
                    var v40 uint16 = 2
                    _ = v39
                    _ = v40
                }
            } else {
                fmt.Println(v3, v9, fi27)
            }
        }
        _ = fi27
    }
    v41 := GenS3{F0: v3, F1: 303, F2: 4294967295}
    v3 = (v3 << ((1) & 63))
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v9
}
