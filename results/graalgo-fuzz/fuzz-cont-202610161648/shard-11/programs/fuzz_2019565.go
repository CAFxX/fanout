package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 bool; F2 uint32 }
func genFn0() uint16 {
    if (((3.14159 + ((0.0 - 1.5) / (0.0 + 1e10))) >= 1e10) || true) {
        v5 := []int{9223372036854775807}
        switch ("foo bar" + "hello") {
        case "foo bar":
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 bool = (!true)
                _ = fi6
                _ = v7
            }
        default:
            var v8 int32 = 42
            _ = v8
        }
        _ = v5
    } else {
        var v9 uint = 65535
        v9 -= v9
        var v10 string = ((("foo bar" + (("\n\t" + "\n\t") + ("\n\t" + ("hello" + "")))) + "hello") + (("foo bar" + "\n\t") + "foo bar"))
        _ = v10
        _ = v9
    }
    var v11 int32 = 2147483647
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS1{F0: 0}
        _ = fi12
        _ = v13
    }
    _ = v11
    return 2
}

func main() {
    v1 := []string{"a", "foo bar", "\n\t"}
    var v2 int16 = 1000
    var v3 bool = false
    var v4 bool = false
    if (("a" + "") >= "") {
        if true {
            fmt.Println(v4, v2, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        switch 32768 {
        case -9223372036854775808:
            if (v3 || (9223372036854775807 >= -1)) {
                v14 := GenS2{F0: -1, F1: 127, F2: 128}
                fmt.Println(v2, v4)
                var v15 uint64 = 4294967295
                _ = v14
                _ = v15
            } else {
                var v16 bool = v4
                switch 602 {
                case 2147483647:
                    v4 = false
                case 0:
                    if v3 {
                        switch 952 {
                        case 2147483647:
                            if v16 {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi17 := 0; fi17 < 15; fi17++ {
                                    for fi18 := 0; fi18 < 1; fi18++ {
                                        var v19 uint = 2
                                        _ = fi18
                                        _ = v19
                                    }
                                    v1 = append(v1, "a")
                                    var v20 float32 = ((float32(1.5) - float32(1e10)) * float32((3.14159 / (float64(((0.0 - 3.14159) + (1e10 + 1e10))) - -2.25))))
                                    _ = fi17
                                    _ = v20
                                }
                                v4 = ((v2 / (((v2 % (((v2 << ((639) & 63))) | 1))) | 1)) != (v2 % (((v2 | -1)) | 1)))
                            }
                        case 940:
                            switch ("" + "a") {
                            case "foo bar":
                                fmt.Println(v4, v3, v3)
                            case "\n\t":
                                fmt.Println(v16, v3)
                            case "":
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch 2147483647 {
                                    case 32768:
                                        var v21 float64 = 0.1
                                        _ = v21
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v1 = append(v1, "\n\t")
                                    default:
                                        var v22 uint8 = 42
                                        _ = v22
                                    }
                                }
                            }
                        case 1:
                            v1 = append(v1, "hello")
                        default:
                            _ = len(v1)
                        }
                    }
                }
                _ = v16
            }
        case 0:
            fmt.Println(v2, v3)
        default:
            var v23 uint = 1
            _ = v23
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            if false {
                switch (fi24 - (fi24 * fi24)) {
                case 869:
                    var v25 float64 = 0.0
                    _ = v25
                case 2147483647:
                    if true {
                        fmt.Println(v2, fi24, fi24)
                    } else {
                        fmt.Println(v3, v4, v4)
                    }
                }
            }
            v4 = v4
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint64 = 18446744073709551615
                fmt.Println(fi24)
                if v4 {
                    v28 := GenS3{F0: -23, F1: v2}
                    _ = v28
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = len(v1)
                    v4 = v3
                }
                _ = fi26
                _ = v27
            }
            _ = fi24
        }
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v29 := GenS1{F0: 175}
    var v30 float64 = 3.14159
    var v31 uint8 = 97
    for fi32 := 0; fi32 < 11; fi32++ {
        v2 -= -937
        v33 := GenS4{F0: -110}
        _ = fi32
        _ = v33
    }
    v31 = 127
    v34 := GenS5{F0: 1, F1: v4, F2: 65535}
    fmt.Println(v3)
    v2 -= 32767
    for fi35 := 0; fi35 < 11; fi35++ {
        if ((!v3) || (2 != 65535)) {
            var v36 float64 = v30
            var v37 bool = (!(true || false))
            _ = v36
            _ = v37
        } else {
            var v38 int16 = 1000
            v2 += -32768
            for fi39 := 0; fi39 < 19; fi39++ {
                var v40 string = (("foo bar" + ("\n\t" + "\n\t")) + "hello")
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v41 uint64 = 2
                v40 = v40
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v38
        }
        _ = fi35
    }
    v2 = int16(v31)
    var v42 float32 = 1e10
    fmt.Println(v31)
    v30 += ((float64(float32(3.14159)) * v30) / ((1e10 / v30) * (float64((v42 * v42)) - (v30 + 1.5))))
    var v43 uint = 943
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v42
    _ = v43
}
