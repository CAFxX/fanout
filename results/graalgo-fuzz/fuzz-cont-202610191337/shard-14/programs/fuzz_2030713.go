package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint32 }
func genFn0() int8 {
    var v7 int64 = -1
    var v8 int = -542
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 int64 = 503
    var v2 string = ""
    var v3 int64 = (v1 / (((v1 & v1)) | 1))
    v4 := make(map[string]int)
    var v5 uint = 0
    v6 := make(map[string]uint32)
    fmt.Println(v3)
    fmt.Println(v5)
    v5 = v5
    v9 := GenS1{F0: (v5 * 807), F1: 1000}
    delete(v6, (v2 + (v2 + v2)))
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS2{F0: -836, F1: -2147483648, F2: 0}
        if false {
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(fi10, fi12, fi10)
                v13 := GenS3{F0: 558, F1: 42}
                var v14 int32 = 1
                v15 := GenS1{F0: (^(v5)), F1: 1000}
                var v16 int64 = v3
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 float64 = 0.1
            _ = v17
        }
        var v18 int64 = 318
        _ = fi10
        _ = v11
        _ = v18
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 3; fi19++ {
        v2 = v2
        for fi20 := 0; fi20 < 13; fi20++ {
            if false {
                v2 = ("a" + (v2 + v2))
                if true {
                    fmt.Println(v5, v1, fi19)
                    var v21 float32 = ((-2.25 * 3.14159) * (0.1 * ((3.14159 / -2.25) + (0.1 - 0.0))))
                    switch (fi19 >> (((^(v5))) & 63)) {
                    case 4294967295:
                        fmt.Println(v1)
                    case 242:
                        var v22 int16 = 42
                        _ = v22
                    default:
                        var v23 uint32 = 2
                        _ = v23
                    }
                    _ = v21
                } else {
                    delete(v4, "")
                    v1 += int64(fi19)
                    if (1 < 128) {
                        v24 := GenS4{F0: 4294967295}
                        var v25 float32 = ((0.1 + 0.1) * 0.1)
                        var v26 bool = false
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                v27 := v4["hello"]
                _ = v27
            } else {
                var v28 int16 = -32768
                switch v2 {
                case "foo bar":
                    switch ("" + ("" + v2)) {
                    case "":
                        v29 := GenS5{F0: 28, F1: v3}
                        _ = v29
                    default:
                        if true {
                            if false {
                                var v30 uint = 2
                                _ = v30
                            }
                            if false {
                                v28 = v28
                            } else {
                                var v31 bool = (false || (true || false))
                                _ = v31
                            }
                            var v32 int = fi20
                            _ = v32
                        }
                    }
                }
                v33 := GenS2{F0: v3, F1: -552, F2: 1}
                _ = v28
                _ = v33
            }
            switch (fi20 * fi20) {
            case 2147483647:
                var v34 int = (fi20 & fi19)
                _ = v34
            }
            _ = fi20
        }
        _ = fi19
    }
    v35 := []int{-9223372036854775808}
    switch "" {
    case "":
        v36 := GenS1{F0: v5, F1: 1000}
        _ = v36
    case "a":
        var v37 int64 = (v3 >> ((837) & 63))
        _ = v37
    default:
        var v38 uint16 = 0
        _ = v38
    }
    v39 := GenS6{F0: 611}
    switch -6 {
    case 1:
        switch -817 {
        case 128:
            fmt.Println(v2)
        case -346:
            v3 += (v3 << ((v5) & 63))
        default:
            v40 := v4["hello"]
            _ = v40
        }
    }
    v41 := GenS4{F0: 0}
    fmt.Println(v5, v1, v2)
    for fi42 := 0; fi42 < 8; fi42++ {
        v5 -= v5
        v3 -= v3
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
