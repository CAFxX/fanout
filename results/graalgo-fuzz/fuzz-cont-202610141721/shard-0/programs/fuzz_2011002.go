package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0() int32 {
    var v5 string = "\n\t"
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{(-(fi6)), (fi6 | (fi7 / ((fi6) | 1))), fi6, (fi7 | fi6)}
            v8 = append(v8, 42)
            v5 = ("hello" + "a")
            _ = fi7
            _ = v8
        }
        switch (-(fi6)) {
        case 42:
            var v9 float32 = 3.14159
            _ = v9
        case 128:
            if false {
                if (100 < -5) {
                    v5 = ((v5 + v5) + "\n\t")
                    var v10 bool = (fi6 == fi6)
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 float32 = float32((1e10 - 1.5))
                        if false {
                            fmt.Println(v10, v12)
                            var v13 uint32 = 1
                            fmt.Println(v13, fi11, v12)
                            _ = v13
                        } else {
                            v14 := GenS1{F0: -1, F1: 42}
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
            } else {
                switch v5 {
                case "\n\t":
                    v5 = "foo bar"
                case "foo bar":
                    if true {
                        for fi15 := 0; fi15 < 18; fi15++ {
                            var v16 uint = 255
                            if true {
                                v5 = v5
                                v17 := []int{}
                                _ = v17
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v5 = "foo bar"
                    }
                case "hello":
                    var v18 int16 = 0
                    _ = v18
                }
            }
        case 32768:
            var v19 uint = 65535
            _ = v19
        }
        v5 = (v5 + v5)
        _ = fi6
    }
    v20 := GenS1{F0: -9223372036854775808, F1: 324}
    _ = v20
    _ = v5
    return -1
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{"a"}
    var v3 uint8 = 127
    var v4 uint32 = 0
    if ((32767 == 4294967295) && true) {
        v21 := GenS1{F0: 9223372036854775807, F1: (v1 << ((787) & 63))}
        v2 = append(v2, "")
        _ = v21
    } else {
        v1 += v1
        for fi22 := 0; fi22 < 19; fi22++ {
            v1 = v1
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 uint64 = 0
                var v25 int64 = -389
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: 0, F1: 255}
            _ = fi22
            _ = v26
        }
        fmt.Println(v3)
    }
    var v27 int = 9223372036854775807
    v28 := GenS1{F0: int64(v3), F1: (^(v1))}
    for fi29 := 0; fi29 < 12; fi29++ {
        if len(v2) > 0 {
            v2[0] = ""
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi29
    }
    fmt.Println(v3, v1, v27)
    var v30 uint16 = (v1 ^ uint16(v4))
    fmt.Println(v30)
    var v31 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
