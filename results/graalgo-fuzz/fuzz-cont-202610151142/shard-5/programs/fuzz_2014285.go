package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int }
func genFn0() bool {
    v4 := []int{32768, -999, -2}
    var v5 float64 = (1.5 - -2.25)
    for fi6 := 0; fi6 < 12; fi6++ {
        switch "" {
        case "":
            switch "\n\t" {
            case "\n\t":
                for fi7 := 0; fi7 < 11; fi7++ {
                    v8 := GenS1{F0: fi7}
                    var v9 uint = 4294967295
                    v9 -= 65535
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        default:
            switch 9223372036854775807 {
            case 792:
                for fi10 := 0; fi10 < 4; fi10++ {
                    if false {
                        v5 -= (1e10 + v5)
                        var v11 uint8 = 127
                        if true {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v12 := GenS1{F0: (fi6 - fi10)}
                            if ((((v5 + (v5 - v5)) + ((1e10 + (-2.25 / 3.14159)) * 0.1)) != 1e10) && false) {
                                var v13 uint = 652
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v13
                            } else {
                                _ = len(v4)
                                var v14 int = 127
                                _ = v14
                            }
                            _ = v12
                        } else {
                            if false {
                                v15 := GenS2{F0: 914, F1: 1}
                                _ = v15
                            }
                            v16 := GenS2{F0: 0, F1: 42}
                            v5 = v5
                            _ = v16
                        }
                        _ = v11
                    } else {
                        var v17 int64 = 42
                        _ = v17
                    }
                    v18 := GenS2{F0: 42, F1: 1}
                    _ = fi10
                    _ = v18
                }
            }
        }
        _ = fi6
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint64 = 587
        _ = fi19
        _ = v20
    }
    var v21 int32 = 2147483647
    _ = v21
    _ = v4
    _ = v5
    return (true || true)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 2
    v3 := make(map[string]uint32)
    v22 := GenS2{F0: 676, F1: 891}
    var v23 uint = 419
    for fi24 := 0; fi24 < 19; fi24++ {
        if (128 == 2) {
            var v25 uint64 = 18446744073709551615
            _ = len(v1)
            v26 := []int{42}
            _ = v25
            _ = v26
        }
        v23 = v23
        v23 = v23
        _ = fi24
    }
    v27 := GenS1{F0: 4294967295}
    v28 := GenS2{F0: -2147483648, F1: -1}
    v23 -= v23
    if false {
        var v29 int64 = -60
        v29 += v29
        _ = v29
    } else {
        var v30 int64 = 9223372036854775807
        fmt.Println(v2, v30)
        _ = v30
    }
    var v31 int32 = 42
    if (!(1 <= 1)) {
        var v32 bool = (0 == 0)
        _ = v32
    }
    v1[v2] = v2
    v31 = v31
    v33 := GenS3{F0: 4294967295}
    delete(v1, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v33
}
