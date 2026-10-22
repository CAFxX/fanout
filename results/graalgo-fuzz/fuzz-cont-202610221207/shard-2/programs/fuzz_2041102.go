package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 int, p2 string) uint64 {
    p0 += p0
    if true {
        v5 := GenS1{F0: 1}
        p2 = p2
        _ = v5
    } else {
        for fi6 := 0; fi6 < 6; fi6++ {
            if (true || true) {
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 float32 = 0.0
                    p0 = 535
                    _ = fi7
                    _ = v8
                }
            }
            p2 = ("\n\t" + p2)
            _ = fi6
        }
    }
    p2 = (p2 + p2)
    p1 -= (-(p1))
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 677
}
func genFn1(p0 int) uint16 {
    var v10 uint8 = 0
    switch p0 {
    case -537:
        p0 -= (-(p0))
    case -494:
        for fi11 := 0; fi11 < 8; fi11++ {
            switch ("" + ("hello" + (((("\n\t" + "\n\t") + "foo bar") + ((("foo bar" + ("foo bar" + "hello")) + (("hello" + "\n\t") + "hello")) + "\n\t")) + "foo bar"))) {
            case "foo bar":
                v12 := GenS2{F0: 203}
                _ = v12
            case "":
                for fi13 := 0; fi13 < 17; fi13++ {
                    p0 = (fi13 + (fi11 << ((18446744073709551615) & 63)))
                    _ = fi13
                }
            case "hello":
                var v14 uint8 = v10
                _ = v14
            }
            v15 := GenS2{F0: 255}
            _ = fi11
            _ = v15
        }
    case 496:
        v10 -= v10
    }
    fmt.Println(p0, v10)
    switch (^(p0)) {
    case 2147483647:
        switch p0 {
        case 2147483647:
            switch 9223372036854775807 {
            case -324:
                for fi16 := 0; fi16 < 4; fi16++ {
                    v10 = (v10 / ((66) | 1))
                    _ = fi16
                }
            case -664:
                var v17 int8 = 18
                _ = v17
            case -1:
                for fi18 := 0; fi18 < 5; fi18++ {
                    p0 += fi18
                    switch p0 {
                    case 980:
                        p0 -= (p0 / (((fi18 * fi18)) | 1))
                    default:
                        v19 := GenS2{F0: 255}
                        _ = v19
                    }
                    v20 := []int{(fi18 >> ((0) & 63)), fi18, (p0 * (fi18 ^ p0))}
                    _ = fi18
                    _ = v20
                }
            default:
                var v21 uint = 18446744073709551615
                _ = v21
            }
        default:
            if (449 > 1) {
                if false {
                    fmt.Println(v10, p0)
                }
            }
        }
    }
    _ = p0
    _ = v10
    return 32767
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    fmt.Println(v1)
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 uint64 = genFn0(70, (fi22 & fi23), "")
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: 380}
        fmt.Println(fi22)
        _ = fi22
        _ = v25
    }
    v26 := GenS3{F0: -38, F1: (^(v1)), F2: 0}
    var v27 bool = (!(true || (!true)))
    v28 := v3[32768]
    _ = len(v2)
    v29 := []int{0}
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int = (fi30 / (((^(fi30))) | 1))
        v1 += (v1 << ((65535) & 63))
        _ = fi30
        _ = v31
    }
    fmt.Println(v28)
    fmt.Println(v28, v1)
    for fi32 := 0; fi32 < 4; fi32++ {
        if len(v29) > 0 {
            v29[0] = 1
        }
        _ = fi32
    }
    var v33 int = -1
    var v34 int32 = 1
    var v35 uint16 = 0
    var v36 int64 = 633
    for fi37 := 0; fi37 < 3; fi37++ {
        v28 = (v28 + v28)
        v38 := GenS1{F0: 12}
        _ = len(v4)
        _ = fi37
        _ = v38
    }
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
