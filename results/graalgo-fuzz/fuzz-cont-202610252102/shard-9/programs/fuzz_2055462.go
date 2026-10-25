package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16 }
func genFn0(p0 int8, p1 string, p2 uint64) int32 {
    v6 := GenS1{F0: 1, F1: -986}
    for fi7 := 0; fi7 < 17; fi7++ {
        switch (fi7 | int(p0)) {
        case 341:
            if true {
                p0 = (p0 % (((p0 >> ((0) & 63))) | 1))
                if true {
                    if true {
                        v8 := GenS2{F0: 42}
                        if true {
                            p0 += p0
                            p0 -= (p0 / (((p0 / (((p0 ^ p0)) | 1))) | 1))
                            var v9 uint32 = 42
                            _ = v9
                        }
                        v10 := []int{fi7}
                        _ = v10
                        _ = v8
                    } else {
                        fmt.Println(p2, p0)
                        for fi11 := 0; fi11 < 19; fi11++ {
                            p1 = p1
                            v12 := []int{(^(fi7)), 1, -805, fi7}
                            v12 = append(v12, 42)
                            _ = fi11
                            _ = v12
                        }
                        v13 := []int{581, (^(fi7)), fi7, (fi7 / ((fi7) | 1))}
                        _ = v13
                    }
                }
            }
        case 42:
            fmt.Println(p0, fi7)
        case -664:
            v14 := GenS1{F0: 0, F1: -995}
            _ = v14
        default:
            p1 = "\n\t"
        }
        if true {
            var v15 int32 = -679
            var v16 uint32 = 2147483647
            v17 := GenS2{F0: 405}
            _ = v15
            _ = v16
            _ = v17
        } else {
            p0 = p0
            p2 = (p2 | p2)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -148
}

func main() {
    var v1 uint8 = 2
    var v2 int8 = -71
    var v3 uint64 = 18446744073709551615
    var v4 int8 = (v2 + v2)
    var v5 float32 = (1e10 * 3.14159)
    var v18 uint32 = 1
    if (true && true) {
        v4 = v2
        switch -619 {
        case -404:
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 uint64 = (^(v3))
                var v21 int32 = -2147483648
                v22 := []int{127}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        default:
            v23 := GenS1{F0: (^(v1)), F1: 0}
            _ = v23
        }
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 uint64 = (v3 >> ((4294967295) & 63))
            _ = fi24
            _ = v25
        }
    }
    v2 = (v2 - (v4 & v4))
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    var v28 int16 = 1000
    switch (("hello" + "foo bar") + "\n\t") {
    case "foo bar":
        var v29 int64 = int64(v28)
        _ = v29
    case "\n\t":
        v30 := GenS3{F0: 0, F1: (^(v28))}
        _ = v30
    case "hello":
        v4 += v2
    }
    fmt.Println(v18, v3, v2)
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 bool = true
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := []int{(fi33 ^ (fi33 - fi31)), (fi31 & fi31)}
            var v35 uint64 = 4294967295
            v3 = (^(v35))
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi31
        _ = v32
    }
    switch 237 {
    case -1:
        var v36 uint32 = uint32(v3)
        _ = v36
    case 902:
        if (!true) {
            v37 := GenS1{F0: 163, F1: v28}
            fmt.Println(v28, v3)
            _ = v37
        } else {
            var v38 uint64 = 4294967295
            v3 = (v3 * (v3 * (v3 / ((741) | 1))))
            _ = v38
        }
    }
    v39 := GenS2{F0: v18}
    fmt.Println(v2, v2)
    v40 := []int{-179, -2147483648, 245}
    switch -2147483648 {
    case -533:
        var v41 string = "foo bar"
        _ = v41
    case -169:
        var v42 int = 4294967295
        _ = v42
    default:
        v1 += (v1 - 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
