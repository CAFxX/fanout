package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int16; F1 int; F2 string }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{533}
    var v3 string = "hello"
    v4 := []string{v3, v3, "a", v3}
    var v5 int32 = -2147483648
    switch v3 {
    case "\n\t":
        if true {
            v3 = v3
            for fi6 := 0; fi6 < 3; fi6++ {
                fmt.Println(v3)
                _ = fi6
            }
        }
    case "foo bar":
        var v7 uint = 42
        _ = v7
    }
    v3 = v3
    var v8 uint32 = 2147483648
    var v9 int = 649
    v3 = ("" + (v3 + (v3 + ((v3 + "") + (v3 + v3)))))
    v10 := GenS1{F0: -1}
    if true {
        fmt.Println(v8, v5, v9)
        for fi11 := 0; fi11 < 13; fi11++ {
            if (!false) {
                for fi12 := 0; fi12 < 13; fi12++ {
                    if (!((!(-439 >= -1)) && ((v5 ^ v5) <= (^(v5))))) {
                        v13 := GenS1{F0: 124}
                        v14 := GenS2{F0: 127}
                        _ = v13
                        _ = v14
                    }
                    _ = fi12
                }
                switch (v3 + (v3 + ("\n\t" + "hello"))) {
                case "a":
                    for fi15 := 0; fi15 < 10; fi15++ {
                        if false {
                            switch 1 {
                            case 961:
                                if (!true) {
                                    v16 := GenS2{F0: 42}
                                    _ = v16
                                }
                            default:
                                fmt.Println(v5, v5)
                                if false {
                                    switch ("hello" + v3) {
                                    case "hello":
                                        var v17 uint16 = 65535
                                        _ = v17
                                    case "a":
                                        v18 := GenS3{F0: fi15, F1: 1}
                                        _ = v18
                                    default:
                                        var v19 int64 = -645
                                        _ = v19
                                    }
                                    var v20 float32 = 0.1
                                    if (1 < 2) {
                                        var v21 uint8 = 0
                                        _ = v21
                                    }
                                    _ = v20
                                } else {
                                    v22 := GenS3{F0: (fi15 >> ((65535) & 63)), F1: 127}
                                    if true {
                                        v5 = (v5 + 42)
                                    } else {
                                        var v23 float32 = 0.1
                                        _ = v23
                                    }
                                    if false {
                                        fmt.Println(fi11, v8)
                                        v24 := v1[v3]
                                        var v25 int16 = 170
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if true {
                                            v5 += (v5 / ((v5) | 1))
                                        }
                                    }
                                    _ = v22
                                }
                                fmt.Println(v5)
                                var v26 float32 = 1.5
                                _ = v26
                            }
                            v3 = v3
                            switch v3 {
                            case "":
                                var v27 int16 = 1000
                                _ = v27
                            case "a":
                                v5 = (v5 ^ -1)
                            case "hello":
                                if (!(v5 < (v5 ^ v5))) {
                                    v3 = (v3 + "hello")
                                }
                            default:
                                if (v8 != 2147483648) {
                                    if ((v9 >> ((318) & 63)) == 32768) {
                                        v8 = v8
                                        v3 = "\n\t"
                                    } else {
                                        fmt.Println(v8)
                                    }
                                    var v28 uint8 = 2
                                    v28 -= (v28 * (v28 / ((20) | 1)))
                                    _ = v28
                                }
                            }
                            v3 = ("\n\t" + v3)
                            var v29 int16 = -370
                            var v30 float64 = 1e10
                            _ = v29
                            _ = v30
                        }
                        var v31 string = "\n\t"
                        v8 = (v8 * (v8 - 4294967295))
                        _ = fi15
                        _ = v31
                    }
                case "":
                    var v32 int = (v9 << ((258) & 63))
                    _ = v32
                case "foo bar":
                    fmt.Println(v9)
                default:
                    for fi33 := 0; fi33 < 12; fi33++ {
                        v2 = append(v2, (v8 * v8))
                        if ((32767 >= 712) && (0 <= 32767)) {
                            v2 = append(v2, (^(v8)))
                            _ = len(v1)
                        } else {
                            fmt.Println(fi33, v8, fi33)
                        }
                        v34 := v1["a"]
                        v35 := GenS4{F0: 1, F1: (v9 + v9), F2: v3}
                        v36 := GenS4{F0: 1, F1: fi33, F2: v3}
                        var v37 int16 = 0
                        v38 := GenS2{F0: 1}
                        var v39 uint8 = 1
                        var v40 float64 = float64(0.1)
                        v39 = v39
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                }
            } else {
                var v41 bool = (!(42 != 0))
                var v42 float64 = float64(-2.25)
                _ = v41
                _ = v42
            }
            switch -2147483648 {
            case 42:
                v43 := v1[v3]
                _ = v43
            case -9223372036854775808:
                if len(v2) > 0 {
                    v2[0] = v8
                }
            default:
                v44 := GenS5{F0: false}
                _ = v44
            }
            v8 = 2
            _ = fi11
        }
        if len(v4) > 0 {
            v4[0] = (v3 + v3)
        }
    }
    var v45 uint16 = 867
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v8
    _ = v9
}
