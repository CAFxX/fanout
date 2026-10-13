package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 int16 = 42
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    var v4 bool = (!(4294967295 <= -480))
    fmt.Println(v4, v1)
    if (4294967295 != 4294967295) {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(v4, v1)
            var v6 string = v2
            _ = fi5
            _ = v6
        }
        switch v2 {
        case "\n\t":
            var v7 uint8 = 0
            _ = v7
        case "a":
            if false {
                if (-2147483648 >= -1) {
                    var v8 int64 = 1
                    var v9 uint64 = 794
                    if v4 {
                        var v10 bool = v4
                        v11 := GenS1{F0: v2, F1: (v1 >> ((65535) & 63))}
                        v12 := GenS1{F0: (v2 + v2), F1: v1}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            switch (v2 + v2) {
                            case "foo bar":
                                v14 := GenS2{F0: 654, F1: 2, F2: 32767}
                                _ = v14
                            }
                            var v15 int = fi13
                            _ = fi13
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v17 := []int{(fi16 / ((fi16) | 1)), fi16, fi16, (fi16 >> ((4294967295) & 63))}
                            switch (fi16 ^ (^(fi16))) {
                            case 128:
                                fmt.Println(v8, v1)
                            default:
                                var v18 int = (fi16 % ((42) | 1))
                                _ = v18
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                    _ = v8
                    _ = v9
                }
                var v19 int8 = -1
                v19 = 42
                _ = v19
            }
        case "foo bar":
            if false {
                var v20 float32 = 0.0
                fmt.Println(v1)
                var v21 bool = (v4 || v4)
                _ = v20
                _ = v21
            } else {
                v2 = v2
                var v22 uint = 18446744073709551615
                _ = v22
            }
        default:
            if (false && (v4 && (("\n\t" + "foo bar") == "a"))) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    switch fi23 {
                    case -732:
                        fmt.Println(v2, v2, v2)
                    }
                    var v24 float32 = 3.14159
                    _ = fi23
                    _ = v24
                }
            } else {
                v4 = (v2 < v2)
                v2 = ((v2 + ("foo bar" + ((v2 + "\n\t") + v2))) + "a")
            }
        }
    } else {
        v25 := []int{9223372036854775807, 128, 123, -328}
        _ = v25
    }
    switch 0 {
    case 9223372036854775807:
        var v26 uint32 = 1
        _ = v26
    case 32767:
        var v27 uint8 = 255
        _ = v27
    default:
        v28 := []int{32767, 32767, 2147483647}
        _ = v28
    }
    if (!v4) {
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS2{F0: 516, F1: 4294967295, F2: 877}
            v4 = v4
            _ = fi29
            _ = v30
        }
    } else {
        var v31 uint64 = 18446744073709551615
        if (!true) {
            if (-2147483648 != -1) {
                v32 := v3[-1]
                var v33 uint8 = 166
                _ = v32
                _ = v33
            } else {
                if (true || false) {
                    if (234 >= 0) {
                        v34 := []int{-52}
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 int = (fi35 & fi35)
                            _ = fi35
                            _ = v36
                        }
                        _ = v34
                    } else {
                        fmt.Println(v2, v31)
                        v37 := v3[311]
                        for fi38 := 0; fi38 < 17; fi38++ {
                            delete(v3, fi38)
                            if (v1 > 1000) {
                                v1 += (v1 ^ v1)
                                switch v37 {
                                case "a":
                                    for fi39 := 0; fi39 < 12; fi39++ {
                                        if false {
                                            v1 = v1
                                            v40 := GenS3{F0: (42 >= (v31 * (v31 << ((4294967295) & 63))))}
                                            _ = v40
                                        }
                                        v41 := GenS3{F0: v4}
                                        v42 := GenS4{F0: (true || false)}
                                        _ = fi39
                                        _ = v41
                                        _ = v42
                                    }
                                case "":
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        switch fi38 {
                                        case 2147483647:
                                            var v44 uint64 = (v31 / (((^(v31))) | 1))
                                            _ = v44
                                        }
                                        _ = fi43
                                    }
                                default:
                                    switch (fi38 | (fi38 + fi38)) {
                                    case 4294967295:
                                        delete(v3, (fi38 + 9223372036854775807))
                                    case 0:
                                        fmt.Println(v31)
                                    case -682:
                                        v37 = (v2 + "foo bar")
                                    }
                                }
                                var v45 int = (fi38 * (fi38 - 1))
                                _ = v45
                            } else {
                                var v46 uint8 = 127
                                _ = v46
                            }
                            v3[fi38] = "a"
                            _ = fi38
                        }
                        _ = v37
                    }
                }
                v1 = v1
            }
            var v47 int16 = 1000
            var v48 int8 = 1
            _ = v47
            _ = v48
        } else {
            v49 := v3[476]
            _ = v49
        }
        _ = v31
    }
    v50 := GenS1{F0: v2, F1: v1}
    for fi51 := 0; fi51 < 19; fi51++ {
        v52 := []int{2147483647}
        fmt.Println(v2, v4, v4)
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        v1 -= 1
        _ = fi53
    }
    for fi54 := 0; fi54 < 13; fi54++ {
        for fi55 := 0; fi55 < 1; fi55++ {
            if true {
                fmt.Println(v2, fi55, fi54)
                v56 := GenS4{F0: true}
                var v57 bool = (188 == 0)
                _ = v56
                _ = v57
            }
            var v58 int = fi54
            var v59 float32 = 1.5
            _ = fi55
            _ = v58
            _ = v59
        }
        _ = fi54
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        v2 = "\n\t"
        var v61 int32 = 0
        _ = fi60
        _ = v61
    }
    var v62 float32 = ((-2.25 / 1.0) / 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v62
}
