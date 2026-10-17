package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int) bool {
    v5 := []int{}
    switch (p0 % ((p0) | 1)) {
    case -250:
        fmt.Println(p0)
    default:
        v6 := GenS1{F0: 42}
        _ = v6
    }
    var v7 int32 = -707
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v7
    return false
}

func main() {
    v1 := []int{9223372036854775807, -9223372036854775808, 0}
    var v2 int8 = 127
    var v3 uint32 = 2147483647
    v4 := make(map[string]string)
    fmt.Println(v2, v3)
    var v10 int8 = (v2 / ((v2) | 1))
    var v11 int = 4294967295
    var v12 uint = 65
    v13 := GenS1{F0: v12}
    if true {
        v14 := GenS1{F0: (^(v12))}
        _ = v14
    }
    var v15 float32 = 1e10
    if ((true && true) || (-1 >= -1)) {
        fmt.Println(v12, v11, v12)
    }
    v16 := GenS1{F0: v12}
    if (0 <= 18446744073709551615) {
        fmt.Println(v2, v2)
        switch "" {
        case "a":
            v3 = (v3 + v3)
        }
        var v17 uint16 = 886
        _ = v17
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            switch int(v3) {
            case -969:
                if true {
                    var v19 float32 = ((((v15 / v15) / (0.0 + v15)) + v15) / 1e10)
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 int64 = -856
                        delete(v4, "foo bar")
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            case -9223372036854775808:
                var v22 uint64 = 754
                _ = v22
            case 42:
                v23 := v4[("a" + "foo bar")]
                _ = v23
            }
            switch fi18 {
            case 9223372036854775807:
                var v24 string = "\n\t"
                _ = v24
            case 1:
                fmt.Println(v10, v10)
            case -919:
                v25 := GenS1{F0: (v12 ^ (v12 % (((v12 | 2)) | 1)))}
                _ = v25
            }
            v11 = (v11 / (((^(fi18))) | 1))
            _ = fi18
        }
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                v26 := GenS1{F0: 0}
                var v27 uint32 = (v3 & v3)
                var v28 uint32 = (^(v27))
                _ = v26
                _ = v27
                _ = v28
            }
            if true {
                fmt.Println(v11, v11, v12)
                if false {
                    v29 := GenS2{F0: 127, F1: 0, F2: 51}
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 8; fi30++ {
                        fmt.Println(fi30)
                        v31 := GenS2{F0: 0, F1: 982, F2: 42}
                        v10 += (v10 / ((v10) | 1))
                        _ = fi30
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 20; fi32++ {
                        v15 = 1e10
                        var v33 int16 = 152
                        _ = fi32
                        _ = v33
                    }
                    v34 := GenS3{F0: 128, F1: (v11 << ((13) & 63)), F2: v10}
                    _ = v34
                }
            } else {
                switch ((("\n\t" + "a") + ("foo bar" + ("\n\t" + "\n\t"))) + "a") {
                case "a":
                    v35 := GenS4{F0: -9223372036854775808}
                    _ = v35
                case "":
                    var v36 float64 = 1e10
                    _ = v36
                default:
                    v10 = (v10 / ((v10) | 1))
                }
            }
        }
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v12 = (v12 % ((65535) | 1))
        if (true || (42 > 42)) {
            v11 = fi37
            if true {
                for fi38 := 0; fi38 < 5; fi38++ {
                    fmt.Println(fi38, fi37)
                    switch -914 {
                    case 9223372036854775807:
                        if false {
                            v15 = v15
                            v39 := GenS3{F0: 128, F1: fi37, F2: -128}
                            if false {
                                delete(v4, "foo bar")
                                v11 = 4294967295
                                delete(v4, "")
                            } else {
                                var v40 uint = v12
                                _ = v40
                            }
                            v41 := GenS2{F0: 42, F1: 42, F2: 152}
                            _ = v39
                            _ = v41
                        } else {
                            v11 = (v11 & v11)
                            fmt.Println(v3, v10, v11)
                        }
                    case -2147483648:
                        v10 = v10
                    }
                    _ = fi38
                }
                v3 -= 0
            } else {
                var v42 uint8 = 1
                _ = v42
            }
            var v43 uint16 = 65535
            _ = v43
        } else {
            v44 := GenS2{F0: 15, F1: uint64(v3), F2: 113}
            v15 -= (v15 * v15)
            v12 = v12
            _ = v44
        }
        var v45 int8 = v2
        _ = fi37
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
