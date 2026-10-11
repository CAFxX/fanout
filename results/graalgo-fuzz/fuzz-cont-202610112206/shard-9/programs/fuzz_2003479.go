package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string; F1 bool; F2 int16 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 float64 = (1.5 / ((((0.0 + 0.0) * 0.1) - 0.1) + 1e10))
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v6 int32 = -2147483648
    var v7 uint8 = 42
    var v8 uint64 = 1
    if (true || true) {
        if (false && false) {
            v9 := GenS1{F0: 9223372036854775807, F1: 32768}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 17; fi10++ {
                if false {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if (1 != 65535) {
                            v3 = (1.5 - ((v3 / v3) * ((v3 / float64((1e10 + 0.0))) - (v3 - float64(0.0)))))
                            var v12 uint = 65535
                            var v13 int64 = 651
                            if (1 > 42) {
                                v3 = 0.0
                            } else {
                                if false {
                                    v3 += -2.25
                                } else {
                                    if false {
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            var v14 float32 = 0.0
                                            _ = v14
                                        }
                                        fmt.Println(v13)
                                        v6 += v6
                                    } else {
                                        v1 -= int8(v7)
                                        var v15 int = int(v8)
                                        _ = v15
                                    }
                                    var v16 uint16 = uint16(v7)
                                    if (42 == -449) {
                                        switch fi11 {
                                        case -148:
                                            fmt.Println(fi11)
                                        }
                                    }
                                    _ = v16
                                }
                                var v17 bool = true
                                v12 -= v12
                                _ = v17
                            }
                            var v18 uint64 = v8
                            fmt.Println(v12)
                            _ = v12
                            _ = v13
                            _ = v18
                        } else {
                            fmt.Println(v6)
                        }
                        fmt.Println(v7, v7, fi10)
                        _ = fi11
                    }
                    for fi19 := 0; fi19 < 10; fi19++ {
                        fmt.Println(v8, v7, fi19)
                        _ = fi19
                    }
                    var v20 int = (fi10 ^ (-(fi10)))
                    _ = v20
                } else {
                    v3 = v3
                    var v21 uint8 = 118
                    _ = v21
                }
                switch fi10 {
                case -9223372036854775808:
                    var v22 uint64 = 408
                    _ = v22
                case 123:
                    if (!false) {
                        v23 := []int{(fi10 << ((42) & 63))}
                        switch (fi10 & 4294967295) {
                        case 32767:
                            var v24 int32 = (^(v6))
                            _ = v24
                        case 257:
                            if false {
                                v6 -= (v6 | -719)
                                var v25 uint8 = (v7 << ((42) & 63))
                                _ = v25
                            }
                        case 190:
                            if (-820 >= 42) {
                                v23 = append(v23, 397)
                            } else {
                                v26 := GenS1{F0: 903, F1: fi10}
                                _ = v26
                            }
                        }
                        _ = v23
                    } else {
                        if false {
                            v1 += (v1 / ((v1) | 1))
                        } else {
                            switch "a" {
                            case "foo bar":
                                var v27 uint8 = 127
                                _ = v27
                            case "":
                                fmt.Println(fi10)
                            }
                        }
                    }
                }
                _ = fi10
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 int64 = 0
                _ = fi28
                _ = v29
            }
        }
    } else {
        v6 -= -1
        v30 := GenS1{F0: 446, F1: 744}
        _ = v30
    }
    v7 = (^(v7))
    if false {
        v8 = v8
        for fi31 := 0; fi31 < 2; fi31++ {
            v6 = (v6 << ((1) & 63))
            _ = fi31
        }
    } else {
        if (9223372036854775807 >= -9223372036854775808) {
            var v32 int8 = v1
            _ = v32
        } else {
            var v33 int16 = 1
            fmt.Println(v1)
            _ = v33
        }
    }
    var v34 float64 = float64(0.1)
    if false {
        fmt.Println(v6, v1, v1)
    } else {
        v8 = v8
    }
    v35 := GenS2{F0: "a", F1: (((-2.25 - v34) * 1e10) >= 1.5), F2: 920}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
