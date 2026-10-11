package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int8 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 int16 = 1000
    var v2 string = "a"
    v3 := make(map[string]int)
    v4 := []uint32{4294967295, 851, 278, 927}
    v5 := []uint32{4294967295}
    var v6 bool = true
    v3[("a" + "a")] = 32767
    if true {
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi8
            }
            v9 := GenS1{F0: 760, F1: 2147483647, F2: 42}
            _ = fi7
            _ = v9
        }
        var v10 uint64 = 313
        _ = v10
    } else {
        fmt.Println(v2, v1, v1)
        fmt.Println(v2, v6)
    }
    var v11 uint = 451
    var v12 int = 42
    v13 := GenS1{F0: 0, F1: 646, F2: 112}
    var v14 float32 = 1e10
    v11 += (v11 / ((18446744073709551615) | 1))
    v15 := GenS1{F0: 2, F1: -114, F2: -19}
    if (true && v6) {
        fmt.Println(v11)
        for fi16 := 0; fi16 < 2; fi16++ {
            if v6 {
                v17 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: 42}
                v18 := GenS2{F0: fi16, F1: 42}
                _ = v17
                _ = v18
            } else {
                v11 = v11
                if v6 {
                    v19 := GenS2{F0: v12, F1: 530}
                    v6 = false
                    _ = v19
                } else {
                    if ((v11 ^ v11) == (v11 - v11)) {
                        v20 := GenS2{F0: -2147483648, F1: 1}
                        _ = v20
                    } else {
                        switch 9223372036854775807 {
                        case 229:
                            var v21 int16 = (v1 / ((v1) | 1))
                            _ = v21
                        case -1:
                            var v22 float32 = -2.25
                            _ = v22
                        }
                    }
                    var v23 uint32 = 929
                    if v6 {
                        fmt.Println(v12, fi16, v12)
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if (12 == 127) {
                                v12 -= 32768
                                var v25 uint64 = 1
                                _ = v25
                            }
                            v5 = append(v5, (^(v23)))
                            v14 = (v14 + v14)
                            _ = fi24
                        }
                        delete(v3, v2)
                    }
                    _ = v23
                }
                if ((42 <= 49) || v6) {
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    var v26 int16 = v1
                    _ = v26
                } else {
                    if v6 {
                        if v6 {
                            v2 = (v2 + v2)
                            for fi27 := 0; fi27 < 12; fi27++ {
                                if (v6 || (v6 || v6)) {
                                    v1 = v1
                                    v6 = v6
                                } else {
                                    v6 = (!v6)
                                    _ = len(v5)
                                }
                                _ = fi27
                            }
                        } else {
                            switch v2 {
                            case "\n\t":
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    fmt.Println(fi16, v11, v2)
                                    fmt.Println(v14, v2)
                                    _ = fi28
                                }
                            case "":
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    switch "" {
                                    case "a":
                                        var v30 int = v12
                                        _ = v30
                                    }
                                    _ = fi29
                                }
                            }
                        }
                        fmt.Println(v6, fi16, v2)
                        v31 := GenS1{F0: 168, F1: 0, F2: 42}
                        _ = v31
                    }
                    var v32 bool = true
                    fmt.Println(v11, v6)
                    _ = v32
                }
            }
            var v33 int64 = 9223372036854775807
            _ = fi16
            _ = v33
        }
        v34 := GenS1{F0: 1, F1: -2147483648, F2: 127}
        _ = v34
    }
    var v35 uint = (^(v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
