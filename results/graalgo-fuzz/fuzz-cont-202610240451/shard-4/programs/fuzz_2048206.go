package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 int64) int16 {
    var v5 bool = ((-344 < 1) || ((!(!true)) && (!((p0 >> ((42) & 63)) != -531))))
    switch ("a" + "\n\t") {
    case "foo bar":
        fmt.Println(v5)
    }
    var v6 bool = v5
    fmt.Println(p0)
    v7 := []int{971}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 284
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = -2.25
    var v3 uint64 = 4294967295
    v4 := make(map[int]uint32)
    var v8 int64 = int64(v3)
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint32 = 0
        for fi11 := 0; fi11 < 4; fi11++ {
            switch int(v1) {
            case -438:
                if (!false) {
                    v12 := GenS1{F0: -1, F1: fi11, F2: 0}
                    _ = v12
                }
                v13 := GenS1{F0: int8(v3), F1: 1, F2: 609}
                v14 := GenS2{F0: 128, F1: 0}
                _ = v13
                _ = v14
            default:
                var v15 uint8 = 128
                _ = v15
            }
            _ = fi11
        }
        v2 -= (1e10 - -2.25)
        _ = fi9
        _ = v10
    }
    v16 := []int{0, 4294967295}
    delete(v4, 32768)
    var v17 int8 = 127
    var v18 uint16 = v1
    v4[127] = 4294967295
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 float64 = 0.0
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v1)
    switch 388 {
    case 128:
        if (1 == (v8 % (((v8 ^ (v8 - (v8 & (^(v8)))))) | 1))) {
            for fi22 := 0; fi22 < 8; fi22++ {
                switch -1 {
                case 32768:
                    v3 = (v3 % (((v3 >> ((0) & 63))) | 1))
                case -1:
                    if false {
                        var v23 uint16 = 2
                        v24 := GenS3{F0: "hello"}
                        if true {
                            for fi25 := 0; fi25 < 2; fi25++ {
                                var v26 uint32 = 1
                                _ = fi25
                                _ = v26
                            }
                            var v27 uint = 768
                            var v28 float32 = 1e10
                            _ = v27
                            _ = v28
                        }
                        _ = v23
                        _ = v24
                    } else {
                        var v29 int32 = 0
                        _ = v29
                    }
                case -134:
                    switch "foo bar" {
                    case "foo bar":
                        v2 += float64(1e10)
                    case "hello":
                        if len(v16) > 0 {
                            v16[0] = 32768
                        }
                    }
                }
                _ = fi22
            }
            v30 := GenS2{F0: 2, F1: -32768}
            v31 := GenS3{F0: "foo bar"}
            _ = v30
            _ = v31
        }
    case -217:
        v17 -= v17
    }
    v18 -= 65535
    v2 = v2
    switch ("" + "a") {
    case "hello":
        if ((1 != 1) && (688 >= 42)) {
            if len(v16) > 0 {
                v16[0] = -1
            }
            var v32 uint64 = v3
            var v33 int32 = 298
            _ = v32
            _ = v33
        } else {
            v34 := GenS4{F0: 2147483648, F1: 2147483647}
            _ = v34
        }
    case "":
        fmt.Println(v2)
    case "foo bar":
        var v35 float64 = (1e10 - float64(0.1))
        _ = v35
    }
    var v36 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v8
}
