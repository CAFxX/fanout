package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 uint16, p1 int16, p2 int32) uint64 {
    switch int(p2) {
    case 0:
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int8 = 105
                if ((false && (42 > 42)) && (544 == 0)) {
                    var v10 int8 = -94
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            if false {
                var v11 int64 = -1
                p2 -= p2
                if (!(false || (true && true))) {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(v11, v11)
                        fmt.Println(fi7)
                        v13 := GenS1{F0: p1}
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v11
            } else {
                p1 = (^(p1))
                fmt.Println(fi7, fi7)
                var v14 uint32 = 4294967295
                _ = v14
            }
            _ = fi7
        }
    case 351:
        for fi15 := 0; fi15 < 2; fi15++ {
            if (42 <= 2) {
                p2 += p2
                for fi16 := 0; fi16 < 16; fi16++ {
                    if (793 != 619) {
                        if (true && (183 < (p0 ^ (p0 * (p0 - (p0 >> ((65535) & 63))))))) {
                            v17 := []int{}
                            switch (fi16 >> ((881) & 63)) {
                            case 32768:
                                switch fi16 {
                                case 336:
                                    var v18 uint = 27
                                    p0 -= 42
                                    _ = v18
                                case 4294967295:
                                    v19 := GenS1{F0: 839}
                                    _ = v19
                                default:
                                    var v20 uint32 = 1
                                    _ = v20
                                }
                            default:
                                p2 -= -749
                            }
                            _ = v17
                        }
                    } else {
                        var v21 uint16 = p0
                        fmt.Println(fi16)
                        p2 = (^(p2))
                        _ = v21
                    }
                    if false {
                        switch -2147483648 {
                        case -316:
                            var v22 uint32 = 42
                            _ = v22
                        }
                        if true {
                            var v23 float32 = 3.14159
                            _ = v23
                        } else {
                            p1 = (^(p1))
                            var v24 int64 = 590
                            p0 = p0
                            var v25 int = fi16
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi16
                }
                v26 := []int{1, 127, (fi15 % (((fi15 - fi15)) | 1))}
                _ = v26
            }
            v27 := []int{fi15}
            _ = fi15
            _ = v27
        }
    default:
        var v28 int64 = -9223372036854775808
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 17; fi29++ {
            p2 += p2
            v30 := GenS2{F0: 13, F1: 4294967295}
            for fi31 := 0; fi31 < 14; fi31++ {
                if false {
                    var v32 uint32 = 0
                    v33 := []int{fi31}
                    _ = v32
                    _ = v33
                } else {
                    if (42 > 18446744073709551615) {
                        var v34 int8 = 42
                        _ = v34
                    }
                    v35 := []int{496, (fi29 >> ((18446744073709551615) & 63))}
                    switch "a" {
                    case "foo bar":
                        switch fi29 {
                        case -658:
                            switch (fi29 ^ fi31) {
                            case -1:
                                p1 = (p1 ^ p1)
                            }
                        case -2147483648:
                            switch (fi31 << ((356) & 63)) {
                            case 2147483647:
                                if (true && (-70 < -58)) {
                                    p0 -= 65535
                                    var v36 bool = (0 < 1)
                                    var v37 uint16 = p0
                                    if true {
                                        p0 = (v37 & v37)
                                        v38 := GenS3{F0: 909}
                                        _ = v38
                                    }
                                    var v39 int16 = p1
                                    var v40 uint32 = 482
                                    _ = v36
                                    _ = v37
                                    _ = v39
                                    _ = v40
                                }
                            default:
                                var v41 uint8 = 0
                                _ = v41
                            }
                        case 127:
                            if false {
                                p2 -= -1
                                p1 = 0
                                v42 := GenS4{F0: 364, F1: 1, F2: (p1 + p1)}
                                if (-2147483648 <= (p2 ^ (p2 << ((0) & 63)))) {
                                    var v43 int16 = p1
                                    _ = v43
                                } else {
                                    fmt.Println(p0)
                                }
                                v44 := GenS3{F0: 4294967295}
                                _ = v42
                                _ = v44
                            }
                        }
                    case "":
                        p1 += (p1 << ((4294967295) & 63))
                    case "hello":
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                    }
                    _ = v35
                }
                var v45 int32 = 1
                v46 := GenS1{F0: 42}
                _ = fi31
                _ = v45
                _ = v46
            }
            _ = fi29
            _ = v30
        }
        p1 += p1
        var v47 uint32 = 614
        _ = v47
    } else {
        for fi48 := 0; fi48 < 17; fi48++ {
            p2 = -559
            _ = fi48
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint64 = 274
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    var v4 int16 = 755
    var v5 int = -997
    var v6 int = v5
    var v49 uint16 = 255
    for fi50 := 0; fi50 < 15; fi50++ {
        v51 := v3[2]
        fmt.Println(v1)
        _ = fi50
        _ = v51
    }
    v4 = (v4 ^ (v4 ^ (v4 - 743)))
    v49 -= 2
    v52 := v2[(v6 ^ (-(v5)))]
    var v53 int16 = (v4 & (v4 % (((v4 - 89)) | 1)))
    v49 += (^(v49))
    if ((v49 - v49) != v49) {
        v54 := []int{(v6 * v6)}
        _ = v54
    } else {
        var v55 int = (^(v5))
        fmt.Println(v55, v53)
        v56 := GenS5{F0: 255, F1: 42, F2: "a"}
        _ = v55
        _ = v56
    }
    v52 = (false && true)
    v5 = v5
    for fi57 := 0; fi57 < 4; fi57++ {
        var v58 int = (v5 - fi57)
        _ = fi57
        _ = v58
    }
    var v59 uint8 = 2
    var v60 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v53
    _ = v59
    _ = v6
    _ = v60
}
