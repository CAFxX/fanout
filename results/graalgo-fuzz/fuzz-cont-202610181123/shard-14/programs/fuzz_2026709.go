package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483647, 0, 701}
    v3 := []int{-305}
    var v4 int = 1
    var v5 int16 = 32767
    v6 := make(map[int]int)
    v7 := v6[v4]
    if true {
        var v8 int64 = 1
        _ = v8
    } else {
        v5 += 42
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 0
        if true {
            var v11 float32 = float32(0.1)
            _ = v11
        }
        v12 := GenS1{F0: 0}
        _ = fi9
        _ = v10
        _ = v12
    }
    _ = len(v1)
    if true {
        _ = len(v3)
        v4 -= (v4 % ((v4) | 1))
        if true {
            var v13 float64 = 3.14159
            for fi14 := 0; fi14 < 5; fi14++ {
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 uint32 = 0
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            _ = v13
        } else {
            switch (v7 * (v7 * (v7 << ((937) & 63)))) {
            case 0:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v5 += 0
            fmt.Println(v4, v7)
        }
    } else {
        v17 := GenS2{F0: 4294967295}
        if (!(1 > -374)) {
            fmt.Println(v5, v5)
            v5 = v5
            for fi18 := 0; fi18 < 18; fi18++ {
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 uint32 = 2147483647
                    _ = fi19
                    _ = v20
                }
                if (true || true) {
                    delete(v1, (^(fi18)))
                    var v21 int64 = 871
                    if false {
                        for fi22 := 0; fi22 < 9; fi22++ {
                            delete(v6, (fi18 | (^(fi18))))
                            if len(v3) > 0 {
                                v3[0] = v7
                            }
                            _ = fi22
                        }
                        var v23 int16 = -404
                        v24 := GenS3{F0: 604, F1: 128}
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 8; fi25++ {
                            fmt.Println(v4, v5)
                            v3 = append(v3, v4)
                            v21 = v21
                            _ = fi25
                        }
                    }
                    _ = v21
                } else {
                    var v26 int64 = -225
                    _ = v26
                }
                _ = fi18
            }
        } else {
            v7 += 657
            v7 -= (v7 ^ -850)
            v27 := GenS3{F0: 4294967295, F1: 0}
            _ = v27
        }
        v28 := GenS1{F0: v5}
        _ = v17
        _ = v28
    }
    v29 := GenS1{F0: (v5 >> ((2) & 63))}
    fmt.Println(v4, v4)
    if false {
        for fi30 := 0; fi30 < 7; fi30++ {
            switch "a" {
            case "":
                v31 := GenS2{F0: 42}
                _ = v31
            case "foo bar":
                v32 := GenS3{F0: 230, F1: 87}
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v5 += v5
                _ = fi33
            }
            v34 := GenS3{F0: 2, F1: 2}
            _ = fi30
            _ = v34
        }
        switch "foo bar" {
        case "a":
            for fi35 := 0; fi35 < 8; fi35++ {
                fmt.Println(v4, fi35)
                v7 = v7
                v2 = append(v2, 42)
                _ = fi35
            }
        case "\n\t":
            var v36 uint8 = 1
            _ = v36
        case "hello":
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 uint64 = 2
                _ = fi37
                _ = v38
            }
        default:
            var v39 uint64 = 929
            _ = v39
        }
        if (!(!false)) {
            fmt.Println(v5, v4)
            switch (v4 << ((65535) & 63)) {
            case 9223372036854775807:
                var v40 uint16 = 2
                _ = v40
            }
        }
    }
    v7 -= v4
    switch (^(v7)) {
    case -1:
        if false {
            for fi41 := 0; fi41 < 10; fi41++ {
                v42 := GenS2{F0: 1}
                fmt.Println(v5, v5, v7)
                _ = len(v3)
                _ = fi41
                _ = v42
            }
        }
    case 893:
        switch "a" {
        case "a":
            var v43 string = "foo bar"
            _ = v43
        default:
            var v44 int16 = v5
            _ = v44
        }
    case -255:
        v5 -= (v5 * v5)
    default:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
