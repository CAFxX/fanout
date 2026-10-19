package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 uint, p2 int64) int {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: -43, F1: p2}
        if (true || (!(p2 < (p2 % (((p2 | p2)) | 1))))) {
            if (!(false && false)) {
                var v7 uint16 = 32767
                p0 += p0
                if (p1 > 42) {
                    p2 -= p2
                    var v8 uint32 = 2147483648
                    for fi9 := 0; fi9 < 11; fi9++ {
                        var v10 int16 = 1000
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                }
                _ = v7
            } else {
                var v11 uint16 = 65535
                _ = v11
            }
            v12 := GenS1{F0: p0, F1: p2}
            _ = v12
        } else {
            v13 := []int{32768}
            fmt.Println(fi5, fi5)
            var v14 uint64 = 4294967295
            _ = v13
            _ = v14
        }
        _ = fi5
        _ = v6
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        switch "a" {
        case "foo bar":
            switch "hello" {
            case "":
                p2 -= p2
            case "hello":
                for fi16 := 0; fi16 < 2; fi16++ {
                    if true {
                        fmt.Println(fi16, fi15, p2)
                        p1 -= (p1 << ((p1) & 63))
                    } else {
                        var v17 int16 = 42
                        _ = v17
                    }
                    _ = fi16
                }
            case "foo bar":
                v18 := []int{(fi15 ^ 602)}
                _ = v18
            default:
                switch 4294967295 {
                case 32768:
                    var v19 int = fi15
                    _ = v19
                case 913:
                    if false {
                        if (p2 <= (^(p2))) {
                            p1 = p1
                            v20 := []int{1, 0, fi15}
                            for fi21 := 0; fi21 < 17; fi21++ {
                                if false {
                                    switch ("\n\t" + "") {
                                    case "a":
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    case "\n\t":
                                        var v22 uint64 = 485
                                        _ = v22
                                    case "":
                                        var v23 int64 = p2
                                        p0 = (p0 & (p0 % ((p0) | 1)))
                                        _ = v23
                                    }
                                }
                                _ = fi21
                            }
                            _ = v20
                        }
                    } else {
                        p0 += (p0 >> ((p1) & 63))
                    }
                }
            }
        default:
            v24 := []int{fi15, (-(fi15))}
            _ = v24
        }
        switch fi15 {
        case -2147483648:
            var v25 uint64 = 679
            _ = v25
        }
        _ = fi15
    }
    var v26 string = "a"
    var v27 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return 42
}
func genFn1(p0 int16) int16 {
    var v28 uint64 = 18446744073709551615
    var v29 float64 = 3.14159
    switch ("" + ("foo bar" + ("foo bar" + "\n\t"))) {
    case "hello":
        v30 := GenS1{F0: 100, F1: -9223372036854775808}
        _ = v30
    case "foo bar":
        fmt.Println(p0)
    case "":
        if true {
            v31 := GenS1{F0: 100, F1: 0}
            _ = v31
        } else {
            v32 := GenS1{F0: 1, F1: -384}
            v33 := []int{814, -2147483648, 530, 42}
            v29 = v29
            _ = v32
            _ = v33
        }
    }
    v34 := []int{9223372036854775807, genFn0(-46, 263, 0), 32767, 128}
    _ = p0
    _ = v28
    _ = v29
    _ = v34
    return -307
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    v3 := []int{32768, 0}
    v4 := make(map[int]string)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    fmt.Println(v1)
    var v35 uint8 = v1
    v36 := GenS2{F0: genFn0(127, 0, -1)}
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS2{F0: (fi37 * (fi37 % (((-(fi37))) | 1)))}
        _ = fi37
        _ = v38
    }
    if false {
        switch ("\n\t" + "hello") {
        case "a":
            v35 -= v35
        }
    } else {
        _ = len(v3)
    }
    fmt.Println(v1)
    var v39 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
