package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    v1 := []string{}
    v2 := []string{}
    v3 := []uint32{}
    var v4 uint32 = 614
    fmt.Println(v4)
    v5 := GenS1{F0: 871}
    fmt.Println(v4)
    var v6 string = "foo bar"
    v6 = v6
    v3 = append(v3, (v4 << ((42) & 63)))
    if false {
        fmt.Println(v6, v6)
        if false {
            switch 1 {
            case -2147483648:
                for fi7 := 0; fi7 < 9; fi7++ {
                    switch (v6 + v6) {
                    case "\n\t":
                        var v8 float64 = -2.25
                        _ = v8
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                    case "a":
                        if ((!(1e10 >= 0.0)) || (26 > 127)) {
                            if (!true) {
                                var v9 uint = 44
                                _ = v9
                            } else {
                                v10 := GenS1{F0: 0}
                                v11 := GenS1{F0: -9223372036854775808}
                                _ = v10
                                _ = v11
                            }
                            fmt.Println(v4)
                        } else {
                            _ = len(v3)
                            v3 = append(v3, v4)
                            v12 := GenS2{F0: fi7, F1: (v6 + v6), F2: 128}
                            _ = v12
                        }
                    }
                    if (false && false) {
                        fmt.Println(fi7, v6)
                        v6 = ("\n\t" + (v6 + "hello"))
                        if len(v2) > 0 {
                            v2[0] = ("foo bar" + v6)
                        }
                    }
                    _ = fi7
                }
            case 127:
                var v13 bool = false
                _ = v13
            case -974:
                fmt.Println(v4, v6)
            default:
                v14 := GenS2{F0: -674, F1: "a", F2: 178}
                _ = v14
            }
        } else {
            if false {
                if len(v1) > 0 {
                    v1[0] = (v6 + (v6 + (v6 + v6)))
                }
            } else {
                v6 = "foo bar"
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        for fi15 := 0; fi15 < 5; fi15++ {
            switch 294 {
            case 366:
                if true {
                    v6 = ("hello" + ("foo bar" + v6))
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                }
            case -71:
                if len(v2) > 0 {
                    v2[0] = (v6 + "foo bar")
                }
            case 42:
                if (1000 > -88) {
                    switch (fi15 ^ (fi15 >> ((uint(v4)) & 63))) {
                    case -9223372036854775808:
                        v16 := GenS3{F0: -540, F1: -1}
                        _ = v16
                    case 32768:
                        fmt.Println(v4, v4)
                    case -646:
                        if true {
                            fmt.Println(v6, v6, v6)
                        } else {
                            var v17 uint64 = 18446744073709551615
                            var v18 uint64 = v17
                            for fi19 := 0; fi19 < 11; fi19++ {
                                var v20 uint8 = 127
                                _ = fi19
                                _ = v20
                            }
                            _ = v17
                            _ = v18
                        }
                    default:
                        v21 := GenS3{F0: -154, F1: -2147483648}
                        _ = v21
                    }
                    fmt.Println(fi15, v6)
                    v4 = 882
                }
            }
            v6 = (((v6 + v6) + "hello") + v6)
            _ = fi15
        }
    }
    v6 = v6
    v22 := GenS2{F0: -1, F1: "foo bar", F2: 25}
    if true {
        if (32767 >= 255) {
            v3 = append(v3, v4)
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if ((-2.25 == 3.14159) && false) {
                    switch (v6 + v6) {
                    case "a":
                        var v23 uint64 = 1
                        _ = v23
                    case "":
                        v4 += 2147483647
                    default:
                        if false {
                            var v24 uint64 = 2
                            if (!true) {
                                var v25 int64 = -389
                                fmt.Println(v25, v4)
                                v26 := GenS3{F0: (v25 ^ 42), F1: 0}
                                _ = v25
                                _ = v26
                            } else {
                                v1 = append(v1, "a")
                            }
                            for fi27 := 0; fi27 < 19; fi27++ {
                                v3 = append(v3, v4)
                                _ = fi27
                            }
                            _ = v24
                        } else {
                            v6 = "\n\t"
                        }
                    }
                }
            }
        }
    }
    var v28 int = 127
    var v29 float32 = 0.1
    v4 -= 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
