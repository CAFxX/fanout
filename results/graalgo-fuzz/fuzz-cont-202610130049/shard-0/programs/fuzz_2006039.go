package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 string }
type GenS4 struct { F0 int64 }
func genFn0() int8 {
    var v4 bool = true
    var v5 int64 = -882
    _ = v4
    _ = v5
    return 107
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    v3 := make(map[int]string)
    v2 += 4294967295
    if false {
        switch -769 {
        case 2147483647:
            v2 -= (v2 + (^(v2)))
        case 42:
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 uint = v2
                _ = fi6
                _ = v7
            }
        }
        var v8 uint16 = 32767
        _ = v8
    } else {
        var v9 float32 = 0.0
        v9 += 0.1
        _ = v9
    }
    var v10 int64 = 1
    var v11 int = 127
    v12 := []int{-2147483648, 2147483647}
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi15
        }
        v13 -= 42
        for fi16 := 0; fi16 < 9; fi16++ {
            if false {
                v17 := GenS1{F0: 0, F1: (v2 | 0), F2: -888}
                var v18 int = (fi16 << ((v2) & 63))
                var v19 uint64 = 802
                _ = v17
                _ = v18
                _ = v19
            }
            var v20 string = (((("\n\t" + "foo bar") + ("hello" + "\n\t")) + ("\n\t" + "a")) + "a")
            var v21 int32 = (v13 & v13)
            if true {
                v22 := GenS1{F0: -1, F1: v2, F2: 42}
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v22
            }
            if (3.14159 != 0.0) {
                v23 := GenS1{F0: -128, F1: (^(v2)), F2: -32768}
                if (914 != 212) {
                    v24 := GenS2{F0: 0}
                    var v25 int64 = 42
                    switch v11 {
                    case -728:
                        fmt.Println(v25)
                    case -63:
                        var v26 int32 = (v13 * v13)
                        _ = v26
                    default:
                        var v27 uint16 = 1
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                }
                if true {
                    if true {
                        delete(v3, (v11 ^ fi16))
                    }
                    v28 := GenS3{F0: fi16, F1: 42, F2: v20}
                    _ = v28
                }
                _ = v23
            } else {
                if true {
                    fmt.Println(v13)
                    if false {
                        var v29 int8 = 0
                        switch (^(fi14)) {
                        case -187:
                            _ = len(v12)
                        case -2147483648:
                            v20 = (("a" + (((((((v20 + "\n\t") + "a") + v20) + ("a" + v20)) + (v20 + v20)) + v20) + v20)) + (v20 + ""))
                        case 32767:
                            fmt.Println(v21, fi16, v13)
                        default:
                            if (!(false || false)) {
                                var v30 int16 = 32767
                                delete(v1, 894)
                                var v31 int64 = (v10 % (((v10 * v10)) | 1))
                                _ = v30
                                _ = v31
                            }
                        }
                        _ = v29
                    } else {
                        v13 += 2147483647
                    }
                    switch "hello" {
                    case "foo bar":
                        switch (v20 + "") {
                        case "foo bar":
                            v20 = v20
                        }
                    default:
                        v20 = "\n\t"
                    }
                    var v32 uint = (v2 << ((v2) & 63))
                    _ = v32
                }
                switch "\n\t" {
                case "\n\t":
                    v33 := GenS2{F0: 706}
                    _ = v33
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            }
            _ = fi16
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    v34 := GenS4{F0: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
}
