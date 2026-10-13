package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint64, p1 uint32, p2 string) int64 {
    v7 := GenS1{F0: 1, F1: 18446744073709551615}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int32 = -786
        var v10 uint32 = p1
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = (p0 | (^(p0)))
    if false {
        if (true || true) {
            for fi11 := 0; fi11 < 13; fi11++ {
                if true {
                    p0 += p0
                } else {
                    var v12 uint8 = 255
                    p0 -= 645
                    _ = v12
                }
                _ = fi11
            }
        }
        var v13 int64 = 1
        _ = v13
    } else {
        if (181 <= 755) {
            p0 = uint64(p1)
        } else {
            var v14 int64 = 42
            switch -903 {
            case -1:
                if true {
                    v15 := []int{32768, -2147483648, -9223372036854775808}
                    _ = v15
                }
            case 9223372036854775807:
                switch p2 {
                case "hello":
                    v16 := []int{}
                    _ = v16
                }
            case -2147483648:
                var v17 int = -1
                _ = v17
            }
            _ = v14
        }
        if (-382 > 0) {
            v18 := GenS1{F0: 128, F1: (p0 << ((255) & 63))}
            p2 = p2
            _ = v18
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 string = "a"
    v6 := make(map[uint32]string)
    v19 := GenS1{F0: 2, F1: 1}
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 bool = (42 <= 128)
        _ = fi20
        _ = v21
    }
    _ = len(v3)
    switch ("" + v5) {
    case "a":
        var v22 int16 = 0
        _ = v22
    default:
        fmt.Println(v5)
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if (!false) {
            fmt.Println(fi23, v5)
        }
        if (1 == 66) {
            var v24 int16 = 0
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 uint32 = 23
                _ = fi25
                _ = v26
            }
            _ = v24
        } else {
            v5 = ""
        }
        _ = fi23
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := v2[""]
        delete(v2, (v5 + "foo bar"))
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v2[v5] = 42
        _ = fi29
    }
    if false {
        var v30 uint = 42
        switch 128 {
        case 42:
            if (!false) {
                switch v5 {
                case "foo bar":
                    _ = len(v2)
                default:
                    if (true || false) {
                        if true {
                            if true {
                                var v31 int = -1
                                var v32 uint8 = 42
                                _ = v31
                                _ = v32
                            }
                        } else {
                            var v33 string = v5
                            v5 = (("a" + v5) + ((v33 + v5) + (v5 + v33)))
                            for fi34 := 0; fi34 < 3; fi34++ {
                                var v35 float32 = 1.5
                                v35 = v35
                                v35 += v35
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                        v6[uint32(v30)] = ""
                    } else {
                        fmt.Println(v30, v5)
                    }
                }
                for fi36 := 0; fi36 < 11; fi36++ {
                    v30 = (v30 >> (((^(v30))) & 63))
                    v37 := []int{fi36, fi36, fi36}
                    _ = fi36
                    _ = v37
                }
            } else {
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 int16 = 0
                    _ = fi38
                    _ = v39
                }
                var v40 uint64 = 511
                _ = v40
            }
        case 0:
            fmt.Println(v30)
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            switch 32767 {
            case 127:
                v5 = ((v5 + "foo bar") + v5)
            case 123:
                v5 = v5
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                switch (v5 + v5) {
                case "\n\t":
                    fmt.Println(fi42)
                case "a":
                    var v43 int64 = 0
                    _ = v43
                case "":
                    v44 := []int{(fi42 * fi42), fi41, (fi42 >> (((v30 + (v30 | (v30 + (v30 % ((v30) | 1)))))) & 63))}
                    _ = v44
                }
                var v45 uint16 = 2
                _ = fi42
                _ = v45
            }
            _ = fi41
        }
        _ = v30
    }
    v46 := GenS1{F0: 233, F1: 2}
    v47 := GenS1{F0: 255, F1: 1}
    v48 := []int{-137, 878, -253}
    var v49 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
