package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(fi7)
        v8 := GenS1{F0: 1}
        _ = fi7
        _ = v8
    }
    var v9 uint = 18446744073709551615
    var v10 uint64 = 0
    _ = v10
    _ = v9
    return -32768
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 122
    v3 := make(map[string]uint32)
    var v4 uint8 = 1
    var v5 int = 32768
    v6 := make(map[uint32]uint32)
    if true {
        v5 += v5
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 string = "a"
            switch v5 {
            case -1:
                if false {
                    delete(v6, 4294967295)
                    v13 := GenS1{F0: 32767}
                    var v14 int32 = -634
                    _ = v13
                    _ = v14
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("a" + (v12 + v12))
                    }
                }
            case 32767:
                v15 := GenS1{F0: 7}
                _ = v15
            default:
                v16 := GenS1{F0: 382}
                _ = v16
            }
            var v17 string = "foo bar"
            _ = fi11
            _ = v12
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        delete(v6, 1)
        for fi19 := 0; fi19 < 7; fi19++ {
            if true {
                if true {
                    var v20 string = "foo bar"
                    var v21 bool = (false || (v2 < (^(v2))))
                    _ = v20
                    _ = v21
                }
                var v22 uint16 = 65535
                _ = v22
            }
            _ = fi19
        }
        fmt.Println(v2, v5, v5)
        _ = fi18
    }
    fmt.Println(v4, v5, v2)
    switch ("a" + "a") {
    case "a":
        switch (v5 % (((v5 << (((^(v2))) & 63))) | 1)) {
        case 4294967295:
            var v23 int16 = 966
            _ = v23
        }
    default:
        v24 := GenS1{F0: 0}
        _ = v24
    }
    v5 -= (v5 << (((v2 % (((^(v2))) | 1))) & 63))
    switch "hello" {
    case "foo bar":
        if false {
            v5 = 326
            v5 = (v5 + (v5 - 1))
            var v25 int64 = 42
            _ = v25
        } else {
            switch (v5 & 169) {
            case 362:
                fmt.Println(v4, v4, v5)
            case 32767:
                v4 -= v4
            case 127:
                fmt.Println(v4, v5)
            default:
                v26 := GenS1{F0: 0}
                _ = v26
            }
            for fi27 := 0; fi27 < 16; fi27++ {
                if true {
                    v2 = (^(v2))
                } else {
                    var v28 int16 = 438
                    switch v5 {
                    case -2147483648:
                        v28 = v28
                    default:
                        for fi29 := 0; fi29 < 11; fi29++ {
                            var v30 float64 = -2.25
                            if true {
                                delete(v6, 649)
                                var v31 int = (v5 & fi27)
                                v30 -= (1e10 * (v30 * 1e10))
                                if false {
                                    fmt.Println(v2)
                                    v32 := v3[("hello" + "hello")]
                                    fmt.Println(fi27, fi29)
                                    _ = v32
                                }
                                _ = v31
                            }
                            _ = fi29
                            _ = v30
                        }
                    }
                    _ = v28
                }
                _ = fi27
            }
        }
    case "":
        v33 := GenS1{F0: 0}
        _ = v33
    default:
        v34 := GenS1{F0: 65535}
        _ = v34
    }
    if true {
        fmt.Println(v2)
    }
    delete(v3, ("" + "a"))
    delete(v3, ("" + ""))
    var v35 int32 = 672
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
