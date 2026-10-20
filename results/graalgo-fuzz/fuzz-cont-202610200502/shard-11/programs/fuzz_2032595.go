package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint; F1 int32; F2 int }
func genFn0(p0 string, p1 uint8, p2 uint64) uint32 {
    p1 -= (p1 % (((p1 / (((p1 & p1)) | 1))) | 1))
    fmt.Println(p2)
    fmt.Println(p2, p1)
    for fi5 := 0; fi5 < 20; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{(^(fi5))}
            if true {
                var v8 bool = ("a" < (p0 + p0))
                var v9 uint64 = 42
                v10 := GenS1{F0: 127, F1: 109, F2: "foo bar"}
                _ = v10
                _ = v8
                _ = v9
            }
            v11 := GenS1{F0: p1, F1: -128, F2: p0}
            _ = fi6
            _ = v11
            _ = v7
        }
        switch fi5 {
        case 32767:
            if (false || false) {
                switch (fi5 ^ fi5) {
                case 0:
                    if (-2.25 == 0.0) {
                        v12 := []int{2147483647, (fi5 - fi5), (-(fi5)), fi5}
                        for fi13 := 0; fi13 < 11; fi13++ {
                            p2 = (p2 & p2)
                            switch "a" {
                            case "a":
                                v14 := GenS2{F0: 449, F1: 255, F2: 2147483647}
                                _ = v14
                            case "hello":
                                v15 := GenS3{F0: 65535, F1: -1}
                                _ = v15
                            }
                            p0 = p0
                            var v16 float64 = (0.0 - 0.0)
                            fmt.Println(v16, v16, fi5)
                            _ = fi13
                            _ = v16
                        }
                        _ = v12
                    }
                default:
                    var v17 int64 = 1
                    _ = v17
                }
                p2 -= p2
                var v18 int = (fi5 >> ((986) & 63))
                _ = v18
            } else {
                for fi19 := 0; fi19 < 3; fi19++ {
                    if ((fi19 * fi19) >= 1) {
                        v20 := []int{33, -2147483648, fi19, fi19}
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 65535, F1: -128}
                        switch (fi5 << ((139) & 63)) {
                        case 4294967295:
                            v22 := []int{fi5}
                            _ = v22
                        case 161:
                            var v23 string = (p0 + "a")
                            _ = v23
                        case 511:
                            var v24 int32 = -516
                            _ = v24
                        }
                        _ = v21
                    }
                    v25 := []int{(fi5 + (fi19 * (fi19 / (((fi5 / (((fi19 % (((fi19 ^ fi5)) | 1))) | 1))) | 1)))), (-(fi5))}
                    fmt.Println(fi19, p1)
                    _ = fi19
                    _ = v25
                }
            }
        case 4294967295:
            var v26 uint = 866
            _ = v26
        }
        if false {
            v27 := []int{(fi5 % (((^(fi5))) | 1)), 1, fi5}
            _ = v27
        } else {
            v28 := GenS1{F0: 133, F1: -41, F2: (((p0 + ("" + "\n\t")) + "a") + (p0 + ((p0 + ("" + "\n\t")) + (("hello" + p0) + p0))))}
            v29 := GenS2{F0: 563, F1: 42, F2: 609}
            _ = v28
            _ = v29
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return 233
}

func main() {
    v1 := []uint32{65535}
    var v2 float64 = -2.25
    v3 := []int{17, 974}
    var v4 bool = true
    v30 := GenS2{F0: 893, F1: 1, F2: -2147483648}
    for fi31 := 0; fi31 < 19; fi31++ {
        _ = len(v1)
        _ = fi31
    }
    var v32 uint32 = 498
    var v33 int16 = -125
    if v4 {
        fmt.Println(v32, v2)
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            _ = len(v1)
            v36 := GenS4{F0: "a", F1: 65535}
            if false {
                v37 := GenS5{F0: 65535, F1: 42, F2: fi35}
                var v38 int64 = 0
                var v39 int32 = -2147483648
                _ = v37
                _ = v38
                _ = v39
            } else {
                _ = len(v3)
                var v40 uint64 = 2
                if ((fi35 >= fi34) || (19 > 1)) {
                    v4 = v4
                    _ = len(v3)
                    var v41 string = ("hello" + ("\n\t" + "foo bar"))
                    v4 = ((1.5 <= 3.14159) || (!(v33 > v33)))
                    switch fi35 {
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v41 = ("a" + ("foo bar" + v41))
                    _ = v41
                }
                _ = v40
            }
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    if false {
        var v42 string = ""
        var v43 int = 0
        _ = v42
        _ = v43
    } else {
        fmt.Println(v4, v33, v2)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    }
    v4 = (-2147483648 == 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
}
