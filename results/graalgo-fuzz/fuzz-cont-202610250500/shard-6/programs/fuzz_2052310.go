package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int32; F2 bool }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []string{"\n\t", "", "foo bar", (("foo bar" + "hello") + ("\n\t" + "hello"))}
    var v2 uint8 = 1
    v3 := make(map[int]int)
    var v4 int = 32767
    fmt.Println(v4, v2)
    var v5 uint = 718
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 float64 = 0.0
        v5 += (v5 + v5)
        if true {
            if ((((float32(3.14159) + float32(v7)) + (1e10 / 0.1)) >= 0.1) && (((1.5 / -2.25) + (float32(float64(0.0)) - float32(1.5))) < (0.0 + (-2.25 / float32(0.1))))) {
                switch v4 {
                case -9223372036854775808:
                    var v8 int = v4
                    _ = v8
                case 42:
                    v9 := GenS1{F0: v5}
                    _ = v9
                }
                fmt.Println(v7)
            }
        } else {
            v2 = (^(v2))
            v3[32768] = (^(fi6))
            v10 := GenS1{F0: (v5 % ((2) | 1))}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, v4)
    if true {
        v11 := GenS1{F0: v5}
        v4 -= 32768
        v12 := GenS2{F0: v4, F1: -876, F2: (0.1 != 1e10)}
        _ = v11
        _ = v12
    }
    if true {
        if (4294967295 == 18446744073709551615) {
            var v13 float32 = ((3.14159 / float32(-2.25)) - 0.0)
            var v14 uint64 = 18446744073709551615
            v15 := v3[v4]
            _ = v13
            _ = v14
            _ = v15
        } else {
            v5 += 463
            if false {
                if true {
                    v5 += v5
                    switch "hello" {
                    case "a":
                        for fi16 := 0; fi16 < 19; fi16++ {
                            var v17 uint = (v5 | 18446744073709551615)
                            var v18 float64 = 1e10
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    case "foo bar":
                        v5 = 159
                    case "hello":
                        for fi19 := 0; fi19 < 8; fi19++ {
                            v20 := v3[fi19]
                            _ = fi19
                            _ = v20
                        }
                    default:
                        v21 := GenS1{F0: (v5 & (^(v5)))}
                        _ = v21
                    }
                    var v22 int16 = 1
                    _ = v22
                } else {
                    fmt.Println(v2, v2, v2)
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS1{F0: (v5 & (v5 / ((v5) | 1)))}
                        switch v4 {
                        case 127:
                            if (3.14159 == 1.5) {
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    v26 := GenS3{F0: 42, F1: -1, F2: 0}
                                    v27 := v3[128]
                                    v28 := GenS1{F0: (v5 % (((v5 & v5)) | 1))}
                                    switch (v4 | 9223372036854775807) {
                                    case 1:
                                        var v29 uint32 = 472
                                        _ = v29
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                v2 = (v2 - (^(v2)))
                                v30 := GenS4{F0: v5}
                                _ = v30
                            } else {
                                switch v4 {
                                case 127:
                                    v31 := GenS2{F0: (fi23 + (fi23 & (^(fi23)))), F1: -2147483648, F2: true}
                                    _ = v31
                                }
                                fmt.Println(v2, fi23, fi23)
                            }
                        case 4294967295:
                            _ = len(v3)
                        case 427:
                            var v32 int32 = 826
                            _ = v32
                        }
                        var v33 string = ""
                        _ = fi23
                        _ = v24
                        _ = v33
                    }
                }
                switch "foo bar" {
                case "hello":
                    var v34 int = 127
                    _ = v34
                default:
                    switch ("hello" + "\n\t") {
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "hello":
                        for fi35 := 0; fi35 < 20; fi35++ {
                            switch 0 {
                            case -434:
                                v36 := GenS4{F0: v5}
                                _ = v36
                            case 1:
                                var v37 int64 = int64(v2)
                                _ = v37
                            case -2147483648:
                                var v38 int16 = 115
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 11; fi39++ {
                                v5 = (v5 % ((v5) | 1))
                                if true {
                                    var v40 uint16 = 60
                                    switch v4 {
                                    case -66:
                                        if true {
                                            _ = len(v3)
                                            var v41 float64 = 3.14159
                                            var v42 int8 = 127
                                            _ = v41
                                            _ = v42
                                        }
                                    case 2147483647:
                                        var v43 uint32 = 1
                                        _ = v43
                                    }
                                    v4 = v4
                                    _ = v40
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        var v44 string = "foo bar"
                                        _ = v44
                                    case "a":
                                        v45 := GenS4{F0: 374}
                                        _ = v45
                                    default:
                                        var v46 string = ("\n\t" + ("foo bar" + ("a" + "foo bar")))
                                        _ = v46
                                    }
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        var v47 int = 742
                                        _ = v47
                                    case "\n\t":
                                        var v48 int64 = -596
                                        _ = v48
                                    case "foo bar":
                                        v5 -= 42
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("" + ("\n\t" + "hello"))
                                }
                                v49 := v3[42]
                                _ = fi39
                                _ = v49
                            }
                            v4 += (v4 | (^(fi35)))
                            _ = fi35
                        }
                    case "\n\t":
                        if false {
                            if false {
                                var v50 float32 = 1.5
                                v5 += (v5 >> ((42) & 63))
                                _ = v50
                            }
                        }
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v51 bool = true
            _ = v51
        }
        var v52 int8 = 42
        _ = v52
    } else {
        v2 = v2
    }
    v53 := GenS3{F0: 38, F1: 0, F2: 42}
    switch (v4 + (^(v4))) {
    case 127:
        var v54 uint32 = 396
        _ = v54
    case 2147483647:
        v4 = (v4 ^ (^(v4)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
