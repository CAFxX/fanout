package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 int64, p2 int16) uint {
    p2 = 1000
    p1 -= (^(p1))
    v5 := []int{722, 32768, 2147483647, -9223372036854775808}
    p2 += 260
    p1 = (p1 - -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    v1 := []int{127, 127}
    var v2 int16 = 0
    var v3 int8 = 1
    v4 := []int{42}
    if (!(("foo bar" > "") && false)) {
        for fi6 := 0; fi6 < 4; fi6++ {
            switch (("" + "") + "hello") {
            case "":
                if (-9223372036854775808 < 42) {
                    v3 = 0
                    for fi7 := 0; fi7 < 17; fi7++ {
                        var v8 int32 = 0
                        var v9 int64 = 953
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    v3 += (v3 ^ -28)
                }
            case "hello":
                var v10 uint32 = 0
                _ = v10
            }
            _ = fi6
        }
        v2 = 1000
        var v11 uint16 = 507
        _ = v11
    }
    if false {
        if true {
            var v12 int64 = -928
            for fi13 := 0; fi13 < 16; fi13++ {
                switch (fi13 / ((32768) | 1)) {
                case 2147483647:
                    v3 = 100
                case -2147483648:
                    var v14 uint64 = 26
                    _ = v14
                case 775:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v12 += (v12 / ((v12) | 1))
                        fmt.Println(fi13)
                        if (true || true) {
                            if (!true) {
                                var v16 int = (^(fi13))
                                if false {
                                    var v17 int = v16
                                    if true {
                                        fmt.Println(v17, v16, fi13)
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            v2 = v2
                                        }
                                    }
                                    v12 = (v12 % ((v12) | 1))
                                    _ = v17
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        v4 = append(v4, 32767)
                                    default:
                                        v2 = (^(v2))
                                    }
                                    var v18 int32 = 2147483647
                                    _ = v18
                                }
                                v3 += 1
                                v12 = v12
                                _ = v16
                            } else {
                                _ = len(v4)
                                var v19 int64 = (v12 - (v12 / (((^(v12))) | 1)))
                                _ = v19
                            }
                        }
                        _ = fi15
                    }
                }
                _ = fi13
            }
            if true {
                v3 += (v3 + v3)
            }
            _ = v12
        }
        var v20 int32 = -1
        if len(v1) > 0 {
            v1[0] = 127
        }
        _ = v20
    }
    v21 := GenS1{F0: 42, F1: 27}
    if false {
        v22 := GenS1{F0: 9223372036854775807, F1: 42}
        if len(v1) > 0 {
            v1[0] = 127
        }
        if true {
            fmt.Println(v2, v3)
            var v23 bool = false
            v23 = (true && v23)
            _ = v23
        } else {
            switch "\n\t" {
            case "\n\t":
                v3 = -1
            }
        }
        _ = v22
    }
    v3 = (v3 - (v3 % ((-1) | 1)))
    v2 = -632
    v24 := GenS2{F0: 0}
    var v25 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
