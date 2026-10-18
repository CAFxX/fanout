package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32) int32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int16 = 0
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 string = ""
            var v11 uint8 = 128
            fmt.Println(fi9, v8)
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    switch (("a" + "") + ("hello" + "a")) {
    case "foo bar":
        p0 = (p0 - 42)
    case "a":
        var v14 uint32 = 2147483647
        _ = v14
    case "hello":
        switch "" {
        case "hello":
            var v15 int32 = (^(p0))
            _ = v15
        }
    }
    if (((-2.25 - (0.1 - 0.1)) - float32(0.0)) >= ((1e10 + -2.25) * float32(1.5))) {
        p0 = (p0 % ((p0) | 1))
        var v16 int32 = 0
        v17 := GenS1{F0: 204}
        _ = v16
        _ = v17
    }
    p0 = 2147483647
    _ = p0
    return (-(p0))
}

func main() {
    v1 := []int{-9223372036854775808, 42}
    v2 := make(map[int]bool)
    var v3 int16 = -918
    var v4 float32 = 1.5
    var v5 int32 = 2147483647
    var v6 uint32 = 42
    var v18 bool = (true || (v5 >= (v5 / ((genFn0(v5)) | 1))))
    if (!(v4 <= 0.1)) {
        if (v3 > 1) {
            v5 = (v5 >> ((77) & 63))
            for fi19 := 0; fi19 < 6; fi19++ {
                v18 = (v4 > v4)
                _ = fi19
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v20 := GenS1{F0: 0}
        _ = v20
    }
    if (1 != 32767) {
        var v21 int16 = (v3 & (v3 - v3))
        v22 := GenS1{F0: 334}
        delete(v2, 4294967295)
        _ = v21
        _ = v22
    } else {
        var v23 float32 = v4
        _ = v23
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 int8 = -18
        _ = fi24
        _ = v25
    }
    var v26 int64 = 1
    for fi27 := 0; fi27 < 5; fi27++ {
        if v18 {
            switch 2147483647 {
            case 595:
                var v28 int64 = v26
                _ = v28
            case 128:
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 string = "\n\t"
                    v1 = append(v1, (fi27 >> ((581) & 63)))
                    if v18 {
                        if v18 {
                            switch fi29 {
                            case -1:
                                var v31 uint32 = v6
                                _ = v31
                            case 2147483647:
                                var v32 int64 = int64(v6)
                                _ = v32
                            }
                            var v33 int64 = (v26 >> ((4294967295) & 63))
                            _ = v33
                        }
                    } else {
                        var v34 int16 = v3
                        v35 := GenS2{F0: 1}
                        _ = v34
                        _ = v35
                    }
                    fmt.Println(v30, fi27)
                    _ = fi29
                    _ = v30
                }
            case 179:
                delete(v2, 676)
            default:
                v6 -= (v6 - 2147483648)
            }
            var v36 uint = 50
            var v37 uint16 = 2
            _ = v36
            _ = v37
        }
        v5 = 914
        v38 := GenS1{F0: (v6 % (((^(v6))) | 1))}
        _ = fi27
        _ = v38
    }
    v26 = 9223372036854775807
    if ((!v18) && true) {
        fmt.Println(v5, v4, v4)
        for fi39 := 0; fi39 < 17; fi39++ {
            v3 -= 0
            _ = fi39
        }
    } else {
        fmt.Println(v3, v5)
    }
    fmt.Println(v3, v3)
    if ((1e10 * v4) < ((-2.25 * v4) / 1.0)) {
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    _ = len(v1)
    var v40 string = "a"
    for fi41 := 0; fi41 < 12; fi41++ {
        switch ("a" + "foo bar") {
        case "":
            var v42 bool = (0.1 <= float64(((v4 + (3.14159 + v4)) - (((v4 / 3.14159) / v4) * 1.5))))
            _ = v42
        default:
            v43 := GenS3{F0: 631}
            _ = v43
        }
        if v18 {
            var v44 int64 = v26
            _ = v44
        }
        _ = fi41
    }
    v5 -= (v5 & (v5 & genFn0(v5)))
    if (v26 > v26) {
        v6 -= (v6 ^ v6)
        for fi45 := 0; fi45 < 13; fi45++ {
            v46 := GenS1{F0: (^(v6))}
            _ = fi45
            _ = v46
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
