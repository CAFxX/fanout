package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 int32 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 uint8; F1 string }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    var v1 int32 = 865
    var v2 bool = false
    v3 := []int{42, -313, 128}
    v4 := make(map[string]int)
    var v5 float64 = (1.5 / 3.14159)
    for fi6 := 0; fi6 < 16; fi6++ {
        switch fi6 {
        case 32768:
            fmt.Println(v2, fi6, fi6)
        default:
            v1 += (v1 & 43)
        }
        v2 = v2
        _ = len(v3)
        _ = fi6
    }
    switch 128 {
    case 127:
        v4[(("\n\t" + ((("a" + (("" + "foo bar") + "foo bar")) + "\n\t") + ((("a" + "a") + (("foo bar" + (("\n\t" + "") + ("a" + "\n\t"))) + "")) + "hello"))) + "\n\t")] = 128
    case 32767:
        switch 843 {
        case -1:
            switch 128 {
            case 614:
                v5 = v5
            case 127:
                var v7 int64 = -1
                _ = v7
            case -9223372036854775808:
                switch "" {
                case "foo bar":
                    var v8 int16 = -1
                    _ = v8
                }
            }
        case 9223372036854775807:
            v9 := GenS1{F0: 255, F1: v2, F2: -1}
            _ = v9
        case 233:
            if (!v2) {
                v10 := GenS2{F0: "a"}
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := v4["\n\t"]
                    var v13 float64 = float64(float32(((0.1 - v5) - float64(1.5))))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
    case 0:
        if true {
            var v14 uint32 = 65535
            v14 += 587
            v15 := GenS3{F0: (v14 & v14), F1: v2, F2: ""}
            _ = v14
            _ = v15
        } else {
            switch 308 {
            case 42:
                v16 := GenS3{F0: 0, F1: v2, F2: ("hello" + ("" + "foo bar"))}
                _ = v16
            case 32767:
                var v17 uint64 = 2
                _ = v17
            case 4294967295:
                v18 := GenS4{F0: 230}
                _ = v18
            default:
                _ = len(v4)
            }
        }
    default:
        v2 = v2
    }
    _ = len(v3)
    v19 := GenS1{F0: 127, F1: v2, F2: 42}
    var v20 int = 32767
    var v21 int8 = 1
    v22 := GenS5{F0: 2, F1: 1, F2: (v21 + (v21 | (^(v21))))}
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS5{F0: 255, F1: 204, F2: (v21 - (v21 - v21))}
            if v2 {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v21)
                if true {
                    if (true && v2) {
                        if v2 {
                            if len(v3) > 0 {
                                v3[0] = (fi23 & (^(fi24)))
                            }
                            fmt.Println(v21, v20, fi23)
                            if v2 {
                                var v26 int16 = -758
                                v20 = fi24
                                _ = v26
                            }
                            if v2 {
                                delete(v4, "")
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    v2 = true
                                case "hello":
                                    fmt.Println(v5, v20, v2)
                                case "":
                                    v27 := GenS4{F0: 127}
                                    _ = v27
                                default:
                                    switch ("foo bar" + "a") {
                                    case "":
                                        v28 := GenS6{F0: (((("" + ("a" + "hello")) + "a") + "") + (("a" + "hello") + "a")), F1: 1, F2: (v1 & (v1 + v1))}
                                        _ = v28
                                    case "\n\t":
                                        var v29 float64 = (v5 - v5)
                                        _ = v29
                                    }
                                }
                                var v30 int8 = v21
                                _ = v30
                            } else {
                                v31 := GenS2{F0: "hello"}
                                var v32 int16 = 539
                                if false {
                                    var v33 float32 = (-2.25 * (1e10 + ((-2.25 - (1e10 * (3.14159 / 3.14159))) / 0.1)))
                                    var v34 int32 = v1
                                    _ = v33
                                    _ = v34
                                }
                                _ = v31
                                _ = v32
                            }
                            var v35 uint16 = 0
                            _ = v35
                        }
                        v1 = (v1 ^ v1)
                        switch "hello" {
                        case "foo bar":
                            delete(v4, ("a" + "foo bar"))
                        case "":
                            v2 = (!((128 > 255) || v2))
                        case "\n\t":
                            _ = len(v4)
                            var v36 uint64 = 42
                            _ = v36
                        }
                    }
                    if true {
                        v5 = float64(((0.1 / float32(v5)) + 1.5))
                        v37 := GenS2{F0: ("foo bar" + "hello")}
                        v38 := GenS7{F0: "a", F1: (v1 | (v1 & 42))}
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS8{F0: 255, F1: "hello"}
                        switch (fi23 << ((65535) & 63)) {
                        case 2147483647:
                            switch (("hello" + ("a" + "a")) + "hello") {
                            case "a":
                                v40 := GenS9{F0: 998, F1: 42}
                                _ = v40
                            }
                        }
                        fmt.Println(v20, v2)
                        v5 = 0.0
                        _ = v39
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v21, fi24, v5)
                    case "foo bar":
                        switch "" {
                        case "a":
                            v41 := GenS8{F0: 1, F1: "foo bar"}
                            _ = v41
                        case "foo bar":
                            v1 -= v1
                        case "hello":
                            v5 = float64(-2.25)
                        }
                    }
                    v42 := GenS10{F0: "\n\t", F1: 42, F2: 18446744073709551615}
                    _ = v42
                } else {
                    v43 := GenS1{F0: 42, F1: (!v2), F2: 9223372036854775807}
                    _ = v43
                }
            }
            v20 = (fi23 << ((171) & 63))
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v44 string = ("\n\t" + "")
    var v45 int16 = 42
    v46 := GenS10{F0: v44, F1: 32767, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
