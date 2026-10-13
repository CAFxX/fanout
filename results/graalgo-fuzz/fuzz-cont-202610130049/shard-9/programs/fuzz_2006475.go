package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int32 = -551
    v2 := []uint32{0, 4294967295}
    var v3 string = "foo bar"
    var v4 float64 = -2.25
    fmt.Println(v4, v1)
    switch -2147483648 {
    case 127:
        v5 := GenS1{F0: 0, F1: 255}
        _ = v5
    case 32768:
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float64 = (((v4 * v4) - (0.1 + 3.14159)) + v4)
            var v8 float32 = (((0.1 * 0.0) / 0.1) / 1.0)
            v8 += (v8 + float32((v4 / ((v4 - 1e10) / (v7 + (v7 - -2.25))))))
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        var v9 int16 = 0
        _ = v9
    }
    v2 = append(v2, 602)
    switch v3 {
    case "":
        if false {
            v10 := GenS1{F0: 1, F1: 366}
            v4 = -2.25
            _ = v10
        } else {
            if true {
                v11 := GenS1{F0: 102, F1: 984}
                if (42 <= -1) {
                    fmt.Println(v4)
                    v1 -= v1
                }
                _ = v11
            }
            var v12 int = -510
            _ = v12
        }
    default:
        v4 += 0.1
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v1)
        _ = fi13
    }
    v14 := GenS1{F0: 42, F1: 25}
    switch 1 {
    case 128:
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint16 = 2
            if (1 <= 1) {
                v3 = v3
            } else {
                v4 = v4
                if false {
                    v17 := GenS2{F0: 44, F1: 1}
                    v18 := GenS2{F0: 0, F1: 0}
                    _ = v17
                    _ = v18
                } else {
                    v4 = v4
                    if (!(143 != 4294967295)) {
                        var v19 uint16 = (v16 % ((v16) | 1))
                        _ = v19
                    }
                }
            }
            v20 := GenS3{F0: 42, F1: (v16 / ((v16) | 1)), F2: 1}
            _ = fi15
            _ = v16
            _ = v20
        }
    case 802:
        if false {
            switch 2147483647 {
            case 32768:
                var v21 int64 = -1
                _ = v21
            default:
                fmt.Println(v3, v1, v1)
            }
            _ = len(v2)
            if (4294967295 <= 65535) {
                var v22 uint8 = 255
                if true {
                    var v23 int32 = v1
                    v3 = "hello"
                    _ = v23
                }
                _ = v22
            } else {
                v24 := GenS3{F0: -9223372036854775808, F1: 255, F2: 127}
                v25 := GenS3{F0: 328, F1: 11, F2: 42}
                _ = v24
                _ = v25
            }
        } else {
            var v26 int8 = -100
            _ = v26
        }
    case -1:
        switch "hello" {
        case "hello":
            for fi27 := 0; fi27 < 9; fi27++ {
                v3 = v3
                v2 = append(v2, 1)
                _ = fi27
            }
        case "a":
            for fi28 := 0; fi28 < 10; fi28++ {
                fmt.Println(v3)
                for fi29 := 0; fi29 < 15; fi29++ {
                    if ("hello" < v3) {
                        _ = len(v2)
                        if (!false) {
                            if true {
                                var v30 int8 = int8(v1)
                                _ = v30
                            }
                        }
                        v31 := GenS4{F0: -32768, F1: "\n\t"}
                        v3 = v3
                        _ = v31
                    } else {
                        var v32 int = fi29
                        var v33 string = v3
                        var v34 uint8 = 121
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    switch 32768 {
                    case 781:
                        v1 -= (v1 % (((v1 + v1)) | 1))
                    case 32767:
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                    default:
                        var v35 int64 = -790
                        _ = v35
                    }
                    v36 := GenS5{F0: (fi28 << ((255) & 63)), F1: -32768}
                    _ = fi29
                    _ = v36
                }
                v37 := GenS6{F0: 742, F1: 255}
                _ = fi28
                _ = v37
            }
        case "\n\t":
            v3 = v3
        }
    }
    var v38 int8 = -82
    v38 = v38
    _ = len(v2)
    v1 -= -1
    if false {
        fmt.Println(v3)
        if true {
            fmt.Println(v38, v4)
            v39 := GenS5{F0: 1, F1: 32767}
            v4 = v4
            _ = v39
        } else {
            switch "hello" {
            case "\n\t":
                _ = len(v2)
            case "foo bar":
                var v40 uint16 = 2
                _ = v40
            }
            v4 -= (v4 - v4)
        }
    } else {
        switch -1 {
        case -724:
            v38 += v38
        case 42:
            for fi41 := 0; fi41 < 10; fi41++ {
                for fi42 := 0; fi42 < 1; fi42++ {
                    if false {
                        var v43 bool = false
                        switch (((v3 + (v3 + v3)) + "\n\t") + v3) {
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 0
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v38 = v38
                            var v44 int64 = -364
                            _ = v44
                        }
                        _ = v43
                    }
                    _ = fi42
                }
                var v45 uint32 = 65535
                if len(v2) > 0 {
                    v2[0] = (v45 ^ (v45 - v45))
                }
                _ = fi41
                _ = v45
            }
        case 835:
            if true {
                switch 32767 {
                case 756:
                    for fi46 := 0; fi46 < 5; fi46++ {
                        for fi47 := 0; fi47 < 1; fi47++ {
                            var v48 bool = (true || true)
                            v49 := GenS2{F0: 2, F1: 255}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        v3 = v3
                        _ = fi46
                    }
                }
            }
        }
        v50 := GenS7{F0: v1}
        v4 = float64(1.5)
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
