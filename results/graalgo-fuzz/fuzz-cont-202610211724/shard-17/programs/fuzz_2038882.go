package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 int16 }
func genFn0(p0 uint64) uint64 {
    var v4 int = -418
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{(fi5 + v4), 0, (fi5 << ((1) & 63)), (v4 ^ (fi5 + (fi5 + fi5)))}
        var v7 int = v4
        p0 = p0
        _ = fi5
        _ = v6
        _ = v7
    }
    switch (v4 + v4) {
    case -9223372036854775808:
        if (!true) {
            var v8 uint64 = p0
            var v9 uint64 = (^(p0))
            _ = v8
            _ = v9
        } else {
            var v10 uint32 = 42
            fmt.Println(v4, v4)
            v11 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: 0}
            _ = v10
            _ = v11
        }
    }
    var v12 int64 = 9223372036854775807
    v13 := GenS2{F0: 79}
    _ = p0
    _ = v12
    _ = v13
    _ = v4
    return 0
}

func main() {
    var v1 bool = true
    var v2 int16 = -1
    var v3 int64 = -964
    var v14 int64 = (v3 / ((42) | 1))
    fmt.Println(v3)
    if (false || (!v1)) {
        var v15 int64 = v14
        v16 := []int{32768, 2147483647, 127}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v1, v14, v1)
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        switch -1 {
        case 875:
            v14 += (v14 >> ((224) & 63))
        case 32768:
            switch fi17 {
            case 1:
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS2{F0: 0}
                    var v20 int32 = 2147483647
                    fmt.Println(v3, v2)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case -445:
                var v21 int8 = -1
                _ = v21
            }
        case 9223372036854775807:
            v22 := GenS2{F0: -1}
            _ = v22
        default:
            fmt.Println(v3, v2, v3)
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := []int{int(v3), fi23}
            var v25 int8 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi17
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        if v1 {
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := GenS2{F0: 486}
                switch 32768 {
                case 2147483647:
                    v14 = v3
                case 9223372036854775807:
                    var v29 int64 = -1
                    _ = v29
                }
                v30 := GenS2{F0: 2147483647}
                switch (^(fi27)) {
                case 127:
                    v2 -= (v2 << ((1) & 63))
                default:
                    var v31 uint64 = 2
                    _ = v31
                }
                v32 := GenS2{F0: -2147483648}
                _ = fi27
                _ = v28
                _ = v30
                _ = v32
            }
        } else {
            if true {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v2 = v2
                        var v34 int16 = v2
                        var v35 string = ("hello" + ("foo bar" + ("hello" + "foo bar")))
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                case "a":
                    if (v2 > (v2 - v2)) {
                        v1 = (0 == -478)
                        var v36 int16 = (v2 / (((-(v2))) | 1))
                        _ = v36
                    } else {
                        var v37 float64 = 1e10
                        var v38 uint8 = 255
                        _ = v37
                        _ = v38
                    }
                }
                var v39 uint8 = 128
                var v40 int16 = -420
                _ = v39
                _ = v40
            }
        }
        _ = fi26
    }
    v3 = 9223372036854775807
    v1 = (32768 > 32767)
    var v41 float32 = 0.0
    v42 := []int{}
    v43 := GenS3{F0: -882, F1: (v14 + v3), F2: 0}
    var v44 int64 = -788
    v45 := GenS3{F0: (v44 % ((v14) | 1)), F1: -939, F2: v2}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
