package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int = 991
    var v2 string = ""
    var v3 int32 = -89
    v4 := make(map[string]uint32)
    var v5 uint8 = 255
    var v6 uint16 = 525
    for fi7 := 0; fi7 < 19; fi7++ {
        v5 += v5
        v2 = v2
        _ = fi7
    }
    v3 = v3
    v2 = "\n\t"
    v8 := GenS1{F0: false, F1: 32767}
    if (true && (!true)) {
        fmt.Println(v6)
    } else {
        v9 := GenS2{F0: false, F1: 0, F2: -1}
        var v10 float64 = 1.5
        _ = v10
        _ = v9
    }
    v3 -= v3
    if (42 > 2) {
        if false {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 float32 = 0.0
                _ = fi11
                _ = v12
            }
            v2 = ""
        } else {
            var v13 float64 = 1e10
            v14 := []int{(v1 % ((-550) | 1)), (v1 - (v1 / (((v1 ^ v1)) | 1))), v1, v1}
            if false {
                var v15 int8 = 42
                _ = v15
            } else {
                _ = len(v14)
                delete(v4, "foo bar")
                if true {
                    fmt.Println(v1, v13, v6)
                }
            }
            _ = v13
            _ = v14
        }
        v6 = (v6 & (v6 | (v6 - (v6 | v6))))
    } else {
        var v16 int = v1
        switch (v2 + "\n\t") {
        case "":
            for fi17 := 0; fi17 < 1; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    v19 := GenS2{F0: false, F1: 2147483647, F2: v3}
                    v20 := []int{32767}
                    fmt.Println(v16, v2, fi17)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 int8 = 15
                v6 = (v6 | (v6 >> ((4294967295) & 63)))
                _ = fi17
                _ = v21
            }
        }
        fmt.Println(v5, v3)
        _ = v16
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 int16 = 42
        for fi24 := 0; fi24 < 13; fi24++ {
            if true {
                v25 := []int{(fi24 * 2147483647), (v1 % ((-951) | 1)), 32767}
                _ = v25
            } else {
                v6 = v6
            }
            var v26 int32 = v3
            if (-2.25 == 0.0) {
                v27 := v4["foo bar"]
                _ = v27
            } else {
                v26 += (v3 % ((-1) | 1))
                v28 := GenS2{F0: (!false), F1: 71, F2: (v3 - v26)}
                switch v2 {
                case "a":
                    v29 := GenS3{F0: v6, F1: 9223372036854775807}
                    _ = v29
                case "foo bar":
                    if (v5 >= (v5 | 2)) {
                        var v30 uint = 36
                        _ = v30
                    }
                case "":
                    v31 := GenS1{F0: (2 == 1), F1: 0}
                    _ = v31
                default:
                    if (v26 < (v3 * (v3 % (((v3 & -1)) | 1)))) {
                        if false {
                            var v32 uint64 = 4294967295
                            _ = v32
                        }
                        v26 = v3
                        if (true || (!true)) {
                            v33 := GenS1{F0: false, F1: v23}
                            fmt.Println(v6, v6)
                            var v34 int32 = 625
                            v6 = (v6 ^ (v6 << ((1) & 63)))
                            _ = v33
                            _ = v34
                        } else {
                            var v35 bool = ((v23 >> ((883) & 63)) < 0)
                            switch v1 {
                            case 0:
                                v1 -= fi24
                            case -307:
                                fmt.Println(v23, v2)
                            }
                            _ = v35
                        }
                        fmt.Println(v3, v26, v23)
                        fmt.Println(v3, fi24, fi24)
                    } else {
                        v36 := []int{fi22, (fi22 - (^(fi22)))}
                        _ = v36
                    }
                }
                _ = v28
            }
            if (true || (false || false)) {
                v37 := []int{4294967295, (fi22 >> ((4294967295) & 63))}
                v38 := GenS1{F0: false, F1: -1}
                var v39 int64 = -1
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi24
            _ = v26
        }
        v3 = (^(v3))
        _ = fi22
        _ = v23
    }
    var v40 string = v2
    var v41 float64 = float64(0.1)
    var v42 string = ("a" + "")
    if true {
        var v43 int = v1
        _ = v43
    } else {
        v44 := v4[(("a" + v40) + "foo bar")]
        switch -2147483648 {
        case 0:
            for fi45 := 0; fi45 < 8; fi45++ {
                v2 = ""
                var v46 uint32 = v44
                _ = fi45
                _ = v46
            }
        default:
            fmt.Println(v42, v6)
        }
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v8
}
