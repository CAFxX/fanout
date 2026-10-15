package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    var v1 uint8 = 128
    var v2 bool = (42 <= 65535)
    var v3 uint16 = 32767
    var v4 float32 = 1e10
    var v5 int32 = 0
    var v6 string = ("a" + "a")
    v4 += 0.0
    v5 = 42
    v2 = ((v2 && (165 > 18446744073709551615)) || (v2 || (!false)))
    v7 := []int{4294967295, 9223372036854775807, 1, -273}
    v3 -= (v3 & v3)
    fmt.Println(v5, v2, v1)
    if (!true) {
        if v2 {
            var v8 string = ((v6 + v6) + v6)
            var v9 float64 = 0.1
            _ = v8
            _ = v9
        } else {
            var v10 int = -471
            v10 -= -2147483648
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            switch v6 {
            case "a":
                var v12 uint64 = 42
                _ = v12
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                v4 += (float32(1e10) * 1e10)
                _ = fi13
            }
            _ = fi11
        }
    }
    v6 = "foo bar"
    switch ("hello" + v6) {
    case "":
        if (2 < 42) {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint8 = v1
                _ = fi14
                _ = v15
            }
            fmt.Println(v5, v2, v6)
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int16 = 0
                _ = fi16
                _ = v17
            }
        }
    case "hello":
        for fi18 := 0; fi18 < 18; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(v3)
                v20 := GenS1{F0: v3, F1: (0.1 < float64(float32(0.0)))}
                _ = fi19
                _ = v20
            }
            var v21 uint64 = 488
            if v2 {
                v3 = v3
            }
            _ = fi18
            _ = v21
        }
    default:
        if (-437 == -1) {
            var v22 string = v6
            for fi23 := 0; fi23 < 12; fi23++ {
                if (-827 < 42) {
                    fmt.Println(v6)
                    switch ((v22 + "a") + (v6 + v22)) {
                    case "":
                        v4 = v4
                    case "a":
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if (!v2) {
                                v3 -= (^(v3))
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = fi24
                                }
                                var v25 uint16 = v3
                                v26 := GenS1{F0: 0, F1: v2}
                                _ = v25
                                _ = v26
                            }
                            var v27 uint64 = 18446744073709551615
                            v22 = (v6 + v6)
                            var v28 string = (v6 + (v6 + v22))
                            var v29 uint = 896
                            if true {
                                switch fi23 {
                                case 406:
                                    v30 := GenS2{F0: int(v3), F1: "a"}
                                    _ = v30
                                }
                                if (2147483647 < -659) {
                                    v31 := GenS3{F0: (v27 | (v27 << ((2) & 63)))}
                                    switch "" {
                                    case "foo bar":
                                        v4 = v4
                                    default:
                                        if (1000 > 1000) {
                                            var v32 float64 = ((1.5 / 3.14159) - 0.1)
                                            var v33 int32 = v5
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                    _ = v31
                                }
                            }
                            var v34 int64 = 1
                            _ = fi24
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v34
                        }
                    case "hello":
                        if (true && v2) {
                            switch "" {
                            case "hello":
                                switch -818 {
                                case 127:
                                    var v35 uint8 = (v1 << ((42) & 63))
                                    _ = v35
                                case -2147483648:
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                            }
                            for fi36 := 0; fi36 < 2; fi36++ {
                                v37 := GenS4{F0: 18446744073709551615, F1: 686}
                                var v38 uint32 = 2147483647
                                var v39 int32 = (v5 % ((-1) | 1))
                                v40 := GenS3{F0: 4294967295}
                                v6 = "a"
                                var v41 uint8 = 128
                                v6 = "hello"
                                _ = fi36
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        } else {
                            switch -2147483648 {
                            case 128:
                                var v42 uint = 42
                                _ = v42
                            }
                        }
                    }
                }
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 float64 = 0.0
                    _ = fi43
                    _ = v44
                }
                _ = fi23
            }
            _ = v22
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v6 = "hello"
            fmt.Println(v3, v6)
        }
    }
    switch v6 {
    case "foo bar":
        var v45 uint64 = 0
        _ = v45
    default:
        var v46 uint = 326
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
