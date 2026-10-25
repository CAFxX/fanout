package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := []uint32{825, 0, 95}
    var v2 string = "hello"
    var v3 int16 = 32767
    var v4 float64 = ((0.0 / float64(3.14159)) - ((-2.25 - -2.25) - ((1.5 * (1.5 - -2.25)) * (1e10 * -2.25))))
    var v5 uint = 2
    var v6 uint = 1
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    fmt.Println(v6)
    if len(v1) > 0 {
        v1[0] = uint32(v5)
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 825 {
    case 32768:
        v7 := GenS1{F0: 367}
        _ = v7
    case 127:
        v2 = (v2 + "foo bar")
    case 752:
        if true {
            if false {
                for fi8 := 0; fi8 < 8; fi8++ {
                    fmt.Println(fi8, v2)
                    v5 = 18446744073709551615
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                    _ = fi8
                }
                if (v5 > (v5 << ((v6) & 63))) {
                    v6 = (v6 | v6)
                } else {
                    v9 := GenS2{F0: 0, F1: -185, F2: 42}
                    _ = len(v1)
                    _ = v9
                }
            }
        } else {
            switch 128 {
            case -394:
                if (!true) {
                    var v10 uint64 = 49
                    v10 -= (v10 / ((v10) | 1))
                    _ = v10
                }
            default:
                var v11 int64 = -1
                _ = v11
            }
            var v12 int32 = -1
            var v13 int = -684
            _ = v12
            _ = v13
        }
    default:
        switch v2 {
        case "":
            if false {
                if ((154 != -1) && (!(-859 <= -173))) {
                    var v14 int32 = 423
                    _ = v14
                } else {
                    v6 -= v6
                    v1 = append(v1, 42)
                    if false {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            var v16 int8 = int8(v3)
                            _ = len(v1)
                            _ = fi15
                            _ = v16
                        }
                        var v17 int = -802
                        if (65535 >= 4294967295) {
                            for fi18 := 0; fi18 < 2; fi18++ {
                                v19 := GenS3{F0: (v3 * -32768)}
                                var v20 int32 = 1
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            for fi21 := 0; fi21 < 10; fi21++ {
                                var v22 int8 = 127
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi21
                                _ = v22
                            }
                            if (!((v17 << ((65535) & 63)) >= v17)) {
                                fmt.Println(v3, v17, v17)
                                v5 = v6
                            } else {
                                var v23 int = v17
                                _ = v23
                            }
                        } else {
                            var v24 int32 = 2147483647
                            v17 -= 127
                            v4 -= v4
                            _ = v24
                        }
                        _ = v17
                    } else {
                        var v25 uint16 = 42
                        v25 = 2
                        _ = v25
                    }
                }
                v2 = ("hello" + "foo bar")
                v3 -= 32767
            }
        case "foo bar":
            v26 := GenS4{F0: 1}
            _ = v26
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
