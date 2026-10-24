package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int }
type GenS3 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 2
    v3 := []uint32{826}
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: 65535, F1: (v2 & v2)}
    var v6 uint16 = v2
    v6 -= v6
    v7 := GenS2{F0: 1, F1: 2147483648, F2: -324}
    var v8 uint8 = 2
    var v9 uint16 = (v6 ^ 2)
    v9 += uint16(v8)
    for fi10 := 0; fi10 < 4; fi10++ {
        switch (fi10 ^ -748) {
        case 32767:
            var v11 int = fi10
            _ = v11
        case 4294967295:
            if true {
                fmt.Println(v6)
                fmt.Println(v6, v6)
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := v4[4294967295]
                    v4[1] = v13
                    v14 := GenS1{F0: 0, F1: (v6 * (v9 - v2))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        }
        _ = fi10
    }
    fmt.Println(v9)
    if (1 > 18446744073709551615) {
        if true {
            var v15 bool = (true || true)
            _ = v15
        }
        if true {
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(fi16, v8)
                v17 := GenS3{F0: 100, F1: 232, F2: 0}
                fmt.Println(v8, v9, v6)
                _ = fi16
                _ = v17
            }
            v2 += (v2 % ((v9) | 1))
            switch 32767 {
            case 2147483647:
                if false {
                    if false {
                        v6 = 2
                        v2 = v6
                    }
                } else {
                    for fi18 := 0; fi18 < 17; fi18++ {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            _ = len(v4)
                            v20 := v4[0]
                            v8 = uint8(v6)
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
            case -1:
                fmt.Println(v9)
            case 614:
                v6 = 32767
            default:
                fmt.Println(v2)
            }
        } else {
            var v21 int64 = 914
            fmt.Println(v9, v8)
            _ = v21
        }
        if (v9 > 1) {
            v22 := GenS2{F0: 18446744073709551615, F1: 2147483647, F2: 1}
            _ = v22
        } else {
            fmt.Println(v9, v9)
            fmt.Println(v2, v2, v2)
        }
    }
    var v23 int16 = 42
    if false {
        var v24 uint64 = uint64(v8)
        var v25 int = -2147483648
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v9, v6)
        }
        _ = v24
        _ = v25
    } else {
        if true {
            var v26 float64 = (0.1 + float64(0.1))
            switch (("\n\t" + ("\n\t" + "")) + "\n\t") {
            case "":
                var v27 uint = 76
                _ = v27
            case "\n\t":
                var v28 int64 = 9223372036854775807
                _ = v28
            case "a":
                if ((v1 ^ 0) <= (v1 | v1)) {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        switch -287 {
                        case 42:
                            v6 += 0
                        case -786:
                            v1 = int32(v8)
                        }
                        fmt.Println(v9, v6, v23)
                        _ = fi29
                    }
                    var v30 uint = 83
                    var v31 int32 = v1
                    _ = v30
                    _ = v31
                } else {
                    if ((v26 * v26) < 0.0) {
                        var v32 int = 128
                        _ = v32
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 376
                        }
                        if false {
                            v1 = v1
                        } else {
                            if true {
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    fmt.Println(v2, fi33)
                                    var v34 float64 = v26
                                    _ = fi33
                                    _ = v34
                                }
                                var v35 uint32 = 2147483648
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v35
                            }
                            v36 := GenS3{F0: 59, F1: 65535, F2: v23}
                            _ = v36
                        }
                    }
                    delete(v4, 778)
                }
            }
            _ = v26
        } else {
            for fi37 := 0; fi37 < 8; fi37++ {
                var v38 float32 = 1e10
                if true {
                    delete(v4, 2147483647)
                    v39 := GenS1{F0: 1, F1: v9}
                    fmt.Println(v2, fi37)
                    _ = v39
                }
                _ = fi37
                _ = v38
            }
            _ = len(v4)
        }
    }
    var v40 float64 = float64(3.14159)
    var v41 uint16 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
