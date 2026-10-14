package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 bool; F2 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    v1 := []string{"a"}
    var v2 string = ""
    var v3 int64 = -9223372036854775808
    var v4 int32 = -2147483648
    var v5 int16 = 32767
    v4 = (v4 % (((v4 - 42)) | 1))
    v6 := GenS1{F0: 0}
    v2 = v2
    v2 = ((v2 + "") + ("\n\t" + "\n\t"))
    switch 32768 {
    case 128:
        _ = len(v1)
    case 9223372036854775807:
        for fi7 := 0; fi7 < 15; fi7++ {
            if len(v1) > 0 {
                v1[0] = (v2 + v2)
            }
            switch (fi7 * (fi7 ^ fi7)) {
            case -471:
                var v8 string = v2
                _ = v8
            case 128:
                v1 = append(v1, (v2 + ("foo bar" + "\n\t")))
            case 1:
                v5 += v5
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: 2}
                if (!true) {
                    if true {
                        if false {
                            if (false && false) {
                                fmt.Println(v5, v2, v3)
                                _ = len(v1)
                                var v11 uint16 = 0
                                _ = v11
                            } else {
                                var v12 float64 = 1.5
                                v1 = append(v1, (v2 + v2))
                                fmt.Println(fi9)
                                _ = v12
                            }
                            switch v2 {
                            case "":
                                switch v2 {
                                case "a":
                                    v13 := GenS1{F0: 0}
                                    _ = v13
                                }
                            }
                        }
                        v2 = v2
                    }
                }
                v4 = v4
                _ = fi9
                _ = v10
            }
            _ = fi7
        }
    case 32767:
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint64 = 126
            _ = fi14
            _ = v15
        }
    default:
        if true {
            var v16 int32 = v4
            var v17 int16 = 1000
            _ = v16
            _ = v17
        } else {
            var v18 int = -420
            v2 = v2
            fmt.Println(v2)
            _ = v18
        }
    }
    fmt.Println(v5, v5, v4)
    switch 32768 {
    case 1:
        if len(v1) > 0 {
            v1[0] = v2
        }
    case -633:
        fmt.Println(v2)
    default:
        var v19 int = -703
        _ = v19
    }
    v2 = "\n\t"
    if (42 == 718) {
        v3 = v3
        var v20 uint32 = 2
        for fi21 := 0; fi21 < 19; fi21++ {
            if true {
                for fi22 := 0; fi22 < 2; fi22++ {
                    if false {
                        switch fi21 {
                        case 1:
                            v5 += v5
                        case 0:
                            var v23 int8 = -128
                            _ = v23
                        }
                    } else {
                        v24 := GenS1{F0: 42}
                        _ = v24
                    }
                    _ = fi22
                }
                if (595 == 65535) {
                    var v25 uint64 = 0
                    _ = v25
                }
                var v26 uint = 255
                _ = v26
            }
            _ = fi21
        }
        _ = v20
    }
    v27 := GenS2{F0: 4294967295, F1: (v2 < "hello"), F2: v4}
    fmt.Println(v2, v2)
    v28 := GenS2{F0: 128, F1: (-9223372036854775808 > v3), F2: v4}
    var v29 int64 = (v3 << ((128) & 63))
    v30 := GenS2{F0: 406, F1: true, F2: (v4 ^ int32(v29))}
    if (false || false) {
        v3 -= -9223372036854775808
        v3 = (v3 % ((v3) | 1))
        switch "foo bar" {
        case "":
            for fi31 := 0; fi31 < 18; fi31++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v1)
                _ = fi31
            }
        default:
            if (!(v4 != (v4 / ((-1) | 1)))) {
                switch ("hello" + "") {
                case "hello":
                    v32 := GenS1{F0: 65535}
                    _ = v32
                case "a":
                    var v33 uint = 0
                    _ = v33
                default:
                    v5 = (v5 | (v5 % (((v5 >> ((0) & 63))) | 1)))
                }
            }
        }
    } else {
        v34 := GenS3{F0: 1, F1: 42, F2: 316}
        fmt.Println(v2)
        var v35 string = "hello"
        _ = v34
        _ = v35
    }
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
