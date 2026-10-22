package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint64 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    if false {
        var v7 uint = 0
        v8 := GenS1{F0: 2147483647, F1: -2147483648, F2: -69}
        if false {
            v9 := GenS1{F0: 714, F1: -183, F2: -9223372036854775808}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 bool = ("a" <= ("foo bar" + ((("foo bar" + "") + "\n\t") + "foo bar")))
                var v12 float32 = 1e10
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint = v7
                v15 := GenS1{F0: 42, F1: -360, F2: 1}
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 int = -385
                    if (true && true) {
                        v18 := GenS1{F0: -1, F1: 2147483647, F2: (fi13 % ((v17) | 1))}
                        v7 = v7
                        v19 := []int{fi13, (fi16 + -9223372036854775808)}
                        v20 := GenS1{F0: 1, F1: 1, F2: v17}
                        v17 -= 585
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    if true {
                        v17 = (v17 & (fi13 % (((-(fi16))) | 1)))
                    }
                    _ = fi16
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            if (v7 <= (v7 * 42)) {
                var v21 float32 = ((1.5 / (0.1 + ((3.14159 - (0.1 + (-2.25 - (0.0 / 0.1)))) + 1e10))) / 1.0)
                if (!true) {
                    switch 0 {
                    case 0:
                        v7 += v7
                    case -1:
                        fmt.Println(v7)
                    case -122:
                        switch ((("" + "foo bar") + "\n\t") + "hello") {
                        case "foo bar":
                            var v22 uint64 = uint64(v7)
                            _ = v22
                        case "":
                            v23 := []int{}
                            _ = v23
                        default:
                            v21 = 1e10
                        }
                    }
                    v24 := []int{}
                    switch 732 {
                    case 328:
                        if (((0 <= 9223372036854775807) || (false || (true && (1 == 1)))) || (((float32(((0.0 + 3.14159) + 0.1)) / 1.0) - (v21 / ((v21 + float32(1e10)) * (v21 / 0.1)))) != 0.1)) {
                            var v25 int = 42
                            switch "\n\t" {
                            case "hello":
                                v26 := GenS2{F0: 2, F1: 187}
                                _ = v26
                            case "\n\t":
                                var v27 string = ("hello" + (((("a" + ("\n\t" + "hello")) + ("a" + "a")) + "foo bar") + ""))
                                _ = v27
                            default:
                                for fi28 := 0; fi28 < 12; fi28++ {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    _ = fi28
                                }
                            }
                            _ = v25
                        }
                    case 899:
                        _ = len(v24)
                    }
                    _ = v24
                }
                v29 := []int{}
                _ = v21
                _ = v29
            }
        }
        _ = v7
        _ = v8
    }
    switch 71 {
    case -910:
        var v30 bool = true
        _ = v30
    case 42:
        var v31 bool = (true || ((!false) && (0 != 32768)))
        _ = v31
    case -2147483648:
        var v32 uint16 = 65535
        _ = v32
    default:
        var v33 float32 = 0.1
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        if (!(-22 < 0)) {
            var v35 float32 = float32(0.1)
            _ = v35
        }
        fmt.Println(fi34)
        if (875 < 2147483648) {
            var v36 int16 = -894
            _ = v36
        }
        _ = fi34
    }
    return 42
}

func main() {
    var v1 uint64 = 790
    v2 := []uint32{42, 2147483648, 2, 4294967295}
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    switch 9223372036854775807 {
    case 42:
        var v37 uint8 = 54
        _ = v37
    case -231:
        var v38 int64 = 865
        _ = v38
    case 265:
        for fi39 := 0; fi39 < 16; fi39++ {
            v1 = (v1 - (v1 + 4294967295))
            _ = fi39
        }
    default:
        v2 = append(v2, 922)
    }
    switch (("" + ("\n\t" + (("" + ("\n\t" + "\n\t")) + "foo bar"))) + ("foo bar" + "foo bar")) {
    case "a":
        var v40 bool = true
        _ = v40
    }
    var v41 uint64 = (^(v1))
    for fi42 := 0; fi42 < 2; fi42++ {
        v43 := GenS2{F0: 127, F1: 442}
        _ = fi42
        _ = v43
    }
    v44 := GenS3{F0: 185, F1: 1}
    _ = len(v2)
    var v45 int32 = -864
    v46 := GenS4{F0: 128, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v46
}
