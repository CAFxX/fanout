package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int32; F2 int }
func genFn0(p0 uint, p1 int16, p2 int64) int64 {
    p1 = p1
    v6 := GenS1{F0: p0, F1: -1, F2: -2147483648}
    if (0 > -2147483648) {
        v7 := GenS2{F0: uint8(p0), F1: -209}
        switch 1 {
        case 127:
            v8 := []int{128}
            _ = v8
        case 32767:
            v9 := []int{-52, 759, -2147483648}
            _ = v9
        case 128:
            var v10 int = 9223372036854775807
            _ = v10
        }
        var v11 int8 = 1
        _ = v11
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 uint8, p1 uint64, p2 int) int64 {
    v12 := []int{(^(p2)), (p2 ^ p2), p2, (p2 << ((447) & 63))}
    fmt.Println(p0, p0, p0)
    p2 += p2
    switch "hello" {
    case "":
        p0 += (^(p0))
    case "hello":
        fmt.Println(p0, p0, p2)
    default:
        fmt.Println(p1, p1, p0)
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        p1 = p1
        v14 := GenS3{F0: p1, F1: 128}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 uint64 = 658
    var v2 uint64 = (^(v1))
    var v3 uint64 = 42
    var v4 uint16 = 42
    var v5 int64 = 913
    if true {
        var v15 int32 = -1
        _ = v15
    }
    if true {
        v16 := []int{}
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                switch (^(fi17)) {
                case 780:
                    if len(v16) > 0 {
                        v16[0] = fi18
                    }
                case 1:
                    _ = len(v16)
                default:
                    if (!false) {
                        if (!(-32768 < 0)) {
                            v16 = append(v16, fi18)
                        }
                    } else {
                        v3 += v1
                        var v19 string = "\n\t"
                        _ = len(v16)
                        _ = v19
                    }
                }
                var v20 bool = false
                _ = fi18
                _ = v20
            }
            v4 += v4
            _ = fi17
        }
        _ = v16
    } else {
        v21 := []int{32767, 1, 4294967295}
        switch (("foo bar" + ("foo bar" + "hello")) + ("\n\t" + "\n\t")) {
        case "":
            v3 = v2
        case "\n\t":
            var v22 float32 = -2.25
            _ = v22
        }
        v2 = (^(v3))
        _ = v21
    }
    v2 = (v2 & (v2 ^ 91))
    v23 := GenS1{F0: 4294967295, F1: 1000, F2: 42}
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(v5, v5, v4)
        var v25 int = (fi24 & fi24)
        v26 := GenS4{F0: 25}
        _ = fi24
        _ = v25
        _ = v26
    }
    v4 -= v4
    v3 = (v3 + 361)
    for fi27 := 0; fi27 < 12; fi27++ {
        for fi28 := 0; fi28 < 2; fi28++ {
            switch (^(fi28)) {
            case 9223372036854775807:
                if true {
                    v29 := GenS3{F0: (v3 - v2), F1: 128}
                    _ = v29
                } else {
                    v30 := GenS2{F0: 223, F1: 4294967295}
                    var v31 uint = 374
                    _ = v30
                    _ = v31
                }
            case 250:
                fmt.Println(v1, fi27)
            default:
                var v32 string = ("\n\t" + "")
                _ = v32
            }
            _ = fi28
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 float64 = 1e10
            v35 := GenS5{F0: 0, F1: 0, F2: (fi33 - (fi33 & int(v5)))}
            v36 := []int{-9223372036854775808, 2147483647, -396, fi33}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        v37 := []int{(fi27 & (fi27 | (fi27 ^ (fi27 + -1)))), fi27}
        _ = fi27
        _ = v37
    }
    var v38 uint64 = 42
    if false {
        if false {
            if false {
                for fi39 := 0; fi39 < 12; fi39++ {
                    v40 := []int{fi39, (fi39 % (((fi39 << ((42) & 63))) | 1)), (fi39 | 2147483647)}
                    var v41 int16 = 1000
                    var v42 float64 = 1.5
                    _ = fi39
                    _ = v40
                    _ = v41
                    _ = v42
                }
                v1 += v1
            } else {
                for fi43 := 0; fi43 < 13; fi43++ {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        v38 += v1
                        fmt.Println(fi44)
                        _ = fi44
                    }
                    _ = fi43
                }
                v1 -= v2
            }
            v45 := []int{}
            _ = v45
        }
        fmt.Println(v2)
        v46 := GenS1{F0: 18446744073709551615, F1: 42, F2: -2147483648}
        _ = v46
    }
    v47 := GenS2{F0: 127, F1: -9223372036854775808}
    if (82 == -1) {
        if true {
            v48 := []int{127, -2147483648}
            _ = v48
        }
    } else {
        v49 := GenS5{F0: 92, F1: 0, F2: 531}
        v5 = 1
        _ = v49
    }
    if (!false) {
        v1 = v1
        v3 -= 18446744073709551615
        v3 = v2
    }
    v50 := []int{42}
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v50
}
