package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }

func main() {
    var v1 string = ("hello" + ("foo bar" + ("\n\t" + "\n\t")))
    var v2 float64 = (float64(1e10) + (float64(1.5) / (1e10 / 0.1)))
    v3 := []int{1, 2147483647}
    var v4 float64 = (float64((-2.25 * (1e10 - 0.1))) * (float64(3.14159) - (v2 - v2)))
    v5 := []string{}
    var v6 bool = true
    v7 := GenS1{F0: 42}
    var v8 int64 = 0
    var v9 uint16 = 255
    var v10 uint = 42
    if false {
        var v11 int8 = 1
        if len(v5) > 0 {
            v5[0] = ""
        }
        v9 -= 76
        _ = v11
    }
    fmt.Println(v2, v6, v10)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 6; fi13++ {
            v14 := GenS1{F0: 1}
            v1 = v1
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: (((v1 + v1) + ("foo bar" + v1)) + v1)}
        v16 := GenS1{F0: 4294967295}
        _ = fi12
        _ = v15
        _ = v16
    }
    if v6 {
        for fi17 := 0; fi17 < 4; fi17++ {
            v1 = v1
            var v18 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
        }
    }
    var v19 uint16 = v9
    var v20 int16 = 32767
    v9 = v9
    v4 = 0.0
    var v21 string = v1
    switch 42 {
    case 128:
        if v6 {
            v9 = v9
            v5 = append(v5, ((((v1 + v21) + (v1 + v1)) + v1) + "hello"))
            var v22 uint32 = 2
            _ = v22
        }
    }
    var v23 uint16 = v9
    fmt.Println(v19, v9)
    if (v2 >= ((-2.25 + v2) / (v4 + (v2 + v2)))) {
        switch 72 {
        case 9223372036854775807:
            v9 += v23
        default:
            if (-1 == -999) {
                for fi24 := 0; fi24 < 15; fi24++ {
                    switch fi24 {
                    case 938:
                        if v6 {
                            if v6 {
                                var v25 uint16 = (^(v19))
                                _ = v25
                            }
                        }
                    case 4294967295:
                        var v26 uint64 = 4294967295
                        _ = v26
                    default:
                        if (42 >= uint32(v19)) {
                            switch ("a" + (v1 + v1)) {
                            case "foo bar":
                                if true {
                                    v9 = 42
                                    for fi27 := 0; fi27 < 6; fi27++ {
                                        if (2 != 843) {
                                            switch (fi27 ^ (^(fi27))) {
                                            case 1:
                                                if true {
                                                    if (((v10 << (((v10 >> (((v10 % (((v10 * v10)) | 1))) & 63))) & 63)) <= v10) && (true && v6)) {
                                                        var v28 int16 = v20
                                                        v20 = (v20 | (v20 | (-(v28))))
                                                        _ = v28
                                                    }
                                                    var v29 uint = v10
                                                    var v30 int16 = v20
                                                    _ = v29
                                                    _ = v30
                                                }
                                            default:
                                                v9 += v19
                                            }
                                            v21 = v1
                                            if v6 {
                                                if v6 {
                                                    fmt.Println(v1, v2)
                                                    switch v21 {
                                                    case "a":
                                                        if (v10 <= 1) {
                                                            var v31 bool = (v6 && ((float32((float64(-2.25) + 0.0)) * (1e10 + 1.5)) >= 1e10))
                                                            v21 = (((v1 + ("\n\t" + v21)) + ("" + "\n\t")) + ("" + v21))
                                                            v32 := GenS1{F0: 618}
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                    case "foo bar":
                                                        var v33 int64 = (v8 + 87)
                                                        _ = v33
                                                    default:
                                                        v21 = v1
                                                    }
                                                    var v34 float64 = v2
                                                    _ = v34
                                                }
                                                var v35 uint32 = 4294967295
                                                v20 = -415
                                                switch v1 {
                                                case "foo bar":
                                                    var v36 uint = v10
                                                    _ = v36
                                                }
                                                v19 += v23
                                                _ = v35
                                            } else {
                                                fmt.Println(v8, v10)
                                                v2 = 1.5
                                                if false {
                                                    var v37 int32 = 42
                                                    v9 = (v19 % ((v9) | 1))
                                                    _ = v37
                                                } else {
                                                    if (!(!true)) {
                                                        if v6 {
                                                            fmt.Println(v1, v19, v6)
                                                            v21 = (("" + (((v21 + v1) + (v1 + v1)) + "foo bar")) + ((v1 + v21) + ("a" + ((v1 + v21) + "a"))))
                                                            var v38 int = (fi24 / ((fi24) | 1))
                                                            _ = v38
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v39 string = v1
                                                        _ = v39
                                                    }
                                                    v6 = (!true)
                                                    v1 = "hello"
                                                }
                                            }
                                        }
                                        _ = fi27
                                    }
                                    v3 = append(v3, fi24)
                                } else {
                                    var v40 int32 = 42
                                    _ = v40
                                }
                            default:
                                v19 -= (^(v9))
                            }
                        } else {
                            if v6 {
                                var v41 int8 = 32
                                for fi42 := 0; fi42 < 20; fi42++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi42
                                }
                                switch -9223372036854775808 {
                                case 9223372036854775807:
                                    if true {
                                        v4 -= (v4 / 1e10)
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            var v44 float64 = v2
                                            v45 := GenS2{F0: v21}
                                            _ = fi43
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = len(v3)
                                    }
                                }
                                _ = v41
                            } else {
                                var v46 int = 264
                                v2 = v2
                                fmt.Println(v6, v9, v20)
                                _ = v46
                            }
                            var v47 uint = v10
                            fmt.Println(v4)
                            _ = v47
                        }
                    }
                    for fi48 := 0; fi48 < 6; fi48++ {
                        fmt.Println(v21)
                        var v49 int8 = 42
                        _ = fi48
                        _ = v49
                    }
                    if v6 {
                        switch "hello" {
                        case "foo bar":
                            var v50 float32 = 1.5
                            _ = v50
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = (-(fi24))
                            }
                        }
                        var v51 string = "hello"
                        v52 := GenS3{F0: (!v6), F1: 255}
                        _ = v51
                        _ = v52
                    } else {
                        v53 := GenS1{F0: 2}
                        _ = v53
                    }
                    _ = fi24
                }
            }
        }
    } else {
        fmt.Println(v1)
    }
    v54 := GenS4{F0: v23, F1: (^(v9)), F2: -676}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
