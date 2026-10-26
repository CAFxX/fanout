package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint8; F2 int64 }
func genFn0(p0 string) uint {
    p0 = "a"
    for fi6 := 0; fi6 < 12; fi6++ {
        switch ((((p0 + ((p0 + "hello") + (p0 + p0))) + p0) + "hello") + "foo bar") {
        case "":
            var v7 string = p0
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: 0, F1: (fi6 * fi6), F2: 2147483647}
            _ = v8
        default:
            v9 := []int{fi6}
            _ = v9
        }
        fmt.Println(fi6, p0)
        _ = fi6
    }
    _ = p0
    return 561
}

func main() {
    var v1 int32 = -1
    var v2 float32 = 0.0
    var v3 int8 = 4
    v4 := make(map[uint32]int)
    var v5 string = ("" + "\n\t")
    if (0 < 0) {
        for fi10 := 0; fi10 < 12; fi10++ {
            if true {
                switch (v5 + v5) {
                case "hello":
                    var v11 uint32 = 424
                    _ = v11
                case "\n\t":
                    var v12 int32 = int32(fi10)
                    _ = v12
                case "foo bar":
                    if (false || (false && (-9223372036854775808 == -456))) {
                        if ((1.5 / float32(1e10)) <= v2) {
                            v1 += v1
                            fmt.Println(v1, v1)
                        } else {
                            fmt.Println(v2, v2)
                        }
                        v3 -= v3
                    }
                default:
                    v13 := GenS1{F0: 42, F1: (fi10 - (fi10 | (fi10 & 638))), F2: (v1 << ((42) & 63))}
                    _ = v13
                }
                var v14 uint32 = 1
                v15 := GenS1{F0: 42, F1: fi10, F2: 0}
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                v3 = 100
                _ = fi16
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                if (255 < 75) {
                    v18 := GenS1{F0: int64(v3), F1: fi17, F2: -1}
                    fmt.Println(fi10, fi10)
                    _ = v18
                } else {
                    v1 -= -1
                }
                if (v1 == (-(v1))) {
                    v19 := GenS2{F0: 2, F1: 381}
                    var v20 int64 = 150
                    var v21 uint32 = 346
                    _ = v19
                    _ = v20
                    _ = v21
                }
                var v22 uint64 = 0
                _ = fi17
                _ = v22
            }
            _ = fi10
        }
        var v23 int32 = v1
        _ = v23
    } else {
        switch v5 {
        case "hello":
            fmt.Println(v1, v3, v3)
        case "a":
            for fi24 := 0; fi24 < 3; fi24++ {
                v5 = v5
                var v25 int16 = 852
                v26 := []int{}
                _ = fi24
                _ = v25
                _ = v26
            }
        case "foo bar":
            v3 = v3
        }
        var v27 int64 = -1
        _ = v27
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        _ = len(v4)
        if true {
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := GenS3{F0: (2 < 18446744073709551615)}
                var v31 int64 = 169
                _ = fi29
                _ = v30
                _ = v31
            }
            fmt.Println(v5)
            delete(v4, 867)
        }
        if false {
            v5 = ((v5 + ((("\n\t" + v5) + (v5 + v5)) + ("foo bar" + (("" + ("a" + "hello")) + v5)))) + v5)
            if (!false) {
                var v32 bool = false
                _ = v32
            } else {
                var v33 int8 = (v3 + 42)
                if true {
                    v5 = "foo bar"
                    for fi34 := 0; fi34 < 10; fi34++ {
                        var v35 uint32 = 0
                        _ = fi34
                        _ = v35
                    }
                } else {
                    v2 += ((float32(1e10) * float32(3.14159)) / (v2 - (v2 / v2)))
                    var v36 uint32 = 2147483647
                    v37 := []int{}
                    _ = v36
                    _ = v37
                }
                _ = v33
            }
        } else {
            v4[2] = fi28
            fmt.Println(v1, fi28, v1)
        }
        _ = fi28
    }
    if true {
        v5 = v5
        for fi38 := 0; fi38 < 4; fi38++ {
            switch fi38 {
            case 128:
                if false {
                    for fi39 := 0; fi39 < 8; fi39++ {
                        var v40 bool = (!(2 < 1))
                        var v41 uint = 1
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    v42 := GenS3{F0: true}
                    var v43 uint16 = 969
                    _ = v42
                    _ = v43
                } else {
                    var v44 uint8 = 98
                    _ = v44
                }
            case -2147483648:
                v45 := GenS1{F0: -9223372036854775808, F1: (fi38 / (((-(fi38))) | 1)), F2: (^(v1))}
                _ = v45
            }
            _ = fi38
        }
    }
    v3 = (v3 / ((v3) | 1))
    delete(v4, 2147483647)
    v46 := GenS1{F0: 0, F1: -9223372036854775808, F2: (v1 - v1)}
    v3 = (-(v3))
    v47 := GenS4{F0: 4294967295, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
