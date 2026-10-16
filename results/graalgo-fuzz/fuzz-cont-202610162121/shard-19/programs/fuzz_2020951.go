package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 int }
func genFn0(p0 int32) uint64 {
    var v5 uint32 = 0
    v5 += (^(v5))
    var v6 int32 = p0
    v7 := []int{32767, 127}
    v8 := GenS1{F0: 255, F1: (v6 + -132)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 uint64 = 0
    var v3 int8 = 42
    var v4 uint64 = (v2 - v2)
    v1 = append(v1, -554)
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 float32 = float32(-2.25)
        _ = fi9
        _ = v10
    }
    if true {
        var v11 uint8 = 42
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                switch "a" {
                case "hello":
                    v2 = (v2 << ((255) & 63))
                }
                _ = fi12
            }
            fmt.Println(v3, v11, v11)
        } else {
            var v13 uint8 = v11
            v4 += (v2 & genFn0(-1))
            _ = v13
        }
        var v14 uint64 = (v4 / ((uint64(v11)) | 1))
        _ = v11
        _ = v14
    } else {
        v4 = v2
    }
    var v15 float32 = 1e10
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS1{F0: 215, F1: -1}
        switch fi16 {
        case 42:
            switch (^(fi16)) {
            case 32768:
                var v18 uint64 = 727
                _ = v18
            case -1:
                if false {
                    var v19 uint = 4294967295
                    fmt.Println(v2)
                    _ = v19
                }
            case -2147483648:
                v2 = v2
            }
        case 0:
            var v20 uint = 65535
            _ = v20
        default:
            switch fi16 {
            case -450:
                v21 := GenS1{F0: 2, F1: 162}
                _ = v21
            case 4294967295:
                for fi22 := 0; fi22 < 19; fi22++ {
                    _ = len(v1)
                    v3 = (v3 * 127)
                    var v23 uint = 18446744073709551615
                    _ = fi22
                    _ = v23
                }
            }
        }
        _ = fi16
        _ = v17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            _ = len(v1)
            var v26 uint = 62
            if (2147483648 <= 4294967295) {
                v15 -= 3.14159
                switch (("a" + ("\n\t" + ("\n\t" + "foo bar"))) + "\n\t") {
                case "":
                    switch "foo bar" {
                    case "\n\t":
                        v2 = 42
                    case "foo bar":
                        v27 := GenS2{F0: -74}
                        _ = v27
                    case "":
                        switch fi25 {
                        case 2147483647:
                            v2 += v4
                        case 663:
                            v28 := GenS1{F0: 123, F1: 2147483647}
                            _ = v28
                        default:
                            v29 := GenS2{F0: 119}
                            if true {
                                if (!false) {
                                    var v30 uint32 = 4294967295
                                    var v31 int32 = 0
                                    v1 = append(v1, -635)
                                    _ = v30
                                    _ = v31
                                }
                                var v32 float64 = 0.1
                                _ = v32
                            }
                            _ = v29
                        }
                    }
                case "foo bar":
                    var v33 uint32 = 2147483648
                    _ = v33
                default:
                    var v34 float32 = v15
                    _ = v34
                }
                var v35 uint = 65535
                _ = v35
            }
            v36 := GenS1{F0: 228, F1: 125}
            switch "a" {
            case "\n\t":
                if (-9223372036854775808 <= int64(v26)) {
                    switch 32768 {
                    case 2147483647:
                        if (!true) {
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                var v37 float64 = 0.0
                                _ = v37
                            default:
                                if len(v1) > 0 {
                                    v1[0] = int(v3)
                                }
                            }
                            v26 -= (v26 % ((623) | 1))
                            v3 = (v3 - v3)
                        }
                    default:
                        var v38 uint64 = 1
                        _ = v38
                    }
                    var v39 float32 = v15
                    _ = v39
                } else {
                    v26 = 2
                    v4 = 1
                    v3 = v3
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                fmt.Println(v2, v4)
            case "hello":
                var v40 int32 = -1
                _ = v40
            }
            _ = fi25
            _ = v26
            _ = v36
        }
        var v41 int32 = -2147483648
        _ = fi24
        _ = v41
    }
    v2 += genFn0(1)
    v42 := GenS2{F0: v3}
    var v43 int8 = (v3 ^ v3)
    v2 = (v2 >> ((65535) & 63))
    v44 := GenS2{F0: v3}
    v45 := GenS3{F0: 1, F1: "\n\t", F2: -9223372036854775808}
    var v46 string = (("" + "hello") + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
