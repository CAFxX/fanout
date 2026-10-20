package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0() string {
    var v7 uint8 = 128
    if (!false) {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{fi8, 1, fi8, -1}
            if true {
                for fi10 := 0; fi10 < 14; fi10++ {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if (59 <= 20) {
                        if false {
                            v7 = (v7 % ((v7) | 1))
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                        var v11 uint32 = 42
                        _ = v11
                    } else {
                        if false {
                            var v12 uint64 = 237
                            v7 -= (v7 & (v7 ^ 195))
                            _ = v12
                        }
                        v7 += (v7 % ((v7) | 1))
                    }
                    fmt.Println(v7, fi8, v7)
                    _ = fi10
                }
                _ = len(v9)
            }
            _ = fi8
            _ = v9
        }
        switch "hello" {
        case "a":
            var v13 bool = true
            _ = v13
        }
    }
    v14 := []int{-185, 42, 385}
    var v15 uint16 = 995
    v16 := GenS1{F0: 1, F1: 32767, F2: (("foo bar" + "hello") + ("foo bar" + "\n\t"))}
    _ = v14
    _ = v15
    _ = v16
    _ = v7
    return ("\n\t" + "")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 100
    var v3 uint64 = 4294967295
    v4 := make(map[int]bool)
    var v5 bool = (127 < 79)
    v6 := make(map[uint32]int)
    var v17 uint32 = 877
    if v5 {
        v17 -= v17
        delete(v6, uint32(v3))
        v18 := v6[65535]
        _ = v18
    } else {
        if v5 {
            v3 -= v3
            var v19 float32 = -2.25
            _ = v19
        } else {
            delete(v1, v17)
            var v20 int64 = -226
            v3 += (v3 - 660)
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := []int{}
            _ = fi21
            _ = v22
        }
        if v5 {
            for fi23 := 0; fi23 < 4; fi23++ {
                switch "hello" {
                case "foo bar":
                    fmt.Println(fi23, fi23)
                }
                _ = fi23
            }
            v24 := v4[4294967295]
            _ = v24
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                if (v5 || v5) {
                    var v26 uint8 = 85
                    v26 = 127
                    _ = v26
                } else {
                    v2 = v2
                }
                _ = fi25
            }
            var v27 int64 = 307
            fmt.Println(v2, v5)
            _ = v27
        }
    }
    delete(v1, (v17 | 65535))
    var v28 int64 = 9223372036854775807
    v29 := v1[v17]
    v30 := []int{}
    v17 += 4294967295
    var v31 uint = 65535
    v32 := GenS1{F0: 1, F1: 42, F2: v29}
    v33 := GenS1{F0: 1, F1: -697, F2: v29}
    if len(v30) > 0 {
        v30[0] = 0
    }
    fmt.Println(v17)
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        for fi35 := 0; fi35 < 18; fi35++ {
            v1[v17] = ((v29 + (("" + v29) + "a")) + "\n\t")
            if (255 == 32767) {
                v3 = 2
            } else {
                _ = len(v30)
                if true {
                    v29 = "hello"
                } else {
                    var v36 uint32 = (v17 << (((v31 + v31)) & 63))
                    v30 = append(v30, (fi34 / ((-654) | 1)))
                    fmt.Println(v29)
                    _ = v36
                }
                var v37 int32 = 42
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS2{F0: 969, F1: v2, F2: (v3 >> (((v31 ^ v31)) & 63))}
        _ = fi34
        _ = v38
    }
    v29 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
