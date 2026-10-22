package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 uint64, p1 bool) string {
    if p1 {
        p0 -= p0
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := []int{-9223372036854775808, (fi5 % (((-(fi5))) | 1)), fi5, (fi5 >> ((687) & 63))}
            _ = fi5
            _ = v6
        }
        fmt.Println(p0)
    } else {
        switch "hello" {
        case "foo bar":
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint8 = 2
                _ = fi7
                _ = v8
            }
        case "a":
            fmt.Println(p0, p1)
        case "hello":
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 12; fi10++ {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        if false {
                            if false {
                                var v11 string = "hello"
                                switch v11 {
                                case "foo bar":
                                    var v12 uint32 = 42
                                    _ = v12
                                case "hello":
                                    var v13 bool = p1
                                    _ = v13
                                default:
                                    var v14 float64 = 0.1
                                    _ = v14
                                }
                                _ = v11
                            }
                            v15 := GenS1{F0: false, F1: -1, F2: 0}
                            _ = v15
                        }
                    case "foo bar":
                        var v16 uint32 = 65535
                        _ = v16
                    default:
                        if p1 {
                            var v17 uint16 = 32767
                            _ = v17
                        }
                    }
                    _ = fi10
                }
                var v18 int = (fi9 << ((18446744073709551615) & 63))
                _ = fi9
                _ = v18
            }
        }
        v19 := GenS2{F0: 42, F1: 32767}
        _ = v19
    }
    v20 := GenS1{F0: p1, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = v20
    return ("foo bar" + ((("" + "foo bar") + "foo bar") + ""))
}
func genFn1(p0 int32, p1 int8) int16 {
    v21 := []int{-914, -2147483648, 2147483647}
    switch 0 {
    case -590:
        if true {
            for fi22 := 0; fi22 < 1; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    v24 := GenS2{F0: 1, F1: (fi23 << ((255) & 63))}
                    p1 -= 0
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            var v25 float32 = 0.0
            _ = v25
        } else {
            var v26 float64 = -2.25
            p0 += -987
            _ = v26
        }
    default:
        v27 := GenS1{F0: false, F1: (p0 / ((p0) | 1)), F2: -872}
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 int64 = 1
        var v30 uint64 = 4294967295
        v31 := GenS2{F0: 488, F1: int(p1)}
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v21
    return -837
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 42
    v3 := []string{"foo bar", (("" + ("" + "foo bar")) + "hello"), "\n\t"}
    v4 := []uint32{241, 4294967295, (v1 % (((v1 * 42)) | 1)), v1}
    var v32 bool = (108 != 2)
    var v33 bool = v32
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 string = "\n\t"
        for fi36 := 0; fi36 < 6; fi36++ {
            fmt.Println(v1, fi36)
            v37 := GenS2{F0: (v2 | 65535), F1: (fi36 - 0)}
            v2 -= (v2 | 752)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v38 int16 = -916
            v35 = (("" + v35) + "\n\t")
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    var v39 bool = ((!(v32 && (968 > 4294967295))) || true)
    var v40 float64 = (float64(0.1) - 3.14159)
    fmt.Println(v40, v33)
    for fi41 := 0; fi41 < 20; fi41++ {
        switch "a" {
        case "":
            if len(v4) > 0 {
                v4[0] = v1
            }
        case "a":
            v42 := GenS2{F0: (^(v2)), F1: (fi41 | (fi41 - (fi41 << (((v2 / (((v2 >> ((v2) & 63))) | 1))) & 63))))}
            _ = v42
        case "\n\t":
            var v43 uint16 = 65535
            _ = v43
        default:
            if v33 {
                v44 := GenS3{F0: 1}
                _ = v44
            }
        }
        if false {
            switch (("hello" + "foo bar") + ("hello" + ("" + "a"))) {
            case "foo bar":
                switch "\n\t" {
                case "hello":
                    if (true || v32) {
                        v32 = ((-9223372036854775808 >= -454) && (false || v33))
                    } else {
                        v45 := GenS3{F0: -32768}
                        v46 := GenS3{F0: 866}
                        var v47 int64 = 9223372036854775807
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                }
            case "a":
                v48 := GenS4{F0: "a", F1: -1}
                _ = v48
            default:
                var v49 float64 = (v40 * float64(3.14159))
                _ = v49
            }
        } else {
            v32 = false
            var v50 int64 = 87
            _ = v50
        }
        _ = fi41
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v51 float64 = (((1.5 - (v40 + 3.14159)) / (float64(1.5) / float64(0.1))) * (1e10 * v40))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v51
}
