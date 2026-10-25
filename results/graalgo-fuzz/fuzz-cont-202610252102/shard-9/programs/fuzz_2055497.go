package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-821, 728}
    var v3 int = 32767
    var v4 int = v3
    v2 = append(v2, v4)
    var v5 uint64 = 18446744073709551615
    var v6 int32 = (v1 % (((^(v1))) | 1))
    _ = len(v2)
    var v7 int64 = 1
    if false {
        _ = len(v2)
        v2 = append(v2, (v3 * (v3 + (v4 % (((v3 ^ v4)) | 1)))))
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int16 = -32768
            v5 -= v5
            _ = fi8
            _ = v9
        }
    }
    var v10 uint64 = v5
    _ = len(v2)
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ((("\n\t" + "hello") + "hello") + "hello") {
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = (fi11 % ((-1) | 1))
                }
            case "foo bar":
                v2 = append(v2, fi12)
            case "a":
                if len(v2) > 0 {
                    v2[0] = (v3 >> ((65535) & 63))
                }
            default:
                v13 := GenS1{F0: true, F1: -32768}
                _ = v13
            }
            _ = fi12
        }
        if len(v2) > 0 {
            v2[0] = v4
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        if false {
            if (4294967295 == 2147483647) {
                v5 = (v5 ^ v5)
                var v15 int32 = -697
                v16 := GenS1{F0: ((42 > 271) && (false && (!(!(v10 <= (v10 * 18446744073709551615)))))), F1: -438}
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch (("" + "foo bar") + ("foo bar" + "")) {
                    case "\n\t":
                        _ = len(v2)
                    case "foo bar":
                        v18 := GenS2{F0: 296}
                        fmt.Println(v7, v10)
                        if (!(65535 > 577)) {
                            v10 -= v5
                        } else {
                            if false {
                                var v19 uint64 = v10
                                v20 := GenS2{F0: 42}
                                v5 = (v19 % (((v19 + v19)) | 1))
                                _ = v19
                                _ = v20
                            } else {
                                if (false || false) {
                                    var v21 uint64 = v10
                                    var v22 bool = (v21 >= v10)
                                    _ = v21
                                    _ = v22
                                } else {
                                    v7 -= (v7 * v7)
                                }
                            }
                            v5 = (v10 << ((1) & 63))
                            var v23 float64 = float64(3.14159)
                            _ = v23
                        }
                        _ = v18
                    case "a":
                        _ = len(v2)
                    default:
                        fmt.Println(v7)
                    }
                    var v24 uint8 = 177
                    v25 := GenS3{F0: -928, F1: (574 <= 4294967295)}
                    _ = fi17
                    _ = v24
                    _ = v25
                }
            }
            fmt.Println(v5, v4, v4)
        } else {
            for fi26 := 0; fi26 < 13; fi26++ {
                v3 = v3
                _ = fi26
            }
            v3 = fi14
            for fi27 := 0; fi27 < 3; fi27++ {
                fmt.Println(fi14, v1, fi27)
                v4 += -575
                _ = fi27
            }
        }
        switch v3 {
        case 127:
            switch fi14 {
            case 128:
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 uint16 = 255
                    _ = fi28
                    _ = v29
                }
            case 9223372036854775807:
                if len(v2) > 0 {
                    v2[0] = 0
                }
            }
        case 32768:
            var v30 int8 = 127
            _ = v30
        }
        _ = fi14
    }
    var v31 uint32 = 2
    v32 := GenS1{F0: false, F1: 1000}
    var v33 bool = true
    if v33 {
        v33 = v33
        fmt.Println(v5, v31, v5)
    } else {
        var v34 float64 = 0.0
        fmt.Println(v31)
        _ = v34
    }
    v35 := GenS3{F0: int64(v31), F1: false}
    if ((2 <= 42) && false) {
        for fi36 := 0; fi36 < 10; fi36++ {
            if (-128 > 84) {
                v10 = v5
                if v33 {
                    v37 := GenS3{F0: (v7 / (((v7 ^ -1)) | 1)), F1: (-104 >= 258)}
                    _ = len(v2)
                    _ = v37
                }
            }
            var v38 int = 32768
            _ = fi36
            _ = v38
        }
        for fi39 := 0; fi39 < 9; fi39++ {
            var v40 int8 = 1
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
