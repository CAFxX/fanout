package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint; F2 int }

func main() {
    var v1 uint8 = 32
    var v2 int8 = 74
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := make(map[string]string)
    v6 := []int{42, -357}
    if false {
        var v7 string = ("" + "a")
        _ = v7
    }
    if (!(-79 <= (^(v2)))) {
        if true {
            fmt.Println(v2, v2)
        } else {
            if (!true) {
                v8 := GenS1{F0: 818}
                var v9 int16 = -373
                _ = v8
                _ = v9
            } else {
                v10 := GenS2{F0: 2}
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 int8 = -33
                    for fi13 := 0; fi13 < 3; fi13++ {
                        switch "hello" {
                        case "":
                            var v14 float64 = 3.14159
                            _ = v14
                        }
                        v15 := GenS3{F0: 32767, F1: v1}
                        _ = fi13
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                }
                if false {
                    if (7 != v1) {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v2 -= v2
                            var v17 float32 = (float32(0.1) - (((1.5 / 3.14159) - ((0.0 - 0.0) + 0.1)) * 1e10))
                            v18 := GenS4{F0: 194, F1: 42}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                        fmt.Println(v1)
                        v1 = v1
                    } else {
                        v19 := GenS1{F0: -977}
                        _ = v19
                    }
                } else {
                    fmt.Println(v2, v2)
                }
                _ = v10
            }
        }
        var v20 float32 = 0.1
        if (false || ((-1 != -901) || true)) {
            var v21 uint16 = 1
            var v22 bool = ("" != "a")
            if v22 {
                _ = len(v4)
                v20 += -2.25
            }
            _ = v21
            _ = v22
        }
        _ = v20
    }
    var v23 int64 = 42
    v1 += v1
    var v24 int8 = v2
    fmt.Println(v1, v24, v1)
    v23 += (^(v23))
    for fi25 := 0; fi25 < 17; fi25++ {
        switch "foo bar" {
        case "foo bar":
            v24 = 127
        case "a":
            switch "foo bar" {
            case "":
                if (!false) {
                    if false {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS5{F0: (!false)}
                            _ = fi26
                            _ = v27
                        }
                        var v28 uint32 = 2147483648
                        _ = v28
                    }
                    v29 := GenS2{F0: 255}
                    var v30 int = int(v1)
                    _ = v29
                    _ = v30
                }
            }
        case "hello":
            v31 := GenS2{F0: 0}
            _ = v31
        default:
            if true {
                if false {
                    if true {
                        var v32 int = (fi25 % ((fi25) | 1))
                        _ = v32
                    } else {
                        fmt.Println(v2)
                        for fi33 := 0; fi33 < 12; fi33++ {
                            var v34 uint8 = v1
                            if true {
                                if ((653 == 65535) || false) {
                                    var v35 uint8 = (^(v1))
                                    _ = v35
                                } else {
                                    v24 -= 100
                                    delete(v5, "foo bar")
                                    fmt.Println(fi25, v23, v24)
                                }
                            } else {
                                var v36 uint16 = 255
                                v24 -= v24
                                var v37 int = fi25
                                _ = v36
                                _ = v37
                            }
                            _ = fi33
                            _ = v34
                        }
                        v24 = 42
                    }
                    v38 := GenS6{F0: -1, F1: 0, F2: int(v1)}
                    _ = v38
                }
            }
        }
        v2 = (v24 ^ (v24 << ((4294967295) & 63)))
        for fi39 := 0; fi39 < 8; fi39++ {
            if ("" != ("hello" + "hello")) {
                _ = len(v3)
                fmt.Println(v2, v1)
                v1 -= (v1 ^ v1)
            } else {
                var v40 uint32 = 4294967295
                var v41 uint8 = (v1 << ((uint(v1)) & 63))
                v42 := v4[1]
                _ = v40
                _ = v41
                _ = v42
            }
            _ = len(v6)
            _ = fi39
        }
        _ = fi25
    }
    var v43 int64 = (v23 >> ((65535) & 63))
    var v44 int = -527
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
