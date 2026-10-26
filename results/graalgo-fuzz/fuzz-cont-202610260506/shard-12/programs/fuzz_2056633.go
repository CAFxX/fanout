package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int8 }

func main() {
    var v1 string = (("\n\t" + "a") + "\n\t")
    v2 := []string{}
    var v3 uint = 4294967295
    v4 := []uint32{1, 65535}
    var v5 int8 = -128
    v3 += 2
    v3 = 2
    v1 = ((("hello" + ((v1 + "") + (v1 + v1))) + v1) + "hello")
    var v6 int8 = v5
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ("foo bar" + ("a" + "hello")) {
        case "":
            for fi8 := 0; fi8 < 17; fi8++ {
                v3 -= (^(v3))
                switch ("hello" + ((v1 + v1) + v1)) {
                case "":
                    var v9 int64 = 42
                    _ = v9
                case "a":
                    v5 = (-(v6))
                }
                v10 := GenS1{F0: 32767}
                switch (("a" + v1) + (v1 + "")) {
                case "a":
                    var v11 uint16 = 200
                    _ = v11
                }
                v12 := GenS2{F0: 306, F1: 2147483647, F2: 42}
                _ = fi8
                _ = v10
                _ = v12
            }
        case "foo bar":
            switch fi7 {
            case -2147483648:
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int = fi13
                    v15 := GenS1{F0: 32767}
                    fmt.Println(v1)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 16; fi16++ {
                    if true {
                        v1 = (v1 + ((("hello" + ("" + v1)) + ("a" + v1)) + "\n\t"))
                        switch fi16 {
                        case 32767:
                            fmt.Println(fi16, v1, v5)
                            v3 = v3
                        }
                        v17 := GenS3{F0: (v3 + (v3 * (^(v3))))}
                        _ = v17
                    } else {
                        v1 = ""
                        fmt.Println(v1, fi16, fi16)
                        fmt.Println(v3, v5, fi7)
                        v3 = 2
                        var v18 uint = (v3 ^ (^(v3)))
                        _ = v18
                    }
                    _ = fi16
                }
            }
        case "\n\t":
            if (549 < 1) {
                _ = len(v4)
                v3 -= v3
            } else {
                v6 = v5
            }
        default:
            if true {
                v3 = (v3 * 65535)
                switch (fi7 / (((fi7 | -151)) | 1)) {
                case 568:
                    if (!(-1 >= -1)) {
                        var v19 int8 = (v6 >> ((v3) & 63))
                        _ = v19
                    }
                case -48:
                    if ((false && true) && ((fi7 * fi7) != (^(fi7)))) {
                        v20 := GenS3{F0: v3}
                        if false {
                            v6 += 56
                            _ = len(v2)
                        }
                        _ = v20
                    }
                case -138:
                    switch (fi7 & fi7) {
                    case -9223372036854775808:
                        for fi21 := 0; fi21 < 12; fi21++ {
                            var v22 float32 = float32(0.1)
                            _ = fi21
                            _ = v22
                        }
                    default:
                        if true {
                            var v23 int64 = 42
                            var v24 float32 = (3.14159 - 1e10)
                            _ = v23
                            _ = v24
                        }
                    }
                }
                v5 += -59
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v1, v6)
        _ = fi7
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 uint8 = 0
        _ = fi25
        _ = v26
    }
    var v27 uint64 = 106
    fmt.Println(v3)
    fmt.Println(v1)
    v28 := GenS2{F0: 42, F1: 926, F2: 42}
    if true {
        var v29 float32 = -2.25
        _ = v29
    } else {
        v27 -= uint64(v3)
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := GenS4{F0: (^(v27)), F1: v3, F2: v6}
            switch (v1 + (("\n\t" + "hello") + v1)) {
            case "\n\t":
                if false {
                    fmt.Println(v5)
                    if true {
                        fmt.Println(v1)
                        var v32 int8 = 42
                        var v33 uint64 = 0
                        _ = v32
                        _ = v33
                    }
                    var v34 int32 = 0
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 17; fi35++ {
                        if true {
                            fmt.Println(v6)
                        }
                        _ = fi35
                    }
                }
            case "":
                var v36 uint32 = uint32(v27)
                _ = v36
            default:
                v5 += v6
            }
            _ = fi30
            _ = v31
        }
        var v37 bool = true
        _ = v37
    }
    v27 -= v27
    v4 = append(v4, 4294967295)
    var v38 int = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        var v39 int = 1
        _ = v39
    }
    switch "\n\t" {
    case "":
        var v40 string = (v1 + ("\n\t" + v1))
        _ = v40
    case "\n\t":
        for fi41 := 0; fi41 < 20; fi41++ {
            var v42 bool = (578 < v3)
            v43 := GenS1{F0: 0}
            _ = fi41
            _ = v42
            _ = v43
        }
    default:
        v44 := GenS1{F0: 32767}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
