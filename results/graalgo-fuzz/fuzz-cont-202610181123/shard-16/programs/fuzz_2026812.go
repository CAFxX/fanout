package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint; F2 uint32 }
func genFn0() uint8 {
    v6 := []int{1}
    v7 := GenS1{F0: 181, F1: 0, F2: 1}
    var v8 float64 = (0.0 - ((-2.25 * (3.14159 + (3.14159 - 0.1))) * 0.0))
    switch -1 {
    case 32768:
        switch "foo bar" {
        case "hello":
            for fi9 := 0; fi9 < 20; fi9++ {
                v8 -= 1e10
                var v10 float32 = -2.25
                _ = fi9
                _ = v10
            }
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = 1
            }
        case "a":
            var v11 uint64 = 269
            _ = v11
        default:
            v12 := GenS1{F0: 919, F1: -1, F2: 32768}
            _ = v12
        }
    default:
        _ = len(v6)
    }
    _ = v6
    _ = v7
    _ = v8
    return 119
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 float64 = 1e10
    var v4 int8 = 100
    v5 := []string{"\n\t", "a", ("hello" + ("hello" + "hello"))}
    var v13 uint32 = 2
    v3 = v3
    var v14 int64 = 851
    v15 := GenS1{F0: 501, F1: 42, F2: 4294967295}
    v13 -= 1
    var v16 int64 = v14
    v17 := GenS2{F0: 151, F1: 9223372036854775807, F2: 18446744073709551615}
    if false {
        if true {
            switch "a" {
            case "\n\t":
                v1["\n\t"] = "hello"
            case "foo bar":
                v18 := v1[((("hello" + "\n\t") + ("\n\t" + ("foo bar" + "foo bar"))) + "foo bar")]
                _ = v18
            default:
                var v19 int16 = -521
                _ = v19
            }
            switch 2147483647 {
            case -2147483648:
                var v20 uint = 0
                _ = v20
            case 448:
                if (!false) {
                    if false {
                        if false {
                            var v21 int16 = 32767
                            _ = v21
                        } else {
                            v22 := v1["foo bar"]
                            _ = v22
                        }
                        fmt.Println(v3)
                    }
                    v13 += 42
                    if len(v5) > 0 {
                        v5[0] = ""
                    }
                } else {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 int = fi23
                        v16 -= 42
                        _ = fi23
                        _ = v24
                    }
                    v16 = v14
                }
            case 32767:
                v16 = (v16 >> ((0) & 63))
            default:
                fmt.Println(v14, v13, v16)
            }
            v4 = 105
        }
        v4 += -1
    }
    if true {
        fmt.Println(v14)
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 int64 = -9223372036854775808
            for fi27 := 0; fi27 < 5; fi27++ {
                switch 787 {
                case 42:
                    if (1 == int32(fi27)) {
                        v3 += -2.25
                    } else {
                        if true {
                            fmt.Println(v16, v16)
                            var v28 uint8 = genFn0()
                            _ = v28
                        }
                        switch (("hello" + "foo bar") + "\n\t") {
                        case "foo bar":
                            v29 := GenS3{F0: 0}
                            _ = v29
                        case "a":
                            var v30 uint32 = (v13 << ((776) & 63))
                            _ = v30
                        default:
                            if true {
                                v31 := GenS4{F0: (fi25 - fi27), F1: 1, F2: v13}
                                v5 = append(v5, "")
                                var v32 string = ""
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint64 = 18446744073709551615
                                var v34 bool = (v26 > (v16 + v16))
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                case 128:
                    v35 := GenS2{F0: (fi27 & (fi27 << ((1) & 63))), F1: (fi27 << ((1) & 63)), F2: 42}
                    _ = v35
                case -1:
                    v3 += v3
                    var v36 uint32 = v13
                    var v37 int64 = v26
                    v13 = (v36 - 4294967295)
                    fmt.Println(v4, fi25, v37)
                    _ = v36
                    _ = v37
                }
                var v38 string = "a"
                var v39 bool = true
                _ = fi27
                _ = v38
                _ = v39
            }
            _ = fi25
            _ = v26
        }
        switch -107 {
        case 32768:
            v40 := GenS4{F0: 1, F1: 0, F2: (v13 >> ((255) & 63))}
            _ = v40
        case 645:
            var v41 uint64 = 0
            _ = v41
        case 0:
            var v42 int32 = 2147483647
            _ = v42
        }
    }
    var v43 uint8 = 42
    v43 -= v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
