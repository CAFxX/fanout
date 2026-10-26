package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{fi7, fi7}
        v9 := GenS1{F0: 1, F1: -1}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := []int{(^(fi11)), fi11, (^(fi11)), fi11}
        var v13 uint16 = 0
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 65535, F1: true}
    var v15 bool = false
    _ = v10
    _ = v14
    _ = v15
    return 42
}
func genFn1(p0 string, p1 int8, p2 int16) uint8 {
    if (true || false) {
        var v16 float32 = (-2.25 + (1e10 + 1e10))
        var v17 uint16 = 0
        p1 -= -118
        _ = v16
        _ = v17
    } else {
        fmt.Println(p1, p0)
        var v18 uint16 = 0
        _ = v18
    }
    if true {
        p0 = ("\n\t" + p0)
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: p1, F1: 42, F2: 1}
            for fi21 := 0; fi21 < 14; fi21++ {
                fmt.Println(p1, p1, p0)
                if ((p1 + (-(p1))) < p1) {
                    if true {
                        v22 := []int{}
                        _ = v22
                    } else {
                        if false {
                            if false {
                                p1 = p1
                                v23 := []int{(fi21 + -2147483648)}
                                fmt.Println(p0, fi19)
                                _ = v23
                            } else {
                                v24 := GenS2{F0: 0, F1: false}
                                _ = v24
                            }
                            var v25 string = "foo bar"
                            p0 = ("foo bar" + v25)
                            _ = v25
                        }
                        var v26 int8 = p1
                        _ = v26
                    }
                    var v27 int64 = -365
                    p2 += (p2 >> ((26) & 63))
                    _ = v27
                } else {
                    p1 = (p1 ^ (p1 + p1))
                    p2 = (p2 + (p2 * p2))
                    p0 = p0
                    var v28 uint = 781
                    if ((2147483648 >= 2147483647) && true) {
                        if false {
                            v29 := GenS3{F0: p1, F1: 627, F2: 4294967295}
                            v30 := []int{fi21, fi21, 42}
                            _ = v29
                            _ = v30
                        } else {
                            v31 := GenS4{F0: fi19}
                            switch (fi19 % ((-2147483648) | 1)) {
                            case -260:
                                if (!false) {
                                    v28 = v28
                                    v28 += v28
                                    var v32 int16 = p2
                                    _ = v32
                                }
                            case 127:
                                v33 := GenS5{F0: (fi19 - -254), F1: 1}
                                _ = v33
                            }
                            fmt.Println(p2, p0, fi19)
                            _ = v31
                        }
                        if true {
                            var v34 uint32 = 2147483647
                            fmt.Println(v28, fi19)
                            _ = v34
                        } else {
                            var v35 uint8 = 90
                            var v36 string = "hello"
                            _ = v35
                            _ = v36
                        }
                        v37 := []int{9223372036854775807, -926, 32768, (^(fi21))}
                        _ = v37
                    }
                    _ = v28
                }
                _ = fi21
            }
            if true {
                switch "" {
                case "a":
                    var v38 uint8 = 127
                    _ = v38
                }
                switch fi19 {
                case 32767:
                    p2 = -1
                case 139:
                    v39 := GenS3{F0: p1, F1: 18446744073709551615, F2: 42}
                    _ = v39
                }
            }
            _ = fi19
            _ = v20
        }
    } else {
        fmt.Println(p0, p0, p2)
        for fi40 := 0; fi40 < 15; fi40++ {
            p1 = p1
            var v41 uint8 = 3
            for fi42 := 0; fi42 < 2; fi42++ {
                p2 -= (p2 << ((42) & 63))
                var v43 uint8 = (v41 - v41)
                p0 = ""
                _ = fi42
                _ = v43
            }
            _ = fi40
            _ = v41
        }
    }
    p1 += (p1 ^ (p1 - p1))
    v44 := GenS5{F0: -560, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    return 2
}

func main() {
    var v1 uint32 = 512
    var v2 string = "a"
    var v3 bool = false
    var v4 uint16 = 255
    v5 := []uint32{(v1 << ((128) & 63)), v1}
    var v6 float32 = 0.0
    v5 = append(v5, (v1 + v1))
    switch (v2 + v2) {
    case "":
        var v45 int32 = -433
        _ = v45
    case "\n\t":
        v6 = (0.1 + (v6 + v6))
    }
    v46 := GenS1{F0: 1, F1: 42}
    fmt.Println(v2)
    switch 128 {
    case 9223372036854775807:
        for fi47 := 0; fi47 < 7; fi47++ {
            v48 := GenS6{F0: 42}
            _ = fi47
            _ = v48
        }
    case 4294967295:
        for fi49 := 0; fi49 < 20; fi49++ {
            var v50 int64 = -235
            v4 = uint16(v1)
            _ = fi49
            _ = v50
        }
    default:
        v1 -= 1
    }
    v1 = (v1 | (v1 ^ (v1 * 0)))
    if v3 {
        v51 := GenS6{F0: 1}
        if (255 != 65535) {
            _ = len(v5)
        }
        var v52 int = int(v4)
        _ = v51
        _ = v52
    }
    fmt.Println(v4, v1)
    if false {
        var v53 int32 = -1
        var v54 bool = v3
        v53 = 0
        _ = v53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
