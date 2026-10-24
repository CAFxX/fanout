package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
func genFn0() int64 {
    v5 := []int{}
    if false {
        var v6 string = "hello"
        v5 = append(v5, -361)
        if false {
            for fi7 := 0; fi7 < 15; fi7++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v8 string = "\n\t"
                _ = fi7
                _ = v8
            }
        } else {
            _ = len(v5)
            v9 := GenS1{F0: 0}
            _ = v9
        }
        _ = v6
    }
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 0
    var v2 int = 32768
    v3 := []int{2147483647, v2, v2}
    var v4 float32 = float32(-2.25)
    if len(v3) > 0 {
        v3[0] = 0
    }
    v4 = 0.1
    v3 = append(v3, (v2 << ((175) & 63)))
    v2 += v2
    var v10 uint32 = 323
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = v2
        fmt.Println(v10, v4)
        _ = fi11
        _ = v12
    }
    if false {
        v1 -= 42
        fmt.Println(v2, v1)
        var v13 int32 = 2147483647
        _ = v13
    } else {
        switch "foo bar" {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = (v2 ^ (v2 / ((v2) | 1)))
            }
        default:
            if false {
                v14 := GenS1{F0: 2}
                var v15 string = "\n\t"
                switch v2 {
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = (-(v2))
                    }
                case 4294967295:
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch "a" {
                        case "\n\t":
                            if (-128 == 100) {
                                switch 951 {
                                case 995:
                                    var v17 int64 = -558
                                    _ = v17
                                }
                                var v18 uint64 = 475
                                switch 2147483647 {
                                case 32767:
                                    fmt.Println(v4)
                                case 0:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = v18
                            } else {
                                fmt.Println(v4)
                                fmt.Println(v10)
                            }
                        }
                        var v19 int32 = 2147483647
                        v20 := GenS1{F0: 1}
                        _ = fi16
                        _ = v19
                        _ = v20
                    }
                case -1:
                    v4 += v4
                }
                _ = v14
                _ = v15
            }
        }
        v21 := GenS2{F0: 2}
        _ = v21
    }
    v2 -= (v2 / ((v2) | 1))
    var v22 int = v2
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v1 = 29
            v1 -= v1
            switch (("" + "\n\t") + ("" + "foo bar")) {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                var v25 float32 = float32(1.5)
                _ = v25
            }
            _ = fi24
        }
        if ((true && (((42 < 1) && (float64(v4) == 1e10)) || (((346 == 1) || true) || (v4 == 0.1)))) && false) {
            _ = len(v3)
        } else {
            var v26 int64 = -256
            _ = v26
        }
        v3 = append(v3, 127)
        _ = fi23
    }
    var v27 uint32 = 65535
    switch (v2 & (v2 * 128)) {
    case 42:
        _ = len(v3)
    case 664:
        var v28 int16 = 1
        _ = v28
    case 32767:
        for fi29 := 0; fi29 < 4; fi29++ {
            for fi30 := 0; fi30 < 9; fi30++ {
                v10 -= 378
                v2 = 32767
                _ = fi30
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 float32 = v4
                fmt.Println(v27, v1)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi31
                _ = v32
            }
            v33 := GenS3{F0: "\n\t"}
            _ = fi29
            _ = v33
        }
    default:
        var v34 int32 = -1
        _ = v34
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v22 += 164
    var v35 uint32 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v27
    _ = v3
    _ = v35
    _ = v4
}
