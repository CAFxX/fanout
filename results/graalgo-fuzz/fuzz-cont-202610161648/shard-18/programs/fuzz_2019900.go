package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 1000
    v3 := make(map[int]string)
    v4 := []uint32{}
    var v5 int16 = (v2 ^ 1000)
    var v6 uint32 = 2
    var v7 uint64 = 2
    v5 += v5
    if true {
        if (!false) {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v6, v7, v5)
                v5 -= v5
            } else {
                v7 += (v7 * 2)
                for fi8 := 0; fi8 < 14; fi8++ {
                    fmt.Println(fi8, v6)
                    _ = fi8
                }
                var v9 uint8 = 255
                _ = v9
            }
            if true {
                v10 := GenS1{F0: true}
                for fi11 := 0; fi11 < 1; fi11++ {
                    v7 += v7
                    v6 = (v6 << ((536) & 63))
                    if false {
                        var v12 bool = (1.5 == 3.14159)
                        switch 634 {
                        case 0:
                            v2 = 32767
                        case -2147483648:
                            for fi13 := 0; fi13 < 15; fi13++ {
                                fmt.Println(v7)
                                fmt.Println(v7, fi13, fi11)
                                _ = fi13
                            }
                        }
                        _ = len(v3)
                        _ = v12
                    }
                    _ = fi11
                }
                _ = v10
            }
        }
    } else {
        v14 := GenS1{F0: false}
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                v2 -= (v2 << ((200) & 63))
                var v17 int = fi16
                if true {
                    _ = len(v4)
                } else {
                    fmt.Println(v5, v5, v7)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi16
                _ = v17
            }
            v5 = v5
            _ = fi15
        }
        v2 += (v5 & (v2 | v2))
        _ = v14
    }
    delete(v3, 654)
    var v18 bool = false
    var v19 uint = 618
    v5 = (v2 - 42)
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v2)
        var v21 uint16 = 65535
        var v22 int8 = 21
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint16 = 32767
    fmt.Println(v6, v5, v18)
    _ = len(v3)
    v6 = (^(v6))
    var v24 bool = (v18 || ((0 == 9223372036854775807) || v18))
    if true {
        var v25 bool = false
        if (188 <= -2147483648) {
            if (760 <= (v6 - (v6 ^ (v6 - v6)))) {
                v1["foo bar"] = v6
            } else {
                var v26 int16 = (v5 + (v2 * (v2 - v5)))
                _ = v26
            }
            var v27 uint = 4294967295
            v23 += (v23 >> (((v19 / (((v27 & (^(v27)))) | 1))) & 63))
            _ = v27
        } else {
            delete(v1, ("\n\t" + "hello"))
            if (-41 >= 0) {
                if (3.14159 <= (((((0.0 - 1e10) / 0.1) / float32(3.14159)) + 3.14159) - float32(1e10))) {
                    v18 = ((v24 && v25) || v24)
                } else {
                    if v25 {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            for fi29 := 0; fi29 < 15; fi29++ {
                                fmt.Println(v19, v6, v5)
                                _ = fi29
                            }
                            fmt.Println(v5)
                            for fi30 := 0; fi30 < 15; fi30++ {
                                v23 = v23
                                v31 := GenS2{F0: 1, F1: v7}
                                if len(v4) > 0 {
                                    v4[0] = (v6 - uint32(v7))
                                }
                                _ = fi30
                                _ = v31
                            }
                            _ = fi28
                        }
                        var v32 uint = (v19 | v19)
                        v33 := GenS3{F0: 9223372036854775807, F1: 42, F2: v24}
                        _ = v32
                        _ = v33
                    }
                }
                var v34 uint32 = 16
                v23 = (v23 - v23)
                _ = v34
            } else {
                var v35 uint64 = v7
                _ = v35
            }
        }
        _ = v25
    } else {
        v24 = false
    }
    v24 = (false || v24)
    if ((^(v19)) == (v19 | v19)) {
        v5 -= v5
        v7 += 868
    } else {
        if (!v18) {
            v6 = 0
            var v36 uint32 = 2
            _ = v36
        }
        var v37 int32 = -756
        _ = v37
    }
    v6 = (v6 >> ((v19) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
